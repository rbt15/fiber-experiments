A simple Go/Fiber API to test the [Fiber](https://gofiber.io/) framework. </br>
Hosted on [Fly.io](https://fly.io/). </br>
PostgreSQL database hosted on [Neon](https://neon.tech/). </br>

Demo: https://fiber-test.fly.dev </br>

### Running locally
```bash
DATABASE_URL="user=postgres ..." go run main.go
```

> [!NOTE]
> If you want to run AutoMigrate, set the `MIGRATE` environment variable to `true`.

### Deploying to Fly.io
First, launch the project:
```bash
fly launch
```

Then, take the database url and create fly secrets:
```bash
fly secrets set DATABASE_URL="user=postgres ..."
```

Then, deploy the app:
```bash
fly deploy
```
