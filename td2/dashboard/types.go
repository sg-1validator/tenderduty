package dash

type ChainStatus struct {
	MsgType      string       `json:"msgType"`
	Name         string       `json:"name"`
	ChainId      string       `json:"chain_id"`
	Moniker      string       `json:"moniker"`
	Bonded       bool         `json:"bonded"`
	Jailed       bool         `json:"jailed"`
	Tombstoned   bool         `json:"tombstoned"`
	Missed       int64        `json:"missed"`
	Window       int64        `json:"window"`
	Nodes        int          `json:"nodes"`
	HealthyNodes int          `json:"healthy_nodes"`
	NodeInfo     []NodeStatus `json:"nodeInfo"`
	ActiveAlerts int          `json:"active_alerts"`
	Height       int64        `json:"height"`
	LastError    string       `json:"last_error"`

	Blocks []int `json:"blocks"`
}

type NodeStatus struct {
	NodeName      string `json:"nodeName"`
	NodeAvailable bool   `json:"nodeStatus"`
}

type LogMessage struct {
	MsgType string `json:"msgType"`
	Ts      int64  `json:"ts"`
	Msg     string `json:"msg"`
}
