package googleclouddeploytarget


type GoogleClouddeployTargetExecutionConfigsPrivatePool struct {
	// Required. Resource name of the Cloud Build worker pool to use. The format is `projects/{project}/locations/{location}/workerPools/{pool}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_target#worker_pool GoogleClouddeployTarget#worker_pool}
	WorkerPool *string `field:"required" json:"workerPool" yaml:"workerPool"`
	// Optional.
	//
	// Cloud Storage location where execution outputs should be stored. This can either be a bucket ("gs://my-bucket") or a path within a bucket ("gs://my-bucket/my-dir"). If unspecified, a default bucket located in the same region will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_target#artifact_storage GoogleClouddeployTarget#artifact_storage}
	ArtifactStorage *string `field:"optional" json:"artifactStorage" yaml:"artifactStorage"`
	// Optional. Google service account to use for execution. If unspecified, the project execution service account (-compute@developer.gserviceaccount.com) will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_target#service_account GoogleClouddeployTarget#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

