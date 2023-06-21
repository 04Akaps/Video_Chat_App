import React from "react";

import "./Home.css"
import {useModal} from "@ebay/nice-modal-react";
import AuthModal from "./Modal/AuthModal";

const Home = (props) => {
    const authModal = useModal(AuthModal)

    const modalOpen = () =>{
        authModal.show()
    }

    return (
        <div className="home-wrapper">
           <button className="login-modal-on-button" onClick={modalOpen}>Login</button>
        </div>
    );
};

export default Home;
