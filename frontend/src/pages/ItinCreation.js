import React, { useState } from 'react';
import itineraryAPI from '../api/itineraryAPI.js';

function Itinerary() {

  const [Name, setName] = useState('');
  const [City, setCity] = useState('');
  const [Country, setCountry] = useState('');
  const [Languages, setLanguages] = useState([]);
  const [Tags, setTags] = useState([]);
  const [Events, setEvents] = useState([]);



  const handleSubmit = async (e) => {
    e.preventDefault();


      console.log('Form data:', {
        Name: Name,
        City: City,
        Country: Country,
        Languages: Languages,
        Tags: Tags,
        Events: Events,
      });

      const Data = {
        Name: Name,
        City: City, 
        Country: Country,
        Languages: Languages.split(','),  // Ensure arrays are sent correctly
        Tags: Tags.split(','),
        Events: Events.split(','),
      };

      try {
        console.log('Sending request to API...');
        const response = await itineraryAPI.post('/itin-creation', Data);
        const { token } = response.data;
        console.log(response)
        console.log('we made it')  
        localStorage.setItem('token', token);  
        console.log('Location created:', response.data);

      } catch (error) {
          console.error('Error creating itinerary:', error);
      }

  };
 

  return (
    <>
        <div>

          <form onSubmit={(e) => { handleSubmit(e) }} className="form">

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
              <input
                type="text"
                placeholder="Languages"
                value={Languages}
                onChange={(e) => setLanguages(e.target.value)}
                required
                className="input"
              />
              <input
                type="text"
                placeholder="Tags"
                value={Tags}
                onChange={(e) => setTags(e.target.value)}
                required
                className="input"
              />
              <input
                type="text"
                placeholder="Events"
                value={Events}
                onChange={(e) => setEvents(e.target.value)}
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
