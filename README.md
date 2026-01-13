# lien

Simple HTTP static server.

![Screenshot](https://github.com/anasrar/lien/assets/38805204/6a0246cc-92cf-45dd-ad08-46480c2c6f95)

## Development

### Setup

```bash
# setup svelte
cd internal/webui
npm install
npm run build # need build directory to run go server

# go server
cd ../..
go run cmd/server/main.go

# svelte as front end and proxy
cd internal/webui
npm run dev
```

### Build

```bash
cd internal/webui
npm run build
# build static binary for linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o lien -trimpath -ldflags="-s -w" cmd/server/main.go
```

## Built With

- [Go](https://go.dev/)
- [Fiber](https://gofiber.io/)
- [HUMA](https://huma.rocks/)
- [Svelte](https://svelte.dev/)
- [shadcn-svelte](https://www.shadcn-svelte.com/)
- [OpenAPI TypeScript](https://openapi-ts.dev/)
- [@tanstack/svelte-query](https://tanstack.com/query/latest)
- [@tanstack/table-core](https://tanstack.com/table/latest)
- [lucide](https://lucide.dev/)
