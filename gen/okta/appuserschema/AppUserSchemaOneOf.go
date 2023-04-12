package appuserschema


type AppUserSchemaOneOf struct {
	// Enum value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_user_schema#const AppUserSchema#const}
	Const *string `field:"required" json:"const" yaml:"const"`
	// Enum title.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_user_schema#title AppUserSchema#title}
	Title *string `field:"required" json:"title" yaml:"title"`
}

