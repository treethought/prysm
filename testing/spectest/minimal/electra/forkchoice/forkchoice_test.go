package forkchoice

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v5/runtime/version"
	"github.com/prysmaticlabs/prysm/v5/testing/spectest/shared/common/forkchoice"
)

func TestMinimal_Electra_Forkchoice(t *testing.T) {
	t.Skip("TODO: Electra")
	forkchoice.Run(t, "minimal", version.Electra)
}