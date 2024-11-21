import React, { useState, useRef, useEffect } from 'react';
import '../styles/styles.css';
import Tags from '../config/tags.json';
import iconMap from '../config/iconMap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import ItineraryGrid from '../components/ItineraryGrid.js';
import searchAPI from '../api/searchAPI';
import boardAPI from '../api/boardAPI.js';
import { faImage } from '@fortawesome/free-solid-svg-icons'; // Add this import at the top

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
    // Fetch itineraries
    const fetchItineraries = async () => {
      const searchData = { country: 'France', city: 'Paris' };
      try {
        console.log("Search API sent:", searchData);
        const response = await searchAPI.get('/search', {
          params: searchData,
        });
        console.log('API response:', response.data.Itineraries);
        setItineraries(response.data || []);
        setFilteredItineraries(response.data || []);
      } catch (error) {
        console.error("Error fetching itineraries:", error);
      }
    };
  
    // Fetch boards
    const fetchBoardsAndImages = async () => {
      await fetchBoards();
    };
  
    fetchItineraries();
    fetchBoardsAndImages();
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

  const handleTagClick = (tag) => {
    setSelectedTags((prevSelectedTags) =>
      prevSelectedTags.includes(tag)
        ? prevSelectedTags.filter((t) => t !== tag)
        : [...prevSelectedTags, tag]
    );
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
