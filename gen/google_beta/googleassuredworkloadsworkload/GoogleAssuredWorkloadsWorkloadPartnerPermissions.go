package googleassuredworkloadsworkload


type GoogleAssuredWorkloadsWorkloadPartnerPermissions struct {
	// Optional. Allow partner to view violation alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_assured_workloads_workload#assured_workloads_monitoring GoogleAssuredWorkloadsWorkload#assured_workloads_monitoring}
	AssuredWorkloadsMonitoring interface{} `field:"optional" json:"assuredWorkloadsMonitoring" yaml:"assuredWorkloadsMonitoring"`
	// Allow the partner to view inspectability logs and monitoring violations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_assured_workloads_workload#data_logs_viewer GoogleAssuredWorkloadsWorkload#data_logs_viewer}
	DataLogsViewer interface{} `field:"optional" json:"dataLogsViewer" yaml:"dataLogsViewer"`
	// Optional. Allow partner to view access approval logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_assured_workloads_workload#service_access_approver GoogleAssuredWorkloadsWorkload#service_access_approver}
	ServiceAccessApprover interface{} `field:"optional" json:"serviceAccessApprover" yaml:"serviceAccessApprover"`
}

