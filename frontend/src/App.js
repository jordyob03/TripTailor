import './styles/App.css';
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Account from './pages/AccountSettings';
import UserSignup from './pages/UserSignup';
import UserProfile from './pages/InitialUserProfile';
import UserLogin from './pages/UserLogin';


function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<UserLogin />} />
        <Route path="/profile-creation" element={<UserProfile />} />
        <Route path="/sign-up" element={<UserSignup />} /> 
      </Routes>
    </Router>
  );
}

export default App;