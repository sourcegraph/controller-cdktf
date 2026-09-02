package monitorv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MonitorV2Config struct {
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
	// The inputs map binds dataset OIDs to labels which can be referenced within stage pipelines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#inputs MonitorV2#inputs}
	Inputs *map[string]*string `field:"required" json:"inputs" yaml:"inputs"`
	// Monitor name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#name MonitorV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Describes the type of each of the rules in the definition (they must all be the same type).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#rule_kind MonitorV2#rule_kind}
	RuleKind *string `field:"required" json:"ruleKind" yaml:"ruleKind"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#rules MonitorV2#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// stage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#stage MonitorV2#stage}
	Stage interface{} `field:"required" json:"stage" yaml:"stage"`
	// OID of the workspace this object is contained in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#workspace MonitorV2#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#actions MonitorV2#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#custom_variables MonitorV2#custom_variables}.
	CustomVariables *string `field:"optional" json:"customVariables" yaml:"customVariables"`
	// expresses the minimum time that should elapse before data is considered "good enough" to evaluate.
	//
	// Choosing a delay really depends on the expectations of latency of data and whether data is expected to arrive later than other data and thus would change previously evaluated results.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#data_stabilization_delay MonitorV2#data_stabilization_delay}
	DataStabilizationDelay *string `field:"optional" json:"dataStabilizationDelay" yaml:"dataStabilizationDelay"`
	// A brief description of the monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#description MonitorV2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable/Disable the monitor (and any underlying transforms).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#disabled MonitorV2#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// groupings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#groupings MonitorV2#groupings}
	Groupings interface{} `field:"optional" json:"groupings" yaml:"groupings"`
	// URL of the monitor icon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#icon_url MonitorV2#icon_url}
	IconUrl *string `field:"optional" json:"iconUrl" yaml:"iconUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#id MonitorV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// optionally describes a duration that must be satisfied by this monitor.
	//
	// It applies to all rules, but is only applicable to rule kinds that utilize it.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#lookback_time MonitorV2#lookback_time}
	LookbackTime *string `field:"optional" json:"lookbackTime" yaml:"lookbackTime"`
	// Overrides the default value of max alerts generated in a single hour before the monitor is deactivated for safety.
	//
	// A value of 0 means "no limit". If unset, defaults to 100 (note that we use -1 in the Terraform state to indicate null/unset due to Terraform limitations).
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#max_alerts_per_hour MonitorV2#max_alerts_per_hour}
	MaxAlertsPerHour *float64 `field:"optional" json:"maxAlertsPerHour" yaml:"maxAlertsPerHour"`
	// no_data_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#no_data_rules MonitorV2#no_data_rules}
	NoDataRules *MonitorV2NoDataRules `field:"optional" json:"noDataRules" yaml:"noDataRules"`
	// scheduling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#scheduling MonitorV2#scheduling}
	Scheduling *MonitorV2Scheduling `field:"optional" json:"scheduling" yaml:"scheduling"`
}

