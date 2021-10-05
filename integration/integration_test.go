package integration_test

import (
	"testing"

	"github.com/nestoroprysk/repl-log/client"
	"github.com/nestoroprysk/repl-log/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestReplLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Replication Log App Suite")
}

var _ = It("Replication Log does the job", func() {
	_, err := client.New(config.Master)
	Expect(err).NotTo(HaveOccurred())

	_, err = client.New(config.SecondaryA)
	Expect(err).NotTo(HaveOccurred())

	_, err = client.New(config.SecondaryB)
	Expect(err).NotTo(HaveOccurred())
})
