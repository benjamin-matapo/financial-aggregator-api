# Financial Aggregator API

A full-stack financial aggregator application built with React (frontend) and Go (backend), designed as a monorepo and deployable on Vercel.

## 🏗️ Project Structure

```
financial-aggregator-api/
├── apps/
│   ├── frontend/          # React + Vite + Tailwind CSS
│   │   ├── src/
│   │   ├── public/
│   │   ├── package.json
│   │   ├── vite.config.ts
│   │   ├── tailwind.config.js
│   │   └── vercel.json
│   └── backend/           # Go API server
│       ├── main.go
│       ├── go.mod
│       ├── package.json
│       └── Dockerfile
├── package.json           # Root package.json with workspace config
├── pnpm-workspace.yaml    # pnpm workspace configuration
├── vercel.json           # Vercel deployment configuration
└── README.md
```

## 🚀 Features

### Frontend (React + Vite + Tailwind)
- Modern React 18 with TypeScript
- Vite for fast development and building
- Tailwind CSS for styling
- Responsive design with mobile-first approach
- Real-time financial data visualization
- Clean, modern UI with summary cards and data tables

### Backend (Go)
- RESTful API with Gin framework
- CORS enabled for frontend communication
- CRUD operations for financial data
- Health check endpoint
- Financial summary calculations
- Docker support for containerization

## 📋 Prerequisites

- **Node.js** (v18 or higher)
- **pnpm** (v8 or higher)
- **Go** (v1.21 or higher)
- **Docker** (optional, for containerized deployment)

## 🛠️ Installation & Setup

### 1. Clone the repository
```bash
git clone <your-repo-url>
cd financial-aggregator-api
```

### 2. Install dependencies
```bash
pnpm install
```

### 3. Development

#### Start both frontend and backend
```bash
pnpm dev
```

#### Start individual services
```bash
# Frontend only (runs on http://localhost:3000)
pnpm frontend:dev

# Backend only (runs on http://localhost:8080)
pnpm backend:dev
```

### 4. Build for production
```bash
pnpm build
```

### 5. Linting
```bash
pnpm lint
```

## 🌐 API Endpoints

The Go backend provides the following REST API endpoints:

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/health` | Health check |
| GET | `/api/financial-data` | Get all financial data |
| GET | `/api/financial-data/:id` | Get specific financial data |
| POST | `/api/financial-data` | Create new financial data |
| PUT | `/api/financial-data/:id` | Update financial data |
| DELETE | `/api/financial-data/:id` | Delete financial data |
| GET | `/api/summary` | Get financial summary |

### Example API Usage

```bash
# Get all financial data
curl http://localhost:8080/api/financial-data

# Create new financial data
curl -X POST http://localhost:8080/api/financial-data \
  -H "Content-Type: application/json" \
  -d '{
    "name": "New Income",
    "amount": 1000,
    "type": "income"
  }'

# Get financial summary
curl http://localhost:8080/api/summary
```

## 🐳 Docker Support

### Build and run with Docker

```bash
# Build the Docker image
cd apps/backend
docker build -t financial-aggregator-api .

# Run the container
docker run -p 8080:8080 financial-aggregator-api
```

## 🚀 Vercel Deployment

### 1. Deploy to Vercel

```bash
# Install Vercel CLI
npm i -g vercel

# Deploy from project root
vercel

# Follow the prompts to configure your project
```

### 2. Environment Variables

Set the following environment variables in your Vercel dashboard:

- `GIN_MODE=release` (for production mode)

### 3. Build Configuration

The project is configured to:
- Build the frontend as a static site
- Deploy the Go backend as serverless functions
- Route API calls to the backend
- Serve the frontend for all other routes

## 📁 File Descriptions

### Root Level Files

- **`package.json`**: Root package configuration with pnpm workspaces
- **`pnpm-workspace.yaml`**: Defines workspace packages
- **`vercel.json`**: Vercel deployment configuration
- **`.eslintrc.js`**: ESLint configuration for code quality
- **`.prettierrc`**: Prettier configuration for code formatting

### Frontend (`/apps/frontend/`)

- **`package.json`**: Frontend dependencies and scripts
- **`vite.config.ts`**: Vite configuration with proxy for API calls
- **`tailwind.config.js`**: Tailwind CSS configuration
- **`tsconfig.json`**: TypeScript configuration
- **`src/App.tsx`**: Main React component with financial dashboard
- **`src/main.tsx`**: React application entry point
- **`src/index.css`**: Global styles with Tailwind imports

### Backend (`/apps/backend/`)

- **`main.go`**: Go API server with Gin framework
- **`go.mod`**: Go module dependencies
- **`package.json`**: Backend scripts and configuration
- **`Dockerfile`**: Multi-stage Docker build for production

## 🎨 Frontend Features

- **Responsive Dashboard**: Clean, modern interface that works on all devices
- **Financial Summary Cards**: Total income, expenses, and net worth at a glance
- **Data Table**: Detailed view of all financial transactions
- **Real-time Updates**: Refresh button to fetch latest data
- **Loading States**: Smooth loading indicators
- **Error Handling**: User-friendly error messages

## 🔧 Development Scripts

| Script | Description |
|--------|-------------|
| `pnpm dev` | Start both frontend and backend in development mode |
| `pnpm build` | Build all applications for production |
| `pnpm lint` | Run ESLint on all packages |
| `pnpm type-check` | Run TypeScript type checking |
| `pnpm clean` | Clean build artifacts |
| `pnpm frontend:dev` | Start only the frontend |
| `pnpm backend:dev` | Start only the backend |

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Make your changes and commit: `git commit -m 'Add feature'`
4. Push to the branch: `git push origin feature-name`
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Troubleshooting

### Common Issues

1. **Port conflicts**: Make sure ports 3000 and 8080 are available
2. **pnpm not found**: Install pnpm globally with `npm install -g pnpm`
3. **Go module issues**: Run `go mod tidy` in the backend directory
4. **Build failures**: Ensure all dependencies are installed with `pnpm install`

### Getting Help

- Check the console for error messages
- Ensure all prerequisites are installed
- Verify that all services are running on the correct ports
- Check the API endpoints directly with curl or Postman

## 🔮 Future Enhancements

- [ ] Database integration (PostgreSQL/MongoDB)
- [ ] User authentication and authorization
- [ ] Real-time data updates with WebSockets
- [ ] Data export functionality (CSV, PDF)
- [ ] Advanced financial analytics and charts
- [ ] Mobile app with React Native
- [ ] Integration with financial APIs
- [ ] Automated testing with Jest and Go testing
