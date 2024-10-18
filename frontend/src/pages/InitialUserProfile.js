import React, { useState, useEffect } from 'react';
import navBarLogo from '../assets/logo-long-transparent.png';
import Tags from '../config/tags.json';
import '../styles/styles.css';  // Import your external CSS

function InitialUserProfile() {
  const allTags = Object.values(Tags.categories).flat();

  const [selectedTags, setSelectedTags] = useState([]);
  const [tagErrorMessage, setTagErrorMessage] = useState('');
  const [countryErrorMessage, setCountryErrorMessage] = useState('');
  const [langErrorMessage, setLangErrorMessage] = useState('');
  const [country, setCountry] = useState('');
  const [languages, setLanguages] = useState([]);
  const [shuffledTags, setShuffledTags] = useState([]);
  const [name, setName] = useState('');
  const [nameErrorMessage, setNameErrorMessage] = useState('');

  const countries = ['USA', 'Canada', 'UK', 'Australia', 'Other'];
  const languageOptions = ['English', 'Spanish', 'French', 'German', 'Chinese'];

  const shuffleArray = (array) => {
    return [...array].sort(() => Math.random() - 0.5);
  };

  useEffect(() => {
    const shuffled = shuffleArray(allTags);
    setShuffledTags(shuffled);
  }, []);

  const handleSubmit = (e) => {
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
      setLangErrorMessage('')
    } else {
      setLangErrorMessage('Please select at least 1 language.');
    }

    if (name && selectedTags.length >= 3 && country.length >= 1 && languages.length >= 1) {
      console.log({ name, selectedTags, country, languages });
    }
  };

  const handleTagSelection = (tag) => {
    if (selectedTags.includes(tag)) {
      setSelectedTags(selectedTags.filter((t) => t !== tag));
    } else {
      setSelectedTags([...selectedTags, tag]);
    }
  };

  const handleLanguageSelection = (e) => {
    const selectedOptions = Array.from(e.target.selectedOptions, option => option.value);
    setLanguages(selectedOptions); 
  };

  return (
    <>
      {/* navBar */}
      <nav className="navBar">
        <img src={navBarLogo} alt="Trip Tailor Logo" className="navBarLogo" />

        {/* Profile Button */}
        <button className="profileButton">
          <i className="fas fa-bars" style={{ fontSize: '16px', color: '#00509e', marginRight: '15px' }}></i>
          <i className="fa-regular fa-user" style={{ fontSize: '24px', color: '#00509e' }}></i>
        </button>
      </nav>

      {/* Main Container */}
      <div className="centeredContainer">
        <div className="centeredBox">
          <h5 className="heading">Tell us more about you</h5>

          <h6 className="subheadingIUP">What's your name?</h6>

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

          <h6 className="subheadingIUP">What tags are important to you on your travels?</h6>

          {/* Error message above tags */}
          {tagErrorMessage && <div className="errorMessage">{tagErrorMessage}</div>}

          {/* Tag selection */}
          <div className="tags">
            {shuffledTags.map((tag) => (
              <div
                key={tag}
                className={`tag ${selectedTags.includes(tag) ? 'selected' : ''}`}
                onClick={() => handleTagSelection(tag)}
              >
                {tag}
              </div>
            ))}
          </div>

          <h6 className="subheadingIUP">Where do you live?</h6>

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

          <h6 className="subheadingIUP">What languages do you speak?</h6>

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

          <button type="submit" className="continueButtonIUP" onClick={handleSubmit}>
            Continue
          </button>
        </div>
      </div>
    </>
  );
}

export default InitialUserProfile;
