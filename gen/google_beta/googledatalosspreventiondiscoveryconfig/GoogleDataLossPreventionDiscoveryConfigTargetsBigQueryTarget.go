package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTarget struct {
	// cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#cadence GoogleDataLossPreventionDiscoveryConfig#cadence}
	Cadence *GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetCadence `field:"optional" json:"cadence" yaml:"cadence"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#conditions GoogleDataLossPreventionDiscoveryConfig#conditions}
	Conditions *GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// disabled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#disabled GoogleDataLossPreventionDiscoveryConfig#disabled}
	Disabled *GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetDisabled `field:"optional" json:"disabled" yaml:"disabled"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#filter GoogleDataLossPreventionDiscoveryConfig#filter}
	Filter *GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilter `field:"optional" json:"filter" yaml:"filter"`
}

