package password

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	// testHashPassword()
	password := RandomString(10)

	hashPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotZero(t, hashPassword)
	require.NoError(t, CheckPassword(password, hashPassword))
	wrongPassword := RandomString(10)
	require.Error(t, CheckPassword(wrongPassword, hashPassword))
}

//func TestGenPassword(t *testing.T) {
//	password := "123456"
//	hashPassword, _ := HashPassword(password)
//	fmt.Println(hashPassword)
//}

const alphabetic = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString 生成一个长度为n的随机字符串
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabetic)
	for i := 0; i < n; i++ {
		c := alphabetic[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
