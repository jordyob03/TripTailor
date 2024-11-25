import React, { useState, useEffect } from "react";
import { useLocation } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faMapMarkerAlt } from "@fortawesome/free-solid-svg-icons";
import iconMap from '../config/iconMap';

function ItineraryDetails() {
  const location = useLocation();
  const itinId = parseInt(location.pathname.split("/").pop(), 10);

  const [itinerary, setItinerary] = useState(null);
  const [errorMessage, setErrorMessage] = useState("");

  const fetchItineraryData = async () => {
    try {
      const dummyData = {
        itineraryId: 6,
        city: "Lahore",
        country: "Pakistan",
        title: "Exploring Lahore's Old City",
        description:
          "My friends and I spent an unforgettable day exploring the rich history and vibrant culture of Lahore's Old City. We wandered through bustling bazaars, admired Mughal architecture, and finished the day with a traditional dinner and a stroll around the iconic Badshahi Mosque.",
        price: 0.0,
        languages: ["Urdu", "Punjabi"],
        tags: [
          "Young Adults",
          "Shopping",
          "City",
          "Beach",
          "Nightlife",
          "Short Walks",
          "Short Walks",
          "Short Walks",
          "Short Walks",
          "Short Walks",
          "Short Walks",
        ],
        events: ["19", "20", "21", "22"],
        postId: 6,
        username: "herobrine",
      };

      setItinerary(dummyData);
    } catch (error) {
      console.error("Error fetching dummy itinerary data:", error);
      setErrorMessage("Failed to fetch itinerary data.");
    }
  };

  useEffect(() => {
    if (!itinId) {
      setErrorMessage("Invalid itinerary ID.");
      return;
    }
    fetchItineraryData();
  }, [itinId]);

  if (errorMessage) {
    return <div className="error">{errorMessage}</div>;
  }

  if (!itinerary) {
    return <div className="loading">Loading...</div>;
  }

  return (
    <div className="itinPageContainer">
      {itinerary && (
        <>
          <div className="itineraryHeader">
            <div className="itineraryText">
              <div className="itineraryDesc">
                <h1>{itinerary.title}</h1>
                <h4>
                  <FontAwesomeIcon
                    icon={faMapMarkerAlt}
                    style={{ color: "#00509e", marginRight: "8px" }}
                  />
                  {itinerary.city}, {itinerary.country}
                </h4>
                <p>{itinerary.description}</p>
                <p className="postedBy">
                  Posted by{" "}
                  <span className="username">@{itinerary.username}</span> on{" "}
                  <span className="date">21st September 2024</span>
                </p>
              </div>
              <div className="eventImages">
                <img
                  className="eventImage"
                  src={`https://via.placeholder.com/400x300?text=Event+0`}
                  alt="Event 0"
                />
                <img
                  className="eventImage"
                  src={`https://via.placeholder.com/400x300?text=Event+1`}
                  alt="Event 1"
                />
                <img
                  className="eventImage"
                  src={`https://via.placeholder.com/400x300?text=Event+2`}
                  alt="Event 2"
                />
              </div>
            </div>
            <img
              src="https://via.placeholder.com/600x600"
              alt="Main"
              className="itineraryMainImage"
            />
          </div>

          <div className="itineraryTags">
            {itinerary.tags.map((tag, index) => (
              <div key={index} className="tagItem">
                <div className="tagIcon1">
                  {iconMap[tag] && <FontAwesomeIcon icon={iconMap[tag]} />}
                </div>
                <span>{tag}</span>
              </div>
            ))}
          </div>

          <div className="itineraryDetails">
            <div className="itineraryOverview">
              <h4>Itinerary Overview</h4>
              <p>
                <strong>Location:</strong> {itinerary.city}, {itinerary.country}
              </p>
              <p>
                <strong>Suggested Travel Party:</strong> Young Adults
              </p>
              <p>
                <strong>Suggested Travel Season:</strong> Spring, Fall
              </p>
              <p>
                <strong>Estimated Cost:</strong> $$$
              </p>
            </div>

            <div className="itineraryTable">
              <h4>Itinerary Events</h4>
              <table>
                <thead>
                  <tr>
                    <th>Time</th>
                    <th>Location</th>
                    <th>Notes</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td>9:00am</td>
                    <td>Shalimar Gardens</td>
                    <td>Peaceful garden, perfect for morning stroll.</td>
                  </tr>
                  <tr>
                    <td>11:00am</td>
                    <td>Food Street</td>
                    <td>Tiny alleyways with traditional street food.</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </>
      )}
    </div>
  );
}

export default ItineraryDetails;
