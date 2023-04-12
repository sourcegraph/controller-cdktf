package userschemaproperty


type UserSchemaPropertyMasterOverridePriority struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_schema_property#value UserSchemaProperty#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_schema_property#type UserSchemaProperty#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

