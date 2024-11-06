import React, { useState} from 'react';
import Tags from '../config/tags.json';
import '../styles/styles.css'; 
import { useNavigate } from 'react-router';
import itineraryAPI from '../api/itineraryAPI.js';
import { convertBytesToImage, convertFileToBase64 } from '../utils/imageHandler.js';

function CreateItinerary() {
  const categories = Tags.categories;
  const [selectedTags, setSelectedTags] = useState([]);
  const [tagErrorMessage, setTagErrorMessage] = useState('');
  const [eventErrorMessage, setEventErrorMessage] = useState('');
  const [basicErrorMessage, setBasicErrorMessage] = useState('');
  const [itineraryImages, setItineraryImages] = useState([]);
  const [itineraryDetails, setItineraryDetails] = useState({
    name: '',
    city: '',
    country: '',
    description: ''
  });

  const navigate = useNavigate()

  const [events, setEvents] = useState([{ name: '', startTime: '1:00 AM', endTime: '2:00 AM', location: '', description: '', cost: '' }]);

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

  const handleMultipleFileChange = async (e) => {
    const files = Array.from(e.target.files);

    const convertedImages = await Promise.all(files.map(async (file) => {
        const previewUrl = URL.createObjectURL(file);
        const base64String = await convertFileToBase64(file); 
        return { file, previewUrl, base64String };
    }));

    setItineraryImages([...itineraryImages, ...convertedImages]);
};



  const removeImage = (index) => {
    const updatedImages = itineraryImages.filter((_, i) => i !== index);
    setItineraryImages(updatedImages);
  };

  const addEvent = () => {

    if (events.length >= 24) {
      setEventErrorMessage('Cannot add more than 24 events in a 24-hour period.');
    } else {
      setEvents([...events, { name: '', startTime: '1:00 AM', endTime: '2:00 AM', location: '', description: '', cost: '' }]);
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

  const handleSubmit = async (e) => {    e.preventDefault();
  
    // Filter out events that have no description and no location
    const filteredEvents = events.filter(
      (event) => event.description.trim() !== '' || event.location.trim() !== ''
    );

  
    // Basic info check
    const { name, city, country, description} = itineraryDetails;
  
    if (name && city && country && description) {
      setBasicErrorMessage('');
    } else {
      setBasicErrorMessage('Please fill out all basic info fields.');
    }
  
    // Check that at least 3 tags have been selected
    if (selectedTags.length >= 3) {
      setTagErrorMessage('');
    } else {
      setTagErrorMessage('Please select at least 3 tags.');
    }
  
    // Events checks
    // Check that there is at least one event with complete details
    const hasValidEvent = filteredEvents.length > 0;

    if (!hasValidEvent) {
      setEventErrorMessage('At least one complete event is required.');
      console.log({ itineraryDetails, itineraryImages, selectedTags, events: filteredEvents, hasValidEvent });
    }
  
    // Check if there is any event with an incomplete state
    const hasIncompleteEvent = filteredEvents.some((event) => {
      const fields = [
        event.name.trim(),
        event.startTime,
        event.endTime,
        event.location.trim(),
        event.description.trim(),
        event.cost
      ];

      if (name && city && country && description && selectedTags.length >= 3 && hasValidEvent)
      {navigate('/my-travels');}
    });

  
    if (hasIncompleteEvent) {
      setEventErrorMessage('Please complete all fields for incomplete events or delete them.');
    }

    if (!hasIncompleteEvent && hasValidEvent) {
      setEventErrorMessage('');
    }
    

    const Data = {
      Name: itineraryDetails.name,
      Username: localStorage.getItem('username'),
      City: itineraryDetails.city,
      Country: itineraryDetails.country,
      Description: itineraryDetails.description,
      Tags: selectedTags,
      Events: filteredEvents,
      Images: itineraryImages.map(img => img.base64String)
    }

    try {
      const response = await itineraryAPI.post('/itin-creation', Data);
      console.log('Location created:', response.data);

    } catch (error) {

    }

    return;
  };
  
  const generateTimeOptions = () => {
    const times = [];
    const periods = ['AM', 'PM'];
    for (let period of periods) {
      for (let hour = 1; hour <= 12; hour++) {
        for (let minute of ['00', '15', '30', '45']) {
          times.push(`${hour}:${minute} ${period}`);
        }
      }
    }
    return times;
  };

    const removeEvent = (index) => {
      const updatedEvents = events.filter((_, i) => i !== index);
      setEvents(updatedEvents);
    };
    
  return (
    <>
      <div className="centeredContainer">
        <div className="leftBox">
          <h2 className="heading">Create Itinerary</h2>
          <form onSubmit={handleSubmit} className="form">
            <h2 className="headingCI">Basic Info</h2>
              <div className="inputGroup">
                <label htmlFor="name" className="subheadingLeft">Name</label>
                <input
                  type="text"
                  id="name"
                  name="name"
                  value={itineraryDetails.name}
                  onChange={handleInputChange}
                  className="input"
                  placeholder='Like "A Day in Rome" or "Sightseeing Road Trip" '
                />
              </div>
            <div className="inputGroup">
              <label htmlFor="city" className="subheadingLeft">City</label>
              <input
                type="text"
                id="city"
                name="city"
                value={itineraryDetails.city}
                onChange={handleInputChange}
                className="input"
                placeholder="Enter city"
              />
            </div>
            <div className="inputGroup">
              <label htmlFor="country" className="subheadingLeft">Country</label>
              <input
                type="text"
                id="country"
                name="country"
                value={itineraryDetails.country}
                onChange={handleInputChange}
                className="input"
                placeholder="Enter country"
              />
            </div>
            <div className="inputGroup">
              <label htmlFor="description" className="subheadingLeft">Description</label>
              <textarea
                id="description"
                name="description"
                value={itineraryDetails.description}
                onChange={handleInputChange}
                className="input"
                placeholder="Enter a brief description of your itinerary"
                maxLength="100"
              />
            </div>
            {basicErrorMessage && <div className="errorMessage">{basicErrorMessage}</div>}
            <div className="inputGroup">
              <label htmlFor="itineraryImages" className="subheadingLeft">Upload Itinerary Images</label>
              <div className="imageUploadContainer">
                <div className="imagePreview">
                  {itineraryImages.map((image, index) => (
                    <div key={index} className="imageContainer">
                        <img src={image.previewUrl} alt={`Itinerary Preview ${index + 1}`} className="previewImage" />
                        <button className="removeButton" onClick={() => removeImage(index)}>x</button>
                    </div>
                  ))}
                  <div className="addImageButton">
                    <input
                      type="file"
                      id="itineraryImages"
                      name="itineraryImages"
                      accept="image/*"
                      multiple
                      onChange={handleMultipleFileChange}
                      className="inputFile"
                    />
                    <label htmlFor="itineraryImages">
                      <i className="fa-regular fa-square-plus fa-3x"></i>
                    </label>
                  </div>
                </div>
              </div>
            </div>
            <div style={{ height: '15px' }}></div>
            <div className="inputGroup">
              <h2 className="headingCI">What tags apply to your itinerary?</h2>
              {Object.keys(categories).map((category) => (
                <div key={category} className="inputGroup">
                  <label className="subheadingLeft">
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
                  {/* Top Row: Event Name and Location */}
                  <div className="inputGroup" style={{ display: 'flex', gap: '20px' }}>
                    <div style={{ flex: 1 }}>
                      <label htmlFor={`name-${index}`} className="subheadingLeft">Event Name</label>
                      <input
                        type="text"
                        id={`name-${index}`}
                        name="name"
                        value={event.name}
                        onChange={(e) => handleEventChange(index, e)}
                        className="input"
                        placeholder="Enter event name"
                      />
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
                    <button className="removeButton" onClick={() => removeEvent(index)}>x</button>
                  </div>
                  
                  {/* Second Row: Start Time, End Time, and cost */}
                  <div className="inputGroup" style={{ display: 'flex', gap: '20px', marginTop: '10px' }}>
                    <div style={{ flex: 1 }}>
                      <label htmlFor={`startTime-${index}`} className="subheadingLeft">Start Time</label>
                      <select
                        id={`startTime-${index}`}
                        name="startTime"
                        value={event.startTime}
                        onChange={(e) => handleEventChange(index, e)}
                        className="input"
                      >
                        {generateTimeOptions().map((time) => (
                          <option key={time} value={time}>{time}</option>
                        ))}
                      </select>
                    </div>
                    <div style={{ flex: 1 }}>
                      <label htmlFor={`endTime-${index}`} className="subheadingLeft">End Time</label>
                      <select
                        id={`endTime-${index}`}
                        name="endTime"
                        value={event.endTime}
                        onChange={(e) => handleEventChange(index, e)}
                        className="input"
                      >
                        {generateTimeOptions().map((time) => (
                          <option key={time} value={time}>{time}</option>
                        ))}
                      </select>
                    </div>
                    <div style={{ flex: 1 }}>
                      <label htmlFor={`cost-${index}`} className="subheadingLeft">Cost</label>
                      <input
                        type="number"
                        id={`cost-${index}`}
                        name="cost"
                        value={event.cost}
                        onChange={(e) => handleEventChange(index, e)}
                        className="input"
                        placeholder="Enter event cost"
                        min="0"
                        step="0.01"
                      />
                    </div>
                  </div>

                  {/* Description Field (optional) */}
                  <div className="inputGroup" style={{ marginTop: '10px' }}>
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
                <span>+</span> Add Event
              </button>
            </div>

            <div style={{ height: '15px' }}></div>
            <button type="submit" className="continueButton">
              Submit Itinerary
            </button>
          </form>
        </div>
      </div>
    </>
  );
}

export default CreateItinerary;
