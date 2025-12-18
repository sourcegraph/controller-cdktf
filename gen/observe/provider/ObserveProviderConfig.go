package provider


type ObserveProviderConfig struct {
	// Your Observe Customer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#customer ObserveProvider#customer}
	Customer *string `field:"required" json:"customer" yaml:"customer"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#alias ObserveProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// An Observe API Token. Used for authenticating requests to API in the absence of `user_email` and `user_password`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#api_token ObserveProvider#api_token}
	ApiToken *string `field:"optional" json:"apiToken" yaml:"apiToken"`
	// Default rematerialization mode for datasets (internal use).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#default_rematerialization_mode ObserveProvider#default_rematerialization_mode}
	DefaultRematerializationMode *string `field:"optional" json:"defaultRematerializationMode" yaml:"defaultRematerializationMode"`
	// Observe API domain. Defaults to `observeinc.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#domain ObserveProvider#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Enable generating object ID-name bindings for cross-tenant export/import (internal use).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#export_object_bindings ObserveProvider#export_object_bindings}
	ExportObjectBindings interface{} `field:"optional" json:"exportObjectBindings" yaml:"exportObjectBindings"`
	// Toggle experimental features.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#flags ObserveProvider#flags}
	Flags *string `field:"optional" json:"flags" yaml:"flags"`
	// HTTP client timeout. Defaults to 2m.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#http_client_timeout ObserveProvider#http_client_timeout}
	HttpClientTimeout *string `field:"optional" json:"httpClientTimeout" yaml:"httpClientTimeout"`
	// Skip TLS certificate validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#insecure ObserveProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// ID of an Observe object that serves as the parent (managing) object for all resources created by the provider (internal use).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#managing_object_id ObserveProvider#managing_object_id}
	ManagingObjectId *string `field:"optional" json:"managingObjectId" yaml:"managingObjectId"`
	// Maximum number of retries on temporary network failures. Defaults to 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#retry_count ObserveProvider#retry_count}
	RetryCount *float64 `field:"optional" json:"retryCount" yaml:"retryCount"`
	// Time between retries. Defaults to 3s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#retry_wait ObserveProvider#retry_wait}
	RetryWait *string `field:"optional" json:"retryWait" yaml:"retryWait"`
	// Skip making dry run API requests for dataset changes during the plan stage (for validation).
	//
	// This can speed up plan time, but means that certain classes of errors will not be detected until applying the changes (such as invalid OPAL).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#skip_dataset_dry_runs ObserveProvider#skip_dataset_dry_runs}
	SkipDatasetDryRuns interface{} `field:"optional" json:"skipDatasetDryRuns" yaml:"skipDatasetDryRuns"`
	// Source identifier comment. If null, fallback to `user_email`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#source_comment ObserveProvider#source_comment}
	SourceComment *string `field:"optional" json:"sourceComment" yaml:"sourceComment"`
	// Source identifier format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#source_format ObserveProvider#source_format}
	SourceFormat *string `field:"optional" json:"sourceFormat" yaml:"sourceFormat"`
	// User email. If supplied, `user_password` is also required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#user_email ObserveProvider#user_email}
	UserEmail *string `field:"optional" json:"userEmail" yaml:"userEmail"`
	// Password for provided `user_email`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs#user_password ObserveProvider#user_password}
	UserPassword *string `field:"optional" json:"userPassword" yaml:"userPassword"`
}

