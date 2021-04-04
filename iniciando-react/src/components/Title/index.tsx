<<<<<<< HEAD
// @flow
import * as React from "react";
//import "./Title.css";

interface TitleProps { 
  color?: string;
  onClick: () => void
}

export const Title: React.FunctionComponent<TitleProps> = (props) => {
  const { color = "red", children, onClick } = props;

  return (
    <h1
      className="Title"
      style={{
        color: color,
      }}
      onClick={onClick}
    >
      {children}
    </h1>
  );
=======
// @flow
import * as React from "react";
//import "./Title.css";

interface TitleProps { 
  color?: string;
  onClick: () => void
}

export const Title: React.FunctionComponent<TitleProps> = (props) => {
  const { color = "red", children, onClick } = props;

  return (
    <h1
      className="Title"
      style={{
        color: color,
      }}
      onClick={onClick}
    >
      {children}
    </h1>
  );
>>>>>>> 9730c3e8e238fcf70168252ad0cf984e24521105
};