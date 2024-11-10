import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/styles.css';
import itineraryAPI from '../api/itineraryAPI.js';

function MyTravels() {
  const [selectedTab, setSelectedTab] = useState('itineraries');
  const [itineraries, setItineraries] = useState([]);
  const [trips, setTrips] = useState([]);
  const [menuOpen, setMenuOpen] = useState(false);
  const navigate = useNavigate();

  const handleTabChange = (tab) => {
    setSelectedTab(tab);
  };

  useEffect(() => {
    const fetchItineraries = async () => {
      if (selectedTab === 'itineraries') {
        try {
          const data = {
            Username: localStorage.getItem('username'),
          };
          const response = await itineraryAPI.post('/get-user-itins', data);
          console.log('Received', response.data);
          setItineraries(response.data.itineraries);
        } catch (error) {
          console.error('Error fetching itineraries:', error);
        }
      }
    };
  
    fetchItineraries();
  }, [selectedTab]);

  return (
    <div>
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
                    <h3>{itinerary.name}</h3>
                    <p>{itinerary.city}</p>
                    <p>{itinerary.country}</p>
                    <p>{itinerary.price}</p>
                    <p>{itinerary.description}</p>
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

export default MyTravels;