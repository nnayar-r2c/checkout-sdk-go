package issuing

import (
	"github.com/checkout/checkout-sdk-go/common"
)

type LifetimeUnit string

const (
	Months LifetimeUnit = "Months"
	Years  LifetimeUnit = "Years"
)

type (
	CardCredentialsQuery struct {
		Credentials string `json:"credentials,omitempty"`
	}
)

type RevokeReason string

const (
	Expired        RevokeReason = "expired"
	ReportedLost   RevokeReason = "reported_lost"
	ReportedStolen RevokeReason = "reported_stolen"
)

type (
	RevokeCardRequest struct {
		Reason RevokeReason `json:"reason,omitempty"`
	}
)

type SuspendReason string

const (
	SuspectedLost   SuspendReason = "suspected_lost"
	SuspectedStolen SuspendReason = "suspected_stolen"
)

type (
	SuspendCardRequest struct {
		Reason SuspendReason `json:"reason,omitempty"`
	}
)

type (
	CardLifetime struct {
		Unit  LifetimeUnit `json:"unit,omitempty"`
		Value int          `json:"value,omitempty"`
	}

	ShippingInstruction struct {
		ShippingRecipient string          `json:"shipping_recipient,omitempty"`
		ShippingAddress   *common.Address `json:"shipping_address,omitempty"`
		AdditionalComment string          `json:"additional_comment,omitempty"`
	}

	CardRequest interface {
		GetType() CardType
	}

	CardTypeRequest struct {
		CardRequest
	}

	CardDetailsRequest struct {
		Type          CardType     `json:"type,omitempty"`
		CardholderId  string       `json:"cardholder_id,omitempty"`
		Lifetime      CardLifetime `json:"lifetime"`
		Reference     string       `json:"reference,omitempty"`
		CardProductId string       `json:"card_product_id,omitempty"`
		DisplayName   string       `json:"display_name,omitempty"`
		ActivateCard  bool         `json:"activate_card,omitempty"`
	}

	physicalCardTypeRequest struct {
		CardDetailsRequest
		ShippingInstructions ShippingInstruction `json:"shipping_instructions,omitempty"`
	}

	virtualCardTypeRequest struct {
		CardDetailsRequest
		IsSingleUse bool `json:"is_single_use,omitempty"`
	}
)

func NewPhysicalCardTypeRequest() *physicalCardTypeRequest {
	return &physicalCardTypeRequest{
		CardDetailsRequest: CardDetailsRequest{Type: Physical},
	}
}

func NewVirtualCardTypeRequest() *virtualCardTypeRequest {
	return &virtualCardTypeRequest{
		CardDetailsRequest: CardDetailsRequest{Type: Virtual},
	}
}

func (c *physicalCardTypeRequest) GetType() CardType {
	return c.Type
}

func (c *virtualCardTypeRequest) GetType() CardType {
	return c.Type
}
