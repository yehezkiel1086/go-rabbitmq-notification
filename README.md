# Go Gin, RabbitMQ & PostgreSQL: Microservice Notification System

This project is a sample implementation of a microservice architecture using Go. It demonstrates how to handle a user registration process where a RESTful API service communicates with a background worker service via a message queue (RabbitMQ) to send a welcome email notification.

## Architecture Overview

The system consists of two main services, a database, and a message broker, all orchestrated using Docker.

1.  **Auth Service**: A REST API built with the Gin framework. It exposes an endpoint for user registration. When a new user registers, it saves the user's data to the PostgreSQL database and publishes a `user_registered` event to a RabbitMQ queue.
2.  **Email Service**: A background worker that listens for messages on the `user_registered` RabbitMQ queue. When it receives a message, it processes it and simulates sending a welcome email to the newly registered user.
3.  **PostgreSQL**: A relational database used to store user information.
4.  **RabbitMQ**: A message broker that facilitates asynchronous communication between the `Auth Service` and the `Email Service`, decoupling the services from each other.

### Data Flow

```
1. Client sends a POST request to /api/v1/register
      |
      v
2. Auth Service (Gin API)
   - Validates the request
   - Saves user data to PostgreSQL
   - Publishes a message to RabbitMQ
      |
      v
3. RabbitMQ (Message Queue)
   - Holds the 'user_registered' message
      |
      v
4. Email Service (Consumer)
   - Receives the message
   - Simulates sending a welcome email
```

## Technologies Used

- **Backend**: Go
- **API Framework**: Gin
- **Database**: PostgreSQL
- **Message Broker**: RabbitMQ
- **Containerization**: Docker & Docker Compose
- **ORM**: GORM (inferred from usage patterns)
- **RabbitMQ Client**: `github.com/rabbitmq/amqp091-go`

## Getting Started

Follow these instructions to get the project running on your local machine.

### Prerequisites

- Docker
- Docker Compose
- Go (for running services locally outside of Docker)

### 1. Clone the Repository

```bash
git clone <your-repository-url>
cd go-gin-rabbitmq-email-notif
```

### 2. Configure Environment Variables

Create a `.env` file in the root of the project by copying the example file:

```bash
cp .env.example .env
```

Now, open the `.env` file and set your desired credentials and ports.

```env
# PostgreSQL
DB_USER=admin
DB_PASS=password
DB_NAME=users_db
DB_PORT=5434 # Host port for Postgres

# RabbitMQ
RABBITMQ_USER=guest
RABBITMQ_PASS=guest
RABBITMQ_PORT=5672 # Host port for RabbitMQ AMQP

# Auth Service
HTTP_HOST=0.0.0.0
HTTP_PORT=8080
```

> **Note**: The `DB_PORT` is the *external* port on your host machine that maps to the PostgreSQL container's port 5432. If you encounter port conflicts, you can change this value.

### 3. Run with Docker Compose

This is the recommended way to run the entire application stack.

```bash
docker-compose up --build
```

This command will:
- Build the Docker images for `auth-service` and `email-service`.
- Start containers for PostgreSQL, RabbitMQ, and both Go services.
- You will see logs from all services in your terminal.

### 4. Verify Services

- **RabbitMQ Management UI**: Open your browser and navigate to `http://localhost:15672`. You can log in with the `RABBITMQ_USER` and `RABBITMQ_PASS` from your `.env` file.
- **PostgreSQL**: You can connect to the database using a client like DBeaver or `psql` on `localhost` with the `DB_PORT` you specified.

### 5. Test the API

Use a tool like `curl` or Postman to send a `POST` request to the registration endpoint.

```bash
curl -X POST http://localhost:8080/api/v1/register \
-H "Content-Type: application/json" \
-d '{
    "email": "test@example.com",
    "name": "Test User"
}'
```

**Expected Response:**

```json
{
    "message": "User registered successfully!"
}
```

After sending the request, check the logs from the `docker-compose up` command. You should see:
1. A log from `auth-service` indicating it published a new user.
2. A log from `email-service` indicating it received the message and is "sending" an email.
