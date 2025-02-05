##################################################
## BASE
##################################################
FROM golang:1.23.4-bookworm AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

##################################################
## DEVELOPMENT
##################################################
FROM base AS development
COPY . .
