package googledataprocmetastorefederationiammember


type GoogleDataprocMetastoreFederationIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_federation_iam_member#expression GoogleDataprocMetastoreFederationIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_federation_iam_member#title GoogleDataprocMetastoreFederationIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_federation_iam_member#description GoogleDataprocMetastoreFederationIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

