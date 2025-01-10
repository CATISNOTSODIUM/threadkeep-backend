# ThreadKeep ⬢  - Your personal archive for online conversations.
![Go](https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge)
![Prisma](https://img.shields.io/badge/Prisma-3982CE?style=for-the-badge&logo=Prisma&logoColor=white)

This is the Go backend for ThreadKeep ⬢. For more information, please refer to this [repository](https://github.com/CATISNOTSODIUM/threadkeep-frontend).

- [ThreadKeep ⬢  - Your personal archive for online conversations.](#threadkeep-----your-personal-archive-for-online-conversations)
	- [Getting Started](#getting-started)
		- [Configure your `.env` file](#configure-your-env-file)
		- [Running locally](#running-locally)
		- [Running with docker](#running-with-docker)
	- [Navigating the code](#navigating-the-code)
		- [Relevant directories/files](#relevant-directoriesfiles)
			- [`middleware/JWT.go`](#middlewarejwtgo)
			- [`handlers`](#handlers)
## Tech stack
- **Go** server hosting
- **PostgreSQL** database management with [go-prisma](https://goprisma.org/) ORM.
- **Docker** and **Google cloud** for deployment
## Getting Started
### Requirements
- `go` (This project is developed based on `go1.23.4 linux/amd64`.)
- `PostgreSQL` database. For this project, you can host your PostgreSQL database locally (via docker) or using Neon database.
- `Docker` (Optional) This is in case you want to build and run with docker.
Make sure to add your `.env` file before starting the server. Here is the example of `.env` file.
```bash
PORT=5000
DATABASE_URL=[YOUR_POSTGRESQL_DB_URL]
JWT_SECRET_KEY=[YOUR_JWT_SECRET_KEY]
```
- `DATABASE_URL`: For this project, you can host your PostgreSQL database locally (via docker) or using Neon database.
- `JWT_SECRET_KEY`: You can choose any string you wish to choose.
### Running locally
```bash
go mod download
go run cmd/server/main.go
```
### Running with docker
To start the server, execute
```bash
docker build --network=host --tag thread-keep:latest .
```
To see a list of built containers, you can use the `docker images` command. You would expect to see something like this.
```
REPOSITORY       TAG       IMAGE ID       CREATED         SIZE
thread-keep      latest    <id>         i <time>          <size>
```
To start the server, execute
```bash
docker run --env-file .env -d --name thread-keep -p 5000:5000 thread-keep:latest
```

If this server is settled properly, you would expect to find the message `"Welcome to our api server!"` at http://localhost:5000/.

To see a list of running containers, you can use the `docker ps` command. You would expect something like this.
```
CONTAINER ID   IMAGE                COMMAND    CREATED         STATUS         PORTS                                       NAMES
<id>           thread-keep:latest   "./main"   2 minutes ago   Up 2 minutes   0.0.0.0:5000->5000/tcp, :::5000->5000/tcp   thread-keep
```
To stop the container, execute `docker stop <id>` or `docker stop thread-keep`. To remove the container, run `docker rm thread-keep`.
## Deployment
This project used Google Cloud Platform (GCP) for deployment. Please follow the instruction from this [link](https://medium.com/novai-cloud-computing/gcp-docker-golang-deploying-a-go-application-to-google-cloud-container-registry-and-cloud-run-b5056324b5d0) for more details. 
## Navigating the code
This is the main file structure of our project, based on [this repository](https://github.com/CVWO/sample-go-app).
```
.
├── cmd
│   ├── server      # Main server
│   ├── tag         # Handle tag management
├── internal
│   ├── api         # Encapsulates types and utilities related to the API
│   ├── dataacess   # Data Access layer accesses data from the database
│   ├── database    # Encapsulates the types and utilities related to the database
│   ├── handlers    # Handler functions to respond to requests
│   ├── middleware	# Handle middleware (such as authentication system)  
│   ├── models      # Definitions of objects used in the application
│   ├── router      # Encapsulates types and utilities related to the router
│   ├── routes      # Defines routes that are used in the application
├── README.md
├── go.mod
└── go.sum
```
### Relevant directories/files
#### `middleware/JWT.go`
This backend server utilizes JSON Web Tokens (JWT) for user authentication and to restrict access to API calls. We employ the golang-jwt library to generate new JWT tokens upon successful user login and to verify existing tokens for authentication purposes.

#### `handlers`
Handler functions are responsible for providing `api` response based on HTTP request. Each subdirectory consists of `types.go`, `messages.go`, and basic CRUD logics (such as `create.go` and `update.go`.)

All handler functions are named by this format: `Handle[OP]`. Here is the example of `HandleDelete` from `handlers/threads/delete.go`.

```go
func HandleDelete(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
    // check if it is POST request or not
	if r.Method != http.MethodPost {
		err := errors.New(ErrInvalidPostRequest)
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

    // parse body
	thread := &ThreadDeleteRequest{}
	err := json.NewDecoder(r.Body).Decode(thread)

	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

    // connect to database
	db, err := database.Connect()
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	defer db.Close()

    // perform database operation
	threadObject, err := mutation.DeleteThread(db, thread.ThreadID)
	
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}
	
    // convert to JSON format
	data, err := json.Marshal(threadObject)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

    // response
	return utils.WrapHTTPPayload(data, SuccessfulDeleteThread)
}
```

Note that function from `utils` are designed to encapsulate error / response messages with function name. 
