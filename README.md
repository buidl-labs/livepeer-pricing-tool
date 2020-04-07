# livepeer-pricing-tool

A tool for enhancing the price visibility in livepeer ecosystem.


# How to Setup

### Prerequisites
1. Setup a livepeer broadcaster/orchestrator node by following the instructions from [here](https://livepeer.readthedocs.io/en/latest/quickstart.html)
2. Install latest version of Go from [here](https://golang.org/doc/install)

### Hosting the API
1. Clone this repository.
2. Move into the api directory by `cd api`
3. Launch the API by `go run ./cmd/main.go`

### Accessing the API endpoints 
Hit these enpoints on your browser:
* localhost:9000/orchestratorStats
* localhost:9000/priceHistory/{orchestrator_address}