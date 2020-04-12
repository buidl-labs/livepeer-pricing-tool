package model

import (
	"math/big"
)

// Orchestrator is the type binding for the result from /registeredOrchestrators endpoint of livepeer broadcaster node.
type Orchestrator struct {
	Address           string `json:"Address"`
	ServiceURI        string `json:"ServiceURI"`
	LastRewardRound   int    `json:"LastRewardRound"`
	RewardCut         int    `json:"RewardCut"`
	FeeShare          int    `json:"FeeShare"`
	DelegatedStake    *big.Int  `json:"DelegatedStake"`
	ActivationRound   int    `json:"ActivationRound"`
	DeactivationRound *big.Int  `json:"DeactivationRound"`
	Active            bool   `json:"Active"`
	Status            string `json:"Status"`
	PricePerPixel     string `json:"PricePerPixel"`
}

// DBOrchestrator is type binding for a row result from Orchestrators table
type DBOrchestrator struct {
	Address           string `json:"Address"`
	ServiceURI        string `json:"ServiceURI"`
	LastRewardRound   int    `json:"LastRewardRound"`
	RewardCut         int    `json:"RewardCut"`
	FeeShare          int    `json:"FeeShare"`
	DelegatedStake    string  `json:"DelegatedStake"`
	ActivationRound   int    `json:"ActivationRound"`
	DeactivationRound string  `json:"DeactivationRound"`
	Active            bool   `json:"Active"`
	Status            string `json:"Status"`
	PricePerPixel     string `json:"PricePerPixel"`
	UpdatedAt         int    `json:"UpdatedAt"`
}

// DBPriceHistory is the type binding for a row result from PriceHistory table
type DBPriceHistory struct {
	Address       string `json:"Address"`
	PricePerPixel string `json:"PricePerPixel"`
	Time          int    `json:"Time"`
}

// APIOrchestrator is type binding for an orchestrators object returned by the pricing tool API endpoints
type APIOrchestrator struct {
	Address           string `json:"Address"`
	ServiceURI        string `json:"ServiceURI"`
	LastRewardRound   int    `json:"LastRewardRound"`
	RewardCut         int    `json:"RewardCut"`
	FeeShare          int    `json:"FeeShare"`
	DelegatedStake    *big.Int  `json:"DelegatedStake"`
	ActivationRound   int    `json:"ActivationRound"`
	DeactivationRound *big.Int  `json:"DeactivationRound"`
	Active            bool   `json:"Active"`
	Status            string `json:"Status"`
	PricePerPixel     float64 `json:"PricePerPixel"`
	UpdatedAt         int    `json:"UpdatedAt"`
}

// APIPriceHistory is the type binding for the price history retured by the pricing tool API endpoints
type APIPriceHistory struct {
	PricePerPixel float64 `json:"PricePerPixel"`
	Time          int    `json:"Time"`
}

