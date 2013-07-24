package lc_go_spec

import (
	. "github.com/bmatsuo/go-spec/spec"
	"github.com/cloudfoundry/gosteno"
	l "gotestit/loggregatorclient"
	"net"
	"testing"
)

func TestLoggregatorClient(t *testing.T) {
	s := NewSpecTest(t)

	s.Spec(1, Should, Equal, 2)
	s.Spec(1, Should, Not, Equal, 2)

	bufferSize := 4096
	loggregatorClient := l.NewLoggregatorClient("localhost:9876", gosteno.NewLogger("TestLogger"), bufferSize)

	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:9876")
	s.Spec(err, Should, Equal, nil)

	udpListener, err := net.ListenUDP("udp", udpAddr)
	s.Spec(err, Should, Equal, nil)

	//Describe, They, It are synonyms
	s.Describe("the loggregator client", func() {
		s.It("sends data", func() {
			expectedOutput := []byte("Important Testmessage")
			loggregatorClient.Send(expectedOutput)

			buffer := make([]byte, bufferSize)
			readCount, _, err := udpListener.ReadFromUDP(buffer)
			s.Spec(err, Should, Equal, nil)

			received := string(buffer[:readCount])
			s.Spec(string(expectedOutput), Should, Equal, received)
		})
		s.It("does not send empty data", func() {
			firstMessage := []byte("")
			secondMessage := []byte("hi")
			loggregatorClient.Send(firstMessage)
			loggregatorClient.Send(secondMessage)

			buffer := make([]byte, bufferSize)
			readCount, _, err := udpListener.ReadFromUDP(buffer)
			s.Spec(err, Should, Equal, nil)

			received := string(buffer[:readCount])
			s.Spec(string(secondMessage), Should, Equal, received)
		})
	})

	udpListener.Close()
}
