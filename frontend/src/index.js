import React from 'react';
import ReactDOM from 'react-dom';
import { Router, Route, browserHistory, IndexRoute } from 'react-router';

// Local Components
import App from './App';
import Home from './views/Home';
import Create from './views/Create';


import registerServiceWorker from './registerServiceWorker';

ReactDOM.render((
      <Router history={browserHistory}>
          <Route path="/" component={App}>
              <IndexRoute component={Home}/>
              <Route path="create" component={Create} />
          </Route>
      </Router>
  ),
  document.getElementById('root')
);

registerServiceWorker();
