package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobInspectConfigLimits struct {
	// max_findings_per_info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#max_findings_per_info_type GoogleDataLossPreventionJobTrigger#max_findings_per_info_type}
	MaxFindingsPerInfoType interface{} `field:"optional" json:"maxFindingsPerInfoType" yaml:"maxFindingsPerInfoType"`
	// Max number of findings that will be returned for each item scanned. The maximum returned is 2000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#max_findings_per_item GoogleDataLossPreventionJobTrigger#max_findings_per_item}
	MaxFindingsPerItem *float64 `field:"optional" json:"maxFindingsPerItem" yaml:"maxFindingsPerItem"`
	// Max number of findings that will be returned per request/job. The maximum returned is 2000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#max_findings_per_request GoogleDataLossPreventionJobTrigger#max_findings_per_request}
	MaxFindingsPerRequest *float64 `field:"optional" json:"maxFindingsPerRequest" yaml:"maxFindingsPerRequest"`
}

