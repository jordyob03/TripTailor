import React, { useState } from 'react';
import itineraryAPI from '../api/itineraryAPI.js';

function Itinerary() {

  const [location, setLocation] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();



      const Data = {
        location: location, 

      }

      try {
        const response = await itineraryAPI.post('/itin-creation', Data);
        const { token } = response.data;  
        localStorage.setItem('token', token);  
        console.log('Location created:', response.data);

      } catch (error) {

      }

  };
 

  return (
    <>
        <div>

          <form onSubmit={handleSubmit} className="form">
              <input
                type="text"
                placeholder="Location"
                value={location}
                onChange={(e) => setLocation(e.target.value)}
                required
                className="input"
              />

              <button
                type="submit"

              >
                Continue
              </button>

          </form>
        </div>
    </>
  );
}

export default Itinerary;
