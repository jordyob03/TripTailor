import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom'; // Import useNavigate
import logo1 from '../assets/logo-long-transparent.png';
import logo2 from '../assets/logo-circle-white.png';


function UserSignup() {
  // State for form fields
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isHovered, setIsHovered] = useState(false);

  const navigate = useNavigate();

  // Handle form submission
  const handleSubmit = (e) => {
    e.preventDefault();
    navigate('/profile-creation');
  };

  // Inline styles
  const styles = {
    navbar: {
      display: 'flex',
      justifyContent: 'space-between', 
      alignItems: 'center',
      padding: '1vw 1vw',
      height: '50px',
      width: '100vw', 
      backgroundColor: 'white',
      boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)', 
      position: 'fixed', 
      top: 0, 
      left: 0,
      zIndex: 2,
    },
    logo1: {
      width: '150px',
      marginLeft: '80px',
      marginTop: '5px',
    },
    profileButton: {
      display: 'flex',        
      alignItems: 'center',   
      justifyContent: 'space-between', 
      padding: '10px 20px',
      backgroundColor: 'white',
      border: '1px solid #dfdfdf', 
      borderRadius: '30px',
      cursor: 'pointer',
      marginRight: '160px',
      boxShadow: '0 2px 2px rgba(0, 0, 0, 0.1)',
    },
    logo2: {
      width: '120px',
      marginBottom: '0px',
    },
    container: {
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
      minHeight: 'calc(100vh - 60px)', 
      width: '100vw',
      backgroundColor: 'white', 
    },
    box: {
      backgroundColor: 'white',
      padding: '40px',
      borderRadius: '20px',
      boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)',
      textAlign: 'center',
      width: '350px',
      zIndex: 1,
      marginTop: '80px', 
      marginBottom: '40px', 
    },
    heading: {
      color: '#002f6c', 
      marginTop: '5px', 
      marginBottom: '10px', 
      fontFamily: "'Red Hat Display', sans-serif",
      fontSize: '24px'
    },
    subheading: {
      color: '#002f6c', 
      marginTop: '0px', 
      marginBottom: '20px', 
      fontFamily: "'Red Hat Display', sans-serif",
      fontSize: '16px'
    },
    separator: {
      width: '100%',      
      border: 'none',   
      borderTop: '2px solid #d4d4d4', 
      margin: '15px auto',  
    },    
    form: {
      display: 'flex',
      flexDirection: 'column',
      gap: '15px',
    },
    input: {
      padding: '12px',
      borderRadius: '4px',
      border: '1px solid #ccc',
      fontSize: '16px',
      fontFamily: "'Red Hat Display', sans-serif", 
    },
    button: {
      padding: '12px',
      backgroundColor: isHovered ? '#00509e' : '#002f6c',
      color: 'white',
      border: 'none',
      borderRadius: '4px',
      fontSize: '16px',
      cursor: 'pointer',
      transition: 'background-color 0.3s ease',
      fontFamily: "'Red Hat Display', sans-serif", 
    },
  };

  return (
    <>
      {/* Navbar */}
      <nav style={styles.navbar}>
        <img src={logo1} alt="Trip Tailor Logo" style={styles.logo1} />

        {/* Profile Button */}
        <button style={styles.profileButton}>
          <i className="fas fa-bars" style={{ fontSize: '16px', color: '#00509e', marginRight: '15px' }}></i> 
          <i className="fa-regular fa-user" style={{ fontSize: '24px', color: '#00509e' }}></i>
        </button>
      </nav>

      {/* Main Container */}
      <div style={styles.container}>
        <div style={styles.box}>
          <img src={logo2} alt="Trip Tailor Logo" style={styles.logo2} />
          <h5 style={styles.heading}>Welcome to Trip Tailor</h5>
          <hr style={styles.separator} />
          <h6 style={styles.subheading}>Sign up</h6>
          <form onSubmit={handleSubmit} style={styles.form}>
            <input
              type="text"
              placeholder="Username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
              style={styles.input}
            />
            <input
              type="email"
              placeholder="Email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              style={styles.input}
            />
            <input
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              style={styles.input}
            />
            <button
              type="submit"
              style={styles.button}
              onMouseEnter={() => setIsHovered(true)}
              onMouseLeave={() => setIsHovered(false)}
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
