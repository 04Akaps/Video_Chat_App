import React from "react";

import MyRoomList from "./subComponents/MyRoomList";
import RecentlyCreated from "./subComponents/RecentlyCreated";

import "./RoomList.css"

const RoomList = (props) => {

    return (
        <div className="room-list-wrapper">
            <div><MyRoomList/></div>
            <div><RecentlyCreated/></div>
        </div>
    );
};

export default RoomList;
