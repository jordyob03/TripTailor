import './styles/App.css';
import HelloWorld from './components/HelloWorld';
import UserAuthentication from './pages/UserAuthetication';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <UserAuthentication/>
      </header>
    </div>
  );
}

export default App;
