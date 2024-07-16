package googlecloudfunctions2function


type GoogleCloudfunctions2FunctionBuildConfig struct {
	// automatic_update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#automatic_update_policy GoogleCloudfunctions2Function#automatic_update_policy}
	AutomaticUpdatePolicy *GoogleCloudfunctions2FunctionBuildConfigAutomaticUpdatePolicy `field:"optional" json:"automaticUpdatePolicy" yaml:"automaticUpdatePolicy"`
	// User managed repository created in Artifact Registry optionally with a customer managed encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#docker_repository GoogleCloudfunctions2Function#docker_repository}
	DockerRepository *string `field:"optional" json:"dockerRepository" yaml:"dockerRepository"`
	// The name of the function (as defined in source code) that will be executed.
	//
	// Defaults to the resource name suffix, if not specified. For backward
	// compatibility, if function with given name is not found, then the system
	// will try to use function named "function". For Node.js this is name of a
	// function exported by the module specified in source_location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#entry_point GoogleCloudfunctions2Function#entry_point}
	EntryPoint *string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// User-provided build-time environment variables for the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#environment_variables GoogleCloudfunctions2Function#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// on_deploy_update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#on_deploy_update_policy GoogleCloudfunctions2Function#on_deploy_update_policy}
	OnDeployUpdatePolicy *GoogleCloudfunctions2FunctionBuildConfigOnDeployUpdatePolicy `field:"optional" json:"onDeployUpdatePolicy" yaml:"onDeployUpdatePolicy"`
	// The runtime in which to run the function. Required when deploying a new function, optional when updating an existing function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#runtime GoogleCloudfunctions2Function#runtime}
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// The fully-qualified name of the service account to be used for building the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#service_account GoogleCloudfunctions2Function#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#source GoogleCloudfunctions2Function#source}
	Source *GoogleCloudfunctions2FunctionBuildConfigSource `field:"optional" json:"source" yaml:"source"`
	// Name of the Cloud Build Custom Worker Pool that should be used to build the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#worker_pool GoogleCloudfunctions2Function#worker_pool}
	WorkerPool *string `field:"optional" json:"workerPool" yaml:"workerPool"`
}

