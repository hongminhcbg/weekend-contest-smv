import React, { Component, useState } from 'react';

function Home() {
    const [name, setName] = useState("");
    const [isError, setError] = useState(false);
    function handlerOnclick() {
        fetch('http://localhost:8080/ping', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
              }    
        }).
        then(r => r.json()).
        then((r) => {
            setName(r.message)
        })
        .catch((err) => {
            console.log(err);
            setError(true);
        })
    }

    function getError() {
        if(isError) {
            return "Error"
        }

        return "Success"
    }

    return (
        <div>
            <h1> This is home, hello {name} </h1>
            <h1>{getError()}</h1>
        
            <button onClick={handlerOnclick}>Ping Server</button>
        </div>
    );
}
export default Home;