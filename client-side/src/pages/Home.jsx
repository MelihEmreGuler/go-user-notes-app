import Layout from "../layouts/Layout.jsx";
import axios from "axios";
import {useEffect, useState} from "react";
import Cookies from "js-cookie";
import {toast} from "react-toastify";
import {Link, useLocation, useNavigate, useParams} from "react-router-dom";
import {RiArrowLeftLine, RiEdit2Fill} from "react-icons/ri";
import classNames from "classnames";


export default function Home(){

    const session_id = (Cookies.get('session_id'));
    const user = localStorage.getItem("user");
    const [note, setNote] = useState({});
    const navigate = useNavigate();
    const path  = useLocation().pathname;


    useEffect(() => {
        const url = `http://localhost:8080/notes`
        axios.get(url,{
            headers: {
                session_id : session_id
            }})
            .then((response)=>{
                const sortedNotes = response.data.notes.sort((a, b) => {
                    return new Date(b.created_at) - new Date(a.created_at);
                });
                sortedNotes && setNote(sortedNotes[0]);
            })
            .catch(error => {
                console.log(error.response)
                if (error.response) {
                    toast.error(error.response.data);
            }});
    }, [navigate]);

    const dateFormat = (dateStr) => {
        const date = new Date(dateStr);
        const day = String(date.getDate()).padStart(2, '0');
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const year = date.getFullYear();
        const hour = String(date.getHours()).padStart(2, '0');
        const minutes = String(date.getMinutes()).padStart(2, '0');
        return `${day}/${month}/${year} ${hour}:${minutes}`;
    }


    return(
        <Layout>
            <main className=" w-full">
                <header className="flex justify-between pb-4 pt-2 border-b">
                    <span className="text-xl font-semibold">
                        Hi {user}! Welcome. Below is the last note you created.
                    </span>
                </header>
                <section className=" pt-10 flex flex-col gap-10">
                    <h1 className="font-bold text-3xl">
                        {note && note.title}
                    </h1>
                    <textarea name="" id="" rows={25}
                              defaultValue={ note ? note.content : ""}
                              readOnly
                              className={classNames({
                                  "p-2 bg-stone-50": true,
                              })}
                    ></textarea>
                </section>
            </main>
        </Layout>
    )
}