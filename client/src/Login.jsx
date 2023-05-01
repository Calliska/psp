import React, {useState} from "react";
import MyInput from "./components/input/MyInput";
import MyButton from "./components/button/MyButton";
import axios from "axios";

function Login(){
    const [Token, setToken] = useState('')
    const handleSubmit = (event) => {
        event.preventDefault();
        const formData = new FormData(event.target);
        axios.post('http://post.nel-it.ru:42113/login', formData)
            .then(res => setToken(res.data.access_token))
            .catch(err => console.log(err))
        // console.log(Token);
        localStorage.setItem('Token', Token);
        console.log(localStorage.getItem('Token'));

    };
    return (
            <form onSubmit={handleSubmit}>
                <div> Login in your account</div>
                <MyInput type="text" name="email" />
                <MyInput type="text" name="password" />
                <MyButton type="submit">Submit</MyButton>
            </form>
    );
}
export default Login;
