package main

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"
)

func startTestServer() {
	go main()
	time.Sleep(500 * time.Millisecond) // Allow server to start
}

func TestRedisServerCommands(t *testing.T) {
	startTestServer()

	conn, err := net.Dial("tcp", "localhost:9736")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	tests := []struct {
		name     string
		command  string
		expected []string // Changed to slice to allow order-independent matching
	}{
		{"SET command", "SET key1 value1\n", []string{"OK"}},
		{"SET command", "SET key2 value2\n", []string{"OK"}},
		{"INCR command", "INCR counter\n", []string{"(integer) 1"}},
		{"INCR command", "INCR counter\n", []string{"(integer) 2"}},
		{"COMPACT command", "COMPACT\n", []string{"SET key1 value1", "SET key2 value2", "SET counter 2"}},
	}

	reader := bufio.NewReader(conn)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := conn.Write([]byte(tt.command))
			if err != nil {
				t.Fatalf("Failed to send command: %v", err)
			}
			resp, err := reader.ReadString('\n')
			if err != nil {
				t.Fatalf("Failed to read response: %v", err)
			}
			resp = strings.TrimSpace(resp)

			// For COMPACT, check that all expected key-value pairs exist somewhere in the response
			if tt.name == "COMPACT command" {
				for _, expected := range tt.expected {
					if !strings.Contains(resp, expected) {
						t.Errorf("Expected COMPACT output to contain: %s, got: %s", expected, resp)
					}
				}
			} else {
				// For other commands, direct match
				if resp != tt.expected[0] {
					t.Errorf("Expected: %s, got: %s", tt.expected[0], resp)
				}
			}
		})
	}
}
