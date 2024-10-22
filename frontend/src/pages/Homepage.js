import React, { useRef } from 'react';
import '../styles/styles.css';
import Tags from '../config/tags.json';
import iconMap from '../config/iconMap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

function HomePage({ selectedTags, onTagClick, tagErrorMessage }) {
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

  return (
    <div className="tagFiltersContainer">
      <button onClick={scrollTagsLeft} className="arrowButton">{'<'}</button>
      <div ref={tagContainerRef} className="tagContainer">
        {allTags.map((tag) => (
          <div
            key={tag}
            onClick={() => onTagClick(tag)}
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
