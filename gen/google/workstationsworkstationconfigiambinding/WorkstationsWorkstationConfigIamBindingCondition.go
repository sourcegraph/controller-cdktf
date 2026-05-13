package workstationsworkstationconfigiambinding


type WorkstationsWorkstationConfigIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config_iam_binding#expression WorkstationsWorkstationConfigIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config_iam_binding#title WorkstationsWorkstationConfigIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config_iam_binding#description WorkstationsWorkstationConfigIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

