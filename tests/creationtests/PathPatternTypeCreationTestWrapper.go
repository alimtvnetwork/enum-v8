package creationtests

import "github.com/alimtvnetwork/enum-v1/pathpatterntype"

type PathPatternTypeCreationTestWrapper struct {
	PathType                      pathpatterntype.Variant
	Name, FullName, CurlyFullName string
	CompiledTemplateFullPath      string
	AssociatedTemplatePaths       []string
}
