package swagger

import "embed"

//go:embed openapi.yml index.html
var FS embed.FS
