package db

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

type Agent struct {
	ID       string    `json:"id"`
	Hostname string    `json:"hostname"`
	IP       string    `json:"ip"`
	LastSeen time.Time `json:"last_seen"`
}

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "x.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		PRAGMA journal_mode=WAL;
		PRAGMA foreign_keys=ON;

		CREATE TABLE IF NOT EXISTS agents (
			id TEXT PRIMARY KEY,
			hostname TEXT,
			ip TEXT,
			last_seen DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			agent_id TEXT,
			command TEXT NOT NULL,
			status TEXT DEFAULT 'pending',
			result TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(agent_id) REFERENCES agents(id)
		);
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func RegisterPing(db *sql.DB, id, hostname, ip string) error {
	_, err := db.Exec(`
		INSERT INTO agents (id, hostname, ip, last_seen)
		VALUES (?, ?, ?, CURRENT_TIMESTAMP)
		ON CONFLICT(id) DO UPDATE SET 
			hostname = excluded.hostname,
			ip = excluded.ip,
			last_seen = CURRENT_TIMESTAMP
	`, id, hostname, ip)
	return err
}

func GetPendingTask(db *sql.DB, agentID string) (int, string, error) {
	var taskID int
	var command string

	err := db.QueryRow(`
		SELECT id, command 
		FROM tasks 
		WHERE agent_id = ? AND status = 'pending' 
		ORDER BY id ASC LIMIT 1
	`, agentID).Scan(&taskID, &command)

	if err == sql.ErrNoRows {
		return 0, "", nil
	}
	return taskID, command, err
}

func UpdateTaskStatus(db *sql.DB, taskID int, result string) error {
	_, err := db.Exec(`
		UPDATE tasks 
		SET status = 'completed', result = ? 
		WHERE id = ?
	`, result, taskID)
	return err
}

func GetAllAgents(db *sql.DB) ([]Agent, error) {
	rows, err := db.Query("SELECT id, hostname, ip, last_seen FROM agents ORDER BY last_seen DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agents []Agent
	for rows.Next() {
		var a Agent
		if err := rows.Scan(&a.ID, &a.Hostname, &a.IP, &a.LastSeen); err == nil {
			agents = append(agents, a)
		}
	}
	return agents, nil
}
