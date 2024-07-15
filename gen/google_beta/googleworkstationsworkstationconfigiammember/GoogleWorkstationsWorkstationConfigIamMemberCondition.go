package googleworkstationsworkstationconfigiammember


type GoogleWorkstationsWorkstationConfigIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config_iam_member#expression GoogleWorkstationsWorkstationConfigIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config_iam_member#title GoogleWorkstationsWorkstationConfigIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config_iam_member#description GoogleWorkstationsWorkstationConfigIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

