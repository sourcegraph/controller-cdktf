package googlemonitoringuptimecheckconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemonitoringuptimecheckconfig/internal"
)

type GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StatusClass() *string
	SetStatusClass(val *string)
	StatusClassInput() *string
	StatusValue() *float64
	SetStatusValue(val *float64)
	StatusValueInput() *float64
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
	ResetStatusClass()
	ResetStatusValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference
type jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) StatusClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) StatusClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) StatusValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) StatusValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference_Override(g GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference)SetStatusClass(val *string) {
	if err := j.validateSetStatusClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statusClass",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference)SetStatusValue(val *float64) {
	if err := j.validateSetStatusValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statusValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) ResetStatusClass() {
	_jsii_.InvokeVoid(
		g,
		"resetStatusClass",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) ResetStatusValue() {
	_jsii_.InvokeVoid(
		g,
		"resetStatusValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

