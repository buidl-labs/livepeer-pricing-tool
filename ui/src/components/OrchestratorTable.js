import React, { Component } from 'react'
import { Link } from 'react-router-dom';
import { Table } from 'antd';

export class OrchestratorTable extends Component {

    columns = [
    {
        title: 'Address',
        dataIndex: 'Address',
        align: 'center',
        render: text => <Link to={{
                    pathname: "/priceHistory",
                    address: text
                }}>
                    {text}
                </Link>
    },
    {
        title: 'Activation Round',
        dataIndex: 'ActivationRound',
        align: 'center',
        sorter: {
            compare: (a, b) => a.ActivationRound - b.ActivationRound
        },
    },
    {
        title: 'Last Reward Round',
        dataIndex: 'LastRewardRound',
        align: 'center',
        sorter: {
            compare: (a, b) => a.LastRewardRound - b.LastRewardRound
        },
    },
    {
        title: 'Delegated Stake (LPT/LPTU)',
        dataIndex: 'DelegatedStake',
        align: 'center',
        sorter: {
            compare: (a, b) => a.DelegatedStakeRaw - b.DelegatedStakeRaw
        },
    },
    {
        title: 'Reward Cut (%)',
        dataIndex: 'RewardCut',
        align: 'center',
        sorter: {
            compare: (a, b) => a.RewardCut - b.RewardCut
        },
    },
    {
        title: 'Fee Share (%)',
        dataIndex: 'FeeShare',
        align: 'center',
        sorter: {
            compare: (a, b) => a.FeeShare - b.FeeShare
        },
    },
    {
        title: 'Price Per Pixel (wei/pixel)',
        dataIndex: 'PricePerPixel',
        align: 'center',
        sorter: {
            compare: (a, b) => a.PricePerPixel - b.PricePerPixel
        },
    },
    {
        title: 'Active',
        dataIndex: 'Active',
        align: 'center',
    },
    {
        title: 'Status',
        dataIndex: 'Status',
        align: 'center',
    },
    ];

    processDelegatedStake(ds) {
        if (ds > 10**12) {
            return (ds / 10**18).toString() + " LPT"
        } else {
            return ds.toString() + " LPTU"
        }
    }

    preprocessData(data) {
        let newdata = []
        data.forEach(element => {
            newdata.push({
                key: element.Address,
                Address: element.Address,
                ServiceURI: element.ServiceURI,
                LastRewardRound: element.LastRewardRound,
                RewardCut: element.RewardCut / 10000,
                FeeShare: element.FeeShare / 10000,
                DelegatedStakeRaw: element.DelegatedStake,
                DelegatedStake: this.processDelegatedStake(element.DelegatedStake),
                ActivationRound: element.ActivationRound,
                DeactivationRound: element.DeactivationRound,
                Active: (element.Active ? "Active" : "Inactive"),
                Status: element.Status,
                PricePerPixel: element.PricePerPixel,
                UpdatedAt: element.UpdatedAt
            })
        });
        return newdata
    }

    onChange (pagination, filters, sorter, extra) {
        console.log('params', pagination, filters, sorter, extra);
    }
    
    
    render() {
        const data = this.preprocessData(this.props.data)
        return (
            <Table columns={this.columns} dataSource={data} onChange={this.onChange} pagination={false}/>
        )
    }
}

export default OrchestratorTable
