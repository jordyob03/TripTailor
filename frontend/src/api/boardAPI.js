import axios from 'axios';

const boardAPI  = axios.create({
  baseURL: 'http://localhost:8086',
  headers: {
    'Content-Type': 'application/json',
  },
});


export default boardAPI ;
