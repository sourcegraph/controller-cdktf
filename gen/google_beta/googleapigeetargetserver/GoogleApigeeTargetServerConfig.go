package googleapigeetargetserver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApigeeTargetServerConfig struct {
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
	// The Apigee environment group associated with the Apigee environment, in the format 'organizations/{{org_name}}/environments/{{env_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_target_server#env_id GoogleApigeeTargetServer#env_id}
	EnvId *string `field:"required" json:"envId" yaml:"envId"`
	// The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_target_server#host GoogleApigeeTargetServer#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The resource id of this reference. Values must match the regular expression [\w\s-.]+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_target_server#name GoogleApigeeTargetServer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_target_server#port GoogleApigeeTargetServer#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// A human-readable description of this TargetServer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_target_server#description GoogleApigeeTargetServer#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_target_server#id GoogleApigeeTargetServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically.
	//
	// Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_target_server#is_enabled GoogleApigeeTargetServer#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Immutable. The protocol used by this TargetServer. Possible values: ["HTTP", "HTTP2", "GRPC_TARGET", "GRPC", "EXTERNAL_CALLOUT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_target_server#protocol GoogleApigeeTargetServer#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// s_sl_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_target_server#s_sl_info GoogleApigeeTargetServer#s_sl_info}
	SSlInfo *GoogleApigeeTargetServerSSlInfo `field:"optional" json:"sSlInfo" yaml:"sSlInfo"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_target_server#timeouts GoogleApigeeTargetServer#timeouts}
	Timeouts *GoogleApigeeTargetServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

