import React from "react";
import "./Home.css"
import NavBar from "./NavBar.jsx";
import RoomList from "./RoomList";

const Home = (props) => {
    return (
        <div className="home-wrapper">
            <NavBar/>
            <RoomList/>
        </div>
    );
};

export default Home;
