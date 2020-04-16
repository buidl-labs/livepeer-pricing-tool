import React, { Component } from 'react'
import { VictoryChart, VictoryLine, VictoryAxis, VictoryBrushContainer, VictoryTooltip, createContainer, VictoryLabel } from 'victory'

export class PriceHistoryGraph extends Component {

    state = {
        zoomDomain: { x: [new Date(1990, 1, 1), new Date(2009, 1, 1)] }
    }

    pppdata = this.reformatData(this.props.data)

    handleZoom(domain) {
        this.setState({ zoomDomain: domain });
    }

    reformatData(data) {
        if (data.length > 72) {
            data = data.slice(0,72)
        }
        data = data.reverse()
        let newdata = []
        data.forEach(element => {
            newdata.push({
                Time: new Date(element.Time * 1000),
                PricePerPixel: element.PricePerPixel,
                label: element.PricePerPixel
            })
        })
        return newdata
    }

    render() {
        const VictoryZoomVoronoiContainer = createContainer("zoom", "voronoi");
        return (
            <div style={{ width: "70%", height: "60%", margin: "50px auto 0px auto"}}>
                <VictoryChart width={600} height={300} scale={{ x: "time" }}
                    padding={{ top: 0, left: 100, right: 100, bottom: 50 }}
                    containerComponent={
                        <VictoryZoomVoronoiContainer
                            zoomDimension="x"
                            zoomDomain={this.state.zoomDomain}
                            onZoomDomainChange={this.handleZoom.bind(this)}
                        />
                    }
                >
                    <VictoryLabel text="Price Per Pixel Chart" x={300} y={10} textAnchor="middle" />
                    <VictoryLine
                        style={{
                            data: { stroke: "#07b35f" }
                        }}
                        labelComponent={<VictoryTooltip/>}
                        data={this.pppdata}
                        x="Time"
                        y="PricePerPixel"
                    />

                </VictoryChart>
                <VictoryChart
                    padding={{ top: 0, left: 50, right: 50, bottom: 30 }}
                    width={400} height={50} scale={{ x: "time" }}
                    containerComponent={
                        <VictoryBrushContainer
                            brushDimension="x"
                            brushDomain={this.state.zoomDomain}
                            onBrushDomainChange={this.handleZoom.bind(this)}
                        />
                    }
                >
                    <VictoryAxis
                        tickFormat={(x) => {
                            let d = new Date(x)
                            return d.getDate() + "/" + d.getMonth()
                        }}
                        style={{tickLabels: {fontSize: 7}}}
                    />
                    <VictoryLine
                        style={{
                            data: { stroke: "#07b35f" }
                        }}
                        labelComponent={<VictoryTooltip/>}
                        data={this.pppdata}
                        x="Time"
                        y="PricePerPixel"
                    />
                </VictoryChart>
            </div>
        );
    }
}

export default PriceHistoryGraph
