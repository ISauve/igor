// +build !dnf,!librepo

package rpm

import (
	"errors"

	"github.com/ISauve/nikos/types"
)

func NewBackend(target *types.Target) (types.Backend, error) {
	return nil, errors.New("dnf backend not supported")
}
