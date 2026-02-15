# Blog Aggregator

A simple RSS feed aggregator CLI built with Go and PostgreSQL.

## Features

- Register and log in users
- Add new RSS feeds
- Follow / unfollow feeds
- List all feeds and followed feeds
- Periodically fetch feeds with `agg`


## Project Structure

```text
.
├── README.md
├── commands.go
├── go.mod
├── go.sum
├── handler_agg.go
├── handler_feed.go
├── handler_feed_following.go
├── handler_feed_follows.go
├── handler_feeds.go
├── handler_reset.go
├── handler_unfollow.go
├── handler_user.go
├── internal
│   ├── config
│   │   └── config.go
│   └── database
│       ├── db.go
│       ├── feed_follows.sql.go
│       ├── feeds.sql.go
│       ├── fetch_feed.sql.go
│       ├── models.go
│       └── users.sql.go
├── main.go
├── rss_feed.go
├── sql
│   ├── queries
│   │   ├── feed_follows.sql
│   │   ├── feeds.sql
│   │   ├── fetch_feed.sql
│   │   └── users.sql
│   └── schema
│       ├── 001_users.sql
│       ├── 002_feeds.sql
│       ├── 003_feed_follows.sql
│       └── 004_feed_lastfetched.sql
└── sqlc.yaml
```

## Usage

Run commands with:

```bash
go run . <command> [arguments]
```

Example flow:

```bash
go run . register kazim
go run . login kazim
go run . addfeed "Hacker News" "https://hnrss.org/frontpage"
go run . feeds
go run . follow "https://hnrss.org/frontpage"
go run . following
go run . agg 30s
```

## Commands

- `register <name>`
- `login <name>`
- `users`
- `reset`
- `addfeed <name> <url>`
- `feeds`
- `follow <url>`
- `following`
- `unfollow <url>`
- `agg <time_between_reqs>` (for example: `10s`, `1m`)

## Notes

- `agg` runs in an infinite loop; stop it with `Ctrl+C`.
- `addfeed` automatically follows the feed after creating it.

## What I Learned

- Building a CLI app in Go with command routing and handler functions.
- Parsing command-line arguments and validating command usage.
- Using a middleware-like pattern (`middlewareLoggedIn`) to protect commands that require authentication.
- Working with PostgreSQL relations (`users`, `feeds`, `feed_follows`) and unique constraints.
- Managing schema changes with SQL migrations (`goose`).
- Generating type-safe database access code using `sqlc`.
- Using `context.Context` in database and HTTP operations.
- Fetching and parsing RSS/XML feeds with Go's `net/http` and `encoding/xml`.
- Handling escaped HTML entities in feed content.
- Running periodic background work with `time.Ticker` and `time.ParseDuration`.
