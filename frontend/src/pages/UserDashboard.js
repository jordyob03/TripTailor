import React from 'react';
import { Link, Route, Routes, useLocation } from 'react-router-dom';
import Boards from './UserBoards';
import Itineraries from './UserItineraries';
import '../styles/styles.css';

function MyTravels() {
  const location = useLocation(); // Get the current path

  return (
    <div className="pageContainer">
      <div className="heading">
        <h2>My Travels</h2>
      </div>

      <div className="tabNavigation">
        <Link to="/my-travels/itineraries">
          <button
            className={location.pathname === '/my-travels/itineraries' ? 'activeTab' : ''}
          >
            Itineraries
          </button>
        </Link>
        <Link to="/my-travels/boards">
          <button
            className={location.pathname === '/my-travels/boards' ? 'activeTab' : ''}
          >
            Boards
          </button>
        </Link>
      </div>

      <Routes>
        <Route path="itineraries" element={<Itineraries />} />
        <Route path="boards" element={<Boards />} />
      </Routes>
    </div>
  );
}

export default MyTravels;
