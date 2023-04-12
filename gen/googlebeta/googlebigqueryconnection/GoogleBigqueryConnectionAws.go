package googlebigqueryconnection


type GoogleBigqueryConnectionAws struct {
	// access_role block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_connection#access_role GoogleBigqueryConnection#access_role}
	AccessRole *GoogleBigqueryConnectionAwsAccessRole `field:"required" json:"accessRole" yaml:"accessRole"`
}

