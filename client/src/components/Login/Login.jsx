import React, {useEffect} from "react";

import "./Login.css"
import {useModal} from "@ebay/nice-modal-react";
import AuthModal from "../Modal/AuthModal";

const Login = (props) => {

    const authModal = useModal(AuthModal)

    const modalOpen = () =>{
        authModal.show()
    }

    return (
        <div className="login-wrapper">
           <button className="login-modal-on-button" onClick={modalOpen}>Login</button>
        </div>
    );
};

export default Login;
