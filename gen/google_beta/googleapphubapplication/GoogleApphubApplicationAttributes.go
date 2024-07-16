package googleapphubapplication


type GoogleApphubApplicationAttributes struct {
	// business_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_application#business_owners GoogleApphubApplication#business_owners}
	BusinessOwners interface{} `field:"optional" json:"businessOwners" yaml:"businessOwners"`
	// criticality block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_application#criticality GoogleApphubApplication#criticality}
	Criticality *GoogleApphubApplicationAttributesCriticality `field:"optional" json:"criticality" yaml:"criticality"`
	// developer_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_application#developer_owners GoogleApphubApplication#developer_owners}
	DeveloperOwners interface{} `field:"optional" json:"developerOwners" yaml:"developerOwners"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_application#environment GoogleApphubApplication#environment}
	Environment *GoogleApphubApplicationAttributesEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// operator_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_application#operator_owners GoogleApphubApplication#operator_owners}
	OperatorOwners interface{} `field:"optional" json:"operatorOwners" yaml:"operatorOwners"`
}

