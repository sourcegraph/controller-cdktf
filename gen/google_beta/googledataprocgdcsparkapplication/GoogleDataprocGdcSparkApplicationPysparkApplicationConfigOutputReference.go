package googledataprocgdcsparkapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocgdcsparkapplication/internal"
)

type GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference interface {
	cdktf.ComplexObject
	ArchiveUris() *[]*string
	SetArchiveUris(val *[]*string)
	ArchiveUrisInput() *[]*string
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
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
	FileUris() *[]*string
	SetFileUris(val *[]*string)
	FileUrisInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataprocGdcSparkApplicationPysparkApplicationConfig
	SetInternalValue(val *GoogleDataprocGdcSparkApplicationPysparkApplicationConfig)
	JarFileUris() *[]*string
	SetJarFileUris(val *[]*string)
	JarFileUrisInput() *[]*string
	MainPythonFileUri() *string
	SetMainPythonFileUri(val *string)
	MainPythonFileUriInput() *string
	PythonFileUris() *[]*string
	SetPythonFileUris(val *[]*string)
	PythonFileUrisInput() *[]*string
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
	ResetArchiveUris()
	ResetArgs()
	ResetFileUris()
	ResetJarFileUris()
	ResetPythonFileUris()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference
type jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) ArchiveUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"archiveUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) ArchiveUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"archiveUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) FileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) FileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) InternalValue() *GoogleDataprocGdcSparkApplicationPysparkApplicationConfig {
	var returns *GoogleDataprocGdcSparkApplicationPysparkApplicationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) JarFileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarFileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) JarFileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarFileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) MainPythonFileUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainPythonFileUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) MainPythonFileUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainPythonFileUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) PythonFileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pythonFileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) PythonFileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pythonFileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocGdcSparkApplication.GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference_Override(g GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocGdcSparkApplication.GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference)SetArchiveUris(val *[]*string) {
	if err := j.validateSetArchiveUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archiveUris",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference)SetFileUris(val *[]*string) {
	if err := j.validateSetFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileUris",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference)SetInternalValue(val *GoogleDataprocGdcSparkApplicationPysparkApplicationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference)SetJarFileUris(val *[]*string) {
	if err := j.validateSetJarFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jarFileUris",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference)SetMainPythonFileUri(val *string) {
	if err := j.validateSetMainPythonFileUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mainPythonFileUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference)SetPythonFileUris(val *[]*string) {
	if err := j.validateSetPythonFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonFileUris",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) ResetArchiveUris() {
	_jsii_.InvokeVoid(
		g,
		"resetArchiveUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		g,
		"resetArgs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) ResetFileUris() {
	_jsii_.InvokeVoid(
		g,
		"resetFileUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) ResetJarFileUris() {
	_jsii_.InvokeVoid(
		g,
		"resetJarFileUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) ResetPythonFileUris() {
	_jsii_.InvokeVoid(
		g,
		"resetPythonFileUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocGdcSparkApplicationPysparkApplicationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

