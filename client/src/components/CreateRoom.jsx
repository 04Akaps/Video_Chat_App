import React from "react";

import "./CreateRoom.css"
import axios from "axios";

const CreateRoom = (props) => {
    const create = async (e) => {
        e.preventDefault();

        const resp = await axios.post("http://localhost:8000/create")
        const { room_id } =  resp.data

        props.history.push(`/room/${room_id}`)
    };

    return (
        <div className="create-room-layout">
            <button onClick={create}>Create Room</button>
        </div>
    );
};

export default CreateRoom;
