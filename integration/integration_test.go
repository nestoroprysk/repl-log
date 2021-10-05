package integration_test

import (
	"testing"

	"github.com/nestoroprysk/repl-log/client"
	"github.com/nestoroprysk/repl-log/config"
	"github.com/nestoroprysk/repl-log/message"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestReplLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Replication Log App Suite")
}

var _ = It("A sample message gets replicated", func() {
	m, err := client.New(config.Master)
	Expect(err).NotTo(HaveOccurred())

	a, err := client.New(config.SecondaryA)
	Expect(err).NotTo(HaveOccurred())

	b, err := client.New(config.SecondaryB)
	Expect(err).NotTo(HaveOccurred())

	msg := message.T("1234")
	Expect(m.PostMessage(msg)).To(Succeed())

	msgs, err := m.GetMessages()
	Expect(err).NotTo(HaveOccurred())
	Expect(msgs).To(ContainElement(msg))

	msgs, err = a.GetMessages()
	Expect(err).NotTo(HaveOccurred())
	Expect(msgs).To(ContainElement(msg))

	msgs, err = b.GetMessages()
	Expect(err).NotTo(HaveOccurred())
	Expect(msgs).To(ContainElement(msg))
})
