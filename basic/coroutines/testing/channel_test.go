package testing

import (
	"basic/coroutines"
	"testing"
)

func TestChanDemo(t *testing.T) {
	coroutines.ChanDemo(2)
}

func TestChannelClose(t *testing.T) {
	coroutines.ChannelCloseDemo2()
}

func TestChannelTickerDemo(t *testing.T) {
	coroutines.ChannelTickerDemo()
}
func TestChannelAfterDemo(t *testing.T) {
	coroutines.ChannelAfterDemo()
}
