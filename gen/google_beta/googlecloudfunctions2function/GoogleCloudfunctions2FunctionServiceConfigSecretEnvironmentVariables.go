package googlecloudfunctions2function


type GoogleCloudfunctions2FunctionServiceConfigSecretEnvironmentVariables struct {
	// Name of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#key GoogleCloudfunctions2Function#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret.
	//
	// If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#project_id GoogleCloudfunctions2Function#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Name of the secret in secret manager (not the full resource name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#secret GoogleCloudfunctions2Function#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// Version of the secret (version number or the string 'latest').
	//
	// It is recommended to use a numeric version for secret environment variables as any updates to the secret value is not reflected until new instances start.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#version GoogleCloudfunctions2Function#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

