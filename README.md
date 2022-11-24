# Read Track

[![build](https://github.com/crispgm/read-track/actions/workflows/ci.yml/badge.svg)](https://github.com/crispgm/read-track/actions/workflows/ci.yml)

Personal reading list inspired by [blog post from Thomas Pain](https://www.tdpain.net/blog/a-year-of-reading).

You may deploy your own Read Track instance and track read from different devices.

## Features

- Simple API Endpoints with easy-to-setup bookmarklet and iOS shortcut.
- Multiple verb support: read, skim, unread, and skip.
- HTTP Basic Auth single user system.
- Admin Pages for Statistics.

## Dev

1. Setup `.env` according to `.env.example`. Make sure your `HTTP_BASIC_AUTH` and `DB_NAME` are correctly configured.
2. Run (DB is auto migrated):
   ```
   script/run.sh
   ```

## Deploy

We only support fly.io right now. And DB is stored in fly volumes.

### Deploy with fly.io

1. Install fly.io cli:
    ```shell
    brew install flyctl
    ```

2. Setup `env` inside `fly.toml`:
    ```shell
    [env]
      INSTANCE = "<instance-name>"
      PORT = "8080"
      MODE = "production"
      TIMEZONE = "Asia/Shanghai"
      HTTP_PORT = ":8080"
      DB_PROVIDER = "sqlite"
      DB_NAME = "/data/read-track-production.db"
    ```

3. Setup secrets with `fly secrets`:
    ```shell
    fly secrets set HTTP_BASIC_AUTH="your-secrets"
    ```

3. Deploy:
    ```shell
    fly deploy
    ```

## License

MIT
