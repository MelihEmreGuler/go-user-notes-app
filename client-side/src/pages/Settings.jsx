import Layout from "../layouts/Layout.jsx";
import {Link, useNavigate} from "react-router-dom";
import {
    RiArrowLeftLine,
    RiCloseCircleFill, RiCloseLine, RiKey2Fill,
    RiKey2Line,
    RiMailLine,
    RiShieldKeyholeLine,
    RiUser3Line
} from "react-icons/ri";
import {useState} from "react";
import {toast} from "react-toastify";
import axios from "axios";
import classNames from "classnames";
import Cookies from "js-cookie";


export default function Settings(){
    const navigate = useNavigate();
    const [username, setUsername] = useState('');
    const [new_password, setNew_password] = useState('');
    const [old_password, setOld_password] = useState('');
    const [password, setPassword] = useState('');
    const [email, setEmail] = useState('');
    const session_id = (Cookies.get('session_id'));
    const [emailchange, setEmailchange] = useState(false);
    const [passchange, setPasschange] = useState(false);

    const ChangeEmail = (e) => {
        e.preventDefault();

        axios.put('http://127.0.0.1:8080/user/email', {
            password,
            email,
            session_id,
        }, {
            withCredentials: true,
            headers: {
                'Content-Type': 'application/json',
            }
        })
            .then(response => {
                if (response.data.success === true){
                    toast.success(response.data.message)
                }
            })
            .catch(error => {
                console.log(error.response)
                if (error.response) {
                    toast.error(error.response.data);
                }
            })
    }

    const ChangePass = (e) => {
        e.preventDefault();

        axios.put('http://127.0.0.1:8080/user/password', {
            old_password,
            new_password,
            session_id,
        }, {
            withCredentials: true,
            headers: {
                'Content-Type': 'application/json',
            }
        })
            .then(response => {
                if (response.data.success === true){
                    toast.success(response.data.message)
                }
            })
            .catch(error => {
                console.log(error.response)
                if (error.response) {
                    toast.error(error.response);
                }
            })
    }

    return(
        <Layout>
            <header className="flex justify-between pb-4 pt-2 border-b">
                <Link to="/" className="bg-gray-100 hover:bg-gray-200 transition-colors duration-300 rounded-full p-3 flex items-center justify-center">
                    <RiArrowLeftLine size={24}/>
                </Link>
                <span className="font-semibold text-center w-full text-2xl">Settings</span>

            </header>


            <article className="w-full grid lg:grid-cols-2 gap-6">

                <section className="w-full flex flex-col gap-4 items-center p-4 pt-16 relative ">
                    <label htmlFor="email" className="bg-gray-100 flex justify-center items-center gap-1 p-3 rounded-md w-full">
                        <RiMailLine size={20}/>
                        <input placeholder="New Email" id="email" onChange={({ target }) => setEmail(target.value)} type="email" className="appearance-none bg-transparent w-full"/>
                    </label>
                    <label htmlFor="password" className="bg-gray-100 flex justify-center items-center gap-1 p-3 rounded-md w-full">
                        <RiKey2Line size={20}/>
                        <input placeholder="Current Password" id="password" onChange={({ target }) => setPassword(target.value)} type="password" className="appearance-none bg-transparent w-full"/>
                    </label>
                    <button type="submit" onClick={ChangeEmail} className="bg-blue-600 text-white p-2 w-full rounded-md hover:bg-blue-500 transition-colors duration-300 ">
                        Change Email
                    </button>
                </section>

                <section className="w-full flex flex-col gap-4 items-center p-4 pt-16 relative ">
                    <label htmlFor="password" className="bg-gray-100 flex justify-center items-center gap-1 p-3 rounded-md w-full">
                        <RiKey2Line size={20}/>
                        <input placeholder="Current Password" id="password" onChange={({ target }) => setOld_password(target.value)} type="password" className="appearance-none bg-transparent w-full"/>
                    </label>
                    <label htmlFor="cpassword" className="bg-gray-100 flex justify-center items-center gap-1 p-3 rounded-md w-full">
                        <RiKey2Fill size={20}/>
                        <input placeholder="New Password" id="cpassword" onChange={({ target }) => setNew_password(target.value)} type="password" className="appearance-none bg-transparent w-full"/>
                    </label>
                    <button type="submit" onClick={ChangePass}  className="bg-blue-600 text-white p-2 w-full rounded-md hover:bg-blue-500 transition-colors duration-300 ">
                        Change Password
                    </button>
                </section>

            </article>

        </Layout>
    )
}