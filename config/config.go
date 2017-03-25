package config

type Config struct {
	NEXT        string `json:"next"`
	PREV        string `json:"prev"`
	INC         string `json:"inc"`
	DEC         string `json:"dec"`
	READ        string `json:"read"`
	WRITE       string `json:"write"`
	OPEN        string `json:"open"`
	CLOSE       string `json:"close"`
	WHITESPACES string `json:"whitespaces"`
}

func (c *Config) Ops() []string {
	return []string{c.NEXT, c.PREV, c.INC, c.DEC, c.READ, c.WRITE, c.OPEN, c.CLOSE}
}
