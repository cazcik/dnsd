# node build
FROM node:18-alpine AS base

FROM base AS deps

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY package.json ./

RUN npm i

FROM base AS node-builder

WORKDIR /app

COPY --from=deps /app/node_modules ./node_modules

COPY . .

RUN npm run build

# golang build
FROM golang:alpine as go-builder

WORKDIR /app 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

# final
FROM scratch

WORKDIR /app

COPY --from=node-builder /app/dist ./dist
COPY --from=go-builder /app/dnsd .

EXPOSE 8080

ENTRYPOINT ["/app/dnsd"]
