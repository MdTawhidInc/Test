package signer

//nolint:tagliatelle
type SignOperationResponse struct {
	PublicKey     string `json:"publicKey"`
	Signature     string `json:"signature"`
	CorrelationID string `json:"correlationId,omitempty"`
	Operation     string `json:"operation,omitempty"`
}

type Signer interface {
	Sign(nickname string, operation []byte) (*SignOperationResponse, error)
}
