import * as React from 'react';
import { Link, RouteComponentProps } from 'react-router-dom';

interface Props extends RouteComponentProps {}

const Page1 = ({ history }: Props) => {
  
  return (
    <div>
      <button onClick={history.goBack}>Previous Page</button>
      <Link to="/">Top</Link>
      <Link to="/page2">Page 2</Link>
    </div>
  );
};

export default Page1;
