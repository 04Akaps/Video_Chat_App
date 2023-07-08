import React from "react";
import "./NavBar.css"

const NavBar = (props) => {
    return (
        <div className="nav-bar-wrapper">
            <div className="nav-bar-box">
                <span>MyDetail</span>
                <span>Create Room</span>
            </div>
        </div>
    );
};

export default NavBar;
