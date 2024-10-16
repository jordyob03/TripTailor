import React, { useState, useEffect } from 'react';
import navBarLogo from '../assets/logo-long-transparent.png';
import Tags from '../config/tags.json';

function CreateItinerary() {
  const allTags = Object.values(Tags.categories).flat();
  const [selectedTags, setSelectedTags] = useState([]);
  const [tagErrorMessage, setTagErrorMessage] = useState('');
  const [itineraryDetails, setItineraryDetails] = useState({
    name: '',
    location: '',
    description: '',
    suggestedParties: '',
    estimatedCost: '',
    suggestedSeason: ''
  });

  const [events, setEvents] = useState([{ ampm: '', time: '', location: '', description: '', image: null }]);
  const [eventErrorMessage, setEventErrorMessage] = useState('');
  const [shuffledTags, setShuffledTags] = useState([]);

  
  const shuffleArray = (array) => {
    return [...array].sort(() => Math.random() - 0.5);
  };

  useEffect(() => {
    const shuffled = shuffleArray(allTags);
    setShuffledTags(shuffled);
  }, []);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setItineraryDetails({ ...itineraryDetails, [name]: value });
  };

  const handleEventChange = (index, e) => {
    const { name, value } = e.target;
    const updatedEvents = [...events];
    updatedEvents[index][name] = value;
    setEvents(updatedEvents);
  };

  const handleFileChange = (index, e) => {
    const file = e.target.files[0];
    const updatedEvents = [...events];
    updatedEvents[index].image = file;
    setEvents(updatedEvents);
  };

  const addEvent = () => {
    if (events.length >= 24) {
      setEventErrorMessage('Cannot add more than 24 events in a 24-hour period.');
    } else {
      setEvents([...events, { ampm: '', time: '', location: '', description: '', image: null }]);
      setEventErrorMessage('');
    }
  };

  const handleTagSelection = (tag) => {
    if (selectedTags.includes(tag)) {
      setSelectedTags(selectedTags.filter((t) => t !== tag));
    } else {
      setSelectedTags([...selectedTags, tag]);
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    
    if (events.length < 1) {
      setEventErrorMessage('At least one event is required.');
      return;
    }

    
    if (selectedTags.length >= 3) {
      setTagErrorMessage('');
      setEventErrorMessage('');
      console.log({ itineraryDetails, selectedTags, events });
    } else {
      setTagErrorMessage('Please select at least 3 tags.');
    }
  };

  const timeOptions = [];
  for (let h = 1; h <= 12; h++) {
    ['00', '15', '30', '45'].forEach(minute => {
      const hour = h.toString().padStart(2, '0');
      timeOptions.push(`${hour}:${minute}`);
    });
  }

  return (
    <>
      {/* navBar */}
      <nav style={{
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        padding: '1vw 1vw',
        height: '50px',
        width: '100vw',
        backgroundColor: 'white',
        boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)',
        position: 'fixed',
        top: '0',
        left: '0',
        zIndex: 2
      }}>
        <img src={navBarLogo} alt="Trip Tailor Logo" style={{ width: '180px', marginLeft: '80px', marginTop: '5px' }} />
        <button style={{
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
          transition: 'box-shadow 0.3s ease',
        }}>
          <i className="fas fa-bars" style={{ fontSize: '16px', color: '#00509e', marginRight: '15px' }}></i>
          <i className="fa-regular fa-user" style={{ fontSize: '24px', color: '#00509e' }}></i>
        </button>
      </nav>

      {}
      <div style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        minHeight: 'calc(100vh - 60px)',
        padding: '20px',
        marginTop: '60px',
        backgroundColor: '#f0f4f8',  
      }}>
        {}
        <div style={{
          backgroundColor: 'white',
          borderRadius: '20px',
          padding: '40px',
          boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)',
          width: '100%',
          maxWidth: '800px',
          transition: 'box-shadow 0.3s ease',
        }}>
          <h2 style={{ color: '#00509e', marginBottom: '20px', fontSize: '24px', textAlign: 'center' }}>Create Itinerary</h2>
          <form onSubmit={handleSubmit}>
            {/* Itinerary Details */}
            <div style={{ marginBottom: '20px' }}>
              <label style={{ display: 'block', marginBottom: '8px', fontWeight: '500', color: '#002f6c', textAlign: 'left' }}>Name</label>
              <input type="text" name="name" value={itineraryDetails.name} onChange={handleInputChange} style={{
                width: '100%',
                padding: '10px',
                borderRadius: '4px',
                border: '1px solid #ddd',
                marginBottom: '15px',
                transition: 'border-color 0.3s ease',
              }} onFocus={(e) => e.target.style.borderColor = '#00509e'} onBlur={(e) => e.target.style.borderColor = '#ddd'} />

              <label style={{ display: 'block', marginBottom: '8px', fontWeight: '500', color: '#002f6c', textAlign: 'left' }}>Location</label>
              <input type="text" name="location" value={itineraryDetails.location} onChange={handleInputChange} style={{
                width: '100%',
                padding: '10px',
                borderRadius: '4px',
                border: '1px solid #ddd',
                marginBottom: '15px',
                transition: 'border-color 0.3s ease',
              }} onFocus={(e) => e.target.style.borderColor = '#00509e'} onBlur={(e) => e.target.style.borderColor = '#ddd'} />

              <label style={{ display: 'block', marginBottom: '8px', fontWeight: '500', color: '#002f6c', textAlign: 'left' }}>Description</label>
              <textarea name="description" value={itineraryDetails.description} onChange={handleInputChange} style={{
                width: '100%',
                padding: '10px',
                borderRadius: '4px',
                border: '1px solid #ddd',
                marginBottom: '15px',
                transition: 'border-color 0.3s ease',
              }} onFocus={(e) => e.target.style.borderColor = '#00509e'} onBlur={(e) => e.target.style.borderColor = '#ddd'} />

              <label style={{ display: 'block', marginBottom: '8px', fontWeight: '500', color: '#002f6c', textAlign: 'left' }}>Suggested Traveling Parties</label>
              <input type="text" name="suggestedParties" value={itineraryDetails.suggestedParties} onChange={handleInputChange} style={{
                width: '100%',
                padding: '10px',
                borderRadius: '4px',
                border: '1px solid #ddd',
                marginBottom: '15px',
                transition: 'border-color 0.3s ease',
              }} onFocus={(e) => e.target.style.borderColor = '#00509e'} onBlur={(e) => e.target.style.borderColor = '#ddd'} />

              <label style={{ display: 'block', marginBottom: '8px', fontWeight: '500', color: '#002f6c', textAlign: 'left' }}>Estimated Cost</label>
              <input type="text" name="estimatedCost" value={itineraryDetails.estimatedCost} onChange={handleInputChange} style={{
                width: '100%',
                padding: '10px',
                borderRadius: '4px',
                border: '1px solid #ddd',
                marginBottom: '15px',
                transition: 'border-color 0.3s ease',
              }} onFocus={(e) => e.target.style.borderColor = '#00509e'} onBlur={(e) => e.target.style.borderColor = '#ddd'} />

              <label style={{ display: 'block', marginBottom: '8px', fontWeight: '500', color: '#002f6c', textAlign: 'left' }}>Suggested Season</label>
              <input type="text" name="suggestedSeason" value={itineraryDetails.suggestedSeason} onChange={handleInputChange} style={{
                width: '100%',
                padding: '10px',
                borderRadius: '4px',
                border: '1px solid #ddd',
                marginBottom: '15px',
                transition: 'border-color 0.3s ease',
              }} onFocus={(e) => e.target.style.borderColor = '#00509e'} onBlur={(e) => e.target.style.borderColor = '#ddd'} />
            </div>

            {/* Tags */}
            <div style={{ marginBottom: '20px' }}>
              <label style={{ display: 'block', marginBottom: '8px', fontWeight: '500', color: '#002f6c', textAlign: 'left' }}>Tags</label>
              <div style={{ display: 'flex', flexWrap: 'wrap', gap: '10px', justifyContent: 'center' }}>
                {shuffledTags.map((tag) => (
                  <div
                    key={tag}
                    onClick={() => handleTagSelection(tag)}
                    style={{
                      padding: '8px 16px',
                      borderRadius: '50px',
                      border: selectedTags.includes(tag) ? '2px solid #00509e' : '1px solid #ccc',
                      backgroundColor: selectedTags.includes(tag) ? '#00509e' : '#C6DFF0',
                      color: selectedTags.includes(tag) ? 'white' : 'black',
                      cursor: 'pointer',
                      transition: 'all 0.3s ease',
                      fontSize: '14px',
                      boxShadow: selectedTags.includes(tag) ? '0 4px 8px rgba(0, 80, 158, 0.3)' : 'none',
                    }}
                  >
                    {tag}
                  </div>
                ))}
              </div>
              {tagErrorMessage && <div style={{ color: 'red', fontSize: '14px', marginTop: '10px' }}>{tagErrorMessage}</div>}
            </div>

            {/* Events Box */}
            <div style={{
              backgroundColor: '#f9f9f9',
              borderRadius: '12px',
              padding: '20px',
              boxShadow: '0 2px 6px rgba(0, 0, 0, 0.1)',
              marginBottom: '20px'
            }}>
              <h2 style={{ color: '#00509e', marginBottom: '15px', fontSize: '24px', textAlign: 'center' }}>Events</h2>
              {events.map((event, index) => (
                <div key={index} style={{ marginBottom: '15px' }}>
                  <label style={{ display: 'block', marginBottom: '8px', fontWeight: '500', color: '#002f6c', textAlign: 'left' }}>Time</label>
                  <div style={{ display: 'flex', gap: '10px', marginBottom: '10px' }}>
                    <select
                      name="ampm"
                      value={event.ampm}
                      onChange={(e) => handleEventChange(index, e)}
                      style={{
                        padding: '10px',
                        borderRadius: '4px',
                        border: '1px solid #ddd',
                        width: '80px'
                      }}
                    >
                      <option value="AM">AM</option>
                      <option value="PM">PM</option>
                    </select>

                    <select
                      name="time"
                      value={event.time}
                      onChange={(e) => handleEventChange(index, e)}
                      style={{
                        padding: '10px',
                        borderRadius: '4px',
                        border: '1px solid #ddd',
                        width: '150px'
                      }}
                    >
                      {timeOptions.map((option) => (
                        <option key={option} value={option}>{option}</option>
                      ))}
                    </select>
                  </div>

                  <label style={{ display: 'block', marginBottom: '8px', fontWeight: '500', color: '#002f6c', textAlign: 'left' }}>Location</label>
                  <input type="text" name="location" value={event.location} onChange={(e) => handleEventChange(index, e)} style={{
                    width: '100%',
                    padding: '10px',
                    borderRadius: '4px',
                    border: '1px solid #ddd',
                    marginBottom: '10px',
                    transition: 'border-color 0.3s ease',
                  }} onFocus={(e) => e.target.style.borderColor = '#00509e'} onBlur={(e) => e.target.style.borderColor = '#ddd'} />

                  <label style={{ display: 'block', marginBottom: '8px', fontWeight: '500', color: '#002f6c', textAlign: 'left' }}>Description</label>
                  <textarea name="description" value={event.description} onChange={(e) => handleEventChange(index, e)} style={{
                    width: '100%',
                    padding: '10px',
                    borderRadius: '4px',
                    border: '1px solid #ddd',
                    marginBottom: '10px',
                    transition: 'border-color 0.3s ease',
                  }} onFocus={(e) => e.target.style.borderColor = '#00509e'} onBlur={(e) => e.target.style.borderColor = '#ddd'} />

                  <label style={{ display: 'block', marginBottom: '8px', fontWeight: '500', color: '#002f6c', textAlign: 'left' }}>Images</label>
                  <input type="file" onChange={(e) => handleFileChange(index, e)} style={{
                    width: '100%',
                    padding: '10px',
                    borderRadius: '4px',
                    marginTop: '10px'
                  }} />
                </div>
              ))}
              {eventErrorMessage && <div style={{ color: 'red', fontSize: '14px', marginTop: '10px' }}>{eventErrorMessage}</div>}
              <button type="button" style={{
                padding: '10px 20px',
                backgroundColor: '#00509e',
                color: 'white',
                border: 'none',
                borderRadius: '4px',
                cursor: 'pointer',
                marginTop: '10px',
                transition: 'background-color 0.3s ease',
              }} onClick={addEvent}>Add Event</button>
            </div>

            {}
            <div style={{ textAlign: 'right' }}> {}
              <button type="submit" style={{
                padding: '12px',
                backgroundColor: '#00509e',
                color: 'white',
                border: 'none',
                borderRadius: '4px',
                fontSize: '16px',
                cursor: 'pointer',
                transition: 'background-color 0.3s ease'
              }}>Create Itinerary</button>
            </div>
          </form>
        </div>
      </div>
    </>
  );
}

export default CreateItinerary;
