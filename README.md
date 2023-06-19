# Go URL Shortener

This is a URL shortener project built with Go lang. It allows you to create short and manageable URLs for your long URLs.

### Run the Project locally using docker

#### With Docker
1. Install [Docker Desktop](https://www.docker.com/products/docker-desktop/)
2. Test the installation by running:
   ```sh
   docker-compose --version
   # Docker Compose version v2.6.1
   ```

##### Steps
1. Make a copy of the `.env.example` file and save it as `.env`:
   ```sh
   cp .env.example .env
   ```
2. Setup the project with single command installation via docker:
   ```sh
   docker-compose up --build -d

### Setting Up the Project locally
Clone the repository to your local machine:

```
sh git clone https://github.com/your-username/go-url-shortener.git
```

#### Requirements
- Golang
- PostgreSQL

#### Pre-requisites
1. Download [Golang](https://go.dev/)
2. Test the installation by running:
   ```sh
   go version
   # go version go1.19.2 darwin/arm64
   ```
#### Steps
1. Navigate to the right directory where your project will be locally saved
2. Clone this repository and move to `go-url-shortener` directory
   ```sh
   cd go-url-shortener/
   ```
3. Install go dependencies
   ```sh
   go get
   ```

4. Install [goose](https://pressly.github.io/goose/) library for running migrations

5. Make a copy of the `.env.example` file and save it as `.env`:
   ```sh
   cp .env.example .env
   ```
6. Update the `.env` file with relevant configurations.

7. Run migrations
   ```sh
   make migration-up-postgres
   ```

8. Start the server
   ```sh
   make server
   ```
