package intx

type Credentials struct {
	AccessKey   string `json:"accessKey"`
	Passphrase  string `json:"passphrase"`
	SigningKey  string `json:"signingKey"`
	PortfolioId string `json:"portfolioId"`
}
