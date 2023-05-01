import axios from "axios";
import MyInput from "./components/input/MyInput";
import MyButton from "./components/button/MyButton";
import React from "react";

function Registration() {
    const handleSubmit = (event) => {
        event.preventDefault();
        const formData = new FormData(event.target);
        axios.post('http://post.nel-it.ru:42113/register', formData)
            .then(res => console.log(res))
            .catch(err => console.log(err))
    };

    return (
        <form onSubmit={handleSubmit}>
            <def> Creat new account</def>
            <MyInput type="text" name="email" />
            <MyInput type="text" name="password" />
            <MyInput type="text" name="firstname" />
            <MyInput type="text" name="secondname" />
            <MyButton type="submit">Submit</MyButton>
        </form>
    );
}
export default Registration;