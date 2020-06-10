import React, { Component } from 'react'
import { Link } from 'react-router-dom';
import { Table } from 'antd';

export class OrchestratorTable extends Component {

    columns = [
    {
        title: 'Address',
        dataIndex: 'Address',
        align: 'center',
    },
    // {
    //     title: 'Activation Round',
    //     dataIndex: 'ActivationRound',
    //     align: 'center',
    //     sorter: {
    //         compare: (a, b) => a.ActivationRound - b.ActivationRound
    //     },
    // },
    // {
    //     title: 'Last Reward Round',
    //     dataIndex: 'LastRewardRound',
    //     align: 'center',
    //     sorter: {
    //         compare: (a, b) => a.LastRewardRound - b.LastRewardRound
    //     },
    // },
    {
        title: 'Delegated Stake',
        dataIndex: 'DelegatedStake',
        align: 'center',
        sorter: {
            compare: (a, b) => a.DelegatedStakeRaw - b.DelegatedStakeRaw
        },
    },
    {
        title: 'Reward Cut',
        dataIndex: 'RewardCut',
        align: 'center',
        sorter: {
            compare: (a, b) => a.RewardCut - b.RewardCut
        },
        render: (text) => text + "%"
    },
    {
        title: 'Fee Share',
        dataIndex: 'FeeShare',
        align: 'center',
        sorter: {
            compare: (a, b) => a.FeeShare - b.FeeShare
        },
        render: (text) => text + "%"
    },
    {
        title: 'Total Fees Earned',
        dataIndex: 'TotalGeneratedFees',
        align: 'center',
        sorter: {
            compare: (a, b) => a.TotalGeneratedFeesRaw - b.TotalGeneratedFeesRaw
        }
    },
    {
        title: 'Price Per Pixel',
        dataIndex: 'PricePerPixel',
        align: 'center',
        sorter: {
            compare: (a, b) => a.PricePerPixelRaw - b.PricePerPixelRaw
        }
    },
    {
        title: 'Price History',
        align: 'center',
        render: function (text, record, index) {
            return (
                <Link to={{
                    pathname: "/priceHistory/"+record.Address,
                    address: record.Address
                }}>
                    { "View" }
                </Link>
            )
        }
    },
    // {
    //     title: 'Active',
    //     dataIndex: 'Active',
    //     align: 'center',
    // },
    // {
    //     title: 'Status',
    //     dataIndex: 'Status',
    //     align: 'center',
    // },
    ];

    formatNumber(num, prec) {
        num = num.toFixed(prec)
        num = num.toString()
        let numstring = ""
        let periodpos = num.indexOf(".")
        let flag = 0
        numstring = num.slice(periodpos, num.length)
        for (let index = periodpos; index > 0; index=index-3) {
            let start = index - 3
            if (start < 0) {
                start = 0
            }
            if(flag===0) {
                numstring = num.slice(start, index) + numstring
                flag = 1
            } else {
                numstring = num.slice(start, index) + "," + numstring
            }
        }
        return numstring
    }

    processDelegatedStake(ds) {
        if (ds > 10**15) {
            return this.formatNumber(ds / 10**18, 3) + " LPT"
        } else {
            return this.formatNumber(ds, 3) + " LPTU"
        }
    }

    processFees(fees) {
        fees = fees * 1.0
        return this.formatNumber(fees / 10**18, 6) + " ETH"
    }

    processPPP(ppp) {
        if (ppp < 0) {
            return "-" + this.formatNumber(Math.abs(ppp), 3) + " wei"
        } else {
            return this.formatNumber(Math.abs(ppp), 3) + " wei"
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
                PricePerPixelRaw: element.PricePerPixel,
                PricePerPixel: this.processPPP(element.PricePerPixel),
                UpdatedAt: element.UpdatedAt,
                TotalGeneratedFeesRaw: element.TotalGeneratedFees,
                TotalGeneratedFees: this.processFees(element.TotalGeneratedFees),
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
