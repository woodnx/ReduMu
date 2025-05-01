FROM mcr.microsoft.com/vscode/devcontainers/base:debian

RUN apt update && apt install -y build-essential

# Install go-task
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d
RUN curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate /usr/local/bin/

WORKDIR /ReduMu