import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useLocation } from 'react-router-dom';
import navBarLogo from '../assets/logo-long-transparent.png'; 
import '../styles/styles.css'; 

function NavBar() {
  const [menuOpen, setMenuOpen] = useState(false);
  const navigate = useNavigate();
  const location = useLocation();
  console.log(location.pathname !== ('/' || '/sign-up' ));
  const toggleMenu = () => {
    setMenuOpen(!menuOpen);
  };

  const handleLogout = () => {
    localStorage.clear();
    navigate('/');
  };

  const noProfile = ['/', '/sign-up'];
  const noLogOut = ['/', '/sign-up'];
  const noCreateItinerary = ['/', '/sign-up', '/profile-creation'];

  return (
    <nav className="navBar">
      <img src={navBarLogo} alt="Trip Tailor Logo" className="navBarLogo" />

      <div className="buttonsContainer">
        {!noCreateItinerary.includes(location.pathname) && (
        <button className="createItineraryButton" onClick={() => navigate('/create-itinerary')}>
          Create Itinerary
        </button>
        )}
        {!noLogOut.includes(location.pathname) && (
        <button className="logoutButton" onClick={handleLogout}>
          <i className="fas fa-sign-out-alt"></i>
          Log Out
        </button>
        )}
        {!noProfile.includes(location.pathname) && (
        <button className="profileButton" onClick={toggleMenu}>
            <i className="fas fa-bars" style={{ fontSize: '16px', color: '#00509e', marginRight: '15px' }}></i>
            <i className="fa-regular fa-user" style={{ fontSize: '16px', color: '#00509e' }}></i>
          </button>
        )}
      </div>
                
      {/* Dropdown Menu */}
      {menuOpen && (
        <div className="dropdownMenu">
          <ul>
            <li onClick={() => navigate('/profile-creation')}>Profile</li>
            <li onClick={() => navigate('/account')}>Account Settings</li>
            <li onClick={() => navigate('/dashboard')}>My Itineraries</li>
            <li onClick={() => navigate('/')}>Home</li>
          </ul>
        </div>
      )}
    </nav>
  );
}

export default NavBar;