import Layout from "../layouts/Layout.jsx";
import axios from "axios";
import {useEffect, useState} from "react";
import Cookies from "js-cookie";
import {toast} from "react-toastify";
import {Link, useLocation, useNavigate, useParams} from "react-router-dom";
import {RiArrowLeftLine, RiCloseCircleFill, RiCloseLine, RiEdit2Fill} from "react-icons/ri";
import classNames from "classnames";

export default function Note(){
    const [session_id,  setSession_id] = useState(Cookies.get('session_id') ? Cookies.get('session_id') : "");
    const [note, setNote] = useState({});
    const navigate = useNavigate();
    const {id,action} = useParams() ?? {id:false, action:false};
    const path  = useLocation().pathname;

    const [date, setDate] = useState("");
    const [title, setTitle] = useState("");
    const [content, setContent] = useState("");

    useEffect(() => {
        if (id){
            const url = `http://localhost:8080/note/:id?id=${id}`
            axios.get(url,{
                headers: {
                    session_id : session_id
                }})
                .then((response)=>{
                    setDate(response.data.note.created_at)
                    setTitle(response.data.note.title)
                    setContent(response.data.note.content)
                })
                .catch(error => {
                    console.log(error.response)
                    if (error.response) {
                        toast.error(error.response.data);
                    }
                });
        }
    }, [navigate,session_id,id]);

    const updateHandle = (e) => {
      e.preventDefault()
        const url = `http://localhost:8080/note/:id?id=${id}`
        axios.put(url,{
            title,
            content,
            session_id
        })
            .then((response)=>{
                if (response.data.success)
                    toast.success("Note updated successfully")
                navigate(`/note/${id}`)
            })
            .catch(error => {
                console.log(error.response)
                if (error.response) {
                    toast.error(error.response.data);
                }
            });
    }

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
                            Created at: {date && dateFormat(date)}
                        </small>
                    </div>

                    { action === "edit"
                        ?
                        <div className="flex items-center justify-between lg:w-1/6 gap-3">
                            <Link to={`/note/${id}`}>
                                <RiCloseCircleFill className="text-red-500" size={36}/>
                            </Link>
                            <button onClick={updateHandle} className="bg-emerald-500 hover:bg-emerald-600 transition-colors duration-300 text-white lg:w-2/3 rounded-md p-2 font-semibold text-xl">
                                Save
                            </button>
                        </div>
                        :
                        <Link to={`/note/${id}/edit`} className="bg-emerald-500 hover:bg-emerald-600 transition-all duration-300 text-white rounded-full p-3">
                            <RiEdit2Fill size={24}/>
                        </Link>
                    }

                </header>
                {note && (
                    <section className="h-screen pt-10 flex flex-col gap-10">
                        <input type="text"
                               defaultValue={title}
                               onChange={({target})=>{setTitle(target.value)}}
                               readOnly={action!=="edit"}
                               className={classNames({
                                   "p-2 font-bold text-3xl": true,
                                   "appearance-none bg-gray-100 bg-opacity-40 border": action ===   "edit"
                               })}
                        />
                        <textarea name="" id="" rows={20}
                                  defaultValue={content}
                                  onChange={({target})=>{setContent(target.value)}}
                                  readOnly={action!=="edit"}
                                  className={classNames({
                                      "p-2": true,
                                      "appearance-none bg-gray-100 bg-opacity-40 border": action ===   "edit"
                                  })}
                        ></textarea>
                    </section>
                )}
            </main>

        </Layout>
    )
}