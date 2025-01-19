# ThreadKeep ⬢  - Your personal archive for online conversations.
![Go](https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge)
![Prisma](https://img.shields.io/badge/Prisma-3982CE?style=for-the-badge&logo=Prisma&logoColor=white)

This is the Go backend for ThreadKeep ⬢. For more information, please refer to this [repository](https://github.com/CATISNOTSODIUM/threadkeep-frontend).

- [ThreadKeep ⬢  - Your personal archive for online conversations.](#threadkeep-----your-personal-archive-for-online-conversations)
	- [Tech stack](#tech-stack)
	- [Getting Started](#getting-started)
		- [Requirements](#requirements)
		- [Setting database](#setting-database)
		- [Running locally](#running-locally)
		- [Running with docker](#running-with-docker)
	- [Deployment](#deployment)
	- [Database design](#database-design)
	- [API Endpoints](#api-endpoints)
		- [User routers `/user`](#user-routers-user)
		- [Thread routers `/threads`](#thread-routers-threads)
			- [Thread reactions](#thread-reactions)
		- [Comment routers `/comments`](#comment-routers-comments)
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
PORT=8080
DATABASE_URL=[YOUR_POSTGRESQL_DB_URL]
JWT_SECRET_KEY=[YOUR_JWT_SECRET_KEY]
```
- `DATABASE_URL`: For this project, you can host your PostgreSQL database locally (via docker) or using Neon database.
- `JWT_SECRET_KEY`: You can choose any string you wish to choose.

### Setting database
- Schema file prisma/schema.ts
- Generate database migration by running `go run github.com/steebchen/prisma-client-go db push`.

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
docker run --env-file .env -d --name thread-keep -p 8080:8080 thread-keep:latest
```

If this server is settled properly, you would expect to find the message `"Welcome to our api server!"` at http://localhost:8080/.

To see a list of running containers, you can use the `docker ps` command. You would expect something like this.
```
CONTAINER ID   IMAGE                COMMAND    CREATED         STATUS         PORTS                                       NAMES
<id>           thread-keep:latest   "./main"   2 minutes ago   Up 2 minutes   0.0.0.0:8080->8080/tcp, :::8080->8080/tcp   thread-keep
```
To stop the container, execute `docker stop <id>` or `docker stop thread-keep`. To remove the container, run `docker rm thread-keep`. The current backend is deployed [here](https://thread-keep-798434152299.asia-southeast1.run.app/).
## Deployment
This project used Google Cloud Platform (GCP) for deployment. Please follow the instruction from this [link](https://medium.com/novai-cloud-computing/gcp-docker-golang-deploying-a-go-application-to-google-cloud-container-registry-and-cloud-run-b5056324b5d0) for more details. 

## Database design
// todo
## API Endpoints
This project mainly use REST API to perform standard database functions. For GET endpoints, all input parameters are parsed as url parameters. For POST endpoints, the requests are handled based on JSON format. Make sure to include relevant headers: `Content-Type: application/json` and `Authorization: Bearer [JWT_TOKEN]` for protected endpoints.
### User routers `/user`
- Handle all logics regarding user information.
- The `user` API endpoint does not require authorization header.

| Method          | Endpoint                        | Description | Input | data (If `http.status = 200`) |
| ------ | --------------- | ------------------------------- | --| --|
| GET          | `/users`                       | List all users in this platform | None | `[]models.User` with `jwt_token` set to `""`
| POST         | `/users/create`                | Create new user (Registeration) | `name` (string), `password` (string) | `[]models.User` with generated token
| POST         | `/users/verify`                | Handling user authentication (Sign in) | `name` (string), `password` (string) | `True \| False`

### Thread routers `/threads`
- Handle all CRUD thread operations and thread reactions.
- Require JWT authorization header (`Authorization: Bearer [JWT_TOKEN]`)

| Method | Endpoint   | Description                           | Input                                             | data (If `http.status = 200`) | 
|--------|------------|---------------------------------------|---------------------------------------------------|-------------------------------|
| GET    | `/threads` | List all threads based on user filter | (Optional) `skip` (int), `max_per_page` (int), `name` (string), `tags` (list of `tagID`)| `[]models.Thread`             |   |   |   |   |   |
| POST   | `/threads` | Obtain individual thread              | `threadID` (string)                               | `models.Thread` 
| GET | `/threads/count` | Count all threads | None | `int`
| POST   | `/threads/create` | Create new thread  | `title` (string) `content` (string) `user` (`models.User`) 
| POST   | `/threads/update` | Update thread details  | `ThreadID` (string) `title` (string) `content` (string) `user` (`models.User`)                             | `models.Thread` 
| POST   | `/threads/delete` | delete individual thread              | `threadID` (string)                               | `0 (Failed) \| 1 (Success)`
| GET | `/threads/tags` | Get all tags | None | `[]models.Tag`
| POST | `/threads/tags` | Get all tags for individual thread | `threadID` | `[]models.Tag`
| POST | `/threads/reaction` | Handle all threads reaction | `userID` (string), `threadID` (string), `reaction` ([ReactionType](#thread-reactions)) | `0 (Failed) \| 1 (Success)`
| POST | `/threads/reaction/isLike` | check if user with `userID` has liked the thread `threadID` or not. | `userID` (string), `threadID` (string) | `True \| False`

#### Thread reactions 
Thread reactions are handled with `ReactionType`.
| ReactionType | `int` |
| --- 		| --- 
| VIEW 		| 0	  
| LIKE		| 1
| UNLIKE	| 2
| SAVED	    | 3
| UNSAVE	| 4	

### Comment routers `/comments`
- Handle all CRUD comment operations.
- Require JWT authorization header (`Authorization: Bearer [JWT_TOKEN]`)

| Method | Endpoint   | Description                           | Input                                             | data (If `http.status = 200`) | 
|--------|------------|---------------------------------------|---------------------------------------------------|-------------------------------|
| POST   | `/comments` | Obtain individual comment              | `commentID` (string)                               | `models.Comment` 
| POST   | `/comments/create` | Create individual comment              |  `content` (string), `user` (`models.User`)               | `models.Comment` 
| POST   | `/comments/update` | Update individual comment              | `commentID` (string),  `content` (string), `user` (`models.User`)                                | `models.Comment` 
| POST   | `/comments/delete` | Delete individual comment              | `commentID` (string)                               | `0 (Failed) \| 1 (Success)` 



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
