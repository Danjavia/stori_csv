import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";
import { KindeProvider } from "@kinde-oss/kinde-auth-react";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement,
);
root.render(
  <React.StrictMode>
    <KindeProvider
      clientId="5d88970b3ecc4f119ea3649941492f18"
      domain="https://csvtest.kinde.com"
      logoutUri={window.location.origin}
      redirectUri={window.location.origin}
    >
      <App />
    </KindeProvider>
  </React.StrictMode>,
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
