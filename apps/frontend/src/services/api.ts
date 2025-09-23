import axios from 'axios';

// Use Vite env for base URL in production (Render or Vercel)
// Set VITE_API_URL to your backend base URL (e.g., https://your-backend.onrender.com)
const API_BASE_URL = (import.meta as any)?.env?.VITE_API_URL || '';
// Log which base URL is used (empty string means same-origin)
// This helps diagnose production issues (e.g., wrong Vercel project root)
// eslint-disable-next-line no-console
console.log('[api] Using baseURL =', API_BASE_URL || '(same-origin)');

const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor
api.interceptors.request.use(
  (config) => {
    console.log(`[api] ${config.method?.toUpperCase()} ${config.baseURL || '(same-origin)'}${config.url}`);
    return config;
  },
  (error) => {
    // eslint-disable-next-line no-console
    console.error('[api] Request error:', error?.message || error);
    return Promise.reject(error);
  }
);

// Response interceptor
api.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    // eslint-disable-next-line no-console
    console.error('[api] Response error:', error.response?.data || error.message);
    return Promise.reject(error);
  }
);

export interface Account {
  id: string;
  name: string;
  bank: string;
  account_type: string;
  balance: number;
  currency: string;
  last_updated: string;
  is_active: boolean;
}

export interface Transaction {
  id: string;
  account_id: string;
  amount: number;
  currency: string;
  type: 'debit' | 'credit' | 'transfer';
  category: string;
  description: string;
  date: string;
  status: string;
  reference?: string;
}

export interface AccountRefreshResponse {
  account_id: string;
  success: boolean;
  message: string;
  last_updated: string;
  new_balance?: number;
}

export interface ApiResponse<T> {
  success: boolean;
  message?: string;
  data?: T;
  error?: string;
}

export interface PaginatedResponse<T> {
  success: boolean;
  data: T[];
  meta: {
    total: number;
    limit: number;
    offset: number;
    pages: number;
  };
}

// API functions
export const apiService = {
  // Health check
  async checkHealth(): Promise<{ status: string; timestamp: string }> {
    const response = await api.get('/health');
    return response.data;
  },

  // Accounts
  async getAccounts(): Promise<Account[]> {
    const response = await api.get<ApiResponse<Account[]>>('/api/accounts');
    return response.data.data || [];
  },

  async getAccount(id: string): Promise<Account> {
    const response = await api.get<ApiResponse<Account>>(`/api/accounts/${id}`);
    if (!response.data.success || !response.data.data) {
      throw new Error(response.data.message || 'Account not found');
    }
    return response.data.data;
  },

  async refreshAccount(id: string): Promise<AccountRefreshResponse> {
    const response = await api.post<ApiResponse<AccountRefreshResponse>>(`/api/accounts/${id}/refresh`);
    if (!response.data.success || !response.data.data) {
      throw new Error(response.data.message || 'Failed to refresh account');
    }
    return response.data.data;
  },

  // Transactions
  async getTransactions(params?: {
    account_id?: string;
    type?: string;
    category?: string;
    status?: string;
    start_date?: string;
    end_date?: string;
    limit?: number;
    offset?: number;
  }): Promise<{ transactions: Transaction[]; meta: any }> {
    const response = await api.get<PaginatedResponse<Transaction>>('/api/transactions', { params });
    return {
      transactions: response.data.data || [],
      meta: response.data.meta || { total: 0, limit: 50, offset: 0, pages: 0 }
    };
  },

  async getTransaction(id: string): Promise<Transaction> {
    const response = await api.get<ApiResponse<Transaction>>(`/api/transactions/${id}`);
    if (!response.data.success || !response.data.data) {
      throw new Error(response.data.message || 'Transaction not found');
    }
    return response.data.data;
  },

  async getAccountTransactions(accountId: string, limit = 50): Promise<Transaction[]> {
    const response = await api.get<ApiResponse<Transaction[]>>(`/api/accounts/${accountId}/transactions`, {
      params: { limit }
    });
    return response.data.data || [];
  },
};

export default apiService;
