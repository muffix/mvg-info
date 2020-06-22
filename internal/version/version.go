package version

// these global variables are set by the build pipeline
// see build/package/api/Dockerfile.
var (
	revision   string //nolint:gochecknoglobals
	pipelineID string //nolint:gochecknoglobals
	buildDate  string //nolint:gochecknoglobals
)
