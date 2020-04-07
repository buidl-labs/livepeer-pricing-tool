# livepeer-pricing-tool

A tool for enhancing the price visibility in livepeer ecosystem.


# How to Setup

### Prerequisites
1. Setup a livepeer broadcaster node by following the instructions from [here](https://livepeer.readthedocs.io/en/latest/quickstart.html)
2. Install latest version of Go from [here](https://golang.org/doc/install)

### Setup of the tool
1. Clone this repo.
2. Move into the server directory by `cd livepeer-pricing-tool/server`
3. Run the backend server by `go run main.go` OR `go build` followed by `./server`

### Accessing the endpoints 
Hit these enpoints on your browser:
* localhost:9000/orchestratorStats
* localhost:9000/priceHistory/{orchestrator_address}