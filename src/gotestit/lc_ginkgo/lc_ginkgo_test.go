package lc_ginkgo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "gotestit/loggregatorclient"

	"github.com/cloudfoundry/gosteno"
	"net"
)

var _ = Describe("LoggregatorClient", func() {
	var (
		loggregatorClient  LoggregatorClient
		bufferSize         int
		loggregatorAddress string
		udpListener        *net.UDPConn
	)

	BeforeEach(func() {
		bufferSize = 4096
		loggregatorAddress = "localhost:9876"
		loggregatorClient = NewLoggregatorClient(loggregatorAddress, gosteno.NewLogger("TestLogger"), bufferSize)

		udpAddr, err := net.ResolveUDPAddr("udp", loggregatorAddress)
		Ω(err).ShouldNot(HaveOccurred())

		udpListener, err = net.ListenUDP("udp", udpAddr)
		Ω(err).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		udpListener.Close()
	})

	readFromUDP := func() []byte {
		buffer := make([]byte, bufferSize)
		readCount, _, err := udpListener.ReadFromUDP(buffer)
		Ω(err).ShouldNot(HaveOccurred())

		return buffer[:readCount]
	}

	Describe("Sending data", func() {
		var data []byte

		BeforeEach(func() {
			data = []byte("Important Testmessage")
		})

		Context("when the data is not empty", func() {
			It("should send data", func() {
				loggregatorClient.Send(data)

				Ω(readFromUDP()).Should(Equal(data))
			})
		})

		Context("when the data is empty", func() {
			It("should not send anything", func() {
				loggregatorClient.Send([]byte(""))
				loggregatorClient.Send(data)

				Ω(readFromUDP()).Should(Equal(data))
			})
		})
	})
})
