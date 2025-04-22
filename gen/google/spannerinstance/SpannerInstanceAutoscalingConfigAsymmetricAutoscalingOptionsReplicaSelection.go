package spannerinstance


type SpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsReplicaSelection struct {
	// The location of the replica to apply asymmetric autoscaling options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/spanner_instance#location SpannerInstance#location}
	Location *string `field:"required" json:"location" yaml:"location"`
}

