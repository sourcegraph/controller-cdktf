package googleapphubworkload


type GoogleApphubWorkloadAttributes struct {
	// business_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_workload#business_owners GoogleApphubWorkload#business_owners}
	BusinessOwners interface{} `field:"optional" json:"businessOwners" yaml:"businessOwners"`
	// criticality block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_workload#criticality GoogleApphubWorkload#criticality}
	Criticality *GoogleApphubWorkloadAttributesCriticality `field:"optional" json:"criticality" yaml:"criticality"`
	// developer_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_workload#developer_owners GoogleApphubWorkload#developer_owners}
	DeveloperOwners interface{} `field:"optional" json:"developerOwners" yaml:"developerOwners"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_workload#environment GoogleApphubWorkload#environment}
	Environment *GoogleApphubWorkloadAttributesEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// operator_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_workload#operator_owners GoogleApphubWorkload#operator_owners}
	OperatorOwners interface{} `field:"optional" json:"operatorOwners" yaml:"operatorOwners"`
}

