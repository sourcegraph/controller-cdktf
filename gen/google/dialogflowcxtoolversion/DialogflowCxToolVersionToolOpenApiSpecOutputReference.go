package dialogflowcxtoolversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowcxtoolversion/internal"
)

type DialogflowCxToolVersionToolOpenApiSpecOutputReference interface {
	cdktf.ComplexObject
	Authentication() DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference
	AuthenticationInput() *DialogflowCxToolVersionToolOpenApiSpecAuthentication
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
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowCxToolVersionToolOpenApiSpec
	SetInternalValue(val *DialogflowCxToolVersionToolOpenApiSpec)
	ServiceDirectoryConfig() DialogflowCxToolVersionToolOpenApiSpecServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *DialogflowCxToolVersionToolOpenApiSpecServiceDirectoryConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextSchema() *string
	SetTextSchema(val *string)
	TextSchemaInput() *string
	TlsConfig() DialogflowCxToolVersionToolOpenApiSpecTlsConfigOutputReference
	TlsConfigInput() *DialogflowCxToolVersionToolOpenApiSpecTlsConfig
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
	PutAuthentication(value *DialogflowCxToolVersionToolOpenApiSpecAuthentication)
	PutServiceDirectoryConfig(value *DialogflowCxToolVersionToolOpenApiSpecServiceDirectoryConfig)
	PutTlsConfig(value *DialogflowCxToolVersionToolOpenApiSpecTlsConfig)
	ResetAuthentication()
	ResetServiceDirectoryConfig()
	ResetTlsConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxToolVersionToolOpenApiSpecOutputReference
type jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) Authentication() DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference {
	var returns DialogflowCxToolVersionToolOpenApiSpecAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) AuthenticationInput() *DialogflowCxToolVersionToolOpenApiSpecAuthentication {
	var returns *DialogflowCxToolVersionToolOpenApiSpecAuthentication
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) InternalValue() *DialogflowCxToolVersionToolOpenApiSpec {
	var returns *DialogflowCxToolVersionToolOpenApiSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) ServiceDirectoryConfig() DialogflowCxToolVersionToolOpenApiSpecServiceDirectoryConfigOutputReference {
	var returns DialogflowCxToolVersionToolOpenApiSpecServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) ServiceDirectoryConfigInput() *DialogflowCxToolVersionToolOpenApiSpecServiceDirectoryConfig {
	var returns *DialogflowCxToolVersionToolOpenApiSpecServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) TextSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) TextSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) TlsConfig() DialogflowCxToolVersionToolOpenApiSpecTlsConfigOutputReference {
	var returns DialogflowCxToolVersionToolOpenApiSpecTlsConfigOutputReference
	_jsii_.Get(
		j,
		"tlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) TlsConfigInput() *DialogflowCxToolVersionToolOpenApiSpecTlsConfig {
	var returns *DialogflowCxToolVersionToolOpenApiSpecTlsConfig
	_jsii_.Get(
		j,
		"tlsConfigInput",
		&returns,
	)
	return returns
}

func NewDialogflowCxToolVersionToolOpenApiSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxToolVersionToolOpenApiSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxToolVersionToolOpenApiSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxToolVersion.DialogflowCxToolVersionToolOpenApiSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxToolVersionToolOpenApiSpecOutputReference_Override(d DialogflowCxToolVersionToolOpenApiSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxToolVersion.DialogflowCxToolVersionToolOpenApiSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) SetInternalValue(val *DialogflowCxToolVersionToolOpenApiSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) SetTextSchema(val *string) {
	if err := j.validateSetTextSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textSchema",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) PutAuthentication(value *DialogflowCxToolVersionToolOpenApiSpecAuthentication) {
	if err := d.validatePutAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAuthentication",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) PutServiceDirectoryConfig(value *DialogflowCxToolVersionToolOpenApiSpecServiceDirectoryConfig) {
	if err := d.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceDirectoryConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) PutTlsConfig(value *DialogflowCxToolVersionToolOpenApiSpecTlsConfig) {
	if err := d.validatePutTlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTlsConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) ResetAuthentication() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) ResetTlsConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetTlsConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOpenApiSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
