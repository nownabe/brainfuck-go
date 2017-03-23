package token

type Token int

const (
	NEXT Token = iota
	PREV
	INC
	DEC
	READ
	WRITE
	OPEN
	CLOSE
	INVALID
	EOF
)
