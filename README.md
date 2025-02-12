# Fresh Products Api

```sh
▗▄▄▄▖▗▄▄▖ ▗▄▄▄▖ ▗▄▄▖ ▗▄▄▖ ▗▄▖  ▗▄▄▖     ▗▄▖ ▗▄▄▖▗▄▄▄▖
▐▌   ▐▌ ▐▌▐▌   ▐▌   ▐▌   ▐▌ ▐▌▐▌       ▐▌ ▐▌▐▌ ▐▌ █  
▐▛▀▀▘▐▛▀▚▖▐▛▀▀▘ ▝▀▚▖▐▌   ▐▌ ▐▌ ▝▀▚▖    ▐▛▀▜▌▐▛▀▘  █  
▐▌   ▐▌ ▐▌▐▙▄▄▖▗▄▄▞▘▝▚▄▄▖▝▚▄▞▘▗▄▄▞▘    ▐▌ ▐▌▐▌  ▗▄█▄▖
```

Welcome to `Frescos Api`!
This is a [chi](https://github.com/go-chi/chi) powered webserver using `Go`, `MYSQL`.

## Swagger

This project uses Swagger documentation powered by [goswag](https://github.com/swaggo/swag).

To explore the app's endpoints, simply raise the server and navigate to `http://localhost:8080/swagger/index.html#/`.

## How to run in development

For development, this project uses `docker compose` and some scripts in a `Makefile`.

Inside the docker container, Hot reloading is enabled because of the use of [air-verse](https://github.com/air-verse/air).

### Commands

#### Up

Start the Docker containers in detached mode:

```bash
make up
```

#### Down

Stop and remove the Docker containers.

```bash
make down
```

#### MySQL

Access the MySQL database running in the frescos-db container.

```bash
make mysql
```

#### Swag

Initialize or updates Swagger documentation for the API.

```bash
make swag
```

#### Lint

Run the [linter](https://github.com/golangci/golangci-lint) on the Go code.

```bash
make lint
```

```bash
make lint-fix // auto apply linter suggestions
```

#### Test

Run the Go tests and format the output with [gotestsum](https://github.com/gotestyourself/gotestsum).

```bash
make test
```

Run the tests with coverage and generate a coverage report, excluding ignored files.

```bash
make test-cover
```

Get the average coverage from the coverage report.

```bash
make cover-avg
```
