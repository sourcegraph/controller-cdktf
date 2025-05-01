package googleclouddeploydeploypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleclouddeploydeploypolicy/internal"
)

type GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference interface {
	cdktf.ComplexObject
	Actions() *[]*string
	SetActions(val *[]*string)
	ActionsInput() *[]*string
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() *GoogleClouddeployDeployPolicyRulesRolloutRestriction
	SetInternalValue(val *GoogleClouddeployDeployPolicyRulesRolloutRestriction)
	Invokers() *[]*string
	SetInvokers(val *[]*string)
	InvokersInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeWindows() GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference
	TimeWindowsInput() *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows
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
	PutTimeWindows(value *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows)
	ResetActions()
	ResetInvokers()
	ResetTimeWindows()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference
type jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) Actions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) ActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) InternalValue() *GoogleClouddeployDeployPolicyRulesRolloutRestriction {
	var returns *GoogleClouddeployDeployPolicyRulesRolloutRestriction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) Invokers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invokers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) InvokersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invokersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) TimeWindows() GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference {
	var returns GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference
	_jsii_.Get(
		j,
		"timeWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) TimeWindowsInput() *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows {
	var returns *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows
	_jsii_.Get(
		j,
		"timeWindowsInput",
		&returns,
	)
	return returns
}


func NewGoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployDeployPolicy.GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference_Override(g GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployDeployPolicy.GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference)SetActions(val *[]*string) {
	if err := j.validateSetActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference)SetInternalValue(val *GoogleClouddeployDeployPolicyRulesRolloutRestriction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference)SetInvokers(val *[]*string) {
	if err := j.validateSetInvokersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invokers",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) PutTimeWindows(value *GoogleClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows) {
	if err := g.validatePutTimeWindowsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeWindows",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) ResetActions() {
	_jsii_.InvokeVoid(
		g,
		"resetActions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) ResetInvokers() {
	_jsii_.InvokeVoid(
		g,
		"resetInvokers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) ResetTimeWindows() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeWindows",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleClouddeployDeployPolicyRulesRolloutRestrictionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

