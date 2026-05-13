package googlecesagent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCesAgentConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#app GoogleCesAgent#app}
	App *string `field:"required" json:"app" yaml:"app"`
	// Display name of the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#display_name GoogleCesAgent#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#location GoogleCesAgent#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// after_agent_callbacks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#after_agent_callbacks GoogleCesAgent#after_agent_callbacks}
	AfterAgentCallbacks interface{} `field:"optional" json:"afterAgentCallbacks" yaml:"afterAgentCallbacks"`
	// after_model_callbacks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#after_model_callbacks GoogleCesAgent#after_model_callbacks}
	AfterModelCallbacks interface{} `field:"optional" json:"afterModelCallbacks" yaml:"afterModelCallbacks"`
	// after_tool_callbacks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#after_tool_callbacks GoogleCesAgent#after_tool_callbacks}
	AfterToolCallbacks interface{} `field:"optional" json:"afterToolCallbacks" yaml:"afterToolCallbacks"`
	// The ID to use for the agent, which will become the final component of the agent's resource name.
	//
	// If not provided, a unique ID will be
	// automatically assigned for the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#agent_id GoogleCesAgent#agent_id}
	AgentId *string `field:"optional" json:"agentId" yaml:"agentId"`
	// before_agent_callbacks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#before_agent_callbacks GoogleCesAgent#before_agent_callbacks}
	BeforeAgentCallbacks interface{} `field:"optional" json:"beforeAgentCallbacks" yaml:"beforeAgentCallbacks"`
	// before_model_callbacks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#before_model_callbacks GoogleCesAgent#before_model_callbacks}
	BeforeModelCallbacks interface{} `field:"optional" json:"beforeModelCallbacks" yaml:"beforeModelCallbacks"`
	// before_tool_callbacks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#before_tool_callbacks GoogleCesAgent#before_tool_callbacks}
	BeforeToolCallbacks interface{} `field:"optional" json:"beforeToolCallbacks" yaml:"beforeToolCallbacks"`
	// List of child agents in the agent tree. Format: 'projects/{project}/locations/{location}/apps/{app}/agents/{agent}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#child_agents GoogleCesAgent#child_agents}
	ChildAgents *[]*string `field:"optional" json:"childAgents" yaml:"childAgents"`
	// Human-readable description of the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#description GoogleCesAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of guardrails for the agent. Format: 'projects/{project}/locations/{location}/apps/{app}/guardrails/{guardrail}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#guardrails GoogleCesAgent#guardrails}
	Guardrails *[]*string `field:"optional" json:"guardrails" yaml:"guardrails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#id GoogleCesAgent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Instructions for the LLM model to guide the agent's behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#instruction GoogleCesAgent#instruction}
	Instruction *string `field:"optional" json:"instruction" yaml:"instruction"`
	// llm_agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#llm_agent GoogleCesAgent#llm_agent}
	LlmAgent *GoogleCesAgentLlmAgent `field:"optional" json:"llmAgent" yaml:"llmAgent"`
	// model_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#model_settings GoogleCesAgent#model_settings}
	ModelSettings *GoogleCesAgentModelSettings `field:"optional" json:"modelSettings" yaml:"modelSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#project GoogleCesAgent#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// remote_dialogflow_agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#remote_dialogflow_agent GoogleCesAgent#remote_dialogflow_agent}
	RemoteDialogflowAgent *GoogleCesAgentRemoteDialogflowAgent `field:"optional" json:"remoteDialogflowAgent" yaml:"remoteDialogflowAgent"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#timeouts GoogleCesAgent#timeouts}
	Timeouts *GoogleCesAgentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// List of available tools for the agent. Format: 'projects/{project}/locations/{location}/apps/{app}/tools/{tool}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#tools GoogleCesAgent#tools}
	Tools *[]*string `field:"optional" json:"tools" yaml:"tools"`
	// toolsets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_agent#toolsets GoogleCesAgent#toolsets}
	Toolsets interface{} `field:"optional" json:"toolsets" yaml:"toolsets"`
}

