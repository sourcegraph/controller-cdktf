package googlelookerinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlelookerinstance/internal"
)

type GoogleLookerInstanceDenyMaintenancePeriodOutputReference interface {
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
	EndDate() GoogleLookerInstanceDenyMaintenancePeriodEndDateOutputReference
	EndDateInput() *GoogleLookerInstanceDenyMaintenancePeriodEndDate
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleLookerInstanceDenyMaintenancePeriod
	SetInternalValue(val *GoogleLookerInstanceDenyMaintenancePeriod)
	StartDate() GoogleLookerInstanceDenyMaintenancePeriodStartDateOutputReference
	StartDateInput() *GoogleLookerInstanceDenyMaintenancePeriodStartDate
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Time() GoogleLookerInstanceDenyMaintenancePeriodTimeOutputReference
	TimeInput() *GoogleLookerInstanceDenyMaintenancePeriodTime
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
	PutEndDate(value *GoogleLookerInstanceDenyMaintenancePeriodEndDate)
	PutStartDate(value *GoogleLookerInstanceDenyMaintenancePeriodStartDate)
	PutTime(value *GoogleLookerInstanceDenyMaintenancePeriodTime)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleLookerInstanceDenyMaintenancePeriodOutputReference
type jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) EndDate() GoogleLookerInstanceDenyMaintenancePeriodEndDateOutputReference {
	var returns GoogleLookerInstanceDenyMaintenancePeriodEndDateOutputReference
	_jsii_.Get(
		j,
		"endDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) EndDateInput() *GoogleLookerInstanceDenyMaintenancePeriodEndDate {
	var returns *GoogleLookerInstanceDenyMaintenancePeriodEndDate
	_jsii_.Get(
		j,
		"endDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) InternalValue() *GoogleLookerInstanceDenyMaintenancePeriod {
	var returns *GoogleLookerInstanceDenyMaintenancePeriod
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) StartDate() GoogleLookerInstanceDenyMaintenancePeriodStartDateOutputReference {
	var returns GoogleLookerInstanceDenyMaintenancePeriodStartDateOutputReference
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) StartDateInput() *GoogleLookerInstanceDenyMaintenancePeriodStartDate {
	var returns *GoogleLookerInstanceDenyMaintenancePeriodStartDate
	_jsii_.Get(
		j,
		"startDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) Time() GoogleLookerInstanceDenyMaintenancePeriodTimeOutputReference {
	var returns GoogleLookerInstanceDenyMaintenancePeriodTimeOutputReference
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) TimeInput() *GoogleLookerInstanceDenyMaintenancePeriodTime {
	var returns *GoogleLookerInstanceDenyMaintenancePeriodTime
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}


func NewGoogleLookerInstanceDenyMaintenancePeriodOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleLookerInstanceDenyMaintenancePeriodOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleLookerInstanceDenyMaintenancePeriodOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference{}

	_jsii_.Create(
		"google-beta.googleLookerInstance.GoogleLookerInstanceDenyMaintenancePeriodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleLookerInstanceDenyMaintenancePeriodOutputReference_Override(g GoogleLookerInstanceDenyMaintenancePeriodOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleLookerInstance.GoogleLookerInstanceDenyMaintenancePeriodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference)SetInternalValue(val *GoogleLookerInstanceDenyMaintenancePeriod) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) PutEndDate(value *GoogleLookerInstanceDenyMaintenancePeriodEndDate) {
	if err := g.validatePutEndDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEndDate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) PutStartDate(value *GoogleLookerInstanceDenyMaintenancePeriodStartDate) {
	if err := g.validatePutStartDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStartDate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) PutTime(value *GoogleLookerInstanceDenyMaintenancePeriodTime) {
	if err := g.validatePutTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTime",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleLookerInstanceDenyMaintenancePeriodOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

