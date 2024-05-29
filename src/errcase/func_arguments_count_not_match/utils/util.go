package utils

import "fmt"

func FormatUserInfo(name string, email string) string {
	return fmt.Sprintf("name: %s, email: %s", name, email)
}
