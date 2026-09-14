package middleware

type CORS interface {
	GetAllowedOrigins() []string
	GetAllowedMethods() []string
	GetAllowedHeaders() []string
	GetallowCredentials() string
}
