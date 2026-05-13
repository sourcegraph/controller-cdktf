package workstationsworkstationiammember


type WorkstationsWorkstationIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_member#expression WorkstationsWorkstationIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_member#title WorkstationsWorkstationIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_member#description WorkstationsWorkstationIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

