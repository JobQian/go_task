package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"strings"

	"golang.org/x/crypto/argon2"
)

func GeneratePassword(password string) string {
	salt := make([]byte, 16)
	rand.Read(salt)

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	return base64.RawStdEncoding.EncodeToString(salt) + "$" +
		base64.RawStdEncoding.EncodeToString(hash)
}

// 验证密码
func VerifyPassword(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 2 {
		return false, errors.New("hash 格式不正确")
	}

	// 提取 salt
	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, err
	}

	// 提取存储的 hash
	storedHash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}

	// 重新计算 hash（参数要和生成时保持一致）
	newHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// 使用 constant-time 比较，避免计时攻击 (恒定的比较时间)
	if subtle.ConstantTimeCompare(newHash, storedHash) == 1 {
		return true, nil
	}
	return false, nil
}
