import React, { Component } from 'react'
import PriceHistoryGraph from '../PriceHistoryGraph'
import axios from 'axios'
import { PageHeader } from 'antd';
import Config from '../../Config';

export class OrchestratorPriceHistory extends Component {

    state = {
        data: null,
        responseStatus: null
    }

    componentDidMount() {
        axios.get(Config.api_url + "/priceHistory/" + this.props.match.params.address)
        .then(res => {
            this.setState({data: res.data, responseStatus: "success"})
        })
        .catch(err => {
            console.log(err)
            this.setState({data: null, responseStatus: "failed"})
        })
    }

    render() {
        if (!this.state.responseStatus) {
            return (
                <React.Fragment> 
                    <p>Fetching data from the server...</p>
                </React.Fragment>
            )
        } else {
            if (this.state.data) {
                return (
                    <div>
                        <PageHeader
                            className="site-page-header"
                            backIcon="false"
                            title="Orchestrator Price History"
                            subTitle={this.props.match.params.address}
                        />
                        <PriceHistoryGraph data={this.state.data} style={{ width: "80vw", height: "75vh"}}/>
                    </div>
                )
            } else {
                return (
                    <div>Error in fetching data from the server. See console for more details.</div>
                )
            }
        }
    }
}

export default OrchestratorPriceHistory