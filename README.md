# E-Commerce Microservices Platform (Go + gRPC)

## Overview

This project is a microservices-based e-commerce backend implemented in **Go**.
It follows **Clean Architecture principles** and uses **gRPC** for inter-service communication.

The system consists of:

* **API Gateway**
* **Inventory Service**
* **Order Service**
* Shared **Protocol Buffers (proto)** definitions

Each service is independently structured with clear separation of:

* `entity`
* `repository`
* `usecase`
* `adapter` (HTTP / gRPC)
* `cmd` (entry point)

---

## Architecture

```
Client
   ↓
API Gateway (HTTP)
   ↓ gRPC
Inventory Service
Order Service
```

### Communication

* HTTP → API Gateway
* gRPC → Between services
* Protobuf → Contract definitions (`proto/`)

---

## Project Structure

```
ecommerce-service/
└── ass2/
    ├── api-gateway/
    │   ├── cmd/
    │   ├── internal/
    │   │   └── routes/
    │   └── services/
    │
    ├── inventory-service/
    │   ├── cmd/
    │   ├── config/
    │   ├── internal/
    │   │   ├── adapter/
    │   │   │   ├── grpc/
    │   │   │   └── http/
    │   │   ├── entity/
    │   │   ├── repository/
    │   │   └── usecase/
    │
    ├── order-service/
    │   ├── cmd/
    │   ├── config/
    │   ├── internal/
    │   │   ├── adapter/
    │   │   │   ├── grpc/
    │   │   │   └── http/
    │   │   ├── entity/
    │   │   ├── repository/
    │   │   └── usecase/
    │
    └── proto/
        ├── inventory.proto
        ├── order.proto
        └── generated *.pb.go files
```

---

## Services

### 1. API Gateway

Acts as a single entry point for clients.

Responsibilities:

* HTTP routing
* Forwarding requests to:

  * Inventory Service
  * Order Service
* Service-level abstraction

Contains:

* `routes/`
* `services/`
* `cmd/main.go`

---

### 2. Inventory Service

Responsible for product and review management.

### Domain Entities

* `Product`
* `Category`
* `Review`

### Layers

* `entity/` → Business models and domain errors
* `repository/` → Data access abstraction
* `usecase/` → Business logic
* `adapter/http/` → REST handlers
* `adapter/grpc/` → gRPC handlers

---

### 3. Order Service

Responsible for order processing.

### Domain Entities

* `Order`

### Layers

* `entity/`
* `repository/`
* `usecase/`
* `adapter/http/`
* `adapter/grpc/`

---

## Clean Architecture Implementation

Each service follows this structure:

```
cmd → adapter → usecase → repository → entity
```

* **Entity**: Core business models
* **Usecase**: Business logic
* **Repository**: Data abstraction layer
* **Adapter**: Interface layer (HTTP / gRPC)
* **Cmd**: Application entry point

This ensures:

* Separation of concerns
* Testability
* Scalability
* Maintainability

---

## gRPC Contracts

Located in:

```
proto/
```

Files:

* `inventory.proto`
* `order.proto`

Generated code:

* `inventorypb/`
* `orderpb/`

These define:

* Service methods
* Request/Response structures
* Inter-service contracts

---

## Technologies Used

* Go
* gRPC
* Protocol Buffers
* Clean Architecture pattern
* Modular monorepo structure

---

## How to Run

Each service must be started independently.

### 1. Start Inventory Service

```
cd inventory-service
go run cmd/main.go
```

### 2. Start Order Service

```
cd order-service
go run cmd/main.go
```

### 3. Start API Gateway

```
cd api-gateway
go run cmd/main.go
```

Make sure services are running before using the API Gateway.

---

## Educational Purpose

This project demonstrates:

* Microservices architecture in Go
* gRPC communication
* Protobuf contract design
* Clean Architecture implementation
* Separation of HTTP and gRPC adapters
* Service-to-service communication via defined contracts
