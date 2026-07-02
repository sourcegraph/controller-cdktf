package kinesisanalyticsv2application

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/kinesisanalyticsv2application/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/kinesisanalyticsv2_application aws_kinesisanalyticsv2_application}.
type Kinesisanalyticsv2Application interface {
	cdktf.TerraformResource
	ApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference
	ApplicationConfigurationInput() *Kinesisanalyticsv2ApplicationApplicationConfiguration
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudwatchLoggingOptions() Kinesisanalyticsv2ApplicationCloudwatchLoggingOptionsOutputReference
	CloudwatchLoggingOptionsInput() *Kinesisanalyticsv2ApplicationCloudwatchLoggingOptions
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
	CreateTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ForceStop() any
	SetForceStop(val any)
	ForceStopInput() any
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
	LastUpdateTimestamp() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	RuntimeEnvironment() *string
	SetRuntimeEnvironment(val *string)
	RuntimeEnvironmentInput() *string
	ServiceExecutionRole() *string
	SetServiceExecutionRole(val *string)
	ServiceExecutionRoleInput() *string
	StartApplication() any
	SetStartApplication(val any)
	StartApplicationInput() any
	Status() *string
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
	Timeouts() Kinesisanalyticsv2ApplicationTimeoutsOutputReference
	TimeoutsInput() any
	VersionId() *float64
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
	PutApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfiguration)
	PutCloudwatchLoggingOptions(value *Kinesisanalyticsv2ApplicationCloudwatchLoggingOptions)
	PutTimeouts(value *Kinesisanalyticsv2ApplicationTimeouts)
	ResetApplicationConfiguration()
	ResetCloudwatchLoggingOptions()
	ResetDescription()
	ResetForceStop()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStartApplication()
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

// The jsii proxy struct for Kinesisanalyticsv2Application
type jsiiProxy_Kinesisanalyticsv2Application struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference
	_jsii_.Get(
		j,
		"applicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ApplicationConfigurationInput() *Kinesisanalyticsv2ApplicationApplicationConfiguration {
	var returns *Kinesisanalyticsv2ApplicationApplicationConfiguration
	_jsii_.Get(
		j,
		"applicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) CloudwatchLoggingOptions() Kinesisanalyticsv2ApplicationCloudwatchLoggingOptionsOutputReference {
	var returns Kinesisanalyticsv2ApplicationCloudwatchLoggingOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) CloudwatchLoggingOptionsInput() *Kinesisanalyticsv2ApplicationCloudwatchLoggingOptions {
	var returns *Kinesisanalyticsv2ApplicationCloudwatchLoggingOptions
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) CreateTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ForceStop() any {
	var returns any
	_jsii_.Get(
		j,
		"forceStop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ForceStopInput() any {
	var returns any
	_jsii_.Get(
		j,
		"forceStopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) LastUpdateTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdateTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) RuntimeEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) RuntimeEnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ServiceExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) ServiceExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) StartApplication() any {
	var returns any
	_jsii_.Get(
		j,
		"startApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) StartApplicationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"startApplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) Timeouts() Kinesisanalyticsv2ApplicationTimeoutsOutputReference {
	var returns Kinesisanalyticsv2ApplicationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) VersionId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/kinesisanalyticsv2_application aws_kinesisanalyticsv2_application} Resource.
func NewKinesisanalyticsv2Application(scope constructs.Construct, id *string, config *Kinesisanalyticsv2ApplicationConfig) Kinesisanalyticsv2Application {
	_init_.Initialize()

	if err := validateNewKinesisanalyticsv2ApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Kinesisanalyticsv2Application{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisanalyticsv2Application.Kinesisanalyticsv2Application",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/kinesisanalyticsv2_application aws_kinesisanalyticsv2_application} Resource.
func NewKinesisanalyticsv2Application_Override(k Kinesisanalyticsv2Application, scope constructs.Construct, id *string, config *Kinesisanalyticsv2ApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisanalyticsv2Application.Kinesisanalyticsv2Application",
		[]any{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetForceStop(val any) {
	if err := j.validateSetForceStopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceStop",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetRuntimeEnvironment(val *string) {
	if err := j.validateSetRuntimeEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEnvironment",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetServiceExecutionRole(val *string) {
	if err := j.validateSetServiceExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceExecutionRole",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetStartApplication(val any) {
	if err := j.validateSetStartApplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startApplication",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2Application) SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a Kinesisanalyticsv2Application resource upon running "cdktf plan <stack-name>".
func Kinesisanalyticsv2Application_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKinesisanalyticsv2Application_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.kinesisanalyticsv2Application.Kinesisanalyticsv2Application",
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
func Kinesisanalyticsv2Application_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateKinesisanalyticsv2Application_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.kinesisanalyticsv2Application.Kinesisanalyticsv2Application",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func Kinesisanalyticsv2Application_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateKinesisanalyticsv2Application_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.kinesisanalyticsv2Application.Kinesisanalyticsv2Application",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func Kinesisanalyticsv2Application_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateKinesisanalyticsv2Application_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.kinesisanalyticsv2Application.Kinesisanalyticsv2Application",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func Kinesisanalyticsv2Application_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.kinesisanalyticsv2Application.Kinesisanalyticsv2Application",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) AddOverride(path *string, value any) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]any{path, value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]any{id, provider},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]any{id},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) MoveTo(moveTarget *string, index any) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]any{id},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) PutApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfiguration) {
	if err := k.validatePutApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putApplicationConfiguration",
		[]any{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) PutCloudwatchLoggingOptions(value *Kinesisanalyticsv2ApplicationCloudwatchLoggingOptions) {
	if err := k.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCloudwatchLoggingOptions",
		[]any{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) PutTimeouts(value *Kinesisanalyticsv2ApplicationTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]any{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetApplicationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetApplicationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetDescription() {
	_jsii_.InvokeVoid(
		k,
		"resetDescription",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetForceStop() {
	_jsii_.InvokeVoid(
		k,
		"resetForceStop",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetStartApplication() {
	_jsii_.InvokeVoid(
		k,
		"resetStartApplication",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetTagsAll() {
	_jsii_.InvokeVoid(
		k,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2Application) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
