import React, { useState } from 'react';
import logo1 from '../assets/logo-long-transparent.png';

function InitialUserProfile() {
  // Tags
  const [selectedTags, setSelectedTags] = useState([]);
  const [errorMessage, setErrorMessage] = useState('');
  const [country, setCountry] = useState('');
  const [language, setLanguage] = useState('');
  const tags = ['Beach', 'Indoors', 'Outdoors', 'City', 'Nightlife', 'Hiking', 'Boating', 'Relaxing', 'Wildlife', 'Shopping', 'Road Trip', 'Sports', 'Arts & Architexture', 'Festivals & Events', 'Backpacking', 'Museums', 'National Park', 'Landmarks & Historical Sites', 'Food', 'Theme Park'];

  // Options for dropdowns
  const countries = ['USA', 'Canada', 'UK', 'Australia', 'Other'];
  const languages = ['English', 'Spanish', 'French', 'German', 'Chinese'];

  // Handle form submission
  const handleSubmit = (e) => {
    e.preventDefault();
    if (selectedTags.length === 3) {
      console.log({ selectedTags, country, language });
      setErrorMessage(''); // Clear error message if valid
    } else {
      setErrorMessage('Please select exactly 3 tags.');
    }
  };

  // Handle tag selection (max of 3 tags)
  const handleTagSelection = (tag) => {
    if (selectedTags.includes(tag)) {
      setSelectedTags(selectedTags.filter((t) => t !== tag));
    } else if (selectedTags.length < 3) {
      setSelectedTags([...selectedTags, tag]);
    }
  };

  // Inline styles
  const styles = {
    navbar: {
      display: 'flex',
      justifyContent: 'space-between', 
      alignItems: 'center',
      padding: '10px 20px',
      height: '60px',
      width: '100vw', 
      backgroundColor: 'white',
      boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)', 
      position: 'fixed', 
      top: 0, 
      left: 0,
      zIndex: 2,
    },
    logo1: {
      width: '150px',
      marginLeft: '80px',
      marginTop: '5px',
    },
    profileButton: {
      display: 'flex',        
      alignItems: 'center',   
      justifyContent: 'space-between', 
      padding: '10px 20px',
      backgroundColor: 'white',
      border: '1px solid #dfdfdf', 
      borderRadius: '30px',
      cursor: 'pointer',
      marginRight: '160px',
      boxShadow: '0 2px 2px rgba(0, 0, 0, 0.1)',
    },
    container: {
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
      minHeight: 'calc(100vh - 60px)', 
      width: '100vw',
      backgroundColor: 'white', 
    },
    box: {
      backgroundColor: 'white',
      padding: '40px',
      borderRadius: '20px',
      boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)',
      textAlign: 'center',
      width: '650px',
      zIndex: 1,
      marginTop: '80px', 
      marginBottom: '40px', 
    },
    heading: {
      color: '#002f6c', 
      marginTop: '5px', 
      marginBottom: '10px', 
      fontFamily: "'Red Hat Display', sans-serif",
    },

    subheading1: {
      color: '#002f6c', 
      marginTop: '40px', 
      marginBottom: '20px', 
      fontFamily: "'Red Hat Display', sans-serif",
      fontSize: '20px',
      fontWeight: '400'
    },
    separator: {
      width: '100%',      
      border: 'none',   
      borderTop: '2px solid #d4d4d4', 
      margin: '15px auto',  
    },
    tags: {
      display: 'flex',
      flexWrap: 'wrap',
      gap: '10px',
      justifyContent: 'center',
      marginBottom: '20px',
    },
    tag: (isSelected) => ({
      padding: '8px 16px',
      borderRadius: '50px',
      border: isSelected ? '2px solid #00509e' : '1px solid #ccc',
      backgroundColor: isSelected ? '#002f6c' : '#C6DFF0',
      cursor: 'pointer',
      transition: 'all 0.3s ease',
      boxShadow: isSelected ? '0 2px 4px rgba(0, 80, 158, 0.3)' : '0 2px 4px rgba(255, 255, 255, 0.3)',
      fontSize: '19px',
      fontWeight: '700',
      color: isSelected ? 'white' : '#002f6c'
    }),
    button: {
      padding: '12px',
      backgroundColor: '#002f6c',
      color: 'white',
      border: 'none',
      borderRadius: '4px',
      fontSize: '16px',
      cursor: 'pointer',
      transition: 'background-color 0.3s ease',
      fontFamily: "'Red Hat Display', sans-serif",
    },
    errorMessage: {
      color: 'red',
      fontSize: '14px',
      marginTop: '-10px',
    },
    dropdown: {
      width: '100%',
      padding: '10px',
      borderRadius: '5px',
      border: '1px solid #ccc',
      marginBottom: '20px',
      fontFamily: "'Red Hat Display', sans-serif",
      fontSize: '16px',
    },
  };

  return (
    <>
      {/* Navbar */}
      <nav style={styles.navbar}>
        <img src={logo1} alt="Trip Tailor Logo" style={styles.logo1} />

        {/* Profile Button */}
        <button style={styles.profileButton}>
          <i className="fas fa-bars" style={{ fontSize: '16px', color: '#00509e', marginRight: '15px' }}></i> 
          <i className="fa-regular fa-user" style={{ fontSize: '24px', color: '#00509e' }}></i>
        </button>
      </nav>

      {/* Main Container */}
      <div style={styles.container}>
        <div style={styles.box}>
          <h5 style={styles.heading}>Tell us more about you</h5>
          <h6 style={styles.subheading1}>What tags are important to you on your travels?</h6>
          
          {/* Tag selection */}
          <div style={styles.tags}>
            {tags.map((tag) => (
              <div
                key={tag}
                style={styles.tag(selectedTags.includes(tag))}
                onClick={() => handleTagSelection(tag)}
              >
                {tag}
              </div>
            ))}
          </div>
          <h6 style={styles.subheading1}>Where do you live?</h6>
      <select
        value={country}
        onChange={(e) => setCountry(e.target.value)}
        style={styles.dropdown}
      >
        <option value="">Select a country</option>
        {countries.map((country) => (
          <option key={country} value={country}>
            {country}
          </option>
        ))}
      </select>

      <h6 style={styles.subheading1}>What is your preferred language?</h6>
      <select
        value={language}
        onChange={(e) => setLanguage(e.target.value)}
        style={styles.dropdown}
      >
        <option value="">Select a language</option>
        {languages.map((language) => (
          <option key={language} value={language}>
            {language}
          </option>
        ))}
      </select>


          {/* Error message */}
          {errorMessage && <div style={styles.errorMessage}>{errorMessage}</div>}

          <button
            type="submit"
            style={styles.button}
            onClick={handleSubmit}
          >
            Continue
          </button>
        </div>
      </div>
    </>
  );
}

export default InitialUserProfile;
