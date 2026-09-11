package workflow

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkflowConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Groups of prerequisite conditions. All conditions in at least one group must be satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#condition_groups Workflow#condition_groups}
	ConditionGroups interface{} `field:"required" json:"conditionGroups" yaml:"conditionGroups"`
	// Whether to continue executing the workflow if a step fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#continue_on_step_error Workflow#continue_on_step_error}
	ContinueOnStepError interface{} `field:"required" json:"continueOnStepError" yaml:"continueOnStepError"`
	// The expressions to be prepared for use by steps and conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#expressions Workflow#expressions}
	Expressions interface{} `field:"required" json:"expressions" yaml:"expressions"`
	// Whether to include private incidents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#include_private_incidents Workflow#include_private_incidents}
	IncludePrivateIncidents interface{} `field:"required" json:"includePrivateIncidents" yaml:"includePrivateIncidents"`
	// Name provided by the user when creating the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#name Workflow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// This workflow will run 'once for' a list of references.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#once_for Workflow#once_for}
	OnceFor *[]*string `field:"required" json:"onceFor" yaml:"onceFor"`
	// Which incident modes should this workflow run on?
	//
	// By default, workflows only run on standard incidents, but can also be configured to run on test and retrospective incidents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#runs_on_incident_modes Workflow#runs_on_incident_modes}
	RunsOnIncidentModes *[]*string `field:"required" json:"runsOnIncidentModes" yaml:"runsOnIncidentModes"`
	// Which incidents should the workflow be applied to?. Possible values are: `newly_created`, `newly_created_and_active`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#runs_on_incidents Workflow#runs_on_incidents}
	RunsOnIncidents *string `field:"required" json:"runsOnIncidents" yaml:"runsOnIncidents"`
	// What state this workflow is in. Possible values are: `active`, `disabled`, `draft`, `error`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#state Workflow#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// Steps that are executed as part of the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#steps Workflow#steps}
	Steps interface{} `field:"required" json:"steps" yaml:"steps"`
	// Unique name of the trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#trigger Workflow#trigger}
	Trigger *string `field:"required" json:"trigger" yaml:"trigger"`
	// Configuration controlling workflow delay behaviour.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#delay Workflow#delay}
	Delay *WorkflowDelay `field:"optional" json:"delay" yaml:"delay"`
	// Folder to display the workflow in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#folder Workflow#folder}
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// The shortform used to trigger this workflow (only applicable for manual triggers).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#shortform Workflow#shortform}
	Shortform *string `field:"optional" json:"shortform" yaml:"shortform"`
}

