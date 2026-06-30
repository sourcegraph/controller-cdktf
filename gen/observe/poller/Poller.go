package poller

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observe/poller/internal"
)

// Represents a {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller observe_poller}.
type Poller interface {
	cdktf.TerraformResource
	AwsSnapshot() PollerAwsSnapshotOutputReference
	AwsSnapshotInput() *PollerAwsSnapshot
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Chunk() PollerChunkOutputReference
	ChunkInput() *PollerChunk
	CloudwatchMetrics() PollerCloudwatchMetricsOutputReference
	CloudwatchMetricsInput() *PollerCloudwatchMetrics
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
	Datastream() *string
	SetDatastream(val *string)
	DatastreamInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disabled() any
	SetDisabled(val any)
	DisabledInput() any
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcpMonitoring() PollerGcpMonitoringOutputReference
	GcpMonitoringInput() *PollerGcpMonitoring
	Http() PollerHttpOutputReference
	HttpInput() *PollerHttp
	Id() *string
	SetId(val *string)
	IdInput() *string
	Interval() *string
	SetInterval(val *string)
	IntervalInput() *string
	Kind() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Mongodbatlas() PollerMongodbatlasOutputReference
	MongodbatlasInput() *PollerMongodbatlas
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Oid() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	Pubsub() PollerPubsubOutputReference
	PubsubInput() *PollerPubsub
	// Experimental.
	RawOverrides() any
	Retries() *float64
	SetRetries(val *float64)
	RetriesInput() *float64
	SkipExternalValidation() any
	SetSkipExternalValidation(val any)
	SkipExternalValidationInput() any
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Workspace() *string
	SetWorkspace(val *string)
	WorkspaceInput() *string
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
	PutAwsSnapshot(value *PollerAwsSnapshot)
	PutChunk(value *PollerChunk)
	PutCloudwatchMetrics(value *PollerCloudwatchMetrics)
	PutGcpMonitoring(value *PollerGcpMonitoring)
	PutHttp(value *PollerHttp)
	PutMongodbatlas(value *PollerMongodbatlas)
	PutPubsub(value *PollerPubsub)
	ResetAwsSnapshot()
	ResetChunk()
	ResetCloudwatchMetrics()
	ResetDatastream()
	ResetDisabled()
	ResetGcpMonitoring()
	ResetHttp()
	ResetId()
	ResetInterval()
	ResetMongodbatlas()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPubsub()
	ResetRetries()
	ResetSkipExternalValidation()
	ResetTags()
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

// The jsii proxy struct for Poller
type jsiiProxy_Poller struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Poller) AwsSnapshot() PollerAwsSnapshotOutputReference {
	var returns PollerAwsSnapshotOutputReference
	_jsii_.Get(
		j,
		"awsSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) AwsSnapshotInput() *PollerAwsSnapshot {
	var returns *PollerAwsSnapshot
	_jsii_.Get(
		j,
		"awsSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Chunk() PollerChunkOutputReference {
	var returns PollerChunkOutputReference
	_jsii_.Get(
		j,
		"chunk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) ChunkInput() *PollerChunk {
	var returns *PollerChunk
	_jsii_.Get(
		j,
		"chunkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) CloudwatchMetrics() PollerCloudwatchMetricsOutputReference {
	var returns PollerCloudwatchMetricsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) CloudwatchMetricsInput() *PollerCloudwatchMetrics {
	var returns *PollerCloudwatchMetrics
	_jsii_.Get(
		j,
		"cloudwatchMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Datastream() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) DatastreamInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastreamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Disabled() any {
	var returns any
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) DisabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) GcpMonitoring() PollerGcpMonitoringOutputReference {
	var returns PollerGcpMonitoringOutputReference
	_jsii_.Get(
		j,
		"gcpMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) GcpMonitoringInput() *PollerGcpMonitoring {
	var returns *PollerGcpMonitoring
	_jsii_.Get(
		j,
		"gcpMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Http() PollerHttpOutputReference {
	var returns PollerHttpOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) HttpInput() *PollerHttp {
	var returns *PollerHttp
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Interval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) IntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Mongodbatlas() PollerMongodbatlasOutputReference {
	var returns PollerMongodbatlasOutputReference
	_jsii_.Get(
		j,
		"mongodbatlas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) MongodbatlasInput() *PollerMongodbatlas {
	var returns *PollerMongodbatlas
	_jsii_.Get(
		j,
		"mongodbatlasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Oid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Pubsub() PollerPubsubOutputReference {
	var returns PollerPubsubOutputReference
	_jsii_.Get(
		j,
		"pubsub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) PubsubInput() *PollerPubsub {
	var returns *PollerPubsub
	_jsii_.Get(
		j,
		"pubsubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Retries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) RetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) SkipExternalValidation() any {
	var returns any
	_jsii_.Get(
		j,
		"skipExternalValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) SkipExternalValidationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"skipExternalValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) Workspace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poller) WorkspaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller observe_poller} Resource.
func NewPoller(scope constructs.Construct, id *string, config *PollerConfig) Poller {
	_init_.Initialize()

	if err := validateNewPollerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Poller{}

	_jsii_.Create(
		"@cdktf/provider-observe.poller.Poller",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller observe_poller} Resource.
func NewPoller_Override(p Poller, scope constructs.Construct, id *string, config *PollerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observe.poller.Poller",
		[]any{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_Poller) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Poller) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Poller) SetDatastream(val *string) {
	if err := j.validateSetDatastreamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datastream",
		val,
	)
}

func (j *jsiiProxy_Poller) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Poller) SetDisabled(val any) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_Poller) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Poller) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Poller) SetInterval(val *string) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_Poller) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Poller) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Poller) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Poller) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Poller) SetRetries(val *float64) {
	if err := j.validateSetRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retries",
		val,
	)
}

func (j *jsiiProxy_Poller) SetSkipExternalValidation(val any) {
	if err := j.validateSetSkipExternalValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipExternalValidation",
		val,
	)
}

func (j *jsiiProxy_Poller) SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Poller) SetWorkspace(val *string) {
	if err := j.validateSetWorkspaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspace",
		val,
	)
}

// Generates CDKTF code for importing a Poller resource upon running "cdktf plan <stack-name>".
func Poller_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePoller_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-observe.poller.Poller",
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
func Poller_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validatePoller_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-observe.poller.Poller",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func Poller_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validatePoller_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-observe.poller.Poller",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func Poller_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validatePoller_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-observe.poller.Poller",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func Poller_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-observe.poller.Poller",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Poller) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (p *jsiiProxy_Poller) AddOverride(path *string, value any) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]any{path, value},
	)
}

func (p *jsiiProxy_Poller) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]any{id, provider},
	)
}

func (p *jsiiProxy_Poller) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]any{id},
	)
}

func (p *jsiiProxy_Poller) MoveTo(moveTarget *string, index any) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (p *jsiiProxy_Poller) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]any{id},
	)
}

func (p *jsiiProxy_Poller) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (p *jsiiProxy_Poller) PutAwsSnapshot(value *PollerAwsSnapshot) {
	if err := p.validatePutAwsSnapshotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAwsSnapshot",
		[]any{value},
	)
}

func (p *jsiiProxy_Poller) PutChunk(value *PollerChunk) {
	if err := p.validatePutChunkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putChunk",
		[]any{value},
	)
}

func (p *jsiiProxy_Poller) PutCloudwatchMetrics(value *PollerCloudwatchMetrics) {
	if err := p.validatePutCloudwatchMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCloudwatchMetrics",
		[]any{value},
	)
}

func (p *jsiiProxy_Poller) PutGcpMonitoring(value *PollerGcpMonitoring) {
	if err := p.validatePutGcpMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGcpMonitoring",
		[]any{value},
	)
}

func (p *jsiiProxy_Poller) PutHttp(value *PollerHttp) {
	if err := p.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHttp",
		[]any{value},
	)
}

func (p *jsiiProxy_Poller) PutMongodbatlas(value *PollerMongodbatlas) {
	if err := p.validatePutMongodbatlasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMongodbatlas",
		[]any{value},
	)
}

func (p *jsiiProxy_Poller) PutPubsub(value *PollerPubsub) {
	if err := p.validatePutPubsubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPubsub",
		[]any{value},
	)
}

func (p *jsiiProxy_Poller) ResetAwsSnapshot() {
	_jsii_.InvokeVoid(
		p,
		"resetAwsSnapshot",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) ResetChunk() {
	_jsii_.InvokeVoid(
		p,
		"resetChunk",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) ResetCloudwatchMetrics() {
	_jsii_.InvokeVoid(
		p,
		"resetCloudwatchMetrics",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) ResetDatastream() {
	_jsii_.InvokeVoid(
		p,
		"resetDatastream",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) ResetDisabled() {
	_jsii_.InvokeVoid(
		p,
		"resetDisabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) ResetGcpMonitoring() {
	_jsii_.InvokeVoid(
		p,
		"resetGcpMonitoring",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) ResetHttp() {
	_jsii_.InvokeVoid(
		p,
		"resetHttp",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) ResetInterval() {
	_jsii_.InvokeVoid(
		p,
		"resetInterval",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) ResetMongodbatlas() {
	_jsii_.InvokeVoid(
		p,
		"resetMongodbatlas",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) ResetPubsub() {
	_jsii_.InvokeVoid(
		p,
		"resetPubsub",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) ResetRetries() {
	_jsii_.InvokeVoid(
		p,
		"resetRetries",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) ResetSkipExternalValidation() {
	_jsii_.InvokeVoid(
		p,
		"resetSkipExternalValidation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poller) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poller) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
