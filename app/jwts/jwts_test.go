package jwts

import (
	"testing"
)

func TestJwt_Check(t *testing.T) {
	j := NewJwt()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJyZWRyb2NrIiwiZXhwIjoiMTU4NDI4Njc1NCIsImlhdCI6IjE1ODQyNzU5NTQiLCJ1c2VybmFtZSI6IiIsInBhc3N3b3JkIjoiIn0%3D.JM50fM3wbKNrw9SWIil1k9gtFahWS7QF3xeQ15D4Z2k%3D"
	j.Check(token, "redrock")
}
