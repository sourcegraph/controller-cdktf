package googlechroniclefeed


type GoogleChronicleFeedDetailsProofpointOnDemandSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Cluster ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#cluster_id GoogleChronicleFeed#cluster_id}
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
}

