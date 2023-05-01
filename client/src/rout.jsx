import React, {useState} from "react";
// import ReactDOM from 'react-dom';
import CreatGroup from "./CreatGroup";
import Login from "./Login";
import Registration from "./Registration";
import NotFoundPage from "./NotFoundPage";
import HomePage from "./HomePage";
import {ErrorContext} from "./Context";
import Logout from "./Logout";
import {
    Routes,
    Route, BrowserRouter
} from 'react-router-dom';
import CreatAnnouncement from "./CreatAnnouncement";
function Rout(){
    const [Error, setError] = useState(0);
            return (
            <ErrorContext.Provider value={{
                Error,
                setError
            }}>
            <header>
                <a href="/"> Login</a>
                <a href="/regestration"> Reg</a>
                <a href="/home_page"> Home</a>
                < a href = "/creat_announcment"> Creat</a>
                <Logout>
                    Log out
                </Logout>
            </header>
            <BrowserRouter>
                <Routes>
                    <Route path = "/" element={<Login />}/>
                    <Route path = "/regestration" element={<Registration />}/>
                    <Route path = "/*" element={<NotFoundPage />}/>
                    <Route path = "/creat_group" element={<CreatGroup />}/>
                    <Route path = "/home_page" element={<HomePage />}/>
                    <Route path = "/creat_announcment" element={<CreatAnnouncement />}/>
                </Routes>
            </BrowserRouter>
            </ErrorContext.Provider>
    );
}
export default Rout;