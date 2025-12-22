package referencetable


type ReferenceTableSchema struct {
	// The name of the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/reference_table#name ReferenceTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the column. See https://docs.observeinc.com/en/latest/content/query-language-reference/OPALUserGuideTypesOperators.html for options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/reference_table#type ReferenceTable#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

