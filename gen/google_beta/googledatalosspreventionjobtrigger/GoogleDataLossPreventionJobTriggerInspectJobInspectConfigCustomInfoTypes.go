package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypes struct {
	// info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#info_type GoogleDataLossPreventionJobTrigger#info_type}
	InfoType *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType `field:"required" json:"infoType" yaml:"infoType"`
	// dictionary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#dictionary GoogleDataLossPreventionJobTrigger#dictionary}
	Dictionary *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary `field:"optional" json:"dictionary" yaml:"dictionary"`
	// If set to EXCLUSION_TYPE_EXCLUDE this infoType will not cause a finding to be returned.
	//
	// It still can be used for rules matching. Possible values: ["EXCLUSION_TYPE_EXCLUDE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#exclusion_type GoogleDataLossPreventionJobTrigger#exclusion_type}
	ExclusionType *string `field:"optional" json:"exclusionType" yaml:"exclusionType"`
	// Likelihood to return for this CustomInfoType.
	//
	// This base value can be altered by a detection rule if the finding meets the criteria
	// specified by the rule. Default value: "VERY_LIKELY" Possible values: ["VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#likelihood GoogleDataLossPreventionJobTrigger#likelihood}
	Likelihood *string `field:"optional" json:"likelihood" yaml:"likelihood"`
	// regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#regex GoogleDataLossPreventionJobTrigger#regex}
	Regex *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesRegex `field:"optional" json:"regex" yaml:"regex"`
	// sensitivity_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#sensitivity_score GoogleDataLossPreventionJobTrigger#sensitivity_score}
	SensitivityScore *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesSensitivityScore `field:"optional" json:"sensitivityScore" yaml:"sensitivityScore"`
	// stored_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#stored_type GoogleDataLossPreventionJobTrigger#stored_type}
	StoredType *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType `field:"optional" json:"storedType" yaml:"storedType"`
	// surrogate_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#surrogate_type GoogleDataLossPreventionJobTrigger#surrogate_type}
	SurrogateType *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType `field:"optional" json:"surrogateType" yaml:"surrogateType"`
}

