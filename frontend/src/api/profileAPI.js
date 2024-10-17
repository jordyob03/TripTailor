import axios from 'axios';

const mainAPI = axios.create({
    baseURL: 'http://localhost:8085', 
    headers: {
      'Content-Type': 'application/json',
    },
  });
  
  export default mainAPI;
  