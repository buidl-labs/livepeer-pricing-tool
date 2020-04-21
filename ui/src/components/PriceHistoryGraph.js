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
                        label: "Price Per Pixel",
                        data: data.ppp,
                        fill: false,
                        borderColor: "#07b35f",
                        borderWidth: 2,
                        pointRadius: 1.5
                    }
                ]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                tooltips: {
                    displayColors: false
                },
                scales: {
                    xAxes: [{
                        type: "time",
                        time: {
                            unit: "day"
                        },
                        gridLines: {
                            drawOnChartArea: false
                        },
                        scaleLabel: {
                            display: true,
                            labelString: "Time"
                        }
                    }],
                    yAxes: [{
                        gridLines: {
                            drawOnChartArea: false
                        },
                        scaleLabel: {
                            display: true,
                            labelString: "Price Per Pixel"
                        }
                    }],
                }
            }
        });
    }

    reformatData (data) {
        data = data.reverse()
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

    render () {
        return (
            <div className={classes.graphContainer}>
                <canvas
                    id="myChart"
                    ref={this.chartRef}
                />
            </div>
        )
    }
}

export default PriceHistoryGraph
