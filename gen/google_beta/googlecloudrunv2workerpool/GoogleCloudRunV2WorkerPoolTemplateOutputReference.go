package googlecloudrunv2workerpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudrunv2workerpool/internal"
)

type GoogleCloudRunV2WorkerPoolTemplateOutputReference interface {
	cdktf.ComplexObject
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
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
	Containers() GoogleCloudRunV2WorkerPoolTemplateContainersList
	ContainersInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EncryptionKey() *string
	SetEncryptionKey(val *string)
	EncryptionKeyInput() *string
	EncryptionKeyRevocationAction() *string
	SetEncryptionKeyRevocationAction(val *string)
	EncryptionKeyRevocationActionInput() *string
	EncryptionKeyShutdownDuration() *string
	SetEncryptionKeyShutdownDuration(val *string)
	EncryptionKeyShutdownDurationInput() *string
	// Experimental.
	Fqn() *string
	GpuZonalRedundancyDisabled() interface{}
	SetGpuZonalRedundancyDisabled(val interface{})
	GpuZonalRedundancyDisabledInput() interface{}
	InternalValue() *GoogleCloudRunV2WorkerPoolTemplate
	SetInternalValue(val *GoogleCloudRunV2WorkerPoolTemplate)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	NodeSelector() GoogleCloudRunV2WorkerPoolTemplateNodeSelectorOutputReference
	NodeSelectorInput() *GoogleCloudRunV2WorkerPoolTemplateNodeSelector
	Revision() *string
	SetRevision(val *string)
	RevisionInput() *string
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Volumes() GoogleCloudRunV2WorkerPoolTemplateVolumesList
	VolumesInput() interface{}
	VpcAccess() GoogleCloudRunV2WorkerPoolTemplateVpcAccessOutputReference
	VpcAccessInput() *GoogleCloudRunV2WorkerPoolTemplateVpcAccess
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
	PutContainers(value interface{})
	PutNodeSelector(value *GoogleCloudRunV2WorkerPoolTemplateNodeSelector)
	PutVolumes(value interface{})
	PutVpcAccess(value *GoogleCloudRunV2WorkerPoolTemplateVpcAccess)
	ResetAnnotations()
	ResetContainers()
	ResetEncryptionKey()
	ResetEncryptionKeyRevocationAction()
	ResetEncryptionKeyShutdownDuration()
	ResetGpuZonalRedundancyDisabled()
	ResetLabels()
	ResetNodeSelector()
	ResetRevision()
	ResetServiceAccount()
	ResetVolumes()
	ResetVpcAccess()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudRunV2WorkerPoolTemplateOutputReference
type jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) Containers() GoogleCloudRunV2WorkerPoolTemplateContainersList {
	var returns GoogleCloudRunV2WorkerPoolTemplateContainersList
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ContainersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) EncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) EncryptionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) EncryptionKeyRevocationAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyRevocationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) EncryptionKeyRevocationActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyRevocationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) EncryptionKeyShutdownDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyShutdownDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) EncryptionKeyShutdownDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyShutdownDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) GpuZonalRedundancyDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gpuZonalRedundancyDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) GpuZonalRedundancyDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gpuZonalRedundancyDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) InternalValue() *GoogleCloudRunV2WorkerPoolTemplate {
	var returns *GoogleCloudRunV2WorkerPoolTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) NodeSelector() GoogleCloudRunV2WorkerPoolTemplateNodeSelectorOutputReference {
	var returns GoogleCloudRunV2WorkerPoolTemplateNodeSelectorOutputReference
	_jsii_.Get(
		j,
		"nodeSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) NodeSelectorInput() *GoogleCloudRunV2WorkerPoolTemplateNodeSelector {
	var returns *GoogleCloudRunV2WorkerPoolTemplateNodeSelector
	_jsii_.Get(
		j,
		"nodeSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) Revision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) RevisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) Volumes() GoogleCloudRunV2WorkerPoolTemplateVolumesList {
	var returns GoogleCloudRunV2WorkerPoolTemplateVolumesList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) VolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) VpcAccess() GoogleCloudRunV2WorkerPoolTemplateVpcAccessOutputReference {
	var returns GoogleCloudRunV2WorkerPoolTemplateVpcAccessOutputReference
	_jsii_.Get(
		j,
		"vpcAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) VpcAccessInput() *GoogleCloudRunV2WorkerPoolTemplateVpcAccess {
	var returns *GoogleCloudRunV2WorkerPoolTemplateVpcAccess
	_jsii_.Get(
		j,
		"vpcAccessInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudRunV2WorkerPoolTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudRunV2WorkerPoolTemplateOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudRunV2WorkerPoolTemplateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudRunV2WorkerPool.GoogleCloudRunV2WorkerPoolTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudRunV2WorkerPoolTemplateOutputReference_Override(g GoogleCloudRunV2WorkerPoolTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudRunV2WorkerPool.GoogleCloudRunV2WorkerPoolTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference)SetEncryptionKey(val *string) {
	if err := j.validateSetEncryptionKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKey",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference)SetEncryptionKeyRevocationAction(val *string) {
	if err := j.validateSetEncryptionKeyRevocationActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKeyRevocationAction",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference)SetEncryptionKeyShutdownDuration(val *string) {
	if err := j.validateSetEncryptionKeyShutdownDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKeyShutdownDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference)SetGpuZonalRedundancyDisabled(val interface{}) {
	if err := j.validateSetGpuZonalRedundancyDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpuZonalRedundancyDisabled",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference)SetInternalValue(val *GoogleCloudRunV2WorkerPoolTemplate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference)SetRevision(val *string) {
	if err := j.validateSetRevisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revision",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) PutContainers(value interface{}) {
	if err := g.validatePutContainersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContainers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) PutNodeSelector(value *GoogleCloudRunV2WorkerPoolTemplateNodeSelector) {
	if err := g.validatePutNodeSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeSelector",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) PutVolumes(value interface{}) {
	if err := g.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVolumes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) PutVpcAccess(value *GoogleCloudRunV2WorkerPoolTemplateVpcAccess) {
	if err := g.validatePutVpcAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVpcAccess",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ResetContainers() {
	_jsii_.InvokeVoid(
		g,
		"resetContainers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ResetEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ResetEncryptionKeyRevocationAction() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionKeyRevocationAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ResetEncryptionKeyShutdownDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionKeyShutdownDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ResetGpuZonalRedundancyDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetGpuZonalRedundancyDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ResetNodeSelector() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeSelector",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ResetRevision() {
	_jsii_.InvokeVoid(
		g,
		"resetRevision",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		g,
		"resetVolumes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ResetVpcAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

