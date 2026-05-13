package googlelustreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlelustreinstance/internal"
)

type GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Hours() *float64
	SetHours(val *float64)
	HoursInput() *float64
	InternalValue() *GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime
	SetInternalValue(val *GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime)
	Minutes() *float64
	SetMinutes(val *float64)
	MinutesInput() *float64
	Nanos() *float64
	SetNanos(val *float64)
	NanosInput() *float64
	Seconds() *float64
	SetSeconds(val *float64)
	SecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetHours()
	ResetMinutes()
	ResetNanos()
	ResetSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference
type jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) Hours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) HoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) InternalValue() *GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime {
	var returns *GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) Minutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) MinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) Nanos() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nanos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) NanosInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nanosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) Seconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"seconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) SecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleLustreInstance.GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference_Override(g GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleLustreInstance.GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference)SetHours(val *float64) {
	if err := j.validateSetHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hours",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference)SetInternalValue(val *GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference)SetMinutes(val *float64) {
	if err := j.validateSetMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minutes",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference)SetNanos(val *float64) {
	if err := j.validateSetNanosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nanos",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference)SetSeconds(val *float64) {
	if err := j.validateSetSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seconds",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ResetHours() {
	_jsii_.InvokeVoid(
		g,
		"resetHours",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ResetMinutes() {
	_jsii_.InvokeVoid(
		g,
		"resetMinutes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ResetNanos() {
	_jsii_.InvokeVoid(
		g,
		"resetNanos",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ResetSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

