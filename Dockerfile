FROM golang:alpine

# adding gcc to run test cases
RUN apk add build-base

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .
COPY ./scripts/entrypoint.sh .

COPY .env .

# Expose port
EXPOSE ${APP_PORT}

ENTRYPOINT ["/app/scripts/entrypoint.sh"]