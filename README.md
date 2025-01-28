# Short Link

The best practice is used DynamoDB and Redis.
And to count request call of urls we can use message broker to count and increment counter.
Then We can add a command to store in redis popular link based on call parameter.

## Install docker command
```bash
  cp .env.example .env
```
```bash
  docker compose up -d
```

## Migrations
```bash
  go run cmd/migration/main.go up
```

## Swagger
```bash
  swag init -g ./cmd/shortlinkserver/main.go
```
- [Swagger Address](http://localhost:2525/swagger/index.html)
- Username: admin
- Password: admin
