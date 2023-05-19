package googleosconfigpatchdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleosconfigpatchdeployment/internal"
)

type GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference interface {
	cdktf.ComplexObject
	AllowedSuccessCodes() *[]*float64
	SetAllowedSuccessCodes(val *[]*float64)
	AllowedSuccessCodesInput() *[]*float64
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
	GcsObject() GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObjectOutputReference
	GcsObjectInput() *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject
	InternalValue() *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig
	SetInternalValue(val *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig)
	Interpreter() *string
	SetInterpreter(val *string)
	InterpreterInput() *string
	LocalPath() *string
	SetLocalPath(val *string)
	LocalPathInput() *string
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
	PutGcsObject(value *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject)
	ResetAllowedSuccessCodes()
	ResetGcsObject()
	ResetInterpreter()
	ResetLocalPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference
type jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) AllowedSuccessCodes() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedSuccessCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) AllowedSuccessCodesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedSuccessCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) GcsObject() GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObjectOutputReference {
	var returns GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObjectOutputReference
	_jsii_.Get(
		j,
		"gcsObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) GcsObjectInput() *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject
	_jsii_.Get(
		j,
		"gcsObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) InternalValue() *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) Interpreter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interpreter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) InterpreterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interpreterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) LocalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) LocalPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference_Override(g GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference)SetAllowedSuccessCodes(val *[]*float64) {
	if err := j.validateSetAllowedSuccessCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedSuccessCodes",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference)SetInternalValue(val *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference)SetInterpreter(val *string) {
	if err := j.validateSetInterpreterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interpreter",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference)SetLocalPath(val *string) {
	if err := j.validateSetLocalPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localPath",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) PutGcsObject(value *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject) {
	if err := g.validatePutGcsObjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsObject",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) ResetAllowedSuccessCodes() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedSuccessCodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) ResetGcsObject() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsObject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) ResetInterpreter() {
	_jsii_.InvokeVoid(
		g,
		"resetInterpreter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) ResetLocalPath() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

