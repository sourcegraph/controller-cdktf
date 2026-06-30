package colabruntimetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/colabruntimetemplate/internal"
)

type ColabRuntimeTemplateSoftwareConfigOutputReference interface {
	cdktf.ComplexObject
	ColabImage() ColabRuntimeTemplateSoftwareConfigColabImageOutputReference
	ColabImageInput() *ColabRuntimeTemplateSoftwareConfigColabImage
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
	Env() ColabRuntimeTemplateSoftwareConfigEnvList
	EnvInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *ColabRuntimeTemplateSoftwareConfig
	SetInternalValue(val *ColabRuntimeTemplateSoftwareConfig)
	PostStartupScriptConfig() ColabRuntimeTemplateSoftwareConfigPostStartupScriptConfigOutputReference
	PostStartupScriptConfigInput() *ColabRuntimeTemplateSoftwareConfigPostStartupScriptConfig
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
	PutColabImage(value *ColabRuntimeTemplateSoftwareConfigColabImage)
	PutEnv(value any)
	PutPostStartupScriptConfig(value *ColabRuntimeTemplateSoftwareConfigPostStartupScriptConfig)
	ResetColabImage()
	ResetEnv()
	ResetPostStartupScriptConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ColabRuntimeTemplateSoftwareConfigOutputReference
type jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) ColabImage() ColabRuntimeTemplateSoftwareConfigColabImageOutputReference {
	var returns ColabRuntimeTemplateSoftwareConfigColabImageOutputReference
	_jsii_.Get(
		j,
		"colabImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) ColabImageInput() *ColabRuntimeTemplateSoftwareConfigColabImage {
	var returns *ColabRuntimeTemplateSoftwareConfigColabImage
	_jsii_.Get(
		j,
		"colabImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) Env() ColabRuntimeTemplateSoftwareConfigEnvList {
	var returns ColabRuntimeTemplateSoftwareConfigEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) EnvInput() any {
	var returns any
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) InternalValue() *ColabRuntimeTemplateSoftwareConfig {
	var returns *ColabRuntimeTemplateSoftwareConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) PostStartupScriptConfig() ColabRuntimeTemplateSoftwareConfigPostStartupScriptConfigOutputReference {
	var returns ColabRuntimeTemplateSoftwareConfigPostStartupScriptConfigOutputReference
	_jsii_.Get(
		j,
		"postStartupScriptConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) PostStartupScriptConfigInput() *ColabRuntimeTemplateSoftwareConfigPostStartupScriptConfig {
	var returns *ColabRuntimeTemplateSoftwareConfigPostStartupScriptConfig
	_jsii_.Get(
		j,
		"postStartupScriptConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewColabRuntimeTemplateSoftwareConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ColabRuntimeTemplateSoftwareConfigOutputReference {
	_init_.Initialize()

	if err := validateNewColabRuntimeTemplateSoftwareConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.colabRuntimeTemplate.ColabRuntimeTemplateSoftwareConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewColabRuntimeTemplateSoftwareConfigOutputReference_Override(c ColabRuntimeTemplateSoftwareConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.colabRuntimeTemplate.ColabRuntimeTemplateSoftwareConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) SetInternalValue(val *ColabRuntimeTemplateSoftwareConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) PutColabImage(value *ColabRuntimeTemplateSoftwareConfigColabImage) {
	if err := c.validatePutColabImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putColabImage",
		[]any{value},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) PutEnv(value any) {
	if err := c.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnv",
		[]any{value},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) PutPostStartupScriptConfig(value *ColabRuntimeTemplateSoftwareConfigPostStartupScriptConfig) {
	if err := c.validatePutPostStartupScriptConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPostStartupScriptConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) ResetColabImage() {
	_jsii_.InvokeVoid(
		c,
		"resetColabImage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		c,
		"resetEnv",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) ResetPostStartupScriptConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPostStartupScriptConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplateSoftwareConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
