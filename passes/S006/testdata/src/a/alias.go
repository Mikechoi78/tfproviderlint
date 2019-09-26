package a

import (
	s "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema of TypeMap should include Elem"
		Type: s.TypeMap,
	}

	_ = map[string]*s.Schema{
		"name": { // want "schema of TypeMap should include Elem"
			Type: s.TypeMap,
		},
	}
}
