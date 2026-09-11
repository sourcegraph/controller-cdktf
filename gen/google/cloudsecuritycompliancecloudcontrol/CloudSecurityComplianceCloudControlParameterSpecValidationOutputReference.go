package cloudsecuritycompliancecloudcontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/cloudsecuritycompliancecloudcontrol/internal"
)

type CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference interface {
	cdktn.ComplexObject
	AllowedValues() CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesOutputReference
	AllowedValuesInput() *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValues
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
	InternalValue() *CloudSecurityComplianceCloudControlParameterSpecValidation
	SetInternalValue(val *CloudSecurityComplianceCloudControlParameterSpecValidation)
	IntRange() CloudSecurityComplianceCloudControlParameterSpecValidationIntRangeOutputReference
	IntRangeInput() *CloudSecurityComplianceCloudControlParameterSpecValidationIntRange
	RegexpPattern() CloudSecurityComplianceCloudControlParameterSpecValidationRegexpPatternOutputReference
	RegexpPatternInput() *CloudSecurityComplianceCloudControlParameterSpecValidationRegexpPattern
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutAllowedValues(value *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValues)
	PutIntRange(value *CloudSecurityComplianceCloudControlParameterSpecValidationIntRange)
	PutRegexpPattern(value *CloudSecurityComplianceCloudControlParameterSpecValidationRegexpPattern)
	ResetAllowedValues()
	ResetIntRange()
	ResetRegexpPattern()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference
type jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) AllowedValues() CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesOutputReference {
	var returns CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesOutputReference
	_jsii_.Get(
		j,
		"allowedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) AllowedValuesInput() *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValues {
	var returns *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValues
	_jsii_.Get(
		j,
		"allowedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) InternalValue() *CloudSecurityComplianceCloudControlParameterSpecValidation {
	var returns *CloudSecurityComplianceCloudControlParameterSpecValidation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) IntRange() CloudSecurityComplianceCloudControlParameterSpecValidationIntRangeOutputReference {
	var returns CloudSecurityComplianceCloudControlParameterSpecValidationIntRangeOutputReference
	_jsii_.Get(
		j,
		"intRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) IntRangeInput() *CloudSecurityComplianceCloudControlParameterSpecValidationIntRange {
	var returns *CloudSecurityComplianceCloudControlParameterSpecValidationIntRange
	_jsii_.Get(
		j,
		"intRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) RegexpPattern() CloudSecurityComplianceCloudControlParameterSpecValidationRegexpPatternOutputReference {
	var returns CloudSecurityComplianceCloudControlParameterSpecValidationRegexpPatternOutputReference
	_jsii_.Get(
		j,
		"regexpPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) RegexpPatternInput() *CloudSecurityComplianceCloudControlParameterSpecValidationRegexpPattern {
	var returns *CloudSecurityComplianceCloudControlParameterSpecValidationRegexpPattern
	_jsii_.Get(
		j,
		"regexpPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference {
	_init_.Initialize()

	if err := validateNewCloudSecurityComplianceCloudControlParameterSpecValidationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cloudSecurityComplianceCloudControl.CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudSecurityComplianceCloudControlParameterSpecValidationOutputReference_Override(c CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cloudSecurityComplianceCloudControl.CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference)SetInternalValue(val *CloudSecurityComplianceCloudControlParameterSpecValidation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) PutAllowedValues(value *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValues) {
	if err := c.validatePutAllowedValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAllowedValues",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) PutIntRange(value *CloudSecurityComplianceCloudControlParameterSpecValidationIntRange) {
	if err := c.validatePutIntRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIntRange",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) PutRegexpPattern(value *CloudSecurityComplianceCloudControlParameterSpecValidationRegexpPattern) {
	if err := c.validatePutRegexpPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRegexpPattern",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) ResetAllowedValues() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) ResetIntRange() {
	_jsii_.InvokeVoid(
		c,
		"resetIntRange",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) ResetRegexpPattern() {
	_jsii_.InvokeVoid(
		c,
		"resetRegexpPattern",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

