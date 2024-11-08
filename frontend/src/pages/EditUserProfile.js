import React, { useState, useEffect } from 'react';
import Tags from '../config/tags.json';
import '../styles/styles.css';  
import { useNavigate } from 'react-router';
import profileAPI from '../api/profileAPI.js';

function InitialUserProfile() {
  const allTags = Object.values(Tags.categories).flat();

  // Sample data to populate fields until data is available from the backend
  const sampleUserData = {
    name: 'John Doe',
    selectedTags: ['Adventure', 'Culture', 'Relaxation'],
    country: 'USA',
    languages: ['English', 'Spanish']
  };

  // Initial state for user profile
  const [name, setName] = useState(sampleUserData.name);
  const [selectedTags, setSelectedTags] = useState(sampleUserData.selectedTags);
  const [country, setCountry] = useState(sampleUserData.country);
  const [languages, setLanguages] = useState(sampleUserData.languages);
  const [shuffledTags, setShuffledTags] = useState([]);
  const [tagErrorMessage, setTagErrorMessage] = useState('');
  const [countryErrorMessage, setCountryErrorMessage] = useState('');
  const [langErrorMessage, setLangErrorMessage] = useState('');
  const [nameErrorMessage, setNameErrorMessage] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();
  const countries = ['USA', 'Canada', 'UK', 'Australia', 'Other'];
  const languageOptions = ['English', 'Spanish', 'French', 'German', 'Chinese'];
  const username = localStorage.getItem('username');

  // Function to shuffle tags
  const shuffleArray = (array) => {
    return [...array].sort(() => Math.random() - 0.5);
  };

  // Shuffle tags once when the component mounts
  useEffect(() => {
    // Filter out the selected tags and shuffle the rest
    const availableTags = allTags.filter(tag => !selectedTags.includes(tag));
    const shuffled = shuffleArray(availableTags);
    setShuffledTags(shuffled);
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
  
    if (!name) {
      setNameErrorMessage('Please enter your name.');
      return;
    } else {
      setNameErrorMessage('');
    }
  
    if (selectedTags.length >= 3) {
      setTagErrorMessage('');
    } else {
      setTagErrorMessage('Please select at least 3 tags.');
    }
    
    if (country.length >= 1) {
      setCountryErrorMessage('');
    } else {
      setCountryErrorMessage('Please select a country.');
    }
  
    if (languages.length >= 1) {
      setLangErrorMessage('');
    } else {
      setLangErrorMessage('Please select at least 1 language.');
    }
  
    if (name && selectedTags.length >= 3 && country.length >= 1 && languages.length >= 1) {
      const profile_data = {
        languages: languages,
        country: country, 
        tags: selectedTags,
        name: name, 
        username: username, 
      };
  
      try {
        const response = await profileAPI.post('/create', profile_data);
        navigate('/home-page');
      } catch (error) {
        if (error.response && error.response.data) {
          setErrorMessage(error.response.data.error);  
        } else {
          setErrorMessage('Saving profile failed');
        }
      }
    }
  };
  
  const handleTagSelection = (tag) => {
    if (selectedTags.includes(tag)) {
      // Remove the tag from selectedTags and add it back to shuffledTags
      setSelectedTags(selectedTags.filter((t) => t !== tag));
      setShuffledTags([...shuffledTags, tag]);
    } else {
      // Add the tag to selectedTags and remove it from shuffledTags
      setSelectedTags([...selectedTags, tag]);
      setShuffledTags(shuffledTags.filter((t) => t !== tag));
    }
  };
  

  const handleLanguageSelection = (e) => {
    const selectedOptions = Array.from(e.target.selectedOptions, option => option.value);
    setLanguages(selectedOptions); 
  };

  return (
    <>
      {/* Main Container */}
      <div className="centeredContainer">
        <div className="centeredBox">
          <h5 className="heading">Edit Personal Info</h5>

          <h6 className="subheadingIUP">Name</h6>

          {/* Error message above name */}
          {nameErrorMessage && <div className="errorMessage">{nameErrorMessage}</div>}

          {/* Name input */}
          <input
            type="text"
            placeholder="Enter your name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
            className="input"
          />

          <h6 className="subheadingIUP">Interested tags</h6>

          {/* Error message above tags */}
          {tagErrorMessage && <div className="errorMessage">{tagErrorMessage}</div>}

          {/* Tag selection */}
          <div className="tags">
            {/* Display selected tags first */}
            {selectedTags.map((tag) => (
              <div
                key={tag}
                className="tag selected"
                onClick={() => handleTagSelection(tag)}
              >
                {tag} <span style={{ color: 'white' }}>✖</span> {/* White X icon */}
              </div>
            ))}
            {/* Display available tags to add */}
            {shuffledTags.map((tag) => (
              <div
                key={tag}
                className="tag"
                onClick={() => handleTagSelection(tag)}
              >
                <span style={{ color: '' }}> +</span> {tag} {/* Yellow Plus icon */}
              </div>
            ))}
          </div>

          <h6 className="subheadingIUP">Country of residence</h6>

          {/* Error message above country */}
          {countryErrorMessage && <div className="errorMessage">{countryErrorMessage}</div>}

          <select
            value={country}
            onChange={(e) => setCountry(e.target.value)}
            className="dropdown"
          >
            <option value="">Select a country</option>
            {countries.map((country) => (
              <option key={country} value={country}>
                {country}
              </option>
            ))}
          </select>

          <h6 className="subheadingIUP">Spoken languages</h6>

          {/* Error message above languages */}
          {langErrorMessage && <div className="errorMessage">{langErrorMessage}</div>}

          <select
            multiple
            value={languages}
            onChange={handleLanguageSelection}
            className="dropdown"
          >
            {languageOptions.map((language) => (
              <option key={language} value={language}>
                {language}
              </option>
            ))}
          </select>

          <button type="submit" className="continueButton" onClick={handleSubmit}>
            Continue
          </button>
        </div>
      </div>
    </>
  );
}

export default InitialUserProfile;