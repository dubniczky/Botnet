# Commander Botnet

A botnet system written in Go capable of running any commands

## Usage

Generate keypair

```bash
./generate.sh
```

Run commander server

```bash
make commander
```

Compile bot

> Note: You need to change the commander address in `bot/main.go` to your own

```bash
cd bot
go build -o bot main.go
```
