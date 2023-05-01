import React, {useState} from "react";
import axios from "axios";
import creatAnnouncement from "./CreatAnnouncement";
function YourAnnouncement (){
    const [announcement, setannouncement] = useState([]);
    axios.get("http://post.nel-it.ru:42113/announcement/get/:token/:group_id").then(res => setannouncement(res.data))
        .catch(err => console.log(err));
    return(
      <label>

      </label>
    );
}