package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartChartDatasource struct {
	// Name(s) of the datasource used in the chart.
	//
	// Available values include:
	// 'UDM', 'ENTITY', 'INGESTION_METRICS', 'RULE_DETECTIONS', 'RULESETS',
	// 'GLOBAL', 'IOC_MATCHES', 'RULES', 'SOAR_CASES', 'SOAR_PLAYBOOKS',
	// 'SOAR_CASE_HISTORY', 'DATA_TABLE', 'INVESTIGATION', 'INVESTIGATION_FEEDBACK'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#data_sources GoogleChronicleDashboardChart#data_sources}
	DataSources *[]*string `field:"optional" json:"dataSources" yaml:"dataSources"`
}

