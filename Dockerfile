#####################
# BUILD ENVIRONMENT #
#####################

FROM golang:1.21-alpine AS build

WORKDIR /build

COPY ./src/ .

RUN set -xe && \
    go get -d -v && \
    go build -o chatgpt-discord-bot

#####################
# FINAL ENVIRONMENT #
#####################

FROM scratch

LABEL maintainer="DJΞRFY <djerfy@gmail.com>" \
      description="ChatGPT Discord Bot" \
      repository="https://github.com/djerfy/chatgpt-discord-bot.git"

COPY --from=build /build/chatgpt-discord-bot /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/chatgpt-discord-bot"]
