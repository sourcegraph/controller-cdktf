package googlenetworkservicesagentgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkServicesAgentGatewayConfig struct {
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
	// The location of the agent gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_agent_gateway#location GoogleNetworkServicesAgentGateway#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of the AgentGateway resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_agent_gateway#name GoogleNetworkServicesAgentGateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of protocols supported by an Agent Gateway. Possible values: ["MCP"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_agent_gateway#protocols GoogleNetworkServicesAgentGateway#protocols}
	Protocols *[]*string `field:"required" json:"protocols" yaml:"protocols"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_agent_gateway#description GoogleNetworkServicesAgentGateway#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// google_managed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_agent_gateway#google_managed GoogleNetworkServicesAgentGateway#google_managed}
	GoogleManaged *GoogleNetworkServicesAgentGatewayGoogleManaged `field:"optional" json:"googleManaged" yaml:"googleManaged"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_agent_gateway#id GoogleNetworkServicesAgentGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the AgentGateway resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_agent_gateway#labels GoogleNetworkServicesAgentGateway#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_agent_gateway#network_config GoogleNetworkServicesAgentGateway#network_config}
	NetworkConfig *GoogleNetworkServicesAgentGatewayNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_agent_gateway#project GoogleNetworkServicesAgentGateway#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A list of Agent registries containing the agents, MCP servers and tools governed by the Agent Gateway.
	//
	// Note: Currently limited to project-scoped registries Must be of format
	// '//agentregistry.googleapis.com/{version}/projects/{{project}}/locations/{{location}}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_agent_gateway#registries GoogleNetworkServicesAgentGateway#registries}
	Registries *[]*string `field:"optional" json:"registries" yaml:"registries"`
	// self_managed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_agent_gateway#self_managed GoogleNetworkServicesAgentGateway#self_managed}
	SelfManaged *GoogleNetworkServicesAgentGatewaySelfManaged `field:"optional" json:"selfManaged" yaml:"selfManaged"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_agent_gateway#timeouts GoogleNetworkServicesAgentGateway#timeouts}
	Timeouts *GoogleNetworkServicesAgentGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

