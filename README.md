# Go Upload Files

This project aims to provide a simple solution to upload files to a server. You can set the target folder, and a list of users allowed to upload files.

It is built in Golang and Gin to provide great performance and be as light as possible and can be deployed using Docker or directly on your server.

## 🛠️ Tech Stack

- [Docker](https://www.docker.com/)
- [Docker Compose v2](https://docs.docker.com/compose/compose-v2/)
- [Golang](https://go.dev/)
- [Go Gin](https://github.com/gin-gonic/gin)
- [Air](https://github.com/cosmtrek/air) (live reload)

## 🛠️ Install Dependencies

```bash
go mod download
```

## 🧑🏻‍💻 Usage

To start the project, you first need configure your `.env` file. You can use the `.env.dist` file as a template:

```bash
cp .env.dist .env
```

Then you can launch your project simply using air command:

```bash
air
```

You can also use docker compose to launch the project:

```bash
cp docker-compose.yml.dist docker-compose.yml
docker-compose up -d --build && docker-compose logs -f
```

> You can now reach your application instance on [http://localhost:8118](http://localhost:8118) 

## 🔧 Configuration

### Environment variables

| Name                | Description                                                               | Default value |
| ------------------- | ------------------------------------------------------------------------- | ------------- |
| `TARGET_FOLDER`     | Target folder where the files will be uploaded. Must end with an `/`      | `./upload/`   |
| `AUTHORIZED_USERID` | List of users allowed to upload files. Must be separated by a `,`         | `user01`      |
| `GOTIFY_URL`        | You can set a gotify url to send notifications when a file is uploaded.   | ``            |
| `GOTIFY_TOKEN`      | You can set a gotify token to send notifications when a file is uploaded. | ``            |
| `GIN_MODE`          | Gin mode. Can be `debug` or `release`.                                    | `debug`       |



> When using docker, you must set the environment variables in the `docker-compose-override.yml` file. If you want to use another folder, you must also update the `volumes` part to target the new folder: `- ./my_target_folder:/upload`


## [License](LICENSE)

## [Contributing](CONTRIBUTING.md)

## [Code of conduct](CODE_OF_CONDUCT.md)
