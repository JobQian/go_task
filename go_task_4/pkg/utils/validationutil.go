package utils

import (
	"fmt"
	"regexp"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("field '%s' %s", e.Field, e.Message)
}

// CheckNotEmpty 校验字段非空
func CheckNotEmpty(fieldName, value string) error {
	if value == "" {
		return &ValidationError{Field: fieldName, Message: "cannot be empty"}
	}
	return nil
}

// CheckMinLength 校验最小长度
func CheckMinLength(fieldName, value string, min int) error {
	if len(value) < min {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at least %d characters", min)}
	}
	return nil
}

// CheckEmail 校验邮箱格式
func CheckEmail(fieldName, value string) error {
	if value == "" {
		return &ValidationError{Field: fieldName, Message: "cannot be empty"}
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(value) {
		return &ValidationError{Field: fieldName, Message: "is not a valid email"}
	}
	return nil
}

// 校验多个字段
func Validate(validators ...error) error {
	for _, v := range validators {
		if v != nil {
			return v
		}
	}
	return nil
}
