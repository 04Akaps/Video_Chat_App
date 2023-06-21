import React, { useEffect, useRef , useState} from "react";


import "./Room.css"

const Room = (props) => {
    const videoRef = useRef(null);
    const [isVideoVisible, setIsVideoVisible] = useState(false);

    useEffect(() => {
        if (isVideoVisible) {
            navigator.mediaDevices.getUserMedia({ video: true })
                .then((stream) => {
                    if (videoRef.current) {
                        videoRef.current.srcObject = stream;
                    }
                })
                .catch((error) => {
                    console.error('Error accessing camera:', error);
                });
        } else {
            if (videoRef.current) {
                const currentStream = videoRef.current.srcObject;
                const tracks = currentStream?.getTracks();
                if (tracks) {
                    tracks.forEach((track) => {
                        track.stop();
                    });
                }
                videoRef.current.srcObject = null;
            }
        }
    }, [isVideoVisible]);

    const toggleVideo = () => {
        setIsVideoVisible(!isVideoVisible);
    };

    return (
        <div >
            <div className={"video-wrapper"}>
                {isVideoVisible && <video ref={videoRef} autoPlay playsInline />}
            </div>

            <button onClick={toggleVideo}>
                {isVideoVisible ? 'Hide' : 'Show'}
            </button>
        </div>
    );
};

export default Room;
