package googledataplexdatascan


type GoogleDataplexDatascanExecutionSpecTrigger struct {
	// on_demand block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#on_demand GoogleDataplexDatascan#on_demand}
	OnDemand *GoogleDataplexDatascanExecutionSpecTriggerOnDemand `field:"optional" json:"onDemand" yaml:"onDemand"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#schedule GoogleDataplexDatascan#schedule}
	Schedule *GoogleDataplexDatascanExecutionSpecTriggerSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

