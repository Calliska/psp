import React, {useContext, useState} from "react";
import axios from "axios";
import {Button, Form} from "react-bootstrap";
import 'bootstrap/dist/css/bootstrap.min.css';
 const CreatGroup = () => {
     const [Error, setError] = useState('');
     const [name, setName] = useState("");
     const handleSubmit = () => {
         axios.post('http://post.nel-it.ru:42113/create/' + localStorage.getItem('Token') + '/' + name)
             .then(res =>
             {
                 console.log(res.data)

             })
             .catch(err => {console.log(err)
             console.log(err.response.data)

             })

     };

     return (
         <Form>
<Form.Group>
    <Form.Label htmlFor="inputPassword5">Group name  </Form.Label>
    <Form.Control
        id="inputPassword5"
        aria-describedby="passwordHelpBlock"
        onChange={e => setName(e.target.value)}
    />
</Form.Group>
             <Button onClick={e => handleSubmit()}>Send</Button>
         </Form>
     );
 }
 export default CreatGroup;