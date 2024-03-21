package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/awsvpc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/awsvpc/internal"
)

// Defines an Awsvpc based on a Terraform module.
//
// Docs at Terraform Registry: {@link https://registry.terraform.io/modules/terraform-aws-modules/vpc/aws/3.19.0 terraform-aws-modules/vpc/aws}
type Awsvpc interface {
	cdktf.TerraformModule
	AmazonSideAsn() *string
	SetAmazonSideAsn(val *string)
	AssignIpv6AddressOnCreation() *bool
	SetAssignIpv6AddressOnCreation(val *bool)
	Azs() *[]*string
	SetAzs(val *[]*string)
	AzsOutput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CgwArnsOutput() *string
	CgwIdsOutput() *string
	Cidr() *string
	SetCidr(val *string)
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CreateDatabaseInternetGatewayRoute() *bool
	SetCreateDatabaseInternetGatewayRoute(val *bool)
	CreateDatabaseNatGatewayRoute() *bool
	SetCreateDatabaseNatGatewayRoute(val *bool)
	CreateDatabaseSubnetGroup() *bool
	SetCreateDatabaseSubnetGroup(val *bool)
	CreateDatabaseSubnetRouteTable() *bool
	SetCreateDatabaseSubnetRouteTable(val *bool)
	CreateEgressOnlyIgw() *bool
	SetCreateEgressOnlyIgw(val *bool)
	CreateElasticacheSubnetGroup() *bool
	SetCreateElasticacheSubnetGroup(val *bool)
	CreateElasticacheSubnetRouteTable() *bool
	SetCreateElasticacheSubnetRouteTable(val *bool)
	CreateFlowLogCloudwatchIamRole() *bool
	SetCreateFlowLogCloudwatchIamRole(val *bool)
	CreateFlowLogCloudwatchLogGroup() *bool
	SetCreateFlowLogCloudwatchLogGroup(val *bool)
	CreateIgw() *bool
	SetCreateIgw(val *bool)
	CreateRedshiftSubnetGroup() *bool
	SetCreateRedshiftSubnetGroup(val *bool)
	CreateRedshiftSubnetRouteTable() *bool
	SetCreateRedshiftSubnetRouteTable(val *bool)
	CreateVpc() *bool
	SetCreateVpc(val *bool)
	CustomerGateways() *map[string]*map[string]interface{}
	SetCustomerGateways(val *map[string]*map[string]interface{})
	CustomerGatewayTags() *map[string]*string
	SetCustomerGatewayTags(val *map[string]*string)
	DatabaseAclTags() *map[string]*string
	SetDatabaseAclTags(val *map[string]*string)
	DatabaseDedicatedNetworkAcl() *bool
	SetDatabaseDedicatedNetworkAcl(val *bool)
	DatabaseInboundAclRules() *[]*map[string]*string
	SetDatabaseInboundAclRules(val *[]*map[string]*string)
	DatabaseInternetGatewayRouteIdOutput() *string
	DatabaseIpv6EgressRouteIdOutput() *string
	DatabaseNatGatewayRouteIdsOutput() *string
	DatabaseNetworkAclArnOutput() *string
	DatabaseNetworkAclIdOutput() *string
	DatabaseOutboundAclRules() *[]*map[string]*string
	SetDatabaseOutboundAclRules(val *[]*map[string]*string)
	DatabaseRouteTableAssociationIdsOutput() *string
	DatabaseRouteTableIdsOutput() *string
	DatabaseRouteTableTags() *map[string]*string
	SetDatabaseRouteTableTags(val *map[string]*string)
	DatabaseSubnetArnsOutput() *string
	DatabaseSubnetAssignIpv6AddressOnCreation() *bool
	SetDatabaseSubnetAssignIpv6AddressOnCreation(val *bool)
	DatabaseSubnetGroupName() *string
	SetDatabaseSubnetGroupName(val *string)
	DatabaseSubnetGroupNameOutput() *string
	DatabaseSubnetGroupOutput() *string
	DatabaseSubnetGroupTags() *map[string]*string
	SetDatabaseSubnetGroupTags(val *map[string]*string)
	DatabaseSubnetIpv6Prefixes() *[]*string
	SetDatabaseSubnetIpv6Prefixes(val *[]*string)
	DatabaseSubnetNames() *[]*string
	SetDatabaseSubnetNames(val *[]*string)
	DatabaseSubnets() *[]*string
	SetDatabaseSubnets(val *[]*string)
	DatabaseSubnetsCidrBlocksOutput() *string
	DatabaseSubnetsIpv6CidrBlocksOutput() *string
	DatabaseSubnetsOutput() *string
	DatabaseSubnetSuffix() *string
	SetDatabaseSubnetSuffix(val *string)
	DatabaseSubnetTags() *map[string]*string
	SetDatabaseSubnetTags(val *map[string]*string)
	DefaultNetworkAclEgress() *[]*map[string]*string
	SetDefaultNetworkAclEgress(val *[]*map[string]*string)
	DefaultNetworkAclIdOutput() *string
	DefaultNetworkAclIngress() *[]*map[string]*string
	SetDefaultNetworkAclIngress(val *[]*map[string]*string)
	DefaultNetworkAclName() *string
	SetDefaultNetworkAclName(val *string)
	DefaultNetworkAclTags() *map[string]*string
	SetDefaultNetworkAclTags(val *map[string]*string)
	DefaultRouteTableIdOutput() *string
	DefaultRouteTableName() *string
	SetDefaultRouteTableName(val *string)
	DefaultRouteTablePropagatingVgws() *[]*string
	SetDefaultRouteTablePropagatingVgws(val *[]*string)
	DefaultRouteTableRoutes() *[]*map[string]*string
	SetDefaultRouteTableRoutes(val *[]*map[string]*string)
	DefaultRouteTableTags() *map[string]*string
	SetDefaultRouteTableTags(val *map[string]*string)
	DefaultSecurityGroupEgress() *[]*map[string]*string
	SetDefaultSecurityGroupEgress(val *[]*map[string]*string)
	DefaultSecurityGroupIdOutput() *string
	DefaultSecurityGroupIngress() *[]*map[string]*string
	SetDefaultSecurityGroupIngress(val *[]*map[string]*string)
	DefaultSecurityGroupName() *string
	SetDefaultSecurityGroupName(val *string)
	DefaultSecurityGroupTags() *map[string]*string
	SetDefaultSecurityGroupTags(val *map[string]*string)
	DefaultVpcArnOutput() *string
	DefaultVpcCidrBlockOutput() *string
	DefaultVpcDefaultNetworkAclIdOutput() *string
	DefaultVpcDefaultRouteTableIdOutput() *string
	DefaultVpcDefaultSecurityGroupIdOutput() *string
	DefaultVpcEnableClassiclink() *bool
	SetDefaultVpcEnableClassiclink(val *bool)
	DefaultVpcEnableDnsHostnames() *bool
	SetDefaultVpcEnableDnsHostnames(val *bool)
	DefaultVpcEnableDnsHostnamesOutput() *string
	DefaultVpcEnableDnsSupport() *bool
	SetDefaultVpcEnableDnsSupport(val *bool)
	DefaultVpcEnableDnsSupportOutput() *string
	DefaultVpcIdOutput() *string
	DefaultVpcInstanceTenancyOutput() *string
	DefaultVpcMainRouteTableIdOutput() *string
	DefaultVpcName() *string
	SetDefaultVpcName(val *string)
	DefaultVpcTags() *map[string]*string
	SetDefaultVpcTags(val *map[string]*string)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DhcpOptionsDomainName() *string
	SetDhcpOptionsDomainName(val *string)
	DhcpOptionsDomainNameServers() *[]*string
	SetDhcpOptionsDomainNameServers(val *[]*string)
	DhcpOptionsIdOutput() *string
	DhcpOptionsNetbiosNameServers() *[]*string
	SetDhcpOptionsNetbiosNameServers(val *[]*string)
	DhcpOptionsNetbiosNodeType() *string
	SetDhcpOptionsNetbiosNodeType(val *string)
	DhcpOptionsNtpServers() *[]*string
	SetDhcpOptionsNtpServers(val *[]*string)
	DhcpOptionsTags() *map[string]*string
	SetDhcpOptionsTags(val *map[string]*string)
	EgressOnlyInternetGatewayIdOutput() *string
	ElasticacheAclTags() *map[string]*string
	SetElasticacheAclTags(val *map[string]*string)
	ElasticacheDedicatedNetworkAcl() *bool
	SetElasticacheDedicatedNetworkAcl(val *bool)
	ElasticacheInboundAclRules() *[]*map[string]*string
	SetElasticacheInboundAclRules(val *[]*map[string]*string)
	ElasticacheNetworkAclArnOutput() *string
	ElasticacheNetworkAclIdOutput() *string
	ElasticacheOutboundAclRules() *[]*map[string]*string
	SetElasticacheOutboundAclRules(val *[]*map[string]*string)
	ElasticacheRouteTableAssociationIdsOutput() *string
	ElasticacheRouteTableIdsOutput() *string
	ElasticacheRouteTableTags() *map[string]*string
	SetElasticacheRouteTableTags(val *map[string]*string)
	ElasticacheSubnetArnsOutput() *string
	ElasticacheSubnetAssignIpv6AddressOnCreation() *bool
	SetElasticacheSubnetAssignIpv6AddressOnCreation(val *bool)
	ElasticacheSubnetGroupName() *string
	SetElasticacheSubnetGroupName(val *string)
	ElasticacheSubnetGroupNameOutput() *string
	ElasticacheSubnetGroupOutput() *string
	ElasticacheSubnetGroupTags() *map[string]*string
	SetElasticacheSubnetGroupTags(val *map[string]*string)
	ElasticacheSubnetIpv6Prefixes() *[]*string
	SetElasticacheSubnetIpv6Prefixes(val *[]*string)
	ElasticacheSubnetNames() *[]*string
	SetElasticacheSubnetNames(val *[]*string)
	ElasticacheSubnets() *[]*string
	SetElasticacheSubnets(val *[]*string)
	ElasticacheSubnetsCidrBlocksOutput() *string
	ElasticacheSubnetsIpv6CidrBlocksOutput() *string
	ElasticacheSubnetsOutput() *string
	ElasticacheSubnetSuffix() *string
	SetElasticacheSubnetSuffix(val *string)
	ElasticacheSubnetTags() *map[string]*string
	SetElasticacheSubnetTags(val *map[string]*string)
	EnableClassiclink() *bool
	SetEnableClassiclink(val *bool)
	EnableClassiclinkDnsSupport() *bool
	SetEnableClassiclinkDnsSupport(val *bool)
	EnableDhcpOptions() *bool
	SetEnableDhcpOptions(val *bool)
	EnableDnsHostnames() *bool
	SetEnableDnsHostnames(val *bool)
	EnableDnsSupport() *bool
	SetEnableDnsSupport(val *bool)
	EnableFlowLog() *bool
	SetEnableFlowLog(val *bool)
	EnableIpv6() *bool
	SetEnableIpv6(val *bool)
	EnableNatGateway() *bool
	SetEnableNatGateway(val *bool)
	EnablePublicRedshift() *bool
	SetEnablePublicRedshift(val *bool)
	EnableVpnGateway() *bool
	SetEnableVpnGateway(val *bool)
	ExternalNatIpIds() *[]*string
	SetExternalNatIpIds(val *[]*string)
	ExternalNatIps() *[]*string
	SetExternalNatIps(val *[]*string)
	FlowLogCloudwatchIamRoleArn() *string
	SetFlowLogCloudwatchIamRoleArn(val *string)
	FlowLogCloudwatchLogGroupKmsKeyId() *string
	SetFlowLogCloudwatchLogGroupKmsKeyId(val *string)
	FlowLogCloudwatchLogGroupNamePrefix() *string
	SetFlowLogCloudwatchLogGroupNamePrefix(val *string)
	FlowLogCloudwatchLogGroupNameSuffix() *string
	SetFlowLogCloudwatchLogGroupNameSuffix(val *string)
	FlowLogCloudwatchLogGroupRetentionInDays() *float64
	SetFlowLogCloudwatchLogGroupRetentionInDays(val *float64)
	FlowLogDestinationArn() *string
	SetFlowLogDestinationArn(val *string)
	FlowLogDestinationType() *string
	SetFlowLogDestinationType(val *string)
	FlowLogFileFormat() *string
	SetFlowLogFileFormat(val *string)
	FlowLogHiveCompatiblePartitions() *bool
	SetFlowLogHiveCompatiblePartitions(val *bool)
	FlowLogLogFormat() *string
	SetFlowLogLogFormat(val *string)
	FlowLogMaxAggregationInterval() *float64
	SetFlowLogMaxAggregationInterval(val *float64)
	FlowLogPerHourPartition() *bool
	SetFlowLogPerHourPartition(val *bool)
	FlowLogTrafficType() *string
	SetFlowLogTrafficType(val *string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IgwArnOutput() *string
	IgwIdOutput() *string
	IgwTags() *map[string]*string
	SetIgwTags(val *map[string]*string)
	InstanceTenancy() *string
	SetInstanceTenancy(val *string)
	IntraAclTags() *map[string]*string
	SetIntraAclTags(val *map[string]*string)
	IntraDedicatedNetworkAcl() *bool
	SetIntraDedicatedNetworkAcl(val *bool)
	IntraInboundAclRules() *[]*map[string]*string
	SetIntraInboundAclRules(val *[]*map[string]*string)
	IntraNetworkAclArnOutput() *string
	IntraNetworkAclIdOutput() *string
	IntraOutboundAclRules() *[]*map[string]*string
	SetIntraOutboundAclRules(val *[]*map[string]*string)
	IntraRouteTableAssociationIdsOutput() *string
	IntraRouteTableIdsOutput() *string
	IntraRouteTableTags() *map[string]*string
	SetIntraRouteTableTags(val *map[string]*string)
	IntraSubnetArnsOutput() *string
	IntraSubnetAssignIpv6AddressOnCreation() *bool
	SetIntraSubnetAssignIpv6AddressOnCreation(val *bool)
	IntraSubnetIpv6Prefixes() *[]*string
	SetIntraSubnetIpv6Prefixes(val *[]*string)
	IntraSubnetNames() *[]*string
	SetIntraSubnetNames(val *[]*string)
	IntraSubnets() *[]*string
	SetIntraSubnets(val *[]*string)
	IntraSubnetsCidrBlocksOutput() *string
	IntraSubnetsIpv6CidrBlocksOutput() *string
	IntraSubnetsOutput() *string
	IntraSubnetSuffix() *string
	SetIntraSubnetSuffix(val *string)
	IntraSubnetTags() *map[string]*string
	SetIntraSubnetTags(val *map[string]*string)
	Ipv4IpamPoolId() *string
	SetIpv4IpamPoolId(val *string)
	Ipv4NetmaskLength() *float64
	SetIpv4NetmaskLength(val *float64)
	Ipv6Cidr() *string
	SetIpv6Cidr(val *string)
	Ipv6IpamPoolId() *string
	SetIpv6IpamPoolId(val *string)
	Ipv6NetmaskLength() *float64
	SetIpv6NetmaskLength(val *float64)
	ManageDefaultNetworkAcl() *bool
	SetManageDefaultNetworkAcl(val *bool)
	ManageDefaultRouteTable() *bool
	SetManageDefaultRouteTable(val *bool)
	ManageDefaultSecurityGroup() *bool
	SetManageDefaultSecurityGroup(val *bool)
	ManageDefaultVpc() *bool
	SetManageDefaultVpc(val *bool)
	MapPublicIpOnLaunch() *bool
	SetMapPublicIpOnLaunch(val *bool)
	Name() *string
	SetName(val *string)
	NameOutput() *string
	NatEipTags() *map[string]*string
	SetNatEipTags(val *map[string]*string)
	NatGatewayDestinationCidrBlock() *string
	SetNatGatewayDestinationCidrBlock(val *string)
	NatGatewayTags() *map[string]*string
	SetNatGatewayTags(val *map[string]*string)
	NatgwIdsOutput() *string
	NatIdsOutput() *string
	NatPublicIpsOutput() *string
	// The tree node.
	Node() constructs.Node
	OneNatGatewayPerAz() *bool
	SetOneNatGatewayPerAz(val *bool)
	OutpostAclTags() *map[string]*string
	SetOutpostAclTags(val *map[string]*string)
	OutpostArn() *string
	SetOutpostArn(val *string)
	OutpostAz() *string
	SetOutpostAz(val *string)
	OutpostDedicatedNetworkAcl() *bool
	SetOutpostDedicatedNetworkAcl(val *bool)
	OutpostInboundAclRules() *[]*map[string]*string
	SetOutpostInboundAclRules(val *[]*map[string]*string)
	OutpostNetworkAclArnOutput() *string
	OutpostNetworkAclIdOutput() *string
	OutpostOutboundAclRules() *[]*map[string]*string
	SetOutpostOutboundAclRules(val *[]*map[string]*string)
	OutpostSubnetArnsOutput() *string
	OutpostSubnetAssignIpv6AddressOnCreation() *bool
	SetOutpostSubnetAssignIpv6AddressOnCreation(val *bool)
	OutpostSubnetIpv6Prefixes() *[]*string
	SetOutpostSubnetIpv6Prefixes(val *[]*string)
	OutpostSubnetNames() *[]*string
	SetOutpostSubnetNames(val *[]*string)
	OutpostSubnets() *[]*string
	SetOutpostSubnets(val *[]*string)
	OutpostSubnetsCidrBlocksOutput() *string
	OutpostSubnetsIpv6CidrBlocksOutput() *string
	OutpostSubnetsOutput() *string
	OutpostSubnetSuffix() *string
	SetOutpostSubnetSuffix(val *string)
	OutpostSubnetTags() *map[string]*string
	SetOutpostSubnetTags(val *map[string]*string)
	PrivateAclTags() *map[string]*string
	SetPrivateAclTags(val *map[string]*string)
	PrivateDedicatedNetworkAcl() *bool
	SetPrivateDedicatedNetworkAcl(val *bool)
	PrivateInboundAclRules() *[]*map[string]*string
	SetPrivateInboundAclRules(val *[]*map[string]*string)
	PrivateIpv6EgressRouteIdsOutput() *string
	PrivateNatGatewayRouteIdsOutput() *string
	PrivateNetworkAclArnOutput() *string
	PrivateNetworkAclIdOutput() *string
	PrivateOutboundAclRules() *[]*map[string]*string
	SetPrivateOutboundAclRules(val *[]*map[string]*string)
	PrivateRouteTableAssociationIdsOutput() *string
	PrivateRouteTableIdsOutput() *string
	PrivateRouteTableTags() *map[string]*string
	SetPrivateRouteTableTags(val *map[string]*string)
	PrivateSubnetArnsOutput() *string
	PrivateSubnetAssignIpv6AddressOnCreation() *bool
	SetPrivateSubnetAssignIpv6AddressOnCreation(val *bool)
	PrivateSubnetIpv6Prefixes() *[]*string
	SetPrivateSubnetIpv6Prefixes(val *[]*string)
	PrivateSubnetNames() *[]*string
	SetPrivateSubnetNames(val *[]*string)
	PrivateSubnets() *[]*string
	SetPrivateSubnets(val *[]*string)
	PrivateSubnetsCidrBlocksOutput() *string
	PrivateSubnetsIpv6CidrBlocksOutput() *string
	PrivateSubnetsOutput() *string
	PrivateSubnetSuffix() *string
	SetPrivateSubnetSuffix(val *string)
	PrivateSubnetTags() *map[string]*string
	SetPrivateSubnetTags(val *map[string]*string)
	PrivateSubnetTagsPerAz() *map[string]*map[string]*string
	SetPrivateSubnetTagsPerAz(val *map[string]*map[string]*string)
	PropagateIntraRouteTablesVgw() *bool
	SetPropagateIntraRouteTablesVgw(val *bool)
	PropagatePrivateRouteTablesVgw() *bool
	SetPropagatePrivateRouteTablesVgw(val *bool)
	PropagatePublicRouteTablesVgw() *bool
	SetPropagatePublicRouteTablesVgw(val *bool)
	// Experimental.
	Providers() *[]interface{}
	PublicAclTags() *map[string]*string
	SetPublicAclTags(val *map[string]*string)
	PublicDedicatedNetworkAcl() *bool
	SetPublicDedicatedNetworkAcl(val *bool)
	PublicInboundAclRules() *[]*map[string]*string
	SetPublicInboundAclRules(val *[]*map[string]*string)
	PublicInternetGatewayIpv6RouteIdOutput() *string
	PublicInternetGatewayRouteIdOutput() *string
	PublicNetworkAclArnOutput() *string
	PublicNetworkAclIdOutput() *string
	PublicOutboundAclRules() *[]*map[string]*string
	SetPublicOutboundAclRules(val *[]*map[string]*string)
	PublicRouteTableAssociationIdsOutput() *string
	PublicRouteTableIdsOutput() *string
	PublicRouteTableTags() *map[string]*string
	SetPublicRouteTableTags(val *map[string]*string)
	PublicSubnetArnsOutput() *string
	PublicSubnetAssignIpv6AddressOnCreation() *bool
	SetPublicSubnetAssignIpv6AddressOnCreation(val *bool)
	PublicSubnetIpv6Prefixes() *[]*string
	SetPublicSubnetIpv6Prefixes(val *[]*string)
	PublicSubnetNames() *[]*string
	SetPublicSubnetNames(val *[]*string)
	PublicSubnets() *[]*string
	SetPublicSubnets(val *[]*string)
	PublicSubnetsCidrBlocksOutput() *string
	PublicSubnetsIpv6CidrBlocksOutput() *string
	PublicSubnetsOutput() *string
	PublicSubnetSuffix() *string
	SetPublicSubnetSuffix(val *string)
	PublicSubnetTags() *map[string]*string
	SetPublicSubnetTags(val *map[string]*string)
	PublicSubnetTagsPerAz() *map[string]*map[string]*string
	SetPublicSubnetTagsPerAz(val *map[string]*map[string]*string)
	PutinKhuylo() *bool
	SetPutinKhuylo(val *bool)
	// Experimental.
	RawOverrides() interface{}
	RedshiftAclTags() *map[string]*string
	SetRedshiftAclTags(val *map[string]*string)
	RedshiftDedicatedNetworkAcl() *bool
	SetRedshiftDedicatedNetworkAcl(val *bool)
	RedshiftInboundAclRules() *[]*map[string]*string
	SetRedshiftInboundAclRules(val *[]*map[string]*string)
	RedshiftNetworkAclArnOutput() *string
	RedshiftNetworkAclIdOutput() *string
	RedshiftOutboundAclRules() *[]*map[string]*string
	SetRedshiftOutboundAclRules(val *[]*map[string]*string)
	RedshiftPublicRouteTableAssociationIdsOutput() *string
	RedshiftRouteTableAssociationIdsOutput() *string
	RedshiftRouteTableIdsOutput() *string
	RedshiftRouteTableTags() *map[string]*string
	SetRedshiftRouteTableTags(val *map[string]*string)
	RedshiftSubnetArnsOutput() *string
	RedshiftSubnetAssignIpv6AddressOnCreation() *bool
	SetRedshiftSubnetAssignIpv6AddressOnCreation(val *bool)
	RedshiftSubnetGroupName() *string
	SetRedshiftSubnetGroupName(val *string)
	RedshiftSubnetGroupOutput() *string
	RedshiftSubnetGroupTags() *map[string]*string
	SetRedshiftSubnetGroupTags(val *map[string]*string)
	RedshiftSubnetIpv6Prefixes() *[]*string
	SetRedshiftSubnetIpv6Prefixes(val *[]*string)
	RedshiftSubnetNames() *[]*string
	SetRedshiftSubnetNames(val *[]*string)
	RedshiftSubnets() *[]*string
	SetRedshiftSubnets(val *[]*string)
	RedshiftSubnetsCidrBlocksOutput() *string
	RedshiftSubnetsIpv6CidrBlocksOutput() *string
	RedshiftSubnetsOutput() *string
	RedshiftSubnetSuffix() *string
	SetRedshiftSubnetSuffix(val *string)
	RedshiftSubnetTags() *map[string]*string
	SetRedshiftSubnetTags(val *map[string]*string)
	ReuseNatIps() *bool
	SetReuseNatIps(val *bool)
	SecondaryCidrBlocks() *[]*string
	SetSecondaryCidrBlocks(val *[]*string)
	SingleNatGateway() *bool
	SetSingleNatGateway(val *bool)
	// Experimental.
	SkipAssetCreationFromLocalModules() *bool
	// Experimental.
	Source() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	ThisCustomerGatewayOutput() *string
	UseIpamPool() *bool
	SetUseIpamPool(val *bool)
	// Experimental.
	Version() *string
	VgwArnOutput() *string
	VgwIdOutput() *string
	VpcArnOutput() *string
	VpcCidrBlockOutput() *string
	VpcEnableDnsHostnamesOutput() *string
	VpcEnableDnsSupportOutput() *string
	VpcFlowLogCloudwatchIamRoleArnOutput() *string
	VpcFlowLogDestinationArnOutput() *string
	VpcFlowLogDestinationTypeOutput() *string
	VpcFlowLogIdOutput() *string
	VpcFlowLogPermissionsBoundary() *string
	SetVpcFlowLogPermissionsBoundary(val *string)
	VpcFlowLogTags() *map[string]*string
	SetVpcFlowLogTags(val *map[string]*string)
	VpcIdOutput() *string
	VpcInstanceTenancyOutput() *string
	VpcIpv6AssociationIdOutput() *string
	VpcIpv6CidrBlockOutput() *string
	VpcMainRouteTableIdOutput() *string
	VpcOwnerIdOutput() *string
	VpcSecondaryCidrBlocksOutput() *string
	VpcTags() *map[string]*string
	SetVpcTags(val *map[string]*string)
	VpnGatewayAz() *string
	SetVpnGatewayAz(val *string)
	VpnGatewayId() *string
	SetVpnGatewayId(val *string)
	VpnGatewayTags() *map[string]*string
	SetVpnGatewayTags(val *map[string]*string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	AddProvider(provider interface{})
	// Experimental.
	GetString(output *string) *string
	// Experimental.
	InterpolationForOutput(moduleOutput *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Awsvpc
type jsiiProxy_Awsvpc struct {
	internal.Type__cdktfTerraformModule
}

func (j *jsiiProxy_Awsvpc) AmazonSideAsn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonSideAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) AssignIpv6AddressOnCreation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignIpv6AddressOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) Azs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"azs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) AzsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CgwArnsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cgwArnsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CgwIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cgwIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CreateDatabaseInternetGatewayRoute() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createDatabaseInternetGatewayRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CreateDatabaseNatGatewayRoute() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createDatabaseNatGatewayRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CreateDatabaseSubnetGroup() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createDatabaseSubnetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CreateDatabaseSubnetRouteTable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createDatabaseSubnetRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CreateEgressOnlyIgw() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createEgressOnlyIgw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CreateElasticacheSubnetGroup() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createElasticacheSubnetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CreateElasticacheSubnetRouteTable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createElasticacheSubnetRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CreateFlowLogCloudwatchIamRole() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createFlowLogCloudwatchIamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CreateFlowLogCloudwatchLogGroup() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createFlowLogCloudwatchLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CreateIgw() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createIgw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CreateRedshiftSubnetGroup() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createRedshiftSubnetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CreateRedshiftSubnetRouteTable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createRedshiftSubnetRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CreateVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CustomerGateways() *map[string]*map[string]interface{} {
	var returns *map[string]*map[string]interface{}
	_jsii_.Get(
		j,
		"customerGateways",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) CustomerGatewayTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customerGatewayTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseAclTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"databaseAclTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseDedicatedNetworkAcl() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"databaseDedicatedNetworkAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseInboundAclRules() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"databaseInboundAclRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseInternetGatewayRouteIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInternetGatewayRouteIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseIpv6EgressRouteIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseIpv6EgressRouteIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseNatGatewayRouteIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNatGatewayRouteIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseNetworkAclArnOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNetworkAclArnOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseNetworkAclIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNetworkAclIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseOutboundAclRules() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"databaseOutboundAclRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseRouteTableAssociationIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRouteTableAssociationIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseRouteTableIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRouteTableIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseRouteTableTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"databaseRouteTableTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseSubnetArnsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseSubnetArnsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseSubnetAssignIpv6AddressOnCreation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"databaseSubnetAssignIpv6AddressOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseSubnetGroupNameOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseSubnetGroupNameOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseSubnetGroupOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseSubnetGroupOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseSubnetGroupTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"databaseSubnetGroupTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseSubnetIpv6Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"databaseSubnetIpv6Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseSubnetNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"databaseSubnetNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"databaseSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseSubnetsCidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseSubnetsCidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseSubnetsIpv6CidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseSubnetsIpv6CidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseSubnetsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseSubnetsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseSubnetSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseSubnetSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DatabaseSubnetTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"databaseSubnetTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultNetworkAclEgress() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"defaultNetworkAclEgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultNetworkAclIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNetworkAclIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultNetworkAclIngress() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"defaultNetworkAclIngress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultNetworkAclName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNetworkAclName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultNetworkAclTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultNetworkAclTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultRouteTableIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRouteTableIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultRouteTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRouteTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultRouteTablePropagatingVgws() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultRouteTablePropagatingVgws",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultRouteTableRoutes() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"defaultRouteTableRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultRouteTableTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultRouteTableTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultSecurityGroupEgress() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"defaultSecurityGroupEgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultSecurityGroupIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSecurityGroupIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultSecurityGroupIngress() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"defaultSecurityGroupIngress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultSecurityGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSecurityGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultSecurityGroupTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultSecurityGroupTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcArnOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVpcArnOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcCidrBlockOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVpcCidrBlockOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcDefaultNetworkAclIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVpcDefaultNetworkAclIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcDefaultRouteTableIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVpcDefaultRouteTableIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcDefaultSecurityGroupIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVpcDefaultSecurityGroupIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcEnableClassiclink() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"defaultVpcEnableClassiclink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcEnableDnsHostnames() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"defaultVpcEnableDnsHostnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcEnableDnsHostnamesOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVpcEnableDnsHostnamesOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcEnableDnsSupport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"defaultVpcEnableDnsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcEnableDnsSupportOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVpcEnableDnsSupportOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVpcIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcInstanceTenancyOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVpcInstanceTenancyOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcMainRouteTableIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVpcMainRouteTableIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVpcName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DefaultVpcTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultVpcTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DhcpOptionsDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhcpOptionsDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DhcpOptionsDomainNameServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dhcpOptionsDomainNameServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DhcpOptionsIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhcpOptionsIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DhcpOptionsNetbiosNameServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dhcpOptionsNetbiosNameServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DhcpOptionsNetbiosNodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhcpOptionsNetbiosNodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DhcpOptionsNtpServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dhcpOptionsNtpServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) DhcpOptionsTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dhcpOptionsTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) EgressOnlyInternetGatewayIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressOnlyInternetGatewayIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheAclTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"elasticacheAclTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheDedicatedNetworkAcl() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"elasticacheDedicatedNetworkAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheInboundAclRules() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"elasticacheInboundAclRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheNetworkAclArnOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticacheNetworkAclArnOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheNetworkAclIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticacheNetworkAclIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheOutboundAclRules() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"elasticacheOutboundAclRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheRouteTableAssociationIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticacheRouteTableAssociationIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheRouteTableIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticacheRouteTableIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheRouteTableTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"elasticacheRouteTableTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheSubnetArnsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticacheSubnetArnsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheSubnetAssignIpv6AddressOnCreation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"elasticacheSubnetAssignIpv6AddressOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticacheSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheSubnetGroupNameOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticacheSubnetGroupNameOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheSubnetGroupOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticacheSubnetGroupOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheSubnetGroupTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"elasticacheSubnetGroupTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheSubnetIpv6Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elasticacheSubnetIpv6Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheSubnetNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elasticacheSubnetNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elasticacheSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheSubnetsCidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticacheSubnetsCidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheSubnetsIpv6CidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticacheSubnetsIpv6CidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheSubnetsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticacheSubnetsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheSubnetSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticacheSubnetSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ElasticacheSubnetTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"elasticacheSubnetTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) EnableClassiclink() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableClassiclink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) EnableClassiclinkDnsSupport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableClassiclinkDnsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) EnableDhcpOptions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableDhcpOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) EnableDnsHostnames() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableDnsHostnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) EnableDnsSupport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableDnsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) EnableFlowLog() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableFlowLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) EnableIpv6() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) EnableNatGateway() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableNatGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) EnablePublicRedshift() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablePublicRedshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) EnableVpnGateway() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableVpnGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ExternalNatIpIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalNatIpIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ExternalNatIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalNatIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) FlowLogCloudwatchIamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogCloudwatchIamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) FlowLogCloudwatchLogGroupKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogCloudwatchLogGroupKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) FlowLogCloudwatchLogGroupNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogCloudwatchLogGroupNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) FlowLogCloudwatchLogGroupNameSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogCloudwatchLogGroupNameSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) FlowLogCloudwatchLogGroupRetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flowLogCloudwatchLogGroupRetentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) FlowLogDestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogDestinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) FlowLogDestinationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogDestinationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) FlowLogFileFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogFileFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) FlowLogHiveCompatiblePartitions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"flowLogHiveCompatiblePartitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) FlowLogLogFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogLogFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) FlowLogMaxAggregationInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flowLogMaxAggregationInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) FlowLogPerHourPartition() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"flowLogPerHourPartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) FlowLogTrafficType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowLogTrafficType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IgwArnOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"igwArnOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IgwIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"igwIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IgwTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"igwTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) InstanceTenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraAclTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"intraAclTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraDedicatedNetworkAcl() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"intraDedicatedNetworkAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraInboundAclRules() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"intraInboundAclRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraNetworkAclArnOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intraNetworkAclArnOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraNetworkAclIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intraNetworkAclIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraOutboundAclRules() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"intraOutboundAclRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraRouteTableAssociationIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intraRouteTableAssociationIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraRouteTableIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intraRouteTableIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraRouteTableTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"intraRouteTableTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraSubnetArnsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intraSubnetArnsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraSubnetAssignIpv6AddressOnCreation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"intraSubnetAssignIpv6AddressOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraSubnetIpv6Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"intraSubnetIpv6Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraSubnetNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"intraSubnetNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"intraSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraSubnetsCidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intraSubnetsCidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraSubnetsIpv6CidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intraSubnetsIpv6CidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraSubnetsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intraSubnetsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraSubnetSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intraSubnetSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) IntraSubnetTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"intraSubnetTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) Ipv4IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4IpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) Ipv4NetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4NetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) Ipv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) Ipv6IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6IpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) Ipv6NetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6NetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ManageDefaultNetworkAcl() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"manageDefaultNetworkAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ManageDefaultRouteTable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"manageDefaultRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ManageDefaultSecurityGroup() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"manageDefaultSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ManageDefaultVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"manageDefaultVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) MapPublicIpOnLaunch() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"mapPublicIpOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) NameOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) NatEipTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"natEipTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) NatGatewayDestinationCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayDestinationCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) NatGatewayTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"natGatewayTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) NatgwIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natgwIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) NatIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) NatPublicIpsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natPublicIpsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OneNatGatewayPerAz() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"oneNatGatewayPerAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostAclTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"outpostAclTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostAz() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostDedicatedNetworkAcl() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"outpostDedicatedNetworkAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostInboundAclRules() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"outpostInboundAclRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostNetworkAclArnOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostNetworkAclArnOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostNetworkAclIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostNetworkAclIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostOutboundAclRules() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"outpostOutboundAclRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostSubnetArnsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostSubnetArnsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostSubnetAssignIpv6AddressOnCreation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"outpostSubnetAssignIpv6AddressOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostSubnetIpv6Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outpostSubnetIpv6Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostSubnetNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outpostSubnetNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outpostSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostSubnetsCidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostSubnetsCidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostSubnetsIpv6CidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostSubnetsIpv6CidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostSubnetsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostSubnetsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostSubnetSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostSubnetSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) OutpostSubnetTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"outpostSubnetTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateAclTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"privateAclTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateDedicatedNetworkAcl() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"privateDedicatedNetworkAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateInboundAclRules() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"privateInboundAclRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateIpv6EgressRouteIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpv6EgressRouteIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateNatGatewayRouteIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateNatGatewayRouteIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateNetworkAclArnOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateNetworkAclArnOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateNetworkAclIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateNetworkAclIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateOutboundAclRules() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"privateOutboundAclRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateRouteTableAssociationIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateRouteTableAssociationIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateRouteTableIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateRouteTableIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateRouteTableTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"privateRouteTableTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateSubnetArnsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateSubnetArnsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateSubnetAssignIpv6AddressOnCreation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"privateSubnetAssignIpv6AddressOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateSubnetIpv6Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateSubnetIpv6Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateSubnetNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateSubnetNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateSubnetsCidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateSubnetsCidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateSubnetsIpv6CidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateSubnetsIpv6CidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateSubnetsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateSubnetsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateSubnetSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateSubnetSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateSubnetTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"privateSubnetTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PrivateSubnetTagsPerAz() *map[string]*map[string]*string {
	var returns *map[string]*map[string]*string
	_jsii_.Get(
		j,
		"privateSubnetTagsPerAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PropagateIntraRouteTablesVgw() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"propagateIntraRouteTablesVgw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PropagatePrivateRouteTablesVgw() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"propagatePrivateRouteTablesVgw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PropagatePublicRouteTablesVgw() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"propagatePublicRouteTablesVgw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) Providers() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicAclTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"publicAclTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicDedicatedNetworkAcl() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"publicDedicatedNetworkAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicInboundAclRules() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"publicInboundAclRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicInternetGatewayIpv6RouteIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicInternetGatewayIpv6RouteIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicInternetGatewayRouteIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicInternetGatewayRouteIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicNetworkAclArnOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicNetworkAclArnOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicNetworkAclIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicNetworkAclIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicOutboundAclRules() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"publicOutboundAclRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicRouteTableAssociationIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicRouteTableAssociationIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicRouteTableIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicRouteTableIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicRouteTableTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"publicRouteTableTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicSubnetArnsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSubnetArnsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicSubnetAssignIpv6AddressOnCreation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"publicSubnetAssignIpv6AddressOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicSubnetIpv6Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicSubnetIpv6Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicSubnetNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicSubnetNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicSubnetsCidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSubnetsCidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicSubnetsIpv6CidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSubnetsIpv6CidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicSubnetsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSubnetsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicSubnetSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSubnetSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicSubnetTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"publicSubnetTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PublicSubnetTagsPerAz() *map[string]*map[string]*string {
	var returns *map[string]*map[string]*string
	_jsii_.Get(
		j,
		"publicSubnetTagsPerAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) PutinKhuylo() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"putinKhuylo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftAclTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"redshiftAclTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftDedicatedNetworkAcl() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"redshiftDedicatedNetworkAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftInboundAclRules() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"redshiftInboundAclRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftNetworkAclArnOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftNetworkAclArnOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftNetworkAclIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftNetworkAclIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftOutboundAclRules() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"redshiftOutboundAclRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftPublicRouteTableAssociationIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftPublicRouteTableAssociationIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftRouteTableAssociationIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftRouteTableAssociationIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftRouteTableIdsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftRouteTableIdsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftRouteTableTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"redshiftRouteTableTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftSubnetArnsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftSubnetArnsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftSubnetAssignIpv6AddressOnCreation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"redshiftSubnetAssignIpv6AddressOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftSubnetGroupOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftSubnetGroupOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftSubnetGroupTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"redshiftSubnetGroupTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftSubnetIpv6Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redshiftSubnetIpv6Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftSubnetNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redshiftSubnetNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redshiftSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftSubnetsCidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftSubnetsCidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftSubnetsIpv6CidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftSubnetsIpv6CidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftSubnetsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftSubnetsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftSubnetSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftSubnetSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) RedshiftSubnetTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"redshiftSubnetTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ReuseNatIps() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"reuseNatIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) SecondaryCidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryCidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) SingleNatGateway() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"singleNatGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) SkipAssetCreationFromLocalModules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipAssetCreationFromLocalModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) ThisCustomerGatewayOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thisCustomerGatewayOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) UseIpamPool() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"useIpamPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VgwArnOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vgwArnOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VgwIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vgwIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcArnOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcArnOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcCidrBlockOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcCidrBlockOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcEnableDnsHostnamesOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEnableDnsHostnamesOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcEnableDnsSupportOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEnableDnsSupportOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcFlowLogCloudwatchIamRoleArnOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcFlowLogCloudwatchIamRoleArnOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcFlowLogDestinationArnOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcFlowLogDestinationArnOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcFlowLogDestinationTypeOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcFlowLogDestinationTypeOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcFlowLogIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcFlowLogIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcFlowLogPermissionsBoundary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcFlowLogPermissionsBoundary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcFlowLogTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"vpcFlowLogTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcInstanceTenancyOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcInstanceTenancyOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcIpv6AssociationIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIpv6AssociationIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcIpv6CidrBlockOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIpv6CidrBlockOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcMainRouteTableIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcMainRouteTableIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcOwnerIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcOwnerIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcSecondaryCidrBlocksOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcSecondaryCidrBlocksOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpcTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"vpcTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpnGatewayAz() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpnGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Awsvpc) VpnGatewayTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"vpnGatewayTags",
		&returns,
	)
	return returns
}


func NewAwsvpc(scope constructs.Construct, id *string, config *AwsvpcConfig) Awsvpc {
	_init_.Initialize()

	if err := validateNewAwsvpcParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Awsvpc{}

	_jsii_.Create(
		"@cdktf/provider-awsvpc.Awsvpc",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

func NewAwsvpc_Override(a Awsvpc, scope constructs.Construct, id *string, config *AwsvpcConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-awsvpc.Awsvpc",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Awsvpc)SetAmazonSideAsn(val *string) {
	_jsii_.Set(
		j,
		"amazonSideAsn",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetAssignIpv6AddressOnCreation(val *bool) {
	_jsii_.Set(
		j,
		"assignIpv6AddressOnCreation",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetAzs(val *[]*string) {
	_jsii_.Set(
		j,
		"azs",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCidr(val *string) {
	_jsii_.Set(
		j,
		"cidr",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCreateDatabaseInternetGatewayRoute(val *bool) {
	_jsii_.Set(
		j,
		"createDatabaseInternetGatewayRoute",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCreateDatabaseNatGatewayRoute(val *bool) {
	_jsii_.Set(
		j,
		"createDatabaseNatGatewayRoute",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCreateDatabaseSubnetGroup(val *bool) {
	_jsii_.Set(
		j,
		"createDatabaseSubnetGroup",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCreateDatabaseSubnetRouteTable(val *bool) {
	_jsii_.Set(
		j,
		"createDatabaseSubnetRouteTable",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCreateEgressOnlyIgw(val *bool) {
	_jsii_.Set(
		j,
		"createEgressOnlyIgw",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCreateElasticacheSubnetGroup(val *bool) {
	_jsii_.Set(
		j,
		"createElasticacheSubnetGroup",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCreateElasticacheSubnetRouteTable(val *bool) {
	_jsii_.Set(
		j,
		"createElasticacheSubnetRouteTable",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCreateFlowLogCloudwatchIamRole(val *bool) {
	_jsii_.Set(
		j,
		"createFlowLogCloudwatchIamRole",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCreateFlowLogCloudwatchLogGroup(val *bool) {
	_jsii_.Set(
		j,
		"createFlowLogCloudwatchLogGroup",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCreateIgw(val *bool) {
	_jsii_.Set(
		j,
		"createIgw",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCreateRedshiftSubnetGroup(val *bool) {
	_jsii_.Set(
		j,
		"createRedshiftSubnetGroup",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCreateRedshiftSubnetRouteTable(val *bool) {
	_jsii_.Set(
		j,
		"createRedshiftSubnetRouteTable",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCreateVpc(val *bool) {
	_jsii_.Set(
		j,
		"createVpc",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCustomerGateways(val *map[string]*map[string]interface{}) {
	_jsii_.Set(
		j,
		"customerGateways",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetCustomerGatewayTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"customerGatewayTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDatabaseAclTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"databaseAclTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDatabaseDedicatedNetworkAcl(val *bool) {
	_jsii_.Set(
		j,
		"databaseDedicatedNetworkAcl",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDatabaseInboundAclRules(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"databaseInboundAclRules",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDatabaseOutboundAclRules(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"databaseOutboundAclRules",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDatabaseRouteTableTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"databaseRouteTableTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDatabaseSubnetAssignIpv6AddressOnCreation(val *bool) {
	_jsii_.Set(
		j,
		"databaseSubnetAssignIpv6AddressOnCreation",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDatabaseSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"databaseSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDatabaseSubnetGroupTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"databaseSubnetGroupTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDatabaseSubnetIpv6Prefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"databaseSubnetIpv6Prefixes",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDatabaseSubnetNames(val *[]*string) {
	_jsii_.Set(
		j,
		"databaseSubnetNames",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDatabaseSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"databaseSubnets",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDatabaseSubnetSuffix(val *string) {
	_jsii_.Set(
		j,
		"databaseSubnetSuffix",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDatabaseSubnetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"databaseSubnetTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultNetworkAclEgress(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"defaultNetworkAclEgress",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultNetworkAclIngress(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"defaultNetworkAclIngress",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultNetworkAclName(val *string) {
	_jsii_.Set(
		j,
		"defaultNetworkAclName",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultNetworkAclTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"defaultNetworkAclTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultRouteTableName(val *string) {
	_jsii_.Set(
		j,
		"defaultRouteTableName",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultRouteTablePropagatingVgws(val *[]*string) {
	_jsii_.Set(
		j,
		"defaultRouteTablePropagatingVgws",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultRouteTableRoutes(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"defaultRouteTableRoutes",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultRouteTableTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"defaultRouteTableTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultSecurityGroupEgress(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"defaultSecurityGroupEgress",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultSecurityGroupIngress(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"defaultSecurityGroupIngress",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultSecurityGroupName(val *string) {
	_jsii_.Set(
		j,
		"defaultSecurityGroupName",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultSecurityGroupTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"defaultSecurityGroupTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultVpcEnableClassiclink(val *bool) {
	_jsii_.Set(
		j,
		"defaultVpcEnableClassiclink",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultVpcEnableDnsHostnames(val *bool) {
	_jsii_.Set(
		j,
		"defaultVpcEnableDnsHostnames",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultVpcEnableDnsSupport(val *bool) {
	_jsii_.Set(
		j,
		"defaultVpcEnableDnsSupport",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultVpcName(val *string) {
	_jsii_.Set(
		j,
		"defaultVpcName",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDefaultVpcTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"defaultVpcTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDhcpOptionsDomainName(val *string) {
	_jsii_.Set(
		j,
		"dhcpOptionsDomainName",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDhcpOptionsDomainNameServers(val *[]*string) {
	_jsii_.Set(
		j,
		"dhcpOptionsDomainNameServers",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDhcpOptionsNetbiosNameServers(val *[]*string) {
	_jsii_.Set(
		j,
		"dhcpOptionsNetbiosNameServers",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDhcpOptionsNetbiosNodeType(val *string) {
	_jsii_.Set(
		j,
		"dhcpOptionsNetbiosNodeType",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDhcpOptionsNtpServers(val *[]*string) {
	_jsii_.Set(
		j,
		"dhcpOptionsNtpServers",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetDhcpOptionsTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"dhcpOptionsTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetElasticacheAclTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"elasticacheAclTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetElasticacheDedicatedNetworkAcl(val *bool) {
	_jsii_.Set(
		j,
		"elasticacheDedicatedNetworkAcl",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetElasticacheInboundAclRules(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"elasticacheInboundAclRules",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetElasticacheOutboundAclRules(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"elasticacheOutboundAclRules",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetElasticacheRouteTableTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"elasticacheRouteTableTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetElasticacheSubnetAssignIpv6AddressOnCreation(val *bool) {
	_jsii_.Set(
		j,
		"elasticacheSubnetAssignIpv6AddressOnCreation",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetElasticacheSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"elasticacheSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetElasticacheSubnetGroupTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"elasticacheSubnetGroupTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetElasticacheSubnetIpv6Prefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"elasticacheSubnetIpv6Prefixes",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetElasticacheSubnetNames(val *[]*string) {
	_jsii_.Set(
		j,
		"elasticacheSubnetNames",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetElasticacheSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"elasticacheSubnets",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetElasticacheSubnetSuffix(val *string) {
	_jsii_.Set(
		j,
		"elasticacheSubnetSuffix",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetElasticacheSubnetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"elasticacheSubnetTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetEnableClassiclink(val *bool) {
	_jsii_.Set(
		j,
		"enableClassiclink",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetEnableClassiclinkDnsSupport(val *bool) {
	_jsii_.Set(
		j,
		"enableClassiclinkDnsSupport",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetEnableDhcpOptions(val *bool) {
	_jsii_.Set(
		j,
		"enableDhcpOptions",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetEnableDnsHostnames(val *bool) {
	_jsii_.Set(
		j,
		"enableDnsHostnames",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetEnableDnsSupport(val *bool) {
	_jsii_.Set(
		j,
		"enableDnsSupport",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetEnableFlowLog(val *bool) {
	_jsii_.Set(
		j,
		"enableFlowLog",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetEnableIpv6(val *bool) {
	_jsii_.Set(
		j,
		"enableIpv6",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetEnableNatGateway(val *bool) {
	_jsii_.Set(
		j,
		"enableNatGateway",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetEnablePublicRedshift(val *bool) {
	_jsii_.Set(
		j,
		"enablePublicRedshift",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetEnableVpnGateway(val *bool) {
	_jsii_.Set(
		j,
		"enableVpnGateway",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetExternalNatIpIds(val *[]*string) {
	_jsii_.Set(
		j,
		"externalNatIpIds",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetExternalNatIps(val *[]*string) {
	_jsii_.Set(
		j,
		"externalNatIps",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetFlowLogCloudwatchIamRoleArn(val *string) {
	_jsii_.Set(
		j,
		"flowLogCloudwatchIamRoleArn",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetFlowLogCloudwatchLogGroupKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"flowLogCloudwatchLogGroupKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetFlowLogCloudwatchLogGroupNamePrefix(val *string) {
	_jsii_.Set(
		j,
		"flowLogCloudwatchLogGroupNamePrefix",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetFlowLogCloudwatchLogGroupNameSuffix(val *string) {
	_jsii_.Set(
		j,
		"flowLogCloudwatchLogGroupNameSuffix",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetFlowLogCloudwatchLogGroupRetentionInDays(val *float64) {
	_jsii_.Set(
		j,
		"flowLogCloudwatchLogGroupRetentionInDays",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetFlowLogDestinationArn(val *string) {
	_jsii_.Set(
		j,
		"flowLogDestinationArn",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetFlowLogDestinationType(val *string) {
	_jsii_.Set(
		j,
		"flowLogDestinationType",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetFlowLogFileFormat(val *string) {
	_jsii_.Set(
		j,
		"flowLogFileFormat",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetFlowLogHiveCompatiblePartitions(val *bool) {
	_jsii_.Set(
		j,
		"flowLogHiveCompatiblePartitions",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetFlowLogLogFormat(val *string) {
	_jsii_.Set(
		j,
		"flowLogLogFormat",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetFlowLogMaxAggregationInterval(val *float64) {
	_jsii_.Set(
		j,
		"flowLogMaxAggregationInterval",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetFlowLogPerHourPartition(val *bool) {
	_jsii_.Set(
		j,
		"flowLogPerHourPartition",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetFlowLogTrafficType(val *string) {
	_jsii_.Set(
		j,
		"flowLogTrafficType",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIgwTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"igwTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetInstanceTenancy(val *string) {
	_jsii_.Set(
		j,
		"instanceTenancy",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIntraAclTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"intraAclTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIntraDedicatedNetworkAcl(val *bool) {
	_jsii_.Set(
		j,
		"intraDedicatedNetworkAcl",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIntraInboundAclRules(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"intraInboundAclRules",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIntraOutboundAclRules(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"intraOutboundAclRules",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIntraRouteTableTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"intraRouteTableTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIntraSubnetAssignIpv6AddressOnCreation(val *bool) {
	_jsii_.Set(
		j,
		"intraSubnetAssignIpv6AddressOnCreation",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIntraSubnetIpv6Prefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"intraSubnetIpv6Prefixes",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIntraSubnetNames(val *[]*string) {
	_jsii_.Set(
		j,
		"intraSubnetNames",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIntraSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"intraSubnets",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIntraSubnetSuffix(val *string) {
	_jsii_.Set(
		j,
		"intraSubnetSuffix",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIntraSubnetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"intraSubnetTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIpv4IpamPoolId(val *string) {
	_jsii_.Set(
		j,
		"ipv4IpamPoolId",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIpv4NetmaskLength(val *float64) {
	_jsii_.Set(
		j,
		"ipv4NetmaskLength",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIpv6Cidr(val *string) {
	_jsii_.Set(
		j,
		"ipv6Cidr",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIpv6IpamPoolId(val *string) {
	_jsii_.Set(
		j,
		"ipv6IpamPoolId",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetIpv6NetmaskLength(val *float64) {
	_jsii_.Set(
		j,
		"ipv6NetmaskLength",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetManageDefaultNetworkAcl(val *bool) {
	_jsii_.Set(
		j,
		"manageDefaultNetworkAcl",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetManageDefaultRouteTable(val *bool) {
	_jsii_.Set(
		j,
		"manageDefaultRouteTable",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetManageDefaultSecurityGroup(val *bool) {
	_jsii_.Set(
		j,
		"manageDefaultSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetManageDefaultVpc(val *bool) {
	_jsii_.Set(
		j,
		"manageDefaultVpc",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetMapPublicIpOnLaunch(val *bool) {
	_jsii_.Set(
		j,
		"mapPublicIpOnLaunch",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetNatEipTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"natEipTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetNatGatewayDestinationCidrBlock(val *string) {
	_jsii_.Set(
		j,
		"natGatewayDestinationCidrBlock",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetNatGatewayTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"natGatewayTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetOneNatGatewayPerAz(val *bool) {
	_jsii_.Set(
		j,
		"oneNatGatewayPerAz",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetOutpostAclTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"outpostAclTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetOutpostArn(val *string) {
	_jsii_.Set(
		j,
		"outpostArn",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetOutpostAz(val *string) {
	_jsii_.Set(
		j,
		"outpostAz",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetOutpostDedicatedNetworkAcl(val *bool) {
	_jsii_.Set(
		j,
		"outpostDedicatedNetworkAcl",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetOutpostInboundAclRules(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"outpostInboundAclRules",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetOutpostOutboundAclRules(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"outpostOutboundAclRules",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetOutpostSubnetAssignIpv6AddressOnCreation(val *bool) {
	_jsii_.Set(
		j,
		"outpostSubnetAssignIpv6AddressOnCreation",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetOutpostSubnetIpv6Prefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"outpostSubnetIpv6Prefixes",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetOutpostSubnetNames(val *[]*string) {
	_jsii_.Set(
		j,
		"outpostSubnetNames",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetOutpostSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"outpostSubnets",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetOutpostSubnetSuffix(val *string) {
	_jsii_.Set(
		j,
		"outpostSubnetSuffix",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetOutpostSubnetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"outpostSubnetTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPrivateAclTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"privateAclTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPrivateDedicatedNetworkAcl(val *bool) {
	_jsii_.Set(
		j,
		"privateDedicatedNetworkAcl",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPrivateInboundAclRules(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"privateInboundAclRules",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPrivateOutboundAclRules(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"privateOutboundAclRules",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPrivateRouteTableTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"privateRouteTableTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPrivateSubnetAssignIpv6AddressOnCreation(val *bool) {
	_jsii_.Set(
		j,
		"privateSubnetAssignIpv6AddressOnCreation",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPrivateSubnetIpv6Prefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"privateSubnetIpv6Prefixes",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPrivateSubnetNames(val *[]*string) {
	_jsii_.Set(
		j,
		"privateSubnetNames",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPrivateSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"privateSubnets",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPrivateSubnetSuffix(val *string) {
	_jsii_.Set(
		j,
		"privateSubnetSuffix",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPrivateSubnetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"privateSubnetTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPrivateSubnetTagsPerAz(val *map[string]*map[string]*string) {
	_jsii_.Set(
		j,
		"privateSubnetTagsPerAz",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPropagateIntraRouteTablesVgw(val *bool) {
	_jsii_.Set(
		j,
		"propagateIntraRouteTablesVgw",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPropagatePrivateRouteTablesVgw(val *bool) {
	_jsii_.Set(
		j,
		"propagatePrivateRouteTablesVgw",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPropagatePublicRouteTablesVgw(val *bool) {
	_jsii_.Set(
		j,
		"propagatePublicRouteTablesVgw",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPublicAclTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"publicAclTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPublicDedicatedNetworkAcl(val *bool) {
	_jsii_.Set(
		j,
		"publicDedicatedNetworkAcl",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPublicInboundAclRules(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"publicInboundAclRules",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPublicOutboundAclRules(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"publicOutboundAclRules",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPublicRouteTableTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"publicRouteTableTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPublicSubnetAssignIpv6AddressOnCreation(val *bool) {
	_jsii_.Set(
		j,
		"publicSubnetAssignIpv6AddressOnCreation",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPublicSubnetIpv6Prefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"publicSubnetIpv6Prefixes",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPublicSubnetNames(val *[]*string) {
	_jsii_.Set(
		j,
		"publicSubnetNames",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPublicSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"publicSubnets",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPublicSubnetSuffix(val *string) {
	_jsii_.Set(
		j,
		"publicSubnetSuffix",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPublicSubnetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"publicSubnetTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPublicSubnetTagsPerAz(val *map[string]*map[string]*string) {
	_jsii_.Set(
		j,
		"publicSubnetTagsPerAz",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetPutinKhuylo(val *bool) {
	_jsii_.Set(
		j,
		"putinKhuylo",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetRedshiftAclTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"redshiftAclTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetRedshiftDedicatedNetworkAcl(val *bool) {
	_jsii_.Set(
		j,
		"redshiftDedicatedNetworkAcl",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetRedshiftInboundAclRules(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"redshiftInboundAclRules",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetRedshiftOutboundAclRules(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"redshiftOutboundAclRules",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetRedshiftRouteTableTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"redshiftRouteTableTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetRedshiftSubnetAssignIpv6AddressOnCreation(val *bool) {
	_jsii_.Set(
		j,
		"redshiftSubnetAssignIpv6AddressOnCreation",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetRedshiftSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"redshiftSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetRedshiftSubnetGroupTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"redshiftSubnetGroupTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetRedshiftSubnetIpv6Prefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"redshiftSubnetIpv6Prefixes",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetRedshiftSubnetNames(val *[]*string) {
	_jsii_.Set(
		j,
		"redshiftSubnetNames",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetRedshiftSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"redshiftSubnets",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetRedshiftSubnetSuffix(val *string) {
	_jsii_.Set(
		j,
		"redshiftSubnetSuffix",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetRedshiftSubnetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"redshiftSubnetTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetReuseNatIps(val *bool) {
	_jsii_.Set(
		j,
		"reuseNatIps",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetSecondaryCidrBlocks(val *[]*string) {
	_jsii_.Set(
		j,
		"secondaryCidrBlocks",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetSingleNatGateway(val *bool) {
	_jsii_.Set(
		j,
		"singleNatGateway",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetUseIpamPool(val *bool) {
	_jsii_.Set(
		j,
		"useIpamPool",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetVpcFlowLogPermissionsBoundary(val *string) {
	_jsii_.Set(
		j,
		"vpcFlowLogPermissionsBoundary",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetVpcFlowLogTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"vpcFlowLogTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetVpcTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"vpcTags",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetVpnGatewayAz(val *string) {
	_jsii_.Set(
		j,
		"vpnGatewayAz",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetVpnGatewayId(val *string) {
	_jsii_.Set(
		j,
		"vpnGatewayId",
		val,
	)
}

func (j *jsiiProxy_Awsvpc)SetVpnGatewayTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"vpnGatewayTags",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Awsvpc_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsvpc_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-awsvpc.Awsvpc",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Awsvpc_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsvpc_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-awsvpc.Awsvpc",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Awsvpc) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_Awsvpc) AddProvider(provider interface{}) {
	if err := a.validateAddProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addProvider",
		[]interface{}{provider},
	)
}

func (a *jsiiProxy_Awsvpc) GetString(output *string) *string {
	if err := a.validateGetStringParameters(output); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Awsvpc) InterpolationForOutput(moduleOutput *string) cdktf.IResolvable {
	if err := a.validateInterpolationForOutputParameters(moduleOutput); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Awsvpc) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Awsvpc) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Awsvpc) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Awsvpc) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Awsvpc) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Awsvpc) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

