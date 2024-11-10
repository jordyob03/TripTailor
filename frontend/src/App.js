import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import NavBar from './components/Navbar';
import AccountSettings from './pages/AccountSettings';
import ProtectedRoute from './components/Protectedroute';
import UserSignup from './pages/UserSignup';
import SearchResults from './pages/SearchResults';
import UserProfile from './pages/InitialUserProfile';
import UserLogin from './pages/UserLogin';
import HomePage from './pages/Homepage';
import MyTravels from './pages/UserDashboard';
import CreateItinerary from './pages/CreateItinerary'; 
import BoardPosts from './pages/BoardPosts';
import EditUserProfile from './pages/EditUserProfile';

function App() {
  const [searchResults, setSearchResults] = useState([]);
  const [searchParams, setSearchParams] = useState({ country: '', city: '' });
  const [countryErrorMessage, setCountryErrorMessage] = useState('');
  const [cityErrorMessage, setCityErrorMessage] = useState('');

  const handleSearch = (results, country, city) => {
    setSearchParams({ country, city });
    setSearchResults(results);
    setCountryErrorMessage('');
    setCityErrorMessage('');
  };

  return (
    <Router>
      <NavBar onSearch={handleSearch} />
      <Routes>
        <Route path="/" element={<UserLogin />} />
        <Route path="/sign-up" element={<UserSignup />} />
        <Route path="/profile-creation" element={<ProtectedRoute><UserProfile /></ProtectedRoute>} />
        <Route path="/create-itinerary" element={<ProtectedRoute><CreateItinerary /></ProtectedRoute>} />
        <Route path="/my-travels/*" element={<ProtectedRoute><MyTravels /></ProtectedRoute>} />
        <Route path="/personal-info" element={<ProtectedRoute><EditUserProfile /></ProtectedRoute>} />
        <Route path="/account-settings" element={<ProtectedRoute><AccountSettings /></ProtectedRoute>} />
        <Route path="/home-page" element={<ProtectedRoute><HomePage /></ProtectedRoute>} />
        <Route path="/search-results" element={<SearchResults searchResults={searchResults} searchParams={searchParams} isSearchPressed={true} />} />
        <Route path="/my-travels/boards/:boardId" element={<BoardPosts />} />
      </Routes>
    </Router>
  );
}

export default App;
