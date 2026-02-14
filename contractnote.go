// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/cas-parser-go/internal/apijson"
	"github.com/stainless-sdks/cas-parser-go/internal/requestconfig"
	"github.com/stainless-sdks/cas-parser-go/option"
	"github.com/stainless-sdks/cas-parser-go/packages/param"
	"github.com/stainless-sdks/cas-parser-go/packages/respjson"
)

// ContractNoteService contains methods and other services that help with
// interacting with the cas-parser API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContractNoteService] method instead.
type ContractNoteService struct {
	Options []option.RequestOption
}

// NewContractNoteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContractNoteService(opts ...option.RequestOption) (r ContractNoteService) {
	r = ContractNoteService{}
	r.Options = opts
	return
}

// This endpoint parses Contract Note PDF files from various brokers including
// Zerodha, Groww, Upstox, ICICI Securities, and others.
//
// **What is a Contract Note?** A contract note is a legal document that provides
// details of all trades executed by an investor. It includes:
//
// - Trade details with timestamps, quantities, and prices
// - Brokerage and charges breakdown
// - Settlement information
// - Regulatory compliance details
//
// **Supported Brokers:**
//
// - Zerodha Broking Limited
// - Groww Invest Tech Private Limited
// - Upstox (RKSV Securities)
// - ICICI Securities Limited
// - Auto-detection for unknown brokers
//
// **Key Features:**
//
//   - **Auto-detection**: Automatically identifies broker type from PDF content
//   - **Comprehensive parsing**: Extracts equity transactions, derivatives
//     transactions, detailed trades, and charges
//   - **Flexible input**: Accepts both file upload and URL-based PDF input
//   - **Password protection**: Supports password-protected PDFs
//
// The API returns structured data including contract note information, client
// details, transaction summaries, and detailed trade-by-trade breakdowns.
func (r *ContractNoteService) Parse(ctx context.Context, body ContractNoteParseParams, opts ...option.RequestOption) (res *ContractNoteParseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/contract_note/parse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContractNoteParseResponse struct {
	Data   ContractNoteParseResponseData `json:"data"`
	Msg    string                        `json:"msg"`
	Status string                        `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Msg         respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractNoteParseResponse) RawJSON() string { return r.JSON.raw }
func (r *ContractNoteParseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractNoteParseResponseData struct {
	BrokerInfo ContractNoteParseResponseDataBrokerInfo `json:"broker_info"`
	// Breakdown of various charges and fees
	ChargesSummary   ContractNoteParseResponseDataChargesSummary   `json:"charges_summary"`
	ClientInfo       ContractNoteParseResponseDataClientInfo       `json:"client_info"`
	ContractNoteInfo ContractNoteParseResponseDataContractNoteInfo `json:"contract_note_info"`
	// Summary of derivatives transactions
	DerivativesTransactions []ContractNoteParseResponseDataDerivativesTransaction `json:"derivatives_transactions"`
	// Detailed breakdown of all individual trades
	DetailedTrades []ContractNoteParseResponseDataDetailedTrade `json:"detailed_trades"`
	// Summary of equity transactions grouped by security
	EquityTransactions []ContractNoteParseResponseDataEquityTransaction `json:"equity_transactions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrokerInfo              respjson.Field
		ChargesSummary          respjson.Field
		ClientInfo              respjson.Field
		ContractNoteInfo        respjson.Field
		DerivativesTransactions respjson.Field
		DetailedTrades          respjson.Field
		EquityTransactions      respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractNoteParseResponseData) RawJSON() string { return r.JSON.raw }
func (r *ContractNoteParseResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractNoteParseResponseDataBrokerInfo struct {
	// Auto-detected or specified broker type
	//
	// Any of "zerodha", "groww", "upstox", "icici", "unknown".
	BrokerType string `json:"broker_type"`
	// Broker company name
	Name string `json:"name"`
	// SEBI registration number of the broker
	SebiRegistration string `json:"sebi_registration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrokerType       respjson.Field
		Name             respjson.Field
		SebiRegistration respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractNoteParseResponseDataBrokerInfo) RawJSON() string { return r.JSON.raw }
func (r *ContractNoteParseResponseDataBrokerInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of various charges and fees
type ContractNoteParseResponseDataChargesSummary struct {
	// Central GST amount
	Cgst float64 `json:"cgst"`
	// Exchange transaction charges
	ExchangeTransactionCharges float64 `json:"exchange_transaction_charges"`
	// Integrated GST amount
	Igst float64 `json:"igst"`
	// Final net amount receivable or payable
	NetAmountReceivablePayable float64 `json:"net_amount_receivable_payable"`
	// Net pay-in/pay-out obligation
	PayInPayOutObligation float64 `json:"pay_in_pay_out_obligation"`
	// SEBI turnover fees
	SebiTurnoverFees float64 `json:"sebi_turnover_fees"`
	// Securities Transaction Tax
	SecuritiesTransactionTax float64 `json:"securities_transaction_tax"`
	// State GST amount
	Sgst float64 `json:"sgst"`
	// Stamp duty charges
	StampDuty float64 `json:"stamp_duty"`
	// Taxable brokerage amount
	TaxableValueBrokerage float64 `json:"taxable_value_brokerage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cgst                       respjson.Field
		ExchangeTransactionCharges respjson.Field
		Igst                       respjson.Field
		NetAmountReceivablePayable respjson.Field
		PayInPayOutObligation      respjson.Field
		SebiTurnoverFees           respjson.Field
		SecuritiesTransactionTax   respjson.Field
		Sgst                       respjson.Field
		StampDuty                  respjson.Field
		TaxableValueBrokerage      respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractNoteParseResponseDataChargesSummary) RawJSON() string { return r.JSON.raw }
func (r *ContractNoteParseResponseDataChargesSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractNoteParseResponseDataClientInfo struct {
	// Client address
	Address string `json:"address"`
	// GST state code
	GstStateCode string `json:"gst_state_code"`
	// Client name
	Name string `json:"name"`
	// Client PAN number
	Pan string `json:"pan"`
	// GST place of supply
	PlaceOfSupply string `json:"place_of_supply"`
	// Unique Client Code
	Ucc string `json:"ucc"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address       respjson.Field
		GstStateCode  respjson.Field
		Name          respjson.Field
		Pan           respjson.Field
		PlaceOfSupply respjson.Field
		Ucc           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractNoteParseResponseDataClientInfo) RawJSON() string { return r.JSON.raw }
func (r *ContractNoteParseResponseDataClientInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractNoteParseResponseDataContractNoteInfo struct {
	// Contract note reference number
	ContractNoteNumber string `json:"contract_note_number"`
	// Settlement date for the trades
	SettlementDate time.Time `json:"settlement_date" format:"date"`
	// Settlement reference number
	SettlementNumber string `json:"settlement_number"`
	// Date when trades were executed
	TradeDate time.Time `json:"trade_date" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContractNoteNumber respjson.Field
		SettlementDate     respjson.Field
		SettlementNumber   respjson.Field
		TradeDate          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractNoteParseResponseDataContractNoteInfo) RawJSON() string { return r.JSON.raw }
func (r *ContractNoteParseResponseDataContractNoteInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractNoteParseResponseDataDerivativesTransaction struct {
	// Brokerage charged per unit
	BrokeragePerUnit float64 `json:"brokerage_per_unit"`
	// Transaction type (Buy/Sell/Bring Forward/Carry Forward)
	BuySellBfCf string `json:"buy_sell_bf_cf"`
	// Closing rate per unit
	ClosingRatePerUnit float64 `json:"closing_rate_per_unit"`
	// Derivatives contract description
	ContractDescription string `json:"contract_description"`
	// Net total amount
	NetTotal float64 `json:"net_total"`
	// Quantity traded
	Quantity float64 `json:"quantity"`
	// Weighted Average Price per unit
	WapPerUnit float64 `json:"wap_per_unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrokeragePerUnit    respjson.Field
		BuySellBfCf         respjson.Field
		ClosingRatePerUnit  respjson.Field
		ContractDescription respjson.Field
		NetTotal            respjson.Field
		Quantity            respjson.Field
		WapPerUnit          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractNoteParseResponseDataDerivativesTransaction) RawJSON() string { return r.JSON.raw }
func (r *ContractNoteParseResponseDataDerivativesTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractNoteParseResponseDataDetailedTrade struct {
	// Brokerage charged for this trade
	Brokerage float64 `json:"brokerage"`
	// Transaction type (B for Buy, S for Sell)
	BuySell string `json:"buy_sell"`
	// Closing rate per unit
	ClosingRatePerUnit float64 `json:"closing_rate_per_unit"`
	// Exchange name
	Exchange string `json:"exchange"`
	// Net rate per unit
	NetRatePerUnit float64 `json:"net_rate_per_unit"`
	// Net total for this trade
	NetTotal float64 `json:"net_total"`
	// Order reference number
	OrderNumber string `json:"order_number"`
	// Time when order was placed
	OrderTime string `json:"order_time"`
	// Quantity traded
	Quantity float64 `json:"quantity"`
	// Additional remarks or notes
	Remarks string `json:"remarks"`
	// Security name with exchange and ISIN
	SecurityDescription string `json:"security_description"`
	// Trade reference number
	TradeNumber string `json:"trade_number"`
	// Time when trade was executed
	TradeTime string `json:"trade_time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brokerage           respjson.Field
		BuySell             respjson.Field
		ClosingRatePerUnit  respjson.Field
		Exchange            respjson.Field
		NetRatePerUnit      respjson.Field
		NetTotal            respjson.Field
		OrderNumber         respjson.Field
		OrderTime           respjson.Field
		Quantity            respjson.Field
		Remarks             respjson.Field
		SecurityDescription respjson.Field
		TradeNumber         respjson.Field
		TradeTime           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractNoteParseResponseDataDetailedTrade) RawJSON() string { return r.JSON.raw }
func (r *ContractNoteParseResponseDataDetailedTrade) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractNoteParseResponseDataEquityTransaction struct {
	// Total quantity purchased
	BuyQuantity float64 `json:"buy_quantity"`
	// Total value of buy transactions
	BuyTotalValue float64 `json:"buy_total_value"`
	// Weighted Average Price for buy transactions
	BuyWap float64 `json:"buy_wap"`
	// ISIN code of the security
	Isin string `json:"isin"`
	// Net amount payable/receivable for this security
	NetObligation float64 `json:"net_obligation"`
	// Name of the security
	SecurityName string `json:"security_name"`
	// Trading symbol
	SecuritySymbol string `json:"security_symbol"`
	// Total quantity sold
	SellQuantity float64 `json:"sell_quantity"`
	// Total value of sell transactions
	SellTotalValue float64 `json:"sell_total_value"`
	// Weighted Average Price for sell transactions
	SellWap float64 `json:"sell_wap"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BuyQuantity    respjson.Field
		BuyTotalValue  respjson.Field
		BuyWap         respjson.Field
		Isin           respjson.Field
		NetObligation  respjson.Field
		SecurityName   respjson.Field
		SecuritySymbol respjson.Field
		SellQuantity   respjson.Field
		SellTotalValue respjson.Field
		SellWap        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractNoteParseResponseDataEquityTransaction) RawJSON() string { return r.JSON.raw }
func (r *ContractNoteParseResponseDataEquityTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractNoteParseParams struct {
	// Password for the PDF file (usually PAN number for Zerodha)
	Password param.Opt[string] `json:"password,omitzero"`
	// Base64 encoded contract note PDF file
	PdfFile param.Opt[string] `json:"pdf_file,omitzero" format:"base64"`
	// URL to the contract note PDF file
	PdfURL param.Opt[string] `json:"pdf_url,omitzero" format:"uri"`
	// Optional broker type override. If not provided, system will auto-detect.
	//
	// Any of "zerodha", "groww", "upstox", "icici".
	BrokerType ContractNoteParseParamsBrokerType `json:"broker_type,omitzero"`
	paramObj
}

func (r ContractNoteParseParams) MarshalJSON() (data []byte, err error) {
	type shadow ContractNoteParseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContractNoteParseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional broker type override. If not provided, system will auto-detect.
type ContractNoteParseParamsBrokerType string

const (
	ContractNoteParseParamsBrokerTypeZerodha ContractNoteParseParamsBrokerType = "zerodha"
	ContractNoteParseParamsBrokerTypeGroww   ContractNoteParseParamsBrokerType = "groww"
	ContractNoteParseParamsBrokerTypeUpstox  ContractNoteParseParamsBrokerType = "upstox"
	ContractNoteParseParamsBrokerTypeIcici   ContractNoteParseParamsBrokerType = "icici"
)
