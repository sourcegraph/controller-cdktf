package googlebigquerydatapolicyv2datapolicy


type GoogleBigqueryDatapolicyv2DataPolicyDataMaskingPolicy struct {
	// A predefined masking expression. Possible values: SHA256 ALWAYS_NULL DEFAULT_MASKING_VALUE LAST_FOUR_CHARACTERS FIRST_FOUR_CHARACTERS EMAIL_MASK DATE_YEAR_MASK RANDOM_HASH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy#predefined_expression GoogleBigqueryDatapolicyv2DataPolicy#predefined_expression}
	PredefinedExpression *string `field:"optional" json:"predefinedExpression" yaml:"predefinedExpression"`
	// The name of the BigQuery routine that contains the custom masking routine, in the format of 'projects/{project_number}/datasets/{dataset_id}/routines/{routine_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy#routine GoogleBigqueryDatapolicyv2DataPolicy#routine}
	Routine *string `field:"optional" json:"routine" yaml:"routine"`
}

