# Short Link

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
