import { useState } from "react";
import { Badge, Button } from 'react-bootstrap';

const DUMMY_DATA = [
    {
        "ip": "129.168.1.1",
        "num": 11
    },
    {
        "ip": "129.168.1.2",
        "num": 2
    },
    {
        "ip": "129.168.5.7",
        "num": 1
    },
    {
        "ip": "129.168.5.2",
        "num": 1
    }
]

const MostVisited = () => {
    const [ips, setIps] = useState([]);
    const [loaded, setLoaded] = useState(false);
    if (!loaded) {
        let url = 'http://35.247.130.194:8080' + '/most-visitors';
        fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            },
        }).
            then(r => r.json()).
            then((r) => {
                setIps(r.data.users);
                setLoaded(true);
            })
            .catch((err) => {
                console.log(err);
            })
    }

    function refresh(){
        let url = 'http://35.247.130.194:8080' + '/most-visitors';
        fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            },
        }).
            then(r => r.json()).
            then((r) => {
                setIps(r.data.users);
                setLoaded(true);
            })
            .catch((err) => {
                console.log(err);
            })
    }

    function getVariantByNum(i) {
        if (i > 50) {
            return "danger"
        }

        if (i >= 40) {
            return "warning"
        }

        if (i >= 30) {
            return "success"
        }

        if (i >= 20) {
            return "primary"
        }

        return "light"
    }

    return (
        <div>
            {ips.map(t => {
                return (
                    <Button variant={getVariantByNum(t.num)} key={t.ip}>
                        {t.ip} <Badge bg="secondary">{t.num}</Badge>
                        <span className="visually-hidden">unread messages</span>
                    </Button>
                )
            })}

            <div>
                <Button variant="success" onClick={refresh}>Refresh</Button>
            </div>

        </div>
    )
}

export default MostVisited;