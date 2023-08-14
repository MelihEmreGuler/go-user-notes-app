import React from 'react'
import ReactDOM from 'react-dom/client'
import axios from "axios";

import "./styles/tailwind.css"


import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.min.css';

import Route from "./route.jsx";

axios.defaults.baseURL = import.meta.env.VITE_API_URL

const options = {
    position: "bottom-right",
    autoClose: 5000,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: true,
    draggable: true,
    progress: undefined,
    theme: "colored",
};

ReactDOM.createRoot(document.getElementById('root')).render(
    <React.StrictMode>
            <Route/>
        <ToastContainer {...options} />
    </React.StrictMode>
)