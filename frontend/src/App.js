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
import MyBoards from './pages/UserBoards';
import BoardPosts from './pages/BoardPosts';

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
        <Route path="/my-travels" element={<ProtectedRoute><MyTravels /></ProtectedRoute>} />
        <Route path="/account-settings" element={<ProtectedRoute><AccountSettings /></ProtectedRoute>} />
        <Route path="/home-page" element={<ProtectedRoute><HomePage /></ProtectedRoute>} />
        <Route path="/search-results" element={<SearchResults searchResults={searchResults} searchParams={searchParams} isSearchPressed={true} />} />
        <Route path="/itincreation" element={<CreateItinerary />} />
        <Route path="/my-boards" element={<ProtectedRoute><MyBoards /></ProtectedRoute>} />
        <Route path="/my-boards/:board_id" element={<ProtectedRoute><BoardPosts /></ProtectedRoute>} />
      </Routes>
    </Router>
  );
}

export default App;
