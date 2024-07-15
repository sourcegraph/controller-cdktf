package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetConditionsOrConditions struct {
	// Duration format. The minimum age a table must have before Cloud DLP can profile it. Value greater than 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#min_age GoogleDataLossPreventionDiscoveryConfig#min_age}
	MinAge *string `field:"optional" json:"minAge" yaml:"minAge"`
	// Minimum number of rows that should be present before Cloud DLP profiles as a table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#min_row_count GoogleDataLossPreventionDiscoveryConfig#min_row_count}
	MinRowCount *float64 `field:"optional" json:"minRowCount" yaml:"minRowCount"`
}

