# README

A simple web application that calculates the Fibonacci sequence of a length specified by the user.

## Requirements

Before running the application, ensure you have the following installed on your system:

1. Golang (Go programming language) - to build or run the application locally

2. Docker - to containerize and run the application easily

3. cURL - to test the running service from the command line

## How to Run

You can run this application in three different ways: directly with Go, using Docker, or via Docker Compose.

1. Run with Golang (local development)
If you want to run the application locally without Docker, navigate to the root directory of the project and run:

```bash
go run main.go
```

This will start the web server on port 8080 by default. You can then access the service locally.

2. Run with Docker
To run the application inside a Docker container, follow these steps:

Build the Docker image by running:

```bash
docker build -t fib .
```

Run the Docker container, exposing port 8080:

```bash
docker run -p 8080:8080 --rm fib
```

The --rm flag will automatically remove the container after it stops.

3. Run with Docker Compose
If you prefer to use Docker Compose (recommended for managing multi-container setups or just convenience), run:

```bash
docker compose up
```

This will build (if necessary) and start the container, mapping port 8080 on your host to the container’s port.

### Testing the Application

Once the application is running via any of the methods above, you can test the Fibonacci sequence endpoint using curl:

```bash
curl http://localhost:8080/api/fibonacci/5 | jq
```

You should get a JSON response with the Fibonacci sequence of length 5:

```json
{
  "numbers": [
    0,
    1,
    1,
    2,
    3
  ]
}
```

Here, /api/fibonacci/5 is the endpoint where 5 is the length of the Fibonacci sequence requested.

## Limitations

The maximum length of the Fibonacci sequence that can be requested depends on the bit-size of the architecture the service is running on.

This is because the Fibonacci numbers grow exponentially and will overflow the integer size limits of the underlying system.

The application automatically detects the architecture bit size and limits the maximum allowed sequence length to prevent integer overflow and unexpected errors.

For 64-bit architectures, the max Fibonacci sequence length is 93.

For 32-bit architectures, the max length is 47.


## K8s

If in some use case this application was to be run in Kubernetes here is scaffolding on how to deploy it:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: fibonacci-service
spec:
  replicas: 1
  selector:
    matchLabels:
      app: fibonacci-service
  template:
    metadata:
      labels:
        app: fibonacci-service
    spec:
      containers:
        - name: fibonacci-service
          image: jasmingacic/fib:latest
          ports:
            - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: fibonacci-service
spec:
  type: LoadBalancer
  selector:
    app: fibonacci-service
  ports:
    - protocol: TCP
      port: 8080
      targetPort: 8080
```


