import React, { useState, useEffect, useRef } from 'react';
import '../styles/styles.css';
import Tags from '../config/tags.json';
import iconMap from '../config/iconMap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

function HomePage() {
  const [selectedTags, setSelectedTags] = useState([]);
  const [itineraries, setItineraries] = useState([]); 
  const tagContainerRef = useRef(null);

  useEffect(() => {
    // Fetch itineraries from the backend
    fetch('/api/itineraries') // Replace 
      .then((response) => response.json())
      .then((data) => setItineraries(data))
      .catch((error) => console.error('Error fetching itineraries:', error));
  }, []);

  const allTags = Object.values(Tags.categories).flat();

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
    setSelectedTags((prevSelectedTags) =>
      prevSelectedTags.includes(tag)
        ? prevSelectedTags.filter((t) => t !== tag)
        : [...prevSelectedTags, tag]
    );
  };

  // Filter itineraries based on selected tags
  const filteredItineraries = selectedTags.length
    ? itineraries.filter((itinerary) =>
        itinerary.tags.some((tag) => selectedTags.includes(tag))
      )
    : itineraries;

  return (
    <div>
      {/* Tag Filter Scrollable Container */}
      <div className="tagFiltersContainer">
        <button onClick={scrollTagsLeft} className="arrowButton">{'<'}</button>
        <div ref={tagContainerRef} className="tagContainer">
          {allTags.map((tag) => (
            <div
              key={tag}
              onClick={() => handleTagClick(tag)}
              className={`tagItem ${selectedTags.includes(tag) ? 'selected' : ''}`}
            >
              <div className="tagIcon">
                {iconMap[tag] && <FontAwesomeIcon icon={iconMap[tag]} />}
              </div>
              <div className="tagLabel">{tag}</div>
            </div>
          ))}
        </div>
        <button onClick={scrollTagsRight} className="arrowButton">{'>'}</button>
      </div>

      {/* Feed Section: Filtered Itineraries */}
      <div className="resultsGrid" style={{ marginTop: '200px', padding: '0 70px' }}>
        {filteredItineraries.map((itinerary) => (
          <div key={itinerary.id} className="resultCard">
            <img src={itinerary.image} alt={itinerary.title} className="resultCardImage" />
            <div className="resultCardContent">
              <h3 className="resultCardTitle">{itinerary.title}</h3>
              <p className="cardLocation">{itinerary.location}</p>
              <p className="resultCardDescription">{itinerary.description}</p>
              <div className="resultTags">
                {itinerary.tags.map((tag) => (
                  <span key={tag} className="resultCardTag">{tag}</span>
                ))}
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default HomePage;
