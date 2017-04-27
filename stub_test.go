package gobmock

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stub", func() {
	It("is created", func() {
		Expect(Stub("foo")).NotTo(BeNil())

	})

	It("includes a basic shell stub", func() {
		stub := Stub("devatio-crederes").MockContents()
		Expect(stub).To(MatchRegexp("^devatio-crederes\\(\\)\\s*{"))
		Expect(stub).To(MatchRegexp("while read -t0.05; do\\s+:\\s+done"))
	})

	It("includes an exported shell stub", func() {
		stub := ExportedStub("shred_four_rhubarbs").MockContents()
		Expect(stub).To(MatchRegexp("}\\s+export -f shred_four_rhubarbs\n"))
	})
})
