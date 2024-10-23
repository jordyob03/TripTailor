import axios from 'axios';

const itineraryAPI  = axios.create({
  baseURL: 'http://localhost:8083',
  headers: {
    'Content-Type': 'application/json',
  },
});


export default itineraryAPI ;
