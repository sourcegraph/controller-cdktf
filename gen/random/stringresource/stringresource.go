package stringresource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/random/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/random/stringresource/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/random/3.5.1/docs/resources/string random_string}.
type StringResource interface {
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
	Id() *string
	Keepers() *map[string]*string
	SetKeepers(val *map[string]*string)
	KeepersInput() *map[string]*string
	Length() *float64
	SetLength(val *float64)
	LengthInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Lower() interface{}
	SetLower(val interface{})
	LowerInput() interface{}
	MinLower() *float64
	SetMinLower(val *float64)
	MinLowerInput() *float64
	MinNumeric() *float64
	SetMinNumeric(val *float64)
	MinNumericInput() *float64
	MinSpecial() *float64
	SetMinSpecial(val *float64)
	MinSpecialInput() *float64
	MinUpper() *float64
	SetMinUpper(val *float64)
	MinUpperInput() *float64
	// The tree node.
	Node() constructs.Node
	Number() interface{}
	SetNumber(val interface{})
	NumberInput() interface{}
	Numeric() interface{}
	SetNumeric(val interface{})
	NumericInput() interface{}
	OverrideSpecial() *string
	SetOverrideSpecial(val *string)
	OverrideSpecialInput() *string
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
	Result() *string
	Special() interface{}
	SetSpecial(val interface{})
	SpecialInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Upper() interface{}
	SetUpper(val interface{})
	UpperInput() interface{}
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
	ResetKeepers()
	ResetLower()
	ResetMinLower()
	ResetMinNumeric()
	ResetMinSpecial()
	ResetMinUpper()
	ResetNumber()
	ResetNumeric()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOverrideSpecial()
	ResetSpecial()
	ResetUpper()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for StringResource
type jsiiProxy_StringResource struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StringResource) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Keepers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"keepers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) KeepersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"keepersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Length() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"length",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) LengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Lower() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lower",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) LowerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lowerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) MinLower() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLower",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) MinLowerInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLowerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) MinNumeric() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumeric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) MinNumericInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumericInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) MinSpecial() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSpecial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) MinSpecialInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSpecialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) MinUpper() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minUpper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) MinUpperInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minUpperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Number() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"number",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) NumberInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Numeric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numeric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) NumericInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numericInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) OverrideSpecial() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideSpecial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) OverrideSpecialInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideSpecialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Result() *string {
	var returns *string
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Special() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"special",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) SpecialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"specialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) Upper() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringResource) UpperInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upperInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/random/3.5.1/docs/resources/string random_string} Resource.
func NewStringResource(scope constructs.Construct, id *string, config *StringResourceConfig) StringResource {
	_init_.Initialize()

	if err := validateNewStringResourceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StringResource{}

	_jsii_.Create(
		"random.stringResource.StringResource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/random/3.5.1/docs/resources/string random_string} Resource.
func NewStringResource_Override(s StringResource, scope constructs.Construct, id *string, config *StringResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"random.stringResource.StringResource",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StringResource) SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetKeepers(val *map[string]*string) {
	if err := j.validateSetKeepersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepers",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetLength(val *float64) {
	if err := j.validateSetLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"length",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetLower(val interface{}) {
	if err := j.validateSetLowerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lower",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetMinLower(val *float64) {
	if err := j.validateSetMinLowerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLower",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetMinNumeric(val *float64) {
	if err := j.validateSetMinNumericParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNumeric",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetMinSpecial(val *float64) {
	if err := j.validateSetMinSpecialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSpecial",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetMinUpper(val *float64) {
	if err := j.validateSetMinUpperParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minUpper",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetNumber(val interface{}) {
	if err := j.validateSetNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"number",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetNumeric(val interface{}) {
	if err := j.validateSetNumericParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numeric",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetOverrideSpecial(val *string) {
	if err := j.validateSetOverrideSpecialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideSpecial",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetSpecial(val interface{}) {
	if err := j.validateSetSpecialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"special",
		val,
	)
}

func (j *jsiiProxy_StringResource) SetUpper(val interface{}) {
	if err := j.validateSetUpperParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upper",
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
func StringResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStringResource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"random.stringResource.StringResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StringResource_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStringResource_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"random.stringResource.StringResource",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StringResource_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStringResource_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"random.stringResource.StringResource",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StringResource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"random.stringResource.StringResource",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StringResource) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StringResource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringResource) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringResource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringResource) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringResource) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringResource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringResource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringResource) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringResource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringResource) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringResource) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StringResource) ResetKeepers() {
	_jsii_.InvokeVoid(
		s,
		"resetKeepers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringResource) ResetLower() {
	_jsii_.InvokeVoid(
		s,
		"resetLower",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringResource) ResetMinLower() {
	_jsii_.InvokeVoid(
		s,
		"resetMinLower",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringResource) ResetMinNumeric() {
	_jsii_.InvokeVoid(
		s,
		"resetMinNumeric",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringResource) ResetMinSpecial() {
	_jsii_.InvokeVoid(
		s,
		"resetMinSpecial",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringResource) ResetMinUpper() {
	_jsii_.InvokeVoid(
		s,
		"resetMinUpper",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringResource) ResetNumber() {
	_jsii_.InvokeVoid(
		s,
		"resetNumber",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringResource) ResetNumeric() {
	_jsii_.InvokeVoid(
		s,
		"resetNumeric",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringResource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringResource) ResetOverrideSpecial() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideSpecial",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringResource) ResetSpecial() {
	_jsii_.InvokeVoid(
		s,
		"resetSpecial",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringResource) ResetUpper() {
	_jsii_.InvokeVoid(
		s,
		"resetUpper",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StringResource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringResource) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringResource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
