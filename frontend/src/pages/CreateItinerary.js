import React, { useState, useEffect } from 'react';
import navBarLogo from '../assets/logo-long-transparent.png';
import Tags from '../config/tags.json';
import '../styles/styles.css'; // Import the CSS file

function CreateItinerary() {
  const categories = Tags.categories;
  const [selectedTags, setSelectedTags] = useState([]);
  const [tagErrorMessage, setTagErrorMessage] = useState('');
  const [itineraryDetails, setItineraryDetails] = useState({
    name: '',
    location: '',
    description: '',
    estimatedCost: '',
    suggestedSeason: ''
  });

  const [events, setEvents] = useState([{ ampm: '', time: '', location: '', description: '', image: null }]);
  const [eventErrorMessage, setEventErrorMessage] = useState('');

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

    const generateTimeOptions = () => {
      const times = [];
      for (let hour = 1; hour <= 12; hour++) {
        times.push(`${hour}:00`);
        times.push(`${hour}:30`);
      }
      return times;
    };
    
  return (
    <>
      <nav className="navBar">
        <img src={navBarLogo} alt="Trip Tailor Logo" className="navBarLogo" />
        <button className="profileButton">
          <i className="fas fa-bars icon"></i>
          <i className="fa-regular fa-user icon"></i>
        </button>
      </nav>
      <div className="centeredContainer">
        <div className="leftBox">
          <h2 className="heading">Create Itinerary</h2>
          <form onSubmit={handleSubmit} className="form">
            <h2 className="headingCI">Basic Info</h2>
              <div className="inputGroup">
                <label htmlFor="name" className="subheading">Name</label>
                <input
                  type="text"
                  id="name"
                  name="name"
                  value={itineraryDetails.name}
                  onChange={handleInputChange}
                  className="input"
                  placeholder="Enter itinerary name"
                />
              </div>
            <div className="inputGroup">
              <label htmlFor="location" className="subheading">Location</label>
              <input
                type="text"
                id="location"
                name="location"
                value={itineraryDetails.location}
                onChange={handleInputChange}
                className="input"
                placeholder="Enter location"
              />
            </div>
            <div className="inputGroup">
              <label htmlFor="description" className="subheading">Description</label>
              <textarea
                id="description"
                name="description"
                value={itineraryDetails.description}
                onChange={handleInputChange}
                className="input"
                placeholder="Enter a brief description"
              />
            </div>
            <div className="inputGroup">
            <label htmlFor="estimatedCost" className="subheading">Estimated Cost</label>
              <input
                type="text"
                id="estimatedCost"
                name="estimatedCost"
                value={itineraryDetails.estimatedCost}
                onChange={handleInputChange}
                className="input"
                placeholder="Enter estimated cost"
              />
            </div>
            <div style={{ height: '15px' }}></div>
            <div className="inputGroup">
              <h2 className="headingCI">What tags apply to your itinerary?</h2>
              {Object.keys(categories).map((category) => (
                <div key={category} className="inputGroup">
                  <label className="subheading">
                    {category.replace('_', ' ').toLowerCase().replace(/^\w/, (c) => c.toUpperCase())}
                  </label>
                  <div className="tagsLeft">
                    {categories[category].map((tag) => (
                      <div
                        key={tag}
                        onClick={() => handleTagSelection(tag)}
                        className={`tag ${selectedTags.includes(tag) ? 'selected' : ''}`}
                      >
                        {tag}
                      </div>
                    ))}
                  </div>
                </div>
              ))}
              {tagErrorMessage && <div className="errorMessage">{tagErrorMessage}</div>}
            </div>
            <h2 className="headingCI">What events are part of your itinerary?</h2>

            <div className="events">
              {events.map((event, index) => (
                <div key={index} className="eventBox">
                  <div className="inputGroup">
                    <div className="timeLocationContainer" style={{ display: 'flex', gap: '20px' }}>
                      <div style={{ flex: 1 }}>
                        <label htmlFor={`time-${index}`} className="subheadingLeft">Time</label>
                        <div className="timeSelect" style={{ display: 'flex', gap: '10px' }}>
                          <select
                            id={`time-${index}`}
                            name="time"
                            value={event.time}
                            onChange={(e) => handleEventChange(index, e)}
                            className="input"
                          >
                            {generateTimeOptions().map((time) => (
                              <option key={time} value={time}>{time}</option>
                            ))}
                          </select>
                          <select
                            id={`ampm-${index}`}
                            name="ampm"
                            value={event.ampm}
                            onChange={(e) => handleEventChange(index, e)}
                            className="input"
                          >
                            <option value="AM">AM</option>
                            <option value="PM">PM</option>
                          </select>
                        </div>
                      </div>
                      <div style={{ flex: 1 }}>
                        <label htmlFor={`location-${index}`} className="subheadingLeft">Location</label>
                        <input
                          type="text"
                          id={`location-${index}`}
                          name="location"
                          value={event.location}
                          onChange={(e) => handleEventChange(index, e)}
                          className="input"
                          placeholder="Enter event location"
                        />
                      </div>
                    </div>
                  </div>
                  <div className="inputGroup">
                    <label htmlFor={`description-${index}`} className="subheadingLeft">Description (100 character max.)</label>
                    <textarea
                      id={`description-${index}`}
                      name="description"
                      value={event.description}
                      onChange={(e) => handleEventChange(index, e)}
                      className="input"
                      placeholder="Enter event description"
                      maxLength="100"
                    />
                  </div>
                </div>
              ))}
              {eventErrorMessage && <div className="errorMessage">{eventErrorMessage}</div>}
              <button type="button" className="continueButton" onClick={addEvent}>
                <span>+</span> Add another event
              </button>
            </div>
          </form>
        </div>
      </div>
    </>
  );
}

export default CreateItinerary;
