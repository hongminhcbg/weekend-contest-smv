import {
    Routes,
    Route
} from 'react-router-dom';
import Home from '../components/home';
import Logs from '../components/logs';
import MostiVisited from '../components/MostVisited';
import {Navbar, Container, Nav} from 'react-bootstrap';
import classes from './main.module.css';

function MainNavigate(){
    return(
        <div>
            <Nav defaultActiveKey="/home" as="ul" className={classes.text}>
            <Nav.Item as="li">
                <Nav.Link href="/">Home</Nav.Link>
            </Nav.Item>
            <Nav.Item as="li">
                <Nav.Link href="/logs">Logs</Nav.Link>
            </Nav.Item>
            <Nav.Item as="li">
                <Nav.Link href="/most-visited">Most Visited</Nav.Link>
            </Nav.Item>
            </Nav>
            <Routes>
                <Route path="/" exact={true} element={<Home/>}>
                </Route>
                <Route path="/logs" element={<Logs/>}>
                </Route>
                <Route path="/most-visited" element={<MostiVisited/>}>
                </Route>
            </Routes>

        </div>
    )
}

export default MainNavigate;