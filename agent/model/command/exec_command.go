package command

type ExecCommand struct {
	Type      string   `json:"type"`
	WorkDir   string   `json:"work_dir"`
	Command   string   `json:"command"`
	Args      []string `json:"args"`
	Timeout   int      `json:"timeout"`
	RequestID string   `json:"request_id"`
}

type ExecResult struct {
	RequestID string `json:"request_id"`
	Code      int    `json:"code"`
	Stdout    string `json:"stdout"`
	Stderr    string `json:"stderr"`
}
