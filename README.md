# Read Track

[![build](https://github.com/crispgm/read-track/actions/workflows/ci.yml/badge.svg)](https://github.com/crispgm/read-track/actions/workflows/ci.yml)

A minimal personal reading list inspired by [blog post from Thomas Pain](https://www.tdpain.net/blog/a-year-of-reading).

You may deploy your own Read Track instance and track read from different devices.

## Features

- Simple API Endpoints with easy-to-setup bookmarklet and iOS shortcut.
- Multiple verb support: read, skim, unread, and skip.
- Designated for single user with Auth0 login.
- Dashboard with statistical data.

## Dev

1. Register an Auth0 account and create an application.
2. Setup `.env` according to `.env.example`. Make sure all the env vars are ready.
3. Run (DB is auto migrated):
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
    fly secrets set HTTP_TOKEN="your-secrets"
    fly secrets set AUTH0_DOMAIN='YOUR_DOMAIN'
    fly secrets set AUTH0_CLIENT_ID='YOUR_CLIENT_ID'
    fly secrets set AUTH0_CLIENT_SECRET='YOUR_CLIENT_SECRET'
    fly secrets set AUTH0_CALLBACK_URL='http://YOUR_DOMAIN/callback'
    fly secrets set AUTH0_USER_ID='YOUR_USER_ID'
    ```

3. Deploy:
    ```shell
    fly deploy
    ```

## License

MIT
