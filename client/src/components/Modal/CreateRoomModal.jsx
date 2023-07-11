import NiceModal, { useModal } from '@ebay/nice-modal-react';
import {Modal, Tooltip} from 'antd';
import React from "react";


import "./CreateRoomModal.css"
import axiosInstance from "../../api/axiosInstance";

export default NiceModal.create(() => {
    const modal = useModal();

    const handleCreateRoom  = async () => {
        axiosInstance.post("/room/create").then((res) => {
            if (res.status == 200) {
                alert("Create New Room Success")
            }else {
                alert("Failed Room Limit")
            }
        });
    }

    return (
        <Modal
            title="Create Room"
            onOk={() => modal.hide()}
            visible={modal.visible}
            onCancel={() => modal.hide()}
            afterClose={() => modal.remove()}
            footer={null}
        >
            <div style={{
                padding : "20px",
                display : "flex",
                flexDirection  : "column",
                justifyContent : "center",
                alignContent : "center",
                alignItems:"center"
            }}>
                <div>방은 최대 3개만 생성 가능 합니다.</div>
               <button className="create_room_button" onClick={handleCreateRoom}>Create New Room</button>
            </div>
        </Modal>
    );
});