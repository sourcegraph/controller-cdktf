package defaultvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/defaultvpc/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/default_vpc aws_default_vpc}.
type DefaultVpc interface {
	cdktf.TerraformResource
	Arn() *string
	AssignGeneratedIpv6CidrBlock() any
	SetAssignGeneratedIpv6CidrBlock(val any)
	AssignGeneratedIpv6CidrBlockInput() any
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CidrBlock() *string
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	DefaultNetworkAclId() *string
	DefaultRouteTableId() *string
	DefaultSecurityGroupId() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DhcpOptionsId() *string
	EnableClassiclink() any
	SetEnableClassiclink(val any)
	EnableClassiclinkDnsSupport() any
	SetEnableClassiclinkDnsSupport(val any)
	EnableClassiclinkDnsSupportInput() any
	EnableClassiclinkInput() any
	EnableDnsHostnames() any
	SetEnableDnsHostnames(val any)
	EnableDnsHostnamesInput() any
	EnableDnsSupport() any
	SetEnableDnsSupport(val any)
	EnableDnsSupportInput() any
	EnableNetworkAddressUsageMetrics() any
	SetEnableNetworkAddressUsageMetrics(val any)
	EnableNetworkAddressUsageMetricsInput() any
	ExistingDefaultVpc() cdktf.IResolvable
	ForceDestroy() any
	SetForceDestroy(val any)
	ForceDestroyInput() any
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceTenancy() *string
	Ipv6AssociationId() *string
	Ipv6CidrBlock() *string
	SetIpv6CidrBlock(val *string)
	Ipv6CidrBlockInput() *string
	Ipv6CidrBlockNetworkBorderGroup() *string
	SetIpv6CidrBlockNetworkBorderGroup(val *string)
	Ipv6CidrBlockNetworkBorderGroupInput() *string
	Ipv6IpamPoolId() *string
	SetIpv6IpamPoolId(val *string)
	Ipv6IpamPoolIdInput() *string
	Ipv6NetmaskLength() *float64
	SetIpv6NetmaskLength(val *float64)
	Ipv6NetmaskLengthInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MainRouteTableId() *string
	// The tree node.
	Node() constructs.Node
	OwnerId() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() any
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAssignGeneratedIpv6CidrBlock()
	ResetEnableClassiclink()
	ResetEnableClassiclinkDnsSupport()
	ResetEnableDnsHostnames()
	ResetEnableDnsSupport()
	ResetEnableNetworkAddressUsageMetrics()
	ResetForceDestroy()
	ResetId()
	ResetIpv6CidrBlock()
	ResetIpv6CidrBlockNetworkBorderGroup()
	ResetIpv6IpamPoolId()
	ResetIpv6NetmaskLength()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for DefaultVpc
type jsiiProxy_DefaultVpc struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DefaultVpc) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) AssignGeneratedIpv6CidrBlock() any {
	var returns any
	_jsii_.Get(
		j,
		"assignGeneratedIpv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) AssignGeneratedIpv6CidrBlockInput() any {
	var returns any
	_jsii_.Get(
		j,
		"assignGeneratedIpv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) DefaultNetworkAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNetworkAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) DefaultRouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRouteTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) DefaultSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) DhcpOptionsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhcpOptionsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) EnableClassiclink() any {
	var returns any
	_jsii_.Get(
		j,
		"enableClassiclink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) EnableClassiclinkDnsSupport() any {
	var returns any
	_jsii_.Get(
		j,
		"enableClassiclinkDnsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) EnableClassiclinkDnsSupportInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableClassiclinkDnsSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) EnableClassiclinkInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableClassiclinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) EnableDnsHostnames() any {
	var returns any
	_jsii_.Get(
		j,
		"enableDnsHostnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) EnableDnsHostnamesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableDnsHostnamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) EnableDnsSupport() any {
	var returns any
	_jsii_.Get(
		j,
		"enableDnsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) EnableDnsSupportInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableDnsSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) EnableNetworkAddressUsageMetrics() any {
	var returns any
	_jsii_.Get(
		j,
		"enableNetworkAddressUsageMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) EnableNetworkAddressUsageMetricsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableNetworkAddressUsageMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) ExistingDefaultVpc() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"existingDefaultVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) ForceDestroy() any {
	var returns any
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) ForceDestroyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) InstanceTenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Ipv6AssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Ipv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Ipv6CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Ipv6CidrBlockNetworkBorderGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockNetworkBorderGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Ipv6CidrBlockNetworkBorderGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockNetworkBorderGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Ipv6IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6IpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Ipv6IpamPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6IpamPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Ipv6NetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6NetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Ipv6NetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6NetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) MainRouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainRouteTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultVpc) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/default_vpc aws_default_vpc} Resource.
func NewDefaultVpc(scope constructs.Construct, id *string, config *DefaultVpcConfig) DefaultVpc {
	_init_.Initialize()

	if err := validateNewDefaultVpcParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DefaultVpc{}

	_jsii_.Create(
		"@cdktf/provider-aws.defaultVpc.DefaultVpc",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/default_vpc aws_default_vpc} Resource.
func NewDefaultVpc_Override(d DefaultVpc, scope constructs.Construct, id *string, config *DefaultVpcConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.defaultVpc.DefaultVpc",
		[]any{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DefaultVpc) SetAssignGeneratedIpv6CidrBlock(val any) {
	if err := j.validateSetAssignGeneratedIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assignGeneratedIpv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetEnableClassiclink(val any) {
	if err := j.validateSetEnableClassiclinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableClassiclink",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetEnableClassiclinkDnsSupport(val any) {
	if err := j.validateSetEnableClassiclinkDnsSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableClassiclinkDnsSupport",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetEnableDnsHostnames(val any) {
	if err := j.validateSetEnableDnsHostnamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDnsHostnames",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetEnableDnsSupport(val any) {
	if err := j.validateSetEnableDnsSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDnsSupport",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetEnableNetworkAddressUsageMetrics(val any) {
	if err := j.validateSetEnableNetworkAddressUsageMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNetworkAddressUsageMetrics",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetForceDestroy(val any) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetIpv6CidrBlock(val *string) {
	if err := j.validateSetIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetIpv6CidrBlockNetworkBorderGroup(val *string) {
	if err := j.validateSetIpv6CidrBlockNetworkBorderGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlockNetworkBorderGroup",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetIpv6IpamPoolId(val *string) {
	if err := j.validateSetIpv6IpamPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6IpamPoolId",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetIpv6NetmaskLength(val *float64) {
	if err := j.validateSetIpv6NetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6NetmaskLength",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DefaultVpc) SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a DefaultVpc resource upon running "cdktf plan <stack-name>".
func DefaultVpc_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDefaultVpc_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.defaultVpc.DefaultVpc",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func DefaultVpc_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateDefaultVpc_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.defaultVpc.DefaultVpc",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DefaultVpc_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateDefaultVpc_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.defaultVpc.DefaultVpc",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DefaultVpc_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateDefaultVpc_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.defaultVpc.DefaultVpc",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func DefaultVpc_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.defaultVpc.DefaultVpc",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DefaultVpc) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (d *jsiiProxy_DefaultVpc) AddOverride(path *string, value any) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]any{path, value},
	)
}

func (d *jsiiProxy_DefaultVpc) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]any{id, provider},
	)
}

func (d *jsiiProxy_DefaultVpc) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]any{id},
	)
}

func (d *jsiiProxy_DefaultVpc) MoveTo(moveTarget *string, index any) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (d *jsiiProxy_DefaultVpc) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]any{id},
	)
}

func (d *jsiiProxy_DefaultVpc) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (d *jsiiProxy_DefaultVpc) ResetAssignGeneratedIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		d,
		"resetAssignGeneratedIpv6CidrBlock",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) ResetEnableClassiclink() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableClassiclink",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) ResetEnableClassiclinkDnsSupport() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableClassiclinkDnsSupport",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) ResetEnableDnsHostnames() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableDnsHostnames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) ResetEnableDnsSupport() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableDnsSupport",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) ResetEnableNetworkAddressUsageMetrics() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableNetworkAddressUsageMetrics",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		d,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) ResetIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6CidrBlock",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) ResetIpv6CidrBlockNetworkBorderGroup() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6CidrBlockNetworkBorderGroup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) ResetIpv6IpamPoolId() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6IpamPoolId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) ResetIpv6NetmaskLength() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6NetmaskLength",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultVpc) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultVpc) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
