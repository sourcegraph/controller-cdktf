package googleappenginestandardappversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleAppEngineStandardAppVersionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#deployment GoogleAppEngineStandardAppVersion#deployment}
	Deployment *GoogleAppEngineStandardAppVersionDeployment `field:"required" json:"deployment" yaml:"deployment"`
	// entrypoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#entrypoint GoogleAppEngineStandardAppVersion#entrypoint}
	Entrypoint *GoogleAppEngineStandardAppVersionEntrypoint `field:"required" json:"entrypoint" yaml:"entrypoint"`
	// Desired runtime. Example python27.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#runtime GoogleAppEngineStandardAppVersion#runtime}
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// AppEngine service resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#service GoogleAppEngineStandardAppVersion#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Allows App Engine second generation runtimes to access the legacy bundled services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#app_engine_apis GoogleAppEngineStandardAppVersion#app_engine_apis}
	AppEngineApis interface{} `field:"optional" json:"appEngineApis" yaml:"appEngineApis"`
	// automatic_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#automatic_scaling GoogleAppEngineStandardAppVersion#automatic_scaling}
	AutomaticScaling *GoogleAppEngineStandardAppVersionAutomaticScaling `field:"optional" json:"automaticScaling" yaml:"automaticScaling"`
	// basic_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#basic_scaling GoogleAppEngineStandardAppVersion#basic_scaling}
	BasicScaling *GoogleAppEngineStandardAppVersionBasicScaling `field:"optional" json:"basicScaling" yaml:"basicScaling"`
	// If set to 'true', the service will be deleted if it is the last version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#delete_service_on_destroy GoogleAppEngineStandardAppVersion#delete_service_on_destroy}
	DeleteServiceOnDestroy interface{} `field:"optional" json:"deleteServiceOnDestroy" yaml:"deleteServiceOnDestroy"`
	// Environment variables available to the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#env_variables GoogleAppEngineStandardAppVersion#env_variables}
	EnvVariables *map[string]*string `field:"optional" json:"envVariables" yaml:"envVariables"`
	// handlers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#handlers GoogleAppEngineStandardAppVersion#handlers}
	Handlers interface{} `field:"optional" json:"handlers" yaml:"handlers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#id GoogleAppEngineStandardAppVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A list of the types of messages that this application is able to receive.
	//
	// Possible values: ["INBOUND_SERVICE_MAIL", "INBOUND_SERVICE_MAIL_BOUNCE", "INBOUND_SERVICE_XMPP_ERROR", "INBOUND_SERVICE_XMPP_MESSAGE", "INBOUND_SERVICE_XMPP_SUBSCRIBE", "INBOUND_SERVICE_XMPP_PRESENCE", "INBOUND_SERVICE_CHANNEL_PRESENCE", "INBOUND_SERVICE_WARMUP"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#inbound_services GoogleAppEngineStandardAppVersion#inbound_services}
	InboundServices *[]*string `field:"optional" json:"inboundServices" yaml:"inboundServices"`
	// Instance class that is used to run this version.
	//
	// Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
	// Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#instance_class GoogleAppEngineStandardAppVersion#instance_class}
	InstanceClass *string `field:"optional" json:"instanceClass" yaml:"instanceClass"`
	// libraries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#libraries GoogleAppEngineStandardAppVersion#libraries}
	Libraries interface{} `field:"optional" json:"libraries" yaml:"libraries"`
	// manual_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#manual_scaling GoogleAppEngineStandardAppVersion#manual_scaling}
	ManualScaling *GoogleAppEngineStandardAppVersionManualScaling `field:"optional" json:"manualScaling" yaml:"manualScaling"`
	// If set to 'true', the application version will not be deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#noop_on_destroy GoogleAppEngineStandardAppVersion#noop_on_destroy}
	NoopOnDestroy interface{} `field:"optional" json:"noopOnDestroy" yaml:"noopOnDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#project GoogleAppEngineStandardAppVersion#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The version of the API in the given runtime environment.
	//
	// Please see the app.yaml reference for valid values at 'https://cloud.google.com/appengine/docs/standard/<language>/config/appref'\
	// Substitute '<language>' with 'python', 'java', 'php', 'ruby', 'go' or 'nodejs'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#runtime_api_version GoogleAppEngineStandardAppVersion#runtime_api_version}
	RuntimeApiVersion *string `field:"optional" json:"runtimeApiVersion" yaml:"runtimeApiVersion"`
	// The identity that the deployed version will run as.
	//
	// Admin API will use the App Engine Appspot service account as default if this field is neither provided in app.yaml file nor through CLI flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#service_account GoogleAppEngineStandardAppVersion#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Whether multiple requests can be dispatched to this version at once.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#threadsafe GoogleAppEngineStandardAppVersion#threadsafe}
	Threadsafe interface{} `field:"optional" json:"threadsafe" yaml:"threadsafe"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#timeouts GoogleAppEngineStandardAppVersion#timeouts}
	Timeouts *GoogleAppEngineStandardAppVersionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Relative name of the version within the service.
	//
	// For example, 'v1'. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#version_id GoogleAppEngineStandardAppVersion#version_id}
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
	// vpc_access_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#vpc_access_connector GoogleAppEngineStandardAppVersion#vpc_access_connector}
	VpcAccessConnector *GoogleAppEngineStandardAppVersionVpcAccessConnector `field:"optional" json:"vpcAccessConnector" yaml:"vpcAccessConnector"`
}

