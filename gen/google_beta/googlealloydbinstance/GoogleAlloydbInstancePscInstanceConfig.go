package googlealloydbinstance


type GoogleAlloydbInstancePscInstanceConfig struct {
	// List of consumer projects that are allowed to create PSC endpoints to service-attachments to this instance.
	//
	// These should be specified as project numbers only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_instance#allowed_consumer_projects GoogleAlloydbInstance#allowed_consumer_projects}
	AllowedConsumerProjects *[]*string `field:"optional" json:"allowedConsumerProjects" yaml:"allowedConsumerProjects"`
}

