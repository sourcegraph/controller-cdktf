package googlecloudsecuritycompliancecloudcontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudsecuritycompliancecloudcontrol/internal"
)

type GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference interface {
	cdktf.ComplexObject
	AllowedValues() GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesOutputReference
	AllowedValuesInput() *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValues
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
	InternalValue() *GoogleCloudSecurityComplianceCloudControlParameterSpecValidation
	SetInternalValue(val *GoogleCloudSecurityComplianceCloudControlParameterSpecValidation)
	IntRange() GoogleCloudSecurityComplianceCloudControlParameterSpecValidationIntRangeOutputReference
	IntRangeInput() *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationIntRange
	RegexpPattern() GoogleCloudSecurityComplianceCloudControlParameterSpecValidationRegexpPatternOutputReference
	RegexpPatternInput() *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationRegexpPattern
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
	PutAllowedValues(value *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValues)
	PutIntRange(value *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationIntRange)
	PutRegexpPattern(value *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationRegexpPattern)
	ResetAllowedValues()
	ResetIntRange()
	ResetRegexpPattern()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference
type jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) AllowedValues() GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesOutputReference {
	var returns GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesOutputReference
	_jsii_.Get(
		j,
		"allowedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) AllowedValuesInput() *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValues {
	var returns *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValues
	_jsii_.Get(
		j,
		"allowedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) InternalValue() *GoogleCloudSecurityComplianceCloudControlParameterSpecValidation {
	var returns *GoogleCloudSecurityComplianceCloudControlParameterSpecValidation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) IntRange() GoogleCloudSecurityComplianceCloudControlParameterSpecValidationIntRangeOutputReference {
	var returns GoogleCloudSecurityComplianceCloudControlParameterSpecValidationIntRangeOutputReference
	_jsii_.Get(
		j,
		"intRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) IntRangeInput() *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationIntRange {
	var returns *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationIntRange
	_jsii_.Get(
		j,
		"intRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) RegexpPattern() GoogleCloudSecurityComplianceCloudControlParameterSpecValidationRegexpPatternOutputReference {
	var returns GoogleCloudSecurityComplianceCloudControlParameterSpecValidationRegexpPatternOutputReference
	_jsii_.Get(
		j,
		"regexpPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) RegexpPatternInput() *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationRegexpPattern {
	var returns *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationRegexpPattern
	_jsii_.Get(
		j,
		"regexpPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudSecurityComplianceCloudControl.GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference_Override(g GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudSecurityComplianceCloudControl.GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference)SetInternalValue(val *GoogleCloudSecurityComplianceCloudControlParameterSpecValidation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) PutAllowedValues(value *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValues) {
	if err := g.validatePutAllowedValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAllowedValues",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) PutIntRange(value *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationIntRange) {
	if err := g.validatePutIntRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIntRange",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) PutRegexpPattern(value *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationRegexpPattern) {
	if err := g.validatePutRegexpPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRegexpPattern",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) ResetAllowedValues() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) ResetIntRange() {
	_jsii_.InvokeVoid(
		g,
		"resetIntRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) ResetRegexpPattern() {
	_jsii_.InvokeVoid(
		g,
		"resetRegexpPattern",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

