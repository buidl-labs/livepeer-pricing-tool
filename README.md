# livepeer-pricing-tool

One liner: A tool for enhancing the price visibility in livepeer ecosystem. 

In detail:
* The Livepeer network contains various nodes (broadcasters, delegators, orchestrators/transcoders, and end-users), which work together to constitute a decentralized streaming service. 
* The orchestrators/transcoders are the heart of this network. These nodes handle the video transcoding operations and help the broadcasters in reaching out to users across various platforms. 
* However, there doesn’t exist a solution for getting an overview of prices and fees being charged for the transcoding process in the network. 
* We aim to fill this void via the development of a Price Monitoring Tool for the network.



# How to Setup

### Prerequisites
1. Setup a livepeer broadcaster/orchestrator node by following the instructions from [here](https://livepeer.readthedocs.io/en/latest/quickstart.html)
2. Install latest version of Go from [here](https://golang.org/doc/install)
3. Clone this repository.

### Hosting the API
1. Move into the "api" directory by `cd api`
2. Launch the API by `go run ./cmd/main.go`

### Accessing the API endpoints 
Hit these enpoints on your browser:
* localhost:9000/orchestratorStats
* localhost:9000/priceHistory/{orchestrator_address}

### Hosting the UI
1. Move into the "ui" directory by `cd ui`
2. Run `npm install`
3. Run `npm start`
