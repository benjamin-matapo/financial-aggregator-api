# Financial Aggregator - Frontend

A modern React + Vite + Tailwind CSS frontend application that provides a beautiful interface for managing financial accounts and transactions.

## 🚀 Features

### 📊 **Account Management**
- View all bank accounts with real-time balance information
- Account type categorization (checking, savings, credit, investment)
- Individual account refresh functionality
- Summary cards with key metrics
- Responsive table design with hover effects

### 💳 **Transaction Tracking**
- Comprehensive transaction history
- Advanced filtering by account, type, category, and status
- Real-time data updates
- Pagination support
- Beautiful transaction cards with color-coded types

### 🎨 **Modern UI/UX**
- Clean, Monzo-inspired design
- Fully responsive (mobile-first approach)
- Loading states and error handling
- Smooth animations and transitions
- Custom scrollbars and hover effects

### 🔧 **Developer Experience**
- TypeScript for type safety
- ESLint + Prettier for code quality
- Hot module replacement with Vite
- API proxy for seamless backend integration
- Comprehensive error boundaries

## 🛠️ Technology Stack

- **React 18** - Modern React with hooks and concurrent features
- **TypeScript** - Type-safe JavaScript
- **Vite** - Lightning-fast build tool and dev server
- **Tailwind CSS** - Utility-first CSS framework
- **React Router** - Client-side routing
- **Axios** - HTTP client for API communication

## 📁 Project Structure

```
src/
├── pages/                 # Page components
│   ├── AccountsPage.tsx   # Account management page
│   ├── TransactionsPage.tsx # Transaction history page
│   └── AboutPage.tsx      # Project information page
├── services/              # API service layer
│   └── api.ts            # API client and types
├── App.tsx               # Main app component with routing
├── main.tsx              # Application entry point
└── index.css             # Global styles and Tailwind imports
```

## 🚀 Getting Started

### Prerequisites
- Node.js 18+ 
- pnpm (recommended) or npm
- Backend API running on port 8080

### Installation

```bash
# Install dependencies
pnpm install

# Start development server
pnpm dev

# Build for production
pnpm build

# Preview production build
pnpm preview
```

### Development Scripts

| Script | Description |
|--------|-------------|
| `pnpm dev` | Start development server with hot reload |
| `pnpm build` | Build for production |
| `pnpm preview` | Preview production build |
| `pnpm lint` | Run ESLint |
| `pnpm type-check` | Run TypeScript type checking |

## 🔌 API Integration

The frontend communicates with the Go backend through a Vite proxy configuration:

```typescript
// vite.config.ts
export default defineConfig({
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        secure: false,
      },
    },
  },
});
```

### API Service

The `apiService` provides a clean interface for all backend communication:

```typescript
// Get all accounts
const accounts = await apiService.getAccounts();

// Get transactions with filters
const { transactions, meta } = await apiService.getTransactions({
  account_id: 'acc_001',
  type: 'debit',
  limit: 10
});

// Refresh account data
const refreshData = await apiService.refreshAccount('acc_001');
```

## 🎨 UI Components

### Account Cards
- Account type badges with color coding
- Balance display with proper formatting
- Refresh button with loading states
- Bank information and account details

### Transaction Table
- Sortable columns
- Filter controls
- Pagination support
- Status indicators
- Amount formatting with color coding

### Loading States
- Skeleton loaders
- Spinner animations
- Progress indicators
- Error boundaries

## 📱 Responsive Design

The application is fully responsive with breakpoints:

- **Mobile**: < 768px
- **Tablet**: 768px - 1024px  
- **Desktop**: > 1024px

### Mobile Features
- Collapsible navigation
- Touch-friendly buttons
- Optimized table layouts
- Swipe gestures support

## 🎯 Key Features

### 1. **Real-time Data Refresh**
```typescript
const handleRefreshAccount = async (accountId: string) => {
  const refreshData = await apiService.refreshAccount(accountId);
  // Update UI with new data
};
```

### 2. **Advanced Filtering**
```typescript
const applyFilters = async () => {
  const data = await apiService.getTransactions(filters);
  setTransactions(data.transactions);
};
```

### 3. **Error Handling**
```typescript
try {
  const data = await apiService.getAccounts();
  setAccounts(data);
} catch (error) {
  setError('Failed to fetch accounts');
}
```

## 🎨 Styling

### Tailwind Configuration
```javascript
// tailwind.config.js
module.exports = {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#eff6ff',
          500: '#3b82f6',
          600: '#2563eb',
          700: '#1d4ed8',
        },
      },
    },
  },
  plugins: [],
}
```

### Custom CSS
- Custom scrollbars
- Smooth transitions
- Loading animations
- Hover effects

## 🔧 Configuration

### Environment Variables
```bash
# .env.local
VITE_API_URL=http://localhost:8080
```

### TypeScript Configuration
- Strict mode enabled
- Path mapping support
- React JSX transform
- Modern ES features

### ESLint Configuration
- React hooks rules
- TypeScript support
- Prettier integration
- Import sorting

## 🚀 Deployment

### Vercel Deployment
```bash
# Install Vercel CLI
npm i -g vercel

# Deploy
vercel

# Production build
vercel --prod
```

### Build Output
- Optimized bundle
- Code splitting
- Asset optimization
- Source maps for debugging

## 🧪 Testing

### Manual Testing
1. Start backend: `cd ../backend && go run main.go`
2. Start frontend: `pnpm dev`
3. Open http://localhost:3000
4. Test all features:
   - Navigate between pages
   - Refresh accounts
   - Filter transactions
   - Test error states

### Browser Support
- Chrome 90+
- Firefox 88+
- Safari 14+
- Edge 90+

## 🐛 Troubleshooting

### Common Issues

1. **API Connection Failed**
   - Ensure backend is running on port 8080
   - Check Vite proxy configuration
   - Verify CORS settings

2. **Build Errors**
   - Clear node_modules and reinstall
   - Check TypeScript errors
   - Verify all imports

3. **Styling Issues**
   - Ensure Tailwind CSS is properly configured
   - Check for conflicting styles
   - Verify responsive breakpoints

## 🔮 Future Enhancements

- [ ] Dark mode support
- [ ] PWA capabilities
- [ ] Offline support
- [ ] Real-time notifications
- [ ] Advanced charts and analytics
- [ ] Export functionality
- [ ] Search capabilities
- [ ] Keyboard shortcuts

## 📄 License

This project is part of the Financial Aggregator monorepo and is licensed under the MIT License.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a pull request

---

**Built with ❤️ using React, TypeScript, and Tailwind CSS**
