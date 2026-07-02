package networkmanagerattachmentaccepter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/networkmanagerattachmentaccepter/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/networkmanager_attachment_accepter aws_networkmanager_attachment_accepter}.
type NetworkmanagerAttachmentAccepter interface {
	cdktf.TerraformResource
	AttachmentId() *string
	SetAttachmentId(val *string)
	AttachmentIdInput() *string
	AttachmentPolicyRuleNumber() *float64
	AttachmentType() *string
	SetAttachmentType(val *string)
	AttachmentTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	CoreNetworkArn() *string
	CoreNetworkId() *string
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EdgeLocation() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OwnerAccountId() *string
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
	ResourceArn() *string
	SegmentName() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NetworkmanagerAttachmentAccepterTimeoutsOutputReference
	TimeoutsInput() any
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
	PutTimeouts(value *NetworkmanagerAttachmentAccepterTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for NetworkmanagerAttachmentAccepter
type jsiiProxy_NetworkmanagerAttachmentAccepter struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) AttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) AttachmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) AttachmentPolicyRuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attachmentPolicyRuleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) AttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) AttachmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) CoreNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) CoreNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) EdgeLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) OwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) SegmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) Timeouts() NetworkmanagerAttachmentAccepterTimeoutsOutputReference {
	var returns NetworkmanagerAttachmentAccepterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/networkmanager_attachment_accepter aws_networkmanager_attachment_accepter} Resource.
func NewNetworkmanagerAttachmentAccepter(scope constructs.Construct, id *string, config *NetworkmanagerAttachmentAccepterConfig) NetworkmanagerAttachmentAccepter {
	_init_.Initialize()

	if err := validateNewNetworkmanagerAttachmentAccepterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkmanagerAttachmentAccepter{}

	_jsii_.Create(
		"@cdktf/provider-aws.networkmanagerAttachmentAccepter.NetworkmanagerAttachmentAccepter",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/networkmanager_attachment_accepter aws_networkmanager_attachment_accepter} Resource.
func NewNetworkmanagerAttachmentAccepter_Override(n NetworkmanagerAttachmentAccepter, scope constructs.Construct, id *string, config *NetworkmanagerAttachmentAccepterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.networkmanagerAttachmentAccepter.NetworkmanagerAttachmentAccepter",
		[]any{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) SetAttachmentId(val *string) {
	if err := j.validateSetAttachmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentId",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) SetAttachmentType(val *string) {
	if err := j.validateSetAttachmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentType",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkmanagerAttachmentAccepter) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a NetworkmanagerAttachmentAccepter resource upon running "cdktf plan <stack-name>".
func NetworkmanagerAttachmentAccepter_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkmanagerAttachmentAccepter_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.networkmanagerAttachmentAccepter.NetworkmanagerAttachmentAccepter",
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
func NetworkmanagerAttachmentAccepter_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateNetworkmanagerAttachmentAccepter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.networkmanagerAttachmentAccepter.NetworkmanagerAttachmentAccepter",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkmanagerAttachmentAccepter_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateNetworkmanagerAttachmentAccepter_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.networkmanagerAttachmentAccepter.NetworkmanagerAttachmentAccepter",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkmanagerAttachmentAccepter_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateNetworkmanagerAttachmentAccepter_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.networkmanagerAttachmentAccepter.NetworkmanagerAttachmentAccepter",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func NetworkmanagerAttachmentAccepter_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.networkmanagerAttachmentAccepter.NetworkmanagerAttachmentAccepter",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) AddOverride(path *string, value any) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]any{path, value},
	)
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]any{id, provider},
	)
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]any{id},
	)
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) MoveTo(moveTarget *string, index any) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]any{id},
	)
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) PutTimeouts(value *NetworkmanagerAttachmentAccepterTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]any{value},
	)
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkmanagerAttachmentAccepter) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
