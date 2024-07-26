# toll-calculator

# Tolling System

## Project Overview

This project implements a comprehensive tolling system that consists of multiple services working together to provide real-time data aggregation and processing. The system includes an On-Board Unit (OBU) simulator, a data receiver, a distance calculator, an aggregator, and a gateway service, all built using Go and various technologies. The architecture leverages microservices and real-time data processing capabilities, making it scalable and efficient.

## Problems the Project Solves

1. **Real-Time Data Processing**: The system processes real-time OBU data to calculate distances traveled and aggregates this data for further analysis.
2. **Scalable Architecture**: By using a microservices architecture, the system can be easily scaled horizontally, ensuring high availability and resilience.
3. **Data Aggregation and Analysis**: The aggregator service collects and processes data from multiple OBUs, providing insights and reports on the data.
4. **Middleware and Observability**: The system includes logging, metrics, and circuit breaker patterns to ensure reliability and observability of the services.
5. **Simulation of OBU Devices**: The OBU simulator generates realistic data for testing and validation purposes.
6. **Distributed Systems Concepts**: The system incorporates distributed systems principles such as data partitioning, replication, and fault tolerance to ensure reliability and scalability.

## All Tech Stacks and Libraries

### Languages and Frameworks
- **Go**: The primary programming language used for all services.
- **gRPC**: For communication between microservices.
- **Go Kit**: A toolkit for microservices in Go, providing tools for building robust services.

### Messaging and Data Handling
- **Kafka**: Used for messaging between the data receiver and other services.
- **Zookeeper**: Manages the Kafka broker.

### Web and Networking
- **WebSocket**: Used by the OBU simulator to send data to the data receiver.
- **HTTP**: Used for communication between the gateway and aggregator services.

### Configuration and Environment
- **Docker Compose**: For setting up and managing the development environment.
- **.env**: For environment configuration.

### Observability and Logging
- **Prometheus**: For metrics collection and monitoring.
- **Logrus**: For structured logging.
- **Promauto**: For creating Prometheus metrics in Go applications.

## How the Code Works Together

### Distributed Systems Principles

1. **Microservices Architecture**: The system is decomposed into small, independent services that communicate over well-defined APIs. This allows for independent deployment and scaling of each service.
2. **Data Partitioning**: Data is partitioned across different services, each responsible for a specific part of the data processing pipeline.
3. **Replication**: Data is replicated across services to ensure high availability and fault tolerance.
4. **Fault Tolerance**: The system uses circuit breakers, retries, and timeouts to handle failures gracefully and maintain service availability.
5. **Scalability**: Each service can be scaled independently to handle increased load, ensuring the system can grow with demand.

### Components

1. **OBU Simulator (`obu/main.go`)**: Generates random latitude and longitude data and sends it to the data receiver via WebSocket. This simulates real-world data from On-Board Units.

2. **Data Receiver (`data_receiver/main.go`)**: Listens for WebSocket connections and receives OBU data. It uses a Kafka producer to forward the received data to a Kafka topic.

3. **Distance Calculator (`distance_calculator`)**: 
   - **Kafka Consumer (`consumer.go`)**: Reads messages from the Kafka topic, calculates the distance based on the received coordinates, and sends the aggregated distance data to the aggregator service using an HTTP client.
   - **Service and Middleware**: Implements the core business logic and includes logging middleware to monitor performance and errors.

4. **Aggregator Service (`aggregator`)**: 
   - **gRPC and HTTP Servers (`grpc.go` and `http.go`)**: Provide endpoints for aggregating distance data and retrieving invoices.
   - **Service and Middleware**: Implements the aggregation logic, storing data in an in-memory store, and includes logging and metrics middleware for observability.

5. **Gateway Service (`gateway/main.go`)**: 
   - **HTTP Server**: Acts as an intermediary, handling HTTP requests from clients and forwarding them to the aggregator service. This allows clients to retrieve invoices and send aggregation requests easily.
   - **Handlers and Middleware**: Implements handlers for processing requests and includes middleware for logging request details and performance metrics.

6. **Protobuf Definitions and Generated Code (`types`)**: 
   - **Protobuf (`ptypes.proto`)**: Defines the messages and services for the gRPC communication.
   - **Generated Code**: Includes the generated Go code for the protobuf definitions and gRPC clients and servers.

7. **Environment Configuration (`.env`)**: Specifies environment variables for configuring the endpoints and store type.

8. **Docker Compose Setup (`docker-compose.yml`)**: Sets up the Zookeeper and Kafka services required for messaging between the components.

9. **Build and Run Automation (`Makefile`)**: Provides targets for building and running each service, as well as generating protobuf code.


```
docker run --name kafka -p 9092:9092 -e ALLOW_PLAINTEXT_LISTENER=yes -e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=true bitnami/kafka:latest 
```

## Installing protobuf compiler (protoc compiler) 
For linux users or (WSL2) 
```
sudo apt install -y protobuf-compiler
```

For Mac users you can use Brew for this
```
brew install protobuff
```

## Installing GRPC and Protobuffer plugins for Golang.
1. Protobuffers
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```

2. GRPC 
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

3. NOTE that you need to set the /go/bin directory in your path
Just like this or whatever your go directly lives.
```
PATH="${PATH}:${HOME}/go/bin"
```

4. install the package dependencies
4.1 protobuffer package
```
go get google.golang.org/protobuf
```
4.2 grpc package
```
go get google.golang.org/grpc/
```

## Installing Prometheus
Install Prometheus in a Docker container
```
docker run -p 9090:9090 -v ./.config/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
```

Installing prometheus golang client
```
go get github.com/prometheus/client_golang/prometheus
```

Installing Prometheus natively on your system
1. Clone the repository
```
git clone https://github.com/promtheus/prometheus.git
```

2. Install
```
cd prometheus
make build
```

3. Run the Prometheus deamon
```
./promtheus --config.file=<your_config_file>yml
```

4. In the projects case that would be (running from inside the project directory)
```
../prometheus/prometheus --config.file=.config/prometheus.yml
```


