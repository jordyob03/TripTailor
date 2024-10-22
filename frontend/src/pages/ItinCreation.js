import React, { useState } from 'react';
import itineraryAPI from '../api/itineraryAPI.js';

function Itinerary() {

  const [Name, setName] = useState('');
  const [City, setCity] = useState('');
  const [Country, setCountry] = useState('');
  const [Languages, setLaguages] = useState([]);
  const [Tags, setTags] = useState([]);
  const [Events, setEvents] = useState([]);



  const handleSubmit = async (e) => {
    e.preventDefault();



      const Data = {
        Name: Name,
        City: City, 
        Country: Country
      }
      console.log(Data);

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
                placeholder="Name"
                value={Name}
                onChange={(e) => setName(e.target.value)}
                required
                className="input"
              />
              <input
                type="text"
                placeholder="City"
                value={City}
                onChange={(e) => setCity(e.target.value)}
                required
                className="input"
              />
              <input
                type="text"
                placeholder="Country"
                value={Country}
                onChange={(e) => setCountry(e.target.value)}
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
