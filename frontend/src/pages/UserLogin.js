import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import navBarLogo from '../assets/logo-long-transparent.png';
import authLogo from '../assets/logo-circle-white.png';
import '../styles/styles.css';  // Import external CSS file

function UserLogin() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [continueIsHovered, setContinueIsHovered] = useState(false);
  const [signUpIsHovered, setSignUpIsHovered] = useState(false);

  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    navigate('/profile-creation');
  };

  const handleSignUpClick = () => {
    navigate('/sign-up');
  };

  return (
    <>
      {/* navBar */}
      <nav className="navBar">
        <img src={navBarLogo} alt="Trip Tailor Logo" className="navBarLogo" />
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
          <h6 className="subheading">Log in</h6>
          <form onSubmit={handleSubmit} className="form">
            <input
              type="text"
              placeholder="Username or Email"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
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
            <button
              type="submit"
              className={`continueButton ${continueIsHovered ? 'hovered' : ''}`}
              onMouseEnter={() => setContinueIsHovered(true)}
              onMouseLeave={() => setContinueIsHovered(false)}
            >
              Continue
            </button>
          </form>
          <button
            type="button"
            className={`signUpButton ${signUpIsHovered ? 'hoveredSignUp' : ''}`}
            onClick={handleSignUpClick}
            onMouseEnter={() => setSignUpIsHovered(true)}
            onMouseLeave={() => setSignUpIsHovered(false)}
          >
            Don't have an account? Sign up
          </button>
        </div>
      </div>
    </>
  );
}

export default UserLogin;
