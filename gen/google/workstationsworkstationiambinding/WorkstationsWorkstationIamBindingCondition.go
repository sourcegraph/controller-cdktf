package workstationsworkstationiambinding


type WorkstationsWorkstationIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_binding#expression WorkstationsWorkstationIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_binding#title WorkstationsWorkstationIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_binding#description WorkstationsWorkstationIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

