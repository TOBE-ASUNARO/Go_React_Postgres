import * as React from 'react';
import { BrowserRouter, Route , Routes} from 'react-router-dom';

import Top from './Features/Top';
import Page1 from './Features/Page1';
import Page2 from './Features/Page2';

const Router = () => {
  return (
    
  <BrowserRouter>
    <Routes>
      <Route exact={true} path="/" component={Top} />
      <Route path="/page1" component={Page1} />
      <Route path="/page1" component={Page2} />
      <Route path="/page2" component={Page2} />
       {/* Not Found */}
      {/* <Route component={() => <Redirect to="/" />} /> */}
    </Routes>
  </BrowserRouter>
    
  );
};

export default Router;
