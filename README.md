# rio

A minimal subset of a Redis server written in Go. Speaks the RESP protocol and listens on port `6379`.

## Run

```sh
go run .
```

Then connect with any Redis client:

```sh
redis-cli
```

## Supported commands

- `PING [msg]`
- `SET key value` / `GET key`
- `HSET hash key value` / `HGET hash key` / `HGETALL hash`

## Layout

- `main.go` — TCP server loop and command dispatch
- `resp.go` — RESP protocol parser
- `value.go` — `Value` type and marshalling
- `writer.go` — RESP response writer
- `handler.go` — command implementations
