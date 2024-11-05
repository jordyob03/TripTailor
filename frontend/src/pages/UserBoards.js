import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import '../styles/styles.css';

function MyBoards() {
  const navigate = useNavigate();

  const handleBoardClick = (boardId) => {
    navigate(`/my-boards/${boardId}`); // Navigate to the board's specific URL
  };

  // Mock data for boards, with images added
  const mockData = [
    {
      boardId: 1,
      name: 'Summer Adventure',
      dateOfCreation: '2024-06-15',
      description: 'A trip to explore beaches and mountains.',
      username: 'user123',
      posts: ['Post 1', 'Post 2'],
      tags: ['beach', 'mountain', 'summer'],
      image: 'https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Block-Column-Image_Boat-Trip_800x800.jpg',
    },
    {
      boardId: 2,
      name: 'Winter Getaway',
      dateOfCreation: '2023-12-05',
      description: 'A cozy retreat in the mountains.',
      username: 'user456',
      posts: ['Post 3', 'Post 4'],
      tags: ['winter', 'retreat', 'cozy'],
      image: 'https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Block-Column-Image_Mining_800x800.jpg',
    },
    {
      boardId: 3,
      name: 'City Exploration',
      dateOfCreation: '2024-08-20',
      description: 'A cultural and historical tour of the city.',
      username: 'user789',
      posts: ['Post 5', 'Post 6'],
      tags: ['city', 'culture', 'history'],
      image: 'https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Block-Column-Image_Winter-Celebration_800x800.jpg',
    },
    {
        boardId: 4,
        name: 'Road Trip Adventures',
        dateOfCreation: '2024-07-10',
        description: 'An epic road trip across the country, visiting national parks and landmarks.',
        username: 'user234',
        posts: ['Post 7', 'Post 8'],
        tags: ['road trip', 'national parks', 'adventure'],
        image: 'https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Media-Block-Image_PC-Bundle-Keyart_800x450.jpg',
      },
      {
        boardId: 5,
        name: 'Tropical Escape',
        dateOfCreation: '2024-03-01',
        description: 'A relaxing beach holiday to escape the winter blues and soak up the sun.',
        username: 'user567',
        posts: ['Post 9', 'Post 10'],
        tags: ['tropical', 'beach', 'vacation'],
        image: 'https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Block-Column-Image_Beach-Cabin_800x800.jpg',
      },
      {
        boardId: 6,
        name: 'Mountain Trekking',
        dateOfCreation: '2024-05-12',
        description: 'A challenging trek through the mountains, exploring breathtaking landscapes.',
        username: 'user890',
        posts: ['Post 11', 'Post 12'],
        tags: ['mountain', 'trekking', 'outdoor'],
        image: 'https://www.minecraft.net/content/dam/games/minecraft/key-art/MC-Vanilla_Media-Block-Image_Java-Keyart_800x450.jpg',
      }
      
  ];

  const boards = mockData;

  return (
    <div className="my-boards">
      {/* Main Heading */}
      <div className="heading">
        <h2>My Boards</h2>
      </div>
      <div className="boards-container">
        {boards.map((board) => (
          <div key={board.boardId} className="board-card" onClick={() => handleBoardClick(board.boardId)}>
            {/* Image section */}
            <img src={board.image} alt={board.name} className="board-image" />
            
            {/* Content section */}
            <div className="board-content">
              <h3>{board.name}</h3>
              <p><strong>{new Date(board.dateOfCreation).toLocaleDateString('en-GB', { day: 'numeric', month: 'long', year: 'numeric' })}</strong></p>
              <p><em>{board.description}</em></p>
              <p><strong>Tags:</strong> {board.tags.join(', ')}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default MyBoards;
