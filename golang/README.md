# Go APP - Golang

Go app is learn about golang.

## Technology Stack

### Backend

- **Language**: Go 1.24
- **Framework**: Fiber

## Installation

1. Clone the repository:

```bash
git clone https://github.com/xamppsul/lab-code.git #basic clone
git clone git@github.com:xamppsul/lab-code.git #with SSH
cd lab-code/golang #open dir golang
```

## Running the Application

Start the server:

```bash
#makefile
make start
#base command
go run index.go
# if want to use of docker engine
make docker-build #start docker service
make docker-stop #stop docker service
```

## If problem about .env for run docker engine

1. delete last docker image
2. change path docker config.EnvConfig(".env") in index.go
3. run make docker-build
4. happy running

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
