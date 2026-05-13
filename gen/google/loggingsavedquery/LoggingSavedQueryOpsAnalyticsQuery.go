package loggingsavedquery


type LoggingSavedQueryOpsAnalyticsQuery struct {
	// A logs analytics SQL query, which generally follows BigQuery format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/logging_saved_query#sql_query_text LoggingSavedQuery#sql_query_text}
	SqlQueryText *string `field:"required" json:"sqlQueryText" yaml:"sqlQueryText"`
}

