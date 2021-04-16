package structs

type CreateChannelOption struct {
	CommonOption
	ChannelConfigPath  string `json:"channel_config_path"`
}

type JoinChannelOption struct {
	CommonOption
	Peers              []string `json:"peers"`
}

// 公共属性
type CommonOption struct {
	ChannelName        string `json:"channel_name" binding:"Required"`
	Orgname            string `json:"orgnization_name"`
}