import React from "react";

class Auth {
  userName;
  updateCallback = null;

  setUserName = (userName) => {
    this.userName = userName;
    if (this.updateCallback) {
      this.updateCallback();
    }
  };

  setUpdateCallback = (callback) => {
    this.updateCallback = callback;
  };
}

export const GlobalAuth = React.createContext(new Auth());

export const useAuth = () => {
  return React.useContext(GlobalAuth);
};

export default Auth;