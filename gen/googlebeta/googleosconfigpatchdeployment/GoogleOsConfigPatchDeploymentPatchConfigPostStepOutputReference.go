package googleosconfigpatchdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleosconfigpatchdeployment/internal"
)

type GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference interface {
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
	InternalValue() *GoogleOsConfigPatchDeploymentPatchConfigPostStep
	SetInternalValue(val *GoogleOsConfigPatchDeploymentPatchConfigPostStep)
	LinuxExecStepConfig() GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference
	LinuxExecStepConfigInput() *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsExecStepConfig() GoogleOsConfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigOutputReference
	WindowsExecStepConfigInput() *GoogleOsConfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfig
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
	PutLinuxExecStepConfig(value *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig)
	PutWindowsExecStepConfig(value *GoogleOsConfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfig)
	ResetLinuxExecStepConfig()
	ResetWindowsExecStepConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference
type jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) InternalValue() *GoogleOsConfigPatchDeploymentPatchConfigPostStep {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigPostStep
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) LinuxExecStepConfig() GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference {
	var returns GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference
	_jsii_.Get(
		j,
		"linuxExecStepConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) LinuxExecStepConfigInput() *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig
	_jsii_.Get(
		j,
		"linuxExecStepConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) WindowsExecStepConfig() GoogleOsConfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigOutputReference {
	var returns GoogleOsConfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfigOutputReference
	_jsii_.Get(
		j,
		"windowsExecStepConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) WindowsExecStepConfigInput() *GoogleOsConfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfig {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfig
	_jsii_.Get(
		j,
		"windowsExecStepConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference_Override(g GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference)SetInternalValue(val *GoogleOsConfigPatchDeploymentPatchConfigPostStep) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) PutLinuxExecStepConfig(value *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig) {
	if err := g.validatePutLinuxExecStepConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLinuxExecStepConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) PutWindowsExecStepConfig(value *GoogleOsConfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfig) {
	if err := g.validatePutWindowsExecStepConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWindowsExecStepConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) ResetLinuxExecStepConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLinuxExecStepConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) ResetWindowsExecStepConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowsExecStepConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

