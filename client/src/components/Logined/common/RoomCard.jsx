import React, {useEffect, useState} from "react";

import "./RoomCard.css"
const RoomCard = (props) => {
    const handleClick = (hash) => {
        window.location.replace(`/room/${hash}`)
    }

    return (
        <div className="card-list-wrapper">
            {props.roomList.map((room, index) => {
                return (
                    <div className="room-list-box" key={index} onClick={() => {
                        handleClick(room.room_hash)
                    }}>
                        <div>owner : {room.owner_name}</div>
                        <div>Created : {room.created_at}</div>
                        <div>방송 여부 : {room.is_broad_cast ? "Yes": "No" }</div>
                        <div>이전 방송 날짜 : {room.before_broad_cast_time}</div>
                    </div>
                )
            })}
        </div>
    );
};

export default RoomCard;
