import logo from './logo.svg'
import React, { Component } from 'react'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom'

import OrchestratorStats from './components/pages/OrchestratorStats'
import OrchestratorPriceHistory from './components/pages/OrchestratorPriceHistory'

import 'antd/dist/antd.css';
import './index.css';
import { Layout } from 'antd';
import { Row, Col } from 'antd';
import { Modal } from 'antd';
const { Header, Content, Footer } = Layout;

export class App extends Component {

    info() {
        Modal.info({
          title: 'There is a known issue for Brave browser.',
          content: (
            <div>
                <br/>
              <p>To resolve rendering issues on Brave,
                  <ul>
                      <li>Go to Brave Shield settings</li>
                      <li>Navigate to Device Recognition setting</li>
                      <li> In the dropdown menu, select "All device recognition attempts allowed".</li>
                  </ul>
                </p>
              <p>If there is any other issue, please report <a href="https://github.com/buidl-labs/livepeer-pricing-tool/issues">here</a>.</p>
            </div>
          ),
          onOk() {},
        });
      }

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
                            <Col span={17}>
                                <Link to="/">
                                    <h1 style={{ color: "white" }}> Livepeer | Pricing Tool </h1>    
                                </Link>
                            </Col>
                            <Col span={4}>
                                <button type="button" className="link-button" onClick={this.info}>Facing rendering issues?</button>
                            </Col>
                            <Col span={2}>
                                <button type="button" className="link-button">
                                    <a href="https://github.com/buidl-labs/livepeer-pricing-tool#faqs" target="_blank" rel="noopener noreferrer" style={{ color: "white" }}>FAQs</a>
                                </button>
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
