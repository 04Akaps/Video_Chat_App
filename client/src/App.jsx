import React, { useEffect } from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";

import Room from "./components/Room"
import Home from "./components/Home";

import "./App.css"

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
