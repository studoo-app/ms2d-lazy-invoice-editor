FROM golang:latest

WORKDIR /app

RUN go mod init pwdGen
RUN go get -u github.com/spf13/cobra@latest

RUN go install github.com/spf13/cobra-cli@latest

EXPOSE 8080



