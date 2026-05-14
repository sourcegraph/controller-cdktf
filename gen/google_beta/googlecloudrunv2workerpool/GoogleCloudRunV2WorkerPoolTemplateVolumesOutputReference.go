package googlecloudrunv2workerpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudrunv2workerpool/internal"
)

type GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference interface {
	cdktf.ComplexObject
	CloudSqlInstance() GoogleCloudRunV2WorkerPoolTemplateVolumesCloudSqlInstanceOutputReference
	CloudSqlInstanceInput() *GoogleCloudRunV2WorkerPoolTemplateVolumesCloudSqlInstance
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
	EmptyDir() GoogleCloudRunV2WorkerPoolTemplateVolumesEmptyDirOutputReference
	EmptyDirInput() *GoogleCloudRunV2WorkerPoolTemplateVolumesEmptyDir
	// Experimental.
	Fqn() *string
	Gcs() GoogleCloudRunV2WorkerPoolTemplateVolumesGcsOutputReference
	GcsInput() *GoogleCloudRunV2WorkerPoolTemplateVolumesGcs
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() GoogleCloudRunV2WorkerPoolTemplateVolumesNfsOutputReference
	NfsInput() *GoogleCloudRunV2WorkerPoolTemplateVolumesNfs
	Secret() GoogleCloudRunV2WorkerPoolTemplateVolumesSecretOutputReference
	SecretInput() *GoogleCloudRunV2WorkerPoolTemplateVolumesSecret
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
	PutCloudSqlInstance(value *GoogleCloudRunV2WorkerPoolTemplateVolumesCloudSqlInstance)
	PutEmptyDir(value *GoogleCloudRunV2WorkerPoolTemplateVolumesEmptyDir)
	PutGcs(value *GoogleCloudRunV2WorkerPoolTemplateVolumesGcs)
	PutNfs(value *GoogleCloudRunV2WorkerPoolTemplateVolumesNfs)
	PutSecret(value *GoogleCloudRunV2WorkerPoolTemplateVolumesSecret)
	ResetCloudSqlInstance()
	ResetEmptyDir()
	ResetGcs()
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

// The jsii proxy struct for GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference
type jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) CloudSqlInstance() GoogleCloudRunV2WorkerPoolTemplateVolumesCloudSqlInstanceOutputReference {
	var returns GoogleCloudRunV2WorkerPoolTemplateVolumesCloudSqlInstanceOutputReference
	_jsii_.Get(
		j,
		"cloudSqlInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) CloudSqlInstanceInput() *GoogleCloudRunV2WorkerPoolTemplateVolumesCloudSqlInstance {
	var returns *GoogleCloudRunV2WorkerPoolTemplateVolumesCloudSqlInstance
	_jsii_.Get(
		j,
		"cloudSqlInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) EmptyDir() GoogleCloudRunV2WorkerPoolTemplateVolumesEmptyDirOutputReference {
	var returns GoogleCloudRunV2WorkerPoolTemplateVolumesEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) EmptyDirInput() *GoogleCloudRunV2WorkerPoolTemplateVolumesEmptyDir {
	var returns *GoogleCloudRunV2WorkerPoolTemplateVolumesEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) Gcs() GoogleCloudRunV2WorkerPoolTemplateVolumesGcsOutputReference {
	var returns GoogleCloudRunV2WorkerPoolTemplateVolumesGcsOutputReference
	_jsii_.Get(
		j,
		"gcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) GcsInput() *GoogleCloudRunV2WorkerPoolTemplateVolumesGcs {
	var returns *GoogleCloudRunV2WorkerPoolTemplateVolumesGcs
	_jsii_.Get(
		j,
		"gcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) Nfs() GoogleCloudRunV2WorkerPoolTemplateVolumesNfsOutputReference {
	var returns GoogleCloudRunV2WorkerPoolTemplateVolumesNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) NfsInput() *GoogleCloudRunV2WorkerPoolTemplateVolumesNfs {
	var returns *GoogleCloudRunV2WorkerPoolTemplateVolumesNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) Secret() GoogleCloudRunV2WorkerPoolTemplateVolumesSecretOutputReference {
	var returns GoogleCloudRunV2WorkerPoolTemplateVolumesSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) SecretInput() *GoogleCloudRunV2WorkerPoolTemplateVolumesSecret {
	var returns *GoogleCloudRunV2WorkerPoolTemplateVolumesSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudRunV2WorkerPoolTemplateVolumesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudRunV2WorkerPool.GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference_Override(g GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudRunV2WorkerPool.GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) PutCloudSqlInstance(value *GoogleCloudRunV2WorkerPoolTemplateVolumesCloudSqlInstance) {
	if err := g.validatePutCloudSqlInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudSqlInstance",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) PutEmptyDir(value *GoogleCloudRunV2WorkerPoolTemplateVolumesEmptyDir) {
	if err := g.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) PutGcs(value *GoogleCloudRunV2WorkerPoolTemplateVolumesGcs) {
	if err := g.validatePutGcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) PutNfs(value *GoogleCloudRunV2WorkerPoolTemplateVolumesNfs) {
	if err := g.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNfs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) PutSecret(value *GoogleCloudRunV2WorkerPoolTemplateVolumesSecret) {
	if err := g.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecret",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) ResetCloudSqlInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudSqlInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		g,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) ResetGcs() {
	_jsii_.InvokeVoid(
		g,
		"resetGcs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		g,
		"resetNfs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		g,
		"resetSecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateVolumesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

