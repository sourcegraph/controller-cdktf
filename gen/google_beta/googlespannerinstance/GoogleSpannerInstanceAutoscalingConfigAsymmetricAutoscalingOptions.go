package googlespannerinstance


type GoogleSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptions struct {
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_instance#overrides GoogleSpannerInstance#overrides}
	Overrides *GoogleSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverrides `field:"required" json:"overrides" yaml:"overrides"`
	// replica_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_instance#replica_selection GoogleSpannerInstance#replica_selection}
	ReplicaSelection *GoogleSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsReplicaSelection `field:"required" json:"replicaSelection" yaml:"replicaSelection"`
}

