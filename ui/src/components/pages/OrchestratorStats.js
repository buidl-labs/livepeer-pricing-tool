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

                axios.post(Config.livepeer_api, {
                    query: `{
                            transcoders {
                                totalGeneratedFees
                                id
                            }
                        }`,
                    headers: {
                        'Content-Type': 'application/json'
                    }
                })
                .then((result) => {
                    result = result.data.data.transcoders
                    let new_data = res.data
                    for (let i = 0; i < new_data.length; i++) {
                        new_data[i]["TotalGeneratedFees"] = null
                        for (let j = 0; j < result.length; j++) { 
                            if (new_data[i].Address === result[j].id) {
                                new_data[i]["TotalGeneratedFees"] = result[j].totalGeneratedFees
                                break
                            }
                        } 
                    }
                    this.setState({ data: new_data, responseStatus: "success" })
                })
                .catch(err => {
                    console.log(err)
                    this.setState({ data: null, responseStatus: "failed" })
                })
                // this.setState({ data: res.data, responseStatus: "success" })
            })
            .catch(err => {
                console.log(err)
                this.setState({ data: null, responseStatus: "failed" })
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
