#####################
# BUILD ENVIRONMENT #
#####################

FROM golang:1.21-alpine AS build

WORKDIR /build

COPY ./src/ .

RUN set -xe && \
    apk add --no-cache ca-certificates && \
    go get -d -v && \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -installsuffix cgo -o chatgpt-discord-bot .

#####################
# FINAL ENVIRONMENT #
#####################

FROM scratch

LABEL maintainer="DJΞRFY <djerfy@gmail.com>" \
      description="ChatGPT Discord Bot" \
      repository="https://github.com/djerfy/chatgpt-discord-bot.git"

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /build/chatgpt-discord-bot /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/chatgpt-discord-bot"]
