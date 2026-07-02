package ec2transitgatewaymulticastdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/ec2transitgatewaymulticastdomain/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ec2_transit_gateway_multicast_domain aws_ec2_transit_gateway_multicast_domain}.
type Ec2TransitGatewayMulticastDomain interface {
	cdktf.TerraformResource
	Arn() *string
	AutoAcceptSharedAssociations() *string
	SetAutoAcceptSharedAssociations(val *string)
	AutoAcceptSharedAssociationsInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	Igmpv2Support() *string
	SetIgmpv2Support(val *string)
	Igmpv2SupportInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	StaticSourcesSupport() *string
	SetStaticSourcesSupport(val *string)
	StaticSourcesSupportInput() *string
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
	Timeouts() Ec2TransitGatewayMulticastDomainTimeoutsOutputReference
	TimeoutsInput() any
	TransitGatewayId() *string
	SetTransitGatewayId(val *string)
	TransitGatewayIdInput() *string
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
	PutTimeouts(value *Ec2TransitGatewayMulticastDomainTimeouts)
	ResetAutoAcceptSharedAssociations()
	ResetId()
	ResetIgmpv2Support()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStaticSourcesSupport()
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

// The jsii proxy struct for Ec2TransitGatewayMulticastDomain
type jsiiProxy_Ec2TransitGatewayMulticastDomain struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) AutoAcceptSharedAssociations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoAcceptSharedAssociations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) AutoAcceptSharedAssociationsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoAcceptSharedAssociationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) Igmpv2Support() *string {
	var returns *string
	_jsii_.Get(
		j,
		"igmpv2Support",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) Igmpv2SupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"igmpv2SupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) StaticSourcesSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticSourcesSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) StaticSourcesSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticSourcesSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) Timeouts() Ec2TransitGatewayMulticastDomainTimeoutsOutputReference {
	var returns Ec2TransitGatewayMulticastDomainTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) TransitGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ec2_transit_gateway_multicast_domain aws_ec2_transit_gateway_multicast_domain} Resource.
func NewEc2TransitGatewayMulticastDomain(scope constructs.Construct, id *string, config *Ec2TransitGatewayMulticastDomainConfig) Ec2TransitGatewayMulticastDomain {
	_init_.Initialize()

	if err := validateNewEc2TransitGatewayMulticastDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2TransitGatewayMulticastDomain{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2TransitGatewayMulticastDomain.Ec2TransitGatewayMulticastDomain",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ec2_transit_gateway_multicast_domain aws_ec2_transit_gateway_multicast_domain} Resource.
func NewEc2TransitGatewayMulticastDomain_Override(e Ec2TransitGatewayMulticastDomain, scope constructs.Construct, id *string, config *Ec2TransitGatewayMulticastDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2TransitGatewayMulticastDomain.Ec2TransitGatewayMulticastDomain",
		[]any{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) SetAutoAcceptSharedAssociations(val *string) {
	if err := j.validateSetAutoAcceptSharedAssociationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAcceptSharedAssociations",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) SetIgmpv2Support(val *string) {
	if err := j.validateSetIgmpv2SupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"igmpv2Support",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) SetStaticSourcesSupport(val *string) {
	if err := j.validateSetStaticSourcesSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticSourcesSupport",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMulticastDomain) SetTransitGatewayId(val *string) {
	if err := j.validateSetTransitGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayId",
		val,
	)
}

// Generates CDKTF code for importing a Ec2TransitGatewayMulticastDomain resource upon running "cdktf plan <stack-name>".
func Ec2TransitGatewayMulticastDomain_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEc2TransitGatewayMulticastDomain_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2TransitGatewayMulticastDomain.Ec2TransitGatewayMulticastDomain",
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
func Ec2TransitGatewayMulticastDomain_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateEc2TransitGatewayMulticastDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2TransitGatewayMulticastDomain.Ec2TransitGatewayMulticastDomain",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2TransitGatewayMulticastDomain_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateEc2TransitGatewayMulticastDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2TransitGatewayMulticastDomain.Ec2TransitGatewayMulticastDomain",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2TransitGatewayMulticastDomain_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateEc2TransitGatewayMulticastDomain_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2TransitGatewayMulticastDomain.Ec2TransitGatewayMulticastDomain",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func Ec2TransitGatewayMulticastDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.ec2TransitGatewayMulticastDomain.Ec2TransitGatewayMulticastDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) AddOverride(path *string, value any) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]any{path, value},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]any{id, provider},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]any{id},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) MoveTo(moveTarget *string, index any) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]any{id},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) PutTimeouts(value *Ec2TransitGatewayMulticastDomainTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]any{value},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) ResetAutoAcceptSharedAssociations() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoAcceptSharedAssociations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) ResetIgmpv2Support() {
	_jsii_.InvokeVoid(
		e,
		"resetIgmpv2Support",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) ResetStaticSourcesSupport() {
	_jsii_.InvokeVoid(
		e,
		"resetStaticSourcesSupport",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMulticastDomain) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
