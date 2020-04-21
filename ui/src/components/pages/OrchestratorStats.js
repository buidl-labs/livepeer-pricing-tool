import React, { Component } from 'react'
import axios from 'axios'
import OrchestratorTable from '../OrchestratorTable'
import { PageHeader } from 'antd';
import Config from '../../Config'

export class OrchestratorStats extends Component {

    state = {
        data: null,
        responseStatus: null
    }

    componentDidMount() {
        axios.get(Config.api_url + '/orchestratorStats')
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
                    <React.Fragment>
                        <PageHeader
                            className="site-page-header"
                            backIcon="false"
                            title="Orchestrator Statistics"
                            subTitle=""
                        />
                        <OrchestratorTable data={this.state.data} />
                    </React.Fragment>
                )
            } else {
                return (
                    <React.Fragment> 
                        <p>Error in fetching data from the server. See console for more details.</p>
                    </React.Fragment>
                )
            }
        }
    }
}

export default OrchestratorStats
