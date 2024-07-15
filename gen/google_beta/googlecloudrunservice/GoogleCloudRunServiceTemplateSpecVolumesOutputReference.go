package googlecloudrunservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudrunservice/internal"
)

type GoogleCloudRunServiceTemplateSpecVolumesOutputReference interface {
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
	Csi() GoogleCloudRunServiceTemplateSpecVolumesCsiOutputReference
	CsiInput() *GoogleCloudRunServiceTemplateSpecVolumesCsi
	EmptyDir() GoogleCloudRunServiceTemplateSpecVolumesEmptyDirOutputReference
	EmptyDirInput() *GoogleCloudRunServiceTemplateSpecVolumesEmptyDir
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() GoogleCloudRunServiceTemplateSpecVolumesNfsOutputReference
	NfsInput() *GoogleCloudRunServiceTemplateSpecVolumesNfs
	Secret() GoogleCloudRunServiceTemplateSpecVolumesSecretOutputReference
	SecretInput() *GoogleCloudRunServiceTemplateSpecVolumesSecret
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
	PutCsi(value *GoogleCloudRunServiceTemplateSpecVolumesCsi)
	PutEmptyDir(value *GoogleCloudRunServiceTemplateSpecVolumesEmptyDir)
	PutNfs(value *GoogleCloudRunServiceTemplateSpecVolumesNfs)
	PutSecret(value *GoogleCloudRunServiceTemplateSpecVolumesSecret)
	ResetCsi()
	ResetEmptyDir()
	ResetNfs()
	ResetSecret()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudRunServiceTemplateSpecVolumesOutputReference
type jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) Csi() GoogleCloudRunServiceTemplateSpecVolumesCsiOutputReference {
	var returns GoogleCloudRunServiceTemplateSpecVolumesCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) CsiInput() *GoogleCloudRunServiceTemplateSpecVolumesCsi {
	var returns *GoogleCloudRunServiceTemplateSpecVolumesCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) EmptyDir() GoogleCloudRunServiceTemplateSpecVolumesEmptyDirOutputReference {
	var returns GoogleCloudRunServiceTemplateSpecVolumesEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) EmptyDirInput() *GoogleCloudRunServiceTemplateSpecVolumesEmptyDir {
	var returns *GoogleCloudRunServiceTemplateSpecVolumesEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) Nfs() GoogleCloudRunServiceTemplateSpecVolumesNfsOutputReference {
	var returns GoogleCloudRunServiceTemplateSpecVolumesNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) NfsInput() *GoogleCloudRunServiceTemplateSpecVolumesNfs {
	var returns *GoogleCloudRunServiceTemplateSpecVolumesNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) Secret() GoogleCloudRunServiceTemplateSpecVolumesSecretOutputReference {
	var returns GoogleCloudRunServiceTemplateSpecVolumesSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) SecretInput() *GoogleCloudRunServiceTemplateSpecVolumesSecret {
	var returns *GoogleCloudRunServiceTemplateSpecVolumesSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudRunServiceTemplateSpecVolumesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleCloudRunServiceTemplateSpecVolumesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudRunServiceTemplateSpecVolumesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudRunService.GoogleCloudRunServiceTemplateSpecVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleCloudRunServiceTemplateSpecVolumesOutputReference_Override(g GoogleCloudRunServiceTemplateSpecVolumesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudRunService.GoogleCloudRunServiceTemplateSpecVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) PutCsi(value *GoogleCloudRunServiceTemplateSpecVolumesCsi) {
	if err := g.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCsi",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) PutEmptyDir(value *GoogleCloudRunServiceTemplateSpecVolumesEmptyDir) {
	if err := g.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) PutNfs(value *GoogleCloudRunServiceTemplateSpecVolumesNfs) {
	if err := g.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNfs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) PutSecret(value *GoogleCloudRunServiceTemplateSpecVolumesSecret) {
	if err := g.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecret",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		g,
		"resetCsi",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		g,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		g,
		"resetNfs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		g,
		"resetSecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudRunServiceTemplateSpecVolumesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

