package appuserschemaproperty


type AppUserSchemaPropertyArrayOneOf struct {
	// Enum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_user_schema_property#const AppUserSchemaProperty#const}
	Const *string `field:"required" json:"const" yaml:"const"`
	// Enum title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/app_user_schema_property#title AppUserSchemaProperty#title}
	Title *string `field:"required" json:"title" yaml:"title"`
}

