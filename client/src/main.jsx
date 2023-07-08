import React from "react";
import ReactDOM from "react-dom";

import Auth, {GlobalAuth} from "./context/user";
import App from "./App"
import NiceModal from "@ebay/nice-modal-react";
import { CookiesProvider } from "react-cookie";

ReactDOM.render(
    <React.StrictMode>
        <GlobalAuth.Provider value={new Auth()}>
            <NiceModal.Provider>
                <CookiesProvider>
                    <App />
                </CookiesProvider>
            </NiceModal.Provider>
        </GlobalAuth.Provider >
    </React.StrictMode>,
    document.getElementById("root")
);
