package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

const defaultDBCount = 16

type Database struct {
	data  map[string]string
	mutex sync.RWMutex
}

type Server struct {
	databases [defaultDBCount]*Database
	clients   map[net.Conn]int
	mutex     sync.Mutex
}

func NewServer() *Server {
	s := &Server{clients: make(map[net.Conn]int)}
	for i := 0; i < defaultDBCount; i++ {
		s.databases[i] = &Database{data: make(map[string]string)}
	}
	return s
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	dbIndex := 0 // Default DB
	reader := bufio.NewReader(conn)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Client disconnected:", conn.RemoteAddr())
			s.mutex.Lock()
			delete(s.clients, conn)
			s.mutex.Unlock()
			return
		}
		input = strings.TrimSpace(input)
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}
		command := strings.ToUpper(args[0])
		response := s.executeCommand(command, args[1:], &dbIndex)
		conn.Write([]byte(response + "\n"))
	}
}

func (s *Server) executeCommand(command string, args []string, dbIndex *int) string {
	db := s.databases[*dbIndex]
	switch command {
	case "SET":
		if len(args) < 2 {
			return "(error) ERR wrong number of arguments for 'set' command"
		}
		key, value := args[0], strings.Join(args[1:], " ")
		db.mutex.Lock()
		db.data[key] = value
		db.mutex.Unlock()
		return "OK"
	case "GET":
		if len(args) != 1 {
			return "(error) ERR wrong number of arguments for 'get' command"
		}
		db.mutex.RLock()
		val, exists := db.data[args[0]]
		db.mutex.RUnlock()
		if !exists {
			return "(nil)"
		}
		return fmt.Sprintf("\"%s\"", val)
	case "DEL":
		if len(args) != 1 {
			return "(error) ERR wrong number of arguments for 'del' command"
		}
		db.mutex.Lock()
		_, exists := db.data[args[0]]
		if exists {
			delete(db.data, args[0])
		}
		db.mutex.Unlock()
		if exists {
			return "(integer) 1"
		}
		return "(integer) 0"
	case "INCR":
		if len(args) != 1 {
			return "(error) ERR wrong number of arguments for 'incr' command"
		}
		key := args[0]
		db.mutex.Lock()
		val, exists := db.data[key]
		if !exists {
			db.data[key] = "1"
			db.mutex.Unlock()
			return "(integer) 1"
		}
		intVal, err := strconv.Atoi(val)
		if err != nil {
			db.mutex.Unlock()
			return "(error) ERR value is not an integer or out of range"
		}
		intVal++
		db.data[key] = strconv.Itoa(intVal)
		db.mutex.Unlock()
		return fmt.Sprintf("(integer) %d", intVal)
	case "SELECT":
		if len(args) != 1 {
			return "(error) ERR wrong number of arguments for 'select' command"
		}
		dbNum, err := strconv.Atoi(args[0])
		if err != nil || dbNum < 0 || dbNum >= defaultDBCount {
			return "(error) ERR DB index is out of range"
		}
		*dbIndex = dbNum
		return "OK"
	case "COMPACT":
		db.mutex.RLock()
		var compacted []string
		for k, v := range db.data {
			compacted = append(compacted, fmt.Sprintf("SET %s %s", k, v))
		}
		db.mutex.RUnlock()
		if len(compacted) == 0 {
			return "(nil)"
		}
		return strings.Join(compacted, "\n")
	default:
		return "(error) ERR unknown command"
	}
}

func main() {
	port := "9736"
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()
	fmt.Println("Redis-like server started on port", port)

	srv := NewServer()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		fmt.Println("Shutting down server...")
		listener.Close()
		os.Exit(0)
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		log.Println("New client connected:", conn.RemoteAddr())
		go srv.handleConnection(conn)
	}
}
