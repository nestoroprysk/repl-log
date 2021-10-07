package integration_test

import (
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/nestoroprysk/repl-log/client"
	"github.com/nestoroprysk/repl-log/config"
	"github.com/nestoroprysk/repl-log/message"
	"github.com/nestoroprysk/repl-log/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var testConfig config.Config

var _ = BeforeSuite(func() {
	c, err := config.Read()
	Expect(err).NotTo(HaveOccurred())

	testConfig = c
})

func TestReplLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Replication Log App Suite")
}

var _ = It("Namespaces get listed", func() {
	m, _, n, clean := env()
	defer clean()

	msg := message.T{Message: "1234", Namespace: n}
	Expect(m.PostMessage(msg)).To(Succeed())

	ns, err := m.GetNamespaces()
	Expect(err).NotTo(HaveOccurred())
	Expect(ns).To(ContainElements(message.DefaultNamespace, n))
})

var _ = It("Deleting a namespace deletes all the related messages", func() {
	m, ss, n, clean := env()
	defer clean()

	msg := message.T{Message: "1234", Namespace: n}
	Expect(m.PostMessage(msg)).To(Succeed())

	ns, err := m.GetNamespaces()
	Expect(err).NotTo(HaveOccurred())
	Expect(ns).To(ContainElements(message.DefaultNamespace, n))

	ok, err := m.DeleteNamespace(n)
	Expect(err).NotTo(HaveOccurred())
	Expect(ok).To(BeTrue())

	for _, c := range append(ss, m) {
		ns, err = c.GetNamespaces()
		Expect(err).NotTo(HaveOccurred())
		Expect(ns).NotTo(ContainElements(n))

		ms, err := m.GetMessages(n)
		Expect(err).NotTo(HaveOccurred())
		Expect(ms).To(BeEmpty())
	}
})

var _ = It("Deleting a namespace that doesn't exist returns false", func() {
	m, _, n, clean := env()
	defer clean()

	ok, err := m.DeleteNamespace(n + "1234")
	Expect(err).NotTo(HaveOccurred())
	Expect(ok).To(BeFalse())
})

var _ = It("A sample message gets replicated", func() {
	m, ss, n, clean := env()
	defer clean()

	msg := message.T{Message: "1234", Namespace: n}
	Expect(m.PostMessage(msg)).To(Succeed())

	for _, c := range append(ss, m) {
		msgs, err := c.GetMessages(n)
		Expect(err).NotTo(HaveOccurred())
		Expect(msgs).To(ConsistOf(msg))
	}
})

var _ = It("Many messages get replicated and respect the order", func() {
	m, ss, n, clean := env()
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

	for _, c := range append(ss, m) {
		result, err := c.GetMessages(n)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(msgs))
	}
})

func env() (*client.T, []*client.T, message.Namespace, func()) {
	m, err := client.New(testConfig.Listen)
	Expect(err).NotTo(HaveOccurred())

	ss, err := util.ToClients(testConfig.Replicate)
	Expect(err).NotTo(HaveOccurred())

	u, err := uuid.NewRandom()
	Expect(err).NotTo(HaveOccurred())

	n := message.Namespace(u.String())

	clean := func() {
		_, err := m.DeleteNamespace(n)
		Expect(err).NotTo(HaveOccurred())
	}

	return m, ss, n, clean
}
