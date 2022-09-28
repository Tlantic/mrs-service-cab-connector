package auth

const (
	ResourceKey = "Auth"
)

type ExternalEndpoint struct {
	BaseURI   string               `json:"base_uri" yaml:"BaseURI"`
	Resources map[string]*Resource `json:"resources" yaml:"Resources"`
}

type Resource struct {
	Path string `json:"path" yaml:"Path"`
}
