import Layout from "../layouts/Layout.jsx";
import axios from "axios";
import {useEffect, useState} from "react";
import Cookies from "js-cookie";
import {toast} from "react-toastify";
import {Link, useLocation, useNavigate, useParams} from "react-router-dom";
import {RiArrowLeftLine, RiEdit2Fill} from "react-icons/ri";
import classNames from "classnames";


export default function Home(){

    const [session_id, setSession_id] = useState(Cookies.get('session_id') ? Cookies.get('session_id') : "");
    const [note, setNote] = useState({});
    const [countNote, setCountNote] = useState(0);
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
                setCountNote(sortedNotes.length)
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
            <main  className="h-screen w-full">
                <header className="flex justify-between pb-4 pt-8 border-b">
                    <div className="flex justify-center items-center gap-3">
                        <Link to="/" className="bg-gray-100 hover:bg-gray-200 transition-colors duration-300 rounded-full p-3 flex items-center justify-center">
                            <RiArrowLeftLine size={24}/>
                        </Link>
                        <small className="font-semibold">
                            Created at: {dateFormat(note.created_at)}
                        </small>
                    </div>

                    <div className="p-2 border-2 border-black rounded-2xl flex justify-center items-center ">
                         Notes Count: { countNote&&countNote}
                    </div>

                    <Link to={`/note/${note.note_id}/edit`} className={classNames({
                            "bg-emerald-500 hover:bg-emerald-600 transition-all duration-300 text-white rounded-full p-3": true,
                            "bg-blue-500": path === "/edit"
                    })}>
                        <RiEdit2Fill size={24}/>
                    </Link>
                </header>
                <section className="h-screen pt-10 flex flex-col gap-10">
                    <h1 className="font-bold text-3xl">
                        {note && note.title}
                    </h1>
                    <div>
                        {note && note.content}
                    </div>
                </section>
            </main>
        </Layout>
    )
}