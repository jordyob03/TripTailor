import React, { useState } from "react";
import searchAPI from "../api/searchAPI";

function Search() {
    const [country, setCountry] = useState('');
    const [city, setCity] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();

        const searchData = {
            country: country, 
            city: city,
        }
        try {
            console.log("search api sent")
            console.log(searchData)
            const response = await searchAPI.get('/search', {
                params: searchData,
            });
            console.log("API response:", response);
            console.log("idk if this is reached")
            console.log(searchData)

        } catch (error) {
            if (error.response && error.response.data) {
                setErrorMessage(error.response.data.error); 
              } else {
                setErrorMessage('Search Failed');
              }
        }
    };

    return (
        <div>
            <h2>Search</h2>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Country:</label>
                    <input 
                        type="text" 
                        value={country} 
                        onChange={(e) => setCountry(e.target.value)} 
                        placeholder="Enter country" 
                    />
                </div>
                <div>
                    <label>City:</label>
                    <input 
                        type="text" 
                        value={city} 
                        onChange={(e) => setCity(e.target.value)} 
                        placeholder="Enter city" 
                    />
                </div>
                <button type="submit">Search</button>
            </form>
        </div>
    );
}

export default Search;
