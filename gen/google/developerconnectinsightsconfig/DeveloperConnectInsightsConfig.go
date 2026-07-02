package developerconnectinsightsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/developerconnectinsightsconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/developer_connect_insights_config google_developer_connect_insights_config}.
type DeveloperConnectInsightsConfig interface {
	cdktf.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	AppHubApplication() *string
	SetAppHubApplication(val *string)
	AppHubApplicationInput() *string
	ArtifactConfigs() DeveloperConnectInsightsConfigArtifactConfigsList
	ArtifactConfigsInput() any
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
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveAnnotations() cdktf.StringMap
	EffectiveLabels() cdktf.StringMap
	Errors() DeveloperConnectInsightsConfigErrorsList
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
	InsightsConfigId() *string
	SetInsightsConfigId(val *string)
	InsightsConfigIdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
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
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	Reconciling() cdktf.IResolvable
	RuntimeConfigs() DeveloperConnectInsightsConfigRuntimeConfigsList
	State() *string
	TargetProjects() DeveloperConnectInsightsConfigTargetProjectsOutputReference
	TargetProjectsInput() *DeveloperConnectInsightsConfigTargetProjects
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DeveloperConnectInsightsConfigTimeoutsOutputReference
	TimeoutsInput() any
	UpdateTime() *string
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
	PutArtifactConfigs(value any)
	PutTargetProjects(value *DeveloperConnectInsightsConfigTargetProjects)
	PutTimeouts(value *DeveloperConnectInsightsConfigTimeouts)
	ResetAnnotations()
	ResetAppHubApplication()
	ResetArtifactConfigs()
	ResetId()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetTargetProjects()
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

// The jsii proxy struct for DeveloperConnectInsightsConfig
type jsiiProxy_DeveloperConnectInsightsConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) AppHubApplication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appHubApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) AppHubApplicationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appHubApplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) ArtifactConfigs() DeveloperConnectInsightsConfigArtifactConfigsList {
	var returns DeveloperConnectInsightsConfigArtifactConfigsList
	_jsii_.Get(
		j,
		"artifactConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) ArtifactConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"artifactConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Errors() DeveloperConnectInsightsConfigErrorsList {
	var returns DeveloperConnectInsightsConfigErrorsList
	_jsii_.Get(
		j,
		"errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) InsightsConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insightsConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) InsightsConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insightsConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) RuntimeConfigs() DeveloperConnectInsightsConfigRuntimeConfigsList {
	var returns DeveloperConnectInsightsConfigRuntimeConfigsList
	_jsii_.Get(
		j,
		"runtimeConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) TargetProjects() DeveloperConnectInsightsConfigTargetProjectsOutputReference {
	var returns DeveloperConnectInsightsConfigTargetProjectsOutputReference
	_jsii_.Get(
		j,
		"targetProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) TargetProjectsInput() *DeveloperConnectInsightsConfigTargetProjects {
	var returns *DeveloperConnectInsightsConfigTargetProjects
	_jsii_.Get(
		j,
		"targetProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) Timeouts() DeveloperConnectInsightsConfigTimeoutsOutputReference {
	var returns DeveloperConnectInsightsConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/developer_connect_insights_config google_developer_connect_insights_config} Resource.
func NewDeveloperConnectInsightsConfig(scope constructs.Construct, id *string, config *DeveloperConnectInsightsConfigConfig) DeveloperConnectInsightsConfig {
	_init_.Initialize()

	if err := validateNewDeveloperConnectInsightsConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeveloperConnectInsightsConfig{}

	_jsii_.Create(
		"@cdktf/provider-google.developerConnectInsightsConfig.DeveloperConnectInsightsConfig",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/developer_connect_insights_config google_developer_connect_insights_config} Resource.
func NewDeveloperConnectInsightsConfig_Override(d DeveloperConnectInsightsConfig, scope constructs.Construct, id *string, config *DeveloperConnectInsightsConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.developerConnectInsightsConfig.DeveloperConnectInsightsConfig",
		[]any{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) SetAppHubApplication(val *string) {
	if err := j.validateSetAppHubApplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appHubApplication",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) SetInsightsConfigId(val *string) {
	if err := j.validateSetInsightsConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insightsConfigId",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectInsightsConfig) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a DeveloperConnectInsightsConfig resource upon running "cdktf plan <stack-name>".
func DeveloperConnectInsightsConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDeveloperConnectInsightsConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.developerConnectInsightsConfig.DeveloperConnectInsightsConfig",
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
func DeveloperConnectInsightsConfig_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateDeveloperConnectInsightsConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.developerConnectInsightsConfig.DeveloperConnectInsightsConfig",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DeveloperConnectInsightsConfig_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateDeveloperConnectInsightsConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.developerConnectInsightsConfig.DeveloperConnectInsightsConfig",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DeveloperConnectInsightsConfig_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateDeveloperConnectInsightsConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.developerConnectInsightsConfig.DeveloperConnectInsightsConfig",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func DeveloperConnectInsightsConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.developerConnectInsightsConfig.DeveloperConnectInsightsConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) AddOverride(path *string, value any) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]any{path, value},
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DeveloperConnectInsightsConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeveloperConnectInsightsConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeveloperConnectInsightsConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeveloperConnectInsightsConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeveloperConnectInsightsConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeveloperConnectInsightsConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeveloperConnectInsightsConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeveloperConnectInsightsConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeveloperConnectInsightsConfig) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]any{id, provider},
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeveloperConnectInsightsConfig) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]any{id},
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) MoveTo(moveTarget *string, index any) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]any{id},
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) PutArtifactConfigs(value any) {
	if err := d.validatePutArtifactConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putArtifactConfigs",
		[]any{value},
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) PutTargetProjects(value *DeveloperConnectInsightsConfigTargetProjects) {
	if err := d.validatePutTargetProjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTargetProjects",
		[]any{value},
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) PutTimeouts(value *DeveloperConnectInsightsConfigTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]any{value},
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) ResetAnnotations() {
	_jsii_.InvokeVoid(
		d,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) ResetAppHubApplication() {
	_jsii_.InvokeVoid(
		d,
		"resetAppHubApplication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) ResetArtifactConfigs() {
	_jsii_.InvokeVoid(
		d,
		"resetArtifactConfigs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) ResetTargetProjects() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetProjects",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectInsightsConfig) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
