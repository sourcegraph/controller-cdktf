package googlevertexaiindexendpointdeployedindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaiindexendpointdeployedindex/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_index_endpoint_deployed_index google_vertex_ai_index_endpoint_deployed_index}.
type GoogleVertexAiIndexEndpointDeployedIndex interface {
	cdktf.TerraformResource
	AutomaticResources() GoogleVertexAiIndexEndpointDeployedIndexAutomaticResourcesOutputReference
	AutomaticResourcesInput() *GoogleVertexAiIndexEndpointDeployedIndexAutomaticResources
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DedicatedResources() GoogleVertexAiIndexEndpointDeployedIndexDedicatedResourcesOutputReference
	DedicatedResourcesInput() *GoogleVertexAiIndexEndpointDeployedIndexDedicatedResources
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeployedIndexAuthConfig() GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigOutputReference
	DeployedIndexAuthConfigInput() *GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig
	DeployedIndexId() *string
	SetDeployedIndexId(val *string)
	DeployedIndexIdInput() *string
	DeploymentGroup() *string
	SetDeploymentGroup(val *string)
	DeploymentGroupInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EnableAccessLogging() interface{}
	SetEnableAccessLogging(val interface{})
	EnableAccessLoggingInput() interface{}
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
	Index() *string
	SetIndex(val *string)
	IndexEndpoint() *string
	SetIndexEndpoint(val *string)
	IndexEndpointInput() *string
	IndexInput() *string
	IndexSyncTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	PrivateEndpoints() GoogleVertexAiIndexEndpointDeployedIndexPrivateEndpointsList
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
	ReservedIpRanges() *[]*string
	SetReservedIpRanges(val *[]*string)
	ReservedIpRangesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleVertexAiIndexEndpointDeployedIndexTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
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
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAutomaticResources(value *GoogleVertexAiIndexEndpointDeployedIndexAutomaticResources)
	PutDedicatedResources(value *GoogleVertexAiIndexEndpointDeployedIndexDedicatedResources)
	PutDeployedIndexAuthConfig(value *GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig)
	PutTimeouts(value *GoogleVertexAiIndexEndpointDeployedIndexTimeouts)
	ResetAutomaticResources()
	ResetDedicatedResources()
	ResetDeployedIndexAuthConfig()
	ResetDeploymentGroup()
	ResetDisplayName()
	ResetEnableAccessLogging()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReservedIpRanges()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleVertexAiIndexEndpointDeployedIndex
type jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) AutomaticResources() GoogleVertexAiIndexEndpointDeployedIndexAutomaticResourcesOutputReference {
	var returns GoogleVertexAiIndexEndpointDeployedIndexAutomaticResourcesOutputReference
	_jsii_.Get(
		j,
		"automaticResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) AutomaticResourcesInput() *GoogleVertexAiIndexEndpointDeployedIndexAutomaticResources {
	var returns *GoogleVertexAiIndexEndpointDeployedIndexAutomaticResources
	_jsii_.Get(
		j,
		"automaticResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DedicatedResources() GoogleVertexAiIndexEndpointDeployedIndexDedicatedResourcesOutputReference {
	var returns GoogleVertexAiIndexEndpointDeployedIndexDedicatedResourcesOutputReference
	_jsii_.Get(
		j,
		"dedicatedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DedicatedResourcesInput() *GoogleVertexAiIndexEndpointDeployedIndexDedicatedResources {
	var returns *GoogleVertexAiIndexEndpointDeployedIndexDedicatedResources
	_jsii_.Get(
		j,
		"dedicatedResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DeployedIndexAuthConfig() GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigOutputReference {
	var returns GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigOutputReference
	_jsii_.Get(
		j,
		"deployedIndexAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DeployedIndexAuthConfigInput() *GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig {
	var returns *GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig
	_jsii_.Get(
		j,
		"deployedIndexAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DeployedIndexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployedIndexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DeployedIndexIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployedIndexIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DeploymentGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DeploymentGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) EnableAccessLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAccessLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) EnableAccessLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAccessLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) IndexEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) IndexEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) IndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) IndexSyncTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexSyncTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) PrivateEndpoints() GoogleVertexAiIndexEndpointDeployedIndexPrivateEndpointsList {
	var returns GoogleVertexAiIndexEndpointDeployedIndexPrivateEndpointsList
	_jsii_.Get(
		j,
		"privateEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ReservedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ReservedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) Timeouts() GoogleVertexAiIndexEndpointDeployedIndexTimeoutsOutputReference {
	var returns GoogleVertexAiIndexEndpointDeployedIndexTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_index_endpoint_deployed_index google_vertex_ai_index_endpoint_deployed_index} Resource.
func NewGoogleVertexAiIndexEndpointDeployedIndex(scope constructs.Construct, id *string, config *GoogleVertexAiIndexEndpointDeployedIndexConfig) GoogleVertexAiIndexEndpointDeployedIndex {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiIndexEndpointDeployedIndexParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiIndexEndpointDeployedIndex.GoogleVertexAiIndexEndpointDeployedIndex",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_index_endpoint_deployed_index google_vertex_ai_index_endpoint_deployed_index} Resource.
func NewGoogleVertexAiIndexEndpointDeployedIndex_Override(g GoogleVertexAiIndexEndpointDeployedIndex, scope constructs.Construct, id *string, config *GoogleVertexAiIndexEndpointDeployedIndexConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiIndexEndpointDeployedIndex.GoogleVertexAiIndexEndpointDeployedIndex",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetDeployedIndexId(val *string) {
	if err := j.validateSetDeployedIndexIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deployedIndexId",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetDeploymentGroup(val *string) {
	if err := j.validateSetDeploymentGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentGroup",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetEnableAccessLogging(val interface{}) {
	if err := j.validateSetEnableAccessLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAccessLogging",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetIndex(val *string) {
	if err := j.validateSetIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetIndexEndpoint(val *string) {
	if err := j.validateSetIndexEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex)SetReservedIpRanges(val *[]*string) {
	if err := j.validateSetReservedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedIpRanges",
		val,
	)
}

// Generates CDKTF code for importing a GoogleVertexAiIndexEndpointDeployedIndex resource upon running "cdktf plan <stack-name>".
func GoogleVertexAiIndexEndpointDeployedIndex_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleVertexAiIndexEndpointDeployedIndex_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleVertexAiIndexEndpointDeployedIndex.GoogleVertexAiIndexEndpointDeployedIndex",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
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
func GoogleVertexAiIndexEndpointDeployedIndex_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiIndexEndpointDeployedIndex_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleVertexAiIndexEndpointDeployedIndex.GoogleVertexAiIndexEndpointDeployedIndex",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleVertexAiIndexEndpointDeployedIndex_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiIndexEndpointDeployedIndex_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleVertexAiIndexEndpointDeployedIndex.GoogleVertexAiIndexEndpointDeployedIndex",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleVertexAiIndexEndpointDeployedIndex_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiIndexEndpointDeployedIndex_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleVertexAiIndexEndpointDeployedIndex.GoogleVertexAiIndexEndpointDeployedIndex",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleVertexAiIndexEndpointDeployedIndex_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleVertexAiIndexEndpointDeployedIndex.GoogleVertexAiIndexEndpointDeployedIndex",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) PutAutomaticResources(value *GoogleVertexAiIndexEndpointDeployedIndexAutomaticResources) {
	if err := g.validatePutAutomaticResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutomaticResources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) PutDedicatedResources(value *GoogleVertexAiIndexEndpointDeployedIndexDedicatedResources) {
	if err := g.validatePutDedicatedResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDedicatedResources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) PutDeployedIndexAuthConfig(value *GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfig) {
	if err := g.validatePutDeployedIndexAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeployedIndexAuthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) PutTimeouts(value *GoogleVertexAiIndexEndpointDeployedIndexTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetAutomaticResources() {
	_jsii_.InvokeVoid(
		g,
		"resetAutomaticResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetDedicatedResources() {
	_jsii_.InvokeVoid(
		g,
		"resetDedicatedResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetDeployedIndexAuthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDeployedIndexAuthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetDeploymentGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetEnableAccessLogging() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableAccessLogging",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetReservedIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetReservedIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexEndpointDeployedIndex) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

