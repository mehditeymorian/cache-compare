package mongodb

type Config struct {
	Uri      string `koanf:"uri"`
	Name     string `koanf:"name"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
}
