package ptable

import (
	"embed"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myName     = "github.com/bitwormhole/ptable"
	myVersion  = "v0.0.4"
	myRevision = 4
)

//go:embed src/main/resources
var theMainRes embed.FS

// Module 导出本模块
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(myName).Version(myVersion).Revision(myRevision)
	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))

	return mb.Create()
}
