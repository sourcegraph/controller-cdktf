package userschema


type UserSchemaMasterOverridePriority struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/user_schema#value UserSchema#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/user_schema#type UserSchema#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

