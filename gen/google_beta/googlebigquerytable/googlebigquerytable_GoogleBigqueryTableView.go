package googlebigquerytable


type GoogleBigqueryTableView struct {
	// A query that BigQuery executes when the view is referenced.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_table#query GoogleBigqueryTable#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// Specifies whether to use BigQuery's legacy SQL for this view.
	//
	// The default value is true. If set to false, the view will use BigQuery's standard SQL
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_table#use_legacy_sql GoogleBigqueryTable#use_legacy_sql}
	UseLegacySql interface{} `field:"optional" json:"useLegacySql" yaml:"useLegacySql"`
}

