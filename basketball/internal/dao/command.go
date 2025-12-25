package dao

type CommandArgs struct {
	HomeDir    string `yaml:"homeDir"`
	ConfigFile string `yaml:"configFile"`
}

type PageInfo struct {
	Page     int64 `json:"page" form:"page"`
	PageSize int64 `json:"pageSize" form:"pageSize"`
}
