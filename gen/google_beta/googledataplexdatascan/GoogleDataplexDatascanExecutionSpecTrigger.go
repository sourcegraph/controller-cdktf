package googledataplexdatascan


type GoogleDataplexDatascanExecutionSpecTrigger struct {
	// on_demand block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#on_demand GoogleDataplexDatascan#on_demand}
	OnDemand *GoogleDataplexDatascanExecutionSpecTriggerOnDemand `field:"optional" json:"onDemand" yaml:"onDemand"`
	// one_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#one_time GoogleDataplexDatascan#one_time}
	OneTime *GoogleDataplexDatascanExecutionSpecTriggerOneTime `field:"optional" json:"oneTime" yaml:"oneTime"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#schedule GoogleDataplexDatascan#schedule}
	Schedule *GoogleDataplexDatascanExecutionSpecTriggerSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

