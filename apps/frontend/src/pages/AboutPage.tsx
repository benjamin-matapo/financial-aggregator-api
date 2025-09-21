export function AboutPage() {
  return (
    <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div className="bg-white shadow-lg rounded-lg overflow-hidden">
        <div className="px-6 py-8 sm:px-8">
          <div className="text-center mb-8">
            <div className="mx-auto h-16 w-16 bg-blue-100 rounded-full flex items-center justify-center mb-4">
              <span className="text-3xl">💰</span>
            </div>
            <h1 className="text-3xl font-bold text-gray-900 mb-2">
              Financial Aggregator
            </h1>
            <p className="text-lg text-gray-600">
              A Monzo-inspired financial management platform
            </p>
          </div>

          <div className="prose prose-lg max-w-none">
            <div className="bg-blue-50 border-l-4 border-blue-400 p-4 mb-8">
              <div className="flex">
                <div className="flex-shrink-0">
                  <svg className="h-5 w-5 text-blue-400" viewBox="0 0 20 20" fill="currentColor">
                    <path fillRule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clipRule="evenodd" />
                  </svg>
                </div>
                <div className="ml-3">
                  <p className="text-sm text-blue-700">
                    <strong>Note:</strong> This is a demo application showcasing modern full-stack development with React, Go, and Vercel deployment.
                  </p>
                </div>
              </div>
            </div>

            <h2 className="text-2xl font-bold text-gray-900 mb-4">About This Project</h2>
            <p className="text-gray-700 mb-6">
              The Financial Aggregator is a modern, full-stack web application inspired by Monzo's innovative approach to banking and financial management. 
              Built as a demonstration of current web development best practices, this project showcases how to create a production-ready application 
              that can aggregate and manage financial data from multiple sources.
            </p>

            <h3 className="text-xl font-semibold text-gray-900 mb-3">Key Features</h3>
            <ul className="list-disc list-inside text-gray-700 mb-6 space-y-2">
              <li><strong>Account Management:</strong> View and manage multiple bank accounts with real-time balance updates</li>
              <li><strong>Transaction Tracking:</strong> Comprehensive transaction history with advanced filtering and categorization</li>
              <li><strong>Real-time Refresh:</strong> Simulate live data updates from external financial sources</li>
              <li><strong>Responsive Design:</strong> Beautiful, mobile-first interface that works on all devices</li>
              <li><strong>Modern Architecture:</strong> Clean separation of concerns with microservices-ready design</li>
            </ul>

            <h3 className="text-xl font-semibold text-gray-900 mb-3">Technology Stack</h3>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
              <div className="bg-gray-50 rounded-lg p-4">
                <h4 className="font-semibold text-gray-900 mb-2">Frontend</h4>
                <ul className="text-sm text-gray-700 space-y-1">
                  <li>• React 18 with TypeScript</li>
                  <li>• Vite for fast development</li>
                  <li>• Tailwind CSS for styling</li>
                  <li>• React Router for navigation</li>
                  <li>• Axios for API communication</li>
                </ul>
              </div>
              <div className="bg-gray-50 rounded-lg p-4">
                <h4 className="font-semibold text-gray-900 mb-2">Backend</h4>
                <ul className="text-sm text-gray-700 space-y-1">
                  <li>• Go with Chi router</li>
                  <li>• RESTful API design</li>
                  <li>• Clean architecture pattern</li>
                  <li>• Comprehensive testing</li>
                  <li>• Docker containerization</li>
                </ul>
              </div>
            </div>

            <h3 className="text-xl font-semibold text-gray-900 mb-3">Architecture Highlights</h3>
            <div className="bg-gray-50 rounded-lg p-6 mb-6">
              <div className="grid grid-cols-1 md:grid-cols-3 gap-4 text-center">
                <div>
                  <div className="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-2">
                    <span className="text-blue-600 font-bold">M</span>
                  </div>
                  <h4 className="font-semibold text-gray-900">Monorepo</h4>
                  <p className="text-sm text-gray-600">pnpm workspaces for efficient dependency management</p>
                </div>
                <div>
                  <div className="w-12 h-12 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-2">
                    <span className="text-green-600 font-bold">A</span>
                  </div>
                  <h4 className="font-semibold text-gray-900">API-First</h4>
                  <p className="text-sm text-gray-600">RESTful design with proper error handling</p>
                </div>
                <div>
                  <div className="w-12 h-12 bg-purple-100 rounded-full flex items-center justify-center mx-auto mb-2">
                    <span className="text-purple-600 font-bold">D</span>
                  </div>
                  <h4 className="font-semibold text-gray-900">Deployable</h4>
                  <p className="text-sm text-gray-600">Ready for Vercel and container platforms</p>
                </div>
              </div>
            </div>

            <h3 className="text-xl font-semibold text-gray-900 mb-3">Development Features</h3>
            <ul className="list-disc list-inside text-gray-700 mb-6 space-y-2">
              <li><strong>Type Safety:</strong> Full TypeScript implementation across frontend and backend</li>
              <li><strong>Code Quality:</strong> ESLint, Prettier, and comprehensive testing</li>
              <li><strong>Hot Reload:</strong> Fast development with Vite and Go's air tool</li>
              <li><strong>Error Handling:</strong> Graceful error states and user feedback</li>
              <li><strong>Loading States:</strong> Smooth user experience with loading indicators</li>
            </ul>

            <h3 className="text-xl font-semibold text-gray-900 mb-3">Getting Started</h3>
            <div className="bg-gray-900 rounded-lg p-4 mb-6">
              <pre className="text-green-400 text-sm overflow-x-auto">
{`# Install dependencies
pnpm install

# Start development servers
pnpm dev

# Or start individually
pnpm frontend:dev  # React app on :3000
pnpm backend:dev   # Go API on :8080`}
              </pre>
            </div>

            <h3 className="text-xl font-semibold text-gray-900 mb-3">API Endpoints</h3>
            <div className="bg-gray-50 rounded-lg p-4 mb-6">
              <div className="space-y-2 text-sm">
                <div className="flex items-center space-x-4">
                  <span className="bg-green-100 text-green-800 px-2 py-1 rounded text-xs font-mono">GET</span>
                  <span className="font-mono">/api/accounts</span>
                  <span className="text-gray-600">List all accounts</span>
                </div>
                <div className="flex items-center space-x-4">
                  <span className="bg-green-100 text-green-800 px-2 py-1 rounded text-xs font-mono">GET</span>
                  <span className="font-mono">/api/transactions</span>
                  <span className="text-gray-600">List transactions with filters</span>
                </div>
                <div className="flex items-center space-x-4">
                  <span className="bg-blue-100 text-blue-800 px-2 py-1 rounded text-xs font-mono">POST</span>
                  <span className="font-mono">/api/accounts/:id/refresh</span>
                  <span className="text-gray-600">Refresh account data</span>
                </div>
              </div>
            </div>

            <div className="border-t border-gray-200 pt-6">
              <h3 className="text-xl font-semibold text-gray-900 mb-3">Future Enhancements</h3>
              <p className="text-gray-700 mb-4">
                This project serves as a foundation for more advanced financial applications. 
                Potential enhancements include:
              </p>
              <ul className="list-disc list-inside text-gray-700 space-y-1">
                <li>Real database integration (PostgreSQL, Cassandra)</li>
                <li>User authentication and authorization</li>
                <li>Real-time data synchronization</li>
                <li>Advanced analytics and reporting</li>
                <li>Mobile app development</li>
                <li>Integration with real financial APIs</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
