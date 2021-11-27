package redis

type Config struct {
	Address  string `koanf:"address"`  // host:port
	Password string `koanf:"password"` // "" means no password
	DB       int    `koanf:"db"`       // 0 means default DB
}
