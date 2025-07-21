# Rivalz Framework Backend

A high-performance Go backend framework built with modern technologies for scalable web applications and microservices.

## 🚀 Features

- **High Performance**: Built with Go 1.23 and Fiber web framework
- **Multi-Database Support**: MongoDB, PostgreSQL, and Redis integration
- **Message Queue**: Kafka integration for event-driven architecture
- **Monitoring**: Elastic APM integration for application performance monitoring
- **Authentication**: JWT-based authentication system
- **API Documentation**: Swagger/OpenAPI documentation
- **Container Ready**: Docker support with multi-stage builds
- **Blockchain Integration**: Ethereum smart contract support
- **Modular Architecture**: Clean separation of concerns with modular design

## 🛠 Tech Stack

- **Language**: Go 1.23
- **Web Framework**: Fiber v2
- **Databases**: 
  - MongoDB (via mongo-driver)
  - PostgreSQL (via GORM)
  - Redis (for caching and pub/sub)
- **Message Queue**: Apache Kafka
- **Monitoring**: Elastic APM
- **Authentication**: JWT
- **Documentation**: Swagger
- **Container**: Docker

## 📋 Prerequisites

- Go 1.23 or higher
- Docker (optional)
- MongoDB
- PostgreSQL
- Redis
- Apache Kafka

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/Rivalz-ai/framework-be.git
cd framework-be
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Environment Configuration

Create a `.env` file in the root directory with your configuration:

```env
# Database Configuration
MONGODB_URI=mongodb://localhost:27017
POSTGRES_URI=postgres://user:password@localhost:5432/dbname
REDIS_URI=redis://localhost:6379

# Kafka Configuration
KAFKA_BROKERS=localhost:9092

# JWT Configuration
JWT_SECRET=your-secret-key

# Server Configuration
PORT=30000
```

### 4. Run the Application

#### Development Mode
```bash
go run main.go
```

#### Using Air (Hot Reload)
```bash
air
```

#### Using Docker
```bash
# Build the image
docker build -t rivalz-framework-be .

# Run the container
docker run -p 30000:30000 rivalz-framework-be
```

## 📚 API Documentation

Once the application is running, you can access the Swagger documentation at:

```
http://localhost:30000/swagger/
```

## 🏗 Project Structure

```
framework-be/
├── acl/                    # Access Control Layer
├── define/                 # Contract definitions and ABIs
│   └── abi/               # Ethereum smart contract ABIs
├── middleware/             # HTTP middleware
├── models/                 # Data models
├── modules/                # Business logic modules
│   ├── agent/             # Agent management
│   ├── project/           # Project management
│   ├── user/              # User management
│   └── x/                 # Extended functionality
├── routes/                 # API route definitions
├── server/                 # Server configuration
├── types/                  # Type definitions
├── main.go                 # Application entry point
├── go.mod                  # Go module file
├── go.sum                  # Go module checksums
├── Dockerfile             # Docker configuration
└── air.toml               # Air hot reload configuration
```

## 🔧 Configuration

The application supports multiple configuration sources:

- Environment variables
- Configuration files
- Vault integration for secrets management

### Key Configuration Areas

- **Database**: MongoDB, PostgreSQL, and Redis connections
- **Kafka**: Message broker configuration
- **APM**: Elastic APM monitoring setup
- **JWT**: Authentication secret keys
- **HTTP**: Server port and middleware settings

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific module tests
go test ./modules/agent/...
```

## 📦 Deployment

### Docker Deployment

```bash
# Build production image
docker build -t rivalz-framework-be:latest .

# Run with environment variables
docker run -d \
  -p 30000:30000 \
  -e MONGODB_URI=mongodb://your-mongo:27017 \
  -e POSTGRES_URI=postgres://user:pass@your-postgres:5432/db \
  -e REDIS_URI=redis://your-redis:6379 \
  rivalz-framework-be:latest
```

### Kubernetes Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: rivalz-framework-be
spec:
  replicas: 3
  selector:
    matchLabels:
      app: rivalz-framework-be
  template:
    metadata:
      labels:
        app: rivalz-framework-be
    spec:
      containers:
      - name: rivalz-framework-be
        image: rivalz-framework-be:latest
        ports:
        - containerPort: 30000
        env:
        - name: MONGODB_URI
          valueFrom:
            secretKeyRef:
              name: db-secrets
              key: mongodb-uri
        - name: POSTGRES_URI
          valueFrom:
            secretKeyRef:
              name: db-secrets
              key: postgres-uri
```

## 🔍 Monitoring

The application integrates with Elastic APM for monitoring:

- Application performance metrics
- Request tracing
- Error tracking
- Database query monitoring

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

For support and questions:

- Create an issue in the GitHub repository
- Contact the development team
- Check the API documentation at `/swagger/`

## 🔄 Version History

- **v1.0.0**: Initial release with core framework features
- Support for MongoDB, PostgreSQL, and Redis
- Kafka integration for event-driven architecture
- Elastic APM monitoring
- JWT authentication
- Swagger API documentation

---

Built with ❤️ by the Rivalz Team
