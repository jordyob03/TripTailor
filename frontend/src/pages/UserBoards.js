import React, { useState, useEffect} from 'react';
import { useNavigate, Link } from 'react-router-dom';
import '../styles/styles.css';
import boardAPI from '../api/boardAPI.js';

const username = localStorage.getItem('username');
console.log('User ID:', username);

function MyBoards() {
  const navigate = useNavigate();
  const [errorMessage, setErrorMessage] = useState('');
  const [boards, setBoards] = useState([]); // Define boards as a state variable

  const handlePostClick = (boardId) => {
    navigate(`/my-boards/${boardId}`);
  };

  const fetchboards = async () => {
    const userData = {
      username: username,
    };
  
    try {
      console.log('Fetching boards with:', userData);
      // Ensure you use axios correctly
      const response = await boardAPI.get('/boards', { params: userData });
      console.log('Boards fetched successfully:', response.data.boards);
  
      setBoards(response.data.boards);
    } catch (error) {
      setErrorMessage(error.message || "An error occurred. Please try again.");
      console.error("Error fetching boards:", error);
    }
  };

  useEffect(() => {
    fetchboards(); // Call fetchboards when the component mounts
  }, []);

  // Array of fallback images
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

  // Function to get a random fallback image
  const getRandomImage = () => {
    const randomIndex = Math.floor(Math.random() * fallbackImages.length);
    return fallbackImages[randomIndex];
  };

  return (
    <div className="my-boards">
      {/* Main Heading */}
      <div className="heading">
        <h2>My Boards</h2>
      </div>
      <div className="boards-container">
        {boards.map((board) => {
          const eventImage = board.image ? board.image : getRandomImage();
          return (
          <div key={board.boardId} className="board-card" onClick={() => handlePostClick(board.boardId)}>
            {/* Image section */}
            <img src={eventImage} alt={board} className="board-image" />
            
            {/* Content section */}
            <div className="board-content">
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
