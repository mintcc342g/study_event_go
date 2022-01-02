package entity

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Entity Suite")
}
