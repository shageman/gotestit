package lc_zen

import (
	"github.com/cloudfoundry/gosteno"
	. "github.com/pranavraja/zen"
	l "gotestit/loggregatorclient"
	"net"
	"testing"
)

func TestLoggregatorClient(t *testing.T) {
	Desc(t, "loggregator client", func(it It) {
		bufferSize := 4096
		loggregatorClient := l.NewLoggregatorClient("localhost:9876", gosteno.NewLogger("TestLogger"), bufferSize)

		udpAddr, err := net.ResolveUDPAddr("udp", "localhost:9876")
		it("does not fail resolving UPD Addr", func(expect Expect) {
			expect(err).ToEqual(nil)
		})

		udpListener, err := net.ListenUDP("udp", udpAddr)
		defer udpListener.Close()
		it("does not fail listening to UPD", func(expect Expect) {
			expect(err).ToEqual(nil)
		})

		it("should send data", func(expect Expect) {
			expectedOutput := []byte("Important Testmessage")
			loggregatorClient.Send(expectedOutput)

			buffer := make([]byte, bufferSize)
			readCount, _, err := udpListener.ReadFromUDP(buffer)
			expect(err).ToEqual(nil)

			received := string(buffer[:readCount])
			expect(string(expectedOutput)).ToEqual(received)
		})
		it("should not send empty data", func(expect Expect) {
			firstMessage := []byte("")
			secondMessage := []byte("hi")
			loggregatorClient.Send(firstMessage)
			loggregatorClient.Send(secondMessage)

			buffer := make([]byte, bufferSize)
			readCount, _, err := udpListener.ReadFromUDP(buffer)
			expect(err).ToEqual(nil)

			received := string(buffer[:readCount])
			expect(string(secondMessage)).ToEqual(received)
		})
	})
}
