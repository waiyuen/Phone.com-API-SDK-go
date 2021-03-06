/* 
 * Phone.com API
 *
 * This is a Phone.com api Swagger definition
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apisupport@phone.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type CreatePricingParams struct {

	// Pricing plan code
	PricingId int32 `json:"pricing_id,omitempty"`

	// Reason this pricing plan is added to the subaccount
	Reason string `json:"reason,omitempty"`

	// Pricing plan expiration time in UNIX format. Disregard or set it to null for plan which never expires
	ExpireDate int32 `json:"expire_date,omitempty"`
}
