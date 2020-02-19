package config

type CfgStruct struct {
	Mysqls []struct {
		MaxOdleConns int64  `json:"MaxOdleConns"`
		MaxOpenConns int64  `json:"MaxOpenConns"`
		Database     string `json:"database"`
		Host         string `json:"host"`
		Password     string `json:"password"`
		Username     string `json:"username"`
	} `json:"mysqls"`
	Servers struct {
		Debug  int64 `json:"debug"`
		Server []struct {
			Blacklist   []string `json:"blacklist"`
			Listen      int64    `json:"listen"`
			Servername  []string `json:"servername"`
			Ssl         string   `json:"ssl"`
			SslCertfile string   `json:"ssl_certfile"`
			SslKeyfile  string   `json:"ssl_keyfile"`
			Whitelist   []string `json:"whitelist"`
		} `json:"server"`
	} `json:"servers"`
	Ws struct {
		Host string `json:"host"`
		Path string `json:"path"`
	} `json:"ws"`
}
