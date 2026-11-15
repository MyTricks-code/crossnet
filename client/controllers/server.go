package controllers

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	HELLO byte = iota + 1
	FILE_OFFER
	ACCEPT
	REJECT
	FILE_DATA
	TRANSFER_COMPLETE
	TRANSFER_ERROR
)

func ServerListener() (net.Listener, error) {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		return nil, err
	}
	fmt.Println("Server listening on 9000")
	return listener, nil
}

func SpawnConnection() (net.Conn, error) {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func HandleConnection(choice int) error {
	/*
		1. Receive
		2. Send
	*/
	if choice == 1 {
		conn, err := SpawnConnection()
		if err != nil {
			return fmt.Errorf("start connection: %w", err)
		}
		defer conn.Close()
		return FileReceive(conn)

	} else if choice == 2 {
		var wg sync.WaitGroup
		listener, err := ServerListener()
		if err != nil {
			return fmt.Errorf("start listener: %w", err)
		}
		defer listener.Close()

		var expectedClients int
		fmt.Print("Expected number of clients: ")
		if _, err := fmt.Scan(&expectedClients); err != nil {
			return fmt.Errorf("read expected client count: %w", err)
		}
		if expectedClients <= 0 {
			return fmt.Errorf("expected client count must be greater than zero")
		}

		var fileLocation string
		fmt.Println("Enter filePath: ")
		if _, err := fmt.Scan(&fileLocation); err != nil {
			return fmt.Errorf("read file path: %w", err)
		}

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		defer signal.Stop(stop)

		const idleTimeout = 20 * time.Second
		type acceptResult struct {
			conn net.Conn
			err  error
		}

		acceptResults := make(chan acceptResult, 1)
		acceptDone := make(chan struct{})
		defer close(acceptDone)
		go func() {
			for {
				conn, err := listener.Accept()
				select {
				case acceptResults <- acceptResult{conn: conn, err: err}:
				case <-acceptDone:
					if conn != nil {
						_ = conn.Close()
					}
					return
				}
				if err != nil {
					return
				}
			}
		}()

		idleTimer := time.NewTimer(idleTimeout)
		defer idleTimer.Stop()
		countdown := time.NewTicker(time.Second)
		defer countdown.Stop()
		deadline := time.Now().Add(idleTimeout)
		acceptedClients := 0
		var acceptErr error

		for acceptedClients < expectedClients {
			select {
			case result := <-acceptResults:
				if result.err != nil {
					if !errors.Is(result.err, net.ErrClosed) {
						acceptErr = fmt.Errorf("accept connection: %w", result.err)
					}
					acceptedClients = expectedClients
					continue
				}

				acceptedClients++
				wg.Add(1)
				go func(conn net.Conn, fileLocation string) {
					defer wg.Done()
					defer conn.Close()
					if err := FileSend(conn, fileLocation); err != nil {
						fmt.Printf("File send to %s failed: %v\n", conn.RemoteAddr(), err)
						return
					}
					fmt.Printf("File sent successfully to %s\n", conn.RemoteAddr())
				}(result.conn, fileLocation)

				if !idleTimer.Stop() {
					select {
					case <-idleTimer.C:
					default:
					}
				}
				idleTimer.Reset(idleTimeout)
				deadline = time.Now().Add(idleTimeout)
			case <-idleTimer.C:
				fmt.Println("No new connection for 10 seconds; closing listener.")
				acceptedClients = expectedClients
			case <-countdown.C:
				remaining := int((time.Until(deadline) + time.Second - 1) / time.Second)
				if remaining < 0 {
					remaining = 0
				}
				fmt.Printf("Closing listener in %d seconds without a new connection.\n", remaining)
			case <-stop:
				fmt.Println("Interrupt received; closing listener.")
				acceptedClients = expectedClients
			}
		}

		if err := listener.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
			return fmt.Errorf("close listener: %w", err)
		}
		wg.Wait()
		fmt.Println("All active transfers finished.")
		return acceptErr
	}
	return fmt.Errorf("invalid connection choice: %d", choice)
}
