import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/styles.css'; // Ensure this includes the correct styling
import boardAPI from '../api/boardAPI'; // Assuming boardAPI is used to fetch events
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'; // Import FontAwesome
import { faMapMarkerAlt } from '@fortawesome/free-solid-svg-icons'; // Import the location icon

const ItineraryGrid = ({ itineraries, getFallbackImage }) => {
  const navigate = useNavigate();
  const [eventImages, setEventImages] = useState({}); // Store first event image by itineraryId

  const handleItinClick = (itineraryId) => {
    navigate(`/itinerary/${itineraryId}`);
  };

  const fallbackImages = [
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Block-Column-Image_Boat-Trip_800x800.jpg",
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Block-Column-Image_Beach-Cabin_800x800.jpg",
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Block-Column-Image_Mining_800x800.jpg",
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Block-Column-Image_Winter-Celebration_800x800.jpg",
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Media-Block-Image_Java-Keyart_800x450.jpg",
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Block-Column-Image_PC-Bundle-Keyart_800x450.jpg",
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Updates-Carousel_Tricky-Trials_800x450.jpg",
    "https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Updates-Carousel_Wild-Update_800x450.jpg",
  ];

  const getRandomImage = () => {
    const randomIndex = Math.floor(Math.random() * fallbackImages.length);
    return fallbackImages[randomIndex];
  };

  const fetchEvents = async (itineraryId) => {
    try {
      const response = await boardAPI.get('/events', { params: { itineraryId } });
      return response.data.Events;
    } catch (error) {
      console.error("Error fetching events:", error);
      return [];
    }
  };

  useEffect(() => {
    const fetchAllEventImages = async () => {
      const imagesMap = {};
      for (const itinerary of itineraries) {
        const events = await fetchEvents(itinerary.itineraryId);
        if (events.length > 0 && events[0].eventImages?.length > 0) {
          imagesMap[itinerary.itineraryId] = `http://localhost:8080/images/${events[0].eventImages[0]}`;
        }
      }
      setEventImages(imagesMap);
    };

    fetchAllEventImages();
  }, [itineraries]);

  return (
    <div className="resultsGrid">
      {Array.isArray(itineraries) && itineraries.length > 0 ? (
        itineraries.map((itinerary) => {
          const imageUrl = eventImages[itinerary.itineraryId] || getFallbackImage?.() || getRandomImage();

          return (
            <div
              key={itinerary.itineraryId}
              className="resultCard"
              onClick={() => handleItinClick(itinerary.itineraryId)}
              style={{ cursor: 'pointer' }}
            >
              <img
                src={imageUrl}
                alt={itinerary.title}
                className="resultCardImage"
              />
              <div className="resultCardContent">
                <h4 className="cardLocation">
                  <FontAwesomeIcon icon={faMapMarkerAlt} style={{ marginRight: '8px', color: '#00509e' }} />
                  {`${itinerary.city}, ${itinerary.country}`}
                </h4>
                <h3 className="cardTitle">{itinerary.title}</h3>
                <p className="cardDescription">{itinerary.description}</p>
                <div className="resultTags">
                  {Array.isArray(itinerary.tags) &&
                    itinerary.tags.map((tag, i) => (
                      <span key={i} className="resultCardTag">
                        {tag.replace(/[^a-zA-Z\s]/g, '')}
                      </span>
                    ))}
                </div>
              </div>
            </div>
          );
        })
      ) : (
        <div className="centeredMessageContainer">
          <div className="noItinerariesMessage">
            No itineraries found. Create a new one to get started!
          </div>
        </div>
      )}
    </div>
  );
};

export default ItineraryGrid;
