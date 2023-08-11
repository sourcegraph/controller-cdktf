package googleiamaccessboundarypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleiamaccessboundarypolicy/internal"
)

type GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference interface {
	cdktf.ComplexObject
	AvailabilityCondition() GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionOutputReference
	AvailabilityConditionInput() *GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition
	AvailablePermissions() *[]*string
	SetAvailablePermissions(val *[]*string)
	AvailablePermissionsInput() *[]*string
	AvailableResource() *string
	SetAvailableResource(val *string)
	AvailableResourceInput() *string
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
	InternalValue() *GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRule
	SetInternalValue(val *GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRule)
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
	PutAvailabilityCondition(value *GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition)
	ResetAvailabilityCondition()
	ResetAvailablePermissions()
	ResetAvailableResource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference
type jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) AvailabilityCondition() GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionOutputReference {
	var returns GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionOutputReference
	_jsii_.Get(
		j,
		"availabilityCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) AvailabilityConditionInput() *GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition {
	var returns *GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition
	_jsii_.Get(
		j,
		"availabilityConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) AvailablePermissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availablePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) AvailablePermissionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availablePermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) AvailableResource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availableResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) AvailableResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availableResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) InternalValue() *GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRule {
	var returns *GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIamAccessBoundaryPolicy.GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference_Override(g GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIamAccessBoundaryPolicy.GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference)SetAvailablePermissions(val *[]*string) {
	if err := j.validateSetAvailablePermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availablePermissions",
		val,
	)
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference)SetAvailableResource(val *string) {
	if err := j.validateSetAvailableResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableResource",
		val,
	)
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference)SetInternalValue(val *GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) PutAvailabilityCondition(value *GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition) {
	if err := g.validatePutAvailabilityConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAvailabilityCondition",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) ResetAvailabilityCondition() {
	_jsii_.InvokeVoid(
		g,
		"resetAvailabilityCondition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) ResetAvailablePermissions() {
	_jsii_.InvokeVoid(
		g,
		"resetAvailablePermissions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) ResetAvailableResource() {
	_jsii_.InvokeVoid(
		g,
		"resetAvailableResource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

