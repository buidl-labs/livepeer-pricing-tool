import React, { Component } from 'react'
import PriceHistoryGraph from '../PriceHistoryGraph'
import axios from 'axios'
import { PageHeader } from 'antd';

export class OrchestratorPriceHistory extends Component {

    state = {
        data: null
    }

    componentDidMount() {
        console.log(this.props.address)
        axios.get("http://localhost:9000/priceHistory/" + this.props.address)
        .then(res => this.setState({data: res.data}))
        .catch(err => console.log(err))
    }

    render() {
        console.log(this.state)
        if (this.state.data) {
            return (
                <div>
                    <PageHeader
                        className="site-page-header"
                        backIcon="false"
                        title="Orchestrator Price History"
                        subTitle={this.props.address}
                    />
                    <PriceHistoryGraph data={this.state.data} />
                </div>
            )
        } else {
            return (
                <div>Error in fetching data. Make sure the API is running at "localhost:9000". See console for more details.</div>
            )
        }
    }
}

export default OrchestratorPriceHistory