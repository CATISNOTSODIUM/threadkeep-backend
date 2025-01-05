# ThreadKeep ⬢  - Your personal archive for online conversations.
This is the Go backend for ThreadKeep ⬢. For more information, please refer to this [repository](https://github.com/CATISNOTSODIUM/threadkeep-frontend).

## Getting Started

### Installing Go

Download and install Go by following the instructions [here](https://go.dev/doc/install).

### Running Locally
1. [Fork](https://docs.github.com/en/get-started/quickstart/fork-a-repo#forking-a-repository) this repo.
2. [Clone](https://docs.github.com/en/get-started/quickstart/fork-a-repo#cloning-your-forked-repository) **your** forked repo.
3. Open your terminal and navigate to the directory containing your cloned project.
4. Run `go run cmd/server/main.go` and head over to http://localhost:8000/users to view the response.

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
### Navigating the code
This is the main file structure. Note that this is simply *one of* various paradigms to organise your code, and is just a bare starting point.
```
.
├── cmd
│   ├── server
├── internal
│   ├── api         # Encapsulates types and utilities related to the API
│   ├── dataacess   # Data Access layer accesses data from the database
│   ├── database    # Encapsulates the types and utilities related to the database
│   ├── handlers    # Handler functions to respond to requests
│   ├── models      # Definitions of objects used in the application
│   ├── router      # Encapsulates types and utilities related to the router
│   ├── routes      # Defines routes that are used in the application
├── README.md
├── go.mod
└── go.sum
```

Main directories/files to note:
* `cmd` contains the main entry point for the application
* `internal` holds most of the functional code for your project that is specific to the core logic of your application
* `README.md` is a form of documentation about the project. It is what you are reading right now.
* `go.mod` contains important metadata, for example, the dependencies in the project. See [here](https://go.dev/ref/mod) for more information
* `go.sum` See [here](https://go.dev/ref/mod) for more information

Try changing some source code and see how the app changes.

## Next Steps

* This project uses [go-chi](https://github.com/go-chi/chi) as a web framework. Feel free to explore other web frameworks such as [gin-gonic](https://github.com/gin-gonic/gin). Compare their pros and cons and use whatever that best justifies the trade-offs.
* Read up more on the [MVC framework](https://developer.mozilla.org/en-US/docs/Glossary/MVC) which this code is designed upon.
* Sometimes, code formatting can get messy and opiniated. Do see how you can incoporate [linters](https://github.com/golangci/golangci-lint) to format your code.
