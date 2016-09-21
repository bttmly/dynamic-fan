package turbofan_test

import (
	"log"

	"github.com/nickb1080/turbofan"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Turbofan", func() {

	var a, b, c chan bool

	BeforeSuite(func() {
		log.SetOutput(GinkgoWriter)
	})

	BeforeEach(func() {
		a = make(chan bool)
		b = make(chan bool)
		c = make(chan bool)
	})

	Context("Turbofan", func() {
		It("Other channels receive when one sends", func(done Done) {
			turbofan.New(a, b, c)
			go func() {
				fromB := <-b
				Expect(fromB).To(BeTrue())
				fromC := <-c
				Expect(fromC).To(BeTrue())
				Consistently(a).ShouldNot(Receive())
				close(done)
			}()
			a <- true
		}, 0.5)

		It("All channels receive on .Broadcast()", func(done Done) {
			t := turbofan.New(a, b, c)
			go func() {
				fromA := <-a
				Expect(fromA).To(BeTrue())
				fromB := <-b
				Expect(fromB).To(BeTrue())
				fromC := <-c
				Expect(fromC).To(BeTrue())
				close(done)
			}()
			t.Broadcast(true)
		}, 0.5)

		It("All channels close on .Close()", func(done Done) {
			t := turbofan.New(a, b, c)
			go func() {
				<-a
				<-b
				<-c
				close(done)
			}()
			t.Close()
		}, 0.5)
	})

})
