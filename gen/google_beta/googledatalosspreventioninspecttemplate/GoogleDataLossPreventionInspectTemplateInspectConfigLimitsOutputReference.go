package googledatalosspreventioninspecttemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventioninspecttemplate/internal"
)

type GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference interface {
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
	InternalValue() *GoogleDataLossPreventionInspectTemplateInspectConfigLimits
	SetInternalValue(val *GoogleDataLossPreventionInspectTemplateInspectConfigLimits)
	MaxFindingsPerInfoType() GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeList
	MaxFindingsPerInfoTypeInput() interface{}
	MaxFindingsPerItem() *float64
	SetMaxFindingsPerItem(val *float64)
	MaxFindingsPerItemInput() *float64
	MaxFindingsPerRequest() *float64
	SetMaxFindingsPerRequest(val *float64)
	MaxFindingsPerRequestInput() *float64
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
	PutMaxFindingsPerInfoType(value interface{})
	ResetMaxFindingsPerInfoType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference
type jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) InternalValue() *GoogleDataLossPreventionInspectTemplateInspectConfigLimits {
	var returns *GoogleDataLossPreventionInspectTemplateInspectConfigLimits
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) MaxFindingsPerInfoType() GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeList {
	var returns GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeList
	_jsii_.Get(
		j,
		"maxFindingsPerInfoType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) MaxFindingsPerInfoTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maxFindingsPerInfoTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) MaxFindingsPerItem() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFindingsPerItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) MaxFindingsPerItemInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFindingsPerItemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) MaxFindingsPerRequest() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFindingsPerRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) MaxFindingsPerRequestInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFindingsPerRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataLossPreventionInspectTemplate.GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference_Override(g GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataLossPreventionInspectTemplate.GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference)SetInternalValue(val *GoogleDataLossPreventionInspectTemplateInspectConfigLimits) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference)SetMaxFindingsPerItem(val *float64) {
	if err := j.validateSetMaxFindingsPerItemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFindingsPerItem",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference)SetMaxFindingsPerRequest(val *float64) {
	if err := j.validateSetMaxFindingsPerRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFindingsPerRequest",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) PutMaxFindingsPerInfoType(value interface{}) {
	if err := g.validatePutMaxFindingsPerInfoTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaxFindingsPerInfoType",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) ResetMaxFindingsPerInfoType() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxFindingsPerInfoType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

