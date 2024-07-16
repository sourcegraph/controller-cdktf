package googleapphubservice


type GoogleApphubServiceAttributes struct {
	// business_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_service#business_owners GoogleApphubService#business_owners}
	BusinessOwners interface{} `field:"optional" json:"businessOwners" yaml:"businessOwners"`
	// criticality block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_service#criticality GoogleApphubService#criticality}
	Criticality *GoogleApphubServiceAttributesCriticality `field:"optional" json:"criticality" yaml:"criticality"`
	// developer_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_service#developer_owners GoogleApphubService#developer_owners}
	DeveloperOwners interface{} `field:"optional" json:"developerOwners" yaml:"developerOwners"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_service#environment GoogleApphubService#environment}
	Environment *GoogleApphubServiceAttributesEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// operator_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_service#operator_owners GoogleApphubService#operator_owners}
	OperatorOwners interface{} `field:"optional" json:"operatorOwners" yaml:"operatorOwners"`
}

