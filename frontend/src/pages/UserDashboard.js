import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function MyTravels() {
  const [selectedTab, setSelectedTab] = useState('itineraries');
  const [itineraries, setItineraries] = useState([]);  // Placeholder for future backend data
  const [trips, setTrips] = useState([]);  // Placeholder for future backend data

  const navigate = useNavigate();

  const handleTabChange = (tab) => {
    setSelectedTab(tab);
  };

  return (
    <div style={{ paddingTop: '100px', display: 'flex', flexDirection: 'column', alignItems: 'center', width: '100%', backgroundColor: '#fff' }}>
      {/* Main Heading */}
      <h2 style={{ color: '#00509e', fontSize: '24px', fontWeight: 'bold' }}>My Travels</h2>

      {/* Tab Navigation */}
      <div style={{ display: 'flex', justifyContent: 'center', gap: '50px', marginBottom: '30px' }}>
        <button
          onClick={() => handleTabChange('itineraries')}
          style={{
            fontSize: '20px',
            fontWeight: 'bold',
            border: 'none',
            background: 'none',
            color: selectedTab === 'itineraries' ? '#00509e' : '#333',
            borderBottom: selectedTab === 'itineraries' ? '3px solid #00509e' : 'none',
            cursor: 'pointer'
          }}
        >
          Itineraries
        </button>
        <button
          onClick={() => handleTabChange('trips')}
          style={{
            fontSize: '20px',
            fontWeight: 'bold',
            border: 'none',
            background: 'none',
            color: selectedTab === 'trips' ? '#00509e' : '#333',
            borderBottom: selectedTab === 'trips' ? '3px solid #00509e' : 'none',
            cursor: 'pointer'
          }}
        >
          Trips
        </button>
      </div>

      {/* Itineraries Tab Content */}
      {selectedTab === 'itineraries' && (
        <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(300px, 1fr))', gap: '20px', width: '100%', maxWidth: '1200px', padding: '0 20px' }}>
          {itineraries.length > 0 ? (
            itineraries.map((itinerary, index) => (
              <div key={index} style={{ backgroundColor: 'white', borderRadius: '12px', boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)', overflow: 'hidden' }}>
                <img
                  src={itinerary.image}
                  alt={itinerary.title}
                  style={{ width: '100%', height: '180px', objectFit: 'cover' }}
                />
                <div style={{ padding: '15px', textAlign: 'left' }}>
                  <span style={{ color: '#000', fontSize: '14px', fontFamily: 'Red Hat Display, sans-serif' }}>{itinerary.location}</span>
                  <h3 style={{ color: '#000', fontSize: '18px', margin: '10px 0', fontFamily: 'Red Hat Display, sans-serif' }}>{itinerary.title}</h3>
                  <p style={{ color: '#555', fontSize: '14px', marginBottom: '10px' }}>{itinerary.description}</p>
                  <div style={{ display: 'flex', flexWrap: 'wrap', gap: '5px' }}>
                    {itinerary.tags.map((tag, i) => (
                      <span key={i} style={{ padding: '5px 10px', backgroundColor: '#E1EFFF', borderRadius: '50px', fontSize: '12px', color: '#00509e' }}>
                        {tag}
                      </span>
                    ))}
                  </div>
                </div>
              </div>
            ))
          ) : (
            <div style={{ backgroundColor: '#f0f8ff', padding: '15px 25px', borderRadius: '20px', textAlign: 'center', fontSize: '18px', fontWeight: 'bold', color: '#00509e', boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)' }}>
              No itineraries found. Create a new one to get started!
            </div>
          )}
        </div>
      )}

      {/* Trips Tab Content */}
      {selectedTab === 'trips' && (
        <div style={{ textAlign: 'center', color: '#666', fontSize: '16px', padding: '20px' }}>
          No trips to display yet.
        </div>
      )}
    </div>
  );
}

export default MyTravels;
