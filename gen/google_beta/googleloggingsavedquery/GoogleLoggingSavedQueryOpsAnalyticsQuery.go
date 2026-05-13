package googleloggingsavedquery


type GoogleLoggingSavedQueryOpsAnalyticsQuery struct {
	// A logs analytics SQL query, which generally follows BigQuery format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#sql_query_text GoogleLoggingSavedQuery#sql_query_text}
	SqlQueryText *string `field:"required" json:"sqlQueryText" yaml:"sqlQueryText"`
}

