package googledataprocclusteriammember


type GoogleDataprocClusterIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_cluster_iam_member#expression GoogleDataprocClusterIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_cluster_iam_member#title GoogleDataprocClusterIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_cluster_iam_member#description GoogleDataprocClusterIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

