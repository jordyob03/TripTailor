import axios from 'axios';

const authAPI = axios.create({
  baseURL: 'http://localhost:8083',
  headers: {
    'Content-Type': 'application/json',
  },
});


authAPI.interceptors.request.use((config) => {
  const token = localStorage.getItem('token'); 
  if (token) {
    config.headers['Itinerary'] = `Bearer ${token}`; 
  }
  return config;
}, (error) => {
  return Promise.reject(error);
});

export default authAPI;
