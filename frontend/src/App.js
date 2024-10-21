import './styles/App.css';
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Account from './pages/AccountSettings';
import ProtectedRoute from './components/Protectedroute';
import UserSignup from './pages/UserSignup';
import UserProfile from './pages/InitialUserProfile';
import UserLogin from './pages/UserLogin';
import Search from './pages/HomeSearch';
import Dashboard from './pages/UserDashboard';
import CreateItinerary from './pages/CreateItinerary'; 

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<UserLogin />} />
        <Route path="/sign-up" element={<UserSignup />} /> 
        <Route path="/profile-creation" element={<ProtectedRoute> <UserProfile /> </ProtectedRoute>} />
        <Route path="/dashboard" element={<ProtectedRoute> <Dashboard /> </ProtectedRoute>} /> 
        <Route path="/create-itinerary" element={<ProtectedRoute> <CreateItinerary /> </ProtectedRoute>} /> 
      </Routes>
    </Router>
  );
}

export default App;
