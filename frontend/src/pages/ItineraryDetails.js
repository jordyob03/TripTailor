import React, { useState, useEffect } from "react";
import { useLocation } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faMapMarkerAlt } from "@fortawesome/free-solid-svg-icons";
import iconMap from '../config/iconMap';
import tagsJson from "../config/tags.json";

function ItineraryDetails() {
  const location = useLocation();
  const itinId = parseInt(location.pathname.split("/").pop(), 10);
  const [events, setEvents] = useState([]);

  const [itinerary, setItinerary] = useState(null);
  const [errorMessage, setErrorMessage] = useState("");

  const categorizeTags = () => {
    const categorizedTags = {};
    Object.keys(tagsJson.categories).forEach((category) => {
      categorizedTags[category] = itinerary.tags.filter((tag) =>
        tagsJson.categories[category].includes(tag)
      );
    });
    return categorizedTags;
  };

  const categorizedTags = categorizeTags();
  const dummyItinerary = {
    itineraryId: 12,
    city: "Islamabad",
    country: "Pakistan",
    title: "Visiting the NUST Campus",
    description:
      "We spent the day exploring the beautiful NUST campus. We visited the library, the sports facilities, and the stunning surroundings. We finished the day with a picnic by the lake and a walk around the campus.",
    price: 0.0,
    languages: ["Urdu", "Punjabi", "Pashto"],
    tags: [
      "Solo",
      "City",
      "Short Walks",
      "Museum Visits",
      "Couples",
      "Historical Sites",
      "Leisure Activities",
      "Photography Tours",
      "Cultural Explorers",
      "Wildlife Watching",
    ],
    events: ["54", "55"],
    postId: 12,
    username: "herobrine",
  };

  const dummyEvents = [
    {
      eventId: 54,
      name: "Library Tour at NUST",
      cost: 0.0,
      address:
        "National University of Sciences and Technology, H-12, Islamabad, Pakistan",
      description:
        "A guided tour of the NUST library, showcasing its impressive collection of resources and modern facilities.",
      startTime: "2024-11-18T10:00:00Z",
      endTime: "2024-11-18T12:00:00Z",
      itineraryId: 12,
      eventImages: ["https://via.placeholder.com/150"],
    },
    {
      eventId: 55,
      name: "Picnic by the NUST Lake",
      cost: 0.0,
      address:
        "NUST Lake, National University of Sciences and Technology, Islamabad, Pakistan",
      description:
        "A relaxing picnic by the serene lake on the NUST campus, perfect for enjoying the natural beauty of the surroundings.",
      startTime: "2024-11-18T15:00:00Z",
      endTime: "2024-11-18T18:00:00Z",
      itineraryId: 12,
      eventImages: ["https://via.placeholder.com/150"],
    },
  ];

  const fetchItineraryData = async () => {
    try {
      // Fetch the specific itinerary and its events
      if (itinId === dummyItinerary.itineraryId) {
        setItinerary(dummyItinerary);
        setEvents(dummyEvents.filter((event) => event.itineraryId === itinId));
      } else {
        throw new Error("Itinerary not found");
      }
    } catch (error) {
      console.error("Error fetching itinerary data:", error);
      setErrorMessage("Failed to fetch itinerary data.");
    }
  };

  const formatTime = (startTime, endTime) => {
    const options = { hour: "numeric", minute: "numeric", hour12: true };
    const start = new Date(startTime).toLocaleTimeString("en-US", options);
    const end = new Date(endTime).toLocaleTimeString("en-US", options);
    return `${start} - ${end}`;
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
    <div className="container">
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
          Posted by <span className="username">@{itinerary.username}</span>
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
      <h5 className="ItinTitle"> Itinerary Overview</h5>
      <p className="boldText">
        <strong>Location</strong>
      </p>
      <p className="smallText">
        {itinerary.city}, {itinerary.country}
      </p>
      <p className="boldText">
        <strong>Price</strong> 
      </p>
      <p className="smallText">
        {itinerary.price === 0 ? "Free" : `$${itinerary.price}`}
      </p>
      <p className="boldText">
        <strong>Languages</strong> 
      </p>
      <p className="smallText">
        {itinerary.languages.join(", ")}
      </p>
      {/* Tags */}
      {Object.keys(categorizedTags).map((category) => (
        <React.Fragment key={category}>
          <p className="boldText">
            <strong>{category.replace(/_/g, " ").replace(/\b\w/g, (c) => c.toUpperCase())}</strong>
          </p>
          <p className="smallText">
            {categorizedTags[category].length > 0
              ? categorizedTags[category].join(", ")
              : "N/A"}
          </p>
        </React.Fragment>
      ))}
    </div>

    <div className="itineraryTable">
      <h5 className="ItinTitle">Itinerary Events</h5>
      <table>
        <thead>
          <tr>
            <th>Time</th>
            <th>Name</th>
            <th>Location</th>
            <th>Notes</th>
            <th>Cost</th>
          </tr>
        </thead>
        <tbody>
          {events.map((event) => (
            <tr key={event.eventId}>
              <td>{formatTime(event.startTime, event.endTime)}</td>
              <td>{event.name}</td>
              <td>{event.address}</td>
              <td>{event.description}</td>
              <td>{event.cost === 0 ? "Free" : `$${event.cost.toFixed(2)}`}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  </div>
</div>
  );
}

export default ItineraryDetails;
