import React, { Component, useState } from 'react';
import {useCookies} from 'react-cookie';
import { detect } from 'detect-browser';

const Home = (props) => {
    const [ctx, setCtx] = useState({});
    const [hasContext, setHasContext] = useState(false);
    const [err, setErr] = useState(null);
    const [registered, setRegistered] = useState(false);  
    const [cookies, setCookie, removeCookie] = useCookies(['token']);

    function getMessage() {
        if(err != null) {
            return 'Đã có lỗi lỗi ' + err; 
        }

        if (!hasContext) {
            return 'Đang xử lý yêu cầu'
        }

        return 'Chào mừng ' + ctx.IPv4 + ', có phải bạn đang ở ' + ctx.city;
    }

    if (!hasContext) {
        if (cookies.IPv4 && cookies.IPv4 !== '') {
            setCtx({
                'IPv4': cookies.IPv4,
                'city': cookies.city,
            });
            setHasContext(true);
            console.log('response ip and city from cookie');
            return;
        }

        fetch("https://geolocation-db.com/json/")
        .then(response => {
            return response.json();
        }, "jsonp")
        .then(res => {
            console.log("[DB] All client information", res);
            setCtx(res);
            setHasContext(true);
            setCookie('IPv4', res.IPv4);
            setCookie('city', res.city);
            })
        .catch(err => {        
                console.log(err);
                setHasContext(false);
                setErr(err);
        })
    }

    if (hasContext && !registered) {
        let url = 'http://35.247.130.194:8080' + '/register';
        let browser = detect();
        console.log(process.env.REACT_APP_SERVER_URL, browser);
        let body = {
            'ip': ctx.IPv4,
            'location': ctx.city,
            'browser': browser.name+':'+browser.version,
        };
        fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(body),
        }).
            then(r => r.json()).
            then((r) => {
                setRegistered(true);
            })
            .catch((err) => {
                console.log(err);
            })
    }
    return (
        <div>
            <h1> {getMessage()} </h1>
        </div>
    );
}
export default Home;