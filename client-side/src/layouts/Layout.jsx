
import { useState, useEffect } from "react";
import classNames from "classnames";
import Aside from "../components/Aside.jsx";
import axios from "axios";
import {Navigate, useNavigate} from "react-router-dom";


export default function Layout({ children }){
    const [offcanvas, setOffcanvas] = useState(0);

    useEffect(() => {
        window.addEventListener('resize', () => {
            if (window.innerWidth > 1023){
                setOffcanvas(0)
            }
        });
    }, []);

    return(
        <div className={classNames({
            "lg:grid grid-cols-5 min-h-screen": true,
            "overflow-hidden": offcanvas
        })}>
            <Aside offcanvas={offcanvas} setOffcanvas={setOffcanvas}/>
            <main className="col-span-4 p-4 lg:p-8 flex flex-col gap-6 lg:gap-8">
                { children }
            </main>
        </div>
    )
}