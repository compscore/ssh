package main

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

func Run(ctx context.Context, target string, command string, expectedOutput string, username string, password string) (bool, string) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {
		client, err := ssh.Dial("tcp", target, config)
		if err != nil {
			errChan <- err
			return
		}
		defer client.Close()

		session, err := client.NewSession()
		if err != nil {
			errChan <- err
			return
		}
		defer session.Close()

		output, err := session.CombinedOutput(command)
		if err != nil {
			errChan <- err
			return
		}

		if string(output) != strings.TrimSpace(expectedOutput) {
			errChan <- fmt.Errorf("expected output %s but got %s", expectedOutput, string(output))
			return
		}

		errChan <- nil
	}()

	select {
	case <-ctx.Done():
		return false, fmt.Sprintf("Timeout exceeded; err: %s", ctx.Err())
	case err := <-errChan:
		if err != nil {
			return false, fmt.Sprintf("Encountered error: %s", err)
		}

		return true, ""
	}
}
