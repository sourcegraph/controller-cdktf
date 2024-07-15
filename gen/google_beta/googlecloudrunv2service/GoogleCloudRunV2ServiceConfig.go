package googlecloudrunv2service

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCloudRunV2ServiceConfig struct {
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
	// The location of the cloud run service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#location GoogleCloudRunV2Service#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of the Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#name GoogleCloudRunV2Service#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#template GoogleCloudRunV2Service#template}
	Template *GoogleCloudRunV2ServiceTemplate `field:"required" json:"template" yaml:"template"`
	// Unstructured key value map that may be set by external tools to store and arbitrary metadata.
	//
	// They are not queryable and should be preserved when modifying objects.
	//
	// Cloud Run API v2 does not support annotations with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected in new resources.
	// All system annotations in v1 now have a corresponding field in v2 Service.
	//
	// This field follows Kubernetes annotations' namespacing, limits, and rules.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#annotations GoogleCloudRunV2Service#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// binary_authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#binary_authorization GoogleCloudRunV2Service#binary_authorization}
	BinaryAuthorization *GoogleCloudRunV2ServiceBinaryAuthorization `field:"optional" json:"binaryAuthorization" yaml:"binaryAuthorization"`
	// Arbitrary identifier for the API client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#client GoogleCloudRunV2Service#client}
	Client *string `field:"optional" json:"client" yaml:"client"`
	// Arbitrary version identifier for the API client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#client_version GoogleCloudRunV2Service#client_version}
	ClientVersion *string `field:"optional" json:"clientVersion" yaml:"clientVersion"`
	// One or more custom audiences that you want this service to support.
	//
	// Specify each custom audience as the full URL in a string. The custom audiences are encoded in the token and used to authenticate requests.
	// For more information, see https://cloud.google.com/run/docs/configuring/custom-audiences.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#custom_audiences GoogleCloudRunV2Service#custom_audiences}
	CustomAudiences *[]*string `field:"optional" json:"customAudiences" yaml:"customAudiences"`
	// Disables public resolution of the default URI of this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#default_uri_disabled GoogleCloudRunV2Service#default_uri_disabled}
	DefaultUriDisabled interface{} `field:"optional" json:"defaultUriDisabled" yaml:"defaultUriDisabled"`
	// User-provided description of the Service. This field currently has a 512-character limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#description GoogleCloudRunV2Service#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#id GoogleCloudRunV2Service#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Provides the ingress settings for this Service.
	//
	// On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active. Possible values: ["INGRESS_TRAFFIC_ALL", "INGRESS_TRAFFIC_INTERNAL_ONLY", "INGRESS_TRAFFIC_INTERNAL_LOAD_BALANCER"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#ingress GoogleCloudRunV2Service#ingress}
	Ingress *string `field:"optional" json:"ingress" yaml:"ingress"`
	// Unstructured key value map that can be used to organize and categorize objects.
	//
	// User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component,
	// environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels.
	//
	// Cloud Run API v2 does not support labels with  'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected.
	// All system labels in v1 now have a corresponding field in v2 Service.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#labels GoogleCloudRunV2Service#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/products#product-launch-stages). Cloud Run supports ALPHA, BETA, and GA. If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features.
	//
	// For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output. Possible values: ["UNIMPLEMENTED", "PRELAUNCH", "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#launch_stage GoogleCloudRunV2Service#launch_stage}
	LaunchStage *string `field:"optional" json:"launchStage" yaml:"launchStage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#project GoogleCloudRunV2Service#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#scaling GoogleCloudRunV2Service#scaling}
	Scaling *GoogleCloudRunV2ServiceScaling `field:"optional" json:"scaling" yaml:"scaling"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#timeouts GoogleCloudRunV2Service#timeouts}
	Timeouts *GoogleCloudRunV2ServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// traffic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#traffic GoogleCloudRunV2Service#traffic}
	Traffic interface{} `field:"optional" json:"traffic" yaml:"traffic"`
}

