import React, { useState } from 'react';


function UserAuthentication() {
 // State for form fields
 const [username, setUsername] = useState('');
 const [email, setEmail] = useState('');
 const [password, setPassword] = useState('');


 // Handle form submission
 const handleSubmit = (e) => {
   e.preventDefault();
   console.log({ username, email, password });
 };


 // Inline styles
 const styles = {
  container: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    height: '100vh',
    width: '100vw', // Ensures the container spans the full width
    backgroundColor: 'white', // White background to cover the full page
  }
  ,
   box: {
     backgroundColor: 'white',
     padding: '40px',
     borderRadius: '8px',
     boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)',
     textAlign: 'center',
     width: '350px',
   },
   logo: {
     width: '120px',
     marginBottom: '20px',
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
     backgroundColor: '#002f6c', // Navy blue
     color: 'white',
     border: 'none',
     borderRadius: '4px',
     fontSize: '16px',
     cursor: 'pointer',
   },
   buttonHover: {
     backgroundColor: '#00509e',
   }
 };


 return (
   <div style={styles.container}>
     <div style={styles.box}>
       {/* Placeholder logo */}
       <img src="trip-tailor-logo.png" alt="Trip Tailor Logo" style={styles.logo} />
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
         <button type="submit" style={styles.button}>
           Continue
         </button>
       </form>
     </div>
   </div>
 );
}


export default UserAuthentication;
