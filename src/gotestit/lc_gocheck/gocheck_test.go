package lc_gocheck

import (
	"github.com/cloudfoundry/gosteno"
	l "gotestit/loggregatorclient"
	. "launchpad.net/gocheck"
	"net"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct {
	bufferSize        int
	loggregatorClient l.LoggregatorClient
	udpListener       *net.UDPConn
}

var _ = Suite(&MySuite{})

func (s *MySuite) SetUpTest(c *C) {
	s.bufferSize = 4096

	s.loggregatorClient = l.NewLoggregatorClient("localhost:9876", gosteno.NewLogger("TestLogger"), s.bufferSize)

	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:9876")
	c.Assert(err, IsNil)

	s.udpListener, err = net.ListenUDP("udp", udpAddr)

	c.Assert(err, IsNil)
}

func (s *MySuite) TearDownTest(c *C) {
	s.udpListener.Close()
}

func (s *MySuite) TestSend(c *C) {
	expectedOutput := []byte("Important Testmessage")
	s.loggregatorClient.Send(expectedOutput)

	buffer := make([]byte, s.bufferSize)
	readCount, _, err := s.udpListener.ReadFromUDP(buffer)
	c.Assert(err, IsNil)

	received := string(buffer[:readCount])
	c.Assert(string(expectedOutput), Equals, received)
}

func (s *MySuite) TestDontSendEmptyData(c *C) {
	firstMessage := []byte("")
	secondMessage := []byte("hi")
	s.loggregatorClient.Send(firstMessage)
	s.loggregatorClient.Send(secondMessage)

	buffer := make([]byte, s.bufferSize)
	readCount, _, err := s.udpListener.ReadFromUDP(buffer)
	c.Assert(err, IsNil)

	received := string(buffer[:readCount])
	c.Assert(string(secondMessage), Equals, received)
}
