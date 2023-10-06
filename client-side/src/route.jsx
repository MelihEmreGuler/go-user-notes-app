import {BrowserRouter, Routes, Route, useNavigate, Navigate} from "react-router-dom";
import Home from "./pages/Home.jsx";
import Login from "./pages/Login.jsx";
import Signup from "./pages/Signup.jsx";
import Create from "./pages/Create.jsx";
import Note from "./pages/Note.jsx";
import Settings from "./pages/Settings.jsx";
import Test from "./pages/Test.jsx";


export default function  route(){

    const hasSessionIdCookie = document.cookie.includes("session_id=");


    return(
        <BrowserRouter>
            <Routes>
                <Route path="/" element={hasSessionIdCookie ? <Home /> : <Navigate to="/login" />} />
                <Route path="/login" element={<Login/>} />
                <Route path="/signup" element={hasSessionIdCookie ? <Signup /> : <Navigate to="/login" />} />
                <Route path="/create" element={hasSessionIdCookie ? <Create/> : <Navigate to="/login" />} />
                <Route path="/settings" element={hasSessionIdCookie ? <Settings/> : <Navigate to="/login" />} />
                <Route path="/note/:id" element={hasSessionIdCookie ? <Note /> : <Navigate to="/login" />} />
                <Route path="/note/:id/:action" element={hasSessionIdCookie ? <Note /> : <Navigate to="/login" />} />
                <Route path="/test" element={hasSessionIdCookie ? <Test /> : <Navigate to="/login" />} />
                <Route path='*' element={<Login />}/>
            </Routes>
        </BrowserRouter>
    )
}