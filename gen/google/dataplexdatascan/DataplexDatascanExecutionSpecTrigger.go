package dataplexdatascan


type DataplexDatascanExecutionSpecTrigger struct {
	// on_demand block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_datascan#on_demand DataplexDatascan#on_demand}
	OnDemand *DataplexDatascanExecutionSpecTriggerOnDemand `field:"optional" json:"onDemand" yaml:"onDemand"`
	// one_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_datascan#one_time DataplexDatascan#one_time}
	OneTime *DataplexDatascanExecutionSpecTriggerOneTime `field:"optional" json:"oneTime" yaml:"oneTime"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_datascan#schedule DataplexDatascan#schedule}
	Schedule *DataplexDatascanExecutionSpecTriggerSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

