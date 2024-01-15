package testconfig

import (
	"errors"

	vrfv2 "github.com/smartcontractkit/chainlink/integration-tests/testconfig/vrfv2"
)

type Config struct {
	Common            *Common                  `toml:"Common"`
	General           *General                 `toml:"General"`
	ExistingEnvConfig *ExistingEnvConfig       `toml:"ExistingEnv"`
	NewEnvConfig      *NewEnvConfig            `toml:"NewEnv"`
	Performance       *vrfv2.PerformanceConfig `toml:"Performance"`
}

func (c *Config) Validate() error {
	if c.Common != nil {
		if err := c.Common.Validate(); err != nil {
			return err
		}
	}
	if c.General != nil {
		if err := c.General.Validate(); err != nil {
			return err
		}
	}
	if c.Performance != nil {
		if err := c.Performance.Validate(); err != nil {
			return err
		}
		if *c.Performance.UseExistingEnv {
			if c.ExistingEnvConfig != nil {
				if err := c.ExistingEnvConfig.Validate(); err != nil {
					return err
				}
			}
		} else {
			if c.NewEnvConfig != nil {
				if err := c.NewEnvConfig.Validate(); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

type Common struct {
	*vrfv2.Common
}

func (c *Common) Validate() error {
	if c.Common == nil {
		return nil
	}
	return c.Common.Validate()
}

type General struct {
	*vrfv2.General
	SubscriptionFundingAmountNative *float64 `toml:"subscription_funding_amount_native"` // Amount of LINK to fund the subscription with
	FulfillmentFlatFeeLinkPPM       *uint32  `toml:"fulfillment_flat_fee_link_ppm"`      // Flat fee in ppm for LINK for the VRF Coordinator config
	FulfillmentFlatFeeNativePPM     *uint32  `toml:"fulfillment_flat_fee_native_ppm"`    // Flat fee in ppm for native currency for the VRF Coordinator config
}

func (c *General) Validate() error {
	if err := c.General.Validate(); err != nil {
		return err
	}
	if c.SubscriptionFundingAmountNative == nil || *c.SubscriptionFundingAmountNative <= 0 {
		return errors.New("subscription_funding_amount_native must be greater than 0")
	}
	if c.FulfillmentFlatFeeLinkPPM == nil || *c.FulfillmentFlatFeeLinkPPM <= 0 {
		return errors.New("fulfillment_flat_fee_link_ppm must be greater than 0")
	}
	if c.FulfillmentFlatFeeNativePPM == nil || *c.FulfillmentFlatFeeNativePPM <= 0 {
		return errors.New("fulfillment_flat_fee_native_ppm must be greater than 0")
	}

	return nil
}

type NewEnvConfig struct {
	*Funding
}

func (c *NewEnvConfig) Validate() error {
	if c.Funding == nil {
		return nil
	}

	return c.Funding.Validate()
}

type ExistingEnvConfig struct {
	*vrfv2.ExistingEnvConfig
	*Funding
}

func (c *ExistingEnvConfig) Validate() error {
	if c.ExistingEnvConfig != nil {
		if err := c.ExistingEnvConfig.Validate(); err != nil {
			return err
		}
	}
	if c.Funding != nil {
		if err := c.Funding.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type Funding struct {
	*SubFunding
	NodeSendingKeyFunding    *float64 `toml:"node_sending_key_funding"`
	NodeSendingKeyFundingMin *float64 `toml:"node_sending_key_funding_min"`
}

func (c *Funding) Validate() error {
	if c.NodeSendingKeyFunding != nil && *c.NodeSendingKeyFunding <= 0 {
		return errors.New("when set node_sending_key_funding must be a positive value")
	}
	if c.NodeSendingKeyFundingMin != nil && *c.NodeSendingKeyFundingMin <= 0 {
		return errors.New("when set node_sending_key_funding_min must be a positive value")
	}

	return c.SubFunding.Validate()
}

type SubFunding struct {
	SubFundsLink   *float64 `toml:"sub_funds_link"`
	SubFundsNative *float64 `toml:"sub_funds_native"`
}

func (c *SubFunding) Validate() error {
	if c.SubFundsLink == nil && c.SubFundsNative == nil {
		return errors.New("at least one of sub_funds_link or sub_funds_native must be set")
	}
	if c.SubFundsLink != nil && *c.SubFundsLink <= 0 {
		return errors.New("sub_funds_link must be greater than 0")
	}
	if c.SubFundsNative != nil && *c.SubFundsNative <= 0 {
		return errors.New("sub_funds_native must be greater than 0")
	}

	return nil
}