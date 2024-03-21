package awsvpc

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AwsvpcConfig struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// The Autonomous System Number (ASN) for the Amazon side of the gateway.
	//
	// By default the virtual private gateway is created with the current default Amazon ASN.
	// 64512.
	AmazonSideAsn *string `field:"optional" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// Assign IPv6 address on subnet, must be disabled to change IPv6 CIDRs.
	//
	// This is the IPv6 equivalent of map_public_ip_on_launch.
	AssignIpv6AddressOnCreation *bool `field:"optional" json:"assignIpv6AddressOnCreation" yaml:"assignIpv6AddressOnCreation"`
	// A list of availability zones names or ids in the region.
	Azs *[]*string `field:"optional" json:"azs" yaml:"azs"`
	// (Optional) The IPv4 CIDR block for the VPC.
	//
	// CIDR can be explicitly set or it can be derived from IPAM using `ipv4_netmask_length` & `ipv4_ipam_pool_id`
	// 0.0.0.0/0
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Controls if an internet gateway route for public database access should be created.
	CreateDatabaseInternetGatewayRoute *bool `field:"optional" json:"createDatabaseInternetGatewayRoute" yaml:"createDatabaseInternetGatewayRoute"`
	// Controls if a nat gateway route should be created to give internet access to the database subnets.
	CreateDatabaseNatGatewayRoute *bool `field:"optional" json:"createDatabaseNatGatewayRoute" yaml:"createDatabaseNatGatewayRoute"`
	// Controls if database subnet group should be created (n.b. database_subnets must also be set) true.
	CreateDatabaseSubnetGroup *bool `field:"optional" json:"createDatabaseSubnetGroup" yaml:"createDatabaseSubnetGroup"`
	// Controls if separate route table for database should be created.
	CreateDatabaseSubnetRouteTable *bool `field:"optional" json:"createDatabaseSubnetRouteTable" yaml:"createDatabaseSubnetRouteTable"`
	// Controls if an Egress Only Internet Gateway is created and its related routes.
	//
	// true.
	CreateEgressOnlyIgw *bool `field:"optional" json:"createEgressOnlyIgw" yaml:"createEgressOnlyIgw"`
	// Controls if elasticache subnet group should be created true.
	CreateElasticacheSubnetGroup *bool `field:"optional" json:"createElasticacheSubnetGroup" yaml:"createElasticacheSubnetGroup"`
	// Controls if separate route table for elasticache should be created.
	CreateElasticacheSubnetRouteTable *bool `field:"optional" json:"createElasticacheSubnetRouteTable" yaml:"createElasticacheSubnetRouteTable"`
	// Whether to create IAM role for VPC Flow Logs.
	CreateFlowLogCloudwatchIamRole *bool `field:"optional" json:"createFlowLogCloudwatchIamRole" yaml:"createFlowLogCloudwatchIamRole"`
	// Whether to create CloudWatch log group for VPC Flow Logs.
	CreateFlowLogCloudwatchLogGroup *bool `field:"optional" json:"createFlowLogCloudwatchLogGroup" yaml:"createFlowLogCloudwatchLogGroup"`
	// Controls if an Internet Gateway is created for public subnets and the related routes that connect them.
	//
	// true.
	CreateIgw *bool `field:"optional" json:"createIgw" yaml:"createIgw"`
	// Controls if redshift subnet group should be created true.
	CreateRedshiftSubnetGroup *bool `field:"optional" json:"createRedshiftSubnetGroup" yaml:"createRedshiftSubnetGroup"`
	// Controls if separate route table for redshift should be created.
	CreateRedshiftSubnetRouteTable *bool `field:"optional" json:"createRedshiftSubnetRouteTable" yaml:"createRedshiftSubnetRouteTable"`
	// Controls if VPC should be created (it affects almost all resources) true.
	CreateVpc *bool `field:"optional" json:"createVpc" yaml:"createVpc"`
	// Maps of Customer Gateway's attributes (BGP ASN and Gateway's Internet-routable external IP address) The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	CustomerGateways *map[string]*map[string]interface{} `field:"optional" json:"customerGateways" yaml:"customerGateways"`
	// Additional tags for the Customer Gateway The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	CustomerGatewayTags *map[string]*string `field:"optional" json:"customerGatewayTags" yaml:"customerGatewayTags"`
	// Additional tags for the database subnets network ACL The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	DatabaseAclTags *map[string]*string `field:"optional" json:"databaseAclTags" yaml:"databaseAclTags"`
	// Whether to use dedicated network ACL (not default) and custom rules for database subnets.
	DatabaseDedicatedNetworkAcl *bool `field:"optional" json:"databaseDedicatedNetworkAcl" yaml:"databaseDedicatedNetworkAcl"`
	// Database subnets inbound network ACL rules [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	DatabaseInboundAclRules *[]*map[string]*string `field:"optional" json:"databaseInboundAclRules" yaml:"databaseInboundAclRules"`
	// Database subnets outbound network ACL rules [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	DatabaseOutboundAclRules *[]*map[string]*string `field:"optional" json:"databaseOutboundAclRules" yaml:"databaseOutboundAclRules"`
	// Additional tags for the database route tables The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	DatabaseRouteTableTags *map[string]*string `field:"optional" json:"databaseRouteTableTags" yaml:"databaseRouteTableTags"`
	// Assign IPv6 address on database subnet, must be disabled to change IPv6 CIDRs.
	//
	// This is the IPv6 equivalent of map_public_ip_on_launch.
	DatabaseSubnetAssignIpv6AddressOnCreation *bool `field:"optional" json:"databaseSubnetAssignIpv6AddressOnCreation" yaml:"databaseSubnetAssignIpv6AddressOnCreation"`
	// Name of database subnet group.
	DatabaseSubnetGroupName *string `field:"optional" json:"databaseSubnetGroupName" yaml:"databaseSubnetGroupName"`
	// Additional tags for the database subnet group The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	DatabaseSubnetGroupTags *map[string]*string `field:"optional" json:"databaseSubnetGroupTags" yaml:"databaseSubnetGroupTags"`
	// Assigns IPv6 database subnet id based on the Amazon provided /56 prefix base 10 integer (0-256).
	//
	// Must be of equal length to the corresponding IPv4 subnet list.
	DatabaseSubnetIpv6Prefixes *[]*string `field:"optional" json:"databaseSubnetIpv6Prefixes" yaml:"databaseSubnetIpv6Prefixes"`
	// Explicit values to use in the Name tag on database subnets.
	//
	// If empty, Name tags are generated.
	DatabaseSubnetNames *[]*string `field:"optional" json:"databaseSubnetNames" yaml:"databaseSubnetNames"`
	// A list of database subnets.
	DatabaseSubnets *[]*string `field:"optional" json:"databaseSubnets" yaml:"databaseSubnets"`
	// Suffix to append to database subnets name db.
	DatabaseSubnetSuffix *string `field:"optional" json:"databaseSubnetSuffix" yaml:"databaseSubnetSuffix"`
	// Additional tags for the database subnets The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	DatabaseSubnetTags *map[string]*string `field:"optional" json:"databaseSubnetTags" yaml:"databaseSubnetTags"`
	// List of maps of egress rules to set on the Default Network ACL [object Object] [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	DefaultNetworkAclEgress *[]*map[string]*string `field:"optional" json:"defaultNetworkAclEgress" yaml:"defaultNetworkAclEgress"`
	// List of maps of ingress rules to set on the Default Network ACL [object Object] [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	DefaultNetworkAclIngress *[]*map[string]*string `field:"optional" json:"defaultNetworkAclIngress" yaml:"defaultNetworkAclIngress"`
	// Name to be used on the Default Network ACL.
	DefaultNetworkAclName *string `field:"optional" json:"defaultNetworkAclName" yaml:"defaultNetworkAclName"`
	// Additional tags for the Default Network ACL The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	DefaultNetworkAclTags *map[string]*string `field:"optional" json:"defaultNetworkAclTags" yaml:"defaultNetworkAclTags"`
	// Name to be used on the default route table.
	DefaultRouteTableName *string `field:"optional" json:"defaultRouteTableName" yaml:"defaultRouteTableName"`
	// List of virtual gateways for propagation.
	DefaultRouteTablePropagatingVgws *[]*string `field:"optional" json:"defaultRouteTablePropagatingVgws" yaml:"defaultRouteTablePropagatingVgws"`
	// Configuration block of routes. See https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/default_route_table#route.
	//
	// The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}
	DefaultRouteTableRoutes *[]*map[string]*string `field:"optional" json:"defaultRouteTableRoutes" yaml:"defaultRouteTableRoutes"`
	// Additional tags for the default route table The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	DefaultRouteTableTags *map[string]*string `field:"optional" json:"defaultRouteTableTags" yaml:"defaultRouteTableTags"`
	// List of maps of egress rules to set on the default security group.
	//
	// The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}
	DefaultSecurityGroupEgress *[]*map[string]*string `field:"optional" json:"defaultSecurityGroupEgress" yaml:"defaultSecurityGroupEgress"`
	// List of maps of ingress rules to set on the default security group.
	//
	// The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}
	DefaultSecurityGroupIngress *[]*map[string]*string `field:"optional" json:"defaultSecurityGroupIngress" yaml:"defaultSecurityGroupIngress"`
	// Name to be used on the default security group.
	DefaultSecurityGroupName *string `field:"optional" json:"defaultSecurityGroupName" yaml:"defaultSecurityGroupName"`
	// Additional tags for the default security group The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	DefaultSecurityGroupTags *map[string]*string `field:"optional" json:"defaultSecurityGroupTags" yaml:"defaultSecurityGroupTags"`
	// [DEPRECATED](https://github.com/hashicorp/terraform/issues/31730) Should be true to enable ClassicLink in the Default VPC.
	DefaultVpcEnableClassiclink *bool `field:"optional" json:"defaultVpcEnableClassiclink" yaml:"defaultVpcEnableClassiclink"`
	// Should be true to enable DNS hostnames in the Default VPC.
	DefaultVpcEnableDnsHostnames *bool `field:"optional" json:"defaultVpcEnableDnsHostnames" yaml:"defaultVpcEnableDnsHostnames"`
	// Should be true to enable DNS support in the Default VPC true.
	DefaultVpcEnableDnsSupport *bool `field:"optional" json:"defaultVpcEnableDnsSupport" yaml:"defaultVpcEnableDnsSupport"`
	// Name to be used on the Default VPC.
	DefaultVpcName *string `field:"optional" json:"defaultVpcName" yaml:"defaultVpcName"`
	// Additional tags for the Default VPC The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	DefaultVpcTags *map[string]*string `field:"optional" json:"defaultVpcTags" yaml:"defaultVpcTags"`
	// Specifies DNS name for DHCP options set (requires enable_dhcp_options set to true).
	DhcpOptionsDomainName *string `field:"optional" json:"dhcpOptionsDomainName" yaml:"dhcpOptionsDomainName"`
	// Specify a list of DNS server addresses for DHCP options set, default to AWS provided (requires enable_dhcp_options set to true) AmazonProvidedDNS.
	DhcpOptionsDomainNameServers *[]*string `field:"optional" json:"dhcpOptionsDomainNameServers" yaml:"dhcpOptionsDomainNameServers"`
	// Specify a list of netbios servers for DHCP options set (requires enable_dhcp_options set to true).
	DhcpOptionsNetbiosNameServers *[]*string `field:"optional" json:"dhcpOptionsNetbiosNameServers" yaml:"dhcpOptionsNetbiosNameServers"`
	// Specify netbios node_type for DHCP options set (requires enable_dhcp_options set to true).
	DhcpOptionsNetbiosNodeType *string `field:"optional" json:"dhcpOptionsNetbiosNodeType" yaml:"dhcpOptionsNetbiosNodeType"`
	// Specify a list of NTP servers for DHCP options set (requires enable_dhcp_options set to true).
	DhcpOptionsNtpServers *[]*string `field:"optional" json:"dhcpOptionsNtpServers" yaml:"dhcpOptionsNtpServers"`
	// Additional tags for the DHCP option set (requires enable_dhcp_options set to true) The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	DhcpOptionsTags *map[string]*string `field:"optional" json:"dhcpOptionsTags" yaml:"dhcpOptionsTags"`
	// Additional tags for the elasticache subnets network ACL The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	ElasticacheAclTags *map[string]*string `field:"optional" json:"elasticacheAclTags" yaml:"elasticacheAclTags"`
	// Whether to use dedicated network ACL (not default) and custom rules for elasticache subnets.
	ElasticacheDedicatedNetworkAcl *bool `field:"optional" json:"elasticacheDedicatedNetworkAcl" yaml:"elasticacheDedicatedNetworkAcl"`
	// Elasticache subnets inbound network ACL rules [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	ElasticacheInboundAclRules *[]*map[string]*string `field:"optional" json:"elasticacheInboundAclRules" yaml:"elasticacheInboundAclRules"`
	// Elasticache subnets outbound network ACL rules [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	ElasticacheOutboundAclRules *[]*map[string]*string `field:"optional" json:"elasticacheOutboundAclRules" yaml:"elasticacheOutboundAclRules"`
	// Additional tags for the elasticache route tables The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	ElasticacheRouteTableTags *map[string]*string `field:"optional" json:"elasticacheRouteTableTags" yaml:"elasticacheRouteTableTags"`
	// Assign IPv6 address on elasticache subnet, must be disabled to change IPv6 CIDRs.
	//
	// This is the IPv6 equivalent of map_public_ip_on_launch.
	ElasticacheSubnetAssignIpv6AddressOnCreation *bool `field:"optional" json:"elasticacheSubnetAssignIpv6AddressOnCreation" yaml:"elasticacheSubnetAssignIpv6AddressOnCreation"`
	// Name of elasticache subnet group.
	ElasticacheSubnetGroupName *string `field:"optional" json:"elasticacheSubnetGroupName" yaml:"elasticacheSubnetGroupName"`
	// Additional tags for the elasticache subnet group The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	ElasticacheSubnetGroupTags *map[string]*string `field:"optional" json:"elasticacheSubnetGroupTags" yaml:"elasticacheSubnetGroupTags"`
	// Assigns IPv6 elasticache subnet id based on the Amazon provided /56 prefix base 10 integer (0-256).
	//
	// Must be of equal length to the corresponding IPv4 subnet list.
	ElasticacheSubnetIpv6Prefixes *[]*string `field:"optional" json:"elasticacheSubnetIpv6Prefixes" yaml:"elasticacheSubnetIpv6Prefixes"`
	// Explicit values to use in the Name tag on elasticache subnets.
	//
	// If empty, Name tags are generated.
	ElasticacheSubnetNames *[]*string `field:"optional" json:"elasticacheSubnetNames" yaml:"elasticacheSubnetNames"`
	// A list of elasticache subnets.
	ElasticacheSubnets *[]*string `field:"optional" json:"elasticacheSubnets" yaml:"elasticacheSubnets"`
	// Suffix to append to elasticache subnets name elasticache.
	ElasticacheSubnetSuffix *string `field:"optional" json:"elasticacheSubnetSuffix" yaml:"elasticacheSubnetSuffix"`
	// Additional tags for the elasticache subnets The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	ElasticacheSubnetTags *map[string]*string `field:"optional" json:"elasticacheSubnetTags" yaml:"elasticacheSubnetTags"`
	// [DEPRECATED](https://github.com/hashicorp/terraform/issues/31730) Should be true to enable ClassicLink for the VPC. Only valid in regions and accounts that support EC2 Classic.
	EnableClassiclink *bool `field:"optional" json:"enableClassiclink" yaml:"enableClassiclink"`
	// [DEPRECATED](https://github.com/hashicorp/terraform/issues/31730) Should be true to enable ClassicLink DNS Support for the VPC. Only valid in regions and accounts that support EC2 Classic.
	EnableClassiclinkDnsSupport *bool `field:"optional" json:"enableClassiclinkDnsSupport" yaml:"enableClassiclinkDnsSupport"`
	// Should be true if you want to specify a DHCP options set with a custom domain name, DNS servers, NTP servers, netbios servers, and/or netbios server type.
	EnableDhcpOptions *bool `field:"optional" json:"enableDhcpOptions" yaml:"enableDhcpOptions"`
	// Should be true to enable DNS hostnames in the VPC.
	EnableDnsHostnames *bool `field:"optional" json:"enableDnsHostnames" yaml:"enableDnsHostnames"`
	// Should be true to enable DNS support in the VPC true.
	EnableDnsSupport *bool `field:"optional" json:"enableDnsSupport" yaml:"enableDnsSupport"`
	// Whether or not to enable VPC Flow Logs.
	EnableFlowLog *bool `field:"optional" json:"enableFlowLog" yaml:"enableFlowLog"`
	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC.
	//
	// You cannot specify the range of IP addresses, or the size of the CIDR block.
	EnableIpv6 *bool `field:"optional" json:"enableIpv6" yaml:"enableIpv6"`
	// Should be true if you want to provision NAT Gateways for each of your private networks.
	EnableNatGateway *bool `field:"optional" json:"enableNatGateway" yaml:"enableNatGateway"`
	// Controls if redshift should have public routing table.
	EnablePublicRedshift *bool `field:"optional" json:"enablePublicRedshift" yaml:"enablePublicRedshift"`
	// Should be true if you want to create a new VPN Gateway resource and attach it to the VPC.
	EnableVpnGateway *bool `field:"optional" json:"enableVpnGateway" yaml:"enableVpnGateway"`
	// List of EIP IDs to be assigned to the NAT Gateways (used in combination with reuse_nat_ips).
	ExternalNatIpIds *[]*string `field:"optional" json:"externalNatIpIds" yaml:"externalNatIpIds"`
	// List of EIPs to be used for `nat_public_ips` output (used in combination with reuse_nat_ips and external_nat_ip_ids).
	ExternalNatIps *[]*string `field:"optional" json:"externalNatIps" yaml:"externalNatIps"`
	// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group.
	//
	// When flow_log_destination_arn is set to ARN of Cloudwatch Logs, this argument needs to be provided.
	FlowLogCloudwatchIamRoleArn *string `field:"optional" json:"flowLogCloudwatchIamRoleArn" yaml:"flowLogCloudwatchIamRoleArn"`
	// The ARN of the KMS Key to use when encrypting log data for VPC flow logs.
	FlowLogCloudwatchLogGroupKmsKeyId *string `field:"optional" json:"flowLogCloudwatchLogGroupKmsKeyId" yaml:"flowLogCloudwatchLogGroupKmsKeyId"`
	// Specifies the name prefix of CloudWatch Log Group for VPC flow logs.
	//
	// /aws/vpc-flow-log/.
	FlowLogCloudwatchLogGroupNamePrefix *string `field:"optional" json:"flowLogCloudwatchLogGroupNamePrefix" yaml:"flowLogCloudwatchLogGroupNamePrefix"`
	// Specifies the name suffix of CloudWatch Log Group for VPC flow logs.
	FlowLogCloudwatchLogGroupNameSuffix *string `field:"optional" json:"flowLogCloudwatchLogGroupNameSuffix" yaml:"flowLogCloudwatchLogGroupNameSuffix"`
	// Specifies the number of days you want to retain log events in the specified log group for VPC flow logs.
	FlowLogCloudwatchLogGroupRetentionInDays *float64 `field:"optional" json:"flowLogCloudwatchLogGroupRetentionInDays" yaml:"flowLogCloudwatchLogGroupRetentionInDays"`
	// The ARN of the CloudWatch log group or S3 bucket where VPC Flow Logs will be pushed.
	//
	// If this ARN is a S3 bucket the appropriate permissions need to be set on that bucket's policy. When create_flow_log_cloudwatch_log_group is set to false this argument must be provided.
	FlowLogDestinationArn *string `field:"optional" json:"flowLogDestinationArn" yaml:"flowLogDestinationArn"`
	// Type of flow log destination.
	//
	// Can be s3 or cloud-watch-logs.
	// cloud-watch-logs.
	FlowLogDestinationType *string `field:"optional" json:"flowLogDestinationType" yaml:"flowLogDestinationType"`
	// (Optional) The format for the flow log.
	//
	// Valid values: `plain-text`, `parquet`.
	// plain-text.
	FlowLogFileFormat *string `field:"optional" json:"flowLogFileFormat" yaml:"flowLogFileFormat"`
	// (Optional) Indicates whether to use Hive-compatible prefixes for flow logs stored in Amazon S3.
	FlowLogHiveCompatiblePartitions *bool `field:"optional" json:"flowLogHiveCompatiblePartitions" yaml:"flowLogHiveCompatiblePartitions"`
	// The fields to include in the flow log record, in the order in which they should appear.
	FlowLogLogFormat *string `field:"optional" json:"flowLogLogFormat" yaml:"flowLogLogFormat"`
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.
	//
	// Valid Values: `60` seconds or `600` seconds.
	// 600.
	FlowLogMaxAggregationInterval *float64 `field:"optional" json:"flowLogMaxAggregationInterval" yaml:"flowLogMaxAggregationInterval"`
	// (Optional) Indicates whether to partition the flow log per hour.
	//
	// This reduces the cost and response time for queries.
	FlowLogPerHourPartition *bool `field:"optional" json:"flowLogPerHourPartition" yaml:"flowLogPerHourPartition"`
	// The type of traffic to capture.
	//
	// Valid values: ACCEPT, REJECT, ALL.
	// ALL.
	FlowLogTrafficType *string `field:"optional" json:"flowLogTrafficType" yaml:"flowLogTrafficType"`
	// Additional tags for the internet gateway The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	IgwTags *map[string]*string `field:"optional" json:"igwTags" yaml:"igwTags"`
	// A tenancy option for instances launched into the VPC default.
	InstanceTenancy *string `field:"optional" json:"instanceTenancy" yaml:"instanceTenancy"`
	// Additional tags for the intra subnets network ACL The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	IntraAclTags *map[string]*string `field:"optional" json:"intraAclTags" yaml:"intraAclTags"`
	// Whether to use dedicated network ACL (not default) and custom rules for intra subnets.
	IntraDedicatedNetworkAcl *bool `field:"optional" json:"intraDedicatedNetworkAcl" yaml:"intraDedicatedNetworkAcl"`
	// Intra subnets inbound network ACLs [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	IntraInboundAclRules *[]*map[string]*string `field:"optional" json:"intraInboundAclRules" yaml:"intraInboundAclRules"`
	// Intra subnets outbound network ACLs [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	IntraOutboundAclRules *[]*map[string]*string `field:"optional" json:"intraOutboundAclRules" yaml:"intraOutboundAclRules"`
	// Additional tags for the intra route tables The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	IntraRouteTableTags *map[string]*string `field:"optional" json:"intraRouteTableTags" yaml:"intraRouteTableTags"`
	// Assign IPv6 address on intra subnet, must be disabled to change IPv6 CIDRs.
	//
	// This is the IPv6 equivalent of map_public_ip_on_launch.
	IntraSubnetAssignIpv6AddressOnCreation *bool `field:"optional" json:"intraSubnetAssignIpv6AddressOnCreation" yaml:"intraSubnetAssignIpv6AddressOnCreation"`
	// Assigns IPv6 intra subnet id based on the Amazon provided /56 prefix base 10 integer (0-256).
	//
	// Must be of equal length to the corresponding IPv4 subnet list.
	IntraSubnetIpv6Prefixes *[]*string `field:"optional" json:"intraSubnetIpv6Prefixes" yaml:"intraSubnetIpv6Prefixes"`
	// Explicit values to use in the Name tag on intra subnets.
	//
	// If empty, Name tags are generated.
	IntraSubnetNames *[]*string `field:"optional" json:"intraSubnetNames" yaml:"intraSubnetNames"`
	// A list of intra subnets.
	IntraSubnets *[]*string `field:"optional" json:"intraSubnets" yaml:"intraSubnets"`
	// Suffix to append to intra subnets name intra.
	IntraSubnetSuffix *string `field:"optional" json:"intraSubnetSuffix" yaml:"intraSubnetSuffix"`
	// Additional tags for the intra subnets The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	IntraSubnetTags *map[string]*string `field:"optional" json:"intraSubnetTags" yaml:"intraSubnetTags"`
	// (Optional) The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR.
	Ipv4IpamPoolId *string `field:"optional" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
	// (Optional) The netmask length of the IPv4 CIDR you want to allocate to this VPC.
	//
	// Requires specifying a ipv4_ipam_pool_id.
	Ipv4NetmaskLength *float64 `field:"optional" json:"ipv4NetmaskLength" yaml:"ipv4NetmaskLength"`
	// (Optional) IPv6 CIDR block to request from an IPAM Pool.
	//
	// Can be set explicitly or derived from IPAM using `ipv6_netmask_length`.
	Ipv6Cidr *string `field:"optional" json:"ipv6Cidr" yaml:"ipv6Cidr"`
	// (Optional) IPAM Pool ID for a IPv6 pool.
	//
	// Conflicts with `assign_generated_ipv6_cidr_block`.
	Ipv6IpamPoolId *string `field:"optional" json:"ipv6IpamPoolId" yaml:"ipv6IpamPoolId"`
	// (Optional) Netmask length to request from IPAM Pool.
	//
	// Conflicts with `ipv6_cidr_block`. This can be omitted if IPAM pool as a `allocation_default_netmask_length` set. Valid values: `56`.
	Ipv6NetmaskLength *float64 `field:"optional" json:"ipv6NetmaskLength" yaml:"ipv6NetmaskLength"`
	// Should be true to adopt and manage Default Network ACL.
	ManageDefaultNetworkAcl *bool `field:"optional" json:"manageDefaultNetworkAcl" yaml:"manageDefaultNetworkAcl"`
	// Should be true to manage default route table.
	ManageDefaultRouteTable *bool `field:"optional" json:"manageDefaultRouteTable" yaml:"manageDefaultRouteTable"`
	// Should be true to adopt and manage default security group.
	ManageDefaultSecurityGroup *bool `field:"optional" json:"manageDefaultSecurityGroup" yaml:"manageDefaultSecurityGroup"`
	// Should be true to adopt and manage Default VPC.
	ManageDefaultVpc *bool `field:"optional" json:"manageDefaultVpc" yaml:"manageDefaultVpc"`
	// Should be false if you do not want to auto-assign public IP on launch true.
	MapPublicIpOnLaunch *bool `field:"optional" json:"mapPublicIpOnLaunch" yaml:"mapPublicIpOnLaunch"`
	// Name to be used on all the resources as identifier.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Additional tags for the NAT EIP The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	NatEipTags *map[string]*string `field:"optional" json:"natEipTags" yaml:"natEipTags"`
	// Used to pass a custom destination route for private NAT Gateway.
	//
	// If not specified, the default 0.0.0.0/0 is used as a destination route.
	// 0.0.0.0/0
	NatGatewayDestinationCidrBlock *string `field:"optional" json:"natGatewayDestinationCidrBlock" yaml:"natGatewayDestinationCidrBlock"`
	// Additional tags for the NAT gateways The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	NatGatewayTags *map[string]*string `field:"optional" json:"natGatewayTags" yaml:"natGatewayTags"`
	// Should be true if you want only one NAT Gateway per availability zone.
	//
	// Requires `var.azs` to be set, and the number of `public_subnets` created to be greater than or equal to the number of availability zones specified in `var.azs`.
	OneNatGatewayPerAz *bool `field:"optional" json:"oneNatGatewayPerAz" yaml:"oneNatGatewayPerAz"`
	// Additional tags for the outpost subnets network ACL The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	OutpostAclTags *map[string]*string `field:"optional" json:"outpostAclTags" yaml:"outpostAclTags"`
	// ARN of Outpost you want to create a subnet in.
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// AZ where Outpost is anchored.
	OutpostAz *string `field:"optional" json:"outpostAz" yaml:"outpostAz"`
	// Whether to use dedicated network ACL (not default) and custom rules for outpost subnets.
	OutpostDedicatedNetworkAcl *bool `field:"optional" json:"outpostDedicatedNetworkAcl" yaml:"outpostDedicatedNetworkAcl"`
	// Outpost subnets inbound network ACLs [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	OutpostInboundAclRules *[]*map[string]*string `field:"optional" json:"outpostInboundAclRules" yaml:"outpostInboundAclRules"`
	// Outpost subnets outbound network ACLs [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	OutpostOutboundAclRules *[]*map[string]*string `field:"optional" json:"outpostOutboundAclRules" yaml:"outpostOutboundAclRules"`
	// Assign IPv6 address on outpost subnet, must be disabled to change IPv6 CIDRs.
	//
	// This is the IPv6 equivalent of map_public_ip_on_launch.
	OutpostSubnetAssignIpv6AddressOnCreation *bool `field:"optional" json:"outpostSubnetAssignIpv6AddressOnCreation" yaml:"outpostSubnetAssignIpv6AddressOnCreation"`
	// Assigns IPv6 outpost subnet id based on the Amazon provided /56 prefix base 10 integer (0-256).
	//
	// Must be of equal length to the corresponding IPv4 subnet list.
	OutpostSubnetIpv6Prefixes *[]*string `field:"optional" json:"outpostSubnetIpv6Prefixes" yaml:"outpostSubnetIpv6Prefixes"`
	// Explicit values to use in the Name tag on outpost subnets.
	//
	// If empty, Name tags are generated.
	OutpostSubnetNames *[]*string `field:"optional" json:"outpostSubnetNames" yaml:"outpostSubnetNames"`
	// A list of outpost subnets inside the VPC.
	OutpostSubnets *[]*string `field:"optional" json:"outpostSubnets" yaml:"outpostSubnets"`
	// Suffix to append to outpost subnets name outpost.
	OutpostSubnetSuffix *string `field:"optional" json:"outpostSubnetSuffix" yaml:"outpostSubnetSuffix"`
	// Additional tags for the outpost subnets The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	OutpostSubnetTags *map[string]*string `field:"optional" json:"outpostSubnetTags" yaml:"outpostSubnetTags"`
	// Additional tags for the private subnets network ACL The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	PrivateAclTags *map[string]*string `field:"optional" json:"privateAclTags" yaml:"privateAclTags"`
	// Whether to use dedicated network ACL (not default) and custom rules for private subnets.
	PrivateDedicatedNetworkAcl *bool `field:"optional" json:"privateDedicatedNetworkAcl" yaml:"privateDedicatedNetworkAcl"`
	// Private subnets inbound network ACLs [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	PrivateInboundAclRules *[]*map[string]*string `field:"optional" json:"privateInboundAclRules" yaml:"privateInboundAclRules"`
	// Private subnets outbound network ACLs [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	PrivateOutboundAclRules *[]*map[string]*string `field:"optional" json:"privateOutboundAclRules" yaml:"privateOutboundAclRules"`
	// Additional tags for the private route tables The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	PrivateRouteTableTags *map[string]*string `field:"optional" json:"privateRouteTableTags" yaml:"privateRouteTableTags"`
	// Assign IPv6 address on private subnet, must be disabled to change IPv6 CIDRs.
	//
	// This is the IPv6 equivalent of map_public_ip_on_launch.
	PrivateSubnetAssignIpv6AddressOnCreation *bool `field:"optional" json:"privateSubnetAssignIpv6AddressOnCreation" yaml:"privateSubnetAssignIpv6AddressOnCreation"`
	// Assigns IPv6 private subnet id based on the Amazon provided /56 prefix base 10 integer (0-256).
	//
	// Must be of equal length to the corresponding IPv4 subnet list.
	PrivateSubnetIpv6Prefixes *[]*string `field:"optional" json:"privateSubnetIpv6Prefixes" yaml:"privateSubnetIpv6Prefixes"`
	// Explicit values to use in the Name tag on private subnets.
	//
	// If empty, Name tags are generated.
	PrivateSubnetNames *[]*string `field:"optional" json:"privateSubnetNames" yaml:"privateSubnetNames"`
	// A list of private subnets inside the VPC.
	PrivateSubnets *[]*string `field:"optional" json:"privateSubnets" yaml:"privateSubnets"`
	// Suffix to append to private subnets name private.
	PrivateSubnetSuffix *string `field:"optional" json:"privateSubnetSuffix" yaml:"privateSubnetSuffix"`
	// Additional tags for the private subnets The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	PrivateSubnetTags *map[string]*string `field:"optional" json:"privateSubnetTags" yaml:"privateSubnetTags"`
	// Additional tags for the private subnets where the primary key is the AZ The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	PrivateSubnetTagsPerAz *map[string]*map[string]*string `field:"optional" json:"privateSubnetTagsPerAz" yaml:"privateSubnetTagsPerAz"`
	// Should be true if you want route table propagation.
	PropagateIntraRouteTablesVgw *bool `field:"optional" json:"propagateIntraRouteTablesVgw" yaml:"propagateIntraRouteTablesVgw"`
	// Should be true if you want route table propagation.
	PropagatePrivateRouteTablesVgw *bool `field:"optional" json:"propagatePrivateRouteTablesVgw" yaml:"propagatePrivateRouteTablesVgw"`
	// Should be true if you want route table propagation.
	PropagatePublicRouteTablesVgw *bool `field:"optional" json:"propagatePublicRouteTablesVgw" yaml:"propagatePublicRouteTablesVgw"`
	// Additional tags for the public subnets network ACL The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	PublicAclTags *map[string]*string `field:"optional" json:"publicAclTags" yaml:"publicAclTags"`
	// Whether to use dedicated network ACL (not default) and custom rules for public subnets.
	PublicDedicatedNetworkAcl *bool `field:"optional" json:"publicDedicatedNetworkAcl" yaml:"publicDedicatedNetworkAcl"`
	// Public subnets inbound network ACLs [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	PublicInboundAclRules *[]*map[string]*string `field:"optional" json:"publicInboundAclRules" yaml:"publicInboundAclRules"`
	// Public subnets outbound network ACLs [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	PublicOutboundAclRules *[]*map[string]*string `field:"optional" json:"publicOutboundAclRules" yaml:"publicOutboundAclRules"`
	// Additional tags for the public route tables The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	PublicRouteTableTags *map[string]*string `field:"optional" json:"publicRouteTableTags" yaml:"publicRouteTableTags"`
	// Assign IPv6 address on public subnet, must be disabled to change IPv6 CIDRs.
	//
	// This is the IPv6 equivalent of map_public_ip_on_launch.
	PublicSubnetAssignIpv6AddressOnCreation *bool `field:"optional" json:"publicSubnetAssignIpv6AddressOnCreation" yaml:"publicSubnetAssignIpv6AddressOnCreation"`
	// Assigns IPv6 public subnet id based on the Amazon provided /56 prefix base 10 integer (0-256).
	//
	// Must be of equal length to the corresponding IPv4 subnet list.
	PublicSubnetIpv6Prefixes *[]*string `field:"optional" json:"publicSubnetIpv6Prefixes" yaml:"publicSubnetIpv6Prefixes"`
	// Explicit values to use in the Name tag on public subnets.
	//
	// If empty, Name tags are generated.
	PublicSubnetNames *[]*string `field:"optional" json:"publicSubnetNames" yaml:"publicSubnetNames"`
	// A list of public subnets inside the VPC.
	PublicSubnets *[]*string `field:"optional" json:"publicSubnets" yaml:"publicSubnets"`
	// Suffix to append to public subnets name public.
	PublicSubnetSuffix *string `field:"optional" json:"publicSubnetSuffix" yaml:"publicSubnetSuffix"`
	// Additional tags for the public subnets The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	PublicSubnetTags *map[string]*string `field:"optional" json:"publicSubnetTags" yaml:"publicSubnetTags"`
	// Additional tags for the public subnets where the primary key is the AZ The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	PublicSubnetTagsPerAz *map[string]*map[string]*string `field:"optional" json:"publicSubnetTagsPerAz" yaml:"publicSubnetTagsPerAz"`
	// Do you agree that Putin doesn't respect Ukrainian sovereignty and territorial integrity?
	//
	// More info: https://en.wikipedia.org/wiki/Putin_khuylo!
	// true.
	PutinKhuylo *bool `field:"optional" json:"putinKhuylo" yaml:"putinKhuylo"`
	// Additional tags for the redshift subnets network ACL The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	RedshiftAclTags *map[string]*string `field:"optional" json:"redshiftAclTags" yaml:"redshiftAclTags"`
	// Whether to use dedicated network ACL (not default) and custom rules for redshift subnets.
	RedshiftDedicatedNetworkAcl *bool `field:"optional" json:"redshiftDedicatedNetworkAcl" yaml:"redshiftDedicatedNetworkAcl"`
	// Redshift subnets inbound network ACL rules [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	RedshiftInboundAclRules *[]*map[string]*string `field:"optional" json:"redshiftInboundAclRules" yaml:"redshiftInboundAclRules"`
	// Redshift subnets outbound network ACL rules [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	RedshiftOutboundAclRules *[]*map[string]*string `field:"optional" json:"redshiftOutboundAclRules" yaml:"redshiftOutboundAclRules"`
	// Additional tags for the redshift route tables The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	RedshiftRouteTableTags *map[string]*string `field:"optional" json:"redshiftRouteTableTags" yaml:"redshiftRouteTableTags"`
	// Assign IPv6 address on redshift subnet, must be disabled to change IPv6 CIDRs.
	//
	// This is the IPv6 equivalent of map_public_ip_on_launch.
	RedshiftSubnetAssignIpv6AddressOnCreation *bool `field:"optional" json:"redshiftSubnetAssignIpv6AddressOnCreation" yaml:"redshiftSubnetAssignIpv6AddressOnCreation"`
	// Name of redshift subnet group.
	RedshiftSubnetGroupName *string `field:"optional" json:"redshiftSubnetGroupName" yaml:"redshiftSubnetGroupName"`
	// Additional tags for the redshift subnet group The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	RedshiftSubnetGroupTags *map[string]*string `field:"optional" json:"redshiftSubnetGroupTags" yaml:"redshiftSubnetGroupTags"`
	// Assigns IPv6 redshift subnet id based on the Amazon provided /56 prefix base 10 integer (0-256).
	//
	// Must be of equal length to the corresponding IPv4 subnet list.
	RedshiftSubnetIpv6Prefixes *[]*string `field:"optional" json:"redshiftSubnetIpv6Prefixes" yaml:"redshiftSubnetIpv6Prefixes"`
	// Explicit values to use in the Name tag on redshift subnets.
	//
	// If empty, Name tags are generated.
	RedshiftSubnetNames *[]*string `field:"optional" json:"redshiftSubnetNames" yaml:"redshiftSubnetNames"`
	// A list of redshift subnets.
	RedshiftSubnets *[]*string `field:"optional" json:"redshiftSubnets" yaml:"redshiftSubnets"`
	// Suffix to append to redshift subnets name redshift.
	RedshiftSubnetSuffix *string `field:"optional" json:"redshiftSubnetSuffix" yaml:"redshiftSubnetSuffix"`
	// Additional tags for the redshift subnets The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	RedshiftSubnetTags *map[string]*string `field:"optional" json:"redshiftSubnetTags" yaml:"redshiftSubnetTags"`
	// Should be true if you don't want EIPs to be created for your NAT Gateways and will instead pass them in via the 'external_nat_ip_ids' variable.
	ReuseNatIps *bool `field:"optional" json:"reuseNatIps" yaml:"reuseNatIps"`
	// List of secondary CIDR blocks to associate with the VPC to extend the IP Address pool.
	SecondaryCidrBlocks *[]*string `field:"optional" json:"secondaryCidrBlocks" yaml:"secondaryCidrBlocks"`
	// Should be true if you want to provision a single shared NAT Gateway across all of your private networks.
	SingleNatGateway *bool `field:"optional" json:"singleNatGateway" yaml:"singleNatGateway"`
	// A map of tags to add to all resources The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Determines whether IPAM pool is used for CIDR allocation.
	UseIpamPool *bool `field:"optional" json:"useIpamPool" yaml:"useIpamPool"`
	// The ARN of the Permissions Boundary for the VPC Flow Log IAM Role.
	VpcFlowLogPermissionsBoundary *string `field:"optional" json:"vpcFlowLogPermissionsBoundary" yaml:"vpcFlowLogPermissionsBoundary"`
	// Additional tags for the VPC Flow Logs The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	VpcFlowLogTags *map[string]*string `field:"optional" json:"vpcFlowLogTags" yaml:"vpcFlowLogTags"`
	// Additional tags for the VPC The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	VpcTags *map[string]*string `field:"optional" json:"vpcTags" yaml:"vpcTags"`
	// The Availability Zone for the VPN Gateway.
	VpnGatewayAz *string `field:"optional" json:"vpnGatewayAz" yaml:"vpnGatewayAz"`
	// ID of VPN Gateway to attach to the VPC.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
	// Additional tags for the VPN gateway The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	VpnGatewayTags *map[string]*string `field:"optional" json:"vpnGatewayTags" yaml:"vpnGatewayTags"`
}

