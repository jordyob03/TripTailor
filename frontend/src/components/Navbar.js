import React, { useState, useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import navBarLogo from '../assets/logo-long-transparent.png';
import '../styles/styles.css';
import { faSearch, faSignOutAlt } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import searchAPI from '../api/searchAPI'; 

function NavBar({ onSearch }) {
  const [SearchValue, setSearchValue] = useState('');
  const [Price, setPrice] = useState(0);
  const [SearchValueErrorMessage, setSearchValueErrorMessage] = useState('');
  const [PriceErrorMessage, setPriceErrorMessage] = useState('');
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

  const handlePriceChange = (e) => {
    const value = parseFloat(e.target.value);
    if (isNaN(value)) {
      setPrice(0); // Default to 0 if not a number
    } else {
      setPrice(value);
    }
  };

  const handleLogout = () => {
    localStorage.clear();
    navigate('/');
  };

  const handleSearch = async (e) => {
    e.preventDefault();
    let hasError = false;

    setSearchValueErrorMessage('');
    setPriceErrorMessage('');
    setErrorMessage('');

    if (!SearchValue) {
      setSearchValueErrorMessage('Please enter a SearchValue.');
      hasError = true;
    }
    if (!Price || isNaN(Price) || Price <= 0) {
      setPriceErrorMessage('Please enter a valid positive Price.');
      hasError = true;
    }

    if (!hasError) {
      const searchData = { SearchValue, Price };

      try {
        console.log("Search API sent:", searchData);
        const response = await searchAPI.get('/search', {
          params: {
            searchValue: SearchValue,
            price: parseFloat(Price),
          },
        });
        
        console.log('API response:', response);

        if (onSearch) {
          onSearch(response.data, SearchValue, Price);
        }

        navigate('/search-results'); 
      } catch (error) {
        console.error('Search API error:', error); // Log full error details
        if (error.response && error.response.data && error.response.data.error) {
          setErrorMessage(error.response.data.error);
        } else {
          setErrorMessage('An unexpected error occurred. Please try again.');
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
              <label className="inputLabel">SearchValue</label>
              <input
                type="text"
                placeholder="Enter keyword (e.g., 'Lahore')"
                value={SearchValue}
                onChange={(e) => setSearchValue(e.target.value)}
                className="inputField"
              />
            </div>
            <div className="inputFieldContainer">
              <label className="inputLabel">Price</label>
              <input
              type="number"
              step="0.01"
              placeholder="Enter Price"
              value={Price}
              onChange={handlePriceChange}
              className="inputField"
              />
            </div>
            <button onClick={handleSearch} className="searchButton">
              <FontAwesomeIcon icon={faSearch} /> Search
            </button>
          </div>
          <div className="errorMessagesContainer">
            {SearchValueErrorMessage && <div className="errorMessageSB">{SearchValueErrorMessage}</div>}
            {PriceErrorMessage && <div className="errorMessageSB">{PriceErrorMessage}</div>}
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
