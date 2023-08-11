package googleprivatecacertificatetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleprivatecacertificatetemplate/internal"
)

type GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference interface {
	cdktf.ComplexObject
	CertSign() interface{}
	SetCertSign(val interface{})
	CertSignInput() interface{}
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
	ContentCommitment() interface{}
	SetContentCommitment(val interface{})
	ContentCommitmentInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CrlSign() interface{}
	SetCrlSign(val interface{})
	CrlSignInput() interface{}
	DataEncipherment() interface{}
	SetDataEncipherment(val interface{})
	DataEnciphermentInput() interface{}
	DecipherOnly() interface{}
	SetDecipherOnly(val interface{})
	DecipherOnlyInput() interface{}
	DigitalSignature() interface{}
	SetDigitalSignature(val interface{})
	DigitalSignatureInput() interface{}
	EncipherOnly() interface{}
	SetEncipherOnly(val interface{})
	EncipherOnlyInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage
	SetInternalValue(val *GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage)
	KeyAgreement() interface{}
	SetKeyAgreement(val interface{})
	KeyAgreementInput() interface{}
	KeyEncipherment() interface{}
	SetKeyEncipherment(val interface{})
	KeyEnciphermentInput() interface{}
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
	ResetCertSign()
	ResetContentCommitment()
	ResetCrlSign()
	ResetDataEncipherment()
	ResetDecipherOnly()
	ResetDigitalSignature()
	ResetEncipherOnly()
	ResetKeyAgreement()
	ResetKeyEncipherment()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference
type jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) CertSign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certSign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) CertSignInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certSignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ContentCommitment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentCommitment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ContentCommitmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentCommitmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) CrlSign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crlSign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) CrlSignInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crlSignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) DataEncipherment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataEncipherment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) DataEnciphermentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataEnciphermentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) DecipherOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decipherOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) DecipherOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decipherOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) DigitalSignature() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"digitalSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) DigitalSignatureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"digitalSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) EncipherOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encipherOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) EncipherOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encipherOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) InternalValue() *GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	var returns *GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) KeyAgreement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyAgreement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) KeyAgreementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyAgreementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) KeyEncipherment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyEncipherment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) KeyEnciphermentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyEnciphermentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference_Override(g GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)SetCertSign(val interface{}) {
	if err := j.validateSetCertSignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certSign",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)SetContentCommitment(val interface{}) {
	if err := j.validateSetContentCommitmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentCommitment",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)SetCrlSign(val interface{}) {
	if err := j.validateSetCrlSignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crlSign",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)SetDataEncipherment(val interface{}) {
	if err := j.validateSetDataEnciphermentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataEncipherment",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)SetDecipherOnly(val interface{}) {
	if err := j.validateSetDecipherOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decipherOnly",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)SetDigitalSignature(val interface{}) {
	if err := j.validateSetDigitalSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digitalSignature",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)SetEncipherOnly(val interface{}) {
	if err := j.validateSetEncipherOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encipherOnly",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)SetInternalValue(val *GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)SetKeyAgreement(val interface{}) {
	if err := j.validateSetKeyAgreementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyAgreement",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)SetKeyEncipherment(val interface{}) {
	if err := j.validateSetKeyEnciphermentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyEncipherment",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ResetCertSign() {
	_jsii_.InvokeVoid(
		g,
		"resetCertSign",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ResetContentCommitment() {
	_jsii_.InvokeVoid(
		g,
		"resetContentCommitment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ResetCrlSign() {
	_jsii_.InvokeVoid(
		g,
		"resetCrlSign",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ResetDataEncipherment() {
	_jsii_.InvokeVoid(
		g,
		"resetDataEncipherment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ResetDecipherOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetDecipherOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ResetDigitalSignature() {
	_jsii_.InvokeVoid(
		g,
		"resetDigitalSignature",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ResetEncipherOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetEncipherOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ResetKeyAgreement() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyAgreement",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ResetKeyEncipherment() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyEncipherment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

