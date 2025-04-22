package dataplexdatascan


type DataplexDatascanDataQualitySpecRulesSqlAssertion struct {
	// The SQL statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/dataplex_datascan#sql_statement DataplexDatascan#sql_statement}
	SqlStatement *string `field:"required" json:"sqlStatement" yaml:"sqlStatement"`
}

