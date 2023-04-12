package googlebigqueryconnection


type GoogleBigqueryConnectionAzure struct {
	// The id of customer's directory that host the data.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_connection#customer_tenant_id GoogleBigqueryConnection#customer_tenant_id}
	CustomerTenantId *string `field:"required" json:"customerTenantId" yaml:"customerTenantId"`
}

