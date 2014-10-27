package shared

type Engine interface {
	Discover(path string) bool
	GetDockerfileTemplate() string
}
