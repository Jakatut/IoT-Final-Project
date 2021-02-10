import './bootstrap.min.css'
import './global.css'

import HeaderBar from './components/HeaderBar';
import Home from './views/home'

import {
  BrowserRouter as Router,
  Switch
} from "react-router-dom";
import ProtectedRoute from "./auth/protected-route";



function App() {
  return (
    <Router>
    <div>
      <HeaderBar/>
      <Switch>
          <ProtectedRoute path="/" exact component={Home} />            
        </Switch>

    </div>
    </Router>  
  );
}

export default App;
