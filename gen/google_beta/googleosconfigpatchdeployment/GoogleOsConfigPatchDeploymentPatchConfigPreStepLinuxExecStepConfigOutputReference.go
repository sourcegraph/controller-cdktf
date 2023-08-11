package googleosconfigpatchdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleosconfigpatchdeployment/internal"
)

type GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference interface {
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
	GcsObject() GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObjectOutputReference
	GcsObjectInput() *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject
	InternalValue() *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig
	SetInternalValue(val *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig)
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
	PutGcsObject(value *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject)
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

// The jsii proxy struct for GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference
type jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) AllowedSuccessCodes() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedSuccessCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) AllowedSuccessCodesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedSuccessCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) GcsObject() GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObjectOutputReference {
	var returns GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObjectOutputReference
	_jsii_.Get(
		j,
		"gcsObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) GcsObjectInput() *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject
	_jsii_.Get(
		j,
		"gcsObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) InternalValue() *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) Interpreter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interpreter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) InterpreterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interpreterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) LocalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) LocalPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference_Override(g GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference)SetAllowedSuccessCodes(val *[]*float64) {
	if err := j.validateSetAllowedSuccessCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedSuccessCodes",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference)SetInternalValue(val *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference)SetInterpreter(val *string) {
	if err := j.validateSetInterpreterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interpreter",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference)SetLocalPath(val *string) {
	if err := j.validateSetLocalPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localPath",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) PutGcsObject(value *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject) {
	if err := g.validatePutGcsObjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsObject",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) ResetAllowedSuccessCodes() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedSuccessCodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) ResetGcsObject() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsObject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) ResetInterpreter() {
	_jsii_.InvokeVoid(
		g,
		"resetInterpreter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) ResetLocalPath() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

