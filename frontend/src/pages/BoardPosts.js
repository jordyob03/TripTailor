import React, { useState, useEffect } from 'react';
import {useNavigate, useParams } from 'react-router-dom';
import '../styles/styles.css';

const userId = localStorage.getItem('userId');

// Mock function to simulate fetching board data
const getBoard = async (boardId) => {
  // Simulate API response
  const boardData = {
    boardId: boardId, // Use the passed boardId correctly
    name: "Sample Board", // Placeholder for actual board name
    creationDate: new Date(),
    description: `Description for board with ID ${boardId}`, // Update description to include boardId
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
  const { board_id } = useParams(); // Retrieve the boardId from URL params
  const [posts, setPosts] = useState([]); // State to store posts
  const [itinerary, setItinerary] = useState(null); // State to store itinerary data
  const [boardDetails, setBoardDetails] = useState(null); // State to store board details

  const navigate = useNavigate();
  
  const handleBoardClick = (postId) => {
    // navigate(postId); // Navigate to the board's specific URL
    window.location.href = `https://www.youtube.com/watch?v=dQw4w9WgXcQ`;
  };

  useEffect(() => {
    const fetchBoardData = async () => {
      const board = await getBoard(board_id); // Correctly pass board_id to getBoard
      setBoardDetails(board);

      // Fetch posts data for each post ID in the board
      const postsData = await Promise.all(board.posts.map(postId => getPost(postId)));
      setPosts(postsData);
    };

    fetchBoardData();
  }, [board_id]); // Run effect when board_id changes

  const fallbackImage = 'https://www.minecraft.net/content/dam/minecraftnet/franchise/logos/Homepage_Download-Launcher_Creeper-Logo_500x500.png';

  return (
    <div className="boardPostsContainer">
      {boardDetails && (
        <div className="boardDetails">
          <h2>{boardDetails.name}</h2>
          <p><em>{boardDetails.description}</em></p>
          <p>Created by: {boardDetails.username}, on {new Date(boardDetails.creationDate).toLocaleDateString('en-GB', { day: 'numeric', month: 'long', year: 'numeric' })}</p>
        </div>
      )}

      <div className="postsGrid">
        {posts.length > 0 ? (
          posts.map((post) => {
            const itinerary = GetItinerary(post.itineraryId);
            const eventImage = GetEvent(itinerary.events[0]).EventImages[0]
              ? GetEvent(itinerary.events[0]).EventImages[0]
              : [fallbackImage]; // Use fallback if no images available

            return (
            <div key={post.postId} className="postCard" onClick={() => handleBoardClick(post.postId)}>
                 <img src={eventImage} alt={boardDetails.name} className="board-image" />
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
                    <span>{post.comments.length} Comments</span>
                  </div>
                </div>
                </div>
            </div>
            );
          })
        ) : (
          <div className="noPostsMessage">No posts available in this board.</div>
        )}
      </div>
    </div>
  );
}

export default BoardPosts;
