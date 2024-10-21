import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import NavBar from '../components/Navbar.js';
import '../styles/styles.css';

function MyDashboard() {
  const [selectedTab, setSelectedTab] = useState('itineraries');
  const [itineraries, setItineraries] = useState([]);
  const [trips, setTrips] = useState([]);
  const [menuOpen, setMenuOpen] = useState(false);
  const navigate = useNavigate();

  const handleTabChange = (tab) => {
    setSelectedTab(tab);
  };

  const handleCreateItinerary = () => {
    navigate('/create-itinerary');
  };

  const toggleMenu = () => {
    setMenuOpen(!menuOpen);
  };

  const handleLogout = () => {
    localStorage.clear(); 
    navigate('/'); 
  };

  return (
    <div>
      <NavBar />

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

      <div className="pageContainer">
      {/* Main Heading */}
      <div className="heading">
        <h2>My Travels</h2>
      </div>

      {/* Tab Navigation */}
      <div className="tabNavigation">
        <button className={selectedTab === 'itineraries' ? 'activeTab' : ''} onClick={() => handleTabChange('itineraries')}>
          Itineraries
        </button>
        <button className={selectedTab === 'trips' ? 'activeTab' : ''} onClick={() => handleTabChange('trips')}>
          Trips
        </button>
      </div>

      {/* Dashboard Content */}
        {selectedTab === 'itineraries' && (
          <div className="itinerariesList">
            {itineraries.length > 0 ? (
              itineraries.map((itinerary, index) => (
                <div key={index} className="itineraryCard">
                  <img src={itinerary.image} alt={itinerary.title} className="itineraryImage" />
                  <div className="itineraryInfo">
                    <h3>{itinerary.title}</h3>
                    <p>{itinerary.description}</p>
                    <span>{itinerary.location}</span>
                  </div>
                </div>
              ))
            ) : (
              <div className="noItinerariesMessage">
                No itineraries found. Create a new one to get started!
              </div>
            )}
          </div>
        )}

        {selectedTab === 'trips' && (
          <div className="tripsList">
            {trips.length > 0 ? (
              trips.map((trip, index) => (
                <div key={index} className="tripCard">
                  <img src={trip.image} alt={trip.title} className="tripImage" />
                  <div className="tripInfo">
                    <h3>{trip.title}</h3>
                    <p>{trip.description}</p>
                    <span>{trip.location}</span>
                  </div>
                </div>
              ))
            ) : (
              <div className="noTripsMessage">
                No trips saved. Add some itineraries to your trips to start planning!
              </div>
            )}
          </div>
        )}
    </div>
    </div>

  );
}

export default MyDashboard;
