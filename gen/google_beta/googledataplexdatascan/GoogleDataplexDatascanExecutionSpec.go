package googledataplexdatascan


type GoogleDataplexDatascanExecutionSpec struct {
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#trigger GoogleDataplexDatascan#trigger}
	Trigger *GoogleDataplexDatascanExecutionSpecTrigger `field:"required" json:"trigger" yaml:"trigger"`
	// The unnested field (of type Date or Timestamp) that contains values which monotonically increase over time.
	//
	// If not specified, a data scan will run for all data in the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#field GoogleDataplexDatascan#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
}

