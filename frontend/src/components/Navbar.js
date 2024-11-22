import React, { useState, useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import navBarLogo from '../assets/logo-long-transparent.png';
import '../styles/styles.css';
import { faSearch, faSignOutAlt } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import searchAPI from '../api/searchAPI'; 

function NavBar({ onSearch }) {
  const [country, setCountry] = useState('');
  const [city, setCity] = useState('');
  const [countryErrorMessage, setCountryErrorMessage] = useState('');
  const [cityErrorMessage, setCityErrorMessage] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const [menuOpen, setMenuOpen] = useState(false);

  const location = useLocation();
  const navigate = useNavigate();

  const noProfile = ['/', '/sign-up'];
  const noLogOut = ['/', '/sign-up'];
  const noCreateItinerary = ['/', '/sign-up'];
  const noSearchBar = ['/', '/sign-up', '/profile-creation'];

  const toggleMenu = () => setMenuOpen((prev) => !prev);

  const closeMenu = () => setMenuOpen(false);

  useEffect(() => {
    // Close menu on route change
    closeMenu();
  }, [location.pathname]);

  const handleLogout = () => {
    localStorage.clear();
    navigate('/');
  };

  const handleSearch = async (e) => {
    e.preventDefault();
    let hasError = false;

    setCountryErrorMessage('');
    setCityErrorMessage('');
    setErrorMessage('');

    if (!country) {
      setCountryErrorMessage('Please enter a country.');
      hasError = true;
    }
    if (!city) {
      setCityErrorMessage('Please enter a city.');
      hasError = true;
    }

    if (!hasError) {
      const searchData = { country, city };

      try {
        console.log("Search API sent:", searchData);
        const response = await searchAPI.get('/search', {
          params: searchData,
        });
        console.log('API response:', response);

        if (onSearch) {
          onSearch(response.data, country, city);
        }

        navigate('/search-results'); 
      } catch (error) {
        if (error.response && error.response.data) {
          setErrorMessage(error.response.data.error);
        } else {
          setErrorMessage('Search Failed');
        }
      }
    }
  };

  return (
    <nav className="navBar">
      <img
        src={navBarLogo}
        alt="Trip Tailor Logo"
        className="navBarLogo"
        onClick={() => navigate('/home-page')}
        style={{ cursor: 'pointer' }}
      />
      {!noSearchBar.includes(location.pathname) && (
        <div className="searchBarContainer">
          <div className="inputGroupNav">
            <div className="inputFieldContainer">
              <label className="inputLabel">Country</label>
              <input
                type="text"
                placeholder="Enter Country"
                value={country}
                onChange={(e) => setCountry(e.target.value)}
                className="inputField"
              />
            </div>
            <div className="inputFieldContainer">
              <label className="inputLabel">City</label>
              <input
                type="text"
                placeholder="Enter City"
                value={city}
                onChange={(e) => setCity(e.target.value)}
                className="inputField"
              />
            </div>
            <button onClick={handleSearch} className="searchButton">
              <FontAwesomeIcon icon={faSearch} /> Search
            </button>
          </div>
          <div className="errorMessagesContainer">
            {countryErrorMessage && <div className="errorMessageSB">{countryErrorMessage}</div>}
            {cityErrorMessage && <div className="errorMessageSB">{cityErrorMessage}</div>}
            {errorMessage && <div className="errorMessageSB">{errorMessage}</div>}
          </div>
        </div>
      )}
      <div className="buttonsContainer">
        {!noCreateItinerary.includes(location.pathname) && (
          <button className="createItineraryButton" onClick={() => navigate('/create-itinerary')}>
            Create Itinerary
          </button>
        )}
        {!noLogOut.includes(location.pathname) && (
          <button className="logoutButton" onClick={handleLogout}>
            <FontAwesomeIcon icon={faSignOutAlt} /> Log Out
          </button>
        )}
        {!noProfile.includes(location.pathname) && (
          <button className="profileButton" onClick={toggleMenu}>
            <i className="fas fa-bars" style={{ fontSize: '16px', color: '#00509e', marginRight: '15px' }}></i>
            <i className="fa-regular fa-user" style={{ fontSize: '16px', color: '#00509e' }}></i>
          </button>
        )}
      </div>

      {menuOpen && (
        <div className="dropdownMenu">
          <ul>
            <li onClick={() => { navigate('/home-page'); closeMenu(); }}>Home</li>
            <li onClick={() => { navigate('/my-travels/itineraries'); closeMenu(); }}>My Travels</li>
            <li onClick={() => { navigate('/account-settings'); closeMenu(); }}>Account Settings</li>
          </ul>
        </div>
      )}
    </nav>
  );
}

export default NavBar;
