import axios from 'axios';

const feedAPI  = axios.create({
  baseURL: 'http://localhost:8093',
  headers: {
    'Content-Type': 'application/json',
  },
});


export default feedAPI ;
