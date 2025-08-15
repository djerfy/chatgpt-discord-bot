package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	openai "github.com/sashabaranov/go-openai"
)

var (
	dg *discordgo.Session
)

func main() {
	if os.Getenv("DISCORD_BOT_TOKEN") == "" || os.Getenv("OPENAI_API_KEY") == "" || os.Getenv("DISCORD_CHANNEL_ID") == "" || os.Getenv("DISCORD_ROLE_REQUIRED") == "" {
		log.Fatal("[error] please set DISCORD_BOT_TOKEN, OPENAI_API_KEY, DISCORD_CHANNEL_ID, and DISCORD_ROLE_REQUIRED environment variables.")
		return
	}

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("[error] creating Discord session: %v", err)
		return
	}

	dg.AddHandler(newMessage)

	err = dg.Open()
	if err != nil {
		log.Fatalf("[error] opening Discord session: %v", err)
		return
	}

	fmt.Println("[info] bot connected and running...")

	defer dg.Close()

	select {}
}

func newMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot || m.ChannelID != os.Getenv("DISCORD_CHANNEL_ID") {
		return
	}

	conversationLog := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: "You are a friendly chatbot",
		},
	}

	messages, err := s.ChannelMessages(m.ChannelID, 15, m.ID, "", "")
	if err != nil {
		log.Printf("[error] fetching messages: %v", err)
		return
	}

	for _, msg := range messages {
		if strings.HasPrefix(msg.Content, "!") || strings.HasPrefix(msg.Content, "/") {
			continue
		}

		if msg.Author.ID != s.State.User.ID && msg.Author.Bot {
			continue
		}

		if msg.Author.ID != m.Author.ID {
			continue
		}

		conversationLog = append(conversationLog, openai.ChatCompletionMessage{
			Role:    "user",
			Content: msg.Content,
		})
	}

	hasRole := false
	for _, role := range m.Member.Roles {
		if role == os.Getenv("DISCORD_ROLE_REQUIRED") {
			hasRole = true
			break
		}
	}

	if !hasRole {
		s.ChannelMessageSend(m.ChannelID, "Sorry, you don't have the required role.")
		log.Printf("[warning] user '%s' (%s) doesn't have the required role", m.Author.Username, m.Author.ID)
		return
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT5,
			Messages: conversationLog,
		},
	)

	if err != nil {
		log.Printf("[error] openai: %v", err)
		return
	}

	s.ChannelMessageSend(m.ChannelID, resp.Choices[0].Message.Content)
}
