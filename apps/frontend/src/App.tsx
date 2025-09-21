import { BrowserRouter as Router, Routes, Route, Link, useLocation } from 'react-router-dom';
import { AccountsPage } from './pages/AccountsPage';
import { TransactionsPage } from './pages/TransactionsPage';
import { AboutPage } from './pages/AboutPage';

function Navbar() {
  const location = useLocation();

  const isActive = (path: string) => {
    return location.pathname === path;
  };

  return (
    <nav className="bg-white shadow-lg border-b border-gray-200">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex items-center">
            <Link to="/" className="flex-shrink-0 flex items-center">
              <h1 className="text-2xl font-bold text-gray-900">
                💰 Financial Aggregator
              </h1>
            </Link>
          </div>
          <div className="flex items-center space-x-8">
            <Link
              to="/accounts"
              className={`px-3 py-2 rounded-md text-sm font-medium transition-colors ${
                isActive('/accounts')
                  ? 'bg-blue-100 text-blue-700'
                  : 'text-gray-500 hover:text-gray-700 hover:bg-gray-100'
              }`}
            >
              Accounts
            </Link>
            <Link
              to="/transactions"
              className={`px-3 py-2 rounded-md text-sm font-medium transition-colors ${
                isActive('/transactions')
                  ? 'bg-blue-100 text-blue-700'
                  : 'text-gray-500 hover:text-gray-700 hover:bg-gray-100'
              }`}
            >
              Transactions
            </Link>
            <Link
              to="/about"
              className={`px-3 py-2 rounded-md text-sm font-medium transition-colors ${
                isActive('/about')
                  ? 'bg-blue-100 text-blue-700'
                  : 'text-gray-500 hover:text-gray-700 hover:bg-gray-100'
              }`}
            >
              About
            </Link>
          </div>
        </div>
      </div>
    </nav>
  );
}

function App() {
  return (
    <Router>
      <div className="min-h-screen bg-gray-50">
        <Navbar />
        <main>
          <Routes>
            <Route path="/" element={<AccountsPage />} />
            <Route path="/accounts" element={<AccountsPage />} />
            <Route path="/transactions" element={<TransactionsPage />} />
            <Route path="/about" element={<AboutPage />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;
