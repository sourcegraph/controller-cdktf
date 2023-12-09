package accessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/accessapplication/internal"
)

type AccessApplicationSaasAppOutputReference interface {
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
	ConsumerServiceUrl() *string
	SetConsumerServiceUrl(val *string)
	ConsumerServiceUrlInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomAttribute() AccessApplicationSaasAppCustomAttributeList
	CustomAttributeInput() interface{}
	// Experimental.
	Fqn() *string
	IdpEntityId() *string
	InternalValue() *AccessApplicationSaasApp
	SetInternalValue(val *AccessApplicationSaasApp)
	NameIdFormat() *string
	SetNameIdFormat(val *string)
	NameIdFormatInput() *string
	PublicKey() *string
	SpEntityId() *string
	SetSpEntityId(val *string)
	SpEntityIdInput() *string
	SsoEndpoint() *string
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
	PutCustomAttribute(value interface{})
	ResetCustomAttribute()
	ResetNameIdFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessApplicationSaasAppOutputReference
type jsiiProxy_AccessApplicationSaasAppOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) ConsumerServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerServiceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) ConsumerServiceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerServiceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) CustomAttribute() AccessApplicationSaasAppCustomAttributeList {
	var returns AccessApplicationSaasAppCustomAttributeList
	_jsii_.Get(
		j,
		"customAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) CustomAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) IdpEntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpEntityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) InternalValue() *AccessApplicationSaasApp {
	var returns *AccessApplicationSaasApp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) NameIdFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameIdFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) NameIdFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameIdFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) SpEntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spEntityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) SpEntityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spEntityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) SsoEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessApplicationSaasAppOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccessApplicationSaasAppOutputReference {
	_init_.Initialize()

	if err := validateNewAccessApplicationSaasAppOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessApplicationSaasAppOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessApplication.AccessApplicationSaasAppOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccessApplicationSaasAppOutputReference_Override(a AccessApplicationSaasAppOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessApplication.AccessApplicationSaasAppOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetConsumerServiceUrl(val *string) {
	if err := j.validateSetConsumerServiceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerServiceUrl",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetInternalValue(val *AccessApplicationSaasApp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetNameIdFormat(val *string) {
	if err := j.validateSetNameIdFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameIdFormat",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetSpEntityId(val *string) {
	if err := j.validateSetSpEntityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spEntityId",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationSaasAppOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) PutCustomAttribute(value interface{}) {
	if err := a.validatePutCustomAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomAttribute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetCustomAttribute() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomAttribute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ResetNameIdFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetNameIdFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationSaasAppOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

