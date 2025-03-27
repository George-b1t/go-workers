package client

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"
)

// Sobrescreve temporariamente a função dial para testes
var dial = net.Dial

// mockConn simula uma conexão de rede para testes

type mockConn struct {
	readBuffer  *strings.Reader
	writeBuffer *strings.Builder
}

func (m *mockConn) Read(b []byte) (n int, err error) {
	return m.readBuffer.Read(b)
}

func (m *mockConn) Write(b []byte) (n int, err error) {
	return m.writeBuffer.Write(b)
}

func (m *mockConn) Close() error {
	return nil
}

func (m *mockConn) LocalAddr() net.Addr {
	return nil
}

func (m *mockConn) RemoteAddr() net.Addr {
	return nil
}

func (m *mockConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *mockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *mockConn) SetWriteDeadline(t time.Time) error {
	return nil
}

// TestNewClient verifica se NewClient retorna uma instância válida
func TestNewClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Fatal("Expected NewClient to return a non-nil Client")
	}
	if client.inputReader == nil {
		t.Fatal("Expected inputReader to be initialized")
	}
}

// TestConnect testa se o cliente conecta corretamente a um servidor mockado
func TestConnect(t *testing.T) {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer listener.Close()

	mock := &mockConn{
		readBuffer:  strings.NewReader(""),
		writeBuffer: &strings.Builder{},
	}
	dial = func(network, address string) (net.Conn, error) {
		return mock, nil
	}
	defer func() { dial = net.Dial }()

	client := NewClient()
	err = client.Connect("localhost:12345")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if client.conn == nil {
		t.Fatal("Expected client.conn to be initialized")
	}

	if client.serverReader == nil {
		t.Fatal("Expected serverReader to be initialized")
	}
}

// TestSendTask verifica se uma tarefa é corretamente enviada para a conexão
func TestSendTask(t *testing.T) {
	mock := &mockConn{
		readBuffer:  strings.NewReader(""),
		writeBuffer: &strings.Builder{},
	}
	client := &Client{
		conn: mock,
	}

	task := "uppercase:hello"
	err := client.SendTask(task)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if mock.writeBuffer.String() != task+"\n" {
		t.Fatalf("Expected task to be written to connection, got %v", mock.writeBuffer.String())
	}
}

// TestReceiveResult testa se o cliente consegue receber a resposta corretamente
func TestReceiveResult(t *testing.T) {
	expectedResult := "HELLO"
	mock := &mockConn{
		readBuffer:  strings.NewReader(expectedResult + "\n"),
		writeBuffer: &strings.Builder{},
	}
	client := &Client{
		serverReader: bufio.NewReader(mock),
	}

	result, err := client.ReceiveResult()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != expectedResult {
		t.Fatalf("Expected result %v, got %v", expectedResult, result)
	}
}
