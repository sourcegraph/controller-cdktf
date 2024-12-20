// @cdktf/provider-awsvpc
package awsvpc

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-awsvpc.Awsvpc",
		reflect.TypeOf((*Awsvpc)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addProvider", GoMethod: "AddProvider"},
			_jsii_.MemberProperty{JsiiProperty: "amazonSideAsn", GoGetter: "AmazonSideAsn"},
			_jsii_.MemberProperty{JsiiProperty: "assignIpv6AddressOnCreation", GoGetter: "AssignIpv6AddressOnCreation"},
			_jsii_.MemberProperty{JsiiProperty: "azs", GoGetter: "Azs"},
			_jsii_.MemberProperty{JsiiProperty: "azsOutput", GoGetter: "AzsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cgwArnsOutput", GoGetter: "CgwArnsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "cgwIdsOutput", GoGetter: "CgwIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "cidr", GoGetter: "Cidr"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "createDatabaseInternetGatewayRoute", GoGetter: "CreateDatabaseInternetGatewayRoute"},
			_jsii_.MemberProperty{JsiiProperty: "createDatabaseNatGatewayRoute", GoGetter: "CreateDatabaseNatGatewayRoute"},
			_jsii_.MemberProperty{JsiiProperty: "createDatabaseSubnetGroup", GoGetter: "CreateDatabaseSubnetGroup"},
			_jsii_.MemberProperty{JsiiProperty: "createDatabaseSubnetRouteTable", GoGetter: "CreateDatabaseSubnetRouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "createEgressOnlyIgw", GoGetter: "CreateEgressOnlyIgw"},
			_jsii_.MemberProperty{JsiiProperty: "createElasticacheSubnetGroup", GoGetter: "CreateElasticacheSubnetGroup"},
			_jsii_.MemberProperty{JsiiProperty: "createElasticacheSubnetRouteTable", GoGetter: "CreateElasticacheSubnetRouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "createFlowLogCloudwatchIamRole", GoGetter: "CreateFlowLogCloudwatchIamRole"},
			_jsii_.MemberProperty{JsiiProperty: "createFlowLogCloudwatchLogGroup", GoGetter: "CreateFlowLogCloudwatchLogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "createIgw", GoGetter: "CreateIgw"},
			_jsii_.MemberProperty{JsiiProperty: "createRedshiftSubnetGroup", GoGetter: "CreateRedshiftSubnetGroup"},
			_jsii_.MemberProperty{JsiiProperty: "createRedshiftSubnetRouteTable", GoGetter: "CreateRedshiftSubnetRouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "createVpc", GoGetter: "CreateVpc"},
			_jsii_.MemberProperty{JsiiProperty: "customerGateways", GoGetter: "CustomerGateways"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayTags", GoGetter: "CustomerGatewayTags"},
			_jsii_.MemberProperty{JsiiProperty: "databaseAclTags", GoGetter: "DatabaseAclTags"},
			_jsii_.MemberProperty{JsiiProperty: "databaseDedicatedNetworkAcl", GoGetter: "DatabaseDedicatedNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInboundAclRules", GoGetter: "DatabaseInboundAclRules"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInternetGatewayRouteIdOutput", GoGetter: "DatabaseInternetGatewayRouteIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseIpv6EgressRouteIdOutput", GoGetter: "DatabaseIpv6EgressRouteIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseNatGatewayRouteIdsOutput", GoGetter: "DatabaseNatGatewayRouteIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseNetworkAclArnOutput", GoGetter: "DatabaseNetworkAclArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseNetworkAclIdOutput", GoGetter: "DatabaseNetworkAclIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseOutboundAclRules", GoGetter: "DatabaseOutboundAclRules"},
			_jsii_.MemberProperty{JsiiProperty: "databaseRouteTableAssociationIdsOutput", GoGetter: "DatabaseRouteTableAssociationIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseRouteTableIdsOutput", GoGetter: "DatabaseRouteTableIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseRouteTableTags", GoGetter: "DatabaseRouteTableTags"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSubnetArnsOutput", GoGetter: "DatabaseSubnetArnsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSubnetAssignIpv6AddressOnCreation", GoGetter: "DatabaseSubnetAssignIpv6AddressOnCreation"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSubnetGroupName", GoGetter: "DatabaseSubnetGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSubnetGroupNameOutput", GoGetter: "DatabaseSubnetGroupNameOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSubnetGroupOutput", GoGetter: "DatabaseSubnetGroupOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSubnetGroupTags", GoGetter: "DatabaseSubnetGroupTags"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSubnetIpv6Prefixes", GoGetter: "DatabaseSubnetIpv6Prefixes"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSubnetNames", GoGetter: "DatabaseSubnetNames"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSubnets", GoGetter: "DatabaseSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSubnetsCidrBlocksOutput", GoGetter: "DatabaseSubnetsCidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSubnetsIpv6CidrBlocksOutput", GoGetter: "DatabaseSubnetsIpv6CidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSubnetsOutput", GoGetter: "DatabaseSubnetsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSubnetSuffix", GoGetter: "DatabaseSubnetSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSubnetTags", GoGetter: "DatabaseSubnetTags"},
			_jsii_.MemberProperty{JsiiProperty: "defaultNetworkAclEgress", GoGetter: "DefaultNetworkAclEgress"},
			_jsii_.MemberProperty{JsiiProperty: "defaultNetworkAclIdOutput", GoGetter: "DefaultNetworkAclIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultNetworkAclIngress", GoGetter: "DefaultNetworkAclIngress"},
			_jsii_.MemberProperty{JsiiProperty: "defaultNetworkAclName", GoGetter: "DefaultNetworkAclName"},
			_jsii_.MemberProperty{JsiiProperty: "defaultNetworkAclTags", GoGetter: "DefaultNetworkAclTags"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteTableIdOutput", GoGetter: "DefaultRouteTableIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteTableName", GoGetter: "DefaultRouteTableName"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteTablePropagatingVgws", GoGetter: "DefaultRouteTablePropagatingVgws"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteTableRoutes", GoGetter: "DefaultRouteTableRoutes"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteTableTags", GoGetter: "DefaultRouteTableTags"},
			_jsii_.MemberProperty{JsiiProperty: "defaultSecurityGroupEgress", GoGetter: "DefaultSecurityGroupEgress"},
			_jsii_.MemberProperty{JsiiProperty: "defaultSecurityGroupIdOutput", GoGetter: "DefaultSecurityGroupIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultSecurityGroupIngress", GoGetter: "DefaultSecurityGroupIngress"},
			_jsii_.MemberProperty{JsiiProperty: "defaultSecurityGroupName", GoGetter: "DefaultSecurityGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "defaultSecurityGroupTags", GoGetter: "DefaultSecurityGroupTags"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcArnOutput", GoGetter: "DefaultVpcArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcCidrBlockOutput", GoGetter: "DefaultVpcCidrBlockOutput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcDefaultNetworkAclIdOutput", GoGetter: "DefaultVpcDefaultNetworkAclIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcDefaultRouteTableIdOutput", GoGetter: "DefaultVpcDefaultRouteTableIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcDefaultSecurityGroupIdOutput", GoGetter: "DefaultVpcDefaultSecurityGroupIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcEnableClassiclink", GoGetter: "DefaultVpcEnableClassiclink"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcEnableDnsHostnames", GoGetter: "DefaultVpcEnableDnsHostnames"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcEnableDnsHostnamesOutput", GoGetter: "DefaultVpcEnableDnsHostnamesOutput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcEnableDnsSupport", GoGetter: "DefaultVpcEnableDnsSupport"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcEnableDnsSupportOutput", GoGetter: "DefaultVpcEnableDnsSupportOutput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcIdOutput", GoGetter: "DefaultVpcIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcInstanceTenancyOutput", GoGetter: "DefaultVpcInstanceTenancyOutput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcMainRouteTableIdOutput", GoGetter: "DefaultVpcMainRouteTableIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcName", GoGetter: "DefaultVpcName"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVpcTags", GoGetter: "DefaultVpcTags"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "dhcpOptionsDomainName", GoGetter: "DhcpOptionsDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "dhcpOptionsDomainNameServers", GoGetter: "DhcpOptionsDomainNameServers"},
			_jsii_.MemberProperty{JsiiProperty: "dhcpOptionsIdOutput", GoGetter: "DhcpOptionsIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dhcpOptionsNetbiosNameServers", GoGetter: "DhcpOptionsNetbiosNameServers"},
			_jsii_.MemberProperty{JsiiProperty: "dhcpOptionsNetbiosNodeType", GoGetter: "DhcpOptionsNetbiosNodeType"},
			_jsii_.MemberProperty{JsiiProperty: "dhcpOptionsNtpServers", GoGetter: "DhcpOptionsNtpServers"},
			_jsii_.MemberProperty{JsiiProperty: "dhcpOptionsTags", GoGetter: "DhcpOptionsTags"},
			_jsii_.MemberProperty{JsiiProperty: "egressOnlyInternetGatewayIdOutput", GoGetter: "EgressOnlyInternetGatewayIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheAclTags", GoGetter: "ElasticacheAclTags"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheDedicatedNetworkAcl", GoGetter: "ElasticacheDedicatedNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheInboundAclRules", GoGetter: "ElasticacheInboundAclRules"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheNetworkAclArnOutput", GoGetter: "ElasticacheNetworkAclArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheNetworkAclIdOutput", GoGetter: "ElasticacheNetworkAclIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheOutboundAclRules", GoGetter: "ElasticacheOutboundAclRules"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheRouteTableAssociationIdsOutput", GoGetter: "ElasticacheRouteTableAssociationIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheRouteTableIdsOutput", GoGetter: "ElasticacheRouteTableIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheRouteTableTags", GoGetter: "ElasticacheRouteTableTags"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheSubnetArnsOutput", GoGetter: "ElasticacheSubnetArnsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheSubnetAssignIpv6AddressOnCreation", GoGetter: "ElasticacheSubnetAssignIpv6AddressOnCreation"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheSubnetGroupName", GoGetter: "ElasticacheSubnetGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheSubnetGroupNameOutput", GoGetter: "ElasticacheSubnetGroupNameOutput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheSubnetGroupOutput", GoGetter: "ElasticacheSubnetGroupOutput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheSubnetGroupTags", GoGetter: "ElasticacheSubnetGroupTags"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheSubnetIpv6Prefixes", GoGetter: "ElasticacheSubnetIpv6Prefixes"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheSubnetNames", GoGetter: "ElasticacheSubnetNames"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheSubnets", GoGetter: "ElasticacheSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheSubnetsCidrBlocksOutput", GoGetter: "ElasticacheSubnetsCidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheSubnetsIpv6CidrBlocksOutput", GoGetter: "ElasticacheSubnetsIpv6CidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheSubnetsOutput", GoGetter: "ElasticacheSubnetsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheSubnetSuffix", GoGetter: "ElasticacheSubnetSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "elasticacheSubnetTags", GoGetter: "ElasticacheSubnetTags"},
			_jsii_.MemberProperty{JsiiProperty: "enableClassiclink", GoGetter: "EnableClassiclink"},
			_jsii_.MemberProperty{JsiiProperty: "enableClassiclinkDnsSupport", GoGetter: "EnableClassiclinkDnsSupport"},
			_jsii_.MemberProperty{JsiiProperty: "enableDhcpOptions", GoGetter: "EnableDhcpOptions"},
			_jsii_.MemberProperty{JsiiProperty: "enableDnsHostnames", GoGetter: "EnableDnsHostnames"},
			_jsii_.MemberProperty{JsiiProperty: "enableDnsSupport", GoGetter: "EnableDnsSupport"},
			_jsii_.MemberProperty{JsiiProperty: "enableFlowLog", GoGetter: "EnableFlowLog"},
			_jsii_.MemberProperty{JsiiProperty: "enableIpv6", GoGetter: "EnableIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "enableNatGateway", GoGetter: "EnableNatGateway"},
			_jsii_.MemberProperty{JsiiProperty: "enablePublicRedshift", GoGetter: "EnablePublicRedshift"},
			_jsii_.MemberProperty{JsiiProperty: "enableVpnGateway", GoGetter: "EnableVpnGateway"},
			_jsii_.MemberProperty{JsiiProperty: "externalNatIpIds", GoGetter: "ExternalNatIpIds"},
			_jsii_.MemberProperty{JsiiProperty: "externalNatIps", GoGetter: "ExternalNatIps"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogCloudwatchIamRoleArn", GoGetter: "FlowLogCloudwatchIamRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogCloudwatchLogGroupKmsKeyId", GoGetter: "FlowLogCloudwatchLogGroupKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogCloudwatchLogGroupNamePrefix", GoGetter: "FlowLogCloudwatchLogGroupNamePrefix"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogCloudwatchLogGroupNameSuffix", GoGetter: "FlowLogCloudwatchLogGroupNameSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogCloudwatchLogGroupRetentionInDays", GoGetter: "FlowLogCloudwatchLogGroupRetentionInDays"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogDestinationArn", GoGetter: "FlowLogDestinationArn"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogDestinationType", GoGetter: "FlowLogDestinationType"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogFileFormat", GoGetter: "FlowLogFileFormat"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogHiveCompatiblePartitions", GoGetter: "FlowLogHiveCompatiblePartitions"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogLogFormat", GoGetter: "FlowLogLogFormat"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogMaxAggregationInterval", GoGetter: "FlowLogMaxAggregationInterval"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogPerHourPartition", GoGetter: "FlowLogPerHourPartition"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogTrafficType", GoGetter: "FlowLogTrafficType"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberProperty{JsiiProperty: "igwArnOutput", GoGetter: "IgwArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "igwIdOutput", GoGetter: "IgwIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "igwTags", GoGetter: "IgwTags"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTenancy", GoGetter: "InstanceTenancy"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForOutput", GoMethod: "InterpolationForOutput"},
			_jsii_.MemberProperty{JsiiProperty: "intraAclTags", GoGetter: "IntraAclTags"},
			_jsii_.MemberProperty{JsiiProperty: "intraDedicatedNetworkAcl", GoGetter: "IntraDedicatedNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "intraInboundAclRules", GoGetter: "IntraInboundAclRules"},
			_jsii_.MemberProperty{JsiiProperty: "intraNetworkAclArnOutput", GoGetter: "IntraNetworkAclArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "intraNetworkAclIdOutput", GoGetter: "IntraNetworkAclIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "intraOutboundAclRules", GoGetter: "IntraOutboundAclRules"},
			_jsii_.MemberProperty{JsiiProperty: "intraRouteTableAssociationIdsOutput", GoGetter: "IntraRouteTableAssociationIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "intraRouteTableIdsOutput", GoGetter: "IntraRouteTableIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "intraRouteTableTags", GoGetter: "IntraRouteTableTags"},
			_jsii_.MemberProperty{JsiiProperty: "intraSubnetArnsOutput", GoGetter: "IntraSubnetArnsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "intraSubnetAssignIpv6AddressOnCreation", GoGetter: "IntraSubnetAssignIpv6AddressOnCreation"},
			_jsii_.MemberProperty{JsiiProperty: "intraSubnetIpv6Prefixes", GoGetter: "IntraSubnetIpv6Prefixes"},
			_jsii_.MemberProperty{JsiiProperty: "intraSubnetNames", GoGetter: "IntraSubnetNames"},
			_jsii_.MemberProperty{JsiiProperty: "intraSubnets", GoGetter: "IntraSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "intraSubnetsCidrBlocksOutput", GoGetter: "IntraSubnetsCidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "intraSubnetsIpv6CidrBlocksOutput", GoGetter: "IntraSubnetsIpv6CidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "intraSubnetsOutput", GoGetter: "IntraSubnetsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "intraSubnetSuffix", GoGetter: "IntraSubnetSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "intraSubnetTags", GoGetter: "IntraSubnetTags"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4IpamPoolId", GoGetter: "Ipv4IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4NetmaskLength", GoGetter: "Ipv4NetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Cidr", GoGetter: "Ipv6Cidr"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6IpamPoolId", GoGetter: "Ipv6IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6NetmaskLength", GoGetter: "Ipv6NetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "manageDefaultNetworkAcl", GoGetter: "ManageDefaultNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "manageDefaultRouteTable", GoGetter: "ManageDefaultRouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "manageDefaultSecurityGroup", GoGetter: "ManageDefaultSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "manageDefaultVpc", GoGetter: "ManageDefaultVpc"},
			_jsii_.MemberProperty{JsiiProperty: "mapPublicIpOnLaunch", GoGetter: "MapPublicIpOnLaunch"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameOutput", GoGetter: "NameOutput"},
			_jsii_.MemberProperty{JsiiProperty: "natEipTags", GoGetter: "NatEipTags"},
			_jsii_.MemberProperty{JsiiProperty: "natGatewayDestinationCidrBlock", GoGetter: "NatGatewayDestinationCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "natGatewayTags", GoGetter: "NatGatewayTags"},
			_jsii_.MemberProperty{JsiiProperty: "natgwIdsOutput", GoGetter: "NatgwIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "natIdsOutput", GoGetter: "NatIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "natPublicIpsOutput", GoGetter: "NatPublicIpsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "oneNatGatewayPerAz", GoGetter: "OneNatGatewayPerAz"},
			_jsii_.MemberProperty{JsiiProperty: "outpostAclTags", GoGetter: "OutpostAclTags"},
			_jsii_.MemberProperty{JsiiProperty: "outpostArn", GoGetter: "OutpostArn"},
			_jsii_.MemberProperty{JsiiProperty: "outpostAz", GoGetter: "OutpostAz"},
			_jsii_.MemberProperty{JsiiProperty: "outpostDedicatedNetworkAcl", GoGetter: "OutpostDedicatedNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "outpostInboundAclRules", GoGetter: "OutpostInboundAclRules"},
			_jsii_.MemberProperty{JsiiProperty: "outpostNetworkAclArnOutput", GoGetter: "OutpostNetworkAclArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "outpostNetworkAclIdOutput", GoGetter: "OutpostNetworkAclIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "outpostOutboundAclRules", GoGetter: "OutpostOutboundAclRules"},
			_jsii_.MemberProperty{JsiiProperty: "outpostSubnetArnsOutput", GoGetter: "OutpostSubnetArnsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "outpostSubnetAssignIpv6AddressOnCreation", GoGetter: "OutpostSubnetAssignIpv6AddressOnCreation"},
			_jsii_.MemberProperty{JsiiProperty: "outpostSubnetIpv6Prefixes", GoGetter: "OutpostSubnetIpv6Prefixes"},
			_jsii_.MemberProperty{JsiiProperty: "outpostSubnetNames", GoGetter: "OutpostSubnetNames"},
			_jsii_.MemberProperty{JsiiProperty: "outpostSubnets", GoGetter: "OutpostSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "outpostSubnetsCidrBlocksOutput", GoGetter: "OutpostSubnetsCidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "outpostSubnetsIpv6CidrBlocksOutput", GoGetter: "OutpostSubnetsIpv6CidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "outpostSubnetsOutput", GoGetter: "OutpostSubnetsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "outpostSubnetSuffix", GoGetter: "OutpostSubnetSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "outpostSubnetTags", GoGetter: "OutpostSubnetTags"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "privateAclTags", GoGetter: "PrivateAclTags"},
			_jsii_.MemberProperty{JsiiProperty: "privateDedicatedNetworkAcl", GoGetter: "PrivateDedicatedNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "privateInboundAclRules", GoGetter: "PrivateInboundAclRules"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpv6EgressRouteIdsOutput", GoGetter: "PrivateIpv6EgressRouteIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateNatGatewayRouteIdsOutput", GoGetter: "PrivateNatGatewayRouteIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateNetworkAclArnOutput", GoGetter: "PrivateNetworkAclArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateNetworkAclIdOutput", GoGetter: "PrivateNetworkAclIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateOutboundAclRules", GoGetter: "PrivateOutboundAclRules"},
			_jsii_.MemberProperty{JsiiProperty: "privateRouteTableAssociationIdsOutput", GoGetter: "PrivateRouteTableAssociationIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateRouteTableIdsOutput", GoGetter: "PrivateRouteTableIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateRouteTableTags", GoGetter: "PrivateRouteTableTags"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetArnsOutput", GoGetter: "PrivateSubnetArnsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetAssignIpv6AddressOnCreation", GoGetter: "PrivateSubnetAssignIpv6AddressOnCreation"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetIpv6Prefixes", GoGetter: "PrivateSubnetIpv6Prefixes"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetNames", GoGetter: "PrivateSubnetNames"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnets", GoGetter: "PrivateSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetsCidrBlocksOutput", GoGetter: "PrivateSubnetsCidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetsIpv6CidrBlocksOutput", GoGetter: "PrivateSubnetsIpv6CidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetsOutput", GoGetter: "PrivateSubnetsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetSuffix", GoGetter: "PrivateSubnetSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetTags", GoGetter: "PrivateSubnetTags"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetTagsPerAz", GoGetter: "PrivateSubnetTagsPerAz"},
			_jsii_.MemberProperty{JsiiProperty: "propagateIntraRouteTablesVgw", GoGetter: "PropagateIntraRouteTablesVgw"},
			_jsii_.MemberProperty{JsiiProperty: "propagatePrivateRouteTablesVgw", GoGetter: "PropagatePrivateRouteTablesVgw"},
			_jsii_.MemberProperty{JsiiProperty: "propagatePublicRouteTablesVgw", GoGetter: "PropagatePublicRouteTablesVgw"},
			_jsii_.MemberProperty{JsiiProperty: "providers", GoGetter: "Providers"},
			_jsii_.MemberProperty{JsiiProperty: "publicAclTags", GoGetter: "PublicAclTags"},
			_jsii_.MemberProperty{JsiiProperty: "publicDedicatedNetworkAcl", GoGetter: "PublicDedicatedNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "publicInboundAclRules", GoGetter: "PublicInboundAclRules"},
			_jsii_.MemberProperty{JsiiProperty: "publicInternetGatewayIpv6RouteIdOutput", GoGetter: "PublicInternetGatewayIpv6RouteIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicInternetGatewayRouteIdOutput", GoGetter: "PublicInternetGatewayRouteIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicNetworkAclArnOutput", GoGetter: "PublicNetworkAclArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicNetworkAclIdOutput", GoGetter: "PublicNetworkAclIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicOutboundAclRules", GoGetter: "PublicOutboundAclRules"},
			_jsii_.MemberProperty{JsiiProperty: "publicRouteTableAssociationIdsOutput", GoGetter: "PublicRouteTableAssociationIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicRouteTableIdsOutput", GoGetter: "PublicRouteTableIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicRouteTableTags", GoGetter: "PublicRouteTableTags"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetArnsOutput", GoGetter: "PublicSubnetArnsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetAssignIpv6AddressOnCreation", GoGetter: "PublicSubnetAssignIpv6AddressOnCreation"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetIpv6Prefixes", GoGetter: "PublicSubnetIpv6Prefixes"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetNames", GoGetter: "PublicSubnetNames"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnets", GoGetter: "PublicSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetsCidrBlocksOutput", GoGetter: "PublicSubnetsCidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetsIpv6CidrBlocksOutput", GoGetter: "PublicSubnetsIpv6CidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetsOutput", GoGetter: "PublicSubnetsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetSuffix", GoGetter: "PublicSubnetSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetTags", GoGetter: "PublicSubnetTags"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetTagsPerAz", GoGetter: "PublicSubnetTagsPerAz"},
			_jsii_.MemberProperty{JsiiProperty: "putinKhuylo", GoGetter: "PutinKhuylo"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftAclTags", GoGetter: "RedshiftAclTags"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftDedicatedNetworkAcl", GoGetter: "RedshiftDedicatedNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftInboundAclRules", GoGetter: "RedshiftInboundAclRules"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftNetworkAclArnOutput", GoGetter: "RedshiftNetworkAclArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftNetworkAclIdOutput", GoGetter: "RedshiftNetworkAclIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftOutboundAclRules", GoGetter: "RedshiftOutboundAclRules"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftPublicRouteTableAssociationIdsOutput", GoGetter: "RedshiftPublicRouteTableAssociationIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftRouteTableAssociationIdsOutput", GoGetter: "RedshiftRouteTableAssociationIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftRouteTableIdsOutput", GoGetter: "RedshiftRouteTableIdsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftRouteTableTags", GoGetter: "RedshiftRouteTableTags"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftSubnetArnsOutput", GoGetter: "RedshiftSubnetArnsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftSubnetAssignIpv6AddressOnCreation", GoGetter: "RedshiftSubnetAssignIpv6AddressOnCreation"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftSubnetGroupName", GoGetter: "RedshiftSubnetGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftSubnetGroupOutput", GoGetter: "RedshiftSubnetGroupOutput"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftSubnetGroupTags", GoGetter: "RedshiftSubnetGroupTags"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftSubnetIpv6Prefixes", GoGetter: "RedshiftSubnetIpv6Prefixes"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftSubnetNames", GoGetter: "RedshiftSubnetNames"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftSubnets", GoGetter: "RedshiftSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftSubnetsCidrBlocksOutput", GoGetter: "RedshiftSubnetsCidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftSubnetsIpv6CidrBlocksOutput", GoGetter: "RedshiftSubnetsIpv6CidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftSubnetsOutput", GoGetter: "RedshiftSubnetsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftSubnetSuffix", GoGetter: "RedshiftSubnetSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftSubnetTags", GoGetter: "RedshiftSubnetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "reuseNatIps", GoGetter: "ReuseNatIps"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryCidrBlocks", GoGetter: "SecondaryCidrBlocks"},
			_jsii_.MemberProperty{JsiiProperty: "singleNatGateway", GoGetter: "SingleNatGateway"},
			_jsii_.MemberProperty{JsiiProperty: "skipAssetCreationFromLocalModules", GoGetter: "SkipAssetCreationFromLocalModules"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "thisCustomerGatewayOutput", GoGetter: "ThisCustomerGatewayOutput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "useIpamPool", GoGetter: "UseIpamPool"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "vgwArnOutput", GoGetter: "VgwArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vgwIdOutput", GoGetter: "VgwIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcArnOutput", GoGetter: "VpcArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlockOutput", GoGetter: "VpcCidrBlockOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEnableDnsHostnamesOutput", GoGetter: "VpcEnableDnsHostnamesOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEnableDnsSupportOutput", GoGetter: "VpcEnableDnsSupportOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcFlowLogCloudwatchIamRoleArnOutput", GoGetter: "VpcFlowLogCloudwatchIamRoleArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcFlowLogDestinationArnOutput", GoGetter: "VpcFlowLogDestinationArnOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcFlowLogDestinationTypeOutput", GoGetter: "VpcFlowLogDestinationTypeOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcFlowLogIdOutput", GoGetter: "VpcFlowLogIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcFlowLogPermissionsBoundary", GoGetter: "VpcFlowLogPermissionsBoundary"},
			_jsii_.MemberProperty{JsiiProperty: "vpcFlowLogTags", GoGetter: "VpcFlowLogTags"},
			_jsii_.MemberProperty{JsiiProperty: "vpcIdOutput", GoGetter: "VpcIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcInstanceTenancyOutput", GoGetter: "VpcInstanceTenancyOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcIpv6AssociationIdOutput", GoGetter: "VpcIpv6AssociationIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcIpv6CidrBlockOutput", GoGetter: "VpcIpv6CidrBlockOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcMainRouteTableIdOutput", GoGetter: "VpcMainRouteTableIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOwnerIdOutput", GoGetter: "VpcOwnerIdOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSecondaryCidrBlocksOutput", GoGetter: "VpcSecondaryCidrBlocksOutput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcTags", GoGetter: "VpcTags"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayAz", GoGetter: "VpnGatewayAz"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayId", GoGetter: "VpnGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayTags", GoGetter: "VpnGatewayTags"},
		},
		func() interface{} {
			j := jsiiProxy_Awsvpc{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformModule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-awsvpc.AwsvpcConfig",
		reflect.TypeOf((*AwsvpcConfig)(nil)).Elem(),
	)
}
