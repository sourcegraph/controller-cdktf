package samlapp


type SamlAppAttributeStatements struct {
	// The reference name of the attribute statement.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#name SamlApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of group attribute filter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#filter_type SamlApp#filter_type}
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
	// Filter value to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#filter_value SamlApp#filter_value}
	FilterValue *string `field:"optional" json:"filterValue" yaml:"filterValue"`
	// The name format of the attribute.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#namespace SamlApp#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The type of attribute statements object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#type SamlApp#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#values SamlApp#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

