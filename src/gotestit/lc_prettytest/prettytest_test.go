package lc_prettytest

import (
	"github.com/cloudfoundry/gosteno"
	"github.com/remogatto/prettytest"
	l "gotestit/loggregatorclient"
	"net"
	"testing"
)

type testSuite struct {
	prettytest.Suite
	bufferSize        int
	loggregatorClient l.LoggregatorClient
	udpListener       *net.UDPConn
}

func TestRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testSuite),
	)
}

func setup(t *testSuite) {
	t.bufferSize = 4096
	t.loggregatorClient = l.NewLoggregatorClient("localhost:9876", gosteno.NewLogger("TestLogger"), t.bufferSize)

	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:9876")
	t.Equal(err, nil)

	t.udpListener, err = net.ListenUDP("udp", udpAddr)
	t.Equal(err, nil)
}

func (t *testSuite) TestSend_WithSetup() {
	setup(t)
	defer t.udpListener.Close()

	expectedOutput := []byte("Important Testmessage")
	t.loggregatorClient.Send(expectedOutput)

	buffer := make([]byte, t.bufferSize)
	readCount, _, err := t.udpListener.ReadFromUDP(buffer)
	t.Equal(err, nil)

	received := string(buffer[:readCount])
	t.Equal(string(expectedOutput), received)
}

func (t *testSuite) TestDontSendEmptyData2_WithSetup() {
	setup(t)
	defer t.udpListener.Close()

	firstMessage := []byte("")
	secondMessage := []byte("hi")
	t.loggregatorClient.Send(firstMessage)
	t.loggregatorClient.Send(secondMessage)

	buffer := make([]byte, t.bufferSize)
	readCount, _, err := t.udpListener.ReadFromUDP(buffer)
	t.Equal(err, nil)

	received := string(buffer[:readCount])
	t.Equal(string(secondMessage), received)
}
