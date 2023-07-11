import React, {useEffect, useState} from "react";
import axiosInstance from "../../../api/axiosInstance";
import RoomCard from "../common/RoomCard";

import "./MyRoomList.css"

const MyRoomList = (props) => {
    const [myRoomList, SetMyRoomList]= useState([])

    useEffect(() => {
        axiosInstance.get("/room/my-room-list").then((res   ) => {
            if (res.status == 200) {
                SetMyRoomList(res.data)
            }else {
                SetMyRoomList(["Error"])
            }
        })
    },[])


    return (
        <div className="my-room-list-wrapper">
            <h1>My Room List</h1>
            <RoomCard roomList={myRoomList}/>
        </div>
    );
};

export default MyRoomList;
