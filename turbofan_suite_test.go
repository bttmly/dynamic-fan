package turbofan_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTurbofan(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Turbofan Suite")
}
