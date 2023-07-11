import React from "react";
import "./NavBar.css"
import {useModal} from "@ebay/nice-modal-react";
import CreateRoomModal from "../Modal/CreateRoomModal";
import Success from "../Tooltip/Success";

const NavBar = (props) => {
    const createModal = useModal(CreateRoomModal)
    const CreateRoomHandler = () =>{
        createModal.show()
    }

    return (
        <div className="nav-bar-wrapper">
            <div className="nav-bar-box">
                <span>MyDetail</span>
                <span onClick={CreateRoomHandler}>Create Room</span>
            </div>
        </div>
    );
};

export default NavBar;
