package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecContainersEnvFromConfigMapRef struct {
	// local_object_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#local_object_reference GoogleCloudRunService#local_object_reference}
	LocalObjectReference *GoogleCloudRunServiceTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference `field:"optional" json:"localObjectReference" yaml:"localObjectReference"`
	// Specify whether the ConfigMap must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#optional GoogleCloudRunService#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

