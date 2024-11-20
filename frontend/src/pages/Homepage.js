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
  const [showModal, setShowModal] = useState(false);
  const [selectedItinerary, setSelectedItinerary] = useState(null);
  const [newBoardName, setNewBoardName] = useState('');
  const [newBoardImage, setNewBoardImage] = useState(null);
  const tagContainerRef = useRef(null);
  const [images, setImages] = useState({});
  const [errorMessage, setErrorMessage] = useState('');

  // Fetch itineraries from the backend
  useEffect(() => {
    // Fetch itineraries
    const fetchItineraries = async () => {
      const searchData = { country: 'Canada', city: 'Toronto' };
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

  const handleSave = (itinerary) => {
    console.log(itinerary); 
    console.log(showModal); 
    setSelectedItinerary(itinerary);
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setSelectedItinerary(null);
    setNewBoardName('');
    setNewBoardImage(null);
  };

  const handleSelectBoard = (boardId) => {
    // SAVE ITINERARY TO BOARD BACKEND INTEGRATION HERE
    console.log(`Itinerary saved to board ${boardId}`);
    handleCloseModal();
  };

  const handleCreateNewBoard = () => {
    if (newBoardName.trim()) {
      const newBoard = {
        id: boards.length + 1, // Mock ID
        name: newBoardName,
        coverImage: newBoardImage, // Add the new board image
      };
      // CREATE NEW BOARD BACKEND INTEGRATION HERE
      setBoards([...boards, newBoard]); // Update boards with the new board
      setNewBoardName(''); // Reset new board name
      setNewBoardImage(null); // Reset the image
    }
  };

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      setNewBoardImage(URL.createObjectURL(file));
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
          <ItineraryGrid itineraries={filteredItineraries} onSave={handleSave} />
        ) : (
          <p>No itineraries found.</p>
        )}
      </div>

      {/* Save Modal */}
      {showModal && (
        <div className="modalOverlay">
          <div className="modalContent">
            <div className="modalHeader">
              <h2>Save Itinerary to Board</h2>
              <button className="closeButton" onClick={handleCloseModal}>×</button>
            </div>
            <div className="modalBody">
            <div className="boardListContainer">
              <div className="boardList">
                {boards.map((board) => {
                  const eventImage = board.coverImage || `http://localhost:8080/images/${images[board.boardId]}`; // Use coverImage if available
                  return (
                    <div key={board.id || board.boardId} className="boardItem" onClick={() => handleSelectBoard(board.id || board.boardId)}>
                      {eventImage ? (
                        <img src={eventImage} alt={board.name} className="boardImage" />
                      ) : (
                        <div className="boardImagePlaceholder">No Image</div>
                      )}
                      <span className="boardName">{board.name}</span>
                    </div>
                  );
                })}
              </div>

              {/* New Board Input Section */}
              <div className="newBoardInput">
                <div className="newBoardInputRow">
                  <label className="imageInputLabel">
                    {newBoardImage ? (
                      <img
                        src={newBoardImage}
                        alt="Selected preview"
                        style={{
                          width: 35,
                          height: 35,
                          borderRadius: '20%',
                          objectFit: 'cover',
                          cursor: 'pointer',
                        }}
                      />
                    ) : (
                      <FontAwesomeIcon
                        icon={faImage}
                        className="imageInputIcon"
                        style={{ color: 'grey', height: 35 }}
                      />
                    )}
                    <input
                      type="file"
                      accept="image/*"
                      onChange={handleImageChange}
                      className="newBoardFileInput"
                    />
                  </label>
                  <input
                    type="text"
                    value={newBoardName}
                    onChange={(e) => setNewBoardName(e.target.value)}
                    placeholder="New board name"
                    className="newBoardInputField"
                  />
                </div>
                <button className="createBoardButton" onClick={handleCreateNewBoard}>
                  Create
                </button>
              </div>
              </div>
            </div>
            </div>
        </div>
      )}
    </div>
  );
}

export default HomePage;
