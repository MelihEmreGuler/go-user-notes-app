
import { useState, useEffect } from "react";
import classNames from "classnames";
import Aside from "../components/Aside.jsx";



export default function Layout({ children }){
    const [offcanvas, setOffcanvas] = useState(true);
    const [aside, setAside] = useState(sessionStorage.getItem("aside") !== null ? sessionStorage.getItem("aside"): "false")



    useEffect(() => {
        window.addEventListener('resize', () => {
            if (window.innerWidth > 1023){
                setOffcanvas(0)
            }
        });
    }, []);

    return(
        <div className={classNames({
            "lg:grid  min-h-screen": true,
            "overflow-hidden": offcanvas,
            "grid-cols-12": aside === "true",
            "grid-cols-5": aside === "false"
        })}>
            <Aside offcanvas={offcanvas} setOffcanvas={setOffcanvas} aside={aside} setAside={setAside}/>
            <main className={classNames({
                " p-4 lg:p-8 flex flex-col gap-6 lg:gap-8": true,
                "col-span-11": aside === "true",
                "col-span-4": aside === "false"

            })}>
                { children }
            </main>
        </div>
    )
}