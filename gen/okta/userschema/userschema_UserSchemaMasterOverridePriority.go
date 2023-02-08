package userschema


type UserSchemaMasterOverridePriority struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_schema#value UserSchema#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_schema#type UserSchema#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

