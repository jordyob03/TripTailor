import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/styles.css';
import testimage from '../assets/logo-long-white.png';

function MyTravels() {
  const [selectedTab, setSelectedTab] = useState('itineraries');
  const [itineraries, setItineraries] = useState([]);
  const [trips, setTrips] = useState([]);
  const navigate = useNavigate();

  // Sample itineraries to use until the backend is ready
  const sampleItineraries = [
    {
      image: testimage,
      location: 'paris',
      title: 'Paris Adventure',
      description: 'Explore the city of lights with this 5-day itinerary.',
      tags: ['Long Walks', 'Family Friendly', 'Gourmet Dining']
    },
    {
      image: testimage,
      location: 'new york',
      title: 'New York Highlights',
      description: 'Discover the top attractions in New York City.',
      tags: ['City', 'Shopping', 'Nightlife']
    }
  ];
  
    const sampleTrips = [
      {
        boardId: 1,
        image: testimage,
        name: "Florida",
        posts: ["Post1", "Post2", "Post3"],
      },
      {
        boardId: 2,
        image: testimage,
        name: "New York",
        posts: ["Post1", "Post2", "Post3"],
      },
      {
        boardId: 3,
        image: testimage,
        name: "Toronto",
        posts: ["Post1", "Post2", "Post3"],
      },
      {
        boardId: 3,
        image: testimage,
        name: "New York",
        posts: ["Post1", "Post2", "Post3"],
      },
      {
        boardId: 4,
        image: testimage,
        name: "Toronto",
        posts: ["Post1", "Post2", "Post3"],
      },
    ];

  const handleItinClick = (id) => {
    navigate(`/itinerary/${id}`); // Navigate to the itinerary details page
  };

  const handleTripClick = (id) => {
    navigate(`/trip/${id}`); // Navigate to the itinerary details page
  };

  // // Example data fetch from the backend
  // useEffect(() => {
  //   const fetchItineraries = async () => {
  //     try {
  //       const response = await itineraryAPI.get('/itineraries'); // Might need to be changed?
  //       setItineraries(response.data);
  //     } catch (error) {
  //       console.error('Error fetching itineraries:', error);
  //     }
  //   };

  useEffect(() => {
    // Set sample data instead of fetching from the API
    setItineraries(sampleItineraries);
    setTrips(sampleTrips);

  }, []);

  const handleTabChange = (tab) => {
    setSelectedTab(tab);
  };

  return (
    <div>
      <div className="pageContainer">
        <div className="heading">
          <h2>My Travels</h2>
        </div>

        <div className="tabNavigation">
          <button className={selectedTab === 'itineraries' ? 'activeTab' : ''} onClick={() => handleTabChange('itineraries')}>
            Itineraries
          </button>
          <button className={selectedTab === 'trips' ? 'activeTab' : ''} onClick={() => handleTabChange('trips')}>
            Trips
          </button>
        </div>

        {selectedTab === 'itineraries' && (
          <div className="resultsGrid">
          {itineraries.length > 0 ? (
            itineraries.map((itinerary, index) => (
              <div
                key={index}
                className="resultCard"
                onClick={() => handleItinClick(itinerary.id)}
                style={{ cursor: 'pointer' }} 
              >
                <img src={itinerary.image} alt={itinerary.title} className="resultCardImage" />
                <div className="resultCardContent">
                  <h4 className="cardLocation">
                    {itinerary.location
                      .split(' ')
                      .map(word => word.charAt(0).toUpperCase() + word.slice(1))
                      .join(' ')}
                  </h4>
                  <h3 className="cardTitle">{itinerary.title}</h3>
                  <p className="cardDescription">{itinerary.description}</p>
                  <div className="resultTags">
                    {itinerary.tags.map((tag, i) => (
                      <span key={i} className="resultCardTag">
                        {tag}
                      </span>
                    ))}
                  </div>
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
          <div className="tripsGrid">
            {trips.length > 0 ? (
              trips.map((trip, index) => (
                <div
                  key={index}
                  className="tripsCard"
                  onClick={() => handleTripClick(trip.boardId)}
                  style={{ cursor: 'pointer' }}
                >
                  <img src={trip.image} alt={trip.name} className="tripCardImage" />
                  <div className="cardInfo">
                    <h3 className="cardTitle">{trip.name}</h3>
                    <p className="itinCount">{trip.posts.length} itineraries</p>
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
