package googlespannerinstancepartition


type GoogleSpannerInstancePartitionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_spanner_instance_partition#create GoogleSpannerInstancePartition#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_spanner_instance_partition#delete GoogleSpannerInstancePartition#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_spanner_instance_partition#update GoogleSpannerInstancePartition#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

