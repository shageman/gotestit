package lc_testing

import (
	"github.com/cloudfoundry/gosteno"
	l "gotestit/loggregatorclient"
	"net"
	"testing"
)

func TestSend(t *testing.T) {
	bufferSize := 4096
	loggregatorClient := l.NewLoggregatorClient("localhost:9876", gosteno.NewLogger("TestLogger"), bufferSize)

	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:9876")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	udpListener, err := net.ListenUDP("udp", udpAddr)
	defer udpListener.Close()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	expectedOutput := []byte("Important Testmessage")
	loggregatorClient.Send(expectedOutput)

	buffer := make([]byte, bufferSize)
	readCount, _, err := udpListener.ReadFromUDP(buffer)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	received := string(buffer[:readCount])
	if string(expectedOutput) != received {
		t.Errorf("Expected %s to equal %s", string(expectedOutput), received)
	}
}

func TestDontSendEmptyData(t *testing.T) {
	bufferSize := 4096
	loggregatorClient := l.NewLoggregatorClient("localhost:9876", gosteno.NewLogger("TestLogger"), bufferSize)

	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:9876")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	udpListener, err := net.ListenUDP("udp", udpAddr)
	defer udpListener.Close()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	firstMessage := []byte("")
	secondMessage := []byte("hi")
	loggregatorClient.Send(firstMessage)
	loggregatorClient.Send(secondMessage)

	buffer := make([]byte, bufferSize)
	readCount, _, err := udpListener.ReadFromUDP(buffer)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	received := string(buffer[:readCount])
	if string(secondMessage) != received {
		t.Errorf("Expected %s to equal %s", string(secondMessage), received)
	}
}
