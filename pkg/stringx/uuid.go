package stringx

import (
	"github.com/google/uuid"
	"strings"
)

func UUID() string {
	s := uuid.New().String()
	return strings.ReplaceAll(s, "-", "")
}
