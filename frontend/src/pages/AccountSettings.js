import React from 'react';
import '../styles/styles.css'; 
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUser, faLock, faBell, faShieldAlt, faGlobe, faBars } from '@fortawesome/free-solid-svg-icons'; // Import all needed icons

function AccountSettings() {

  const sections = [
    { title: 'Personal Info', description: 'Provide personal details and how we can reach you', icon: faUser },
    { title: 'Login & Security', description: 'Update your password and secure your account', icon: faLock },
    { title: 'Notifications', description: 'Choose notification preferences', icon: faBell },
    { title: 'Privacy & Sharing', description: 'Manage your personal data and sharing settings', icon: faShieldAlt },
    { title: 'Global Preferences', description: 'Default language, currency, and time zone', icon: faGlobe },
  ];

  return (
    <>
      {/* Main Container */}
      <div className="centeredContainer">
        <div className="gridWrapper">
          <div className="grid">
            {sections.map((section) => (
              <div key={section.title} className="card">
                <div className="cardContent">
                  <h3 className="cardTitle">{section.title}</h3>
                  <p className="cardDescription">{section.description}</p>
                </div>
                <FontAwesomeIcon icon={section.icon} className="icon" /> 
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  );
}

export default AccountSettings;
