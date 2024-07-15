package googleworkstationsworkstationiambinding


type GoogleWorkstationsWorkstationIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_iam_binding#expression GoogleWorkstationsWorkstationIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_iam_binding#title GoogleWorkstationsWorkstationIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_iam_binding#description GoogleWorkstationsWorkstationIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

