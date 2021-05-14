package token

// token model
type TokenModel struct {
	ID    uint64 `json:"id"`
	Typ   int8   `json:"typ"` // 0 admin,1 client,
	Token string `json:"token"`
}
