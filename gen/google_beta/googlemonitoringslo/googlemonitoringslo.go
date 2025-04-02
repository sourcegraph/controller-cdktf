package googlemonitoringslo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemonitoringslo/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_monitoring_slo google_monitoring_slo}.
type GoogleMonitoringSlo interface {
	cdktf.TerraformResource
	BasicSli() GoogleMonitoringSloBasicSliOutputReference
	BasicSliInput() *GoogleMonitoringSloBasicSli
	CalendarPeriod() *string
	SetCalendarPeriod(val *string)
	CalendarPeriodInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Goal() *float64
	SetGoal(val *float64)
	GoalInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RequestBasedSli() GoogleMonitoringSloRequestBasedSliOutputReference
	RequestBasedSliInput() *GoogleMonitoringSloRequestBasedSli
	RollingPeriodDays() *float64
	SetRollingPeriodDays(val *float64)
	RollingPeriodDaysInput() *float64
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	SloId() *string
	SetSloId(val *string)
	SloIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleMonitoringSloTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserLabels() *map[string]*string
	SetUserLabels(val *map[string]*string)
	UserLabelsInput() *map[string]*string
	WindowsBasedSli() GoogleMonitoringSloWindowsBasedSliOutputReference
	WindowsBasedSliInput() *GoogleMonitoringSloWindowsBasedSli
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
	PutBasicSli(value *GoogleMonitoringSloBasicSli)
	PutRequestBasedSli(value *GoogleMonitoringSloRequestBasedSli)
	PutTimeouts(value *GoogleMonitoringSloTimeouts)
	PutWindowsBasedSli(value *GoogleMonitoringSloWindowsBasedSli)
	ResetBasicSli()
	ResetCalendarPeriod()
	ResetDisplayName()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRequestBasedSli()
	ResetRollingPeriodDays()
	ResetSloId()
	ResetTimeouts()
	ResetUserLabels()
	ResetWindowsBasedSli()
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

// The jsii proxy struct for GoogleMonitoringSlo
type jsiiProxy_GoogleMonitoringSlo struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleMonitoringSlo) BasicSli() GoogleMonitoringSloBasicSliOutputReference {
	var returns GoogleMonitoringSloBasicSliOutputReference
	_jsii_.Get(
		j,
		"basicSli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) BasicSliInput() *GoogleMonitoringSloBasicSli {
	var returns *GoogleMonitoringSloBasicSli
	_jsii_.Get(
		j,
		"basicSliInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) CalendarPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calendarPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) CalendarPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calendarPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) Goal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"goal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) GoalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"goalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) RequestBasedSli() GoogleMonitoringSloRequestBasedSliOutputReference {
	var returns GoogleMonitoringSloRequestBasedSliOutputReference
	_jsii_.Get(
		j,
		"requestBasedSli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) RequestBasedSliInput() *GoogleMonitoringSloRequestBasedSli {
	var returns *GoogleMonitoringSloRequestBasedSli
	_jsii_.Get(
		j,
		"requestBasedSliInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) RollingPeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rollingPeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) RollingPeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rollingPeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) SloId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) SloIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) Timeouts() GoogleMonitoringSloTimeoutsOutputReference {
	var returns GoogleMonitoringSloTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) UserLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) UserLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) WindowsBasedSli() GoogleMonitoringSloWindowsBasedSliOutputReference {
	var returns GoogleMonitoringSloWindowsBasedSliOutputReference
	_jsii_.Get(
		j,
		"windowsBasedSli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringSlo) WindowsBasedSliInput() *GoogleMonitoringSloWindowsBasedSli {
	var returns *GoogleMonitoringSloWindowsBasedSli
	_jsii_.Get(
		j,
		"windowsBasedSliInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_monitoring_slo google_monitoring_slo} Resource.
func NewGoogleMonitoringSlo(scope constructs.Construct, id *string, config *GoogleMonitoringSloConfig) GoogleMonitoringSlo {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringSloParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringSlo{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleMonitoringSlo.GoogleMonitoringSlo",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_monitoring_slo google_monitoring_slo} Resource.
func NewGoogleMonitoringSlo_Override(g GoogleMonitoringSlo, scope constructs.Construct, id *string, config *GoogleMonitoringSloConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleMonitoringSlo.GoogleMonitoringSlo",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetCalendarPeriod(val *string) {
	if err := j.validateSetCalendarPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"calendarPeriod",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetGoal(val *float64) {
	if err := j.validateSetGoalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"goal",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetRollingPeriodDays(val *float64) {
	if err := j.validateSetRollingPeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rollingPeriodDays",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetSloId(val *string) {
	if err := j.validateSetSloIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sloId",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringSlo)SetUserLabels(val *map[string]*string) {
	if err := j.validateSetUserLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userLabels",
		val,
	)
}

// Generates CDKTF code for importing a GoogleMonitoringSlo resource upon running "cdktf plan <stack-name>".
func GoogleMonitoringSlo_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleMonitoringSlo_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleMonitoringSlo.GoogleMonitoringSlo",
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
func GoogleMonitoringSlo_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleMonitoringSlo_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleMonitoringSlo.GoogleMonitoringSlo",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleMonitoringSlo_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleMonitoringSlo_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleMonitoringSlo.GoogleMonitoringSlo",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleMonitoringSlo_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleMonitoringSlo_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleMonitoringSlo.GoogleMonitoringSlo",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleMonitoringSlo_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleMonitoringSlo.GoogleMonitoringSlo",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleMonitoringSlo) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringSlo) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringSlo) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringSlo) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringSlo) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringSlo) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringSlo) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringSlo) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringSlo) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringSlo) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSlo) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringSlo) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) PutBasicSli(value *GoogleMonitoringSloBasicSli) {
	if err := g.validatePutBasicSliParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBasicSli",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) PutRequestBasedSli(value *GoogleMonitoringSloRequestBasedSli) {
	if err := g.validatePutRequestBasedSliParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestBasedSli",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) PutTimeouts(value *GoogleMonitoringSloTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) PutWindowsBasedSli(value *GoogleMonitoringSloWindowsBasedSli) {
	if err := g.validatePutWindowsBasedSliParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWindowsBasedSli",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) ResetBasicSli() {
	_jsii_.InvokeVoid(
		g,
		"resetBasicSli",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) ResetCalendarPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetCalendarPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) ResetRequestBasedSli() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestBasedSli",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) ResetRollingPeriodDays() {
	_jsii_.InvokeVoid(
		g,
		"resetRollingPeriodDays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) ResetSloId() {
	_jsii_.InvokeVoid(
		g,
		"resetSloId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) ResetUserLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetUserLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) ResetWindowsBasedSli() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowsBasedSli",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringSlo) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSlo) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSlo) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSlo) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSlo) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringSlo) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

