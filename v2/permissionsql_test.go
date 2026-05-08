package permissionsql

import (
	"testing"

	"github.com/xyproto/pinterface/v2"
)

func TestInterface(t *testing.T) {
	perm, err := NewWithConf(connectionString)
	if err != nil {
		t.Error(err)
	}
	// Check that the value qualifies for the interface
	var _ pinterface.IPermissions = perm
}
