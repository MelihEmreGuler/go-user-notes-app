import axios from "axios";
import {toast} from "react-toastify";
import {Link, NavLink, useLocation, useNavigate, useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import Cookies from 'js-cookie';
import {
    RiArchiveDrawerFill, RiArrowLeftLine, RiArrowRightLine,
    RiBook3Line,
    RiFileAddLine,
    RiHome2Line,
    RiLogoutBoxLine, RiMenuLine,
    RiSettings3Line,
    RiTestTubeLine
} from "react-icons/ri";
import classNames from "classnames";
import {BiMenu, BiX} from "react-icons/bi";




// eslint-disable-next-line react/prop-types
export default function Aside({ offcanvas, setOffcanvas, aside, setAside }){

    const navigate = useNavigate();
    const session_id = (Cookies.get('session_id'));
    const [notes, setNotes] = useState([]);
    const [modal, setModal] = useState(false);
    const { id } = useParams() ?? {id:false};
    const path  = useLocation().pathname;
    const [countNote, setCountNote] = useState(0);

    useEffect(() => {
        sessionStorage.setItem("aside", aside)
    }, [aside]);




    useEffect(() => {
        axios.get("http://127.0.0.1:8080/notes",{
            headers: {
               session_id : session_id
           }
       })
           .then(response =>{
               const sortedNotes = response.data.notes.sort((a, b) => {
                   return new Date(b.created_at) - new Date(a.created_at);
               });
               setNotes(sortedNotes);
               setCountNote(sortedNotes.length)

           })
           .catch(error => {
               console.log(error.response)
               if (error.response) {
                   toast.error(error.response.data);
               }
           });

    }, [navigate,session_id]);



    const logoutHandle = (e) => {
        e.preventDefault()
        axios.post("http://127.0.0.1:8080/logout",{
            session_id
        })
            .then(response =>{
                if (response.data.success === true){
                    Cookies.remove('session_id');
                    localStorage.removeItem("user")
                    sessionStorage.removeItem("aside")
                    toast.success(response.data.message)
                    navigate("/login")
                }
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


    const asideHandle = (e) => {
      e.preventDefault()
        if (aside === "true")
            setAside("false")
        else if (aside === "false")
            setAside("true")
    }



    return(
        <aside className={classNames({
            "py-4 lg:h-screen lg:py-8 px-5 lg:px-2 mb-3 lg:mb-0 gap-6 lg:border-r relative flex justify-center items-center lg:justify-start  lg:flex-col bg-secondary  border-primary-100 z-10": true,
            "col-span-1 lg:justify-center": aside=== "true"
        })}>

            <button onClick={asideHandle} className={classNames({
                "absolute p-1 border rounded-full hidden lg:block right-1 bottom-1": true,

            })}>
                {aside === "false"
                    ? <RiArrowLeftLine size={16}/>
                    : <RiArrowRightLine size={16}/>
                }

            </button>

            <div className={classNames({
                "flex justify-between items-center w-4/5 gap-3 p-1": true,
                "flex-col": aside=== "true"
            })}>

                <Link to="/" className={classNames({
                    "p-2 text-center bg-emerald-500 rounded-full text-white text-lg hover:scale-125 transition-transform duration-300": true,
                    "scale-125": path === "/"
                })}>
                    <RiHome2Line size={20}/>
                </Link>

                <Link to="/create" className={classNames({
                    "p-2 text-center bg-blue-500 rounded-full text-white text-lg hover:scale-125 transition-transform duration-300": true,
                    "scale-125": path === "/create"
                })}>
                    <RiFileAddLine size={20}/>
                </Link>

                <Link to="/settings" className={classNames({
                    "p-2 text-center bg-gray-500 rounded-full text-white text-lg hover:scale-125 transition-transform duration-300": true,
                    "scale-125": path === "/settings"
                })}>
                    <RiSettings3Line size={20}/>
                </Link>

                <button onClick={()=>{setModal(!modal)}} className="p-2 text-center bg-red-500 rounded-full left-4 text-white text-lg hover:scale-125 transition-transform duration-300">
                    <RiLogoutBoxLine size={20}/>
                </button>

                <Link to="/test" className={classNames({
                    "p-2 text-center bg-gray-500 rounded-full text-white text-lg hover:scale-125 transition-transform duration-300 hidden lg:block": true,
                    "scale-125": path === "/test"
                })}>
                    <RiTestTubeLine size={20}/>
                </Link>

                <button onClick={()=>{setOffcanvas(!offcanvas)}}  className={classNames({
                    "p-2 text-center bg-black rounded-full text-white lg:hidden text-lg hover:scale-125 transition-transform duration-300": true,
                })}>
                    { offcanvas === true
                        ? <BiX size="20"/>
                        : <BiMenu size="20"/>
                    }
                </button>
            </div>



            {aside=== "false" &&
                <nav className="flex-col gap-1 hidden lg:flex p-1 h-full overflow-y-auto w-full relative">
                <span className="flex justify-center items-center gap-1">
                    <RiBook3Line size={20}/>
                    {countNote}
                </span>

                    {countNote === 0 &&
                        <p className="absolute top-12 ">You don't have any notes! <Link to="/create" className="text-blue-500">Click here</Link> for create first note.</p>
                    }

                    {notes && notes.map((note)=>
                        <Link to={`/note/${note.note_id}`} key={note.note_id} className={classNames({
                            "text-black flex flex-col gap-1 border-b pb-2 p-1 rounded-lg hover:bg-gray-100 transition-colors duration-300" : true,
                            "bg-gray-300 bg-opacity-80 hover:bg-gray-300 shadow-md": id === note.note_id,
                        })}>
                            <span className="font-semibold w-full">{note.title}</span>
                            <small> {note.content.substring(0,50)}...</small>
                            <small className="text-end w-full">{dateFormat(note.created_at)}</small>
                        </Link>
                    )}
                </nav>
            }


            {offcanvas === true && aside=== "false" &&
                <div className="absolute z-30 h-screen w-screen right-0 lg:hidden top-20  grid grid-cols-2">
                    <div onClick={()=>{setOffcanvas(!offcanvas)}} className="bg-transparent">

                    </div>
                    <nav className="flex flex-col gap-4 p-1 bg-gray-100 bg-opacity-95 overflow-y-auto rounded-2xl relative">
                        <span className="flex justify-center items-center gap-1">
                            <RiBook3Line size={20}/>
                            {countNote}
                        </span>

                        {countNote === 0 &&
                            <p className="absolute top-12 ">You don't have any notes! <Link to="/create" className="text-blue-500">Click here</Link> for create first note.</p>
                        }

                        {notes && notes.map((note)=>
                            <NavLink to={`/note/${note.note_id}`} key={note.note_id} className={classNames({
                                "text-black flex flex-col gap-1 border pb-2 p-1 rounded-lg hover:bg-gray-100 transition-colors duration-300" : true,
                                "bg-gray-300 bg-opacity-80 hover:bg-gray-300 shadow-md": id === note.note_id,
                            })}>
                                <span className="font-semibold w-full">{note.title}</span>
                                <small> {note.content.substring(0,50)}...</small>
                                <small className="text-end w-full">{dateFormat(note.created_at)}</small>
                            </NavLink>
                        )}
                    </nav>
                </div>

            }

            {modal &&
                <div className="absolute h-screen w-screen top-0 left-0 bg-black bg-opacity-60 flex items-start justify-center">
                    <div className="bg-white p-6 flex flex-col mt-20 gap-8 rounded-md">
                        <span className="font-semibold text-xl">
                            Are you sure you want to log out?
                        </span>
                        <div className="flex justify-center items-center gap-10">
                            <button onClick={()=>{setModal(false)}} className="border p-2 bg-blue-50 border-blue-500 text-blue-600 rounded-md w-1/4">
                                Cancel
                            </button>
                            <button onClick={logoutHandle} className="border p-2 bg-gray-50 border-gray-500 text-gray-600 rounded-md w-1/4">
                                Yes
                            </button>
                        </div>
                    </div>
                </div>
            }

        </aside>
    )
}