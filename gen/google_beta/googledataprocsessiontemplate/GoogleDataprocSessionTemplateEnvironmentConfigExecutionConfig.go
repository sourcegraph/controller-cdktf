package googledataprocsessiontemplate


type GoogleDataprocSessionTemplateEnvironmentConfigExecutionConfig struct {
	// authentication_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#authentication_config GoogleDataprocSessionTemplate#authentication_config}
	AuthenticationConfig *GoogleDataprocSessionTemplateEnvironmentConfigExecutionConfigAuthenticationConfig `field:"optional" json:"authenticationConfig" yaml:"authenticationConfig"`
	// The duration to keep the session alive while it's idling.
	//
	// Exceeding this threshold causes the session to terminate. Minimum value is 10 minutes; maximum value is 14 day.
	// Defaults to 1 hour if not set. If both ttl and idleTtl are specified for an interactive session, the conditions
	// are treated as OR conditions: the workload will be terminated when it has been idle for idleTtl or when ttl has
	// been exceeded, whichever occurs first.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#idle_ttl GoogleDataprocSessionTemplate#idle_ttl}
	IdleTtl *string `field:"optional" json:"idleTtl" yaml:"idleTtl"`
	// The Cloud KMS key to use for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#kms_key GoogleDataprocSessionTemplate#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Tags used for network traffic control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#network_tags GoogleDataprocSessionTemplate#network_tags}
	NetworkTags *[]*string `field:"optional" json:"networkTags" yaml:"networkTags"`
	// Service account that used to execute workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#service_account GoogleDataprocSessionTemplate#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// A Cloud Storage bucket used to stage workload dependencies, config files, and store workload output and other ephemeral data, such as Spark history files.
	//
	// If you do not specify a staging bucket,
	// Cloud Dataproc will determine a Cloud Storage location according to the region where your workload is running,
	// and then create and manage project-level, per-location staging and temporary buckets.
	// This field requires a Cloud Storage bucket name, not a gs://... URI to a Cloud Storage bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#staging_bucket GoogleDataprocSessionTemplate#staging_bucket}
	StagingBucket *string `field:"optional" json:"stagingBucket" yaml:"stagingBucket"`
	// Subnetwork configuration for workload execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#subnetwork_uri GoogleDataprocSessionTemplate#subnetwork_uri}
	SubnetworkUri *string `field:"optional" json:"subnetworkUri" yaml:"subnetworkUri"`
	// The duration after which the workload will be terminated.
	//
	// When the workload exceeds this duration, it will be unconditionally terminated without waiting for ongoing
	// work to finish. If ttl is not specified for a session workload, the workload will be allowed to run until it
	// exits naturally (or run forever without exiting). If ttl is not specified for an interactive session,
	// it defaults to 24 hours. If ttl is not specified for a batch that uses 2.1+ runtime version, it defaults to 4 hours.
	// Minimum value is 10 minutes; maximum value is 14 days. If both ttl and idleTtl are specified (for an interactive session),
	// the conditions are treated as OR conditions: the workload will be terminated when it has been idle for idleTtl or
	// when ttl has been exceeded, whichever occurs first.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#ttl GoogleDataprocSessionTemplate#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

