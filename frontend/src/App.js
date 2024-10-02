import './styles/App.css';

import UserAuthentication from './pages/UserAuthetication';
import UserProfile from './pages/InitialUserProfile';


function App() {
  return (
    <div className="App">
      <header className="App-header">
        <UserProfile/>
      </header>
    </div>
  );
}
export default App;
