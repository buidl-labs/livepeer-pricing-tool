import logo from './logo.svg'
import React, { Component } from 'react'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom'

import OrchestratorStats from './components/pages/OrchestratorStats'
import OrchestratorPriceHistory from './components/pages/OrchestratorPriceHistory'

import 'antd/dist/antd.css';
import './index.css';
import { Layout } from 'antd';
import { Row, Col } from 'antd';
const { Header, Content, Footer } = Layout;

export class App extends Component {
    render() {
        return (
            <Router>
                <Layout className="layout">
                    <Header>
                        <Row>
                            <Col span={1}>
                                <Link to="/">
                                    <img src={logo} alt="Logo" style={{ transform: "rotate(-90deg)" }} />
                                </Link>
                            </Col>
                            <Col span={11}>
                                <Link to="/">
                                    <h1 style={{ color: "white" }}>Livepeer | Pricing Tool </h1>    
                                </Link>
                            </Col>
                        </Row>
                    </Header>
                    <Content style={{ padding: '0 50px', minHeight: '85vh' }}>
                        <div className="site-layout-content">
                            <Route exact path="/" render={props => (
                                <OrchestratorStats />
                            )} />

                            <Route path="/priceHistory/:address" render={props => (
                                <OrchestratorPriceHistory {...props} />
                            )} />
                        </div>
                    </Content>
                    <Footer style={{ textAlign: 'center' }}>
                        Made with <span style={{ color: 'orangered' }}>&hearts;</span> by {' '}
                        <a target="_blank" rel="noopener noreferrer" href="http://buidllabs.io/" className="card-link">
                        BUIDL Labs
                        </a>.
                    </Footer>
                </Layout>
            </Router>
        )
    }
}

export default App
