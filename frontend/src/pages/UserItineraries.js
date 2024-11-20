import React, { useEffect, useState } from 'react';
import itineraryAPI from '../api/itineraryAPI.js';
import ItineraryGrid from '../components/ItineraryGrid.js'; 
import '../styles/styles.css';


function Itineraries() {
  const [itineraries, setItineraries] = useState([]);

  const fetchItins = async () => {
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
  };
  
  useEffect(() => {
    fetchItins();
  }, []);


  return (
    <div className="results">
      <ItineraryGrid itineraries={itineraries} showSaveButton={true} />
    </div>
  );
}

export default Itineraries;
