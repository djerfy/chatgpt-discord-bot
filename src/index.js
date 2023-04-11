/**
 * ChatGPT Discord Bot
 */

require('dotenv/config');

const { Client, IntentsBitField } = require('discord.js');
const { Configuration, OpenAIApi } = require('openai');

const client = new Client({
    intents: [
        IntentsBitField.Flags.Guilds,
        IntentsBitField.Flags.GuildMessages,
        IntentsBitField.Flags.MessageContent
    ]
});

const configuration = new Configuration({
    apiKey: process.env.OPENAI_API_KEY
});

const openai = new OpenAIApi(configuration);

client.on('ready', () => {
    console.log(`Successfully, ChatGPT bot is ready and online`);
});

client.on('messageCreate', async (message) => {
    if (message.author.bot) return;
    if (message.channel.id !== process.env.DISCORD_CHANNEL_ID) return;
    if (message.content.startsWith('!') || message.content.startsWith('/')) return;

    let conversationLog = [{ role: 'system', content: 'You are a friendly chatbot' }];

    try {
        await message.channel.sendTyping();

        let prevMessages = await message.channel.messages.fetch({ limit: 15 });
        prevMessages.reverse();

        prevMessages.forEach((msg) => {
            if (message.content.startsWith('!') || message.content.startsWith('/')) return;
            if (msg.author.id !== client.user.id && message.author.bot) return;
            if (msg.author.id !== message.author.id) return;

            conversationLog.push({
                role: 'user',
                content: msg.content
            });
        });

        if (message.member.roles.cache.has(process.env.DISCORD_ROLE_REQUIRED)) {
            const result = await openai.createChatCompletion({
                model: 'gpt-3.5-turbo',
                messages: conversationLog
            }).catch((error) => {
                console.log(`[error] openai: ${error}`);
            });
            message.reply(result.data.choices[0].message);
        } else {
            message.reply(`Sorry, you don't have **openai** role`);
            console.log(`[warning] user '${message.author.username}' (${message.author.id}) doesn't have 'openai' role`);
        }

    } catch (error) {
        console.log(`[error] discord: ${error}`);
    }
});

client.login(process.env.DISCORD_BOT_TOKEN);
