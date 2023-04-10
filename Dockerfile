############################
# Docker build environment #
############################

FROM node:19-buster-slim AS build

WORKDIR /build

COPY . .

RUN npm i

############################
# Docker final environment #
############################

FROM node:19-buster-slim

LABEL maintainer="DJΞRFY <djerfy@gmail.com>" \
      description="ChatGPT Discord Bot" \
      repository="https://github.com/djerfy/chatgpt-discord-bot.git"

WORKDIR /app

COPY --from=build /build/node_modules .
COPY . .

ENTRYPOINT ["npm", "run", "start"]

