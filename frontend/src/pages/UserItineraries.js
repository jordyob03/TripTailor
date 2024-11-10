import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import itineraryAPI from '../api/itineraryAPI.js';
import '../styles/styles.css';

const username = localStorage.getItem('username');

function Itineraries() {
  const [itineraries, setItineraries] = useState([]);
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();



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
  const handleItinClick = (id) => {
    navigate(`/itinerary/${id}`);
  };

  useEffect(() => {
    fetchItins();
  }, []);

  const fallbackImages = [
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Block-Column-Image_Boat-Trip_800x800.jpg",
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Block-Column-Image_Beach-Cabin_800x800.jpg",
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Block-Column-Image_Mining_800x800.jpg",
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Block-Column-Image_Winter-Celebration_800x800.jpg",
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Media-Block-Image_Java-Keyart_800x450.jpg",
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Media-Block-Image_PC-Bundle-Keyart_800x450.jpg",
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Updates-Carousel_Tricky-Trials_800x450.jpg",
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Updates-Carousel_Wild-Update_800x450.jpg",
  ];

  const getRandomImage = () => {
    const randomIndex = Math.floor(Math.random() * fallbackImages.length);
    return fallbackImages[randomIndex];
  };

  return (
    <div className="results">
      {Array.isArray(itineraries) && itineraries.length > 0 ? (
        // Render the results grid only when itineraries are present
        <div className="resultsGrid">
          {itineraries.map((itinerary) => (
            <div key={itinerary.itineraryID} className="resultCard" onClick={() => handleItinClick(itinerary.itineraryID)} style={{ cursor: 'pointer' }}>
              <img src={getRandomImage()} alt={itinerary.title} className="resultCardImage" />
              <div className="resultCardContent">
                <h4 className="cardLocation">{itinerary.city + ', ' + itinerary.country}</h4>
                <h3 className="cardTitle">{itinerary.title}</h3>
                <p className="cardDescription">{itinerary.description}</p>
                <div className="resultTags">
                  {itinerary.tags.map((tag, i) => (
                    <span key={i} className="resultCardTag">{tag}</span>
                  ))}
                </div>
              </div>
            </div>
          ))}
        </div>
      ) : (
        // Render the centered message container when no itineraries are found
        <div className="centeredMessageContainer">
          <div className="noItinerariesMessage">No itineraries found. Create a new one to get started!</div>
        </div>
      )}
    </div>
  );
}

export default Itineraries;
