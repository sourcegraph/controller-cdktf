package googledataprocsessiontemplate


type GoogleDataprocSessionTemplateEnvironmentConfigExecutionConfigAuthenticationConfig struct {
	// Authentication type for the user workload running in containers. Possible values: ["SERVICE_ACCOUNT", "END_USER_CREDENTIALS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#user_workload_authentication_type GoogleDataprocSessionTemplate#user_workload_authentication_type}
	UserWorkloadAuthenticationType *string `field:"optional" json:"userWorkloadAuthenticationType" yaml:"userWorkloadAuthenticationType"`
}

