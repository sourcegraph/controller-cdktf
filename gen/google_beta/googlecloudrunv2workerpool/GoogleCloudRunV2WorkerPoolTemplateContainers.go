package googlecloudrunv2workerpool


type GoogleCloudRunV2WorkerPoolTemplateContainers struct {
	// URL of the Container image in Google Container Registry or Google Artifact Registry. More info: https://kubernetes.io/docs/concepts/containers/images.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#image GoogleCloudRunV2WorkerPool#image}
	Image *string `field:"required" json:"image" yaml:"image"`
	// Arguments to the entrypoint.
	//
	// The docker image's CMD is used if this is not provided. Variable references are not supported in Cloud Run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#args GoogleCloudRunV2WorkerPool#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Entrypoint array.
	//
	// Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#command GoogleCloudRunV2WorkerPool#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Names of the containers that must start before this container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#depends_on GoogleCloudRunV2WorkerPool#depends_on}
	DependsOn *[]*string `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#env GoogleCloudRunV2WorkerPool#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// liveness_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#liveness_probe GoogleCloudRunV2WorkerPool#liveness_probe}
	LivenessProbe *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	// Name of the container specified as a DNS_LABEL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#name GoogleCloudRunV2WorkerPool#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#resources GoogleCloudRunV2WorkerPool#resources}
	Resources *GoogleCloudRunV2WorkerPoolTemplateContainersResources `field:"optional" json:"resources" yaml:"resources"`
	// startup_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#startup_probe GoogleCloudRunV2WorkerPool#startup_probe}
	StartupProbe *GoogleCloudRunV2WorkerPoolTemplateContainersStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	// volume_mounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#volume_mounts GoogleCloudRunV2WorkerPool#volume_mounts}
	VolumeMounts interface{} `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	// Container's working directory.
	//
	// If not specified, the container runtime's default will be used, which might be configured in the container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#working_dir GoogleCloudRunV2WorkerPool#working_dir}
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

