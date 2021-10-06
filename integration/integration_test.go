package integration_test

import (
	"testing"

	"github.com/google/uuid"
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

var _ = It("Namespaces get listed", func() {
	m, _, _, n, clean := env()
	defer clean()

	msg := message.T{Message: "1234", Namespace: n}
	Expect(m.PostMessage(msg)).To(Succeed())

	ns, err := m.GetNamespaces()
	Expect(err).NotTo(HaveOccurred())
	Expect(ns).To(ContainElements(message.DefaultNamespace, n))
})

var _ = It("Deleting a namespace deletes all the related messages", func() {
	m, a, b, n, clean := env()
	defer clean()

	msg := message.T{Message: "1234", Namespace: n}
	Expect(m.PostMessage(msg)).To(Succeed())

	ns, err := m.GetNamespaces()
	Expect(err).NotTo(HaveOccurred())
	Expect(ns).To(ContainElements(message.DefaultNamespace, n))

	ok, err := m.DeleteNamespace(n)
	Expect(err).NotTo(HaveOccurred())
	Expect(ok).To(BeTrue())

	for _, c := range []*client.T{m, a, b} {
		ns, err = c.GetNamespaces()
		Expect(err).NotTo(HaveOccurred())
		Expect(ns).NotTo(ContainElements(n))

		ms, err := m.GetMessages(n)
		Expect(err).NotTo(HaveOccurred())
		Expect(ms).To(BeEmpty())
	}
})

var _ = It("Deleting a namespace that doesn't exist returns false", func() {
	m, _, _, n, clean := env()
	defer clean()

	ok, err := m.DeleteNamespace(n + "1234")
	Expect(err).NotTo(HaveOccurred())
	Expect(ok).To(BeFalse())
})

var _ = It("A sample message gets replicated", func() {
	m, a, b, n, clean := env()
	defer clean()

	msg := message.T{Message: "1234", Namespace: n}
	Expect(m.PostMessage(msg)).To(Succeed())

	msgs, err := m.GetMessages(n)
	Expect(err).NotTo(HaveOccurred())
	Expect(msgs).To(ConsistOf(msg))

	msgs, err = a.GetMessages(n)
	Expect(err).NotTo(HaveOccurred())
	Expect(msgs).To(ConsistOf(msg))

	msgs, err = b.GetMessages(n)
	Expect(err).NotTo(HaveOccurred())
	Expect(msgs).To(ConsistOf(msg))
})

func env() (*client.T, *client.T, *client.T, message.Namespace, func()) {
	m, err := client.New(config.Master)
	Expect(err).NotTo(HaveOccurred())

	a, err := client.New(config.SecondaryA)
	Expect(err).NotTo(HaveOccurred())

	b, err := client.New(config.SecondaryB)
	Expect(err).NotTo(HaveOccurred())

	u, err := uuid.NewUUID()
	Expect(err).NotTo(HaveOccurred())

	n := message.Namespace(u.String())

	clean := func() {
		_, err := m.DeleteNamespace(n)
		Expect(err).NotTo(HaveOccurred())
	}

	return m, a, b, n, clean
}
