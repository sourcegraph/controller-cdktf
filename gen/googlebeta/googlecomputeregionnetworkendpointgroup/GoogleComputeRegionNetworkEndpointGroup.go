package googlecomputeregionnetworkendpointgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecomputeregionnetworkendpointgroup/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_network_endpoint_group google_compute_region_network_endpoint_group}.
type GoogleComputeRegionNetworkEndpointGroup interface {
	cdktf.TerraformResource
	AppEngine() GoogleComputeRegionNetworkEndpointGroupAppEngineOutputReference
	AppEngineInput() *GoogleComputeRegionNetworkEndpointGroupAppEngine
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudFunction() GoogleComputeRegionNetworkEndpointGroupCloudFunctionOutputReference
	CloudFunctionInput() *GoogleComputeRegionNetworkEndpointGroupCloudFunction
	CloudRun() GoogleComputeRegionNetworkEndpointGroupCloudRunOutputReference
	CloudRunInput() *GoogleComputeRegionNetworkEndpointGroupCloudRun
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	Network() *string
	SetNetwork(val *string)
	NetworkEndpointType() *string
	SetNetworkEndpointType(val *string)
	NetworkEndpointTypeInput() *string
	NetworkInput() *string
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
	PscTargetService() *string
	SetPscTargetService(val *string)
	PscTargetServiceInput() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SelfLink() *string
	ServerlessDeployment() GoogleComputeRegionNetworkEndpointGroupServerlessDeploymentOutputReference
	ServerlessDeploymentInput() *GoogleComputeRegionNetworkEndpointGroupServerlessDeployment
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeRegionNetworkEndpointGroupTimeoutsOutputReference
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
	PutAppEngine(value *GoogleComputeRegionNetworkEndpointGroupAppEngine)
	PutCloudFunction(value *GoogleComputeRegionNetworkEndpointGroupCloudFunction)
	PutCloudRun(value *GoogleComputeRegionNetworkEndpointGroupCloudRun)
	PutServerlessDeployment(value *GoogleComputeRegionNetworkEndpointGroupServerlessDeployment)
	PutTimeouts(value *GoogleComputeRegionNetworkEndpointGroupTimeouts)
	ResetAppEngine()
	ResetCloudFunction()
	ResetCloudRun()
	ResetDescription()
	ResetId()
	ResetNetwork()
	ResetNetworkEndpointType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPscTargetService()
	ResetServerlessDeployment()
	ResetSubnetwork()
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

// The jsii proxy struct for GoogleComputeRegionNetworkEndpointGroup
type jsiiProxy_GoogleComputeRegionNetworkEndpointGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) AppEngine() GoogleComputeRegionNetworkEndpointGroupAppEngineOutputReference {
	var returns GoogleComputeRegionNetworkEndpointGroupAppEngineOutputReference
	_jsii_.Get(
		j,
		"appEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) AppEngineInput() *GoogleComputeRegionNetworkEndpointGroupAppEngine {
	var returns *GoogleComputeRegionNetworkEndpointGroupAppEngine
	_jsii_.Get(
		j,
		"appEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) CloudFunction() GoogleComputeRegionNetworkEndpointGroupCloudFunctionOutputReference {
	var returns GoogleComputeRegionNetworkEndpointGroupCloudFunctionOutputReference
	_jsii_.Get(
		j,
		"cloudFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) CloudFunctionInput() *GoogleComputeRegionNetworkEndpointGroupCloudFunction {
	var returns *GoogleComputeRegionNetworkEndpointGroupCloudFunction
	_jsii_.Get(
		j,
		"cloudFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) CloudRun() GoogleComputeRegionNetworkEndpointGroupCloudRunOutputReference {
	var returns GoogleComputeRegionNetworkEndpointGroupCloudRunOutputReference
	_jsii_.Get(
		j,
		"cloudRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) CloudRunInput() *GoogleComputeRegionNetworkEndpointGroupCloudRun {
	var returns *GoogleComputeRegionNetworkEndpointGroupCloudRun
	_jsii_.Get(
		j,
		"cloudRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) NetworkEndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkEndpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) NetworkEndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkEndpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) PscTargetService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pscTargetService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) PscTargetServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pscTargetServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ServerlessDeployment() GoogleComputeRegionNetworkEndpointGroupServerlessDeploymentOutputReference {
	var returns GoogleComputeRegionNetworkEndpointGroupServerlessDeploymentOutputReference
	_jsii_.Get(
		j,
		"serverlessDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ServerlessDeploymentInput() *GoogleComputeRegionNetworkEndpointGroupServerlessDeployment {
	var returns *GoogleComputeRegionNetworkEndpointGroupServerlessDeployment
	_jsii_.Get(
		j,
		"serverlessDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) Timeouts() GoogleComputeRegionNetworkEndpointGroupTimeoutsOutputReference {
	var returns GoogleComputeRegionNetworkEndpointGroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_network_endpoint_group google_compute_region_network_endpoint_group} Resource.
func NewGoogleComputeRegionNetworkEndpointGroup(scope constructs.Construct, id *string, config *GoogleComputeRegionNetworkEndpointGroupConfig) GoogleComputeRegionNetworkEndpointGroup {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionNetworkEndpointGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionNetworkEndpointGroup{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeRegionNetworkEndpointGroup.GoogleComputeRegionNetworkEndpointGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_network_endpoint_group google_compute_region_network_endpoint_group} Resource.
func NewGoogleComputeRegionNetworkEndpointGroup_Override(g GoogleComputeRegionNetworkEndpointGroup, scope constructs.Construct, id *string, config *GoogleComputeRegionNetworkEndpointGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeRegionNetworkEndpointGroup.GoogleComputeRegionNetworkEndpointGroup",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetNetworkEndpointType(val *string) {
	if err := j.validateSetNetworkEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkEndpointType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetPscTargetService(val *string) {
	if err := j.validateSetPscTargetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pscTargetService",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
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
func GoogleComputeRegionNetworkEndpointGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRegionNetworkEndpointGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googleComputeRegionNetworkEndpointGroup.GoogleComputeRegionNetworkEndpointGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeRegionNetworkEndpointGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRegionNetworkEndpointGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googleComputeRegionNetworkEndpointGroup.GoogleComputeRegionNetworkEndpointGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeRegionNetworkEndpointGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRegionNetworkEndpointGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googleComputeRegionNetworkEndpointGroup.GoogleComputeRegionNetworkEndpointGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeRegionNetworkEndpointGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-googlebeta.googleComputeRegionNetworkEndpointGroup.GoogleComputeRegionNetworkEndpointGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) PutAppEngine(value *GoogleComputeRegionNetworkEndpointGroupAppEngine) {
	if err := g.validatePutAppEngineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAppEngine",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) PutCloudFunction(value *GoogleComputeRegionNetworkEndpointGroupCloudFunction) {
	if err := g.validatePutCloudFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudFunction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) PutCloudRun(value *GoogleComputeRegionNetworkEndpointGroupCloudRun) {
	if err := g.validatePutCloudRunParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudRun",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) PutServerlessDeployment(value *GoogleComputeRegionNetworkEndpointGroupServerlessDeployment) {
	if err := g.validatePutServerlessDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServerlessDeployment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) PutTimeouts(value *GoogleComputeRegionNetworkEndpointGroupTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ResetAppEngine() {
	_jsii_.InvokeVoid(
		g,
		"resetAppEngine",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ResetCloudFunction() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudFunction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ResetCloudRun() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudRun",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ResetNetworkEndpointType() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkEndpointType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ResetPscTargetService() {
	_jsii_.InvokeVoid(
		g,
		"resetPscTargetService",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ResetServerlessDeployment() {
	_jsii_.InvokeVoid(
		g,
		"resetServerlessDeployment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionNetworkEndpointGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

