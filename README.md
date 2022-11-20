# Read Track

Personal reading list inspired by <https://www.tdpain.net/blog/a-year-of-reading>.

## Features

- Simple API Endpoints.
- Admin Pages for Statistics and Token Management.
- Easy access for Static Site Generation.

## Dev

1. Setup `.env` according to `.env.example`. Make sure your `HTTP_BASIC_AUTH` and `DB_NAME` are correctly configured.
2. Run
   ```
   script/run.sh
   ```

## Deploy

### Deploy with fly.io

1. Install fly.io cli
    ```shell
    brew install flyctl
    ```

2. Setup ENV with `fly.toml` and `fly secrets`
    ```shell
    fly secrets set HTTP_BASIC_AUTH="your-secrets"
    ```

3. Deploy
    ```shell
    fly deploy
    ```

## License

MIT
