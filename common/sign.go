package common

import (
	"path"

	"go_test/internal/sign"
)

func Sign(obj string, parent string) string {
	return sign.Sign(path.Join(parent, obj))
}
