import React, { useState, useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import navBarLogo from '../assets/logo-long-transparent.png';
import '../styles/styles.css';
import { faSearch, faSignOutAlt } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import searchAPI from '../api/searchAPI'; 

function NavBar({ onSearch }) {
  const [SearchValue, setSearchValue] = useState('');
  const [Price, setPrice] = useState('');
  const [SearchValueErrorMessage, setSearchValueErrorMessage] = useState('');
  const [PriceErrorMessage, setPriceErrorMessage] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const [menuOpen, setMenuOpen] = useState(false);
  const [typingTimeout, setTypingTimeout] = useState(null);

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
    const value = e.target.value;
  
    // Allow only numbers, decimal points, and empty values
    if (/^\d*\.?\d*$/.test(value)) {
      setPrice(value); // Keep value as string for precise user input handling
      handleSearchDebounced(SearchValue, value); // Call the search function with updated Price
    }
  };

  const handleLogout = () => {
    localStorage.clear();
    navigate('/');
    window.location.reload();
  };

  const handleSearchDebounced = (searchFieldEntry, price) => {

    if (typingTimeout) {
      clearTimeout(typingTimeout); // Clear any existing timeout
    }

    const priceTemp = (!price || isNaN(price) || price <= 0) ? 999999 : price;
    const numericPrice = parseFloat(priceTemp);
    const searchTemp = (!searchFieldEntry || searchFieldEntry <= '') ? ' ' : searchFieldEntry;
    const searchValue = searchTemp;

    const timeout = setTimeout(async () => {
      let hasError = false;
  
      setSearchValueErrorMessage('');
      setPriceErrorMessage('');
      setErrorMessage('');
  
  
      if (!hasError) {
        const searchData = { SearchValue: searchValue, Price: numericPrice };
  
        try {
          console.log("Search API sent:", searchData);
          const response = await searchAPI.get('/search', {
            params: {
              searchValue,
              price: numericPrice,
            },
          });
  
          console.log('API response:', response);
  
          if (onSearch) {
            onSearch(response.data, searchValue, price);
          }
  
          // Optionally navigate to results
          if (location.pathname !== '/search-results') {
            navigate('/search-results');
            window.location.reload();
          }
        } catch (error) {
          console.error('Search API error:', error); // Log full error details
          if (error.response && error.response.data && error.response.data.error) {
            setErrorMessage(error.response.data.error);
          } else {
            setErrorMessage('An unexpected error occurred. Please try again.');
          }
        }
      }
    }, 100); // Delay API call by 500ms
  
    setTypingTimeout(timeout); // Store the timeout ID
  };
  
  

  return (
    <nav className="navBar">
      <img
        src={navBarLogo}
        alt="Trip Tailor Logo"
        className="navBarLogo"
        onClick={() => {navigate('/home-page'); window.location.reload();}}
        style={{ cursor: 'pointer' }}
      />
      {!noSearchBar.includes(location.pathname) && (
        <div className="searchBarContainer">
          <div className="inputGroupNav">
            <div className="inputFieldContainer">
              <label className="inputLabel">Search</label>
              <input
                type="text"
                placeholder="Enter keyword (e.g., 'Lahore')"
                value={SearchValue}
                onChange={(e) => {
                  const newValue = e.target.value;
                  setSearchValue(newValue);
                  handleSearchDebounced(newValue, Price);
                }}
                className="inputField"
                style={{ width: '300px' }}
              />
            </div>
            <div className="inputFieldContainer">
              <label className="inputLabel">Price</label>
              <input
                type="text"
                placeholder="Enter price"
                value={Price}
                onChange={handlePriceChange}
                className="inputField"
                style={{ width: '150px' }}
              />
            </div>
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
          <button className="createItineraryButton" onClick={() => {navigate('/create-itinerary'); window.location.reload();}}>
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
            <li onClick={() => { navigate('/home-page'); window.location.reload(); closeMenu(); }}>Home</li>
            <li onClick={() => { navigate('/my-travels/itineraries'); window.location.reload(); closeMenu(); }}>My Travels</li>
            <li onClick={() => { navigate('/personal-info'); window.location.reload(); closeMenu(); }}>Account Settings</li>
          </ul>
        </div>
      )}
    </nav>
  );
}

export default NavBar;
