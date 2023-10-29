package ping

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

func Run(ctx context.Context, target string, command string, expectedOutput string, username string, password string, options map[string]interface{}) (bool, string) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if !strings.Contains(target, ":") {
		target = fmt.Sprintf("%s:22", target)
	}

	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)
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

		outputString := strings.TrimSpace(string(output))
		expectedOutputString := strings.TrimSpace(expectedOutput)

		if outputString != expectedOutputString {
			errChan <- fmt.Errorf("expected output \"%s\" but got \"%s\"", expectedOutputString, outputString)
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
