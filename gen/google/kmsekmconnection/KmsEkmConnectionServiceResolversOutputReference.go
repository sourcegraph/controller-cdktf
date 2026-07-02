package kmsekmconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/kmsekmconnection/internal"
)

type KmsEkmConnectionServiceResolversOutputReference interface {
	cdktf.ComplexObject
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
	EndpointFilter() *string
	SetEndpointFilter(val *string)
	EndpointFilterInput() *string
	// Experimental.
	Fqn() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	InternalValue() any
	SetInternalValue(val any)
	ServerCertificates() KmsEkmConnectionServiceResolversServerCertificatesList
	ServerCertificatesInput() any
	ServiceDirectoryService() *string
	SetServiceDirectoryService(val *string)
	ServiceDirectoryServiceInput() *string
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
	PutServerCertificates(value any)
	ResetEndpointFilter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KmsEkmConnectionServiceResolversOutputReference
type jsiiProxy_KmsEkmConnectionServiceResolversOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) EndpointFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) EndpointFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) ServerCertificates() KmsEkmConnectionServiceResolversServerCertificatesList {
	var returns KmsEkmConnectionServiceResolversServerCertificatesList
	_jsii_.Get(
		j,
		"serverCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) ServerCertificatesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"serverCertificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) ServiceDirectoryService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceDirectoryService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) ServiceDirectoryServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceDirectoryServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewKmsEkmConnectionServiceResolversOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KmsEkmConnectionServiceResolversOutputReference {
	_init_.Initialize()

	if err := validateNewKmsEkmConnectionServiceResolversOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KmsEkmConnectionServiceResolversOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.kmsEkmConnection.KmsEkmConnectionServiceResolversOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKmsEkmConnectionServiceResolversOutputReference_Override(k KmsEkmConnectionServiceResolversOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.kmsEkmConnection.KmsEkmConnectionServiceResolversOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) SetEndpointFilter(val *string) {
	if err := j.validateSetEndpointFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointFilter",
		val,
	)
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) SetServiceDirectoryService(val *string) {
	if err := j.validateSetServiceDirectoryServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceDirectoryService",
		val,
	)
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) PutServerCertificates(value any) {
	if err := k.validatePutServerCertificatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putServerCertificates",
		[]any{value},
	)
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) ResetEndpointFilter() {
	_jsii_.InvokeVoid(
		k,
		"resetEndpointFilter",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		k,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsEkmConnectionServiceResolversOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
