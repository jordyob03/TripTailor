import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import boardAPI from '../api/boardAPI.js';
import '../styles/styles.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTimes } from '@fortawesome/free-solid-svg-icons';

const username = localStorage.getItem('username');

function Boards() {
  const [boards, setBoards] = useState([]);
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

  const fetchBoards = async () => {
    const userData = { username: username };
    try {
      const response = await boardAPI.get('/boards', { params: userData });
      setBoards(response.data.boards);
    } catch (error) {
      setErrorMessage(error.message || "An error occurred. Please try again.");
    }
  };

  const handleBoardClick = (id) => {
    navigate(`/my-travels/boards/${id}`);
  };

  useEffect(() => {
    fetchBoards();
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
      {boards.length > 0 ? (
        <div className="resultsGrid">
          {boards.map((board) => {
            const eventImage = board.image || getRandomImage();
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
                  <div className="resultTags">
                    {board.tags.map((tag, i) => (
                      <span key={i} className="resultCardTag">{tag}</span>
                    ))}
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
