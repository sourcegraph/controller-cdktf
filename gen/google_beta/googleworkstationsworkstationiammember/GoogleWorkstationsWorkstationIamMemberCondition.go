package googleworkstationsworkstationiammember


type GoogleWorkstationsWorkstationIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_iam_member#expression GoogleWorkstationsWorkstationIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_iam_member#title GoogleWorkstationsWorkstationIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_iam_member#description GoogleWorkstationsWorkstationIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

