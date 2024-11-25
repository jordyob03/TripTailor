import { faImage } from '@fortawesome/free-solid-svg-icons'; // Add this import at the top
import React, { useState, useRef, useEffect } from 'react';
import '../styles/styles.css';
import Tags from '../config/tags.json';
import iconMap from '../config/iconMap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import feedAPI from '../api/feedAPI.js';
import ItineraryGrid from '../components/ItineraryGrid.js';
import itineraryAPI from '../api/itineraryAPI.js';
import boardAPI from '../api/boardAPI.js';
import profileAPI from '../api/profileAPI';

const fetchfeed = async (tags) => {
  try {
    // Convert tags array to JSON string and encode it
    const tagsParam = JSON.stringify(tags);
    console.log(tags, username);
    const response = await feedAPI.get('/feed', { params: { tags: tagsParam } });
    if (response && response.data) {
      console.log(response.data);
      return response.data.itineraries
    } else {
      console.log("No data received in response.");
    }
  } catch (error) {
    console.log(error.message || "An error occurred. Please try again.");
  }
};

const username = localStorage.getItem('username');

function HomePage() {
  const [selectedTags, setSelectedTags] = useState([]);
  const [itineraries, setItineraries] = useState([]);
  const [filteredItineraries, setFilteredItineraries] = useState([]);
  const [boards, setBoards] = useState([]);
  const tagContainerRef = useRef(null);
  const [images, setImages] = useState({});
  const [errorMessage, setErrorMessage] = useState('');

  // Fetch itineraries from the backend
  useEffect(() => {
    const fetchItineraries = async () => {
      const response = [];
      const users = ["testuser", "herobrine", "jordy_ob", "mitchro"];
  
      const user = await profileAPI.get("/user", { params: { username } });
      const userData = user.data;
  
      // Users tags
      const userTags = userData.tags;
  
      try {
        // Loop through each user and fetch their itineraries
        for (const user of users) {
          const data = { Username: user };
          const tempResponse = await itineraryAPI.post('/get-user-itins', data);
          if (Array.isArray(tempResponse.data.itineraries)) {
            response.push(...tempResponse.data.itineraries);
          } else {
            console.error(`Itineraries for ${user} is not an array`, tempResponse.data.Itineraries);
          }
        }
  
        // Filter itineraries based on user tags
        const filteredItineraries = response.filter((itinerary) =>
          itinerary.tags.some((tag) => userTags.includes(tag))
        );
  
        // Only set filtered itineraries here, no need to set `itineraries` unless needed
        setFilteredItineraries(filteredItineraries);
  
      } catch (error) {
        console.error("Error fetching itineraries:", error);
      }
    };
  
    fetchItineraries();
  }, []);
  
  // Handle tag filtering
  useEffect(() => {
    if (selectedTags.length > 0) {
      const filtered = itineraries.filter((itinerary) =>
        itinerary.tags.some((tag) => selectedTags.includes(tag))
      );
      setFilteredItineraries(filtered);
    } else {
      setFilteredItineraries(itineraries);
    }
  }, [selectedTags, itineraries]);
  

  // Handle tag filtering
  useEffect(() => {
    if (selectedTags.length > 0) {
      const filtered = itineraries.filter((itinerary) =>
        itinerary.tags.some((tag) => selectedTags.includes(tag))
      );
      setFilteredItineraries(filtered);
    } else {
      setFilteredItineraries(itineraries);
    }
  }, [selectedTags, itineraries]);

  const fetchEvents = async (itineraryId) => {
    try {
      const response = await boardAPI.get('/events', { params: { itineraryId } });
      return response.data.Events;
    } catch (error) {
      console.error("Error fetching events:", error);
      return [];
    }
  };

  const fetchBoards = async () => {
    const userData = { username };
    const imagesMap = {};
    try {
      const response = await boardAPI.get('/boards', { params: userData });
      const boardsData = response.data.boards;
      setBoards(boardsData);
  
      const imagePromises = boardsData.map(async (board) => {
        try {
          const postsResponse = await boardAPI.get('/posts', { params: { boardId: board.boardId } });
          const firstPostId = postsResponse.data.Posts?.[0]?.postId;
  
          if (firstPostId) {
            const itinerary = await boardAPI.get('/itineraries', { params: { postId: firstPostId } });
            const events = await fetchEvents(itinerary.data.Itinerary?.itineraryId);
  
            if (events.length > 0 && events[0].eventImages?.length > 0) {
              // Use the first image ID from the event
              const imageId = events[0].eventImages[0];
              imagesMap[board.boardId] = imageId; // Store imageId for the board
            }
          }
        } catch (error) {
          console.error(`Error fetching data for board ${board.boardId}:`, error);
        }
        return { boardId: board.boardId, image: null };
      });
  
      await Promise.all(imagePromises);
      setImages(imagesMap);
    } catch (error) {
      console.error('Error fetching boards:', error);
      setErrorMessage(error.message || "An error occurred. Please try again.");
    }
  };

  const scrollTagsLeft = () => {
    if (tagContainerRef.current) {
      tagContainerRef.current.scrollBy({ left: -150, behavior: 'smooth' });
    }
  };

  const scrollTagsRight = () => {
    if (tagContainerRef.current) {
      tagContainerRef.current.scrollBy({ left: 150, behavior: 'smooth' });
    }
  };

  const handleTagClick = async (tag) => {
    // Toggle the selection of the tag
    let updatedTags;
    // Logic for handeling if user already clicked tag or not
    if (selectedTags.includes(tag)) {
      // If tag is already selected, keep the existing selected tags
      updatedTags = [...selectedTags];
    } else {
      // If tag is not selected, add it to the selected tags
      updatedTags = [...selectedTags, tag];
    }
  
    setSelectedTags(updatedTags);
  
    if (updatedTags.length >= 1) {
      // Fetch feed with the updated tags and wait for the result
      const fetchedData = await fetchfeed(updatedTags);
  
      if (Array.isArray(fetchedData)) {
        setItineraries(fetchedData); 
      } else {
        setItineraries([]); 
      }
    } else {
      
      setItineraries([]); 
    }
  };
  
  

  return (
    <div>
      {/* Tag Filter Scrollable Container */}
      <div className="tagFiltersContainer">
        <button onClick={scrollTagsLeft} className="arrowButton">{'<'}</button>
        <div ref={tagContainerRef} className="tagContainer">
          {Object.values(Tags.categories).flat().map((tag) => (
            <div
              key={tag}
              onClick={() => handleTagClick(tag)}
              className={`tagItem ${selectedTags.includes(tag) ? 'selected' : ''}`}
            >
              <div className="tagIcon">
                {iconMap[tag] && <FontAwesomeIcon icon={iconMap[tag]} />}
              </div>
              <div className="tagLabel">{tag}</div>
            </div>
          ))}
        </div>
        <button onClick={scrollTagsRight} className="arrowButton">{'>'}</button>
      </div>

      {/* Feed Section: Filtered Itineraries */}
      <div className="pageContainer" style={{ marginTop: '100px' }}>
        {filteredItineraries.length > 0 ? (
          <ItineraryGrid itineraries={filteredItineraries} showSaveButton={true} />
        ) : (
          <p>No itineraries found.</p>
        )}
      </div>
    </div>
  );
}

export default HomePage;