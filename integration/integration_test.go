package integration_test

import (
	"os"
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
	_, err := client.New(config.T{Host: os.Getenv("MASTER_HOST"), Port: os.Getenv("MASTER_PORT")})
	Expect(err).NotTo(HaveOccurred())

	_, err = client.New(config.T{Host: os.Getenv("SECONDARY_1_HOST"), Port: os.Getenv("SECONDARY_1_PORT")})
	Expect(err).NotTo(HaveOccurred())

	_, err = client.New(config.T{Host: os.Getenv("SECONDARY_2_HOST"), Port: os.Getenv("SECONDARY_2_PORT")})
	Expect(err).NotTo(HaveOccurred())
})
