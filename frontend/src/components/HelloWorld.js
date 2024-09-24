import React, { useState, useEffect } from "react";

const HelloWorld = () => {
  const [message, setMessage] = useState("");

  useEffect(() => {
    fetch("http://localhost:8080/hello")
      .then((response) => response.text())
      .then((data) => setMessage(data))
      .catch((error) => console.error("Error fetching message:", error));
  }, []);

  return (
    <div>
      <h1>{message || "Loading..."}</h1>
    </div>
  );
};

export default HelloWorld;
