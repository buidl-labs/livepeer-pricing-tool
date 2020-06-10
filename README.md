# Livepeer Pricing Tool

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
Please refer documentation for API usage [here](https://github.com/buidl-labs/livepeer-pricing-tool/tree/master/api).

### Hosting the UI
1. Move into the "ui" directory by `cd ui`
2. Run `npm install`
3. Run `npm start`


# FAQs

- **What is price per pixel? How to factor in when making a decision about which orchestrator to choose?** 

Price per pixel (wei per pixel) represents the fundamental cost unit used for the calculation of total transcoding fees payable by a broadcaster for a transcoding job.

Please refer [here](https://livepeer.readthedocs.io/en/latest/transcoding.html#configuring-payment-parameters) for more info about cost calculations based on number of streams and other video properties.

- **What is delegated stake, and why is it relevant?**

Delegated stake is crucial in determining whether an orchestrator will remain in a round's active orchestrator set or not. The process is as follows:

The active orchestrator set is locked in at the beginning of each round. Any staking activity that occurs during a round will impact membership of the active set in the following round. So, if an orchestrator accumulates stake such that it is in the top 100 orchestrators with the most stake it will join the active set in the next round. If an orchestrator previously was in the top 100 orchestrators with the most stake, but no longer is in the top 100, it will exit the active set in the next round.

- **What is reward cut % and fee share %? How to factor in when making a decision about which orchestrator to choose?**

Reward cut is the percentage of inflationary LPT rewards that the orchestrator will keep. The rest is distributed among its delegators.

Fee share is the percentage of ETH transcoding fees that the orchestrator will share with its delegators. The remainder is kept by the orchestrator with itself. 

A combination of reward cut and fee share gives orchestrators a powerful economic advantage over traditional centralized video providers since the value of the token offsets what they need to charge broadcasters to break even. With traditional centralized video providers, they have to charge you their cost of service for transcoding and distributing video plus a margin.

An explaination around the workings of entire Livepeer network could be found [here](https://livepeer.org/primer/)

- **What is total fees earned?**

Total fees earned represents the total fees earned by an orchestrator via transcoding (before distribution to delegators).


- **Why pricing tool doesn't show the full list of orchestrators?**

The pricing tool excludes the orchestrators for which the broadcaster recieves the price per pixel value as 0, considering them unusable for transcoding purposes. This may happen because of various reasons like an orchestrator being not responsive, insufficient broadcaster reserve etc. 

- **Why is there negative pricing being displayed on the tool?**

On the tool, negative pricing values for some orchestrators will be encountered sometimes. Negative pricing value was an internal issue with go-livepeer, which was resolved in the 0.5.6 release. Since this release, if any orchestrator hasn't updated its go-livepeer, negative values could be seen for these orchestrators sometimes. 