package googleprivatecacertificatetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleprivatecacertificatetemplate/internal"
)

type GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference interface {
	cdktf.ComplexObject
	AdditionalExtensions() GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsList
	AdditionalExtensionsInput() interface{}
	AiaOcspServers() *[]*string
	SetAiaOcspServers(val *[]*string)
	AiaOcspServersInput() *[]*string
	CaOptions() GooglePrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference
	CaOptionsInput() *GooglePrivatecaCertificateTemplatePredefinedValuesCaOptions
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
	InternalValue() *GooglePrivatecaCertificateTemplatePredefinedValues
	SetInternalValue(val *GooglePrivatecaCertificateTemplatePredefinedValues)
	KeyUsage() GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference
	KeyUsageInput() *GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsage
	PolicyIds() GooglePrivatecaCertificateTemplatePredefinedValuesPolicyIdsList
	PolicyIdsInput() interface{}
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
	PutAdditionalExtensions(value interface{})
	PutCaOptions(value *GooglePrivatecaCertificateTemplatePredefinedValuesCaOptions)
	PutKeyUsage(value *GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsage)
	PutPolicyIds(value interface{})
	ResetAdditionalExtensions()
	ResetAiaOcspServers()
	ResetCaOptions()
	ResetKeyUsage()
	ResetPolicyIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference
type jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) AdditionalExtensions() GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsList {
	var returns GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsList
	_jsii_.Get(
		j,
		"additionalExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) AdditionalExtensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalExtensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) AiaOcspServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aiaOcspServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) AiaOcspServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aiaOcspServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) CaOptions() GooglePrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference {
	var returns GooglePrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference
	_jsii_.Get(
		j,
		"caOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) CaOptionsInput() *GooglePrivatecaCertificateTemplatePredefinedValuesCaOptions {
	var returns *GooglePrivatecaCertificateTemplatePredefinedValuesCaOptions
	_jsii_.Get(
		j,
		"caOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) InternalValue() *GooglePrivatecaCertificateTemplatePredefinedValues {
	var returns *GooglePrivatecaCertificateTemplatePredefinedValues
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) KeyUsage() GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference {
	var returns GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference
	_jsii_.Get(
		j,
		"keyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) KeyUsageInput() *GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsage {
	var returns *GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsage
	_jsii_.Get(
		j,
		"keyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) PolicyIds() GooglePrivatecaCertificateTemplatePredefinedValuesPolicyIdsList {
	var returns GooglePrivatecaCertificateTemplatePredefinedValuesPolicyIdsList
	_jsii_.Get(
		j,
		"policyIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) PolicyIdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGooglePrivatecaCertificateTemplatePredefinedValuesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCertificateTemplatePredefinedValuesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCertificateTemplatePredefinedValuesOutputReference_Override(g GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference)SetAiaOcspServers(val *[]*string) {
	if err := j.validateSetAiaOcspServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aiaOcspServers",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference)SetInternalValue(val *GooglePrivatecaCertificateTemplatePredefinedValues) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) PutAdditionalExtensions(value interface{}) {
	if err := g.validatePutAdditionalExtensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdditionalExtensions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) PutCaOptions(value *GooglePrivatecaCertificateTemplatePredefinedValuesCaOptions) {
	if err := g.validatePutCaOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCaOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) PutKeyUsage(value *GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsage) {
	if err := g.validatePutKeyUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKeyUsage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) PutPolicyIds(value interface{}) {
	if err := g.validatePutPolicyIdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPolicyIds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) ResetAdditionalExtensions() {
	_jsii_.InvokeVoid(
		g,
		"resetAdditionalExtensions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) ResetAiaOcspServers() {
	_jsii_.InvokeVoid(
		g,
		"resetAiaOcspServers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) ResetCaOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetCaOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) ResetKeyUsage() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyUsage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) ResetPolicyIds() {
	_jsii_.InvokeVoid(
		g,
		"resetPolicyIds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

