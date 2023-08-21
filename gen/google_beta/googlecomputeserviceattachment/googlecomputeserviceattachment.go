package googlecomputeserviceattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeserviceattachment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_compute_service_attachment google_compute_service_attachment}.
type GoogleComputeServiceAttachment interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConnectedEndpoints() GoogleComputeServiceAttachmentConnectedEndpointsList
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionPreference() *string
	SetConnectionPreference(val *string)
	ConnectionPreferenceInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConsumerAcceptLists() GoogleComputeServiceAttachmentConsumerAcceptListsList
	ConsumerAcceptListsInput() interface{}
	ConsumerRejectLists() *[]*string
	SetConsumerRejectLists(val *[]*string)
	ConsumerRejectListsInput() *[]*string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DomainNames() *[]*string
	SetDomainNames(val *[]*string)
	DomainNamesInput() *[]*string
	EnableProxyProtocol() interface{}
	SetEnableProxyProtocol(val interface{})
	EnableProxyProtocolInput() interface{}
	Fingerprint() *string
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	NatSubnets() *[]*string
	SetNatSubnets(val *[]*string)
	NatSubnetsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReconcileConnections() interface{}
	SetReconcileConnections(val interface{})
	ReconcileConnectionsInput() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SelfLink() *string
	TargetService() *string
	SetTargetService(val *string)
	TargetServiceInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeServiceAttachmentTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutConsumerAcceptLists(value interface{})
	PutTimeouts(value *GoogleComputeServiceAttachmentTimeouts)
	ResetConsumerAcceptLists()
	ResetConsumerRejectLists()
	ResetDescription()
	ResetDomainNames()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetReconcileConnections()
	ResetRegion()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleComputeServiceAttachment
type jsiiProxy_GoogleComputeServiceAttachment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) ConnectedEndpoints() GoogleComputeServiceAttachmentConnectedEndpointsList {
	var returns GoogleComputeServiceAttachmentConnectedEndpointsList
	_jsii_.Get(
		j,
		"connectedEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) ConnectionPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) ConnectionPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) ConsumerAcceptLists() GoogleComputeServiceAttachmentConsumerAcceptListsList {
	var returns GoogleComputeServiceAttachmentConsumerAcceptListsList
	_jsii_.Get(
		j,
		"consumerAcceptLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) ConsumerAcceptListsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consumerAcceptListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) ConsumerRejectLists() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"consumerRejectLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) ConsumerRejectListsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"consumerRejectListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) DomainNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) DomainNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) EnableProxyProtocol() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableProxyProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) EnableProxyProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableProxyProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) NatSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"natSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) NatSubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"natSubnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) ReconcileConnections() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reconcileConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) ReconcileConnectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reconcileConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) TargetService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) TargetServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) Timeouts() GoogleComputeServiceAttachmentTimeoutsOutputReference {
	var returns GoogleComputeServiceAttachmentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeServiceAttachment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_compute_service_attachment google_compute_service_attachment} Resource.
func NewGoogleComputeServiceAttachment(scope constructs.Construct, id *string, config *GoogleComputeServiceAttachmentConfig) GoogleComputeServiceAttachment {
	_init_.Initialize()

	if err := validateNewGoogleComputeServiceAttachmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeServiceAttachment{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeServiceAttachment.GoogleComputeServiceAttachment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_compute_service_attachment google_compute_service_attachment} Resource.
func NewGoogleComputeServiceAttachment_Override(g GoogleComputeServiceAttachment, scope constructs.Construct, id *string, config *GoogleComputeServiceAttachmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeServiceAttachment.GoogleComputeServiceAttachment",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetConnectionPreference(val *string) {
	if err := j.validateSetConnectionPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionPreference",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetConsumerRejectLists(val *[]*string) {
	if err := j.validateSetConsumerRejectListsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerRejectLists",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetDomainNames(val *[]*string) {
	if err := j.validateSetDomainNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainNames",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetEnableProxyProtocol(val interface{}) {
	if err := j.validateSetEnableProxyProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableProxyProtocol",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetNatSubnets(val *[]*string) {
	if err := j.validateSetNatSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natSubnets",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetReconcileConnections(val interface{}) {
	if err := j.validateSetReconcileConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reconcileConnections",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeServiceAttachment)SetTargetService(val *string) {
	if err := j.validateSetTargetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetService",
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
func GoogleComputeServiceAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeServiceAttachment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeServiceAttachment.GoogleComputeServiceAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeServiceAttachment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeServiceAttachment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeServiceAttachment.GoogleComputeServiceAttachment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeServiceAttachment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeServiceAttachment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeServiceAttachment.GoogleComputeServiceAttachment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeServiceAttachment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleComputeServiceAttachment.GoogleComputeServiceAttachment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) PutConsumerAcceptLists(value interface{}) {
	if err := g.validatePutConsumerAcceptListsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConsumerAcceptLists",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) PutTimeouts(value *GoogleComputeServiceAttachmentTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) ResetConsumerAcceptLists() {
	_jsii_.InvokeVoid(
		g,
		"resetConsumerAcceptLists",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) ResetConsumerRejectLists() {
	_jsii_.InvokeVoid(
		g,
		"resetConsumerRejectLists",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) ResetDomainNames() {
	_jsii_.InvokeVoid(
		g,
		"resetDomainNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) ResetReconcileConnections() {
	_jsii_.InvokeVoid(
		g,
		"resetReconcileConnections",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeServiceAttachment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

