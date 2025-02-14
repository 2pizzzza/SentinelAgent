package logs

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type SystemdLog struct {
	Unit       string `json:"_SYSTEMD_UNIT"`
	Message    string `json:"MESSAGE"`
	Priority   string `json:"PRIORITY"`
	Timestamp  string `json:"__REALTIME_TIMESTAMP"`
	BootID     string `json:"_BOOT_ID"`
	PID        string `json:"_PID"`
	UID        string `json:"_UID"`
	Executable string `json:"_EXE"`
}

func GetSystemdLogs() ([]SystemdLog, error) {
	cmd := exec.Command("journalctl", "-o", "json", "-n", "100")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute journalctl: %v", err)
	}

	logStrings := strings.Split(string(output), "\n")
	var logs []SystemdLog

	for _, logStr := range logStrings {
		if logStr == "" {
			continue
		}

		var log SystemdLog
		if err := json.Unmarshal([]byte(logStr), &log); err != nil {
			return nil, fmt.Errorf("failed to unmarshal log entry: %v", err)
		}
		logs = append(logs, log)
	}

	return logs, nil
}
