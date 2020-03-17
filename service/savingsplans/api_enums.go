// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package savingsplans

type CurrencyCode string

// Enum values for CurrencyCode
const (
	CurrencyCodeCny CurrencyCode = "CNY"
	CurrencyCodeUsd CurrencyCode = "USD"
)

func (enum CurrencyCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CurrencyCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SavingsPlanOfferingFilterAttribute string

// Enum values for SavingsPlanOfferingFilterAttribute
const (
	SavingsPlanOfferingFilterAttributeRegion         SavingsPlanOfferingFilterAttribute = "region"
	SavingsPlanOfferingFilterAttributeInstanceFamily SavingsPlanOfferingFilterAttribute = "instanceFamily"
)

func (enum SavingsPlanOfferingFilterAttribute) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SavingsPlanOfferingFilterAttribute) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SavingsPlanOfferingPropertyKey string

// Enum values for SavingsPlanOfferingPropertyKey
const (
	SavingsPlanOfferingPropertyKeyRegion         SavingsPlanOfferingPropertyKey = "region"
	SavingsPlanOfferingPropertyKeyInstanceFamily SavingsPlanOfferingPropertyKey = "instanceFamily"
)

func (enum SavingsPlanOfferingPropertyKey) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SavingsPlanOfferingPropertyKey) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SavingsPlanPaymentOption string

// Enum values for SavingsPlanPaymentOption
const (
	SavingsPlanPaymentOptionAllUpfront     SavingsPlanPaymentOption = "All Upfront"
	SavingsPlanPaymentOptionPartialUpfront SavingsPlanPaymentOption = "Partial Upfront"
	SavingsPlanPaymentOptionNoUpfront      SavingsPlanPaymentOption = "No Upfront"
)

func (enum SavingsPlanPaymentOption) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SavingsPlanPaymentOption) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SavingsPlanProductType string

// Enum values for SavingsPlanProductType
const (
	SavingsPlanProductTypeEc2     SavingsPlanProductType = "EC2"
	SavingsPlanProductTypeFargate SavingsPlanProductType = "Fargate"
	SavingsPlanProductTypeLambda  SavingsPlanProductType = "Lambda"
)

func (enum SavingsPlanProductType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SavingsPlanProductType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SavingsPlanRateFilterAttribute string

// Enum values for SavingsPlanRateFilterAttribute
const (
	SavingsPlanRateFilterAttributeRegion             SavingsPlanRateFilterAttribute = "region"
	SavingsPlanRateFilterAttributeInstanceFamily     SavingsPlanRateFilterAttribute = "instanceFamily"
	SavingsPlanRateFilterAttributeInstanceType       SavingsPlanRateFilterAttribute = "instanceType"
	SavingsPlanRateFilterAttributeProductDescription SavingsPlanRateFilterAttribute = "productDescription"
	SavingsPlanRateFilterAttributeTenancy            SavingsPlanRateFilterAttribute = "tenancy"
	SavingsPlanRateFilterAttributeProductId          SavingsPlanRateFilterAttribute = "productId"
)

func (enum SavingsPlanRateFilterAttribute) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SavingsPlanRateFilterAttribute) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SavingsPlanRateFilterName string

// Enum values for SavingsPlanRateFilterName
const (
	SavingsPlanRateFilterNameRegion             SavingsPlanRateFilterName = "region"
	SavingsPlanRateFilterNameInstanceType       SavingsPlanRateFilterName = "instanceType"
	SavingsPlanRateFilterNameProductDescription SavingsPlanRateFilterName = "productDescription"
	SavingsPlanRateFilterNameTenancy            SavingsPlanRateFilterName = "tenancy"
	SavingsPlanRateFilterNameProductType        SavingsPlanRateFilterName = "productType"
	SavingsPlanRateFilterNameServiceCode        SavingsPlanRateFilterName = "serviceCode"
	SavingsPlanRateFilterNameUsageType          SavingsPlanRateFilterName = "usageType"
	SavingsPlanRateFilterNameOperation          SavingsPlanRateFilterName = "operation"
)

func (enum SavingsPlanRateFilterName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SavingsPlanRateFilterName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SavingsPlanRatePropertyKey string

// Enum values for SavingsPlanRatePropertyKey
const (
	SavingsPlanRatePropertyKeyRegion             SavingsPlanRatePropertyKey = "region"
	SavingsPlanRatePropertyKeyInstanceType       SavingsPlanRatePropertyKey = "instanceType"
	SavingsPlanRatePropertyKeyInstanceFamily     SavingsPlanRatePropertyKey = "instanceFamily"
	SavingsPlanRatePropertyKeyProductDescription SavingsPlanRatePropertyKey = "productDescription"
	SavingsPlanRatePropertyKeyTenancy            SavingsPlanRatePropertyKey = "tenancy"
)

func (enum SavingsPlanRatePropertyKey) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SavingsPlanRatePropertyKey) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SavingsPlanRateServiceCode string

// Enum values for SavingsPlanRateServiceCode
const (
	SavingsPlanRateServiceCodeAmazonEc2 SavingsPlanRateServiceCode = "AmazonEC2"
	SavingsPlanRateServiceCodeAmazonEcs SavingsPlanRateServiceCode = "AmazonECS"
	SavingsPlanRateServiceCodeAwslambda SavingsPlanRateServiceCode = "AWSLambda"
)

func (enum SavingsPlanRateServiceCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SavingsPlanRateServiceCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SavingsPlanRateUnit string

// Enum values for SavingsPlanRateUnit
const (
	SavingsPlanRateUnitHrs            SavingsPlanRateUnit = "Hrs"
	SavingsPlanRateUnitLambdaGbSecond SavingsPlanRateUnit = "Lambda-GB-Second"
	SavingsPlanRateUnitRequest        SavingsPlanRateUnit = "Request"
)

func (enum SavingsPlanRateUnit) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SavingsPlanRateUnit) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SavingsPlanState string

// Enum values for SavingsPlanState
const (
	SavingsPlanStatePaymentPending SavingsPlanState = "payment-pending"
	SavingsPlanStatePaymentFailed  SavingsPlanState = "payment-failed"
	SavingsPlanStateActive         SavingsPlanState = "active"
	SavingsPlanStateRetired        SavingsPlanState = "retired"
)

func (enum SavingsPlanState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SavingsPlanState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SavingsPlanType string

// Enum values for SavingsPlanType
const (
	SavingsPlanTypeCompute     SavingsPlanType = "Compute"
	SavingsPlanTypeEc2instance SavingsPlanType = "EC2Instance"
)

func (enum SavingsPlanType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SavingsPlanType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SavingsPlansFilterName string

// Enum values for SavingsPlansFilterName
const (
	SavingsPlansFilterNameRegion            SavingsPlansFilterName = "region"
	SavingsPlansFilterNameEc2InstanceFamily SavingsPlansFilterName = "ec2-instance-family"
	SavingsPlansFilterNameCommitment        SavingsPlansFilterName = "commitment"
	SavingsPlansFilterNameUpfront           SavingsPlansFilterName = "upfront"
	SavingsPlansFilterNameTerm              SavingsPlansFilterName = "term"
	SavingsPlansFilterNameSavingsPlanType   SavingsPlansFilterName = "savings-plan-type"
	SavingsPlansFilterNamePaymentOption     SavingsPlansFilterName = "payment-option"
	SavingsPlansFilterNameStart             SavingsPlansFilterName = "start"
	SavingsPlansFilterNameEnd               SavingsPlansFilterName = "end"
)

func (enum SavingsPlansFilterName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SavingsPlansFilterName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
