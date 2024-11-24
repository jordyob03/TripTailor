import React, { useRef } from 'react';
import '../styles/styles.css';
import ItineraryGrid from '../components/ItineraryGrid.js'; 

function SearchResults({ searchResults = [], searchParams = {}, isSearchPressed = false }) {
  const tagContainerRef = useRef(null);

  return (
    <div className="searchResultsContainer">
      {/* Display search parameters if provided */}
      {searchParams.Price && searchParams.SearchValue && (
        <h2 className="searchResultsHeader">
          Search Results for {searchParams.Price.charAt(0).toUpperCase() + searchParams.Price.slice(1)}, {searchParams.SearchValue.charAt(0).toUpperCase() + searchParams.SearchValue.slice(1)}
        </h2>
      )}
      {searchResults.length > 0 ? (
        <ItineraryGrid itineraries={searchResults} showSaveButton={true} />
      ) : (
        // Display a message if no results are found
        isSearchPressed && <p className="noResultsMessage">No search results found.</p>
      )}
    </div>
  );
}

export default SearchResults;
