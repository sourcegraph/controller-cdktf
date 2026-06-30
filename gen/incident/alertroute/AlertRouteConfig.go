package alertroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlertRouteConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Which alert sources should this alert route match?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#alert_sources AlertRoute#alert_sources}
	AlertSources any `field:"required" json:"alertSources" yaml:"alertSources"`
	// Groups of prerequisite conditions. All conditions in at least one group must be satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#condition_groups AlertRoute#condition_groups}
	ConditionGroups any `field:"required" json:"conditionGroups" yaml:"conditionGroups"`
	// Whether this alert route is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#enabled AlertRoute#enabled}
	Enabled any `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#escalation_config AlertRoute#escalation_config}.
	EscalationConfig *AlertRouteEscalationConfig `field:"required" json:"escalationConfig" yaml:"escalationConfig"`
	// The expressions to be prepared for use by steps and conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#expressions AlertRoute#expressions}
	Expressions any `field:"required" json:"expressions" yaml:"expressions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#incident_config AlertRoute#incident_config}.
	IncidentConfig *AlertRouteIncidentConfig `field:"required" json:"incidentConfig" yaml:"incidentConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#incident_template AlertRoute#incident_template}.
	IncidentTemplate *AlertRouteIncidentTemplate `field:"required" json:"incidentTemplate" yaml:"incidentTemplate"`
	// Whether this alert route is private. Private alert routes will only create private incidents from alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#is_private AlertRoute#is_private}
	IsPrivate any `field:"required" json:"isPrivate" yaml:"isPrivate"`
	// The name of this alert route config, for the user's reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#name AlertRoute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The channel configuration for this alert route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#channel_config AlertRoute#channel_config}
	ChannelConfig any `field:"optional" json:"channelConfig" yaml:"channelConfig"`
}
