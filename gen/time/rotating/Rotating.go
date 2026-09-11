package rotating

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/time/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/time/rotating/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/rotating time_rotating}.
type Rotating interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	Day() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hour() *float64
	Id() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Minute() *float64
	Month() *float64
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Rfc3339() *string
	SetRfc3339(val *string)
	Rfc3339Input() *string
	RotationDays() *float64
	SetRotationDays(val *float64)
	RotationDaysInput() *float64
	RotationHours() *float64
	SetRotationHours(val *float64)
	RotationHoursInput() *float64
	RotationMinutes() *float64
	SetRotationMinutes(val *float64)
	RotationMinutesInput() *float64
	RotationMonths() *float64
	SetRotationMonths(val *float64)
	RotationMonthsInput() *float64
	RotationRfc3339() *string
	SetRotationRfc3339(val *string)
	RotationRfc3339Input() *string
	RotationYears() *float64
	SetRotationYears(val *float64)
	RotationYearsInput() *float64
	Second() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Triggers() *map[string]*string
	SetTriggers(val *map[string]*string)
	TriggersInput() *map[string]*string
	Unix() *float64
	Year() *float64
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
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
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRfc3339()
	ResetRotationDays()
	ResetRotationHours()
	ResetRotationMinutes()
	ResetRotationMonths()
	ResetRotationRfc3339()
	ResetRotationYears()
	ResetTriggers()
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for Rotating
type jsiiProxy_Rotating struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Rotating) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Day() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"day",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Hour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Minute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Month() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"month",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Rfc3339() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rfc3339",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Rfc3339Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rfc3339Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationMonths() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationMonths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationMonthsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationMonthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationRfc3339() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationRfc3339",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationRfc3339Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationRfc3339Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationYears() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationYears",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) RotationYearsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationYearsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Second() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"second",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Triggers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) TriggersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Unix() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rotating) Year() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"year",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/rotating time_rotating} Resource.
func NewRotating(scope constructs.Construct, id *string, config *RotatingConfig) Rotating {
	_init_.Initialize()

	if err := validateNewRotatingParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Rotating{}

	_jsii_.Create(
		"@cdktn/provider-time.rotating.Rotating",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/rotating time_rotating} Resource.
func NewRotating_Override(r Rotating, scope constructs.Construct, id *string, config *RotatingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-time.rotating.Rotating",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_Rotating)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetRfc3339(val *string) {
	if err := j.validateSetRfc3339Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rfc3339",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetRotationDays(val *float64) {
	if err := j.validateSetRotationDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationDays",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetRotationHours(val *float64) {
	if err := j.validateSetRotationHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationHours",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetRotationMinutes(val *float64) {
	if err := j.validateSetRotationMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationMinutes",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetRotationMonths(val *float64) {
	if err := j.validateSetRotationMonthsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationMonths",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetRotationRfc3339(val *string) {
	if err := j.validateSetRotationRfc3339Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationRfc3339",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetRotationYears(val *float64) {
	if err := j.validateSetRotationYearsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationYears",
		val,
	)
}

func (j *jsiiProxy_Rotating)SetTriggers(val *map[string]*string) {
	if err := j.validateSetTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

// Generates CDKTN code for importing a Rotating resource upon running "cdktn plan <stack-name>".
func Rotating_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateRotating_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-time.rotating.Rotating",
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
func Rotating_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRotating_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-time.rotating.Rotating",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Rotating_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRotating_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-time.rotating.Rotating",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Rotating_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRotating_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-time.rotating.Rotating",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Rotating_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-time.rotating.Rotating",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Rotating) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_Rotating) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_Rotating) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_Rotating) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := r.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Rotating) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_Rotating) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Rotating) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_Rotating) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := r.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (r *jsiiProxy_Rotating) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetRfc3339() {
	_jsii_.InvokeVoid(
		r,
		"resetRfc3339",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetRotationDays() {
	_jsii_.InvokeVoid(
		r,
		"resetRotationDays",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetRotationHours() {
	_jsii_.InvokeVoid(
		r,
		"resetRotationHours",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetRotationMinutes() {
	_jsii_.InvokeVoid(
		r,
		"resetRotationMinutes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetRotationMonths() {
	_jsii_.InvokeVoid(
		r,
		"resetRotationMonths",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetRotationRfc3339() {
	_jsii_.InvokeVoid(
		r,
		"resetRotationRfc3339",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetRotationYears() {
	_jsii_.InvokeVoid(
		r,
		"resetRotationYears",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) ResetTriggers() {
	_jsii_.InvokeVoid(
		r,
		"resetTriggers",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Rotating) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rotating) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		r,
		"with",
		args,
		&returns,
	)

	return returns
}

