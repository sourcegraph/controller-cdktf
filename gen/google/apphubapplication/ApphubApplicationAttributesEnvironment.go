package apphubapplication


type ApphubApplicationAttributesEnvironment struct {
	// Environment type. Possible values: ["PRODUCTION", "STAGING", "TEST", "DEVELOPMENT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/apphub_application#type ApphubApplication#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

