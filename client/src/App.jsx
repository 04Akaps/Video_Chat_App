import React, { useState } from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";

import CreateRoom from "./components/CreateRoom"
import Room from "./components/Room"

import "./App.css"
import Home from "./components/Home";

function App() {



    return <div className="App">
        <BrowserRouter>
            <Switch>
                {/*<Route path="/" exact component={CreateRoom}></Route>*/}
                <Route path="/" exact component={Home}></Route>
                <Route path="/room/:roomID" component={Room}></Route>
            </Switch>
        </BrowserRouter>
    </div>;
}

export default App;
