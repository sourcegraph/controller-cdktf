package googledialogflowcxsecuritysettings


type GoogleDialogflowCxSecuritySettingsInsightsExportSettings struct {
	// If enabled, we will automatically exports conversations to Insights and Insights runs its analyzers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_security_settings#enable_insights_export GoogleDialogflowCxSecuritySettings#enable_insights_export}
	EnableInsightsExport interface{} `field:"required" json:"enableInsightsExport" yaml:"enableInsightsExport"`
}

