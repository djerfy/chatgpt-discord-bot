# ChatGPT Discord Bot

This is simple ChatGPT Discord Bot built using `gpt5` model of OpenAI.

## How to run

You need to have 4 variables in your environment:

* `DISCORD_CHANNEL_ID`: the Discord channel where the bot will do the answers
* `DISCORD_BOT_TOKEN`: the Discord bot token
* `DISCORD_ROLE_REQUIRED`: require role (id) to ask questions and get answers
* `OPENAI_API_KEY`: OpenAI token [link](https://platform.openai.com/account/api-keys)

### Manually

```
$ cd ./src/
$ go run .
```

> Don't forget to edit `.env` with your configuration

### Docker

```text
$ docker run --rm --name chatgpt-discord-bot \
    -e DISCORD_CHANNEL_ID="" \
    -e DISCORD_BOT_TOKEN="" \
    -e DISCORD_ROLE_REQUIRED="" \
    -e OPENAI_API_KEY="" \
    ghcr.io/djerfy/chatgpt-discord-bot
```
