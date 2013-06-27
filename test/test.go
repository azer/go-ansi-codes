package main

import (
	"github.com/azer/ansi-codes.go"
	"github.com/azer/test.go"
)

func main () {

	test.It("exports ansi codes of styles", func (expect test.Expect) {
		expect(ansicodes.Red, "\033[31m")
		expect(ansicodes.Blue, "\033[34m")
		expect(ansicodes.Bold, "\033[1m")
		expect(ansicodes.YellowBg, "\033[43m")
	})

}
