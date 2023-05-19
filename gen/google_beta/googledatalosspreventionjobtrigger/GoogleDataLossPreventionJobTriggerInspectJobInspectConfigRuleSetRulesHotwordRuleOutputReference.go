package googledatalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventionjobtrigger/internal"
)

type GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference interface {
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
	HotwordRegex() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexOutputReference
	HotwordRegexInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex
	InternalValue() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule
	SetInternalValue(val *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule)
	LikelihoodAdjustment() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentOutputReference
	LikelihoodAdjustmentInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment
	Proximity() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityOutputReference
	ProximityInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity
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
	PutHotwordRegex(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex)
	PutLikelihoodAdjustment(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment)
	PutProximity(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity)
	ResetHotwordRegex()
	ResetLikelihoodAdjustment()
	ResetProximity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference
type jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) HotwordRegex() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexOutputReference
	_jsii_.Get(
		j,
		"hotwordRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) HotwordRegexInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex
	_jsii_.Get(
		j,
		"hotwordRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) InternalValue() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) LikelihoodAdjustment() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentOutputReference
	_jsii_.Get(
		j,
		"likelihoodAdjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) LikelihoodAdjustmentInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment
	_jsii_.Get(
		j,
		"likelihoodAdjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) Proximity() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityOutputReference
	_jsii_.Get(
		j,
		"proximity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) ProximityInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity
	_jsii_.Get(
		j,
		"proximityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference_Override(g GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference)SetInternalValue(val *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) PutHotwordRegex(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) {
	if err := g.validatePutHotwordRegexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHotwordRegex",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) PutLikelihoodAdjustment(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) {
	if err := g.validatePutLikelihoodAdjustmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLikelihoodAdjustment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) PutProximity(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) {
	if err := g.validatePutProximityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProximity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) ResetHotwordRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetHotwordRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) ResetLikelihoodAdjustment() {
	_jsii_.InvokeVoid(
		g,
		"resetLikelihoodAdjustment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) ResetProximity() {
	_jsii_.InvokeVoid(
		g,
		"resetProximity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

