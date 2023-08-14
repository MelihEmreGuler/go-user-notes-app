import {toast} from "react-toastify";
import {useEffect, useState} from "react";
import axios from "axios";
import {Link, useNavigate} from "react-router-dom";
import {RiMailLine} from "@react-icons/all-files/ri/RiMailLine.js";
import {RiKey2Line} from "@react-icons/all-files/ri/RiKey2Line.js";
import Cookies from 'js-cookie';


export default function Login(){
    const navigate = useNavigate();
    const [username_or_email, setUsername_or_email] = useState('');
    const [password, setPassword] = useState('');

    useEffect(() => {
        const hasSessionIdCookie = document.cookie.includes("session_id=");
        hasSessionIdCookie && navigate("/")
    }, [navigate]);



    const submitHandle = (e) => {
        e.preventDefault()

        axios.post('http://127.0.0.1:8080/login', {
            username_or_email,
            password,
            }, {
            withCredentials: true,
            headers: {
                'Content-Type': 'application/json',
            }
        })
            .then(response => {
                Cookies.set('session_id', response.data.session_id, { expires: 1 });

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
        <main className="bg-gray-50 h-screen w-screen flex justify-center items-center">
            <div className="bg-white shadow-lg p-4 rounded-lg grid grid-cols-3 gap-4 w-1/2 lg:w-2/3 lg:h-1/2">

                <div className="hidden lg:flex bg-blue-600 rounded-l-lg text-white p-2 py-10  flex-col gap-4">
                    <h2 className="font-bold text-2xl">Join the Note Taking App</h2>
                    <div>
                        <strong>
                            Discover the world of efficiency and organization with the Note Taking App!
                        </strong>
                        <p>
                            This innovative application is designed to simplify your life by providing a seamless note-taking experience, empowering you to capture your ideas effortlessly, wherever you go.
                        </p>
                    </div>
                </div>

                <form className="flex flex-col justify-center items-center gap-6 col-span-3 lg:col-span-2">
                    <div className="w-full">
                        <span>Username or Email</span>
                        <label htmlFor="email" className="bg-gray-100 flex justify-center items-center gap-1 p-2 rounded-md w-full">
                            <RiMailLine size={20}/>
                            <input id="email" onChange={({ target }) => setUsername_or_email(target.value)} type="email" className="appearance-none bg-transparent w-full"/>
                        </label>
                    </div>

                    <div className="w-full">
                        <span>Password</span>
                        <label htmlFor="password" className="bg-gray-100 flex justify-center items-center gap-1 p-2 rounded-md w-full">
                            <RiKey2Line size={20}/>
                            <input id="password" onChange={({ target }) => setPassword(target.value)} type="password" className="appearance-none bg-transparent w-full"/>
                        </label>
                    </div>

                    <button type="submit" onClick={submitHandle} className="bg-blue-600 text-white p-2 w-full rounded-md hover:bg-blue-500 transition-colors duration-300">
                        Login
                    </button>

                    <div className="flex gap-1.5 w-full">
                        <span className="text-gray-500">New here?</span>
                        <Link to="/signup" className="text-blue-500 hover:scale-105 transition-transform duration-300">
                            Sign up now
                        </Link>
                    </div>

                </form>

            </div>
        </main>
    )
}