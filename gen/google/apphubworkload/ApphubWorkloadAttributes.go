package apphubworkload

type ApphubWorkloadAttributes struct {
	// business_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/apphub_workload#business_owners ApphubWorkload#business_owners}
	BusinessOwners any `field:"optional" json:"businessOwners" yaml:"businessOwners"`
	// criticality block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/apphub_workload#criticality ApphubWorkload#criticality}
	Criticality *ApphubWorkloadAttributesCriticality `field:"optional" json:"criticality" yaml:"criticality"`
	// developer_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/apphub_workload#developer_owners ApphubWorkload#developer_owners}
	DeveloperOwners any `field:"optional" json:"developerOwners" yaml:"developerOwners"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/apphub_workload#environment ApphubWorkload#environment}
	Environment *ApphubWorkloadAttributesEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// operator_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/apphub_workload#operator_owners ApphubWorkload#operator_owners}
	OperatorOwners any `field:"optional" json:"operatorOwners" yaml:"operatorOwners"`
}
