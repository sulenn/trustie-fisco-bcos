package structs

type InstallChaincodeOption struct {
	Orgname            string    `json:"orgnization_name"`
	Peers              []string  `json:"peers"`
	ChaincodeName      string    `json:"chaincode_name"`
	ChaincodePath      string    `json:"chaincode_path"`
	ChaincodeVersion   string    `json:"chaincode_version"`
	ChaincodeType      string    `json:"chaincode_type"`   // 链码所用语言
}
