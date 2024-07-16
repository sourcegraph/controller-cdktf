package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigOrgConfig struct {
	// location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#location GoogleDataLossPreventionDiscoveryConfig#location}
	Location *GoogleDataLossPreventionDiscoveryConfigOrgConfigLocation `field:"optional" json:"location" yaml:"location"`
	// The project that will run the scan.
	//
	// The DLP service account that exists within this project must have access to all resources that are profiled, and the cloud DLP API must be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#project_id GoogleDataLossPreventionDiscoveryConfig#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

