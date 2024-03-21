package offset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/time/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/time/offset/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/offset time_offset}.
type Offset interface {
	cdktf.TerraformResource
	BaseRfc3339() *string
	SetBaseRfc3339(val *string)
	BaseRfc3339Input() *string
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
	Day() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hour() *float64
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Minute() *float64
	Month() *float64
	// The tree node.
	Node() constructs.Node
	OffsetDays() *float64
	SetOffsetDays(val *float64)
	OffsetDaysInput() *float64
	OffsetHours() *float64
	SetOffsetHours(val *float64)
	OffsetHoursInput() *float64
	OffsetMinutes() *float64
	SetOffsetMinutes(val *float64)
	OffsetMinutesInput() *float64
	OffsetMonths() *float64
	SetOffsetMonths(val *float64)
	OffsetMonthsInput() *float64
	OffsetSeconds() *float64
	SetOffsetSeconds(val *float64)
	OffsetSecondsInput() *float64
	OffsetYears() *float64
	SetOffsetYears(val *float64)
	OffsetYearsInput() *float64
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
	Rfc3339() *string
	Second() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
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
	ResetBaseRfc3339()
	ResetOffsetDays()
	ResetOffsetHours()
	ResetOffsetMinutes()
	ResetOffsetMonths()
	ResetOffsetSeconds()
	ResetOffsetYears()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTriggers()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Offset
type jsiiProxy_Offset struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Offset) BaseRfc3339() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseRfc3339",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) BaseRfc3339Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseRfc3339Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Day() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"day",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Hour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Minute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Month() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"month",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) OffsetDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) OffsetDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) OffsetHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) OffsetHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) OffsetMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) OffsetMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) OffsetMonths() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetMonths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) OffsetMonthsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetMonthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) OffsetSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) OffsetSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) OffsetYears() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetYears",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) OffsetYearsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetYearsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Rfc3339() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rfc3339",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Second() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"second",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Triggers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) TriggersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Unix() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Offset) Year() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"year",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/offset time_offset} Resource.
func NewOffset(scope constructs.Construct, id *string, config *OffsetConfig) Offset {
	_init_.Initialize()

	if err := validateNewOffsetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Offset{}

	_jsii_.Create(
		"@cdktf/provider-time.offset.Offset",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/offset time_offset} Resource.
func NewOffset_Override(o Offset, scope constructs.Construct, id *string, config *OffsetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-time.offset.Offset",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_Offset)SetBaseRfc3339(val *string) {
	if err := j.validateSetBaseRfc3339Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseRfc3339",
		val,
	)
}

func (j *jsiiProxy_Offset)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Offset)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Offset)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Offset)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Offset)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Offset)SetOffsetDays(val *float64) {
	if err := j.validateSetOffsetDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offsetDays",
		val,
	)
}

func (j *jsiiProxy_Offset)SetOffsetHours(val *float64) {
	if err := j.validateSetOffsetHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offsetHours",
		val,
	)
}

func (j *jsiiProxy_Offset)SetOffsetMinutes(val *float64) {
	if err := j.validateSetOffsetMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offsetMinutes",
		val,
	)
}

func (j *jsiiProxy_Offset)SetOffsetMonths(val *float64) {
	if err := j.validateSetOffsetMonthsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offsetMonths",
		val,
	)
}

func (j *jsiiProxy_Offset)SetOffsetSeconds(val *float64) {
	if err := j.validateSetOffsetSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offsetSeconds",
		val,
	)
}

func (j *jsiiProxy_Offset)SetOffsetYears(val *float64) {
	if err := j.validateSetOffsetYearsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offsetYears",
		val,
	)
}

func (j *jsiiProxy_Offset)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Offset)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Offset)SetTriggers(val *map[string]*string) {
	if err := j.validateSetTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

// Generates CDKTF code for importing a Offset resource upon running "cdktf plan <stack-name>".
func Offset_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOffset_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-time.offset.Offset",
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
func Offset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOffset_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-time.offset.Offset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Offset_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOffset_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-time.offset.Offset",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Offset_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOffset_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-time.offset.Offset",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Offset_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-time.offset.Offset",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_Offset) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_Offset) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_Offset) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Offset) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Offset) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Offset) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Offset) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Offset) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Offset) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Offset) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Offset) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Offset) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Offset) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_Offset) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Offset) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_Offset) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_Offset) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_Offset) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_Offset) ResetBaseRfc3339() {
	_jsii_.InvokeVoid(
		o,
		"resetBaseRfc3339",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Offset) ResetOffsetDays() {
	_jsii_.InvokeVoid(
		o,
		"resetOffsetDays",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Offset) ResetOffsetHours() {
	_jsii_.InvokeVoid(
		o,
		"resetOffsetHours",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Offset) ResetOffsetMinutes() {
	_jsii_.InvokeVoid(
		o,
		"resetOffsetMinutes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Offset) ResetOffsetMonths() {
	_jsii_.InvokeVoid(
		o,
		"resetOffsetMonths",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Offset) ResetOffsetSeconds() {
	_jsii_.InvokeVoid(
		o,
		"resetOffsetSeconds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Offset) ResetOffsetYears() {
	_jsii_.InvokeVoid(
		o,
		"resetOffsetYears",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Offset) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Offset) ResetTriggers() {
	_jsii_.InvokeVoid(
		o,
		"resetTriggers",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Offset) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Offset) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Offset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Offset) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

