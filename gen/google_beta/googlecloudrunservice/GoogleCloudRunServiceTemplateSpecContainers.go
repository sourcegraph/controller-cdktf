package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecContainers struct {
	// Docker image name. This is most often a reference to a container located in the container registry, such as gcr.io/cloudrun/hello.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#image GoogleCloudRunService#image}
	Image *string `field:"required" json:"image" yaml:"image"`
	// Arguments to the entrypoint. The docker image's CMD is used if this is not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#args GoogleCloudRunService#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#command GoogleCloudRunService#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#env GoogleCloudRunService#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// env_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#env_from GoogleCloudRunService#env_from}
	EnvFrom interface{} `field:"optional" json:"envFrom" yaml:"envFrom"`
	// liveness_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#liveness_probe GoogleCloudRunService#liveness_probe}
	LivenessProbe *GoogleCloudRunServiceTemplateSpecContainersLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	// Name of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#name GoogleCloudRunService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#ports GoogleCloudRunService#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#resources GoogleCloudRunService#resources}
	Resources *GoogleCloudRunServiceTemplateSpecContainersResources `field:"optional" json:"resources" yaml:"resources"`
	// startup_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#startup_probe GoogleCloudRunService#startup_probe}
	StartupProbe *GoogleCloudRunServiceTemplateSpecContainersStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	// volume_mounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#volume_mounts GoogleCloudRunService#volume_mounts}
	VolumeMounts interface{} `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	// Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#working_dir GoogleCloudRunService#working_dir}
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

