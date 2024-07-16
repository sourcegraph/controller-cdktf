package googledatafusioninstanceiammember


type GoogleDataFusionInstanceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance_iam_member#expression GoogleDataFusionInstanceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance_iam_member#title GoogleDataFusionInstanceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance_iam_member#description GoogleDataFusionInstanceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

