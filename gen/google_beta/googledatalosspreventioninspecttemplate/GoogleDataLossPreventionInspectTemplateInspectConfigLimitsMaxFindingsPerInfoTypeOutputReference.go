package googledatalosspreventioninspecttemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventioninspecttemplate/internal"
)

type GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference interface {
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
	InfoType() GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeOutputReference
	InfoTypeInput() *GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxFindings() *float64
	SetMaxFindings(val *float64)
	MaxFindingsInput() *float64
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
	PutInfoType(value *GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType)
	ResetInfoType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference
type jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) InfoType() GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeOutputReference {
	var returns GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeOutputReference
	_jsii_.Get(
		j,
		"infoType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) InfoTypeInput() *GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	var returns *GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType
	_jsii_.Get(
		j,
		"infoTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) MaxFindings() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) MaxFindingsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataLossPreventionInspectTemplate.GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference_Override(g GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataLossPreventionInspectTemplate.GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference)SetMaxFindings(val *float64) {
	if err := j.validateSetMaxFindingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFindings",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) PutInfoType(value *GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) {
	if err := g.validatePutInfoTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInfoType",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) ResetInfoType() {
	_jsii_.InvokeVoid(
		g,
		"resetInfoType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

