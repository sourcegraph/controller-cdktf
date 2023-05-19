package googleworkstationsworkstationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleworkstationsworkstationconfig/internal"
)

type GoogleWorkstationsWorkstationConfigContainerOutputReference interface {
	cdktf.ComplexObject
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	Env() *map[string]*string
	SetEnv(val *map[string]*string)
	EnvInput() *map[string]*string
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InternalValue() *GoogleWorkstationsWorkstationConfigContainer
	SetInternalValue(val *GoogleWorkstationsWorkstationConfigContainer)
	RunAsUser() *float64
	SetRunAsUser(val *float64)
	RunAsUserInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkingDir() *string
	SetWorkingDir(val *string)
	WorkingDirInput() *string
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
	ResetArgs()
	ResetCommand()
	ResetEnv()
	ResetImage()
	ResetRunAsUser()
	ResetWorkingDir()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleWorkstationsWorkstationConfigContainerOutputReference
type jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) Env() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) EnvInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) InternalValue() *GoogleWorkstationsWorkstationConfigContainer {
	var returns *GoogleWorkstationsWorkstationConfigContainer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) RunAsUser() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) RunAsUserInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runAsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) WorkingDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirInput",
		&returns,
	)
	return returns
}


func NewGoogleWorkstationsWorkstationConfigContainerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleWorkstationsWorkstationConfigContainerOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleWorkstationsWorkstationConfigContainerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference{}

	_jsii_.Create(
		"google-beta.googleWorkstationsWorkstationConfig.GoogleWorkstationsWorkstationConfigContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleWorkstationsWorkstationConfigContainerOutputReference_Override(g GoogleWorkstationsWorkstationConfigContainerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleWorkstationsWorkstationConfig.GoogleWorkstationsWorkstationConfigContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference)SetEnv(val *map[string]*string) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference)SetInternalValue(val *GoogleWorkstationsWorkstationConfigContainer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference)SetRunAsUser(val *float64) {
	if err := j.validateSetRunAsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUser",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference)SetWorkingDir(val *string) {
	if err := j.validateSetWorkingDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDir",
		val,
	)
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		g,
		"resetArgs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		g,
		"resetCommand",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		g,
		"resetEnv",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		g,
		"resetImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) ResetRunAsUser() {
	_jsii_.InvokeVoid(
		g,
		"resetRunAsUser",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) ResetWorkingDir() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkingDir",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleWorkstationsWorkstationConfigContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

