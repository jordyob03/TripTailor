import React, { useState, useEffect } from 'react';
import Tags from '../config/tags.json';
import '../styles/styles.css';  
import { useNavigate } from 'react-router';
import profileAPI from '../api/profileAPI.js';

function EditUserProfile() {
  const allTags = Object.values(Tags.categories).flat();
  const [name, setName] = useState('');
  const [selectedTags, setSelectedTags] = useState([]);
  const [country, setCountry] = useState('');
  const [languages, setLanguages] = useState([]);
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

  // Fetch user data from backend
  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await profileAPI.get(`/user`, { params: { username } }); // Idk if this endpoint exists
        const userData = response.data;
        setName(userData.name || '');
        setSelectedTags(userData.tags || []);
        setCountry(userData.country || '');
        setLanguages(userData.languages || []);
      } catch (error) {
        console.error("Failed to fetch user data:", error);
        setErrorMessage("Failed to load user profile");
      }
    };

    fetchUserData();

    // Shuffle available tags
    const availableTags = allTags.filter(tag => !selectedTags.includes(tag));
    setShuffledTags(shuffleArray(availableTags));
  }, [username]);

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
      const profileData = {
        name,
        country,
        languages,
        tags: selectedTags,
      };
  
      try {
        await profileAPI.put(`/user`, profileData, { params: { username } });
        navigate('/home-page');
      } catch (error) {
        console.error("Failed to save profile:", error);
        setErrorMessage("Failed to save profile changes");
      }
    }
  };
  
  const handleTagSelection = (tag) => {
    if (selectedTags.includes(tag)) {
      setSelectedTags(selectedTags.filter((t) => t !== tag));
      setShuffledTags([...shuffledTags, tag]);
    } else {
      setSelectedTags([...selectedTags, tag]);
      setShuffledTags(shuffledTags.filter((t) => t !== tag));
    }
  };

  const handleLanguageSelection = (e) => {
    const selectedOptions = Array.from(e.target.selectedOptions, option => option.value);
    setLanguages(selectedOptions); 
  };

  return (
    <div className="centeredContainer">
      <div className="centeredBox">
        <h5 className="heading">Edit Personal Info</h5>
        {errorMessage && <div className="errorMessage">{errorMessage}</div>}

        <h6 className="subheadingIUP">Name</h6>
        {nameErrorMessage && <div className="errorMessage">{nameErrorMessage}</div>}
        <input
          type="text"
          placeholder="Enter your name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
          className="input"
        />

        <h6 className="subheadingIUP">Interested tags</h6>
        {tagErrorMessage && <div className="errorMessage">{tagErrorMessage}</div>}
        <div className="tags">
          {selectedTags.map((tag) => (
            <div key={tag} className="tag selected" onClick={() => handleTagSelection(tag)}>
              {tag} <span style={{ color: 'white' }}>✖</span>
            </div>
          ))}
          {shuffledTags.map((tag) => (
            <div key={tag} className="tag" onClick={() => handleTagSelection(tag)}>
              <span style={{ color: '' }}>+</span> {tag}
            </div>
          ))}
        </div>

        <h6 className="subheadingIUP">Country of residence</h6>
        {countryErrorMessage && <div className="errorMessage">{countryErrorMessage}</div>}
        <select value={country} onChange={(e) => setCountry(e.target.value)} className="dropdown">
          <option value="">Select a country</option>
          {countries.map((country) => (
            <option key={country} value={country}>{country}</option>
          ))}
        </select>

        <h6 className="subheadingIUP">Spoken languages</h6>
        {langErrorMessage && <div className="errorMessage">{langErrorMessage}</div>}
        <select multiple value={languages} onChange={handleLanguageSelection} className="dropdown">
          {languageOptions.map((language) => (
            <option key={language} value={language}>{language}</option>
          ))}
        </select>

        <button type="submit" className="continueButton" onClick={handleSubmit}>
          Save
        </button>
      </div>
    </div>
  );
}

export default EditUserProfile;
