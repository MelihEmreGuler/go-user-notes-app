import axios from "axios";
import {toast} from "react-toastify";
import {Link, useNavigate} from "react-router-dom";
import {useLayoutEffect, useState} from "react";
import Cookies from 'js-cookie';
import {RiArrowLeftCircleFill, RiArrowLeftLine} from "react-icons/ri";
import Layout from "../layouts/Layout.jsx";



export default function Create(){
    const navigate = useNavigate();
    const session_id = (Cookies.get('session_id'));
    const [title, setTitle] = useState('');
    const [content, setContent] = useState('');

    const createHandle = (e) => {
        e.preventDefault()

        if (title === "" || content === "")
            toast.warning("Please complete all fields!")
        else
            axios.post("http://127.0.0.1:8080/note",{
                title,
                content,
                session_id
            }).then(response =>{
                if (response.data.success === true){
                    toast.success(response.data.message)
                    navigate("/")
                }
            })
                .catch(error => {
                    console.log(error.response)
                    if (error.response) {
                        toast.error(error.response.data);
                    }
                });
    }


    return(
        <Layout>
            <main>
                <header className="flex justify-between pb-4 pt-2 border-b">
                    <Link to="/" className="bg-gray-100 hover:bg-gray-200 transition-colors duration-300 rounded-full p-3 flex items-center justify-center">
                        <RiArrowLeftLine size={24}/>
                    </Link>
                    <span className="font-semibold text-2xl">Create New Note!</span>
                    <button type="submit" onClick={createHandle} className="bg-emerald-500 hover:bg-emerald-600 transition-colors duration-300 text-white rounded-md p-2 w-1/6 font-semibold text-xl">
                        Save
                    </button>
                </header>

                <form action="" className="flex flex-col justify-start p-10 items-center gap-10">
                    <label htmlFor="title" className="w-full flex flex-col gap-2">
                        <span className="font-semibold text-2xl">Title</span>
                        <input required id="title" type="text" onChange={({ target }) => setTitle(target.value)} className="appearance-none p-3 border focus:shadow-md w-full rounded-md"/>
                    </label>
                    <label htmlFor="content" className="w-full flex flex-col gap-2">
                        <span className="font-semibold text-2xl">Content</span>
                        <textarea required id="content" rows={18} onChange={({ target }) => setContent(target.value)} className="appearance-none p-3 border focus:shadow-md w-full rounded-md"/>
                    </label>
                </form>
            </main>
        </Layout>

    )
}