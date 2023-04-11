FROM node:19-buster-slim

LABEL maintainer="DJΞRFY <djerfy@gmail.com>" \
      description="ChatGPT Discord Bot" \
      repository="https://github.com/djerfy/chatgpt-discord-bot.git"

WORKDIR /app

COPY ./src/ .

RUN npm install

ENTRYPOINT ["npm", "run", "start"]
