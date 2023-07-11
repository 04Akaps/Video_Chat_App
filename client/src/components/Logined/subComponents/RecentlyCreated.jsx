import React, {useEffect, useState} from "react";
import RoomCard from "../common/RoomCard";

import "./MyRoomList.css"
import axiosInstance from "../../../api/axiosInstance";

const RecentlyCreated = (props) => {
    const [roomList, SetRoomLIst]= useState([])

    useEffect(() => {
        axiosInstance.get("/room/recently-created-room-list").then((res   ) => {
            if (res.status == 200) {
                SetRoomLIst(res.data)
            }else {
                SetRoomLIst(["Error"])
            }
        })
    },[])
    return (
        <div className="my-room-list-wrapper">
            <h1>RecentlyCreated Room List</h1>
            <RoomCard roomList={roomList}/>
        </div>
    );
};

export default RecentlyCreated;
