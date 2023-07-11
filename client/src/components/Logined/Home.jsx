import React from "react";
import "./Home.css"
import RoomList from "./RoomList";

const Home = (props) => {
    return (
        <div className="home-wrapper">
            <RoomList/>
        </div>
    );
};

export default Home;
