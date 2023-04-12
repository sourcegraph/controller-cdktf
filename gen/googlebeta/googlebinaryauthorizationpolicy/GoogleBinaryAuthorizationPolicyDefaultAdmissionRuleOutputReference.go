package googlebinaryauthorizationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlebinaryauthorizationpolicy/internal"
)

type GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference interface {
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
	EnforcementMode() *string
	SetEnforcementMode(val *string)
	EnforcementModeInput() *string
	EvaluationMode() *string
	SetEvaluationMode(val *string)
	EvaluationModeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBinaryAuthorizationPolicyDefaultAdmissionRule
	SetInternalValue(val *GoogleBinaryAuthorizationPolicyDefaultAdmissionRule)
	RequireAttestationsBy() *[]*string
	SetRequireAttestationsBy(val *[]*string)
	RequireAttestationsByInput() *[]*string
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
	ResetRequireAttestationsBy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference
type jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) EnforcementMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforcementMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) EnforcementModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforcementModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) EvaluationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) EvaluationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) InternalValue() *GoogleBinaryAuthorizationPolicyDefaultAdmissionRule {
	var returns *GoogleBinaryAuthorizationPolicyDefaultAdmissionRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) RequireAttestationsBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requireAttestationsBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) RequireAttestationsByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requireAttestationsByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleBinaryAuthorizationPolicy.GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference_Override(g GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleBinaryAuthorizationPolicy.GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetEnforcementMode(val *string) {
	if err := j.validateSetEnforcementModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforcementMode",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetEvaluationMode(val *string) {
	if err := j.validateSetEvaluationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationMode",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetInternalValue(val *GoogleBinaryAuthorizationPolicyDefaultAdmissionRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetRequireAttestationsBy(val *[]*string) {
	if err := j.validateSetRequireAttestationsByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAttestationsBy",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) ResetRequireAttestationsBy() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireAttestationsBy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

