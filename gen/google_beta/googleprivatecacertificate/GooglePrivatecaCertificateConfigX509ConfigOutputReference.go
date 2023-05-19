package googleprivatecacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleprivatecacertificate/internal"
)

type GooglePrivatecaCertificateConfigX509ConfigOutputReference interface {
	cdktf.ComplexObject
	AdditionalExtensions() GooglePrivatecaCertificateConfigX509ConfigAdditionalExtensionsList
	AdditionalExtensionsInput() interface{}
	AiaOcspServers() *[]*string
	SetAiaOcspServers(val *[]*string)
	AiaOcspServersInput() *[]*string
	CaOptions() GooglePrivatecaCertificateConfigX509ConfigCaOptionsOutputReference
	CaOptionsInput() *GooglePrivatecaCertificateConfigX509ConfigCaOptions
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
	InternalValue() *GooglePrivatecaCertificateConfigX509Config
	SetInternalValue(val *GooglePrivatecaCertificateConfigX509Config)
	KeyUsage() GooglePrivatecaCertificateConfigX509ConfigKeyUsageOutputReference
	KeyUsageInput() *GooglePrivatecaCertificateConfigX509ConfigKeyUsage
	NameConstraints() GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference
	NameConstraintsInput() *GooglePrivatecaCertificateConfigX509ConfigNameConstraints
	PolicyIds() GooglePrivatecaCertificateConfigX509ConfigPolicyIdsList
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
	PutCaOptions(value *GooglePrivatecaCertificateConfigX509ConfigCaOptions)
	PutKeyUsage(value *GooglePrivatecaCertificateConfigX509ConfigKeyUsage)
	PutNameConstraints(value *GooglePrivatecaCertificateConfigX509ConfigNameConstraints)
	PutPolicyIds(value interface{})
	ResetAdditionalExtensions()
	ResetAiaOcspServers()
	ResetCaOptions()
	ResetNameConstraints()
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

// The jsii proxy struct for GooglePrivatecaCertificateConfigX509ConfigOutputReference
type jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) AdditionalExtensions() GooglePrivatecaCertificateConfigX509ConfigAdditionalExtensionsList {
	var returns GooglePrivatecaCertificateConfigX509ConfigAdditionalExtensionsList
	_jsii_.Get(
		j,
		"additionalExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) AdditionalExtensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalExtensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) AiaOcspServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aiaOcspServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) AiaOcspServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aiaOcspServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) CaOptions() GooglePrivatecaCertificateConfigX509ConfigCaOptionsOutputReference {
	var returns GooglePrivatecaCertificateConfigX509ConfigCaOptionsOutputReference
	_jsii_.Get(
		j,
		"caOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) CaOptionsInput() *GooglePrivatecaCertificateConfigX509ConfigCaOptions {
	var returns *GooglePrivatecaCertificateConfigX509ConfigCaOptions
	_jsii_.Get(
		j,
		"caOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) InternalValue() *GooglePrivatecaCertificateConfigX509Config {
	var returns *GooglePrivatecaCertificateConfigX509Config
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) KeyUsage() GooglePrivatecaCertificateConfigX509ConfigKeyUsageOutputReference {
	var returns GooglePrivatecaCertificateConfigX509ConfigKeyUsageOutputReference
	_jsii_.Get(
		j,
		"keyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) KeyUsageInput() *GooglePrivatecaCertificateConfigX509ConfigKeyUsage {
	var returns *GooglePrivatecaCertificateConfigX509ConfigKeyUsage
	_jsii_.Get(
		j,
		"keyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) NameConstraints() GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference {
	var returns GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference
	_jsii_.Get(
		j,
		"nameConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) NameConstraintsInput() *GooglePrivatecaCertificateConfigX509ConfigNameConstraints {
	var returns *GooglePrivatecaCertificateConfigX509ConfigNameConstraints
	_jsii_.Get(
		j,
		"nameConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) PolicyIds() GooglePrivatecaCertificateConfigX509ConfigPolicyIdsList {
	var returns GooglePrivatecaCertificateConfigX509ConfigPolicyIdsList
	_jsii_.Get(
		j,
		"policyIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) PolicyIdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGooglePrivatecaCertificateConfigX509ConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePrivatecaCertificateConfigX509ConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCertificateConfigX509ConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googlePrivatecaCertificate.GooglePrivatecaCertificateConfigX509ConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCertificateConfigX509ConfigOutputReference_Override(g GooglePrivatecaCertificateConfigX509ConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googlePrivatecaCertificate.GooglePrivatecaCertificateConfigX509ConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference)SetAiaOcspServers(val *[]*string) {
	if err := j.validateSetAiaOcspServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aiaOcspServers",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference)SetInternalValue(val *GooglePrivatecaCertificateConfigX509Config) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) PutAdditionalExtensions(value interface{}) {
	if err := g.validatePutAdditionalExtensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdditionalExtensions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) PutCaOptions(value *GooglePrivatecaCertificateConfigX509ConfigCaOptions) {
	if err := g.validatePutCaOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCaOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) PutKeyUsage(value *GooglePrivatecaCertificateConfigX509ConfigKeyUsage) {
	if err := g.validatePutKeyUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKeyUsage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) PutNameConstraints(value *GooglePrivatecaCertificateConfigX509ConfigNameConstraints) {
	if err := g.validatePutNameConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNameConstraints",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) PutPolicyIds(value interface{}) {
	if err := g.validatePutPolicyIdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPolicyIds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) ResetAdditionalExtensions() {
	_jsii_.InvokeVoid(
		g,
		"resetAdditionalExtensions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) ResetAiaOcspServers() {
	_jsii_.InvokeVoid(
		g,
		"resetAiaOcspServers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) ResetCaOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetCaOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) ResetNameConstraints() {
	_jsii_.InvokeVoid(
		g,
		"resetNameConstraints",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) ResetPolicyIds() {
	_jsii_.InvokeVoid(
		g,
		"resetPolicyIds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

