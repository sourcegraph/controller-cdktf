package googlestoragetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlestoragetransferjob/internal"
)

type GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference interface {
	cdktf.ComplexObject
	Acl() *string
	SetAcl(val *string)
	AclInput() *string
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
	Gid() *string
	SetGid(val *string)
	GidInput() *string
	InternalValue() *GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptions
	SetInternalValue(val *GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptions)
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	StorageClass() *string
	SetStorageClass(val *string)
	StorageClassInput() *string
	Symlink() *string
	SetSymlink(val *string)
	SymlinkInput() *string
	TemporaryHold() *string
	SetTemporaryHold(val *string)
	TemporaryHoldInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeCreated() *string
	SetTimeCreated(val *string)
	TimeCreatedInput() *string
	Uid() *string
	SetUid(val *string)
	UidInput() *string
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
	ResetAcl()
	ResetGid()
	ResetKmsKey()
	ResetMode()
	ResetStorageClass()
	ResetSymlink()
	ResetTemporaryHold()
	ResetTimeCreated()
	ResetUid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference
type jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) Acl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) AclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) Gid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) InternalValue() *GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptions {
	var returns *GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) Symlink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"symlink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) SymlinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"symlinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) TemporaryHold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"temporaryHold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) TemporaryHoldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"temporaryHoldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) TimeCreated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeCreated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) TimeCreatedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeCreatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) UidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uidInput",
		&returns,
	)
	return returns
}


func NewGoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageTransferJob.GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference_Override(g GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageTransferJob.GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetAcl(val *string) {
	if err := j.validateSetAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acl",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetGid(val *string) {
	if err := j.validateSetGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gid",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetInternalValue(val *GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetSymlink(val *string) {
	if err := j.validateSetSymlinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"symlink",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetTemporaryHold(val *string) {
	if err := j.validateSetTemporaryHoldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temporaryHold",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetTimeCreated(val *string) {
	if err := j.validateSetTimeCreatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeCreated",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetUid(val *string) {
	if err := j.validateSetUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uid",
		val,
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetAcl() {
	_jsii_.InvokeVoid(
		g,
		"resetAcl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetGid() {
	_jsii_.InvokeVoid(
		g,
		"resetGid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetKmsKey() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		g,
		"resetMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetStorageClass() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetSymlink() {
	_jsii_.InvokeVoid(
		g,
		"resetSymlink",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetTemporaryHold() {
	_jsii_.InvokeVoid(
		g,
		"resetTemporaryHold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetTimeCreated() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeCreated",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetUid() {
	_jsii_.InvokeVoid(
		g,
		"resetUid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

