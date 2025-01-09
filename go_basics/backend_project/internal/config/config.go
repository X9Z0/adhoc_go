package config

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env string
	HTTPServer
}
