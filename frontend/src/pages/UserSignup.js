import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import navBarLogo from '../assets/logo-long-transparent.png';
import authLogo from '../assets/logo-circle-white.png';
import '../styles/styles.css'; 
import authAPI from '../api/authAPI.js';

function UserSignup() {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const [continueIsHovered, setContinueIsHovered] = useState(false);

  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (password === confirmPassword) {
      setErrorMessage('');

      const userData = {
        username: username, 
        email: email,
        password: password,
        dateOfBirth: '1990-01-01'
      }

      try {
        const response = await authAPI.post('/signup', userData);
        const { token } = response.data;  
        localStorage.setItem('token', token);  
        console.log('User signed up successfully:', response.data);

        navigate('/profile-creation');
      } catch (error) {
        if (error.response && error.response.data) {
          setErrorMessage(error.response.data.error); 
        } else {
          setErrorMessage('Signup failed. Please try again.');
        }
      }


    } else {
      setErrorMessage('Passwords do not match. Please try again.');
    }
  };

  return (
    <>
      {/* navBar */}
      <nav className="navBar">
        <img src={navBarLogo} alt="Trip Tailor Logo" className="navBarLogo" />

        {/* Profile Button */}
        <button className="profileButton">
          <i className="fas fa-bars" style={{ fontSize: '16px', color: '#00509e', marginRight: '15px' }}></i>
          <i className="fa-regular fa-user" style={{ fontSize: '24px', color: '#00509e' }}></i>
        </button>
      </nav>

      {/* Main Container */}
      <div className="centeredContainer">
        <div className="centeredBox">
          <img src={authLogo} alt="Trip Tailor Logo" className="authLogo" />
          <h5 className="heading">Welcome to Trip Tailor</h5>
          <hr className="separatorLine" />
          <h6 className="subheading">Sign up</h6>
          <form onSubmit={handleSubmit} className="form">
            <input
              type="text"
              placeholder="Username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
              className="input"
            />
            <input
              type="email"
              placeholder="Email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              className="input"
            />
            <input
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              className="input"
            />
            <input
              type="password"
              placeholder="Confirm Password"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)} 
              required
              className="input"
            />

            {/* Display error message if passwords don't match */}
            {errorMessage && <div className="errorMessage">{errorMessage}</div>}

            <button
              type="submit"
              className={`continueButton ${continueIsHovered ? 'hovered' : ''}`}
              onMouseEnter={() => setContinueIsHovered(true)}
              onMouseLeave={() => setContinueIsHovered(false)}
            >
              Continue
            </button>
          </form>
        </div>
      </div>
    </>
  );
}

export default UserSignup;
