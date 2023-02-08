package userschema


type UserSchemaArrayOneOf struct {
	// Enum value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_schema#const UserSchema#const}
	Const *string `field:"required" json:"const" yaml:"const"`
	// Enum title.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_schema#title UserSchema#title}
	Title *string `field:"required" json:"title" yaml:"title"`
}

