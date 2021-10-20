package ptable

import "github.com/bitwormhole/starter/application"

const (
	myName     = "github.com/bitwormhole/propertydb"
	myVersion  = "v0.0.1"
	myRevision = 1
)

// Module 导出本模块
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(myName).Version(myVersion).Revision(myRevision)

	return mb.Create()
}
