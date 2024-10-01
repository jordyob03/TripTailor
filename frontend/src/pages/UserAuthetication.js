import React, { useState } from 'react';

function UserAuthentication() {
  // State for form fields
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isHovered, setIsHovered] = useState(false);

  // Handle form submission
  const handleSubmit = (e) => {
    e.preventDefault();
    console.log({ username, email, password });
    // Add any form validation or API call here
  };

  // Inline styles
  const styles = {
    navbar: {
      display: 'flex',
      justifyContent: 'flex-start',
      alignItems: 'center',
      padding: '10px 20px',
      height: '60px',
      width: '100vw', // Extend the navbar to fit the full page width
      backgroundColor: 'white',
      boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)', // Light shadow for the navbar
      position: 'fixed', // Keep navbar fixed at the top
      top: 0, // Align it to the top of the page
      left: 0,
      zIndex: 2, // Ensure the navbar stays on top of other elements
    },
    logo: {
      width: '120px',
    },
    container: {
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
      height: 'calc(100vh - 60px)', // Adjust height to account for navbar
      width: '100vw',
      backgroundColor: 'white', // Extend white background to the full page
    },
    box: {
      backgroundColor: 'white',
      padding: '40px',
      borderRadius: '8px',
      boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)',
      textAlign: 'center',
      width: '350px',
      zIndex: 1, // Ensure the form stays on top
    },
    heading: {
      color: '#002f6c', // Navy blue
      marginBottom: '20px',
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
    },
    button: {
      padding: '12px',
      backgroundColor: isHovered ? '#00509e' : '#002f6c', // Change on hover
      color: 'white',
      border: 'none',
      borderRadius: '4px',
      fontSize: '16px',
      cursor: 'pointer',
      transition: 'background-color 0.3s ease',
    },
  };

  return (
    <>
      {/* Navbar */}
      <nav style={styles.navbar}>
        <img src="trip-tailor-logo.png" alt="Trip Tailor Logo" style={styles.logo} />
      </nav>

      {/* Main Container */}
      <div style={styles.container}>
        <div style={styles.box}>
          <h2 style={styles.heading}>Log in or Sign Up</h2>
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

export default UserAuthentication;
