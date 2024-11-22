import React, { useState, useRef, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/styles.css';
import boardAPI from '../api/boardAPI';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'; // Import FontAwesome
import { faMapMarkerAlt } from '@fortawesome/free-solid-svg-icons'; // Import the location icon
import { faImage } from '@fortawesome/free-solid-svg-icons'; // Add this import at the top

const username = localStorage.getItem('username');

const ItineraryGrid = ({ itineraries, showSaveButton, editMode, onDeletePost}) => {
  const [selectedTags, setSelectedTags] = useState([]);
  const [filteredItineraries, setFilteredItineraries] = useState([]);
  const [boards, setBoards] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [selectedItinerary, setSelectedItinerary] = useState(null);
  const [newBoardName, setNewBoardName] = useState('');
  const [newBoardImage, setNewBoardImage] = useState(null);
  const tagContainerRef = useRef(null);
  const [images, setImages] = useState({});
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();
  const [eventImages, setEventImages] = useState({}); // Store first event image by itineraryId

  const handleItinClick = (itineraryId) => {
    navigate(`/itinerary/${itineraryId}`);
  };

  const fallbackImage = "https://t4.ftcdn.net/jpg/08/34/00/03/360_F_834000314_tLfhX7N7wnZpMkPIy02pqbRt8JFKiUuG.jpg";

  const fetchEvents = async (itineraryId) => {
    try {
      const response = await boardAPI.get('/events', { params: { itineraryId } });
      return response.data.Events;
    } catch (error) {
      console.error("Error fetching events:", error);
      return [];
    }
  };

  const fetchBoardsAndImages = async () => {
    await fetchBoards();
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
    fetchBoardsAndImages();
  }, [itineraries]);

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

  const handleSelectBoard = (boardId, postId) => {
    boardAPI.post('/addpost', {boardId, postId})
      .then((response) => {
        console.log('Itinerary saved successfully:', response.data);
        handleCloseModal();
      })
      .catch((error) => {
        console.error('Error saving itinerary to board:', error);
      });
    
    console.log(`Itinerary ${postId} saved to board ${boardId}`);
    handleCloseModal();
  };

  const handleCreateNewBoard = async (itinerary) => {
    if (newBoardName.trim()) {
      const data = {
        Username: username,
        BoardName: newBoardName
      };

      console.log("data: ", data)
      const response = await boardAPI.post('/boards', data);
      console.log(response)
      const board = response.data.boardId
      handleSelectBoard(board, itinerary.itineraryId)
      console.log("Itinerary, ", itinerary.itineraryId, " saved to board, ", board)
      window.location.reload()
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

  return (
    <div className="resultsGrid">
      {Array.isArray(itineraries) && itineraries.length > 0 ? (
        itineraries.map((itinerary) => {
          const imageUrl = eventImages[itinerary.itineraryId] || fallbackImage;
          return (
            <div
              key={itinerary.itineraryId}
              className="resultCard"
              onClick={() => handleItinClick(itinerary.itineraryId)}
              style={{ position: 'relative' }}
            >
              {showSaveButton && (
                <button
                  onClick={(e) => {
                    e.stopPropagation();
                    handleSave(itinerary);
                  }}
                  className="saveButton"
                >
                  Save
                </button>
              )}
              <img
                src={imageUrl}
                alt={itinerary.title}
                className="resultCardImage"
              />
              <div className="resultCardContent">
                <h4 className="cardLocation">
                  <FontAwesomeIcon
                    icon={faMapMarkerAlt}
                    style={{ marginRight: '8px', color: '#00509e' }}
                  />
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
                {editMode && (
                  <button
                    className="deleteButton"
                    onClick={(e) => {
                      e.stopPropagation();
                      onDeletePost(itinerary.postId);
                    }}
                  >
                    X
                  </button>
                )}
              </div>
            </div>
          );
        })
      ) : (
        <div className="centeredMessageContainer">
          <div className="noItinerariesMessage">No itineraries found.</div>
        </div>
      )}
      {/* Save Modal */}
      {showModal && (
        <div className="modalOverlay">
          <div className="modalContent">
            <div className="modalHeader">
              <h2>Save Itinerary to Board</h2>
              <button className="closeButton" onClick={handleCloseModal}>
                ×
              </button>
            </div>
            <div className="modalBody">
              <div className="boardListContainer">
                <div className="boardList">
                  {boards.map((board) => {
                    const eventImage =
                      board.coverImage ||
                      `http://localhost:8080/images/${images[board.boardId]}`;
                    return (
                      <div
                        key={board.id || board.boardId}
                        className="boardItem"
                        onClick={() =>
                          handleSelectBoard(board.id || board.boardId, selectedItinerary.itineraryId)
                        }
                      >
                        {eventImage ? (
                          <img
                            src={eventImage}
                            alt={board.name}
                            className="boardImage"
                          />
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
                  <button className="createBoardButton" onClick={() => handleCreateNewBoard(selectedItinerary)}>
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

export default ItineraryGrid;
