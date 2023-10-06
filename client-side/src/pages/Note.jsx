import Layout from "../layouts/Layout.jsx";
import axios from "axios";
import {useEffect, useState} from "react";
import Cookies from "js-cookie";
import {toast} from "react-toastify";
import {Link, useNavigate, useParams} from "react-router-dom";
import {RiArrowLeftLine, RiCloseLine, RiEdit2Fill} from "react-icons/ri";
import classNames from "classnames";


export default function Note(){
    const session_id = (Cookies.get('session_id'));
    const [note, setNote] = useState({});
    const [cancelModal, setCancelModal] = useState(false);
    const [deleteModal, setDeleteModal] = useState(false);
    const navigate = useNavigate();
    const {id,action} = useParams() ?? {id:false, action:false};

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

    const deleteHandle = (e) => {
        e.preventDefault()
        const url = `http://localhost:8080/note/:id?id=${id}`
        axios.delete(url,{
            headers: {
                session_id : session_id
            }})
            .then((response)=>{
                if (response.data.success){
                    toast.success("Note deleted successfully")
                    setDeleteModal(false)
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
            <main  className=" w-full">
                <header className="flex justify-between pb-4 pt-2 border-b">
                    <Link to="/" className="bg-gray-100 hover:bg-gray-200 transition-colors duration-300 rounded-full p-3 flex items-center justify-center">
                        <RiArrowLeftLine size={24}/>
                    </Link>

                    { action === "edit"
                        ?
                        <div className="flex items-center justify-between lg:w-1/6 gap-3">
                            <button onClick={()=>{setCancelModal(!cancelModal)}} className="bg-red-200 text-red-600 border border-red-600 hover:bg-red-300 transition-colors duration-300 rounded-full p-2 flex items-center justify-center">
                                <RiCloseLine size={28}/>
                            </button>
                            <button onClick={updateHandle} className="bg-emerald-500 hover:bg-emerald-600 transition-colors duration-300 text-white lg:w-2/3 rounded-md p-3 font-semibold text-xl">
                                Save
                            </button>
                        </div>
                        :
                        <div className="flex items-center justify-between gap-3">
                            <button onClick={()=>{setDeleteModal(!deleteModal)}} className="p-2 border border-red-600 bg-red-200 text-red-600 rounded-md">
                                Delete
                            </button>
                            <Link to={`/note/${id}/edit`} className="bg-emerald-500 hover:bg-emerald-600 transition-all duration-300 text-white rounded-full p-3">
                                <RiEdit2Fill size={28}/>
                            </Link>
                        </div>
                    }

                </header>
                {note && (
                    <section className=" pt-10 flex flex-col gap-6 ">
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
                                      "p-2 bg-stone-50": true,
                                      "appearance-none bg-gray-100 bg-opacity-40 border": action ===   "edit"
                                  })}
                        ></textarea>
                    </section>
                )}

                {cancelModal &&
                    <div className="absolute h-screen w-screen top-0 left-0 bg-black bg-opacity-60 flex items-start justify-center z-50">
                        <div className="bg-white p-6 flex flex-col mt-20 gap-8 rounded-md">
                        <span className="font-semibold text-xl">
                            Are you sure you want to cancel changes?
                        </span>
                            <div className="flex justify-center items-center gap-10">
                                <button onClick={()=>{setCancelModal(false)}} className="border p-2 bg-blue-50 border-blue-500 text-blue-600 rounded-md w-1/4">
                                    Cancel
                                </button>
                                <Link to={`/note/${id}`} onClick={()=>{setCancelModal(false)}} className="border p-2 bg-gray-50 border-gray-500 text-gray-600 text-center rounded-md w-1/4">
                                    Yes
                                </Link>
                            </div>
                        </div>
                    </div>
                }

                {deleteModal &&
                    <div className="absolute h-screen w-screen top-0 left-0 bg-black bg-opacity-60 flex items-start justify-center z-50">
                        <div className="bg-white p-6 flex flex-col mt-20 gap-8 rounded-md">
                        <span className="font-semibold text-xl">
                            Are you sure you want to delete this note?
                        </span>
                            <div className="flex justify-center items-center gap-10">
                                <button onClick={()=>{setDeleteModal(false)}} className="border p-2 bg-blue-50 border-blue-500 text-blue-600 rounded-md w-1/4">
                                    Cancel
                                </button>
                                <button onClick={deleteHandle} className="border p-2 bg-gray-50 border-gray-500 text-gray-600 text-center rounded-md w-1/4">
                                    Yes
                                </button>
                            </div>
                        </div>
                    </div>
                }
            </main>

        </Layout>
    )
}