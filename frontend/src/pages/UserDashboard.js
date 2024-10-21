import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import navBarLogo from '../assets/logo-long-transparent.png';
import '../styles/styles.css';


function MyDashboard() {
 const [selectedTab, setSelectedTab] = useState('itineraries');
 const [itineraries, setItineraries] = useState([]);
 const [trips, setTrips] = useState([]);
 const [menuOpen, setMenuOpen] = useState(false);
 const navigate = useNavigate();


 const handleTabChange = (tab) => {
   setSelectedTab(tab);
 };


 const handleCreateItinerary = () => {
   navigate('/create-itinerary');
 };


 const toggleMenu = () => {
   setMenuOpen(!menuOpen);
 };


 return (
   <div>
     {/* Navbar */}
     <nav className="navBar">
       <img src={navBarLogo} alt="Trip Tailor Logo" className="navBarLogo" />


       <div style={{ marginLeft: 'auto', display: 'flex', alignItems: 'center' }}>
         <button
           className="createItineraryButton"
           onClick={handleCreateItinerary}
           style={{
             backgroundColor: '#00509e',
             color: '#ffffff',
             borderRadius: '25px',
             padding: '10px 20px',
             fontSize: '16px',
             fontWeight: 'bold',
             border: 'none',
             cursor: 'pointer',
             marginRight: '15px',
           }}
         >
           Create Itinerary
         </button>


         <button className="profileButton" onClick={toggleMenu}>
           <i className="fas fa-bars" style={{ fontSize: '16px', color: '#00509e', marginRight: '15px' }}></i>
           <i className="fa-regular fa-user" style={{ fontSize: '24px', color: '#00509e' }}></i>
         </button>
       </div>


       {/* Dropdown Menu */}
       {menuOpen && (
         <div style={{
           position: 'absolute',
           top: '60px',
           right: '10px',
           transform: 'translateX(-50%)',
           backgroundColor: '#00509e',
           color: 'white',
           boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)',
           borderRadius: '8px',
           width: '150px',
           zIndex: 3
         }}>
           <ul style={{
             listStyle: 'none',
             padding: '10px 0',
             margin: '0',
             textAlign: 'left'
           }}>
             <li style={{ padding: '10px 20px', cursor: 'pointer', color: 'white' }} onClick={() => navigate('/profile-creation')}>Profile</li>
             <li style={{ padding: '10px 20px', cursor: 'pointer', color: 'white' }} onClick={() => navigate('/account')}>Account Settings</li>
             <li style={{ padding: '10px 20px', cursor: 'pointer', color: 'white' }} onClick={() => navigate('/dashboard')}>My Itineraries</li>
             <li style={{ padding: '10px 20px', cursor: 'pointer', color: 'white' }} onClick={() => navigate('/')}>Home</li>
           </ul>
         </div>
       )}
     </nav>


     {/* Main Heading */}
     <div style={{ textAlign: 'center', marginTop: '100px' }}>
       <h2 style={{ fontSize: '32px', fontWeight: 'bold', marginBottom: '10px', color: '#00509e' }}>My Travels</h2>
       <hr style={{ width: '90%', border: '1px solid #cccccc', margin: '0 auto 20px auto' }} />
     </div>


     {/* Tab Navigation */}
     <div style={{
       display: 'flex',
       justifyContent: 'center',
       gap: '40px',
       marginBottom: '30px'
     }}>
       <button
         className={selectedTab === 'itineraries' ? 'activeTab' : ''}
         onClick={() => handleTabChange('itineraries')}
         style={{
           fontSize: '18px',
           fontWeight: 'bold',
           border: 'none',
           background: 'none',
           color: '#00509e',
           borderBottom: selectedTab === 'itineraries' ? '3px solid #00509e' : 'none',
           paddingBottom: '5px',
           cursor: 'pointer'
         }}
       >
         Itineraries
       </button>
       <button
         className={selectedTab === 'trips' ? 'activeTab' : ''}
         onClick={() => handleTabChange('trips')}
         style={{
           fontSize: '18px',
           fontWeight: 'bold',
           border: 'none',
           background: 'none',
           color: '#00509e',
           borderBottom: selectedTab === 'trips' ? '3px solid #00509e' : 'none',
           paddingBottom: '5px',
           cursor: 'pointer'
         }}
       >
         Trips
       </button>
     </div>


     {/* Dashboard Content */}
     <div className="dashboardContent" style={{ marginTop: '20px' }}>
       {/* Content Based on Selected Tab */}
       <div className="tabContent">
         {selectedTab === 'itineraries' && (
           <div className="itinerariesList" style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '200px' }}>
             {itineraries.length > 0 ? (
               itineraries.map((itinerary, index) => (
                 <div key={index} className="itineraryCard">
                   <img src={itinerary.image} alt={itinerary.title} className="itineraryImage" />
                   <div className="itineraryInfo">
                     <h3 style={{ color: '#00509e' }}>{itinerary.title}</h3>
                     <p style={{ color: '#00509e' }}>{itinerary.description}</p>
                     <span style={{ color: '#00509e' }}>{itinerary.location}</span>
                     {/* Tags or other info */}
                   </div>
                 </div>
               ))
             ) : (
               <div style={{
                 backgroundColor: '#f0f8ff',
                 padding: '15px 25px',
                 borderRadius: '20px',
                 textAlign: 'center',
                 fontSize: '18px',
                 fontWeight: 'bold',
                 color: '#00509e',
                 boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)'
               }}>
                 No itineraries found. Create a new one to get started!
               </div>
             )}
           </div>
         )}


         {selectedTab === 'trips' && (
           <div className="tripsList" style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '200px' }}>
             {trips.length > 0 ? (
               trips.map((trip, index) => (
                 <div key={index} className="tripCard">
                   <img src={trip.image} alt={trip.title} className="tripImage" />
                   <div className="tripInfo">
                     <h3 style={{ color: '#00509e' }}>{trip.title}</h3>
                     <p style={{ color: '#00509e' }}>{trip.description}</p>
                     <span style={{ color: '#00509e' }}>{trip.location}</span>
                     {}
                   </div>
                 </div>
               ))
             ) : (
               <div style={{
                 backgroundColor: '#f0f8ff',
                 padding: '15px 25px',
                 borderRadius: '20px',
                 textAlign: 'center',
                 fontSize: '18px',
                 fontWeight: 'bold',
                 color: '#00509e',
                 boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)'
               }}>
                 No trips saved. Add some itineraries to your trips to start planning!
               </div>
             )}
           </div>
         )}
       </div>
     </div>
   </div>
 );
}


export default MyDashboard;