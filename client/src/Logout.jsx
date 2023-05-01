import React from "react";
import MyButton from "./components/button/MyButton";
function Logout(){
    return (
        <MyButton onClick={e => localStorage.setItem('Token', '')}>
            Log out
        </MyButton>

    );
}
export default Logout;