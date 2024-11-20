import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import boardAPI from '../api/boardAPI.js';
import '../styles/styles.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTimes } from '@fortawesome/free-solid-svg-icons';

const username = localStorage.getItem('username');

function Boards() {
  const [boards, setBoards] = useState([]);
  const [images, setImages] = useState({});
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();

  const deleteBoard = async (boardId) => {
    try {
      await boardAPI.delete(`/boards/${boardId}`);
      setBoards(boards.filter((board) => board.boardId !== boardId));
      window.location.reload();
    } catch (error) {
      console.error("Error deleting board:", error);
      setErrorMessage("Failed to delete board");
    }
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
  
  const handleBoardClick = (id) => {
    navigate(`/my-travels/boards/${id}`);
  };

  useEffect(() => {
    fetchBoards();
  }, []);

  const fallbackImage = "https://t4.ftcdn.net/jpg/08/34/00/03/360_F_834000314_tLfhX7N7wnZpMkPIy02pqbRt8JFKiUuG.jpg";

  return (
    <div className="results">
      {boards.length > 0 ? (
        <div className="resultsGrid">
          {boards.map((board) => {
            const eventImage = `http://localhost:8080/images/${images[board.boardId] || fallbackImage}`
            return (
              <div key={board.boardId} className="boardsCard" onClick={() => handleBoardClick(board.boardId)}>
                <img src={eventImage} alt={board.name} className="resultCardImage" />
                <div className="resultCardContent">
                  <button
                    className="deleteButton"
                    onClick={(e) => {
                      e.stopPropagation();
                      const confirmDelete = window.confirm("Are you sure you want to delete this board?");
                      if (confirmDelete) deleteBoard(board.boardId);
                    }}
                  >
                    <FontAwesomeIcon icon={faTimes} style={{ color: 'white' }} />
                  </button>
                  <h3 className="cardTitle">{board.name}</h3>
                  <div className="countAndDate">
                    <h4 className="cardPostCount">{board.posts.length} itineraries</h4>
                    <p className="cardDescription">
                      {new Date(board.dateOfCreation).toLocaleDateString('en-GB', { day: 'numeric', month: 'long', year: 'numeric' })}
                    </p>
                  </div>
                </div>
              </div>
            );
          })}
        </div>
      ) : (
        <div className="centeredMessageContainer">
          <div className="noBoardsMessage">No boards saved. Add some itineraries to your boards to start planning!</div>
        </div>
      )}
    </div>
  );
}

export default Boards;
