package config

const (
	swaggerEnabledEnv string = "SWAGGER_ENABLED"
)

type swagger struct {
	Enabled bool `conf:"required"`
}

func (s swagger) GetSwaggerEnabled() bool {
	return s.Enabled
}
