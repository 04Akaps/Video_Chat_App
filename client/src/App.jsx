import React, {useEffect, useState} from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";

import Room from "./components/Room"
import Login from "./components/Login/Login";

import "./App.css"
import axiosInstance from "./api/axiosInstance";
import {useAuth} from "./context/user";
import Home from "./components/Logined/Home";
import {renderToStaticMarkup} from "react-dom/server";

function App() {
    const auth = useAuth()
    const [, setRender] = useState(0);

    useEffect(() => {
        axiosInstance.get("/auth/check-token").then((res) => {
            if (res.status != 200 ){
                auth.setUserName(null);
            }else {
                auth.setUserName(res.data);
            }
        });
    }, [auth]);

    useEffect(() => {
        auth.setUpdateCallback(() => {
            setRender((prev) => prev + 1);
        });
    }, [auth]);

    return <div className="App">
        <BrowserRouter>
            <Switch>
                {auth.userName ?
                    (
                        <Route path="/" exact component={Home}></Route>
                    ) :
                    (
                        <Route path="/" exact component={Login}></Route>
                    )
                }
                <Route path="/room/:roomID" component={Room}></Route>
            </Switch>
        </BrowserRouter>
    </div>;
}

export default App;
