package accessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/accessapplication/internal"
)

type AccessApplicationCorsHeadersOutputReference interface {
	cdktf.ComplexObject
	AllowAllHeaders() interface{}
	SetAllowAllHeaders(val interface{})
	AllowAllHeadersInput() interface{}
	AllowAllMethods() interface{}
	SetAllowAllMethods(val interface{})
	AllowAllMethodsInput() interface{}
	AllowAllOrigins() interface{}
	SetAllowAllOrigins(val interface{})
	AllowAllOriginsInput() interface{}
	AllowCredentials() interface{}
	SetAllowCredentials(val interface{})
	AllowCredentialsInput() interface{}
	AllowedHeaders() *[]*string
	SetAllowedHeaders(val *[]*string)
	AllowedHeadersInput() *[]*string
	AllowedMethods() *[]*string
	SetAllowedMethods(val *[]*string)
	AllowedMethodsInput() *[]*string
	AllowedOrigins() *[]*string
	SetAllowedOrigins(val *[]*string)
	AllowedOriginsInput() *[]*string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxAge() *float64
	SetMaxAge(val *float64)
	MaxAgeInput() *float64
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
	ResetAllowAllHeaders()
	ResetAllowAllMethods()
	ResetAllowAllOrigins()
	ResetAllowCredentials()
	ResetAllowedHeaders()
	ResetAllowedMethods()
	ResetAllowedOrigins()
	ResetMaxAge()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessApplicationCorsHeadersOutputReference
type jsiiProxy_AccessApplicationCorsHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) AllowAllHeaders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) AllowAllHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) AllowAllMethods() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) AllowAllMethodsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) AllowAllOrigins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) AllowAllOriginsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) AllowCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) AllowCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) AllowedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) AllowedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) AllowedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) AllowedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) AllowedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) AllowedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) MaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) MaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessApplicationCorsHeadersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AccessApplicationCorsHeadersOutputReference {
	_init_.Initialize()

	if err := validateNewAccessApplicationCorsHeadersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessApplicationCorsHeadersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessApplication.AccessApplicationCorsHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAccessApplicationCorsHeadersOutputReference_Override(a AccessApplicationCorsHeadersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accessApplication.AccessApplicationCorsHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference)SetAllowAllHeaders(val interface{}) {
	if err := j.validateSetAllowAllHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAllHeaders",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference)SetAllowAllMethods(val interface{}) {
	if err := j.validateSetAllowAllMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAllMethods",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference)SetAllowAllOrigins(val interface{}) {
	if err := j.validateSetAllowAllOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAllOrigins",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference)SetAllowCredentials(val interface{}) {
	if err := j.validateSetAllowCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowCredentials",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference)SetAllowedHeaders(val *[]*string) {
	if err := j.validateSetAllowedHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedHeaders",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference)SetAllowedMethods(val *[]*string) {
	if err := j.validateSetAllowedMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedMethods",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference)SetAllowedOrigins(val *[]*string) {
	if err := j.validateSetAllowedOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOrigins",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference)SetMaxAge(val *float64) {
	if err := j.validateSetMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAge",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessApplicationCorsHeadersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) ResetAllowAllHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowAllHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) ResetAllowAllMethods() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowAllMethods",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) ResetAllowAllOrigins() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowAllOrigins",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) ResetAllowCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) ResetAllowedHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) ResetAllowedMethods() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedMethods",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) ResetAllowedOrigins() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedOrigins",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) ResetMaxAge() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxAge",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccessApplicationCorsHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

