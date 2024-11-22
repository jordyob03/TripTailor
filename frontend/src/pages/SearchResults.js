import React, { useRef } from 'react';
import '../styles/styles.css';
import ItineraryGrid from '../components/ItineraryGrid.js'; 

function SearchResults({ searchResults = [], searchParams = {}, isSearchPressed = false }) {
  const tagContainerRef = useRef(null);

  return (
    <div className="searchResultsContainer">
      {/* Display search parameters if provided */}
      {searchParams.city && searchParams.country && (
        <h2 className="searchResultsHeader">
          Search Results for {searchParams.city.charAt(0).toUpperCase() + searchParams.city.slice(1)}, {searchParams.country.charAt(0).toUpperCase() + searchParams.country.slice(1)}
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
