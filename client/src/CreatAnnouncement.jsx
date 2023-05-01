import React, {useState} from "react";
import axios from "axios";
import MyButton from "./components/button/MyButton";
import MyInput from "./components/input/MyInput";
function CreatAnnouncement(){
    const [announcement, setAnnouncement] = useState([]);
    const [groupID, setgroupID] = useState(0);
    const handleSubmit = (event) => {
        event.preventDefault();
        const formData = new FormData(event.target);
        formData.append('groupId', groupID);
        axios.post('http://post.nel-it.ru:42113/announcement/create/' + localStorage.getItem('Token'), formData)
            .then(res => console.log(res))
            .catch(err => console.log(err))
    };
    function GroupSelect (props) {
        axios.get( 'http://post.nel-it.ru:42113/groups/get/forUser/' + localStorage.getItem('Token'))
            .then(res =>  setAnnouncement(res.data))
            .catch(err => console.log(err))
        return (
            <div>
                {props.children}
            <select onChange={e => setgroupID(e.target.value)}>
                {announcement.map((announcement)=><option value={announcement.id} key={announcement.id}>{announcement.name}</option>)}
            </select>
            </div>
        );
    }


    return (
        <form onSubmit={handleSubmit}>
            <h1> Creat new announcement</h1>
            <MyInput type="text" name="text" />
            <GroupSelect>
                Select
            </GroupSelect>
            <MyButton type="submit">Submit</MyButton>
        </form>
    );
}
export default CreatAnnouncement;