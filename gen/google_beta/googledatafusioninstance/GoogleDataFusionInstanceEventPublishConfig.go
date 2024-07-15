package googledatafusioninstance


type GoogleDataFusionInstanceEventPublishConfig struct {
	// Option to enable Event Publishing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#enabled GoogleDataFusionInstance#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The resource name of the Pub/Sub topic. Format: projects/{projectId}/topics/{topic_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#topic GoogleDataFusionInstance#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

