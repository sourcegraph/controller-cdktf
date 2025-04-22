package googleosconfigpatchdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleosconfigpatchdeployment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_patch_deployment google_os_config_patch_deployment}.
type GoogleOsConfigPatchDeployment interface {
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
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
	InstanceFilter() GoogleOsConfigPatchDeploymentInstanceFilterOutputReference
	InstanceFilterInput() *GoogleOsConfigPatchDeploymentInstanceFilter
	LastExecuteTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	OneTimeSchedule() GoogleOsConfigPatchDeploymentOneTimeScheduleOutputReference
	OneTimeScheduleInput() *GoogleOsConfigPatchDeploymentOneTimeSchedule
	PatchConfig() GoogleOsConfigPatchDeploymentPatchConfigOutputReference
	PatchConfigInput() *GoogleOsConfigPatchDeploymentPatchConfig
	PatchDeploymentId() *string
	SetPatchDeploymentId(val *string)
	PatchDeploymentIdInput() *string
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
	RecurringSchedule() GoogleOsConfigPatchDeploymentRecurringScheduleOutputReference
	RecurringScheduleInput() *GoogleOsConfigPatchDeploymentRecurringSchedule
	Rollout() GoogleOsConfigPatchDeploymentRolloutOutputReference
	RolloutInput() *GoogleOsConfigPatchDeploymentRollout
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleOsConfigPatchDeploymentTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
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
	PutInstanceFilter(value *GoogleOsConfigPatchDeploymentInstanceFilter)
	PutOneTimeSchedule(value *GoogleOsConfigPatchDeploymentOneTimeSchedule)
	PutPatchConfig(value *GoogleOsConfigPatchDeploymentPatchConfig)
	PutRecurringSchedule(value *GoogleOsConfigPatchDeploymentRecurringSchedule)
	PutRollout(value *GoogleOsConfigPatchDeploymentRollout)
	PutTimeouts(value *GoogleOsConfigPatchDeploymentTimeouts)
	ResetDescription()
	ResetDuration()
	ResetId()
	ResetOneTimeSchedule()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPatchConfig()
	ResetProject()
	ResetRecurringSchedule()
	ResetRollout()
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

// The jsii proxy struct for GoogleOsConfigPatchDeployment
type jsiiProxy_GoogleOsConfigPatchDeployment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) InstanceFilter() GoogleOsConfigPatchDeploymentInstanceFilterOutputReference {
	var returns GoogleOsConfigPatchDeploymentInstanceFilterOutputReference
	_jsii_.Get(
		j,
		"instanceFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) InstanceFilterInput() *GoogleOsConfigPatchDeploymentInstanceFilter {
	var returns *GoogleOsConfigPatchDeploymentInstanceFilter
	_jsii_.Get(
		j,
		"instanceFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) LastExecuteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastExecuteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) OneTimeSchedule() GoogleOsConfigPatchDeploymentOneTimeScheduleOutputReference {
	var returns GoogleOsConfigPatchDeploymentOneTimeScheduleOutputReference
	_jsii_.Get(
		j,
		"oneTimeSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) OneTimeScheduleInput() *GoogleOsConfigPatchDeploymentOneTimeSchedule {
	var returns *GoogleOsConfigPatchDeploymentOneTimeSchedule
	_jsii_.Get(
		j,
		"oneTimeScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) PatchConfig() GoogleOsConfigPatchDeploymentPatchConfigOutputReference {
	var returns GoogleOsConfigPatchDeploymentPatchConfigOutputReference
	_jsii_.Get(
		j,
		"patchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) PatchConfigInput() *GoogleOsConfigPatchDeploymentPatchConfig {
	var returns *GoogleOsConfigPatchDeploymentPatchConfig
	_jsii_.Get(
		j,
		"patchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) PatchDeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchDeploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) PatchDeploymentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchDeploymentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) RecurringSchedule() GoogleOsConfigPatchDeploymentRecurringScheduleOutputReference {
	var returns GoogleOsConfigPatchDeploymentRecurringScheduleOutputReference
	_jsii_.Get(
		j,
		"recurringSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) RecurringScheduleInput() *GoogleOsConfigPatchDeploymentRecurringSchedule {
	var returns *GoogleOsConfigPatchDeploymentRecurringSchedule
	_jsii_.Get(
		j,
		"recurringScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) Rollout() GoogleOsConfigPatchDeploymentRolloutOutputReference {
	var returns GoogleOsConfigPatchDeploymentRolloutOutputReference
	_jsii_.Get(
		j,
		"rollout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) RolloutInput() *GoogleOsConfigPatchDeploymentRollout {
	var returns *GoogleOsConfigPatchDeploymentRollout
	_jsii_.Get(
		j,
		"rolloutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) Timeouts() GoogleOsConfigPatchDeploymentTimeoutsOutputReference {
	var returns GoogleOsConfigPatchDeploymentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_patch_deployment google_os_config_patch_deployment} Resource.
func NewGoogleOsConfigPatchDeployment(scope constructs.Construct, id *string, config *GoogleOsConfigPatchDeploymentConfig) GoogleOsConfigPatchDeployment {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigPatchDeploymentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigPatchDeployment{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeployment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_patch_deployment google_os_config_patch_deployment} Resource.
func NewGoogleOsConfigPatchDeployment_Override(g GoogleOsConfigPatchDeployment, scope constructs.Construct, id *string, config *GoogleOsConfigPatchDeploymentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeployment",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment)SetPatchDeploymentId(val *string) {
	if err := j.validateSetPatchDeploymentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchDeploymentId",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeployment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleOsConfigPatchDeployment resource upon running "cdktf plan <stack-name>".
func GoogleOsConfigPatchDeployment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleOsConfigPatchDeployment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeployment",
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
func GoogleOsConfigPatchDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleOsConfigPatchDeployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleOsConfigPatchDeployment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleOsConfigPatchDeployment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeployment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleOsConfigPatchDeployment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleOsConfigPatchDeployment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeployment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleOsConfigPatchDeployment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeployment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) PutInstanceFilter(value *GoogleOsConfigPatchDeploymentInstanceFilter) {
	if err := g.validatePutInstanceFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstanceFilter",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) PutOneTimeSchedule(value *GoogleOsConfigPatchDeploymentOneTimeSchedule) {
	if err := g.validatePutOneTimeScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOneTimeSchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) PutPatchConfig(value *GoogleOsConfigPatchDeploymentPatchConfig) {
	if err := g.validatePutPatchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPatchConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) PutRecurringSchedule(value *GoogleOsConfigPatchDeploymentRecurringSchedule) {
	if err := g.validatePutRecurringScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRecurringSchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) PutRollout(value *GoogleOsConfigPatchDeploymentRollout) {
	if err := g.validatePutRolloutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRollout",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) PutTimeouts(value *GoogleOsConfigPatchDeploymentTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ResetDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ResetOneTimeSchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetOneTimeSchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ResetPatchConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPatchConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ResetRecurringSchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetRecurringSchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ResetRollout() {
	_jsii_.InvokeVoid(
		g,
		"resetRollout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeployment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

