import React, { Component } from 'react'
import Chart from "chart.js";
import classes from "./LineGraph.module.css";


export class PriceHistoryGraph extends Component {

    chartRef = React.createRef();

    componentDidMount() {
        const myChartRef = this.chartRef.current.getContext("2d");
        
        const data = this.reformatData(this.props.data)

        new Chart(myChartRef, {
            type: "line",
            data: {
                labels: data.time,
                datasets: [
                    {
                        label: "PricePerPixel",
                        data: data.ppp,
                        fill: false,
                        borderColor: "#26e98b"
                    }
                ]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                tooltips: {
                    displayColors: false
                }
            }
        });
    }

    reformatData (data) {
        let newdata = {
            time: [],
            ppp: []
        }
        data.forEach(element => {
            let time = new Date(element.Time*1000)
            time = time.toISOString().slice(0,-5).replace("T", " ")
            newdata.time.push(time)
            newdata.ppp.push(element.PricePerPixel)
        })
        return newdata
    }

    render() {
        let { data } = this.props.data
        console.log(data)
        data = this.reformatData(this.props.data)
        console.log(data)
        return (
            <div className={classes.graphContainer} style={graphStyle}>
                <canvas
                    id="myChart"
                    ref={this.chartRef}
                    width="800px"
                    height="600px"
                />
            </div>
        )
    }
}

const graphStyle = {
    margin: "0 auto"
}

export default PriceHistoryGraph
