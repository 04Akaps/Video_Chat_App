import React, { useLayoutEffect, useRef, useEffect, useState} from "react";

import "./Room.css"
import {useAuth} from "../context/user";
import axiosInstance from "../api/axiosInstance";

import { io, Socket } from "socket.io-client";


const Room = (props) => {
    const [isOwner, setIsOwner] = useState(false);
    const [userName, setUser] = useState("")

    const [initLoading, setInitLoading] = useState(true)

    const user = useAuth()
    const currentPath = window.location.pathname;
    const roomID = currentPath.split("/")[2];

    const socketRef = useRef();
    const myVideoRef = useRef(null);
    const pcRef = useRef();

    const getMedia = async () => {
        try {
            const stream = await navigator.mediaDevices.getUserMedia({
                video: true,
                audio: true,
            });

            myVideoRef.current.srcObject = stream;

            if (!(pcRef.current && socketRef.current)) {
                return;
            }
            stream.getTracks().forEach((track) => {
                if (!pcRef.current) {
                    return;
                }
                pcRef.current.addTrack(track, stream);
            });

            pcRef.current.onicecandidate = (e) => {
                if (e.candidate) {
                    if (!socketRef.current) {
                        return;
                    }
                    console.log("recv candidate");
                    socketRef.current.emit("candidate", e.candidate, roomID);
                }
            };


        } catch (e) {
            console.error(e);
        }
    };

    const createOffer = async () => {
        console.log("create Offer");
        if (!(pcRef.current && socketRef.current)) {
            return;
        }
        try {
            const sdp = await pcRef.current.createOffer();
            pcRef.current.setLocalDescription(sdp);
            console.log("sent the offer");
            socketRef.current.emit("offer", sdp, roomID);
        } catch (e) {
            console.error(e);
        }
    };


    const createAnswer = async (sdp) => {
        console.log("createAnswer");
        if (!(pcRef.current && socketRef.current)) {
            return;
        }

        try {
            pcRef.current.setRemoteDescription(sdp);
            const answerSdp = await pcRef.current.createAnswer();
            pcRef.current.setLocalDescription(answerSdp);

            console.log("sent the answer");
            socketRef.current.emit("answer", answerSdp, roomName);
        } catch (e) {
            console.error(e);
        }
    };


    useEffect(() => {
        socketRef.current = io("localhost:9000");

        pcRef.current = new RTCPeerConnection({
            iceServers: [
                {
                    urls: "stun:stun.l.google.com:19302",
                },
            ],
        });

        socketRef.current.on("all_users", (allUsers) => {
            if (allUsers.length > 0) {
                createOffer();
            }
        });

        socketRef.current.on("getOffer", (sdp) => {
            console.log("recv Offer");
            createAnswer(sdp);
        });

        socketRef.current.on("getAnswer", (sdp) => {
            console.log("recv Answer")
            if (!pcRef.current) {
                return;
            }
            pcRef.current.setRemoteDescription(sdp);
        });

        socketRef.current.on("getCandidate", async (candidate) => {
            if (!pcRef.current) {
                return;
            }

            await pcRef.current.addIceCandidate(candidate);
        });

        socketRef.current.emit("join_room", {
            room: roomID,
        });

        getMedia();
    }, []);


    useLayoutEffect(() => {
        axiosInstance.get(`/room/room-by-hash/${roomID}`).then((res)=> {
            if (res.status == 200){
                 if ( res.data.room.owner_name == user.userName) {
                     setIsOwner(true)
                 }
                setUser(user.userName)
            }else {
                alert("잘못된 Room 정보")
            }
            setInitLoading(false)
        })

    }, [])

    if (initLoading) {
        return (
            <div>Loading</div>
        )
    }

    console.log(myVideoRef.current)
    return (
        <div>
            <div className="room-video-wrapper">
                <div style={{
                    width : 500,
                    height : 500
                }}>
                    <video
                        id="remotevideo"
                        style={{
                            width: "100%",
                            height: "100%",
                            backgroundColor: "black",
                        }}
                        ref={myVideoRef}
                        autoPlay
                        muted={true}
                    />
                </div>
                <div style={{
                    margin : "10px",
                    border: "1px solid black",
                    width :"400px"
                }}>
                    name
                </div>
            </div>

        </div>

    );
};

export default Room;
