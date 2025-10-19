package model

type Config struct {
	ServerIp   string `json:"serverIp"`
	ServerPort uint64 `json:"serverPort"`
	WsPort     uint64 `json:"wsPort"`
	JwtSecret  string `json:"jwtSecret"`
}

type GoogleOauth struct {
	Web struct {
		ClientId                string `json:"client_id"`
		ProjectId               string `json:"project_id"`
		OauthUri                string `json:"oauth_uri"`
		TokenUri                string `json:"token_uri"`
		AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
		ClientSecret            string `json:"client_secret"`
	} `json:"web"`
}
