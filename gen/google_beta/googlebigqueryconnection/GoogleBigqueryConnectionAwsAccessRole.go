package googlebigqueryconnection


type GoogleBigqueryConnectionAwsAccessRole struct {
	// The user’s AWS IAM Role that trusts the Google-owned AWS IAM user Connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#iam_role_id GoogleBigqueryConnection#iam_role_id}
	IamRoleId *string `field:"required" json:"iamRoleId" yaml:"iamRoleId"`
}

