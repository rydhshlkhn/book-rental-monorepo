// Code generated by candi v1.17.15.

package shared

// Environment additional in this service
type Environment struct {
	// more additional environment with struct tag is environment key example:
	// ExampleHost string `env:"EXAMPLE_HOST"`
	AuthServiceKey  string `env:"AUTH_SERVICE_KEY"`
	AuthServiceHost string `env:"AUTH_SERVICE_HOST"`
	
}

var sharedEnv Environment

// GetEnv get global additional environment
func GetEnv() Environment {
	return sharedEnv
}

// SetEnv get global additional environment
func SetEnv(env Environment) {
	sharedEnv = env
}
