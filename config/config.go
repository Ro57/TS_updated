package config

type Config struct {
	RpcPort  string `yaml:"rpc_port" json:"rpc_port"`
	HttpPort string `yaml:"http_port" json:"http_port"`
	Domain   string `yaml:"domain" json:"domain"`
}
