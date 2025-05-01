package googlenetworkmanagementvpcflowlogsconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkManagementVpcFlowLogsConfigConfig struct {
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
	// Resource ID segment making up resource 'name'.
	//
	// It identifies the resource
	// within its parent collection as described in https://google.aip.dev/122. See documentation
	// for resource type 'networkmanagement.googleapis.com/VpcFlowLogsConfig'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#location GoogleNetworkManagementVpcFlowLogsConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Required. ID of the 'VpcFlowLogsConfig'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#vpc_flow_logs_config_id GoogleNetworkManagementVpcFlowLogsConfig#vpc_flow_logs_config_id}
	VpcFlowLogsConfigId *string `field:"required" json:"vpcFlowLogsConfigId" yaml:"vpcFlowLogsConfigId"`
	// Optional.
	//
	// The aggregation interval for the logs. Default value is
	// INTERVAL_5_SEC.   Possible values:  AGGREGATION_INTERVAL_UNSPECIFIED INTERVAL_5_SEC INTERVAL_30_SEC INTERVAL_1_MIN INTERVAL_5_MIN INTERVAL_10_MIN INTERVAL_15_MIN"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#aggregation_interval GoogleNetworkManagementVpcFlowLogsConfig#aggregation_interval}
	AggregationInterval *string `field:"optional" json:"aggregationInterval" yaml:"aggregationInterval"`
	// Optional. The user-supplied description of the VPC Flow Logs configuration. Maximum of 512 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#description GoogleNetworkManagementVpcFlowLogsConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional. Export filter used to define which VPC Flow Logs should be logged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#filter_expr GoogleNetworkManagementVpcFlowLogsConfig#filter_expr}
	FilterExpr *string `field:"optional" json:"filterExpr" yaml:"filterExpr"`
	// Optional.
	//
	// The value of the field must be in (0, 1]. The sampling rate
	// of VPC Flow Logs where 1.0 means all collected logs are reported. Setting the
	// sampling rate to 0.0 is not allowed. If you want to disable VPC Flow Logs, use
	// the state field instead. Default value is 1.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#flow_sampling GoogleNetworkManagementVpcFlowLogsConfig#flow_sampling}
	FlowSampling *float64 `field:"optional" json:"flowSampling" yaml:"flowSampling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#id GoogleNetworkManagementVpcFlowLogsConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Traffic will be logged from the Interconnect Attachment. Format: projects/{project_id}/regions/{region}/interconnectAttachments/{name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#interconnect_attachment GoogleNetworkManagementVpcFlowLogsConfig#interconnect_attachment}
	InterconnectAttachment *string `field:"optional" json:"interconnectAttachment" yaml:"interconnectAttachment"`
	// Optional. Resource labels to represent user-provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#labels GoogleNetworkManagementVpcFlowLogsConfig#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Optional.
	//
	// Configures whether all, none or a subset of metadata fields
	// should be added to the reported VPC flow logs. Default value is INCLUDE_ALL_METADATA.
	//   Possible values:  METADATA_UNSPECIFIED INCLUDE_ALL_METADATA EXCLUDE_ALL_METADATA CUSTOM_METADATA
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#metadata GoogleNetworkManagementVpcFlowLogsConfig#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Optional.
	//
	// Custom metadata fields to include in the reported VPC flow
	// logs. Can only be specified if \"metadata\" was set to CUSTOM_METADATA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#metadata_fields GoogleNetworkManagementVpcFlowLogsConfig#metadata_fields}
	MetadataFields *[]*string `field:"optional" json:"metadataFields" yaml:"metadataFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#project GoogleNetworkManagementVpcFlowLogsConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Optional.
	//
	// The state of the VPC Flow Log configuration. Default value
	// is ENABLED. When creating a new configuration, it must be enabled.   Possible
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#state GoogleNetworkManagementVpcFlowLogsConfig#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#timeouts GoogleNetworkManagementVpcFlowLogsConfig#timeouts}
	Timeouts *GoogleNetworkManagementVpcFlowLogsConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Traffic will be logged from the VPN Tunnel. Format: projects/{project_id}/regions/{region}/vpnTunnels/{name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_management_vpc_flow_logs_config#vpn_tunnel GoogleNetworkManagementVpcFlowLogsConfig#vpn_tunnel}
	VpnTunnel *string `field:"optional" json:"vpnTunnel" yaml:"vpnTunnel"`
}

