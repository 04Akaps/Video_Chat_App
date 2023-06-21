import React from "react";
import ReactDOM from "react-dom";

import App from "./App"
import NiceModal from "@ebay/nice-modal-react";

ReactDOM.render(
    <React.StrictMode>
        <NiceModal.Provider>
            <App />
        </NiceModal.Provider>
    </React.StrictMode>,
    document.getElementById("root")
);
