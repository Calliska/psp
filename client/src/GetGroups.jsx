import React from "react";
import axios from "axios";
import MyButton from "./components/button/MyButton";
function GetGroups(){
    function getting() {
        axios.get('http://post.nel-it.ru:42113/groups/get/forUser/' + localStorage.getItem('Token'))
            .then(res => console.log(res))
            .catch(err => console.log(err))
    }
    return (
        <div>
            help
        </div>

    );
}
// export default GetGroups;