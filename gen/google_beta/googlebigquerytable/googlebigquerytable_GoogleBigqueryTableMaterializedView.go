package googlebigquerytable


type GoogleBigqueryTableMaterializedView struct {
	// A query whose result is persisted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_table#query GoogleBigqueryTable#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// Specifies if BigQuery should automatically refresh materialized view when the base table is updated. The default is true.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_table#enable_refresh GoogleBigqueryTable#enable_refresh}
	EnableRefresh interface{} `field:"optional" json:"enableRefresh" yaml:"enableRefresh"`
	// Specifies maximum frequency at which this materialized view will be refreshed. The default is 1800000.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_table#refresh_interval_ms GoogleBigqueryTable#refresh_interval_ms}
	RefreshIntervalMs *float64 `field:"optional" json:"refreshIntervalMs" yaml:"refreshIntervalMs"`
}

