import React, { FormEvent, useRef } from "react";
import logo from "./logo.svg";
import "./App.css";
import { Form } from "./components/Form";
import { Title } from "./components/Title";


function App() {
  return (
    <div>
      <Title onClick={() => alert('Primeiro')}>Iniciando com React</Title>
      <Title onClick={() => alert('Segundo')} color="green">Iniciando com React</Title>
      <Form />
    </div>
  );
}

export default App;