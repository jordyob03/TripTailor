import React, { useState } from 'react';
import logo1 from '../assets/logo-long-transparent.png';

function InitialUserProfile() {
  const [selectedTags, setSelectedTags] = useState([]);
  const [errorMessage, setErrorMessage] = useState('');

  // Tags
  const tags = ['Beach', 'Indoors', 'Outdoors', 'City', 'Nightlife', 'Hiking'];

  // Handle form submission
  const handleSubmit = (e) => {
    e.preventDefault();
    if (selectedTags.length === 3) {
      console.log({ selectedTags });
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
      width: '1000px',
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
    subheading: {
      color: '#002f6c', 
      marginTop: '0px', 
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
      borderRadius: '20px',
      border: isSelected ? '2px solid #00509e' : '1px solid #ccc',
      backgroundColor: isSelected ? '#e0f0ff' : '#f0f0f0',
      cursor: 'pointer',
      transition: 'all 0.3s ease',
      boxShadow: isSelected ? '0 2px 4px rgba(0, 80, 158, 0.3)' : 'none',
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
          <h6 style={styles.subheading}>What tags are important to you on your travels?</h6>
          
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
