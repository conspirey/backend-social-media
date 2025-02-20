FROM golang:1.19 as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM node:20 AS frontend-build
WORKDIR /app/frontend
COPY ./frontend ./
RUN npm install && npm run build



FROM ubuntu:latest
LABEL authors="rihar"
WORKDIR /app

COPY .env /app/.env

RUN apt update && apt install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the Go binary from the build stage
COPY --from=build /app/main /app/main

# Copy the frontend build from the frontend build stage
COPY --from=frontend-build /app/frontend/dist /app/frontend/dist

ENV PORT=313

# Expose the port
EXPOSE $PORT
CMD ["sh", "-c", "./main"]

