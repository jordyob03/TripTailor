import React, { useState, useRef } from 'react';
import '../styles/styles.css';
import Tags from '../config/tags.json';
import iconMap from '../config/iconMap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

function HomePage() {
  const [selectedTags, setSelectedTags] = useState([]); // Manage selected tags state
  const [tagErrorMessage, setTagErrorMessage] = useState(''); // Manage error message state
  const tagContainerRef = useRef(null);
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
    // Toggle the selection of the tag
    if (selectedTags.includes(tag)) {
      setSelectedTags(selectedTags.filter((t) => t !== tag));
    } else {
      setSelectedTags([...selectedTags, tag]);
    }
    setTagErrorMessage(''); // Clear error message when a tag is clicked
  };

  return (
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
            <div className="tagLabel">
              {tag}
            </div>
          </div>
        ))}
      </div>
      <button onClick={scrollTagsRight} className="arrowButton">{'>'}</button>
      {tagErrorMessage && <div className="errorMessage">{tagErrorMessage}</div>}
    </div>
  );
}

export default HomePage;
