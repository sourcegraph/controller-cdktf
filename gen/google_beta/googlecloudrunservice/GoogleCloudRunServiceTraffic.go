package googlecloudrunservice


type GoogleCloudRunServiceTraffic struct {
	// Percent specifies percent of the traffic to this Revision or Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#percent GoogleCloudRunService#percent}
	Percent *float64 `field:"required" json:"percent" yaml:"percent"`
	// LatestRevision may be optionally provided to indicate that the latest ready Revision of the Configuration should be used for this traffic target.
	//
	// When
	// provided LatestRevision must be true if RevisionName is empty; it must be
	// false when RevisionName is non-empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#latest_revision GoogleCloudRunService#latest_revision}
	LatestRevision interface{} `field:"optional" json:"latestRevision" yaml:"latestRevision"`
	// RevisionName of a specific revision to which to send this portion of traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#revision_name GoogleCloudRunService#revision_name}
	RevisionName *string `field:"optional" json:"revisionName" yaml:"revisionName"`
	// Tag is optionally used to expose a dedicated url for referencing this target exclusively.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#tag GoogleCloudRunService#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

