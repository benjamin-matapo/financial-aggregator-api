# API Testing Guide

This guide provides examples of how to test the Financial Aggregator API endpoints.

## 🚀 Quick Start

1. **Start the server:**
```bash
cd apps/backend
go run main.go
```

2. **Test the health endpoint:**
```bash
curl http://localhost:8080/health
```

## 📊 API Endpoints Testing

### 1. Health Check
```bash
curl http://localhost:8080/health
```
**Expected Response:**
```json
{"status":"healthy","timestamp":"2025-09-21T17:46:04+01:00"}
```

### 2. Get All Accounts
```bash
curl http://localhost:8080/api/accounts
```
**Expected Response:**
```json
{
  "success": true,
  "message": "Accounts retrieved successfully",
  "data": [
    {
      "id": "acc_001",
      "name": "Primary Checking",
      "bank": "Chase Bank",
      "account_type": "checking",
      "balance": 2500.75,
      "currency": "USD",
      "last_updated": "2025-09-21T15:45:57.182901+01:00",
      "is_active": true
    }
    // ... more accounts
  ]
}
```

### 3. Get Specific Account
```bash
curl http://localhost:8080/api/accounts/acc_001
```

### 4. Refresh Account Data
```bash
curl -X POST http://localhost:8080/api/accounts/acc_001/refresh
```
**Expected Response:**
```json
{
  "success": true,
  "message": "account data refreshed successfully",
  "data": {
    "account_id": "acc_001",
    "success": true,
    "message": "account data refreshed successfully",
    "last_updated": "2025-09-21T17:46:20.301899+01:00",
    "new_balance": 2499.75
  }
}
```

### 5. Get All Transactions
```bash
curl http://localhost:8080/api/transactions
```

### 6. Get Transactions with Filters
```bash
# Filter by account
curl "http://localhost:8080/api/transactions?account_id=acc_001"

# Filter by type
curl "http://localhost:8080/api/transactions?type=debit"

# Filter by category
curl "http://localhost:8080/api/transactions?category=food"

# Filter by date range
curl "http://localhost:8080/api/transactions?start_date=2024-01-01&end_date=2024-01-31"

# Limit results
curl "http://localhost:8080/api/transactions?limit=5"

# Combine filters
curl "http://localhost:8080/api/transactions?account_id=acc_001&type=debit&limit=3"
```

### 7. Get Account Transactions
```bash
curl http://localhost:8080/api/accounts/acc_001/transactions
```

### 8. Get Specific Transaction
```bash
curl http://localhost:8080/api/transactions/txn_001
```

## 🧪 Testing with jq (JSON processor)

For better formatted output, install `jq` and pipe responses:

```bash
# Install jq (macOS)
brew install jq

# Test with formatted output
curl -s http://localhost:8080/api/accounts | jq .
curl -s http://localhost:8080/api/transactions | jq .
```

## 🔍 Testing Error Cases

### 1. Invalid Account ID
```bash
curl http://localhost:8080/api/accounts/invalid_id
```
**Expected Response:**
```json
{
  "success": false,
  "message": "Account not found",
  "error": "account not found"
}
```

### 2. Invalid Transaction ID
```bash
curl http://localhost:8080/api/transactions/invalid_id
```

### 3. Refresh Non-existent Account
```bash
curl -X POST http://localhost:8080/api/accounts/invalid_id/refresh
```

## 📊 Performance Testing

### 1. Load Testing with Apache Bench
```bash
# Install Apache Bench
brew install httpd

# Test accounts endpoint
ab -n 100 -c 10 http://localhost:8080/api/accounts

# Test transactions endpoint
ab -n 100 -c 10 http://localhost:8080/api/transactions
```

### 2. Concurrent Testing
```bash
# Test multiple endpoints simultaneously
for i in {1..10}; do
  curl -s http://localhost:8080/api/accounts > /dev/null &
  curl -s http://localhost:8080/api/transactions > /dev/null &
done
wait
```

## 🐳 Docker Testing

### 1. Build and Run Docker Container
```bash
# Build the image
docker build -t financial-aggregator-api .

# Run the container
docker run -p 8080:8080 financial-aggregator-api

# Test the API
curl http://localhost:8080/health
```

### 2. Test Docker Health Check
```bash
# Check container health
docker ps
docker inspect <container_id> | grep Health
```

## 🔧 Development Testing

### 1. Run Tests
```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run tests with race detection
make test-race
```

### 2. Code Quality Checks
```bash
# Run linter
make lint

# Format code
make fmt

# Run all checks
make check
```

## 📝 Test Data

The API comes with pre-loaded mock data:

### Accounts
- **acc_001**: Primary Checking (Chase Bank) - $2,500.75
- **acc_002**: High Yield Savings (Ally Bank) - $15,000.00
- **acc_003**: Credit Card (Capital One) - -$1,200.50
- **acc_004**: Investment Account (Fidelity) - $45,000.25
- **acc_005**: Business Checking (Wells Fargo) - $8,500.00
- **acc_006**: Emergency Fund (Marcus) - $25,000.00

### Transaction Categories
- **Income**: salary, investment, business
- **Expenses**: food, utilities, transportation, entertainment, healthcare
- **Transfers**: between accounts

### Transaction Types
- **debit**: money going out
- **credit**: money coming in
- **transfer**: between accounts

## 🚨 Troubleshooting

### Common Issues

1. **Port already in use**
```bash
lsof -ti:8080 | xargs kill -9
```

2. **Server not responding**
```bash
# Check if server is running
ps aux | grep "go run main.go"

# Check port usage
lsof -i:8080
```

3. **JSON parsing errors**
```bash
# Use jq to validate JSON
curl -s http://localhost:8080/api/accounts | jq .
```

4. **CORS issues**
- The API is configured to allow requests from `http://localhost:3000` and `https://*.vercel.app`
- Check browser console for CORS errors

## 📈 Monitoring

### 1. Check Server Logs
```bash
# Run server in foreground to see logs
go run main.go
```

### 2. Monitor Resource Usage
```bash
# Check memory and CPU usage
top -p $(pgrep -f "go run main.go")
```

### 3. Test Response Times
```bash
# Time API responses
time curl -s http://localhost:8080/api/accounts > /dev/null
```

This guide should help you thoroughly test the Financial Aggregator API and ensure it's working correctly in all scenarios.
