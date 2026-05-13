package bigquerydatapolicyv2datapolicyiammember


type BigqueryDatapolicyv2DataPolicyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_datapolicyv2_data_policy_iam_member#expression BigqueryDatapolicyv2DataPolicyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_datapolicyv2_data_policy_iam_member#title BigqueryDatapolicyv2DataPolicyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_datapolicyv2_data_policy_iam_member#description BigqueryDatapolicyv2DataPolicyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

