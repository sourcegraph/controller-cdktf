package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceScalingConfig struct {
	// autoscaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#autoscaling_config GoogleDataprocMetastoreService#autoscaling_config}
	AutoscalingConfig *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfig `field:"optional" json:"autoscalingConfig" yaml:"autoscalingConfig"`
	// Metastore instance sizes. Possible values: ["EXTRA_SMALL", "SMALL", "MEDIUM", "LARGE", "EXTRA_LARGE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#instance_size GoogleDataprocMetastoreService#instance_size}
	InstanceSize *string `field:"optional" json:"instanceSize" yaml:"instanceSize"`
	// Scaling factor, in increments of 0.1 for values less than 1.0, and increments of 1.0 for values greater than 1.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#scaling_factor GoogleDataprocMetastoreService#scaling_factor}
	ScalingFactor *float64 `field:"optional" json:"scalingFactor" yaml:"scalingFactor"`
}

