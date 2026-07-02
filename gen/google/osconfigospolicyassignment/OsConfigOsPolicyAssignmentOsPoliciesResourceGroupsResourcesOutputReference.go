package osconfigospolicyassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/osconfigospolicyassignment/internal"
)

type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference interface {
	cdktf.ComplexObject
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
	Exec() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference
	ExecInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec
	File() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileOutputReference
	FileInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() any
	SetInternalValue(val any)
	Pkg() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference
	PkgInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg
	Repository() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference
	RepositoryInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository
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
	PutExec(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec)
	PutFile(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile)
	PutPkg(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg)
	PutRepository(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository)
	ResetExec()
	ResetFile()
	ResetPkg()
	ResetRepository()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference
type jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) Exec() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference {
	var returns OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference
	_jsii_.Get(
		j,
		"exec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ExecInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec {
	var returns *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec
	_jsii_.Get(
		j,
		"execInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) File() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileOutputReference {
	var returns OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) FileInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile {
	var returns *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) Pkg() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference {
	var returns OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgOutputReference
	_jsii_.Get(
		j,
		"pkg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) PkgInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg {
	var returns *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg
	_jsii_.Get(
		j,
		"pkgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) Repository() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference {
	var returns OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) RepositoryInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository {
	var returns *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigOsPolicyAssignment.OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference_Override(o OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigOsPolicyAssignment.OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) PutExec(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) {
	if err := o.validatePutExecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putExec",
		[]any{value},
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) PutFile(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile) {
	if err := o.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putFile",
		[]any{value},
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) PutPkg(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg) {
	if err := o.validatePutPkgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPkg",
		[]any{value},
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) PutRepository(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) {
	if err := o.validatePutRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRepository",
		[]any{value},
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ResetExec() {
	_jsii_.InvokeVoid(
		o,
		"resetExec",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		o,
		"resetFile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ResetPkg() {
	_jsii_.InvokeVoid(
		o,
		"resetPkg",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ResetRepository() {
	_jsii_.InvokeVoid(
		o,
		"resetRepository",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		o,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
