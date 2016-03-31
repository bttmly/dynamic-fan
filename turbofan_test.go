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
		It("Works as expected", func(done Done) {
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
	})

})
