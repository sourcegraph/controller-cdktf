package userschemaproperty


type UserSchemaPropertyOneOf struct {
	// Enum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/user_schema_property#const UserSchemaProperty#const}
	Const *string `field:"required" json:"const" yaml:"const"`
	// Enum title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/user_schema_property#title UserSchemaProperty#title}
	Title *string `field:"required" json:"title" yaml:"title"`
}

