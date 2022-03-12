import { useState } from 'react';
import { Table, Button } from 'react-bootstrap';

const MOCK_DATA = [
    {
        "id": 36,
        "ip": "113.173.120.78",
        "browser": "chrome:99.0.4844",
        "location": "Ho Chi Minh City",
        "time": "2022-03-12T14:30:50Z"
    },
    {
        "id": 35,
        "ip": "113.173.120.78",
        "browser": "chrome:99.0.4844",
        "location": "Ho Chi Minh City",
        "time": "2022-03-12T14:30:50Z"
    },
    {
        "id": 34,
        "ip": "113.173.120.78",
        "browser": "server auto detact",
        "location": "Ho Chi Minh City",
        "time": "2022-03-12T14:30:01Z"
    },
    {
        "id": 33,
        "ip": "113.173.120.78",
        "browser": "server auto detact",
        "location": "Ho Chi Minh City",
        "time": "2022-03-12T14:30:01Z"
    },
    {
        "id": 32,
        "ip": "113.173.120.78",
        "browser": "server auto detact",
        "location": "Ho Chi Minh City",
        "time": "2022-03-12T14:29:36Z"
    },
    {
        "id": 31,
        "ip": "113.173.120.78",
        "browser": "server auto detact",
        "location": "Ho Chi Minh City",
        "time": "2022-03-12T14:29:36Z"
    },
    {
        "id": 30,
        "ip": "113.173.120.78",
        "browser": "server auto detact",
        "location": "Ho Chi Minh City",
        "time": "2022-03-12T14:23:11Z"
    },
    {
        "id": 29,
        "ip": "113.173.120.78",
        "browser": "server auto detact",
        "location": "Ho Chi Minh City",
        "time": "2022-03-12T14:23:11Z"
    },
    {
        "id": 28,
        "ip": "113.173.120.78",
        "browser": "eee",
        "location": "Ho Chi Minh City",
        "time": "2022-03-12T14:22:10Z"
    },
    {
        "id": 27,
        "ip": "113.173.120.78",
        "browser": "eee",
        "location": "Ho Chi Minh City",
        "time": "2022-03-12T14:22:10Z"
    }
]

const Logs = () => {
    const [logs, setLogs] = useState([]);
    const [hasMoreData, setHasMoreData] = useState(true);
    const [lastId, setLastId] = useState(999999);
    function loadMoreData() {
        if(!hasMoreData) {
            // do nothing
            return
        }

        let url = 'http://127.0.0.1:8080' + '/visitors?last_id=' + lastId;
        fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            },
        }).
            then(r => r.json()).
            then((r) => {
                let newLogs = [...logs, ...r.data.users];
                setLogs(newLogs);
                setHasMoreData(r.data.has_more_data);
                if(r.data.has_more_data) {
                    let last_id = r.data.users[r.data.users.length-1].id;
                    setLastId(last_id);
                }
            })
            .catch((err) => {
                console.log(err);
            })

    }
    function getButtonVariant() {
        if (hasMoreData) {
            return "success"
        }

        return "danger"
    }

    function getButtonText() {
        if (hasMoreData) {
            return "Load more data"
        }

        return "Out of data"
    }

    return (
        <div>
            <Table striped bordered hover>
                <thead>
                    <tr>
                        <th>id</th>
                        <th>IPV4</th>
                        <th>Location</th>
                        <th>Browser</th>
                        <th>Time</th>
                    </tr>
                </thead>
                <tbody>
                    {logs.map(t => {
                        return (
                            <tr>
                                <td>{t.id}</td>
                                <td>{t.ip}</td>
                                <td>{t.location}</td>
                                <td>{t.browser}</td>
                                <td>{t.time}</td>
                            </tr>
                        )
                    })}
                </tbody>
            </Table>

            <Button variant={getButtonVariant()} onClick={loadMoreData}>{getButtonText()}</Button>
        </div>
    )
}

export default Logs;
