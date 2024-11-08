import React, { useState } from 'react';
import { Link, Route, Routes } from 'react-router-dom';
import Boards from './UserBoards';
import Itineraries from './UserItineraries';
import '../styles/styles.css';

function MyTravels() {
  const [selectedTab, setSelectedTab] = useState('itineraries');

  return (
    <div className="pageContainer">
      <div className="heading">
        <h2>My Travels</h2>
      </div>

      <div className="tabNavigation">
        <Link to="/my-travels/itineraries">
          <button className={selectedTab === 'itineraries' ? 'activeTab' : ''} onClick={() => setSelectedTab('itineraries')}>
            Itineraries
          </button>
        </Link>
        <Link to="/my-travels/boards">
          <button className={selectedTab === 'boards' ? 'activeTab' : ''} onClick={() => setSelectedTab('boards')}>
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
