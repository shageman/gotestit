package lc_gospec

import (
	"github.com/cloudfoundry/gosteno"
	"github.com/orfjackal/gospec/src/gospec"
	. "github.com/orfjackal/gospec/src/gospec"
	l "gotestit/loggregatorclient"
	"net"
)

func LoggregatorClientSpec(c gospec.Context) {
	bufferSize := 4096
	loggregatorClient := l.NewLoggregatorClient("localhost:9876", gosteno.NewLogger("TestLogger"), bufferSize)

	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:9876")
	c.Expect(err, IsNil)
	c.Expect(err, Not(Equals), 1)

	udpListener, err := net.ListenUDP("udp", udpAddr)
	defer udpListener.Close()
	c.Expect(err, IsNil)

	c.Specify("That it sends data", func() {
		expectedOutput := []byte("Important Testmessage")
		loggregatorClient.Send(expectedOutput)

		buffer := make([]byte, bufferSize)
		readCount, _, err := udpListener.ReadFromUDP(buffer)
		c.Expect(err, IsNil)

		received := string(buffer[:readCount])
		c.Expect(string(expectedOutput), Equals, received)
	})
	c.Specify("That it does not send empty data", func() {
		firstMessage := []byte("")
		secondMessage := []byte("hi")
		loggregatorClient.Send(firstMessage)
		loggregatorClient.Send(secondMessage)

		buffer := make([]byte, bufferSize)
		readCount, _, err := udpListener.ReadFromUDP(buffer)
		c.Expect(err, Equals, nil)

		received := string(buffer[:readCount])
		c.Expect(string(secondMessage), Equals, received)
	})
}
