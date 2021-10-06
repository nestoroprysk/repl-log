package integration_test

import (
	"sync"
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

	for _, c := range []*client.T{m, a, b} {
		msgs, err := c.GetMessages(n)
		Expect(err).NotTo(HaveOccurred())
		Expect(msgs).To(ConsistOf(msg))
	}
})

var _ = It("Many messages get replicated and respect the order", func() {
	m, a, b, n, clean := env()
	defer clean()

	var msgs []message.T
	for i := 0; i < 100; i++ {
		u, err := uuid.NewRandom()
		Expect(err).NotTo(HaveOccurred())
		msgs = append(msgs, message.T{Message: u.String(), Namespace: n})
	}

	var wg sync.WaitGroup
	for _, _msg := range msgs {
		wg.Add(1)
		msg := _msg
		go func() {
			Expect(m.PostMessage(msg)).To(Succeed())
			wg.Done()
		}()
	}
	wg.Wait()

	for _, c := range []*client.T{m, a, b} {
		result, err := c.GetMessages(n)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(msgs))
	}
})

func env() (*client.T, *client.T, *client.T, message.Namespace, func()) {
	m, err := client.New(config.Master)
	Expect(err).NotTo(HaveOccurred())

	a, err := client.New(config.SecondaryA)
	Expect(err).NotTo(HaveOccurred())

	b, err := client.New(config.SecondaryB)
	Expect(err).NotTo(HaveOccurred())

	u, err := uuid.NewRandom()
	Expect(err).NotTo(HaveOccurred())

	n := message.Namespace(u.String())

	clean := func() {
		_, err := m.DeleteNamespace(n)
		Expect(err).NotTo(HaveOccurred())
	}

	return m, a, b, n, clean
}
