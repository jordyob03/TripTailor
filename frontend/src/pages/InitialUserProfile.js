import React, { useState, useEffect } from 'react';
import navBarLogo from '../assets/logo-long-transparent.png';
import Tags from '../config/tags.json';
import '../styles/styles.css';  
import { useNavigate } from 'react-router';

function InitialUserProfile() {
  const allTags = Object.values(Tags.categories).flat();

  const [selectedTags, setSelectedTags] = useState([]);
  const [tagErrorMessage, setTagErrorMessage] = useState('');
  const [countryErrorMessage, setCountryErrorMessage] = useState('');
  const [langErrorMessage, setLangErrorMessage] = useState('');
  const [country, setCountry] = useState('');
  const [languages, setLanguages] = useState([]);
  const [shuffledTags, setShuffledTags] = useState([]);

  const countries = ['USA', 'Canada', 'UK', 'Australia', 'Other'];
  const languageOptions = ['English', 'Spanish', 'French', 'German', 'Chinese'];

  const shuffleArray = (array) => {
    return [...array].sort(() => Math.random() - 0.5);
  };

  const navigate = useNavigate()

  useEffect(() => {
    const shuffled = shuffleArray(allTags);
    setShuffledTags(shuffled);
  }, []);

  const handleSubmit = (e) => {
    e.preventDefault();
    if (selectedTags.length >= 3) {
      console.log({ selectedTags, country, languages });
      setTagErrorMessage(''); // Clear error message if valid
    } else {
      console.log({ selectedTags, country, languages });
      setTagErrorMessage('Please select at least 3 tags.');
    }
    if (country.length >= 1) {
      console.log({ selectedTags, country, languages });
      setCountryErrorMessage(''); // Clear error message if valid
    } else {
      setCountryErrorMessage('Please select a country.');
    }
    if (languages.length >= 1) {
      console.log({ selectedTags, country, languages });
      setLangErrorMessage(''); 
    } else {
      setLangErrorMessage('Please select at least 1 language.');
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

  const handleLogout = () => {
    localStorage.removeItem('token'); 
    navigate('/'); 
  };

  return (
    <>
      {/* navBar */}
      <nav className="navBar">
        <img src={navBarLogo} alt="Trip Tailor Logo" className="navBarLogo" />
          <div className="buttonsContainer">
          {/* Logout Button */}
          <button className="logoutButton" onClick={handleLogout}>
            <i className="fas fa-sign-out-alt" style={{ fontSize: '24px', color: '#00509e', marginLeft: '5px', marginRight: '10px' }}></i>
            Log Out
          </button>
          {/* Profile Button */}
          <button className="profileButton">
            <i className="fas fa-bars" style={{ fontSize: '16px', color: '#00509e', marginRight: '15px' }}></i>
            <i className="fa-regular fa-user" style={{ fontSize: '24px', color: '#00509e' }}></i>
          </button>
        </div>
      </nav>

      {/* Main Container */}
      <div className="centeredContainer">
        <div className="centeredBox">
          <h5 className="heading">Tell us more about you</h5>
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
                    
          {/* Error message above tags */}
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

          {/* Error message above tags */}
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
