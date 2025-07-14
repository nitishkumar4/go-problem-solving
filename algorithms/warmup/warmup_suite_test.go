package warmup

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWarmup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Warmup Suite")
}
