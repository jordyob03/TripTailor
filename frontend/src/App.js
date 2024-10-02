import './styles/App.css';
import Account from './pages/AccountSettings';
import UserAuthentication from './pages/UserSignup';
import UserProfile from './pages/InitialUserProfile';
import Login from './pages/UserLogin';


function App() {
  return (
    <div className="App">
      <header className="App-header">
        <UserProfile/>
        <UserAuthentication/>
        <Account/>
        <Login/>
      </header>
    </div>
  );
}
export default App;
