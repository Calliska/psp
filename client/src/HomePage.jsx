import React from "react";
function HomePage (){
    function CheckToken(){
        return (
          <div>
              {localStorage.getItem('Token')}
          </div>
        );
    }
    return(
        <>
            <a href="/creat_group">
                Group creation
            </a>
            <h1>
                its your home page
            </h1>
            <CheckToken>

            </CheckToken>

        </>
    );
}
export default HomePage;