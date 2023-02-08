package googledatalosspreventioninspecttemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventioninspecttemplate/internal"
)

type GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference interface {
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
	InternalValue() *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity
	SetInternalValue(val *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowAfter() *float64
	SetWindowAfter(val *float64)
	WindowAfterInput() *float64
	WindowBefore() *float64
	SetWindowBefore(val *float64)
	WindowBeforeInput() *float64
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
	ResetWindowAfter()
	ResetWindowBefore()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference
type jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) InternalValue() *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
	var returns *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) WindowAfter() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"windowAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) WindowAfterInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"windowAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) WindowBefore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"windowBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) WindowBeforeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"windowBeforeInput",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataLossPreventionInspectTemplate.GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference_Override(g GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataLossPreventionInspectTemplate.GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference)SetInternalValue(val *GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference)SetWindowAfter(val *float64) {
	if err := j.validateSetWindowAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowAfter",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference)SetWindowBefore(val *float64) {
	if err := j.validateSetWindowBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowBefore",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) ResetWindowAfter() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowAfter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) ResetWindowBefore() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowBefore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

