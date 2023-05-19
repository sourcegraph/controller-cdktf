package userschema


type UserSchemaArrayOneOf struct {
	// Enum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/user_schema#const UserSchema#const}
	Const *string `field:"required" json:"const" yaml:"const"`
	// Enum title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/user_schema#title UserSchema#title}
	Title *string `field:"required" json:"title" yaml:"title"`
}

