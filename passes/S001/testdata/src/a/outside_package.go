package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Type: schema.TypeList,
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
	}
}
