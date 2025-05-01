package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpec struct {
	// ContainerConcurrency specifies the maximum allowed in-flight (concurrent) requests per container of the Revision.
	//
	// If not specified or 0, defaults to 80 when
	// requested CPU >= 1 and defaults to 1 when requested CPU < 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_run_service#container_concurrency GoogleCloudRunService#container_concurrency}
	ContainerConcurrency *float64 `field:"optional" json:"containerConcurrency" yaml:"containerConcurrency"`
	// containers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_run_service#containers GoogleCloudRunService#containers}
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// Node Selector describes the hardware requirements of the resources.
	//
	// Use the following node selector keys to configure features on a Revision:
	//   - 'run.googleapis.com/accelerator' sets the [type of GPU](https://cloud.google.com/run/docs/configuring/services/gpu) required by the Revision to run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_run_service#node_selector GoogleCloudRunService#node_selector}
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	// Email address of the IAM service account associated with the revision of the service.
	//
	// The service account represents the identity of the running revision,
	// and determines what permissions the revision has. If not provided, the revision
	// will use the project's default service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_run_service#service_account_name GoogleCloudRunService#service_account_name}
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// TimeoutSeconds holds the max duration the instance is allowed for responding to a request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_run_service#timeout_seconds GoogleCloudRunService#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_run_service#volumes GoogleCloudRunService#volumes}
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

