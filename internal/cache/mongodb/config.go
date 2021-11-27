package mongodb

type Config struct {
	Uri  string `koanf:"uri"`
	Name string `koanf:"name"`
}
