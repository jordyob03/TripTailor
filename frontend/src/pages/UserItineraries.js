import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import itineraryAPI from '../api/itineraryAPI.js';
import boardAPI from '../api/boardAPI.js';
import '../styles/styles.css';

const username = localStorage.getItem('username');

function Itineraries() {
  const [itineraries, setItineraries] = useState([]);
  const [itineraryEvents, setItineraryEvents] = useState([]); // 2D array for itineraries and events
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();

  const fetchItins = async () => {
      try {
        const data = {
          Username: localStorage.getItem('username'),
        };
        const response = await itineraryAPI.post('/get-user-itins', data);
        console.log('Received', response.data);
        const fetchedItineraries = response.data.itineraries;
        setItineraries(fetchedItineraries);

        const eventsData = await Promise.all(
          fetchedItineraries.map(async (itinerary) => {
            const events = await fetchevents(itinerary.itineraryId);
            return { itinerary, events };
          })
        );
        setItineraryEvents(eventsData);
      } catch (error) {
        console.error('Error fetching itineraries:', error);
      }
    };

  const fetchevents = async (itineraryId) => {
    try {
      const response = await boardAPI.get('/events', { params: { itineraryId: itineraryId } }); 
      return response.data.Events;
    } catch (error) {
      console.error("Error fetching events:", error);
      setErrorMessage("Failed to fetch events");
      return [];
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
      {Array.isArray(itineraryEvents) && itineraryEvents.length > 0 ? (
        <div className="resultsGrid">
          {itineraryEvents.map(({ itinerary, events }) => {
            const eventImages = events.flatMap(event => event.eventImages || []);
            const randomImageNumber = eventImages.length > 0 
              ? eventImages[Math.floor(Math.random() * eventImages.length)] 
              : null;
          
            const imageUrl = randomImageNumber 
              ? `http://localhost:8080/images/${randomImageNumber}` 
              : getRandomImage();
            
            return (
            <div key={itinerary.itineraryId} className="resultCard" onClick={() => handleItinClick(itinerary.itineraryId)} style={{ cursor: 'pointer' }}>
              <img src={imageUrl} alt={itinerary.title} className="resultCardImage" />
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
          )})}
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
