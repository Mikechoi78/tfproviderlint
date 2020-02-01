package V006

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/bflad/tfproviderlint/passes/selectorexpr/helper/validation/validatelistuniquestrings"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V006",
	validatelistuniquestrings.Analyzer,
	validation.PackageName,
	validation.FuncNameValidateListUniqueStrings,
	validation.PackageName,
	validation.FuncNameListOfUniqueStrings,
)
