package googlecomputenetworkpeering

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputenetworkpeering/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_network_peering google_compute_network_peering}.
type GoogleComputeNetworkPeering interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExportCustomRoutes() interface{}
	SetExportCustomRoutes(val interface{})
	ExportCustomRoutesInput() interface{}
	ExportSubnetRoutesWithPublicIp() interface{}
	SetExportSubnetRoutesWithPublicIp(val interface{})
	ExportSubnetRoutesWithPublicIpInput() interface{}
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
	ImportCustomRoutes() interface{}
	SetImportCustomRoutes(val interface{})
	ImportCustomRoutesInput() interface{}
	ImportSubnetRoutesWithPublicIp() interface{}
	SetImportSubnetRoutesWithPublicIp(val interface{})
	ImportSubnetRoutesWithPublicIpInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	// The tree node.
	Node() constructs.Node
	PeerNetwork() *string
	SetPeerNetwork(val *string)
	PeerNetworkInput() *string
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
	State() *string
	StateDetails() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeNetworkPeeringTimeoutsOutputReference
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
	PutTimeouts(value *GoogleComputeNetworkPeeringTimeouts)
	ResetExportCustomRoutes()
	ResetExportSubnetRoutesWithPublicIp()
	ResetId()
	ResetImportCustomRoutes()
	ResetImportSubnetRoutesWithPublicIp()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for GoogleComputeNetworkPeering
type jsiiProxy_GoogleComputeNetworkPeering struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) ExportCustomRoutes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportCustomRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) ExportCustomRoutesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportCustomRoutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) ExportSubnetRoutesWithPublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportSubnetRoutesWithPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) ExportSubnetRoutesWithPublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportSubnetRoutesWithPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) ImportCustomRoutes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importCustomRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) ImportCustomRoutesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importCustomRoutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) ImportSubnetRoutesWithPublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importSubnetRoutesWithPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) ImportSubnetRoutesWithPublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importSubnetRoutesWithPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) PeerNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) PeerNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) StateDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) Timeouts() GoogleComputeNetworkPeeringTimeoutsOutputReference {
	var returns GoogleComputeNetworkPeeringTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkPeering) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_network_peering google_compute_network_peering} Resource.
func NewGoogleComputeNetworkPeering(scope constructs.Construct, id *string, config *GoogleComputeNetworkPeeringConfig) GoogleComputeNetworkPeering {
	_init_.Initialize()

	if err := validateNewGoogleComputeNetworkPeeringParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeNetworkPeering{}

	_jsii_.Create(
		"google-beta.googleComputeNetworkPeering.GoogleComputeNetworkPeering",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_network_peering google_compute_network_peering} Resource.
func NewGoogleComputeNetworkPeering_Override(g GoogleComputeNetworkPeering, scope constructs.Construct, id *string, config *GoogleComputeNetworkPeeringConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeNetworkPeering.GoogleComputeNetworkPeering",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetExportCustomRoutes(val interface{}) {
	if err := j.validateSetExportCustomRoutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportCustomRoutes",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetExportSubnetRoutesWithPublicIp(val interface{}) {
	if err := j.validateSetExportSubnetRoutesWithPublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportSubnetRoutesWithPublicIp",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetImportCustomRoutes(val interface{}) {
	if err := j.validateSetImportCustomRoutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importCustomRoutes",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetImportSubnetRoutesWithPublicIp(val interface{}) {
	if err := j.validateSetImportSubnetRoutesWithPublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importSubnetRoutesWithPublicIp",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetPeerNetwork(val *string) {
	if err := j.validateSetPeerNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerNetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkPeering)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func GoogleComputeNetworkPeering_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeNetworkPeering_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeNetworkPeering.GoogleComputeNetworkPeering",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeNetworkPeering_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleComputeNetworkPeering.GoogleComputeNetworkPeering",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkPeering) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeNetworkPeering) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeNetworkPeering) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeNetworkPeering) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeNetworkPeering) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeNetworkPeering) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeNetworkPeering) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeNetworkPeering) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeNetworkPeering) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeNetworkPeering) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeNetworkPeering) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeNetworkPeering) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeNetworkPeering) PutTimeouts(value *GoogleComputeNetworkPeeringTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeNetworkPeering) ResetExportCustomRoutes() {
	_jsii_.InvokeVoid(
		g,
		"resetExportCustomRoutes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkPeering) ResetExportSubnetRoutesWithPublicIp() {
	_jsii_.InvokeVoid(
		g,
		"resetExportSubnetRoutesWithPublicIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkPeering) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkPeering) ResetImportCustomRoutes() {
	_jsii_.InvokeVoid(
		g,
		"resetImportCustomRoutes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkPeering) ResetImportSubnetRoutesWithPublicIp() {
	_jsii_.InvokeVoid(
		g,
		"resetImportSubnetRoutesWithPublicIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkPeering) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkPeering) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkPeering) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkPeering) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkPeering) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkPeering) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

