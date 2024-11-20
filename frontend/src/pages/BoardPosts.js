import React, { useState, useEffect } from 'react';
import boardAPI from '../api/boardAPI.js';
import ItineraryGrid from '../components/ItineraryGrid.js';
import { useLocation } from 'react-router-dom';
import '../styles/styles.css';

function BoardPosts() {
  const location = useLocation();
  const boardId = parseInt(location.pathname.split("/").pop(), 10);

  const [selectedBoard, setSelectedBoard] = useState(null);
  const [itineraries, setItineraries] = useState([]);
  const [errorMessage, setErrorMessage] = useState('');

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

      const postsResponse = await boardAPI.get('/posts', { params: { boardId } });
      const postIds = postsResponse.data.Posts.map((post) => post.postId);

      const itinerariesData = [];
      for (const postId of postIds) {
        const itineraryResponse = await boardAPI.get('/itineraries', { params: { postId } });
        itinerariesData.push(itineraryResponse.data.Itinerary);
      }

      setItineraries(itinerariesData);
    } catch (error) {
      console.error('Error fetching board data:', error);
      setErrorMessage('Failed to fetch board data.');
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
      {selectedBoard && (
        <div className="boardDetails">
          <h2>{selectedBoard.name}</h2>
          <p>{selectedBoard.description}</p>
          <p>
            Created by {selectedBoard.username} on{' '}
            {new Date(selectedBoard.dateOfCreation).toLocaleDateString('en-GB', {
              day: 'numeric',
              month: 'long',
              year: 'numeric',
            })}
          </p>
        </div>
      )}

      <div className="results">
        {itineraries.length > 0 ? (
          <ItineraryGrid itineraries={itineraries} showSaveButton={false} />
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
