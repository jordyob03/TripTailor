import React, { useState, useRef } from 'react';
import navBarLogo from '../assets/logo-long-transparent.png';
import '../styles/styles.css'; 
import Tags from '../config/tags.json';

function HomeSearch() {
  const allTags = Object.values(Tags.categories).flat();
  const [selectedTags, setSelectedTags] = useState([]); 
  const [country, setCountry] = useState('');
  const [city, setCity] = useState('');
  const [tagErrorMessage, setTagErrorMessage] = useState('');
  const [countryErrorMessage, setCountryErrorMessage] = useState('');
  const [cityErrorMessage, setCityErrorMessage] = useState('');
  const [menuOpen, setMenuOpen] = useState(false); 
  const [searchResults, setSearchResults] = useState([]); 
  const tagContainerRef = useRef(null);

  const handleTagClick = (tag) => {
    // Toggle the selection of the tag
    if (selectedTags.includes(tag)) {
      setSelectedTags(selectedTags.filter((t) => t !== tag));
    } else {
      setSelectedTags([...selectedTags, tag]);
    }
    setTagErrorMessage('');
  };

  const handleSearch = () => {
    if (selectedTags.length === 0) {
      setTagErrorMessage('Please select at least one tag.');
    }
    if (!country) {
      setCountryErrorMessage('Please enter a country.');
    }
    if (!city) {
      setCityErrorMessage('Please enter a city.');
    }

    // Mock search results for demonstration
    if (selectedTags.length > 0 && country && city) {
      setSearchResults([
        {
          location: `${city}, ${country}`,
          title: 'Hiking and Hot Springs',
          description: 'Immerse yourself in the beauty of hot springs surrounded by volcanic views.',
          tags: ['Hiking', 'Hot Springs', 'Outdoors'],
          image: 'https://via.placeholder.com/300x180', // Placeholder image
        },
        {
          location: `${city}, ${country}`,
          title: 'Beautiful Evening at Bob Kerrey Bridge',
          description: 'Enjoy a stunning evening stroll across the Bob Kerrey Bridge.',
          tags: ['Outdoors', 'Sightseeing'],
          image: 'https://via.placeholder.com/300x180', // Placeholder image
        },
      ]);
    }
  };

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
    <div>
      {/* Navbar */}
      <nav className="navBar">
        <img src={navBarLogo} alt="Trip Tailor Logo" className="navBarLogo" />

        {/* Profile Button */}
        <button className="profileButton" onClick={() => setMenuOpen(!menuOpen)}>
          <i className="fas fa-bars" style={{ fontSize: '16px', color: '#00509e', marginRight: '15px' }}></i>
          <i className="fa-regular fa-user" style={{ fontSize: '24px', color: '#00509e' }}></i>
        </button>

        {/* Dropdown Menu */}
        {menuOpen && (
          <div style={{
            position: 'absolute',
            top: '50px',
            right: '0',
            backgroundColor: 'white',
            boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)',
            borderRadius: '8px',
            width: '150px',
            zIndex: 3
          }}>
            <ul style={{
              listStyle: 'none',
              padding: '10px 0',
              margin: '0',
              textAlign: 'left'
            }}>
              <li style={{ padding: '10px 20px', cursor: 'pointer' }}>Profile</li>
              <li style={{ padding: '10px 20px', cursor: 'pointer' }}>Account Settings</li>
              <li style={{ padding: '10px 20px', cursor: 'pointer' }}>My Itineraries</li>
              <li style={{ padding: '10px 20px', cursor: 'pointer' }}>Home</li>
            </ul>
          </div>
        )}
      </nav>

      {/* Tag Filters */}
      <div style={{
        position: 'fixed',
        top: '100px',
        left: '0',
        right: '0',
        backgroundColor: 'white',
        padding: '10px 0',
        zIndex: 1,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-evenly',
        boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)',
      }}>
        <div style={{ display: 'flex', alignItems: 'center' }}>
          <button onClick={scrollTagsLeft} style={{ marginRight: '10px' }}>{'<'}</button>
          <div ref={tagContainerRef} style={{
            display: 'flex',
            overflowX: 'auto',
            maxWidth: '900px',
            whiteSpace: 'nowrap',
          }}>
            {allTags.map((tag) => (
              <div
                key={tag}
                style={{
                  padding: '6px 12px',
                  borderRadius: '50px',
                  border: selectedTags.includes(tag) ? '2px solid #00509e' : '1px solid #ccc',
                  backgroundColor: selectedTags.includes(tag) ? '#00509e' : '#C6DFF0',
                  color: selectedTags.includes(tag) ? 'white' : 'black',
                  cursor: 'pointer',
                  transition: 'all 0.3s ease',
                  margin: '0 5px',
                  fontSize: '14px',
                  display: 'inline-block'
                }}
                onClick={() => handleTagClick(tag)}
              >
                {tag}
              </div>
            ))}
          </div>
          <button onClick={scrollTagsRight} style={{ marginLeft: '10px' }}>{'>'}</button>
        </div>

        {/* Country and City Filters */}
        <div style={{ display: 'flex', alignItems: 'center' }}>
          <input
            type="text"
            placeholder="Enter Country"
            value={country}
            onChange={(e) => setCountry(e.target.value)}
            style={{
              padding: '6px 12px',
              borderRadius: '4px',
              border: '1px solid #ccc',
              marginRight: '10px'
            }}
          />
          <input
            type="text"
            placeholder="Enter City"
            value={city}
            onChange={(e) => setCity(e.target.value)}
            style={{
              padding: '6px 12px',
              borderRadius: '4px',
              border: '1px solid #ccc'
            }}
          />
        </div>

        {/* Search Button */}
        <div style={{ marginLeft: '10px' }}>
          <button
            onClick={handleSearch}
            style={{
              padding: '8px 16px',
              backgroundColor: '#00509e',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer'
            }}
          >
            Search
          </button>
        </div>
      </div>

      {/* Error Messages */}
      <div style={{ marginTop: '170px', textAlign: 'center' }}>
        {tagErrorMessage && <div style={{ color: 'red' }}>{tagErrorMessage}</div>}
        {countryErrorMessage && <div style={{ color: 'red' }}>{countryErrorMessage}</div>}
        {cityErrorMessage && <div style={{ color: 'red' }}>{cityErrorMessage}</div>}
      </div>

      {/* Search Results */}
      <div style={{ marginTop: '20px', padding: '20px' }}>
        {searchResults.length > 0 ? (
          <div style={{
            display: 'grid',
            gridTemplateColumns: 'repeat(auto-fill, minmax(300px, 1fr))',
            gap: '20px',
          }}>
            {searchResults.map((itinerary, index) => (
              <div key={index} style={{
                backgroundColor: 'white',
                borderRadius: '12px',
                boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)',
                overflow: 'hidden',
                position: 'relative',
              }}>
                <input type="checkbox" style={{
                  position: 'absolute',
                  top: '10px',
                  right: '10px',
                  cursor: 'pointer',
                }} />
                {itinerary.image && (
                  <img src={itinerary.image} alt={itinerary.title} style={{
                    width: '100%',
                    height: '180px',
                    objectFit: 'cover',
                  }} />
                )}
                <div style={{ padding: '15px', textAlign: 'left' }}>
                  <h4 style={{ color: '#00509e', fontSize: '18px', marginBottom: '5px' }}>
                    {itinerary.location}
                  </h4>
                  <h3 style={{ color: '#333', fontSize: '16px', fontWeight: 'bold', marginBottom: '10px' }}>
                    {itinerary.title}
                  </h3>
                  <p style={{ fontSize: '14px', color: '#666', marginBottom: '15px' }}>
                    {itinerary.description}
                  </p>
                  <div style={{ display: 'flex', flexWrap: 'wrap', gap: '5px', marginBottom: '10px' }}>
                    {itinerary.tags.map((tag, i) => (
                      <span key={i} style={{
                        padding: '5px 10px',
                        backgroundColor: '#E1EFFF',
                        borderRadius: '50px',
                        fontSize: '12px',
                        color: '#00509e',
                      }}>
                        {tag}
                      </span>
                    ))}
                  </div>
                  <button style={{
                    padding: '10px 20px',
                    backgroundColor: '#00509e',
                    color: 'white',
                    border: 'none',
                    borderRadius: '4px',
                    cursor: 'pointer',
                    width: '100%',
                    fontSize: '14px',
                    fontWeight: 'bold',
                  }}>
                    View Details
                  </button>
                </div>
              </div>
            ))}
          </div>
        ) : (
          <p style={{ textAlign: 'center', color: '#666' }}>No search results found.</p>
        )}
      </div>
    </div>
  );
}

export default HomeSearch;
