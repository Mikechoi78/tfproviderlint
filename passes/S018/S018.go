// Package S018 defines an Analyzer that checks for
// Schema that should prefer TypeList with MaxItems 1
package S018

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/schemaschema"
)

const Doc = `check for Schema that should prefer TypeList with MaxItems 1

The S018 analyzer reports cases of schema including MaxItems 1 and TypeSet
that should be simplified.`

const analyzerName = "S018"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		schemaschema.Analyzer,
		commentignore.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemas := pass.ResultOf[schemaschema.Analyzer].([]*ast.CompositeLit)
	for _, schema := range schemas {
		if ignorer.ShouldIgnore(analyzerName, schema) {
			continue
		}

		if !terraformtype.HelperSchemaTypeSchemaContainsTypes(schema, pass.TypesInfo, terraformtype.SchemaValueTypeSet) {
			continue
		}

		maxItems := terraformtype.HelperSchemaTypeSchemaMaxItems(schema)

		if maxItems == nil || *maxItems != 1 {
			continue
		}

		switch t := schema.Type.(type) {
		default:
			pass.Reportf(schema.Lbrace, "%s: schema should use TypeList with MaxItems 1", analyzerName)
		case *ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(), "%s: schema should use TypeList with MaxItems 1", analyzerName)
		}
	}

	return nil, nil
}
