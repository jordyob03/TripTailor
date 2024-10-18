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
  const [dob, setDob] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const [continueIsHovered, setContinueIsHovered] = useState(false);

  const navigate = useNavigate();

  const validateUsername = (username) => {
    const noWhitespace = !/\s/.test(username); 
    const validChars = /^[A-Za-z0-9_]+$/.test(username); 
    const lengthValid = username.length > 0 && username.length <= 24; 
    
    if (!noWhitespace) {
      return 'Username must not contain spaces.';
    }
    if (!validChars) {
      return 'Username can only contain letters, numbers, and underscores.';
    }
    if (!lengthValid) {
      return 'Username must be less than 24 characters.';
    }
    return ''; 
  };

  const validatePassword = (password) => {
    const containsLetter = /[a-zA-Z]/.test(password);
    const containsNumber = /\d/.test(password);
    const containsSpecial = /[^a-zA-Z0-9]/.test(password);
    const noSpaces = !/\s/.test(password); 
    const lengthValid = password.length >= 8 && password.length <= 24;

    if (!containsLetter) {
      return 'Password must contain at least one letter.';
    }
    if (!containsNumber) {
      return 'Password must contain at least one number.';
    }
    if (!containsSpecial) {
      return 'Password must contain at least one special character.';
    }
    if (!noSpaces) {
      return 'Password must not contain spaces.';
    }
    if (!lengthValid) {
      return 'Password must be between 8 and 24 characters.';
    }
    return ''; 
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    // Validate the username
    const usernameError = validateUsername(username);
    if (usernameError) {
      setErrorMessage(usernameError);
      return;
    }

    // Validate the password
    const passwordError = validatePassword(password);
    if (passwordError) {
      setErrorMessage(passwordError);
      return;
    }

    // Confirm passwords match
    if (password !== confirmPassword) {
      setErrorMessage('Passwords do not match. Please try again.');
      return;
    }
    
    setErrorMessage('');

    const userData = {
      username: username,
      email: email,
      password: password,
      dateOfBirth: dob
    };

    try {
      const response = await authAPI.post('/signup', userData);
      const { token, username } = response.data;
      localStorage.setItem('token', token);  
      localStorage.setItem('username', username);
      console.log('User signed up successfully:', response.data);

      navigate('/profile-creation');
    } catch (error) {
      if (error.response && error.response.data) {
        setErrorMessage(error.response.data.error); 
      } else {
        setErrorMessage('Signup failed. Please try again.');
      }
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

            {/* Display error message */}
            {errorMessage && <div className="errorMessage">{errorMessage}</div>}

            <hr className="separatorLine" />

            <h6 className="subheading1">Date of Birth</h6>

            {/* Date of Birth Input */}
            <input
              type="date"
              value={dob}
              onChange={(e) => setDob(e.target.value)}
              required
              className="input"
            />

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
