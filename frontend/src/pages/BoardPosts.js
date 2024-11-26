import React, { useState, useEffect } from 'react';
import boardAPI from '../api/boardAPI.js';
import ItineraryGrid from '../components/ItineraryGrid.js';
import { useLocation, useNavigate } from 'react-router-dom';
import '../styles/styles.css';

function BoardPosts() {
  const location = useLocation();
  const navigate = useNavigate();

  const boardId = parseInt(location.pathname.split("/").pop(), 10);

  const [selectedBoard, setSelectedBoard] = useState(null);
  const [itineraries, setItineraries] = useState([]);
  const [errorMessage, setErrorMessage] = useState('');
  const [editMode, setEditMode] = useState(false); // Tracks edit mode
  const [editedDescription, setEditedDescription] = useState(''); // For editing description
  const [posts, setPosts] = useState([]);

  const fetchBoardData = async () => {
    try {
      const userData = { username: localStorage.getItem('username') };
      const boardsResponse = await boardAPI.get('/boards', { params: userData });
      const board = boardsResponse.data.boards.find((b) => b.boardId === boardId);

      if (!board) {
        setErrorMessage('Board not found.');
        return;
      }

      setSelectedBoard({
        name: board.name,
        description: board.description,
        username: board.username,
        dateOfCreation: board.dateOfCreation,
      });

      setEditedDescription(board.description); 

      const postsResponse = await boardAPI.get('/posts', { params: { boardId } });
      const postIds = postsResponse.data.Posts.map((post) => post.postId);

      const itinerariesData = [];
      for (const postId of postIds) {
        const itineraryResponse = await boardAPI.get('/itineraries', { params: { postId } });
        const itinerary = itineraryResponse.data.Itinerary;
        itinerariesData.push({ ...itinerary, postId, boardId }); 
      }

      setItineraries(itinerariesData);
    } catch (error) {
      console.error('Error fetching board data:', error);
      setErrorMessage('Failed to fetch board data.');
    }
  };

  const handleDeletePost = async (boardId, postId) => {
    try {
      await boardAPI.delete(`/boards/${boardId}/posts/${postId}`);
      setPosts(posts.filter(post => post.postId !== postId));
      window.location.reload();
    } catch (error) {
      console.error("Error deleting post:", error);
      setErrorMessage("Failed to delete post");
    }
  }; 

  const handleSaveChanges = async () => {
    try {
      const updatedBoard = {
        boardId: boardId,
        name: selectedBoard.name,
        description: editedDescription,
      };

      await boardAPI.post('/editboard', updatedBoard);

      setSelectedBoard((prev) => ({
        ...prev,
        description: editedDescription,
      }));

      setEditMode(false); // Exit edit mode
      
    } catch (error) {
      console.error('Error saving changes:', error);
      setErrorMessage('Failed to save changes.');
    }
  };

  useEffect(() => {
    if (isNaN(boardId)) {
      setErrorMessage('Invalid board ID.');
      return;
    }

    fetchBoardData();
  }, [boardId]);

  return (
    <div className="pageContainer">
      <button
        className="backButton"
        onClick={() => navigate('/my-travels/boards')}
      >
        {"< Back to My Travels"}
      </button>
      {selectedBoard && (
        <div className="boardDetails">
          <h2 style={{ marginTop: '0px' }}>{selectedBoard.name}</h2>
          <p>
            Created by {selectedBoard.username} on{' '}
            {new Date(selectedBoard.dateOfCreation).toLocaleDateString('en-GB', {
              day: 'numeric',
              month: 'long',
              year: 'numeric',
            })}
          </p>
          {editMode ? (
            <>
              <textarea
                className="editDescriptionField"
                value={editedDescription}
                onChange={(e) => setEditedDescription(e.target.value)}
              />
              <br />
              <button className="editButton" onClick={() => handleSaveChanges()}>
                Save
              </button>
            </>
          ) : (
            <>
              <p>{selectedBoard.description}</p>
              <button className="editButton" onClick={() => setEditMode(true)}>
                Edit
              </button>
            </>
          )}
        </div>
      )}

      <div className="results">
        {itineraries.length > 0 ? (
          <ItineraryGrid
            itineraries={itineraries}
            showSaveButton={false}
            editMode={editMode} // Pass edit mode to ItineraryGrid
            onDeletePost={handleDeletePost} // Pass delete handler
          />
        ) : errorMessage ? (
          <p className="message">{errorMessage}</p>
        ) : (
          <p className="noResultsMessage">No itineraries found for this board.</p>
        )}
      </div>
    </div>
  );
}

export default BoardPosts;
