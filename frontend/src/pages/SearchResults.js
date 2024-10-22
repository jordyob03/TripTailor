import React, { useState, useRef } from 'react';
import '../styles/styles.css';
import Tags from '../config/tags.json';
import iconMap from '../config/iconMap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

function SearchResults({ searchResults = [], searchParams = {}, isSearchPressed = false }) {
  const tagContainerRef = useRef(null);
  const allTags = Object.values(Tags.categories).flat();
  const [selectedTags, setSelectedTags] = useState([]); 
  const [tagErrorMessage, setTagErrorMessage] = useState('');

  const scrollTagsLeft = () => {
    if (tagContainerRef.current) {
      tagContainerRef.current.scrollBy({ left: -150, behavior: 'smooth' });
    }
  };

  const scrollTagsRight = () => {
    if (tagContainerRef.current) {
      tagContainerRef.current.scrollBy({ left: 150, behavior: 'smooth' });
    }
  };

  const handleTagClick = (tag) => {
    // Toggle the selection of the tag
    
    if (selectedTags.includes(tag)) {
      setSelectedTags(selectedTags.filter((t) => t !== tag));
    } else {
      setSelectedTags([...selectedTags, tag]);
    }
    setTagErrorMessage('');
  };

  return (
    <div className="searchResultsContainer">

      {searchParams.city && searchParams.country && (
        <h2 className="searchResultsHeader">
          Search Results for {searchParams.city.charAt(0).toUpperCase() + searchParams.city.slice(1)}, {searchParams.country.charAt(0).toUpperCase() + searchParams.country.slice(1)}
        </h2>
      )}
      {searchResults.length > 0 ? (
        <div className="resultsGrid">
          {searchResults.map((result, index) => (
            <div key={index} className="resultCard">
              <img src={result.image} alt={result.title} className="resultCardImage" />
              <div className="resultCardContent">
              <h4 className="cardLocation">
                {result.location
                  .split(' ')
                  .map(word => word.charAt(0).toUpperCase() + word.slice(1))
                  .join(' ')}
              </h4>
                <h3 className="cardTitle" > {result.title}</h3>
                <p className="cardDescription">{result.description}</p>
                <div className="resultTags">
                  {result.tags.map((tag, i) => (
                    <span key={i} className="resultCardTag">
                      {tag}
                    </span>
                  ))}
                </div>
              </div>
            </div>
          ))}
        </div>
      ) : (
        isSearchPressed && <p className="noResultsMessage">No search results found.</p>
      )}
    </div>
  );
}

export default SearchResults;
