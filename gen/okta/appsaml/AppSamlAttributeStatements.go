package appsaml


type AppSamlAttributeStatements struct {
	// The reference name of the attribute statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#name AppSaml#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of group attribute filter. Valid values are: `STARTS_WITH`, `EQUALS`, `CONTAINS`, or `REGEX`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#filter_type AppSaml#filter_type}
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
	// Filter value to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#filter_value AppSaml#filter_value}
	FilterValue *string `field:"optional" json:"filterValue" yaml:"filterValue"`
	// The attribute namespace. It can be set to `urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified`, `urn:oasis:names:tc:SAML:2.0:attrname-format:uri`, or `urn:oasis:names:tc:SAML:2.0:attrname-format:basic`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#namespace AppSaml#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The type of attribute statements object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#type AppSaml#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#values AppSaml#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

