# Redis-like Key-Value Store

## Overview
This project is an in-memory key-value store similar to Redis, built using Go. It supports basic Redis commands and can be accessed via a TCP server.

## Features
- **Basic Commands**: `SET`, `GET`, `DEL`
- **Numeric Operations**: `INCR`, `INCRBY`
- **Transactions**: `MULTI`, `EXEC`, `DISCARD`
- **Compaction**: `COMPACT`
- **Multi-Client Support**
- **Database Selection**: `SELECT`
- **TCP Server Support**

## Installation
### Prerequisites
- Go 1.18+
- A terminal or command prompt

### Clone the Repository
```sh
git clone https://github.com/your-repo/redis-like-db.git
cd redis-like-db
```

### Build and Run
```sh
go run main.go
```

## Usage
### Connecting to the Server
Use one of the following methods to connect to the server:

#### **Telnet (Windows/Linux/macOS)**
```sh
telnet localhost 9736
```

#### **Netcat (Linux/macOS)**
```sh
nc localhost 9736
```

#### **PowerShell (Windows)**
```powershell
$tcpClient = New-Object System.Net.Sockets.TcpClient("localhost", 9736)
$stream = $tcpClient.GetStream()
$writer = New-Object System.IO.StreamWriter($stream)
$reader = New-Object System.IO.StreamReader($stream)

$writer.WriteLine("SET name John")
$writer.Flush()
$reader.ReadLine()
```

### Example Commands
```sh
SET name John
> OK

GET name
> "John"

INCR counter
> (integer) 1

DEL name
> (integer) 1

MULTI
> OK

INCR counter
> QUEUED

SET foo bar
> QUEUED

EXEC
> 1) (integer) 2
   2) OK

COMPACT
> SET counter 2
  SET foo bar
```

## Error Handling
The server follows Redis-like error messages. Some examples:
```sh
SET surname foo bar
> (error) ERR syntax error

INCR age
> (error) ERR value is not an integer or out of range

SELECT 20
> (error) ERR DB index is out of range
```

## Graceful Shutdown
To stop the server, use **Ctrl+C** or send a `SIGINT` signal.