package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecContainersEnvFrom struct {
	// config_map_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#config_map_ref GoogleCloudRunService#config_map_ref}
	ConfigMapRef *GoogleCloudRunServiceTemplateSpecContainersEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	// An optional identifier to prepend to each key in the ConfigMap.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#prefix GoogleCloudRunService#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#secret_ref GoogleCloudRunService#secret_ref}
	SecretRef *GoogleCloudRunServiceTemplateSpecContainersEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

