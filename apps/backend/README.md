# Financial Aggregator API - Backend

A production-ready Go REST API built with chi router, featuring clean architecture, comprehensive testing, and Docker support.

## 🏗️ Architecture

```
apps/backend/
├── handlers/           # HTTP request handlers
│   ├── account_handler.go
│   ├── account_handler_test.go
│   ├── transaction_handler.go
│   └── transaction_handler_test.go
├── services/           # Business logic layer
│   ├── account_service.go
│   └── transaction_service.go
├── models/             # Data models and DTOs
│   ├── account.go
│   ├── transaction.go
│   └── response.go
├── internal/           # Internal server configuration
│   └── server.go
├── main.go            # Application entry point
├── go.mod             # Go module dependencies
├── go.sum             # Dependency checksums
├── Dockerfile         # Container configuration
├── Makefile           # Build and development commands
├── .golangci.yml      # Linting configuration
└── README.md          # This file
```

## 🚀 Features

### REST API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check endpoint |
| GET | `/api/accounts` | Get all bank accounts |
| GET | `/api/accounts/{id}` | Get specific account |
| POST | `/api/accounts/{id}/refresh` | Refresh account data |
| GET | `/api/accounts/{id}/transactions` | Get account transactions |
| GET | `/api/transactions` | Get all transactions with filters |
| GET | `/api/transactions/{id}` | Get specific transaction |

### Query Parameters for `/api/transactions`

- `account_id` - Filter by account ID
- `type` - Filter by transaction type (debit, credit, transfer)
- `category` - Filter by category (food, salary, etc.)
- `status` - Filter by status (pending, completed, failed)
- `start_date` - Filter by start date (YYYY-MM-DD)
- `end_date` - Filter by end date (YYYY-MM-DD)
- `limit` - Limit number of results (default: 50)
- `offset` - Pagination offset (default: 0)

## 🛠️ Prerequisites

- **Go** 1.21 or higher
- **Make** (for build commands)
- **Docker** (optional, for containerized deployment)

## 📦 Installation & Setup

### 1. Clone and navigate to backend directory
```bash
cd apps/backend
```

### 2. Install dependencies
```bash
make deps
# or
go mod download
```

### 3. Set up development environment
```bash
make setup
```

This will install:
- `air` for hot reload during development
- `golangci-lint` for code linting

## 🚀 Running the Application

### Development Mode (with hot reload)
```bash
make dev
# or
air
```

### Production Mode
```bash
make run
# or
go run main.go
```

### Build and Run
```bash
make build
./bin/financial-aggregator-api
```

## 🧪 Testing

### Run all tests
```bash
make test
```

### Run tests with coverage
```bash
make test-coverage
# Generates coverage.html report
```

### Run tests with race detection
```bash
make test-race
```

### Run benchmarks
```bash
make benchmark
```

## 🔍 Code Quality

### Run linter
```bash
make lint
```

### Format code
```bash
make fmt
```

### Vet code
```bash
make vet
```

### Run all checks
```bash
make check
```

## 🐳 Docker Support

### Build Docker image
```bash
make docker-build
```

### Run Docker container
```bash
make docker-run
```

### Clean Docker image
```bash
make docker-clean
```

## 📊 API Usage Examples

### Get all accounts
```bash
curl http://localhost:8080/api/accounts
```

### Get specific account
```bash
curl http://localhost:8080/api/accounts/acc_001
```

### Refresh account data
```bash
curl -X POST http://localhost:8080/api/accounts/acc_001/refresh
```

### Get all transactions
```bash
curl http://localhost:8080/api/transactions
```

### Get transactions with filters
```bash
# Filter by account
curl "http://localhost:8080/api/transactions?account_id=acc_001"

# Filter by type and limit
curl "http://localhost:8080/api/transactions?type=debit&limit=10"

# Filter by date range
curl "http://localhost:8080/api/transactions?start_date=2024-01-01&end_date=2024-01-31"
```

### Get account transactions
```bash
curl http://localhost:8080/api/accounts/acc_001/transactions
```

## 📝 Response Format

### Success Response
```json
{
  "success": true,
  "message": "Accounts retrieved successfully",
  "data": [...]
}
```

### Error Response
```json
{
  "success": false,
  "message": "Account not found",
  "error": "account not found"
}
```

### Paginated Response
```json
{
  "success": true,
  "data": [...],
  "meta": {
    "total": 100,
    "limit": 50,
    "offset": 0,
    "pages": 2
  }
}
```

## 🔧 Configuration

### Environment Variables

- `PORT` - Server port (default: 8080)

### CORS Configuration

The API is configured to allow requests from:
- `http://localhost:3000` (React frontend)
- `https://*.vercel.app` (Vercel deployments)

## 🏗️ Architecture Details

### Clean Architecture

The application follows clean architecture principles:

1. **Handlers** - Handle HTTP requests and responses
2. **Services** - Contain business logic
3. **Models** - Define data structures
4. **Internal** - Server configuration and setup

### Concurrency Safety

- All services use `sync.RWMutex` for thread-safe operations
- Mock data is stored in memory with proper locking

### Error Handling

- Consistent error response format
- Proper HTTP status codes
- Detailed error messages for debugging

### Testing Strategy

- Unit tests for all handlers
- Mock services for isolated testing
- Comprehensive test coverage
- Race condition detection

## 🚀 Production Deployment

### Using Docker

1. Build the image:
```bash
docker build -t financial-aggregator-api .
```

2. Run the container:
```bash
docker run -p 8080:8080 financial-aggregator-api
```

### Using Vercel

The API is configured for Vercel serverless deployment:

1. Deploy from the project root
2. Vercel will automatically detect the Go application
3. The API will be available at `/api/*` endpoints

## 🔮 Future Enhancements

- [ ] Database integration (PostgreSQL/Cassandra)
- [ ] Authentication and authorization
- [ ] Rate limiting
- [ ] Request/response logging
- [ ] Metrics and monitoring
- [ ] API versioning
- [ ] Swagger/OpenAPI documentation
- [ ] Integration tests
- [ ] Performance optimization

## 🆘 Troubleshooting

### Common Issues

1. **Port already in use**: Change the PORT environment variable
2. **Go module issues**: Run `go mod tidy`
3. **Test failures**: Ensure all dependencies are installed
4. **Docker build fails**: Check Dockerfile and ensure Go version compatibility

### Debug Mode

Run with verbose logging:
```bash
go run main.go -v
```

## 📄 License

This project is licensed under the MIT License.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a pull request

## 📞 Support

For issues and questions:
- Check the troubleshooting section
- Review the test files for usage examples
- Open an issue in the repository
