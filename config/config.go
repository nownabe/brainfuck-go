package config

type Config struct {
	NEXT  string `json:"next"`
	PREV  string `json:"prev"`
	INC   string `json:"inc"`
	DEC   string `json:"dec"`
	READ  string `json:"read"`
	WRITE string `json:"write"`
	OPEN  string `json:"open"`
	CLOSE string `json:"close"`
}
