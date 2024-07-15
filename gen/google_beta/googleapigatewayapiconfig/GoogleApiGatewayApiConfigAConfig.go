package googleapigatewayapiconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApiGatewayApiConfigAConfig struct {
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
	// The API to attach the config to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config#api GoogleApiGatewayApiConfigA#api}
	Api *string `field:"required" json:"api" yaml:"api"`
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config#api_config_id GoogleApiGatewayApiConfigA#api_config_id}
	ApiConfigId *string `field:"optional" json:"apiConfigId" yaml:"apiConfigId"`
	// Creates a unique name beginning with the specified prefix.
	//
	// If this and api_config_id are unspecified, a random value is chosen for the name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config#api_config_id_prefix GoogleApiGatewayApiConfigA#api_config_id_prefix}
	ApiConfigIdPrefix *string `field:"optional" json:"apiConfigIdPrefix" yaml:"apiConfigIdPrefix"`
	// A user-visible name for the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config#display_name GoogleApiGatewayApiConfigA#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// gateway_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config#gateway_config GoogleApiGatewayApiConfigA#gateway_config}
	GatewayConfig *GoogleApiGatewayApiConfigGatewayConfig `field:"optional" json:"gatewayConfig" yaml:"gatewayConfig"`
	// grpc_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config#grpc_services GoogleApiGatewayApiConfigA#grpc_services}
	GrpcServices interface{} `field:"optional" json:"grpcServices" yaml:"grpcServices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config#id GoogleApiGatewayApiConfigA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels to represent user-provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config#labels GoogleApiGatewayApiConfigA#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// managed_service_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config#managed_service_configs GoogleApiGatewayApiConfigA#managed_service_configs}
	ManagedServiceConfigs interface{} `field:"optional" json:"managedServiceConfigs" yaml:"managedServiceConfigs"`
	// openapi_documents block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config#openapi_documents GoogleApiGatewayApiConfigA#openapi_documents}
	OpenapiDocuments interface{} `field:"optional" json:"openapiDocuments" yaml:"openapiDocuments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config#project GoogleApiGatewayApiConfigA#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config#timeouts GoogleApiGatewayApiConfigA#timeouts}
	Timeouts *GoogleApiGatewayApiConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

