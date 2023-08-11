package googlecloudrunv2service

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudrunv2service/internal"
)

type GoogleCloudRunV2ServiceTemplateOutputReference interface {
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
	Containers() GoogleCloudRunV2ServiceTemplateContainersList
	ContainersInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EncryptionKey() *string
	SetEncryptionKey(val *string)
	EncryptionKeyInput() *string
	ExecutionEnvironment() *string
	SetExecutionEnvironment(val *string)
	ExecutionEnvironmentInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCloudRunV2ServiceTemplate
	SetInternalValue(val *GoogleCloudRunV2ServiceTemplate)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	MaxInstanceRequestConcurrency() *float64
	SetMaxInstanceRequestConcurrency(val *float64)
	MaxInstanceRequestConcurrencyInput() *float64
	Revision() *string
	SetRevision(val *string)
	RevisionInput() *string
	Scaling() GoogleCloudRunV2ServiceTemplateScalingOutputReference
	ScalingInput() *GoogleCloudRunV2ServiceTemplateScaling
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	SessionAffinity() interface{}
	SetSessionAffinity(val interface{})
	SessionAffinityInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
	Volumes() GoogleCloudRunV2ServiceTemplateVolumesList
	VolumesInput() interface{}
	VpcAccess() GoogleCloudRunV2ServiceTemplateVpcAccessOutputReference
	VpcAccessInput() *GoogleCloudRunV2ServiceTemplateVpcAccess
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
	PutScaling(value *GoogleCloudRunV2ServiceTemplateScaling)
	PutVolumes(value interface{})
	PutVpcAccess(value *GoogleCloudRunV2ServiceTemplateVpcAccess)
	ResetAnnotations()
	ResetContainers()
	ResetEncryptionKey()
	ResetExecutionEnvironment()
	ResetLabels()
	ResetMaxInstanceRequestConcurrency()
	ResetRevision()
	ResetScaling()
	ResetServiceAccount()
	ResetSessionAffinity()
	ResetTimeout()
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

// The jsii proxy struct for GoogleCloudRunV2ServiceTemplateOutputReference
type jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) Containers() GoogleCloudRunV2ServiceTemplateContainersList {
	var returns GoogleCloudRunV2ServiceTemplateContainersList
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ContainersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) EncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) EncryptionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ExecutionEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ExecutionEnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) InternalValue() *GoogleCloudRunV2ServiceTemplate {
	var returns *GoogleCloudRunV2ServiceTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) MaxInstanceRequestConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceRequestConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) MaxInstanceRequestConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceRequestConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) Revision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) RevisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) Scaling() GoogleCloudRunV2ServiceTemplateScalingOutputReference {
	var returns GoogleCloudRunV2ServiceTemplateScalingOutputReference
	_jsii_.Get(
		j,
		"scaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ScalingInput() *GoogleCloudRunV2ServiceTemplateScaling {
	var returns *GoogleCloudRunV2ServiceTemplateScaling
	_jsii_.Get(
		j,
		"scalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) SessionAffinity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) SessionAffinityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) Volumes() GoogleCloudRunV2ServiceTemplateVolumesList {
	var returns GoogleCloudRunV2ServiceTemplateVolumesList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) VolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) VpcAccess() GoogleCloudRunV2ServiceTemplateVpcAccessOutputReference {
	var returns GoogleCloudRunV2ServiceTemplateVpcAccessOutputReference
	_jsii_.Get(
		j,
		"vpcAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) VpcAccessInput() *GoogleCloudRunV2ServiceTemplateVpcAccess {
	var returns *GoogleCloudRunV2ServiceTemplateVpcAccess
	_jsii_.Get(
		j,
		"vpcAccessInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudRunV2ServiceTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudRunV2ServiceTemplateOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudRunV2ServiceTemplateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudRunV2Service.GoogleCloudRunV2ServiceTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudRunV2ServiceTemplateOutputReference_Override(g GoogleCloudRunV2ServiceTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudRunV2Service.GoogleCloudRunV2ServiceTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference)SetEncryptionKey(val *string) {
	if err := j.validateSetEncryptionKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKey",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference)SetExecutionEnvironment(val *string) {
	if err := j.validateSetExecutionEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionEnvironment",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference)SetInternalValue(val *GoogleCloudRunV2ServiceTemplate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference)SetMaxInstanceRequestConcurrency(val *float64) {
	if err := j.validateSetMaxInstanceRequestConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstanceRequestConcurrency",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference)SetRevision(val *string) {
	if err := j.validateSetRevisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revision",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference)SetSessionAffinity(val interface{}) {
	if err := j.validateSetSessionAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAffinity",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) PutContainers(value interface{}) {
	if err := g.validatePutContainersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContainers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) PutScaling(value *GoogleCloudRunV2ServiceTemplateScaling) {
	if err := g.validatePutScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScaling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) PutVolumes(value interface{}) {
	if err := g.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVolumes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) PutVpcAccess(value *GoogleCloudRunV2ServiceTemplateVpcAccess) {
	if err := g.validatePutVpcAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVpcAccess",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ResetContainers() {
	_jsii_.InvokeVoid(
		g,
		"resetContainers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ResetEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ResetExecutionEnvironment() {
	_jsii_.InvokeVoid(
		g,
		"resetExecutionEnvironment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ResetMaxInstanceRequestConcurrency() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxInstanceRequestConcurrency",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ResetRevision() {
	_jsii_.InvokeVoid(
		g,
		"resetRevision",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ResetScaling() {
	_jsii_.InvokeVoid(
		g,
		"resetScaling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ResetSessionAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetSessionAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		g,
		"resetVolumes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ResetVpcAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

