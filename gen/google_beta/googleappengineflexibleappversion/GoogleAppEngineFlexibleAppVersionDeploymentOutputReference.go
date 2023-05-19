package googleappengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleappengineflexibleappversion/internal"
)

type GoogleAppEngineFlexibleAppVersionDeploymentOutputReference interface {
	cdktf.ComplexObject
	CloudBuildOptions() GoogleAppEngineFlexibleAppVersionDeploymentCloudBuildOptionsOutputReference
	CloudBuildOptionsInput() *GoogleAppEngineFlexibleAppVersionDeploymentCloudBuildOptions
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
	Container() GoogleAppEngineFlexibleAppVersionDeploymentContainerOutputReference
	ContainerInput() *GoogleAppEngineFlexibleAppVersionDeploymentContainer
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Files() GoogleAppEngineFlexibleAppVersionDeploymentFilesList
	FilesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleAppEngineFlexibleAppVersionDeployment
	SetInternalValue(val *GoogleAppEngineFlexibleAppVersionDeployment)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Zip() GoogleAppEngineFlexibleAppVersionDeploymentZipOutputReference
	ZipInput() *GoogleAppEngineFlexibleAppVersionDeploymentZip
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
	PutCloudBuildOptions(value *GoogleAppEngineFlexibleAppVersionDeploymentCloudBuildOptions)
	PutContainer(value *GoogleAppEngineFlexibleAppVersionDeploymentContainer)
	PutFiles(value interface{})
	PutZip(value *GoogleAppEngineFlexibleAppVersionDeploymentZip)
	ResetCloudBuildOptions()
	ResetContainer()
	ResetFiles()
	ResetZip()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAppEngineFlexibleAppVersionDeploymentOutputReference
type jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) CloudBuildOptions() GoogleAppEngineFlexibleAppVersionDeploymentCloudBuildOptionsOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionDeploymentCloudBuildOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudBuildOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) CloudBuildOptionsInput() *GoogleAppEngineFlexibleAppVersionDeploymentCloudBuildOptions {
	var returns *GoogleAppEngineFlexibleAppVersionDeploymentCloudBuildOptions
	_jsii_.Get(
		j,
		"cloudBuildOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) Container() GoogleAppEngineFlexibleAppVersionDeploymentContainerOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionDeploymentContainerOutputReference
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) ContainerInput() *GoogleAppEngineFlexibleAppVersionDeploymentContainer {
	var returns *GoogleAppEngineFlexibleAppVersionDeploymentContainer
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) Files() GoogleAppEngineFlexibleAppVersionDeploymentFilesList {
	var returns GoogleAppEngineFlexibleAppVersionDeploymentFilesList
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) FilesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) InternalValue() *GoogleAppEngineFlexibleAppVersionDeployment {
	var returns *GoogleAppEngineFlexibleAppVersionDeployment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) Zip() GoogleAppEngineFlexibleAppVersionDeploymentZipOutputReference {
	var returns GoogleAppEngineFlexibleAppVersionDeploymentZipOutputReference
	_jsii_.Get(
		j,
		"zip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) ZipInput() *GoogleAppEngineFlexibleAppVersionDeploymentZip {
	var returns *GoogleAppEngineFlexibleAppVersionDeploymentZip
	_jsii_.Get(
		j,
		"zipInput",
		&returns,
	)
	return returns
}


func NewGoogleAppEngineFlexibleAppVersionDeploymentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAppEngineFlexibleAppVersionDeploymentOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAppEngineFlexibleAppVersionDeploymentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference{}

	_jsii_.Create(
		"google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersionDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAppEngineFlexibleAppVersionDeploymentOutputReference_Override(g GoogleAppEngineFlexibleAppVersionDeploymentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersionDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference)SetInternalValue(val *GoogleAppEngineFlexibleAppVersionDeployment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) PutCloudBuildOptions(value *GoogleAppEngineFlexibleAppVersionDeploymentCloudBuildOptions) {
	if err := g.validatePutCloudBuildOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudBuildOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) PutContainer(value *GoogleAppEngineFlexibleAppVersionDeploymentContainer) {
	if err := g.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContainer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) PutFiles(value interface{}) {
	if err := g.validatePutFilesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFiles",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) PutZip(value *GoogleAppEngineFlexibleAppVersionDeploymentZip) {
	if err := g.validatePutZipParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putZip",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) ResetCloudBuildOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudBuildOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) ResetContainer() {
	_jsii_.InvokeVoid(
		g,
		"resetContainer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) ResetFiles() {
	_jsii_.InvokeVoid(
		g,
		"resetFiles",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) ResetZip() {
	_jsii_.InvokeVoid(
		g,
		"resetZip",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionDeploymentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

