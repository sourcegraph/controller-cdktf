package googledataprocmetastoretableiammember


type GoogleDataprocMetastoreTableIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_metastore_table_iam_member#expression GoogleDataprocMetastoreTableIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_metastore_table_iam_member#title GoogleDataprocMetastoreTableIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_metastore_table_iam_member#description GoogleDataprocMetastoreTableIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

