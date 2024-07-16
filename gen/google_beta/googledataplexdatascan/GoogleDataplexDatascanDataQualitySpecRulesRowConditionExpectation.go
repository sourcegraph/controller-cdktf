package googledataplexdatascan


type GoogleDataplexDatascanDataQualitySpecRulesRowConditionExpectation struct {
	// The SQL expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#sql_expression GoogleDataplexDatascan#sql_expression}
	SqlExpression *string `field:"required" json:"sqlExpression" yaml:"sqlExpression"`
}

