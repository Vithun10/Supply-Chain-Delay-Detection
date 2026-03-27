import { useState } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Navbar from './components/Navbar';
import Sidebar from './components/Sidebar';

// Existing Pages
import OwnerDashboard from './pages/OwnerDashboard';
import AdminDashboard from './pages/AdminDashboard';
import CustomerTracking from './pages/CustomerTracking';
import ShipmentsPage from './pages/ShipmentsPage';
import DelaysPage from './pages/DelaysPage';
import AnalyticsPage from './pages/AnalyticsPage';
import UploadDataset from './pages/UploadDataset';

// New Auth Pages
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';

function App() {
  const [role, setRole] = useState('Owner');
  const [isAuthenticated, setIsAuthenticated] = useState(
    !!localStorage.getItem('token') // auto login if token exists
  );

  // If not authenticated, only show Login and Register pages
  if (!isAuthenticated) {
    return (
      <Router>
        <Routes>
          <Route
            path="/login"
            element={
              <LoginPage
                setRole={setRole}
                setIsAuthenticated={setIsAuthenticated}
              />
            }
          />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="*" element={<Navigate to="/login" />} />
        </Routes>
      </Router>
    );
  }

  // Authenticated — render full app with existing routes
  const renderRoutes = () => {
    switch (role) {
      case 'Owner':
        return (
          <Routes>
            <Route path="/" element={<OwnerDashboard />} />
            <Route path="/shipments" element={<ShipmentsPage />} />
            <Route path="/delays" element={<DelaysPage type="all" />} />
            <Route path="/delays/high-risk" element={<DelaysPage type="high-risk" />} />
            <Route path="/analytics" element={<AnalyticsPage />} />
            <Route path="/upload-dataset" element={<UploadDataset />} />
            <Route path="*" element={<Navigate to="/" />} />
          </Routes>
        );
      case 'Admin':
        return (
          <Routes>
            <Route path="/" element={<AdminDashboard />} />
            <Route path="/shipments" element={<ShipmentsPage />} />
            <Route path="/delays" element={<DelaysPage type="all" />} />
            <Route path="/delays/high-risk" element={<DelaysPage type="high-risk" />} />
            <Route path="/upload-dataset" element={<UploadDataset />} />
            <Route path="*" element={<Navigate to="/" />} />
          </Routes>
        );
      case 'Customer':
        return (
          <Routes>
            <Route path="/" element={<CustomerTracking />} />
            <Route path="*" element={<Navigate to="/" />} />
          </Routes>
        );
      default:
        return null;
    }
  };

  return (
    <Router>
      <div className="flex flex-col h-screen bg-slate-50">
        <Navbar
          role={role}
          setRole={setRole}
          setIsAuthenticated={setIsAuthenticated}
        />
        <div className="flex flex-1 overflow-hidden">
          {role !== 'Customer' && <Sidebar role={role} />}
          <main className="flex-1 overflow-x-hidden overflow-y-auto bg-slate-50 p-6">
            {renderRoutes()}
          </main>
        </div>
      </div>
    </Router>
  );
}

export default App;