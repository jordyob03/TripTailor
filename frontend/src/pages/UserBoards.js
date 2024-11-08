import React, { useState, useEffect} from 'react';
import { useNavigate, Link } from 'react-router-dom';
import '../styles/styles.css';
import boardAPI from '../api/boardAPI.js';

const username = localStorage.getItem('username');

function MyBoards() {
  const navigate = useNavigate();
  const [errorMessage, setErrorMessage] = useState('');
  const [boards, setBoards] = useState([]);

  const deleteBoard = async (boardId) => {
    try {
      await boardAPI.delete(`/boards/${boardId}`);
      setBoards(boards.filter(board => board.boardId !== boardId));
      window.location.reload();
    } catch (error) {
      console.error("Error deleting board:", error);
      setErrorMessage("Failed to delete board");
    }
  }; 

  const handlePostClick = (boardId) => {
    navigate(`/my-boards/${boardId}`);
  };

  const fetchboards = async () => {
    const userData = {
      username: username,
    };
  
    try {
      const response = await boardAPI.get('/boards', { params: userData }); 
      setBoards(response.data.boards);
    } catch (error) {
      setErrorMessage(error.message || "An error occurred. Please try again.");
    }
  };

  useEffect(() => {
    fetchboards();
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
    <div className="my-boards">
      <div className="heading">
        <h2>My Boards</h2>
      </div>
      <div className="boards-container">
        {boards.map((board) => {
          const eventImage = board.image ? board.image : getRandomImage();
          return (
          <div key={board.boardId} className="board-card" onClick={() => handlePostClick(board.boardId)}>
            <div className="boardImageContainer">
              <img src={eventImage} alt={board} className="board-image" />
            </div>
            <div className="board-content">
              <button className="deleteButton" onClick={(e) => {e.stopPropagation(); deleteBoard(board.boardId);}}>X</button>
              <h3>{board.name}</h3>
              <p><strong>{new Date(board.dateOfCreation).toLocaleDateString('en-GB', { day: 'numeric', month: 'long', year: 'numeric' })}</strong></p>
              <p><em>{board.description}</em></p>
              <p><strong>Tags:</strong> {board.tags.join(', ')}</p>
            </div>
          </div>
        )})}
      </div>
    </div>
  );
}

export default MyBoards;
