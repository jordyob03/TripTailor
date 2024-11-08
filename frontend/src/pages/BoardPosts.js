import React, { useState, useEffect } from 'react';
import {useNavigate, useParams } from 'react-router-dom';
import '../styles/styles.css';
import boardAPI from '../api/boardAPI.js';
import { useLocation } from 'react-router-dom';

const userId = localStorage.getItem('userId');

function BoardPosts() {
  const navigate = useNavigate();
  const location = useLocation();
  const boardId = parseInt(location.pathname.split("/").pop());
  const [data, setData] = useState([]);
  const [errorMessage, setErrorMessage] = useState('');
  const [selectedBoard, setBoard] = useState(null);
  const [posts, setPosts] = useState([]);
  const [itineraries, setItineraries] = useState([]);
  const [events, setEvents] = useState([]); 

  const deleteBoardPost = async (boardId, postId) => {
    try {
      await boardAPI.delete(`/boards/${boardId}/posts/${postId}`);
      setPosts(posts.filter(post => post.postId !== postId));
      window.location.reload();
    } catch (error) {
      console.error("Error deleting post:", error);
      setErrorMessage("Failed to delete post");
    }
  }; 
  
  const handleBoardClick = (postId) => {
    // navigate(postId); // Navigate to the selectedBoard's specific URL
    window.location.href = `https://www.youtube.com/watch?v=dQw4w9WgXcQ`;
  };

  const fetchboards = async () => {
    const userData = { username: localStorage.getItem('username') };
    try {
      const response = await boardAPI.get('/boards', { params: userData });
      const selectedBoard = response.data.boards.find(board => board.boardId === boardId);
      setBoard(selectedBoard);
      return selectedBoard;
    } catch (error) {
      console.error("Error fetching boards:", error);
      setErrorMessage("Failed to fetch boards");
      return null;
    }
  };
  
  const fetchposts = async (boardId) => {
    try {
      const response = await boardAPI.get('/posts', { params: { boardId: boardId } });
      setPosts(response.data.Posts);
      return response.data.Posts;
    } catch (error) {
      console.error("Error fetching posts:", error.response ? error.response.data : error.message);
      setErrorMessage("Failed to fetch posts");
      return [];
    }
  };
  
  const fetchitineraries = async (postId) => {
    try {
      const response = await boardAPI.get('/itineraries', { params: { postId: postId } });
      return response.data.Itinerary;
    } catch (error) {
      console.error("Error fetching itineraries:", error);
      setErrorMessage("Failed to fetch itineraries");
      return null;
    }
  };
  
  const fetchevents = async (itineraryId) => {
    try {
      const response = await boardAPI.get('/events', { params: { itineraryId: itineraryId } }); 
      return response.data.Events;
    } catch (error) {
      console.error("Error fetching events:", error);
      setErrorMessage("Failed to fetch events");
      return [];
    }
  };
  
  const fetchAllData = async () => {
    try {
      const selectedBoard = await fetchboards();
      if (!selectedBoard) return;
      const postsData = await fetchposts(boardId);
      const structuredData = [];
      for (let post of postsData) {
        const itinerary = await fetchitineraries(post.postId); 
        const eventsData = await fetchevents(itinerary.itineraryId);
        structuredData.push({ post, itinerary, events: eventsData });
      }
      setData(structuredData); // Set the 3D array to the state
    } catch (error) {
      setErrorMessage("An error occurred while fetching data");
      console.error("Error fetching data:", error);
    }
  };
  
  useEffect(() => {
    if (isNaN(boardId)) {
      console.error("Invalid boardId:", boardId);
      return;
    }
    fetchAllData();
  }, [boardId]);
  
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
    <div className="boardPostsContainer">
      {selectedBoard && (
        <div className="boardDetails">
          <h2>{selectedBoard.name}</h2>
          <p><em>{selectedBoard.description}</em></p>
          <p>Created by: {selectedBoard.username}, on {new Date(selectedBoard.dateOfCreation).toLocaleDateString('en-GB', { day: 'numeric', month: 'long', year: 'numeric' })}</p>
        </div>
      )}
  
  <div className="postsGrid">
  {Array.isArray(data) && data.length > 0 ? (
    data.map((dataEntry, index) => {
      const { post, itinerary, events } = dataEntry;

      const eventImages = events.flatMap(event => event.eventImages || []);
      const randomImageNumber = eventImages.length > 0 
        ? eventImages[Math.floor(Math.random() * eventImages.length)]
        : null;

      const imageUrl = randomImageNumber 
        ? `http://localhost:8080/images/${randomImageNumber}` 
        : getRandomImage();

      return (
        <div key={index} className="postCard" onClick={() => handleBoardClick(post.postId)}>
          <button className="deleteButton" onClick={(e) => {e.stopPropagation(); deleteBoardPost(selectedBoard.boardId, post.postId);}}>X</button>
          <div className="imageContainer">
            <img src={imageUrl} alt={selectedBoard.name} className="post-image" />
          </div>
          <div className="postInfo">
            <h3>{itinerary.title}</h3>
            <span>{new Date(post.dateOfCreation).toLocaleDateString('en-GB', { day: 'numeric', month: 'long', year: 'numeric' })}</span>
            <p><strong>City:</strong> {itinerary.city}</p>
            <p><strong>Country:</strong> {itinerary.country}</p>
            <p><strong>Price:</strong> ${itinerary?.price !== undefined ? itinerary.price : 0}</p>
            <p><strong>Languages:</strong> {Array.isArray(itinerary.languages) ? itinerary.languages.join(", ") : "N/A"}</p>
            <p><strong>Tags:</strong> {Array.isArray(itinerary.tags) ? itinerary.tags.join(", ") : "N/A"}</p>
            <div className="postStats">
              <span>{post.likes || 0} Likes</span>
              <span>{post.comments ? post.comments.length : 0} Comments</span>
            </div>
          </div>
        </div>
      );
    })
  ) : (<></>)}
  </div>
    {errorMessage && <div className="message">{errorMessage}</div>}
  </div>
);
}

export default BoardPosts;
