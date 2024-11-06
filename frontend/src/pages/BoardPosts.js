import React, { useState, useEffect } from 'react';
import {useNavigate, useParams } from 'react-router-dom';
import '../styles/styles.css';
import boardAPI from '../api/boardAPI.js';

const userId = localStorage.getItem('userId');

// Mock function to simulate fetching selectedBoard data
const getBoard = async (boardId) => {
  // Simulate API response
  const boardData = {
    boardId: boardId, // Use the passed boardId correctly
    name: "Sample Board", // Placeholder for actual selectedBoard name
    creationDate: new Date(),
    description: `Description for selectedBoard with ID ${boardId}`, // Update description to include boardId
    username: "board_owner",
    posts: ["1", "2", "3", "4", "5", "6"], // Placeholder for actual post IDs
    tags: ["Travel", "Adventure"]
  };
  return boardData;
};

// Mock function to simulate fetching post data
const getPost = async (postId) => {
  // Simulate API response
  const postData = {
    postId: postId,
    itineraryId: 101 + parseInt(postId), // Just an example
    creationDate: new Date(),
    username: `user_${postId}`,
    boards: ["Sample Board"],
    likes: Math.floor(Math.random() * 100), // Random number for likes
    comments: [`Comment 1 for post ${postId}`, `Comment 2 for post ${postId}`]
  };
  return postData;
};

const GetItinerary = (itineraryId) => {
    // Simulate fetching itinerary data from an API or database
    const mockItineraryData = {
      itineraryId: itineraryId,
      name: "Summer in Paris",
      city: "Paris",
      country: "France",
      title: "Amazing Summer Vacation",
      description: "A detailed itinerary for an amazing summer vacation in Paris.",
      price: 2000.00,
      languages: ["English", "French"],
      tags: ["Vacation", "Summer", "Paris"],
      events: ["Eiffel Tower Visit", "Seine River Cruise"],
      postId: 1,
      username: "john_doe",
    };

    return mockItineraryData;
}

const GetEvent = (EventId) => {
    // Simulate fetching itinerary data from an API or database
    const mockEventData = {
        EventId: EventId,
        Name: "Eiffel Tower Visit",
        Cost: 50.00,
        Address: "Champ de Mars, 5 Avenue Anatole, 75007 Paris, France",
        Description: "Visit the iconic Eiffel Tower and enjoy panoramic views of Paris.",
        StartTime: "2024-07-01T10:00:00",
        EndTime: "2024-07-01T12:00:00",
        ItineraryId: 101,
        EventImages: ["", ""],
    };

    return mockEventData;
}

function BoardPosts() {
  const navigate = useNavigate();
  const { boardIdStr } = useParams();
  const boardId = boardIdStr ? parseInt(boardIdStr, 10) : 1;
  const [data, setData] = useState([]);
  const [errorMessage, setErrorMessage] = useState('');
  const [selectedBoard, setBoard] = useState(null); // Store the selected selectedBoard
  const [posts, setPosts] = useState([]);
  const [itineraries, setItineraries] = useState([]);
  const [events, setEvents] = useState([]);
  
  const handleBoardClick = (postId) => {
    // navigate(postId); // Navigate to the selectedBoard's specific URL
    window.location.href = `https://www.youtube.com/watch?v=dQw4w9WgXcQ`;
  };

  const fetchboards = async () => {
    const userData = { username: localStorage.getItem('username') };
    console.log("Fetching boards with username:", userData.username);
  
    try {
      const response = await boardAPI.get('/boards', { params: userData });
      console.log("Fetched boards:", response.data);
  
      const selectedBoard = response.data.boards.find(board => board.boardId === parseInt(boardId));
      console.log("Selected board:", selectedBoard);
  
      setBoard(selectedBoard);
      return selectedBoard;
    } catch (error) {
      console.error("Error fetching boards:", error);
      setErrorMessage("Failed to fetch boards");
      return null;
    }
  };
  
  const fetchposts = async (boardId) => {
    console.log("Fetching posts for boardId:", boardId);
  
    try {
      const response = await boardAPI.get('/posts', { params: { boardId: boardId } });
      console.log("Fetched posts:", response.data.Posts);
      setPosts(response.data.Posts);
      return response.data.Posts;
    } catch (error) {
      console.error("Error fetching posts:", error.response ? error.response.data : error.message);
      setErrorMessage("Failed to fetch posts");
      return [];
    }
  };
  
  const fetchitineraries = async (postId) => {
    console.log("Fetching itineraries for postId:", postId);
  
    try {
      const response = await boardAPI.get('/itineraries', { params: { postId: postId } });
      console.log("Fetched itineraries:", response.data.itineraries);
  
      return response.data.itineraries;
    } catch (error) {
      console.error("Error fetching itineraries:", error);
      setErrorMessage("Failed to fetch itineraries");
      return [];
    }
  };
  
  const fetchevents = async (itineraryId) => {
    console.log("Fetching events for itineraryId:", itineraryId);
  
    try {
      const response = await boardAPI.get('/events', { params: { itineraryId: itineraryId } });
      console.log("Fetched events:", response.data.events);
  
      return response.data.events;
    } catch (error) {
      console.error("Error fetching events:", error);
      setErrorMessage("Failed to fetch events");
      return [];
    }
  };
  
  const fetchAllData = async () => {
    console.log("Fetching all data for boardId:", boardId);
    try {
      const selectedBoard = await fetchboards();
      if (!selectedBoard) return;
  
      const postsData = await fetchposts(selectedBoard.boardId);
      const structuredData = [];
  
      // For each post, fetch itineraries, then events
      for (let post of postsData) {
        console.log("Fetching itineraries for postId:", post.postId);
  
        const itinerariesData = await fetchitineraries(post.postId);
        console.log("Fetched itineraries for postId:", post.postId, itinerariesData);
  
        const itineraryEvents = [];
  
        for (let itinerary of itinerariesData) {
          console.log("Fetching events for itineraryId:", itinerary.itineraryId);
          const eventsData = await fetchevents(itinerary.itineraryId);
          console.log("Fetched events for itineraryId:", itinerary.itineraryId, eventsData);
  
          itineraryEvents.push(eventsData);
        }
  
        // Push the post's itineraries and events into the structured data
        structuredData.push(itineraryEvents);
      }
  
      console.log("Structured data:", structuredData);
      setData(structuredData); // Set the 3D array to the state
  
    } catch (error) {
      setErrorMessage("An error occurred while fetching data");
      console.error("Error fetching data:", error);
    }
  };
  
  useEffect(() => {
    console.log("useEffect: Fetching data for boardId:", boardId);
    if (isNaN(boardId)) {
      console.error("Invalid boardId:", boardId);
      return;
    }
    fetchAllData(); // Fetch data when the component mounts
  }, [boardId]); // Re-fetch data if boardId changes
  

  const fallbackImage = 'https://www.minecraft.net/content/dam/minecraftnet/franchise/logos/Homepage_Download-Launcher_Creeper-Logo_500x500.png';

  return (
    <div className="boardPostsContainer">
      {selectedBoard && (
        <div className="selectedBoard">
          <h2>{selectedBoard.name}</h2>
          <p><em>{selectedBoard.description}</em></p>
          <p>Created by: {selectedBoard.username}, on {new Date(selectedBoard.creationDate).toLocaleDateString('en-GB', { day: 'numeric', month: 'long', year: 'numeric' })}</p>
        </div>
      )}

      <div className="postsGrid">
        {Array.isArray(posts) && posts.length > 0 ? (
          posts.map((post) => {
            const itinerary = GetItinerary(post.itineraryId);
            const eventImage = GetEvent(itinerary.events[0]).EventImages[0]
              ? GetEvent(itinerary.events[0]).EventImages[0]
              : [fallbackImage];

            return (
            <div key={post.postId} className="postCard" onClick={() => handleBoardClick(post.postId)}>
                 <img src={eventImage} alt={selectedBoard.name} className="selectedBoard-image" />
              <div key={post.postId} className="postCard">
                <div className="postInfo">
                <h3>{itinerary.title}</h3>
                  <span>{new Date(post.creationDate).toLocaleDateString('en-GB', { day: 'numeric', month: 'long', year: 'numeric' })}</span>
                  <p><strong>City:</strong> {itinerary.city}</p>
                  <p><strong>Country:</strong> {itinerary.country}</p>
                  <p><strong>Price:</strong> ${itinerary.price}</p>
                  <p><strong>Languages:</strong> {itinerary.languages.join(", ")}</p>
                  <p><strong>Tags:</strong> {itinerary.tags.join(", ")}</p>
                  <div className="postStats">
                    <span>{post.likes} Likes</span>
                    {/* <span>{post.comments ? post.comments.length : 0} Comments</span> */}
                  </div>
                </div>
                </div>
            </div>
            );
          })
        ) : (
          <div className="noPostsMessage">No posts available in this selectedBoard.</div>
        )}
      </div>
    </div>
  );
}

export default BoardPosts;
