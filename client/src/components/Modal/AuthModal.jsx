import NiceModal, { useModal } from '@ebay/nice-modal-react';
import { Modal } from 'antd';
import React from "react";

import axios from "axios"

import "./AuthModal.css"

export default NiceModal.create(() => {
    const modal = useModal();

    const handleGoogleLogin  = async () => {

        const res = await axios.get("http://localhost:8000/login")
        console.log((res.data))

        window.open(res.data)

        // modal.hide()
    }

    return (
        <Modal
            title="Auth Login"
            onOk={() => modal.hide()}
            visible={modal.visible}
            onCancel={() => modal.hide()}
            afterClose={() => modal.remove()}
            footer={null}
        >
            <div style={{
                padding : "20px",
                display : "flex",
                justifyContent : "center",
                alignContent : "center"
            }}>
                <button className="google_login_button" onClick={handleGoogleLogin}>Google</button>
            </div>
        </Modal>
    );
});