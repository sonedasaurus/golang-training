package main // main
import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type Conn struct {
	rootDir    string
	workdir    string
	reqUser    string
	user       string
	currentDir string
	granted    bool
	dataType   string
	ctrlConn   *net.TCPConn
	clientAddr *net.TCPAddr
}

func (conn *Conn) writeMessage(code int, message string, v ...interface{}) error {
	msg := fmt.Sprintf(message, v...)
	_, err := fmt.Fprintf(conn.ctrlConn, "%d %s\r\n", code, msg)
	return err
}

func (conn *Conn) handleCommand(line string) {
	tokens := strings.Fields(line)
	opc := strings.ToUpper(tokens[0])
	opr := tokens[1:]

	switch opc {

	// ACCESS CONTROL COMMANDS
	case "USER":
		conn.handleUserCommand(opc, opr)
		return
	case "PASS":
		conn.handlePassCommand(opc, opr)
		return
	case "CWD":
		conn.handleCwdCommand(opc, opr)
		return
	case "PWD":
		conn.handlePwdCommand(opc, opr)
		return
	case "QUIT":
		conn.handleQuitCommand(opc, opr)
		return

	// TRANSFER PARAMETER COMMANDS
	case "PORT":
		// return s.handlePortCommand(opc, opr)
	case "TYPE":
		// return s.handleTypeCommand(opc, opr)
	case "STRU":
		// return s.handleStruCommand(opc, opr)
	case "MODE":
		// return s.handleModeCommand(opc, opr)

	// FTP SERVICE COMMANDS
	case "RETR":
		// return s.handleRetrCommand(opc, opr)

	default:
		conn.writeMessage(500, "%s not understood", opc)
		return
	}
}

func (conn *Conn) handleUserCommand(opc string, opr []string) {
	conn.reqUser = opr[0]
	conn.writeMessage(331, "User name ok, password required")
}

func (conn *Conn) handlePassCommand(opc string, opr []string) {
	ok, err := conn.checkPasswd(conn.user, opr[0])
	if err != nil {
		conn.writeMessage(550, "Checking password error")
		return
	}

	if ok {
		conn.user = conn.reqUser
		conn.reqUser = ""
		conn.writeMessage(230, "Password ok, continue")
	} else {
		conn.writeMessage(530, "Incorrect password, not logged in")
	}
}

func (conn *Conn) handleCwdCommand(opc string, opr []string) {
	path := conn.buildPath(opr[0])
	if f, err := os.Stat(conn.rootDir + path); err != nil || !f.IsDir() {
		conn.writeMessage(550, "Failed to change directory.")
		return
	}
	conn.currentDir = path
	conn.writeMessage(250, "Directory changed to "+path)
}

func (conn *Conn) handlePwdCommand(opc string, opr []string) {
	message := fmt.Sprintf("\"%s\" is the current directory", conn.currentDir)
	conn.writeMessage(257, message)
}

func (conn *Conn) handleRetlCommand(opc string, opr []string) {
}

func (conn *Conn) handleQuitCommand(opc string, opr []string) {
	conn.writeMessage(221, "Goodbye")
	conn.ctrlConn.Close()
}

func (conn *Conn) checkPasswd(user string, pass string) (bool, error) {
	return true, nil
}

func (conn *Conn) buildPath(filename string) string {
	fullPath := ""
	if len(filename) > 0 && filename[0:1] == "/" {
		fullPath = filepath.Clean(filename)
	} else if len(filename) > 0 && filename != "-a" {
		fullPath = filepath.Clean(conn.currentDir + "/" + filename)
	} else {
		fullPath = filepath.Clean(conn.currentDir)
	}
	fullPath = strings.Replace(fullPath, "//", "/", -1)
	return fullPath
}

func startSession(conn *net.TCPConn) {
	var c Conn
	c.ctrlConn = conn
	if rootDir, err := os.Getwd(); err == nil {
		fmt.Println(rootDir)
		c.rootDir = rootDir
	} else {
		log.Fatal(err)
	}

	c.writeMessage(220, "FTP Server Ready")

	go func() {
		defer func() {
			conn.Close()
			log.Println("close controll connection")
		}()

		input := bufio.NewScanner(conn)
		for input.Scan() {
			c.handleCommand(input.Text())
		}
	}()
}

func main() {
	port := flag.Int("port", 21, "port to bind")
	flag.Parse()

	addr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	ctrlListener, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ctrlListener.AcceptTCP()
		if err != nil {
			log.Print(err)
			continue
		}
		startSession(conn)
	}
}
