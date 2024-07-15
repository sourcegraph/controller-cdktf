package googledataprocmetastoreserviceiammember


type GoogleDataprocMetastoreServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service_iam_member#expression GoogleDataprocMetastoreServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service_iam_member#title GoogleDataprocMetastoreServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service_iam_member#description GoogleDataprocMetastoreServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

