package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPreStart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "vxlan-policy-agent pre-start Integration Suite")
}
