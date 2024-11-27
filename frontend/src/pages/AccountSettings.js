import React from 'react';
import '../styles/styles.css'; 
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUser, faLock, faBell, faShieldAlt, faGlobe } from '@fortawesome/free-solid-svg-icons';
import { useNavigate } from 'react-router-dom';

function AccountSettings() {
  const navigate = useNavigate();

  // Updated sections with paths for navigation
  const sections = [
    { title: 'Personal Info', description: 'Provide or update personal details', icon: faUser, path: '/personal-info' },
    { title: 'Login & Security', description: 'Update your password and secure your account', icon: faLock, path: '/login-security' },
    { title: 'Notifications', description: 'Choose notification preferences', icon: faBell, path: '/notifications' },
    { title: 'Privacy & Sharing', description: 'Manage your personal data and sharing settings', icon: faShieldAlt, path: '/privacy-sharing' },
    { title: 'Global Preferences', description: 'Default language, currency, and time zone', icon: faGlobe, path: '/global-preferences' },
  ];

  return (
    <div className="centeredContainer">
      <div className="gridWrapper">
        <div className="grid">
          {sections.map((section) => (
            <div
              key={section.title}
              className="card"
              onClick={() => {navigate(section.path); window.location.reload();}} // Navigate to the path
              style={{ cursor: 'pointer' }} // Add a pointer cursor for better UX
            >
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
  );
}

export default AccountSettings;
