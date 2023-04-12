package groupschemaproperty


type GroupSchemaPropertyMasterOverridePriority struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#value GroupSchemaProperty#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#type GroupSchemaProperty#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

