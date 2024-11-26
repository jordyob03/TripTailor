import React, { useState, useRef, useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faMapMarkerAlt } from "@fortawesome/free-solid-svg-icons";
import iconMap from '../config/iconMap';
import tagsJson from "../config/tags.json";
import boardAPI from '../api/boardAPI';
import { faImage } from '@fortawesome/free-solid-svg-icons'; // Add this import at the top

const username = localStorage.getItem('username');
const fallbackImage =
  'https://t4.ftcdn.net/jpg/08/34/00/03/360_F_834000314_tLfhX7N7wnZpMkPIy02pqbRt8JFKiUuG.jpg';

function ItineraryDetails() {
  const location = useLocation();
  const itinId = parseInt(location.pathname.split("/").pop(), 10);
  const [events, setEvents] = useState([]);
  const [itinerary, setItinerary] = useState(null);
  const [errorMessage, setErrorMessage] = useState("");
  const [eventImages, setEventImages] = useState([]); // Store first event image by itineraryId
  const [isPopupOpen, setIsPopupOpen] = useState(false); // Controls the visibility of the popup
  const [currentImageIndex, setCurrentImageIndex] = useState(0); // Tracks the currently displayed image
  const [boards, setBoards] = useState([]);
  const [newBoardName, setNewBoardName] = useState("");
  const [newBoardImage, setNewBoardImage] = useState(null);
  const [showModal, setShowModal] = useState(false);
  const [selectedItinerary, setSelectedItinerary] = useState(null);
  const [images, setImages] = useState({});

  // Open the popup and set the clicked image index
  const openPopup = (index) => {
    setCurrentImageIndex(index);
    setIsPopupOpen(true);
  };

  // Close the popup
  const closePopup = () => {
    setIsPopupOpen(false);
  };

 // Show the next image and loop back to the first
    const showNextImage = () => {
        setCurrentImageIndex((prevIndex) => (prevIndex + 1) % popupImages.length);
    };
    
    // Show the previous image and loop back to the last
    const showPrevImage = () => {
        setCurrentImageIndex(
        (prevIndex) => (prevIndex - 1 + popupImages.length) % popupImages.length
        );
    };

  const categorizeTags = () => {
    if (!itinerary) return {}; // Return an empty object if itinerary is null
    const categorizedTags = {};
    Object.keys(tagsJson.categories).forEach((category) => {
      categorizedTags[category] = itinerary.tags.filter((tag) =>
        tagsJson.categories[category].includes(tag)
      );
    });
    return categorizedTags;
  };

  useEffect(() => {
    if (!itinId) {
      setErrorMessage('Invalid itinerary ID.');
      return;
    }
    fetchItineraryData();
  }, [itinId]);

  const fetchEvents = async (itineraryId) => {
    try {
      const response = await boardAPI.get('/events', { params: { itineraryId } });
      const fetchedEvents = response.data.Events || [];
      setEvents(fetchedEvents);

      // Extract event images
      const images = fetchedEvents.flatMap((event) =>
        event.eventImages?.length > 0
          ? event.eventImages.map(
              (img) => `http://localhost:8080/images/${img}`
            )
          : []
      );
      setEventImages(images); // Keep all images
    } catch (error) {
      console.error("Error fetching events:", error);
    }
  };

  const fetchItineraryData = async () => {
    try {
      const response = await boardAPI.get('/itineraries', { params: { postId: itinId } });
      const fetchedItinerary = response.data.Itinerary;
      setItinerary(fetchedItinerary);

      if (fetchedItinerary?.itineraryId) {
        await fetchEvents(fetchedItinerary.itineraryId); // Fetch events for this itinerary
      }
    } catch (error) {
      console.error('Error fetching itinerary data:', error);
      setErrorMessage('Failed to fetch itinerary data.');
    }
  };

  const categorizedTags = categorizeTags();

  const formatTime = (startTime, endTime) => {
    const options = { hour: "numeric", minute: "numeric", hour12: true };
    const start = new Date(startTime).toLocaleTimeString("en-US", options);
    const end = new Date(endTime).toLocaleTimeString("en-US", options);
    return `${start}-${end}`;
  };
  
  const fetchBoards = async () => {
    const userData = { username };
    const imagesMap = {};
    try {
      const response = await boardAPI.get("/boards", { params: userData });
      const boardsData = response.data.boards;
      setBoards(boardsData);
    } catch (error) {
      console.error("Error fetching boards:", error);
      setErrorMessage("An error occurred. Please try again.");
    }
  };

  useEffect(() => {
    fetchBoards();
  }, []);

  const handleSave = (itinerary) => {
    setSelectedItinerary(itinerary);
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setSelectedItinerary(null);
    setNewBoardName("");
    setNewBoardImage(null);
  };

  const handleSelectBoard = (boardId, postId) => {
    boardAPI
      .post("/addboardpost", { boardId, postId })
      .then((response) => {
        console.log("Itinerary saved successfully:", response.data);
        handleCloseModal();
      })
      .catch((error) => {
        console.error("Error saving itinerary to board:", error);
      });
  };

  const handleCreateNewBoard = async (itinerary) => {
    if (newBoardName.trim()) {
      const data = {
        Username: username,
        BoardName: newBoardName,
      };

      try {
        const response = await boardAPI.post("/addboard", data);
        const board = response.data.boardId;
        handleSelectBoard(board, itinerary.itineraryId);
        setNewBoardName("");
        setNewBoardImage(null);
      } catch (error) {
        console.error("Error creating new board:", error);
      }
    }
  };

  if (errorMessage) {
    return <div className="error">{errorMessage}</div>;
  }

  if (!itinerary) {
    return <div className="loading">Loading...</div>;
  }

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      setNewBoardImage(URL.createObjectURL(file));
    }
  };

  const displayImages = [
    ...(Array.isArray(eventImages) ? eventImages.slice(0, 4) : []), // Ensure eventImages has at most 4 items
    ...Array(Math.max(4 - (eventImages?.length || 0), 0)).fill(fallbackImage), // Fill the remaining slots with fallback images
  ].slice(0, 4); // Ensure the array length is exactly 4

  const popupImages = eventImages.filter((img) => img !== fallbackImage);

  return (
    <div className="container">
      <div className="itineraryHeader">
        <div className="itineraryText">
          <div className="itineraryDesc">
            <button className="saveButton1" onClick={() => handleSave(itinerary)}>
              Save to Board
            </button>
            <h1>{itinerary.title}</h1>
            <h4>
              <FontAwesomeIcon
                icon={faMapMarkerAlt}
                style={{ color: "#00509e", marginRight: "8px" }}
              />
              {itinerary.city}, {itinerary.country}
            </h4>
            <p>{itinerary.description}</p>
            <p className="postedBy">
              Posted by <span className="username">@{itinerary.username}</span>
            </p>
          </div>
          <div className="eventImages">
            {displayImages.slice(1).map((image, index) => (
              <img
                key={index}
                className="eventImage"
                src={image}
                alt={`Event ${index + 1}`}
                onClick={() => openPopup(index + 1)} // Open popup on image click
                style={{
                  cursor: "pointer", // Indicate the image is clickable
                  objectFit: "cover",
                  aspectRatio: "4 / 3",
                }}
              />
            ))}
          </div>
        </div>
        <img
          src={displayImages[0]}
          alt="Main Event"
          className="itineraryMainImage"
          onClick={() => openPopup(0)} // Main image is also clickable
          style={{ cursor: "pointer" }}
        />
      </div>
  
      <div className="itineraryTags">
        {itinerary.tags.map((tag, index) => (
          <div key={index} className="tagItem1">
            <div className="tagIcon1">
              {iconMap[tag] && <FontAwesomeIcon icon={iconMap[tag]} />}
            </div>
            <span>{tag}</span>
          </div>
        ))}
      </div>
  
      <div className="itineraryDetails">
        <div className="itineraryOverview">
          <h5 className="ItinTitle"> Itinerary Overview</h5>
          <p className="boldText">
            <strong>Location</strong>
          </p>
          <p className="smallText">
            {itinerary.city}, {itinerary.country}
          </p>
          <p className="boldText">
            <strong>Price</strong>
          </p>
          <p className="smallText">
            {itinerary.price === 0 ? "Free" : `$${itinerary.price}`}
          </p>
          <p className="boldText">
            <strong>Languages</strong>
          </p>
          <p className="smallText">{itinerary.languages.join(", ")}</p>
          {Object.keys(categorizedTags).map((category) => (
            <React.Fragment key={category}>
              <p className="boldText">
                <strong>
                  {category
                    .replace(/_/g, " ")
                    .replace(/\b\w/g, (c) => c.toUpperCase())}
                </strong>
              </p>
              <p className="smallText">
                {categorizedTags[category].length > 0
                  ? categorizedTags[category].join(", ")
                  : "N/A"}
              </p>
            </React.Fragment>
          ))}
        </div>
  
        <div className="itineraryTable">
          <h5 className="ItinTitle">Itinerary Events</h5>
          <table>
            <thead>
              <tr>
                <th>Time</th>
                <th>Name</th>
                <th>Location</th>
                <th>Notes</th>
                <th>Cost</th>
              </tr>
            </thead>
            <tbody>
              {events.map((event) => (
                <tr key={event.eventId}>
                  <td className="smallText">
                    {formatTime(event.startTime, event.endTime)}
                  </td>
                  <td>{event.name}</td>
                  <td className="smallText">{event.address}</td>
                  <td className="smallText">{event.description}</td>
                  <td>
                    {event.cost === 0 ? "Free" : `$${event.cost.toFixed(0)}`}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
  
      {/* Popup Modal for Image Slideshow */}
      {isPopupOpen && popupImages.length > 0 && (
        <div className="popupOverlay" onClick={closePopup}>
          <div
            className="popupContent"
            onClick={(e) => e.stopPropagation()} // Prevent closing on content click
          >
            <button className="prevButton" onClick={showPrevImage}>
              &lt;
            </button>
            <img
              className="popupImage"
              src={popupImages[currentImageIndex]}
              alt={`Event ${currentImageIndex}`}
              style={{ maxWidth: "100%", maxHeight: "90vh" }}
            />
            <button className="nextButton" onClick={showNextImage}>
              &gt;
            </button>
          </div>
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
                          handleSelectBoard(
                            board.id || board.boardId,
                            selectedItinerary.itineraryId
                          )
                        }
                      >
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
                            borderRadius: "20%",
                            objectFit: "cover",
                            cursor: "pointer",
                          }}
                        />
                      ) : (
                        <FontAwesomeIcon
                          icon={faImage}
                          className="imageInputIcon"
                          style={{ color: "grey", height: 35 }}
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
                  <button
                    className="createBoardButton"
                    onClick={() => handleCreateNewBoard(selectedItinerary)}
                  >
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

export default ItineraryDetails;
