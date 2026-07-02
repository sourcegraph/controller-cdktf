package defaultsubnet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/defaultsubnet/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/default_subnet aws_default_subnet}.
type DefaultSubnet interface {
	cdktf.TerraformResource
	Arn() *string
	AssignIpv6AddressOnCreation() any
	SetAssignIpv6AddressOnCreation(val any)
	AssignIpv6AddressOnCreationInput() any
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneId() *string
	AvailabilityZoneInput() *string
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
	CustomerOwnedIpv4Pool() *string
	SetCustomerOwnedIpv4Pool(val *string)
	CustomerOwnedIpv4PoolInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableDns64() any
	SetEnableDns64(val any)
	EnableDns64Input() any
	EnableResourceNameDnsAaaaRecordOnLaunch() any
	SetEnableResourceNameDnsAaaaRecordOnLaunch(val any)
	EnableResourceNameDnsAaaaRecordOnLaunchInput() any
	EnableResourceNameDnsARecordOnLaunch() any
	SetEnableResourceNameDnsARecordOnLaunch(val any)
	EnableResourceNameDnsARecordOnLaunchInput() any
	ExistingDefaultSubnet() cdktf.IResolvable
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
	Ipv6CidrBlock() *string
	SetIpv6CidrBlock(val *string)
	Ipv6CidrBlockAssociationId() *string
	Ipv6CidrBlockInput() *string
	Ipv6Native() any
	SetIpv6Native(val any)
	Ipv6NativeInput() any
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MapCustomerOwnedIpOnLaunch() any
	SetMapCustomerOwnedIpOnLaunch(val any)
	MapCustomerOwnedIpOnLaunchInput() any
	MapPublicIpOnLaunch() any
	SetMapPublicIpOnLaunch(val any)
	MapPublicIpOnLaunchInput() any
	// The tree node.
	Node() constructs.Node
	OutpostArn() *string
	OwnerId() *string
	PrivateDnsHostnameTypeOnLaunch() *string
	SetPrivateDnsHostnameTypeOnLaunch(val *string)
	PrivateDnsHostnameTypeOnLaunchInput() *string
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
	Timeouts() DefaultSubnetTimeoutsOutputReference
	TimeoutsInput() any
	VpcId() *string
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
	PutTimeouts(value *DefaultSubnetTimeouts)
	ResetAssignIpv6AddressOnCreation()
	ResetCustomerOwnedIpv4Pool()
	ResetEnableDns64()
	ResetEnableResourceNameDnsAaaaRecordOnLaunch()
	ResetEnableResourceNameDnsARecordOnLaunch()
	ResetForceDestroy()
	ResetId()
	ResetIpv6CidrBlock()
	ResetIpv6Native()
	ResetMapCustomerOwnedIpOnLaunch()
	ResetMapPublicIpOnLaunch()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateDnsHostnameTypeOnLaunch()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
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

// The jsii proxy struct for DefaultSubnet
type jsiiProxy_DefaultSubnet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DefaultSubnet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) AssignIpv6AddressOnCreation() any {
	var returns any
	_jsii_.Get(
		j,
		"assignIpv6AddressOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) AssignIpv6AddressOnCreationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"assignIpv6AddressOnCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) CustomerOwnedIpv4Pool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerOwnedIpv4Pool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) CustomerOwnedIpv4PoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerOwnedIpv4PoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) EnableDns64() any {
	var returns any
	_jsii_.Get(
		j,
		"enableDns64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) EnableDns64Input() any {
	var returns any
	_jsii_.Get(
		j,
		"enableDns64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) EnableResourceNameDnsAaaaRecordOnLaunch() any {
	var returns any
	_jsii_.Get(
		j,
		"enableResourceNameDnsAaaaRecordOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) EnableResourceNameDnsAaaaRecordOnLaunchInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableResourceNameDnsAaaaRecordOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) EnableResourceNameDnsARecordOnLaunch() any {
	var returns any
	_jsii_.Get(
		j,
		"enableResourceNameDnsARecordOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) EnableResourceNameDnsARecordOnLaunchInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableResourceNameDnsARecordOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) ExistingDefaultSubnet() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"existingDefaultSubnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) ForceDestroy() any {
	var returns any
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) ForceDestroyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Ipv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Ipv6CidrBlockAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Ipv6CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Ipv6Native() any {
	var returns any
	_jsii_.Get(
		j,
		"ipv6Native",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Ipv6NativeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ipv6NativeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) MapCustomerOwnedIpOnLaunch() any {
	var returns any
	_jsii_.Get(
		j,
		"mapCustomerOwnedIpOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) MapCustomerOwnedIpOnLaunchInput() any {
	var returns any
	_jsii_.Get(
		j,
		"mapCustomerOwnedIpOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) MapPublicIpOnLaunch() any {
	var returns any
	_jsii_.Get(
		j,
		"mapPublicIpOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) MapPublicIpOnLaunchInput() any {
	var returns any
	_jsii_.Get(
		j,
		"mapPublicIpOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) OutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) PrivateDnsHostnameTypeOnLaunch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsHostnameTypeOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) PrivateDnsHostnameTypeOnLaunchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsHostnameTypeOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Timeouts() DefaultSubnetTimeoutsOutputReference {
	var returns DefaultSubnetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/default_subnet aws_default_subnet} Resource.
func NewDefaultSubnet(scope constructs.Construct, id *string, config *DefaultSubnetConfig) DefaultSubnet {
	_init_.Initialize()

	if err := validateNewDefaultSubnetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DefaultSubnet{}

	_jsii_.Create(
		"@cdktf/provider-aws.defaultSubnet.DefaultSubnet",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/default_subnet aws_default_subnet} Resource.
func NewDefaultSubnet_Override(d DefaultSubnet, scope constructs.Construct, id *string, config *DefaultSubnetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.defaultSubnet.DefaultSubnet",
		[]any{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetAssignIpv6AddressOnCreation(val any) {
	if err := j.validateSetAssignIpv6AddressOnCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assignIpv6AddressOnCreation",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetCustomerOwnedIpv4Pool(val *string) {
	if err := j.validateSetCustomerOwnedIpv4PoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerOwnedIpv4Pool",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetEnableDns64(val any) {
	if err := j.validateSetEnableDns64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDns64",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetEnableResourceNameDnsAaaaRecordOnLaunch(val any) {
	if err := j.validateSetEnableResourceNameDnsAaaaRecordOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceNameDnsAaaaRecordOnLaunch",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetEnableResourceNameDnsARecordOnLaunch(val any) {
	if err := j.validateSetEnableResourceNameDnsARecordOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceNameDnsARecordOnLaunch",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetForceDestroy(val any) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetIpv6CidrBlock(val *string) {
	if err := j.validateSetIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetIpv6Native(val any) {
	if err := j.validateSetIpv6NativeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Native",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetMapCustomerOwnedIpOnLaunch(val any) {
	if err := j.validateSetMapCustomerOwnedIpOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapCustomerOwnedIpOnLaunch",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetMapPublicIpOnLaunch(val any) {
	if err := j.validateSetMapPublicIpOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapPublicIpOnLaunch",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetPrivateDnsHostnameTypeOnLaunch(val *string) {
	if err := j.validateSetPrivateDnsHostnameTypeOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsHostnameTypeOnLaunch",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet) SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a DefaultSubnet resource upon running "cdktf plan <stack-name>".
func DefaultSubnet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDefaultSubnet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.defaultSubnet.DefaultSubnet",
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
func DefaultSubnet_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateDefaultSubnet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.defaultSubnet.DefaultSubnet",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DefaultSubnet_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateDefaultSubnet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.defaultSubnet.DefaultSubnet",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DefaultSubnet_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateDefaultSubnet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.defaultSubnet.DefaultSubnet",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func DefaultSubnet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.defaultSubnet.DefaultSubnet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DefaultSubnet) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (d *jsiiProxy_DefaultSubnet) AddOverride(path *string, value any) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]any{path, value},
	)
}

func (d *jsiiProxy_DefaultSubnet) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DefaultSubnet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DefaultSubnet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DefaultSubnet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DefaultSubnet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DefaultSubnet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DefaultSubnet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DefaultSubnet) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DefaultSubnet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DefaultSubnet) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultSubnet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]any{id, provider},
	)
}

func (d *jsiiProxy_DefaultSubnet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DefaultSubnet) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]any{id},
	)
}

func (d *jsiiProxy_DefaultSubnet) MoveTo(moveTarget *string, index any) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (d *jsiiProxy_DefaultSubnet) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]any{id},
	)
}

func (d *jsiiProxy_DefaultSubnet) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (d *jsiiProxy_DefaultSubnet) PutTimeouts(value *DefaultSubnetTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]any{value},
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetAssignIpv6AddressOnCreation() {
	_jsii_.InvokeVoid(
		d,
		"resetAssignIpv6AddressOnCreation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetCustomerOwnedIpv4Pool() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomerOwnedIpv4Pool",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetEnableDns64() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableDns64",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetEnableResourceNameDnsAaaaRecordOnLaunch() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableResourceNameDnsAaaaRecordOnLaunch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetEnableResourceNameDnsARecordOnLaunch() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableResourceNameDnsARecordOnLaunch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		d,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6CidrBlock",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetIpv6Native() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6Native",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetMapCustomerOwnedIpOnLaunch() {
	_jsii_.InvokeVoid(
		d,
		"resetMapCustomerOwnedIpOnLaunch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetMapPublicIpOnLaunch() {
	_jsii_.InvokeVoid(
		d,
		"resetMapPublicIpOnLaunch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetPrivateDnsHostnameTypeOnLaunch() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateDnsHostnameTypeOnLaunch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultSubnet) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultSubnet) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultSubnet) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultSubnet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultSubnet) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
