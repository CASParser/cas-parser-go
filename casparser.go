// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/CASParser/cas-parser-go/internal/apijson"
	"github.com/CASParser/cas-parser-go/internal/requestconfig"
	"github.com/CASParser/cas-parser-go/option"
	"github.com/CASParser/cas-parser-go/packages/param"
	"github.com/CASParser/cas-parser-go/packages/respjson"
)

// CasParserService contains methods and other services that help with interacting
// with the CAS Parser API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasParserService] method instead.
type CasParserService struct {
	Options []option.RequestOption
}

// NewCasParserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCasParserService(opts ...option.RequestOption) (r CasParserService) {
	r = CasParserService{}
	r.Options = opts
	return
}

// This endpoint specifically parses CAMS/KFintech CAS (Consolidated Account
// Statement) PDF files and returns data in a unified format. Use this endpoint
// when you know the PDF is from CAMS or KFintech.
func (r *CasParserService) CamsKfintech(ctx context.Context, body CasParserCamsKfintechParams, opts ...option.RequestOption) (res *UnifiedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/cams_kfintech/parse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint specifically parses CDSL CAS (Consolidated Account Statement) PDF
// files and returns data in a unified format. Use this endpoint when you know the
// PDF is from CDSL.
func (r *CasParserService) Cdsl(ctx context.Context, body CasParserCdslParams, opts ...option.RequestOption) (res *UnifiedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/cdsl/parse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint specifically parses NSDL CAS (Consolidated Account Statement) PDF
// files and returns data in a unified format. Use this endpoint when you know the
// PDF is from NSDL.
func (r *CasParserService) Nsdl(ctx context.Context, body CasParserNsdlParams, opts ...option.RequestOption) (res *UnifiedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/nsdl/parse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint parses CAS (Consolidated Account Statement) PDF files from NSDL,
// CDSL, or CAMS/KFintech and returns data in a unified format. It auto-detects the
// CAS type and transforms the data into a consistent structure regardless of the
// source.
func (r *CasParserService) SmartParse(ctx context.Context, body CasParserSmartParseParams, opts ...option.RequestOption) (res *UnifiedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/smart/parse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UnifiedResponse struct {
	DematAccounts []UnifiedResponseDematAccount `json:"demat_accounts"`
	Insurance     UnifiedResponseInsurance      `json:"insurance"`
	Investor      UnifiedResponseInvestor       `json:"investor"`
	Meta          UnifiedResponseMeta           `json:"meta"`
	MutualFunds   []UnifiedResponseMutualFund   `json:"mutual_funds"`
	// List of NPS accounts
	Nps     []UnifiedResponseNp    `json:"nps"`
	Summary UnifiedResponseSummary `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DematAccounts respjson.Field
		Insurance     respjson.Field
		Investor      respjson.Field
		Meta          respjson.Field
		MutualFunds   respjson.Field
		Nps           respjson.Field
		Summary       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponse) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseDematAccount struct {
	// Additional information specific to the demat account type
	AdditionalInfo UnifiedResponseDematAccountAdditionalInfo `json:"additional_info"`
	// Beneficiary Owner ID (primarily for CDSL)
	BoID string `json:"bo_id"`
	// Client ID
	ClientID string `json:"client_id"`
	// Type of demat account
	//
	// Any of "NSDL", "CDSL".
	DematType string `json:"demat_type"`
	// Depository Participant ID
	DpID string `json:"dp_id"`
	// Depository Participant name
	DpName   string                              `json:"dp_name"`
	Holdings UnifiedResponseDematAccountHoldings `json:"holdings"`
	// List of account holders linked to this demat account
	LinkedHolders []UnifiedResponseDematAccountLinkedHolder `json:"linked_holders"`
	// Total value of the demat account
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		BoID           respjson.Field
		ClientID       respjson.Field
		DematType      respjson.Field
		DpID           respjson.Field
		DpName         respjson.Field
		Holdings       respjson.Field
		LinkedHolders  respjson.Field
		Value          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccount) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseDematAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information specific to the demat account type
type UnifiedResponseDematAccountAdditionalInfo struct {
	// Beneficiary Owner status (CDSL)
	BoStatus string `json:"bo_status"`
	// Beneficiary Owner sub-status (CDSL)
	BoSubStatus string `json:"bo_sub_status"`
	// Beneficiary Owner type (CDSL)
	BoType string `json:"bo_type"`
	// Basic Services Demat Account status (CDSL)
	Bsda string `json:"bsda"`
	// Email associated with the demat account (CDSL)
	Email string `json:"email" format:"email"`
	// List of linked PAN numbers (NSDL)
	LinkedPans []string `json:"linked_pans"`
	// Nominee details (CDSL)
	Nominee string `json:"nominee"`
	// Account status (CDSL)
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoStatus    respjson.Field
		BoSubStatus respjson.Field
		BoType      respjson.Field
		Bsda        respjson.Field
		Email       respjson.Field
		LinkedPans  respjson.Field
		Nominee     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountAdditionalInfo) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseDematAccountAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseDematAccountHoldings struct {
	Aifs                 []UnifiedResponseDematAccountHoldingsAif                `json:"aifs"`
	CorporateBonds       []UnifiedResponseDematAccountHoldingsCorporateBond      `json:"corporate_bonds"`
	DematMutualFunds     []UnifiedResponseDematAccountHoldingsDematMutualFund    `json:"demat_mutual_funds"`
	Equities             []UnifiedResponseDematAccountHoldingsEquity             `json:"equities"`
	GovernmentSecurities []UnifiedResponseDematAccountHoldingsGovernmentSecurity `json:"government_securities"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aifs                 respjson.Field
		CorporateBonds       respjson.Field
		DematMutualFunds     respjson.Field
		Equities             respjson.Field
		GovernmentSecurities respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldings) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseDematAccountHoldings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseDematAccountHoldingsAif struct {
	// Additional information specific to the AIF
	AdditionalInfo UnifiedResponseDematAccountHoldingsAifAdditionalInfo `json:"additional_info"`
	// ISIN code of the AIF
	Isin string `json:"isin"`
	// Name of the AIF
	Name string `json:"name"`
	// List of transactions for this holding (beta)
	Transactions []UnifiedResponseDematAccountHoldingsAifTransaction `json:"transactions"`
	// Number of units held
	Units float64 `json:"units"`
	// Current market value of the holding
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Isin           respjson.Field
		Name           respjson.Field
		Transactions   respjson.Field
		Units          respjson.Field
		Value          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsAif) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseDematAccountHoldingsAif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information specific to the AIF
type UnifiedResponseDematAccountHoldingsAifAdditionalInfo struct {
	// Closing balance units for the statement period (beta)
	CloseUnits float64 `json:"close_units,nullable"`
	// Opening balance units for the statement period (beta)
	OpenUnits float64 `json:"open_units,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CloseUnits  respjson.Field
		OpenUnits   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsAifAdditionalInfo) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseDematAccountHoldingsAifAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unified transaction schema for all holding types (MF folios, equities, bonds,
// etc.)
type UnifiedResponseDematAccountHoldingsAifTransaction struct {
	// Additional transaction-specific fields that vary by source
	AdditionalInfo UnifiedResponseDematAccountHoldingsAifTransactionAdditionalInfo `json:"additional_info"`
	// Transaction amount in currency (computed from units × price/NAV)
	Amount float64 `json:"amount,nullable"`
	// Balance units after transaction
	Balance float64 `json:"balance"`
	// Transaction date (YYYY-MM-DD)
	Date time.Time `json:"date" format:"date"`
	// Transaction description/particulars
	Description string `json:"description"`
	// Dividend rate (for DIVIDEND_PAYOUT transactions)
	DividendRate float64 `json:"dividend_rate,nullable"`
	// NAV/price per unit on transaction date
	Nav float64 `json:"nav,nullable"`
	// Transaction type. Possible values are PURCHASE, PURCHASE_SIP, REDEMPTION,
	// SWITCH_IN, SWITCH_IN_MERGER, SWITCH_OUT, SWITCH_OUT_MERGER, DIVIDEND_PAYOUT,
	// DIVIDEND_REINVEST, SEGREGATION, STAMP_DUTY_TAX, TDS_TAX, STT_TAX, MISC,
	// REVERSAL, UNKNOWN.
	//
	// Any of "PURCHASE", "PURCHASE_SIP", "REDEMPTION", "SWITCH_IN",
	// "SWITCH_IN_MERGER", "SWITCH_OUT", "SWITCH_OUT_MERGER", "DIVIDEND_PAYOUT",
	// "DIVIDEND_REINVEST", "SEGREGATION", "STAMP_DUTY_TAX", "TDS_TAX", "STT_TAX",
	// "MISC", "REVERSAL", "UNKNOWN".
	Type string `json:"type"`
	// Number of units involved in transaction
	Units float64 `json:"units"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Amount         respjson.Field
		Balance        respjson.Field
		Date           respjson.Field
		Description    respjson.Field
		DividendRate   respjson.Field
		Nav            respjson.Field
		Type           respjson.Field
		Units          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsAifTransaction) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseDematAccountHoldingsAifTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional transaction-specific fields that vary by source
type UnifiedResponseDematAccountHoldingsAifTransactionAdditionalInfo struct {
	// Capital withdrawal amount (CDSL MF transactions)
	CapitalWithdrawal float64 `json:"capital_withdrawal"`
	// Units credited (demat transactions)
	Credit float64 `json:"credit"`
	// Units debited (demat transactions)
	Debit float64 `json:"debit"`
	// Income distribution amount (CDSL MF transactions)
	IncomeDistribution float64 `json:"income_distribution"`
	// Order/transaction reference number (demat transactions)
	OrderNo string `json:"order_no"`
	// Price per unit (NSDL/CDSL MF transactions)
	Price float64 `json:"price"`
	// Stamp duty charged
	StampDuty float64 `json:"stamp_duty"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapitalWithdrawal  respjson.Field
		Credit             respjson.Field
		Debit              respjson.Field
		IncomeDistribution respjson.Field
		OrderNo            respjson.Field
		Price              respjson.Field
		StampDuty          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsAifTransactionAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *UnifiedResponseDematAccountHoldingsAifTransactionAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseDematAccountHoldingsCorporateBond struct {
	// Additional information specific to the corporate bond
	AdditionalInfo UnifiedResponseDematAccountHoldingsCorporateBondAdditionalInfo `json:"additional_info"`
	// ISIN code of the corporate bond
	Isin string `json:"isin"`
	// Name of the corporate bond
	Name string `json:"name"`
	// List of transactions for this holding (beta)
	Transactions []UnifiedResponseDematAccountHoldingsCorporateBondTransaction `json:"transactions"`
	// Number of units held
	Units float64 `json:"units"`
	// Current market value of the holding
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Isin           respjson.Field
		Name           respjson.Field
		Transactions   respjson.Field
		Units          respjson.Field
		Value          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsCorporateBond) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseDematAccountHoldingsCorporateBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information specific to the corporate bond
type UnifiedResponseDematAccountHoldingsCorporateBondAdditionalInfo struct {
	// Closing balance units for the statement period (beta)
	CloseUnits float64 `json:"close_units,nullable"`
	// Opening balance units for the statement period (beta)
	OpenUnits float64 `json:"open_units,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CloseUnits  respjson.Field
		OpenUnits   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsCorporateBondAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *UnifiedResponseDematAccountHoldingsCorporateBondAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unified transaction schema for all holding types (MF folios, equities, bonds,
// etc.)
type UnifiedResponseDematAccountHoldingsCorporateBondTransaction struct {
	// Additional transaction-specific fields that vary by source
	AdditionalInfo UnifiedResponseDematAccountHoldingsCorporateBondTransactionAdditionalInfo `json:"additional_info"`
	// Transaction amount in currency (computed from units × price/NAV)
	Amount float64 `json:"amount,nullable"`
	// Balance units after transaction
	Balance float64 `json:"balance"`
	// Transaction date (YYYY-MM-DD)
	Date time.Time `json:"date" format:"date"`
	// Transaction description/particulars
	Description string `json:"description"`
	// Dividend rate (for DIVIDEND_PAYOUT transactions)
	DividendRate float64 `json:"dividend_rate,nullable"`
	// NAV/price per unit on transaction date
	Nav float64 `json:"nav,nullable"`
	// Transaction type. Possible values are PURCHASE, PURCHASE_SIP, REDEMPTION,
	// SWITCH_IN, SWITCH_IN_MERGER, SWITCH_OUT, SWITCH_OUT_MERGER, DIVIDEND_PAYOUT,
	// DIVIDEND_REINVEST, SEGREGATION, STAMP_DUTY_TAX, TDS_TAX, STT_TAX, MISC,
	// REVERSAL, UNKNOWN.
	//
	// Any of "PURCHASE", "PURCHASE_SIP", "REDEMPTION", "SWITCH_IN",
	// "SWITCH_IN_MERGER", "SWITCH_OUT", "SWITCH_OUT_MERGER", "DIVIDEND_PAYOUT",
	// "DIVIDEND_REINVEST", "SEGREGATION", "STAMP_DUTY_TAX", "TDS_TAX", "STT_TAX",
	// "MISC", "REVERSAL", "UNKNOWN".
	Type string `json:"type"`
	// Number of units involved in transaction
	Units float64 `json:"units"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Amount         respjson.Field
		Balance        respjson.Field
		Date           respjson.Field
		Description    respjson.Field
		DividendRate   respjson.Field
		Nav            respjson.Field
		Type           respjson.Field
		Units          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsCorporateBondTransaction) RawJSON() string {
	return r.JSON.raw
}
func (r *UnifiedResponseDematAccountHoldingsCorporateBondTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional transaction-specific fields that vary by source
type UnifiedResponseDematAccountHoldingsCorporateBondTransactionAdditionalInfo struct {
	// Capital withdrawal amount (CDSL MF transactions)
	CapitalWithdrawal float64 `json:"capital_withdrawal"`
	// Units credited (demat transactions)
	Credit float64 `json:"credit"`
	// Units debited (demat transactions)
	Debit float64 `json:"debit"`
	// Income distribution amount (CDSL MF transactions)
	IncomeDistribution float64 `json:"income_distribution"`
	// Order/transaction reference number (demat transactions)
	OrderNo string `json:"order_no"`
	// Price per unit (NSDL/CDSL MF transactions)
	Price float64 `json:"price"`
	// Stamp duty charged
	StampDuty float64 `json:"stamp_duty"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapitalWithdrawal  respjson.Field
		Credit             respjson.Field
		Debit              respjson.Field
		IncomeDistribution respjson.Field
		OrderNo            respjson.Field
		Price              respjson.Field
		StampDuty          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsCorporateBondTransactionAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *UnifiedResponseDematAccountHoldingsCorporateBondTransactionAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseDematAccountHoldingsDematMutualFund struct {
	// Additional information specific to the mutual fund
	AdditionalInfo UnifiedResponseDematAccountHoldingsDematMutualFundAdditionalInfo `json:"additional_info"`
	// ISIN code of the mutual fund
	Isin string `json:"isin"`
	// Name of the mutual fund
	Name string `json:"name"`
	// List of transactions for this holding (beta)
	Transactions []UnifiedResponseDematAccountHoldingsDematMutualFundTransaction `json:"transactions"`
	// Number of units held
	Units float64 `json:"units"`
	// Current market value of the holding
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Isin           respjson.Field
		Name           respjson.Field
		Transactions   respjson.Field
		Units          respjson.Field
		Value          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsDematMutualFund) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseDematAccountHoldingsDematMutualFund) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information specific to the mutual fund
type UnifiedResponseDematAccountHoldingsDematMutualFundAdditionalInfo struct {
	// Closing balance units for the statement period (beta)
	CloseUnits float64 `json:"close_units,nullable"`
	// Opening balance units for the statement period (beta)
	OpenUnits float64 `json:"open_units,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CloseUnits  respjson.Field
		OpenUnits   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsDematMutualFundAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *UnifiedResponseDematAccountHoldingsDematMutualFundAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unified transaction schema for all holding types (MF folios, equities, bonds,
// etc.)
type UnifiedResponseDematAccountHoldingsDematMutualFundTransaction struct {
	// Additional transaction-specific fields that vary by source
	AdditionalInfo UnifiedResponseDematAccountHoldingsDematMutualFundTransactionAdditionalInfo `json:"additional_info"`
	// Transaction amount in currency (computed from units × price/NAV)
	Amount float64 `json:"amount,nullable"`
	// Balance units after transaction
	Balance float64 `json:"balance"`
	// Transaction date (YYYY-MM-DD)
	Date time.Time `json:"date" format:"date"`
	// Transaction description/particulars
	Description string `json:"description"`
	// Dividend rate (for DIVIDEND_PAYOUT transactions)
	DividendRate float64 `json:"dividend_rate,nullable"`
	// NAV/price per unit on transaction date
	Nav float64 `json:"nav,nullable"`
	// Transaction type. Possible values are PURCHASE, PURCHASE_SIP, REDEMPTION,
	// SWITCH_IN, SWITCH_IN_MERGER, SWITCH_OUT, SWITCH_OUT_MERGER, DIVIDEND_PAYOUT,
	// DIVIDEND_REINVEST, SEGREGATION, STAMP_DUTY_TAX, TDS_TAX, STT_TAX, MISC,
	// REVERSAL, UNKNOWN.
	//
	// Any of "PURCHASE", "PURCHASE_SIP", "REDEMPTION", "SWITCH_IN",
	// "SWITCH_IN_MERGER", "SWITCH_OUT", "SWITCH_OUT_MERGER", "DIVIDEND_PAYOUT",
	// "DIVIDEND_REINVEST", "SEGREGATION", "STAMP_DUTY_TAX", "TDS_TAX", "STT_TAX",
	// "MISC", "REVERSAL", "UNKNOWN".
	Type string `json:"type"`
	// Number of units involved in transaction
	Units float64 `json:"units"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Amount         respjson.Field
		Balance        respjson.Field
		Date           respjson.Field
		Description    respjson.Field
		DividendRate   respjson.Field
		Nav            respjson.Field
		Type           respjson.Field
		Units          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsDematMutualFundTransaction) RawJSON() string {
	return r.JSON.raw
}
func (r *UnifiedResponseDematAccountHoldingsDematMutualFundTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional transaction-specific fields that vary by source
type UnifiedResponseDematAccountHoldingsDematMutualFundTransactionAdditionalInfo struct {
	// Capital withdrawal amount (CDSL MF transactions)
	CapitalWithdrawal float64 `json:"capital_withdrawal"`
	// Units credited (demat transactions)
	Credit float64 `json:"credit"`
	// Units debited (demat transactions)
	Debit float64 `json:"debit"`
	// Income distribution amount (CDSL MF transactions)
	IncomeDistribution float64 `json:"income_distribution"`
	// Order/transaction reference number (demat transactions)
	OrderNo string `json:"order_no"`
	// Price per unit (NSDL/CDSL MF transactions)
	Price float64 `json:"price"`
	// Stamp duty charged
	StampDuty float64 `json:"stamp_duty"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapitalWithdrawal  respjson.Field
		Credit             respjson.Field
		Debit              respjson.Field
		IncomeDistribution respjson.Field
		OrderNo            respjson.Field
		Price              respjson.Field
		StampDuty          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsDematMutualFundTransactionAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *UnifiedResponseDematAccountHoldingsDematMutualFundTransactionAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseDematAccountHoldingsEquity struct {
	// Additional information specific to the equity
	AdditionalInfo UnifiedResponseDematAccountHoldingsEquityAdditionalInfo `json:"additional_info"`
	// ISIN code of the equity
	Isin string `json:"isin"`
	// Name of the equity
	Name string `json:"name"`
	// List of transactions for this holding (beta)
	Transactions []UnifiedResponseDematAccountHoldingsEquityTransaction `json:"transactions"`
	// Number of units held
	Units float64 `json:"units"`
	// Current market value of the holding
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Isin           respjson.Field
		Name           respjson.Field
		Transactions   respjson.Field
		Units          respjson.Field
		Value          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsEquity) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseDematAccountHoldingsEquity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information specific to the equity
type UnifiedResponseDematAccountHoldingsEquityAdditionalInfo struct {
	// Closing balance units for the statement period (beta)
	CloseUnits float64 `json:"close_units,nullable"`
	// Opening balance units for the statement period (beta)
	OpenUnits float64 `json:"open_units,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CloseUnits  respjson.Field
		OpenUnits   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsEquityAdditionalInfo) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseDematAccountHoldingsEquityAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unified transaction schema for all holding types (MF folios, equities, bonds,
// etc.)
type UnifiedResponseDematAccountHoldingsEquityTransaction struct {
	// Additional transaction-specific fields that vary by source
	AdditionalInfo UnifiedResponseDematAccountHoldingsEquityTransactionAdditionalInfo `json:"additional_info"`
	// Transaction amount in currency (computed from units × price/NAV)
	Amount float64 `json:"amount,nullable"`
	// Balance units after transaction
	Balance float64 `json:"balance"`
	// Transaction date (YYYY-MM-DD)
	Date time.Time `json:"date" format:"date"`
	// Transaction description/particulars
	Description string `json:"description"`
	// Dividend rate (for DIVIDEND_PAYOUT transactions)
	DividendRate float64 `json:"dividend_rate,nullable"`
	// NAV/price per unit on transaction date
	Nav float64 `json:"nav,nullable"`
	// Transaction type. Possible values are PURCHASE, PURCHASE_SIP, REDEMPTION,
	// SWITCH_IN, SWITCH_IN_MERGER, SWITCH_OUT, SWITCH_OUT_MERGER, DIVIDEND_PAYOUT,
	// DIVIDEND_REINVEST, SEGREGATION, STAMP_DUTY_TAX, TDS_TAX, STT_TAX, MISC,
	// REVERSAL, UNKNOWN.
	//
	// Any of "PURCHASE", "PURCHASE_SIP", "REDEMPTION", "SWITCH_IN",
	// "SWITCH_IN_MERGER", "SWITCH_OUT", "SWITCH_OUT_MERGER", "DIVIDEND_PAYOUT",
	// "DIVIDEND_REINVEST", "SEGREGATION", "STAMP_DUTY_TAX", "TDS_TAX", "STT_TAX",
	// "MISC", "REVERSAL", "UNKNOWN".
	Type string `json:"type"`
	// Number of units involved in transaction
	Units float64 `json:"units"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Amount         respjson.Field
		Balance        respjson.Field
		Date           respjson.Field
		Description    respjson.Field
		DividendRate   respjson.Field
		Nav            respjson.Field
		Type           respjson.Field
		Units          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsEquityTransaction) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseDematAccountHoldingsEquityTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional transaction-specific fields that vary by source
type UnifiedResponseDematAccountHoldingsEquityTransactionAdditionalInfo struct {
	// Capital withdrawal amount (CDSL MF transactions)
	CapitalWithdrawal float64 `json:"capital_withdrawal"`
	// Units credited (demat transactions)
	Credit float64 `json:"credit"`
	// Units debited (demat transactions)
	Debit float64 `json:"debit"`
	// Income distribution amount (CDSL MF transactions)
	IncomeDistribution float64 `json:"income_distribution"`
	// Order/transaction reference number (demat transactions)
	OrderNo string `json:"order_no"`
	// Price per unit (NSDL/CDSL MF transactions)
	Price float64 `json:"price"`
	// Stamp duty charged
	StampDuty float64 `json:"stamp_duty"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapitalWithdrawal  respjson.Field
		Credit             respjson.Field
		Debit              respjson.Field
		IncomeDistribution respjson.Field
		OrderNo            respjson.Field
		Price              respjson.Field
		StampDuty          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsEquityTransactionAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *UnifiedResponseDematAccountHoldingsEquityTransactionAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseDematAccountHoldingsGovernmentSecurity struct {
	// Additional information specific to the government security
	AdditionalInfo UnifiedResponseDematAccountHoldingsGovernmentSecurityAdditionalInfo `json:"additional_info"`
	// ISIN code of the government security
	Isin string `json:"isin"`
	// Name of the government security
	Name string `json:"name"`
	// List of transactions for this holding (beta)
	Transactions []UnifiedResponseDematAccountHoldingsGovernmentSecurityTransaction `json:"transactions"`
	// Number of units held
	Units float64 `json:"units"`
	// Current market value of the holding
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Isin           respjson.Field
		Name           respjson.Field
		Transactions   respjson.Field
		Units          respjson.Field
		Value          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsGovernmentSecurity) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseDematAccountHoldingsGovernmentSecurity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information specific to the government security
type UnifiedResponseDematAccountHoldingsGovernmentSecurityAdditionalInfo struct {
	// Closing balance units for the statement period (beta)
	CloseUnits float64 `json:"close_units,nullable"`
	// Opening balance units for the statement period (beta)
	OpenUnits float64 `json:"open_units,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CloseUnits  respjson.Field
		OpenUnits   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsGovernmentSecurityAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *UnifiedResponseDematAccountHoldingsGovernmentSecurityAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unified transaction schema for all holding types (MF folios, equities, bonds,
// etc.)
type UnifiedResponseDematAccountHoldingsGovernmentSecurityTransaction struct {
	// Additional transaction-specific fields that vary by source
	AdditionalInfo UnifiedResponseDematAccountHoldingsGovernmentSecurityTransactionAdditionalInfo `json:"additional_info"`
	// Transaction amount in currency (computed from units × price/NAV)
	Amount float64 `json:"amount,nullable"`
	// Balance units after transaction
	Balance float64 `json:"balance"`
	// Transaction date (YYYY-MM-DD)
	Date time.Time `json:"date" format:"date"`
	// Transaction description/particulars
	Description string `json:"description"`
	// Dividend rate (for DIVIDEND_PAYOUT transactions)
	DividendRate float64 `json:"dividend_rate,nullable"`
	// NAV/price per unit on transaction date
	Nav float64 `json:"nav,nullable"`
	// Transaction type. Possible values are PURCHASE, PURCHASE_SIP, REDEMPTION,
	// SWITCH_IN, SWITCH_IN_MERGER, SWITCH_OUT, SWITCH_OUT_MERGER, DIVIDEND_PAYOUT,
	// DIVIDEND_REINVEST, SEGREGATION, STAMP_DUTY_TAX, TDS_TAX, STT_TAX, MISC,
	// REVERSAL, UNKNOWN.
	//
	// Any of "PURCHASE", "PURCHASE_SIP", "REDEMPTION", "SWITCH_IN",
	// "SWITCH_IN_MERGER", "SWITCH_OUT", "SWITCH_OUT_MERGER", "DIVIDEND_PAYOUT",
	// "DIVIDEND_REINVEST", "SEGREGATION", "STAMP_DUTY_TAX", "TDS_TAX", "STT_TAX",
	// "MISC", "REVERSAL", "UNKNOWN".
	Type string `json:"type"`
	// Number of units involved in transaction
	Units float64 `json:"units"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Amount         respjson.Field
		Balance        respjson.Field
		Date           respjson.Field
		Description    respjson.Field
		DividendRate   respjson.Field
		Nav            respjson.Field
		Type           respjson.Field
		Units          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsGovernmentSecurityTransaction) RawJSON() string {
	return r.JSON.raw
}
func (r *UnifiedResponseDematAccountHoldingsGovernmentSecurityTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional transaction-specific fields that vary by source
type UnifiedResponseDematAccountHoldingsGovernmentSecurityTransactionAdditionalInfo struct {
	// Capital withdrawal amount (CDSL MF transactions)
	CapitalWithdrawal float64 `json:"capital_withdrawal"`
	// Units credited (demat transactions)
	Credit float64 `json:"credit"`
	// Units debited (demat transactions)
	Debit float64 `json:"debit"`
	// Income distribution amount (CDSL MF transactions)
	IncomeDistribution float64 `json:"income_distribution"`
	// Order/transaction reference number (demat transactions)
	OrderNo string `json:"order_no"`
	// Price per unit (NSDL/CDSL MF transactions)
	Price float64 `json:"price"`
	// Stamp duty charged
	StampDuty float64 `json:"stamp_duty"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapitalWithdrawal  respjson.Field
		Credit             respjson.Field
		Debit              respjson.Field
		IncomeDistribution respjson.Field
		OrderNo            respjson.Field
		Price              respjson.Field
		StampDuty          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountHoldingsGovernmentSecurityTransactionAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *UnifiedResponseDematAccountHoldingsGovernmentSecurityTransactionAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseDematAccountLinkedHolder struct {
	// Name of the account holder
	Name string `json:"name"`
	// PAN of the account holder
	Pan string `json:"pan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Pan         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseDematAccountLinkedHolder) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseDematAccountLinkedHolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseInsurance struct {
	LifeInsurancePolicies []UnifiedResponseInsuranceLifeInsurancePolicy `json:"life_insurance_policies"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LifeInsurancePolicies respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseInsurance) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseInsurance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseInsuranceLifeInsurancePolicy struct {
	// Additional information specific to the policy
	AdditionalInfo any `json:"additional_info"`
	// Name of the life assured
	LifeAssured string `json:"life_assured"`
	// Name of the insurance policy
	PolicyName string `json:"policy_name"`
	// Insurance policy number
	PolicyNumber string `json:"policy_number"`
	// Premium amount
	PremiumAmount float64 `json:"premium_amount"`
	// Frequency of premium payment (e.g., Annual, Monthly)
	PremiumFrequency string `json:"premium_frequency"`
	// Insurance company name
	Provider string `json:"provider"`
	// Status of the policy (e.g., Active, Lapsed)
	Status string `json:"status"`
	// Sum assured amount
	SumAssured float64 `json:"sum_assured"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo   respjson.Field
		LifeAssured      respjson.Field
		PolicyName       respjson.Field
		PolicyNumber     respjson.Field
		PremiumAmount    respjson.Field
		PremiumFrequency respjson.Field
		Provider         respjson.Field
		Status           respjson.Field
		SumAssured       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseInsuranceLifeInsurancePolicy) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseInsuranceLifeInsurancePolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseInvestor struct {
	// Address of the investor
	Address string `json:"address"`
	// CAS ID of the investor (only for NSDL and CDSL)
	CasID string `json:"cas_id"`
	// Email address of the investor
	Email string `json:"email" format:"email"`
	// Mobile number of the investor
	Mobile string `json:"mobile"`
	// Name of the investor
	Name string `json:"name"`
	// PAN (Permanent Account Number) of the investor
	Pan string `json:"pan"`
	// Postal code of the investor's address
	Pincode string `json:"pincode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		CasID       respjson.Field
		Email       respjson.Field
		Mobile      respjson.Field
		Name        respjson.Field
		Pan         respjson.Field
		Pincode     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseInvestor) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseInvestor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseMeta struct {
	// Type of CAS detected and processed
	//
	// Any of "NSDL", "CDSL", "CAMS_KFINTECH".
	CasType string `json:"cas_type"`
	// Timestamp when the response was generated
	GeneratedAt     time.Time                          `json:"generated_at" format:"date-time"`
	StatementPeriod UnifiedResponseMetaStatementPeriod `json:"statement_period"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CasType         respjson.Field
		GeneratedAt     respjson.Field
		StatementPeriod respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseMetaStatementPeriod struct {
	// Start date of the statement period
	From time.Time `json:"from" format:"date"`
	// End date of the statement period
	To time.Time `json:"to" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseMetaStatementPeriod) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseMetaStatementPeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseMutualFund struct {
	// Additional folio information
	AdditionalInfo UnifiedResponseMutualFundAdditionalInfo `json:"additional_info"`
	// Asset Management Company name
	Amc string `json:"amc"`
	// Folio number
	FolioNumber string `json:"folio_number"`
	// List of account holders linked to this mutual fund folio
	LinkedHolders []UnifiedResponseMutualFundLinkedHolder `json:"linked_holders"`
	// Registrar and Transfer Agent name
	Registrar string                            `json:"registrar"`
	Schemes   []UnifiedResponseMutualFundScheme `json:"schemes"`
	// Total value of the folio
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Amc            respjson.Field
		FolioNumber    respjson.Field
		LinkedHolders  respjson.Field
		Registrar      respjson.Field
		Schemes        respjson.Field
		Value          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseMutualFund) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseMutualFund) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional folio information
type UnifiedResponseMutualFundAdditionalInfo struct {
	// KYC status of the folio
	KYC string `json:"kyc"`
	// PAN associated with the folio
	Pan string `json:"pan"`
	// PAN KYC status
	Pankyc string `json:"pankyc"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KYC         respjson.Field
		Pan         respjson.Field
		Pankyc      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseMutualFundAdditionalInfo) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseMutualFundAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseMutualFundLinkedHolder struct {
	// Name of the account holder
	Name string `json:"name"`
	// PAN of the account holder
	Pan string `json:"pan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Pan         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseMutualFundLinkedHolder) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseMutualFundLinkedHolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseMutualFundScheme struct {
	// Additional information specific to the scheme
	AdditionalInfo UnifiedResponseMutualFundSchemeAdditionalInfo `json:"additional_info"`
	// Cost of investment
	Cost float64                             `json:"cost"`
	Gain UnifiedResponseMutualFundSchemeGain `json:"gain"`
	// ISIN code of the scheme
	Isin string `json:"isin"`
	// Scheme name
	Name string `json:"name"`
	// Net Asset Value per unit
	Nav float64 `json:"nav"`
	// List of nominees
	Nominees     []string                                     `json:"nominees"`
	Transactions []UnifiedResponseMutualFundSchemeTransaction `json:"transactions"`
	// Type of mutual fund scheme
	//
	// Any of "Equity", "Debt", "Hybrid", "Other".
	Type string `json:"type"`
	// Number of units held
	Units float64 `json:"units"`
	// Current market value of the holding
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Cost           respjson.Field
		Gain           respjson.Field
		Isin           respjson.Field
		Name           respjson.Field
		Nav            respjson.Field
		Nominees       respjson.Field
		Transactions   respjson.Field
		Type           respjson.Field
		Units          respjson.Field
		Value          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseMutualFundScheme) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseMutualFundScheme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information specific to the scheme
type UnifiedResponseMutualFundSchemeAdditionalInfo struct {
	// Financial advisor name (CAMS/KFintech)
	Advisor string `json:"advisor"`
	// AMFI code for the scheme (CAMS/KFintech)
	Amfi string `json:"amfi"`
	// Closing balance units for the statement period
	CloseUnits float64 `json:"close_units,nullable"`
	// Opening balance units for the statement period
	OpenUnits float64 `json:"open_units,nullable"`
	// RTA code for the scheme (CAMS/KFintech)
	RtaCode string `json:"rta_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Advisor     respjson.Field
		Amfi        respjson.Field
		CloseUnits  respjson.Field
		OpenUnits   respjson.Field
		RtaCode     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseMutualFundSchemeAdditionalInfo) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseMutualFundSchemeAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseMutualFundSchemeGain struct {
	// Absolute gain or loss
	Absolute float64 `json:"absolute"`
	// Percentage gain or loss
	Percentage float64 `json:"percentage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Absolute    respjson.Field
		Percentage  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseMutualFundSchemeGain) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseMutualFundSchemeGain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unified transaction schema for all holding types (MF folios, equities, bonds,
// etc.)
type UnifiedResponseMutualFundSchemeTransaction struct {
	// Additional transaction-specific fields that vary by source
	AdditionalInfo UnifiedResponseMutualFundSchemeTransactionAdditionalInfo `json:"additional_info"`
	// Transaction amount in currency (computed from units × price/NAV)
	Amount float64 `json:"amount,nullable"`
	// Balance units after transaction
	Balance float64 `json:"balance"`
	// Transaction date (YYYY-MM-DD)
	Date time.Time `json:"date" format:"date"`
	// Transaction description/particulars
	Description string `json:"description"`
	// Dividend rate (for DIVIDEND_PAYOUT transactions)
	DividendRate float64 `json:"dividend_rate,nullable"`
	// NAV/price per unit on transaction date
	Nav float64 `json:"nav,nullable"`
	// Transaction type. Possible values are PURCHASE, PURCHASE_SIP, REDEMPTION,
	// SWITCH_IN, SWITCH_IN_MERGER, SWITCH_OUT, SWITCH_OUT_MERGER, DIVIDEND_PAYOUT,
	// DIVIDEND_REINVEST, SEGREGATION, STAMP_DUTY_TAX, TDS_TAX, STT_TAX, MISC,
	// REVERSAL, UNKNOWN.
	//
	// Any of "PURCHASE", "PURCHASE_SIP", "REDEMPTION", "SWITCH_IN",
	// "SWITCH_IN_MERGER", "SWITCH_OUT", "SWITCH_OUT_MERGER", "DIVIDEND_PAYOUT",
	// "DIVIDEND_REINVEST", "SEGREGATION", "STAMP_DUTY_TAX", "TDS_TAX", "STT_TAX",
	// "MISC", "REVERSAL", "UNKNOWN".
	Type string `json:"type"`
	// Number of units involved in transaction
	Units float64 `json:"units"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Amount         respjson.Field
		Balance        respjson.Field
		Date           respjson.Field
		Description    respjson.Field
		DividendRate   respjson.Field
		Nav            respjson.Field
		Type           respjson.Field
		Units          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseMutualFundSchemeTransaction) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseMutualFundSchemeTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional transaction-specific fields that vary by source
type UnifiedResponseMutualFundSchemeTransactionAdditionalInfo struct {
	// Capital withdrawal amount (CDSL MF transactions)
	CapitalWithdrawal float64 `json:"capital_withdrawal"`
	// Units credited (demat transactions)
	Credit float64 `json:"credit"`
	// Units debited (demat transactions)
	Debit float64 `json:"debit"`
	// Income distribution amount (CDSL MF transactions)
	IncomeDistribution float64 `json:"income_distribution"`
	// Order/transaction reference number (demat transactions)
	OrderNo string `json:"order_no"`
	// Price per unit (NSDL/CDSL MF transactions)
	Price float64 `json:"price"`
	// Stamp duty charged
	StampDuty float64 `json:"stamp_duty"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapitalWithdrawal  respjson.Field
		Credit             respjson.Field
		Debit              respjson.Field
		IncomeDistribution respjson.Field
		OrderNo            respjson.Field
		Price              respjson.Field
		StampDuty          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseMutualFundSchemeTransactionAdditionalInfo) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseMutualFundSchemeTransactionAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseNp struct {
	// Additional information specific to the NPS account
	AdditionalInfo any `json:"additional_info"`
	// Central Record Keeping Agency name
	Cra   string                  `json:"cra"`
	Funds []UnifiedResponseNpFund `json:"funds"`
	// List of account holders linked to this NPS account
	LinkedHolders []UnifiedResponseNpLinkedHolder `json:"linked_holders"`
	// Permanent Retirement Account Number (PRAN)
	Pran string `json:"pran"`
	// Total value of the NPS account
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Cra            respjson.Field
		Funds          respjson.Field
		LinkedHolders  respjson.Field
		Pran           respjson.Field
		Value          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseNp) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseNp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseNpFund struct {
	// Additional information specific to the NPS fund
	AdditionalInfo UnifiedResponseNpFundAdditionalInfo `json:"additional_info"`
	// Cost of investment
	Cost float64 `json:"cost"`
	// Name of the NPS fund
	Name string `json:"name"`
	// Net Asset Value per unit
	Nav float64 `json:"nav"`
	// Number of units held
	Units float64 `json:"units"`
	// Current market value of the holding
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfo respjson.Field
		Cost           respjson.Field
		Name           respjson.Field
		Nav            respjson.Field
		Units          respjson.Field
		Value          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseNpFund) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseNpFund) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information specific to the NPS fund
type UnifiedResponseNpFundAdditionalInfo struct {
	// Fund manager name
	Manager string `json:"manager"`
	// NPS tier (Tier I or Tier II)
	//
	// Any of 1, 2.
	Tier float64 `json:"tier,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Manager     respjson.Field
		Tier        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseNpFundAdditionalInfo) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseNpFundAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseNpLinkedHolder struct {
	// Name of the account holder
	Name string `json:"name"`
	// PAN of the account holder
	Pan string `json:"pan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Pan         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseNpLinkedHolder) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseNpLinkedHolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseSummary struct {
	Accounts UnifiedResponseSummaryAccounts `json:"accounts"`
	// Total portfolio value across all accounts
	TotalValue float64 `json:"total_value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		TotalValue  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseSummaryAccounts struct {
	Demat       UnifiedResponseSummaryAccountsDemat       `json:"demat"`
	Insurance   UnifiedResponseSummaryAccountsInsurance   `json:"insurance"`
	MutualFunds UnifiedResponseSummaryAccountsMutualFunds `json:"mutual_funds"`
	Nps         UnifiedResponseSummaryAccountsNps         `json:"nps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Demat       respjson.Field
		Insurance   respjson.Field
		MutualFunds respjson.Field
		Nps         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseSummaryAccounts) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseSummaryAccounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseSummaryAccountsDemat struct {
	// Number of demat accounts
	Count int64 `json:"count"`
	// Total value of demat accounts
	TotalValue float64 `json:"total_value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		TotalValue  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseSummaryAccountsDemat) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseSummaryAccountsDemat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseSummaryAccountsInsurance struct {
	// Number of insurance policies
	Count int64 `json:"count"`
	// Total value of insurance policies
	TotalValue float64 `json:"total_value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		TotalValue  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseSummaryAccountsInsurance) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseSummaryAccountsInsurance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseSummaryAccountsMutualFunds struct {
	// Number of mutual fund folios
	Count int64 `json:"count"`
	// Total value of mutual funds
	TotalValue float64 `json:"total_value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		TotalValue  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseSummaryAccountsMutualFunds) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseSummaryAccountsMutualFunds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnifiedResponseSummaryAccountsNps struct {
	// Number of NPS accounts
	Count int64 `json:"count"`
	// Total value of NPS accounts
	TotalValue float64 `json:"total_value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		TotalValue  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnifiedResponseSummaryAccountsNps) RawJSON() string { return r.JSON.raw }
func (r *UnifiedResponseSummaryAccountsNps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CasParserCamsKfintechParams struct {
	// Password for the PDF file (if required)
	Password param.Opt[string] `json:"password,omitzero"`
	// Base64 encoded CAS PDF file
	PdfFile param.Opt[string] `json:"pdf_file,omitzero" format:"base64"`
	// URL to the CAS PDF file
	PdfURL param.Opt[string] `json:"pdf_url,omitzero" format:"uri"`
	paramObj
}

func (r CasParserCamsKfintechParams) MarshalJSON() (data []byte, err error) {
	type shadow CasParserCamsKfintechParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CasParserCamsKfintechParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CasParserCdslParams struct {
	// Password for the PDF file (if required)
	Password param.Opt[string] `json:"password,omitzero"`
	// Base64 encoded CAS PDF file
	PdfFile param.Opt[string] `json:"pdf_file,omitzero" format:"base64"`
	// URL to the CAS PDF file
	PdfURL param.Opt[string] `json:"pdf_url,omitzero" format:"uri"`
	paramObj
}

func (r CasParserCdslParams) MarshalJSON() (data []byte, err error) {
	type shadow CasParserCdslParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CasParserCdslParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CasParserNsdlParams struct {
	// Password for the PDF file (if required)
	Password param.Opt[string] `json:"password,omitzero"`
	// Base64 encoded CAS PDF file
	PdfFile param.Opt[string] `json:"pdf_file,omitzero" format:"base64"`
	// URL to the CAS PDF file
	PdfURL param.Opt[string] `json:"pdf_url,omitzero" format:"uri"`
	paramObj
}

func (r CasParserNsdlParams) MarshalJSON() (data []byte, err error) {
	type shadow CasParserNsdlParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CasParserNsdlParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CasParserSmartParseParams struct {
	// Password for the PDF file (if required)
	Password param.Opt[string] `json:"password,omitzero"`
	// Base64 encoded CAS PDF file
	PdfFile param.Opt[string] `json:"pdf_file,omitzero" format:"base64"`
	// URL to the CAS PDF file
	PdfURL param.Opt[string] `json:"pdf_url,omitzero" format:"uri"`
	paramObj
}

func (r CasParserSmartParseParams) MarshalJSON() (data []byte, err error) {
	type shadow CasParserSmartParseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CasParserSmartParseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
