import React, { useEffect, useRef } from "react";

const Room = (props) => {
    const userVideo = useRef();
    const partnerVideo = useRef();

    return (
        <div>
            <video autoPlay controls={true} ref={userVideo}></video>
            <video autoPlay controls={true} ref={partnerVideo}></video>
        </div>
    );
};

export default Room;
