package lc_testify

import (
	"github.com/cloudfoundry/gosteno"
	"github.com/stretchr/testify/assert"
	l "gotestit/loggregatorclient"
	"net"
	"testing"
)

func TestSend(t *testing.T) {
	bufferSize := 4096
	loggregatorClient := l.NewLoggregatorClient("localhost:9876", gosteno.NewLogger("TestLogger"), bufferSize)

	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:9876")
	assert.NoError(t, err)

	udpListener, err := net.ListenUDP("udp", udpAddr)
	defer udpListener.Close()
	assert.NoError(t, err)

	expectedOutput := []byte("Important Testmessage")
	loggregatorClient.Send(expectedOutput)

	buffer := make([]byte, bufferSize)
	readCount, _, err := udpListener.ReadFromUDP(buffer)
	assert.NoError(t, err)

	received := string(buffer[:readCount])
	assert.Equal(t, string(expectedOutput), received)
}

func TestDontSendEmptyData(t *testing.T) {
	bufferSize := 4096
	loggregatorClient := l.NewLoggregatorClient("localhost:9876", gosteno.NewLogger("TestLogger"), bufferSize)

	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:9876")
	assert.NoError(t, err)

	udpListener, err := net.ListenUDP("udp", udpAddr)
	defer udpListener.Close()
	assert.NoError(t, err)

	firstMessage := []byte("")
	secondMessage := []byte("hi")
	loggregatorClient.Send(firstMessage)
	loggregatorClient.Send(secondMessage)

	buffer := make([]byte, bufferSize)
	readCount, _, err := udpListener.ReadFromUDP(buffer)

	assert.NoError(t, err)

	received := string(buffer[:readCount])
	assert.Equal(t, string(secondMessage), received)
}
