package data

import (
	"encoding/json"
)

type Asset struct {
	AssetID       int     `json:"asset_id"`
	ExecContract  int     `json:"exec_contract"`
	AssetContract int     `json:"asset_contract"`
	ImgPreview    string  `json:"img_preview"`
	Executor      int     `json:"executor"`
	Desc          string  `json:"desc"`
	Price         float64 `json:"price"`
}

type Contract struct {
	ContractID int    `json:"contract_id"`
	Addr       string `json:"addr"`
	IsExec     bool   `json:"is_exec"`
	IsFinished bool   `json:"is_finished"`
}

type User struct {
	UserID        int    `json:"user_id"`
	WalletPubkey  string `json:"wallet_pubkey"`
	ProfilePicUrl string `json:"profile_pic_url"`
	Bio           string `json:"bio"`
}

// Status of the contribution tx.
type ContribStatus int

const (
	// Contribution is still pending.
	ContribPending ContribStatus = 0
	// Contirbution succesfully committed.
	ContribSuccess ContribStatus = 1
	// Contribution tx failed.
	ContribFailure ContribStatus = 2
)

type Contribution struct {
	ContribID    int           `json:"contrib_id"`
	ExecContract int           `json:"exec_contract"`
	User         int           `json:"user"`
	Amt          float64       `json:"amt"`
	Status       ContribStatus `json:"status"`
}

// Convert types into JSON objects

func (asset Asset) ToJSON() (string, error) {
	assetJson, err := json.Marshal(&asset)
	if err != nil {
		return "", err
	}

	return string(assetJson), nil
}

func (contract Contract) ToJSON() (string, error) {
	contractJson, err := json.Marshal(&contract)
	if err != nil {
		return "", err
	}

	return string(contractJson), nil
}

func (user User) ToJSON() (string, error) {
	userJson, err := json.Marshal(&user)
	if err != nil {
		return "", err
	}

	return string(userJson), nil
}

func (contrib Contribution) ToJSON() (string, error) {
	contribJson, err := json.Marshal(&contrib)
	if err != nil {
		return "", err
	}

	return string(contribJson), nil
}
