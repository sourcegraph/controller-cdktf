package appmeshvirtualnode

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/appmeshvirtualnode/internal"
)

type AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference interface {
	cdktf.ComplexObject
	Certificate() AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference
	CertificateInput() *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	Enforce() any
	SetEnforce(val any)
	EnforceInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTls
	SetInternalValue(val *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTls)
	Ports() *[]*float64
	SetPorts(val *[]*float64)
	PortsInput() *[]*float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Validation() AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationOutputReference
	ValidationInput() *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidation
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutCertificate(value *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate)
	PutValidation(value *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidation)
	ResetCertificate()
	ResetEnforce()
	ResetPorts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference
type jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) Certificate() AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference {
	var returns AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) CertificateInput() *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate {
	var returns *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) Enforce() any {
	var returns any
	_jsii_.Get(
		j,
		"enforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) EnforceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) InternalValue() *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTls {
	var returns *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTls
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) Ports() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) PortsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"portsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) Validation() AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationOutputReference {
	var returns AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationOutputReference
	_jsii_.Get(
		j,
		"validation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) ValidationInput() *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidation {
	var returns *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidation
	_jsii_.Get(
		j,
		"validationInput",
		&returns,
	)
	return returns
}

func NewAppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference {
	_init_.Initialize()

	if err := validateNewAppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualNode.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference_Override(a AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualNode.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference",
		[]any{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) SetEnforce(val any) {
	if err := j.validateSetEnforceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforce",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) SetInternalValue(val *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTls) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) SetPorts(val *[]*float64) {
	if err := j.validateSetPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ports",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) PutCertificate(value *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate) {
	if err := a.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCertificate",
		[]any{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) PutValidation(value *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidation) {
	if err := a.validatePutValidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putValidation",
		[]any{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) ResetEnforce() {
	_jsii_.InvokeVoid(
		a,
		"resetEnforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) ResetPorts() {
	_jsii_.InvokeVoid(
		a,
		"resetPorts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		a,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
