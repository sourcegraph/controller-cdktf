package groupschemaproperty


type GroupSchemaPropertyOneOf struct {
	// Enum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/group_schema_property#const GroupSchemaProperty#const}
	Const *string `field:"required" json:"const" yaml:"const"`
	// Enum title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/group_schema_property#title GroupSchemaProperty#title}
	Title *string `field:"required" json:"title" yaml:"title"`
}

