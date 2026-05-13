package googlenetworkmanagementorganizationvpcflowlogsconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkManagementOrganizationVpcFlowLogsConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_management_organization_vpc_flow_logs_config#location GoogleNetworkManagementOrganizationVpcFlowLogsConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_management_organization_vpc_flow_logs_config#organization GoogleNetworkManagementOrganizationVpcFlowLogsConfig#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Required. ID of the 'VpcFlowLogsConfig'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_management_organization_vpc_flow_logs_config#vpc_flow_logs_config_id GoogleNetworkManagementOrganizationVpcFlowLogsConfig#vpc_flow_logs_config_id}
	VpcFlowLogsConfigId *string `field:"required" json:"vpcFlowLogsConfigId" yaml:"vpcFlowLogsConfigId"`
	// Optional.
	//
	// The aggregation interval for the logs. Default value is
	// INTERVAL_5_SEC.   Possible values: INTERVAL_5_SEC INTERVAL_30_SEC INTERVAL_1_MIN INTERVAL_5_MIN INTERVAL_10_MIN INTERVAL_15_MIN
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_management_organization_vpc_flow_logs_config#aggregation_interval GoogleNetworkManagementOrganizationVpcFlowLogsConfig#aggregation_interval}
	AggregationInterval *string `field:"optional" json:"aggregationInterval" yaml:"aggregationInterval"`
	// Determines whether to include cross project annotations in the logs.
	//
	// This field is available only for organization configurations. If not
	// specified in org configs will be set to CROSS_PROJECT_METADATA_ENABLED.
	// Possible values:
	// CROSS_PROJECT_METADATA_ENABLED
	// CROSS_PROJECT_METADATA_DISABLED Possible values: ["CROSS_PROJECT_METADATA_ENABLED", "CROSS_PROJECT_METADATA_DISABLED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_management_organization_vpc_flow_logs_config#cross_project_metadata GoogleNetworkManagementOrganizationVpcFlowLogsConfig#cross_project_metadata}
	CrossProjectMetadata *string `field:"optional" json:"crossProjectMetadata" yaml:"crossProjectMetadata"`
	// Optional. The user-supplied description of the VPC Flow Logs configuration. Maximum of 512 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_management_organization_vpc_flow_logs_config#description GoogleNetworkManagementOrganizationVpcFlowLogsConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional. Export filter used to define which VPC Flow Logs should be logged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_management_organization_vpc_flow_logs_config#filter_expr GoogleNetworkManagementOrganizationVpcFlowLogsConfig#filter_expr}
	FilterExpr *string `field:"optional" json:"filterExpr" yaml:"filterExpr"`
	// Optional.
	//
	// The value of the field must be in (0, 1]. The sampling rate
	// of VPC Flow Logs where 1.0 means all collected logs are reported. Setting the
	// sampling rate to 0.0 is not allowed. If you want to disable VPC Flow Logs, use
	// the state field instead. Default value is 1.0
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_management_organization_vpc_flow_logs_config#flow_sampling GoogleNetworkManagementOrganizationVpcFlowLogsConfig#flow_sampling}
	FlowSampling *float64 `field:"optional" json:"flowSampling" yaml:"flowSampling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_management_organization_vpc_flow_logs_config#id GoogleNetworkManagementOrganizationVpcFlowLogsConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. Resource labels to represent the user-provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_management_organization_vpc_flow_logs_config#labels GoogleNetworkManagementOrganizationVpcFlowLogsConfig#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Optional.
	//
	// Configures whether all, none or a subset of metadata fields
	// should be added to the reported VPC flow logs. Default value is INCLUDE_ALL_METADATA.
	//   Possible values:  METADATA_UNSPECIFIED INCLUDE_ALL_METADATA EXCLUDE_ALL_METADATA CUSTOM_METADATA
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_management_organization_vpc_flow_logs_config#metadata GoogleNetworkManagementOrganizationVpcFlowLogsConfig#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Optional.
	//
	// Custom metadata fields to include in the reported VPC flow
	// logs. Can only be specified if \"metadata\" was set to CUSTOM_METADATA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_management_organization_vpc_flow_logs_config#metadata_fields GoogleNetworkManagementOrganizationVpcFlowLogsConfig#metadata_fields}
	MetadataFields *[]*string `field:"optional" json:"metadataFields" yaml:"metadataFields"`
	// Optional.
	//
	// The state of the VPC Flow Log configuration. Default value
	// is ENABLED. When creating a new configuration, it must be enabled.
	// Possible values: ENABLED DISABLED
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_management_organization_vpc_flow_logs_config#state GoogleNetworkManagementOrganizationVpcFlowLogsConfig#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_management_organization_vpc_flow_logs_config#timeouts GoogleNetworkManagementOrganizationVpcFlowLogsConfig#timeouts}
	Timeouts *GoogleNetworkManagementOrganizationVpcFlowLogsConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

