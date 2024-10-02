import React from 'react';
import logo1 from '../assets/logo-long-transparent.png'; // Same logo


function Account() {
 const styles = {
   navbar: {
     display: 'flex',
     justifyContent: 'space-between',
     alignItems: 'center',
     padding: '10px 20px',
     height: '60px',
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
     padding: '10px 20px',
     backgroundColor: 'white',
     border: '1px solid #dfdfdf',
     borderRadius: '30px',
     cursor: 'pointer',
     marginRight: '160px',
     boxShadow: '0 2px 2px rgba(0, 0, 0, 0.1)',
   },
   container: {
     display: 'flex',
     justifyContent: 'center',
     alignItems: 'center',
     flexDirection: 'column',
     minHeight: 'calc(100vh - 60px)', // Adjust for full screen
     width: '100vw',
     backgroundColor: 'white',
     paddingTop: '80px', // Add 80px padding at the top to make room for the navbar
   },
   gridWrapper: {
     display: 'flex',
     justifyContent: 'center',
     alignItems: 'center',
     width: '100%',
     maxWidth: '900px', // Fixed width to fit the boxes within a set size
     padding: '40px',
     boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)',
     backgroundColor: 'white',
     borderRadius: '20px',
   },
   grid: {
     display: 'grid',
     gridTemplateColumns: '1fr 1fr', // Two columns of boxes
     gap: '20px',
     width: '100%',
   },
   card: {
     backgroundColor: 'white',
     padding: '20px',
     borderRadius: '12px',
     boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)',
     display: 'flex',
     alignItems: 'center',
     cursor: 'pointer',
     transition: 'box-shadow 0.3s ease',
     justifyContent: 'space-between',
   },
   cardContent: {
     display: 'flex',
     flexDirection: 'column',
   },
   cardTitle: {
     fontSize: '18px',
     color: '#002f6c',
     fontFamily: "'Red Hat Display', sans-serif",
     marginBottom: '5px',
   },
   cardDescription: {
     fontSize: '14px',
     color: '#555',
   },
   icon: {
     fontSize: '24px',
     color: '#00509e',
     marginRight: '15px',
   },
 };


 // List of sections (you can add actual links later)
 const sections = [
   { title: 'Personal Info', description: 'Provide personal details and how we can reach you', icon: 'fa-user' },
   { title: 'Login & Security', description: 'Update your password and secure your account', icon: 'fa-lock' },
   { title: 'Notifications', description: 'Choose notification preferences', icon: 'fa-bell' },
   { title: 'Privacy & Sharing', description: 'Manage your personal data and sharing settings', icon: 'fa-shield-alt' },
   { title: 'Global Preferences', description: 'Default language, currency, and time zone', icon: 'fa-globe' },
 ];


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
       <div style={styles.gridWrapper}>
         <div style={styles.grid}>
           {sections.map((section) => (
             <div key={section.title} style={styles.card}>
               <div style={styles.cardContent}>
                 <h3 style={styles.cardTitle}>{section.title}</h3>
                 <p style={styles.cardDescription}>{section.description}</p>
               </div>
               <i className={`fas ${section.icon}`} style={styles.icon}></i>
             </div>
           ))}
         </div>
       </div>
     </div>
   </>
 );
}


export default Account;


