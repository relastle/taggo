FROM golang:alpine
LABEL maintainer "Hiroki Konishi <relastle@gmail.com>"

# Install zsh
RUN apk update && \
    apk add --no-cache zsh git && \
    rm -f /tmp/* /etc/apk/cache/*
ENV SHELL /bin/zsh

# Install fzf
RUN go get -u github.com/junegunn/fzf
RUN go get -u github.com/fatih/color

# Install taggo
WORKDIR /go/src/github.com/relastle/taggo
COPY . .
RUN go install


