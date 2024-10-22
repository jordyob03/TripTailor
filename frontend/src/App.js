import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import NavBar from './components/Navbar';
import Account from './pages/AccountSettings';
import ProtectedRoute from './components/Protectedroute';
import UserSignup from './pages/UserSignup';
import SearchResults from './pages/SearchResults';
import UserProfile from './pages/InitialUserProfile';
import UserLogin from './pages/UserLogin';
import Dashboard from './pages/UserDashboard';
import CreateItinerary from './pages/CreateItinerary'; 

function App() {
  const [searchResults, setSearchResults] = useState([]);
  const [searchParams, setSearchParams] = useState({ country: '', city: '' });
  const [countryErrorMessage, setCountryErrorMessage] = useState('');
  const [cityErrorMessage, setCityErrorMessage] = useState('');

  const handleSearch = (country, city) => {
    // Mock search results for demonstration
    const results = [
      {
        location: `${city}, ${country}`,
        title: 'Hiking and Hot Springs',
        description: 'Immerse yourself in the beauty of hot springs surrounded by volcanic views.',
        tags: ['Backpacker', 'Long Walks', 'Hiking'],
        image: 'https://via.placeholder.com/300x180',
      },
      {
        location: `${city}, ${country}`,
        title: 'Beautiful Evening at Bob Kerrey Bridge',
        description: 'Enjoy a stunning evening stroll across the Bob Kerrey Bridge.',
        tags: ['Short Walks', 'Historical Sites'],
        image: 'https://via.placeholder.com/300x180',
      },
    ];

    setSearchParams({ country, city });
    setSearchResults(results);
    setCountryErrorMessage(''); // Clear previous errors when search is successful
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
        <Route path="/search-results" element={<SearchResults searchResults={searchResults} searchParams={searchParams} />} />
      </Routes>
    </Router>
  );
}

export default App;
