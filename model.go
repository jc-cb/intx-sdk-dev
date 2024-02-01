package intx

type ErrorMessage struct {
	Value string `json:"message"`
}

type Portfolio struct {
	PortfolioId    string `json:"portfolio_id"`
	PortfolioUuid  string `json:"portfolio_uuid"`
	Name           string `json:"name"`
	UserUuid       string `json:"user_uuid"`
	MakerFeeRate   string `json:"maker_fee_rate"`
	TakerFeeRate   string `json:"taker_fee_rate"`
	TradingLock    bool   `json:"trading_lock"`
	BorrowDisabled bool   `json:"borrow_disabled"`
	IsLsp          bool   `json:"is_lsp"`
	IsDefault      bool   `json:"is_default"`
}

type Asset struct {
	AssetId                  string  `json:"asset_id"`
	AssetUuid                string  `json:"asset_uuid"`
	AssetName                string  `json:"asset_name"`
	Status                   string  `json:"status"`
	CollateralWeight         float64 `json:"collateral_weight"`
	SupportedNetworksEnabled bool    `json:"supported_networks_enabled"`
}

type Network struct {
	AssetId          string `json:"asset_id"`
	AssetUuid        string `json:"asset_uuid"`
	AssetName        string `json:"asset_name"`
	IsDefault        bool   `json:"is_default"`
	NetworkName      string `json:"network_name"`
	DisplayName      string `json:"display_name"`
	NetworkArnId     string `json:"network_arn_id"`
	MinWithdrawalAmt string `json:"min_withdrawal_amt"`
	MaxWithdrawalAmt string `json:"max_withdrawal_amt"`
	NetworkConfirms  int    `json:"network_confirms"`
	ProcessingTime   int    `json:"processing_time"`
}

type Instrument struct {
	InstrumentId       string  `json:"instrument_id"`
	InstrumentUuid     string  `json:"instrument_uuid"`
	Symbol             string  `json:"symbol"`
	Type               string  `json:"type"`
	BaseAssetId        string  `json:"base_asset_id"`
	BaseAssetUuid      string  `json:"base_asset_uuid"`
	BaseAssetName      string  `json:"base_asset_name"`
	QuoteAssetId       string  `json:"quote_asset_id"`
	QuoteAssetUuid     string  `json:"quote_asset_uuid"`
	QuoteAssetName     string  `json:"quote_asset_name"`
	BaseIncrement      string  `json:"base_increment"`
	QuoteIncrement     string  `json:"quote_increment"`
	MarketOrderPercent float64 `json:"market_order_percent"`
	PriceBandPercent   float64 `json:"price_band_percent"`
	Qty24hr            string  `json:"qty_24hr"`
	Notional24hr       string  `json:"notional_24hr"`
	AvgDailyQty        string  `json:"avg_daily_qty"`
	AvgDailyNotional   string  `json:"avg_daily_notional"`
	PreviousDayQty     string  `json:"previous_day_qty"`
	PositionLimitQty   string  `json:"position_limit_qty"`
	PositionLimitAdv   float64 `json:"position_limit_adv"`
	InitialMarginAdv   float64 `json:"initial_margin_adv"`
	ReplacementCost    string  `json:"replacement_cost"`
	BaseImf            float64 `json:"base_imf"`
	MinNotionalValue   string  `json:"min_notional_value"`
	FundingInterval    string  `json:"funding_interval"`
	TradingState       string  `json:"trading_state"`
	OpenInterest       string  `json:"open_interest"`
	Quote              Quote   `json:"quote"`
}

type Quote struct {
	BestBidPrice     string `json:"best_bid_price"`
	BestBidSize      string `json:"best_bid_size"`
	BestAskPrice     string `json:"best_ask_price"`
	BestAskSize      string `json:"best_ask_size"`
	TradePrice       string `json:"trade_price"`
	TradeQty         string `json:"trade_qty"`
	IndexPrice       string `json:"index_price"`
	MarkPrice        string `json:"mark_price"`
	SettlementPrice  string `json:"settlement_price"`
	LimitUp          string `json:"limit_up"`
	LimitDown        string `json:"limit_down"`
	PredictedFunding string `json:"predicted_funding"`
	Timestamp        string `json:"timestamp"`
}

type HistoricalFundingRate struct {
	InstrumentId string  `json:"instrument_id"`
	FundingRate  float64 `json:"funding_rate"`
	MarkPrice    float64 `json:"mark_price"`
	EventTime    string  `json:"event_time"`
}

type Summary struct {
	Collateral                         string  `json:"collateral"`
	UnrealizedPnl                      string  `json:"unrealized_pnl"`
	PositionNotional                   string  `json:"position_notional"`
	OpenPositionNotional               string  `json:"open_position_notional"`
	PendingFees                        string  `json:"pending_fees"`
	Borrow                             string  `json:"borrow"`
	AccruedInterest                    string  `json:"accrued_interest"`
	RollingDebt                        string  `json:"rolling_debt"`
	Balance                            string  `json:"balance"`
	BuyingPower                        string  `json:"buying_power"`
	PortfolioCurrentMargin             float64 `json:"portfolio_current_margin"`
	PortfolioInitialMargin             float64 `json:"portfolio_initial_margin"`
	PortfolioMaintenanceMargin         float64 `json:"portfolio_maintenance_margin"`
	PortfolioCloseOutMargin            float64 `json:"portfolio_close_out_margin"`
	InLiquidation                      bool    `json:"in_liquidation"`
	PortfolioCurrentMarginNotional     float64 `json:"portfolio_current_margin_notional"`
	PortfolioInitialMarginNotional     float64 `json:"portfolio_initial_margin_notional"`
	PortfolioMaintenanceMarginNotional float64 `json:"portfolio_maintenance_margin_notional"`
	PortfolioCloseOutMarginNotional    float64 `json:"portfolio_close_out_margin_notional"`
	MarginOverride                     float64 `json:"margin_override"`
	LockupInitialMargin                float64 `json:"lockup_initial_margin"`
}

type Balance struct {
	AssetId           string `json:"asset_id"`
	AssetUuid         string `json:"asset_uuid"`
	AssetName         string `json:"asset_name"`
	Quantity          string `json:"quantity"`
	Hold              string `json:"hold"`
	TransferHold      string `json:"transfer_hold"`
	CollateralValue   string `json:"collateral_value"`
	MaxWithdrawAmount string `json:"max_withdraw_amount"`
}

type Position struct {
	InstrumentId   string `json:"instrument_id"`
	InstrumentUuid string `json:"instrument_uuid"`
	Symbol         string `json:"symbol"`
	Vwap           string `json:"vwap"`
	NetSize        string `json:"net_size"`
	BuyOrderSize   string `json:"buy_order_size"`
	SellOrderSize  string `json:"sell_order_size"`
	ImContribution string `json:"im_contribution"`
	UnrealizedPnl  string `json:"unrealized_pnl"`
	MarkPrice      string `json:"mark_price"`
}

type Details struct {
	Summary   Summary    `json:"summary"`
	Balances  []Balance  `json:"balances"`
	Positions []Position `json:"positions"`
}

type CreateOrderRequest struct {
	ClientOrderId string  `json:"client_order_id"`
	Side          string  `json:"side"`
	Size          string  `json:"size"`
	Tif           string  `json:"tif"`
	Instrument    string  `json:"instrument"`
	Type          string  `json:"type"`
	Price         string  `json:"price,omitempty"`
	StopPrice     *string `json:"stop_price,omitempty"`
	ExpireTime    *string `json:"expire_time,omitempty"`
	Portfolio     string  `json:"portfolio"`
	User          *string `json:"user,omitempty"`
	StpMode       *string `json:"stp_mode,omitempty"`
	PostOnly      *bool   `json:"post_only,omitempty"`
}

type Order struct {
	OrderId        string `json:"order_id"`
	ClientOrderId  string `json:"client_order_id"`
	Side           string `json:"side"`
	InstrumentId   string `json:"instrument_id"`
	InstrumentUuid string `json:"instrument_uuid"`
	Symbol         string `json:"symbol"`
	PortfolioId    string `json:"portfolio_id"`
	PortfolioUuid  string `json:"portfolio_uuid"`
	Type           string `json:"type"`
	Price          string `json:"price"`
	StopPrice      string `json:"stop_price"`
	Size           string `json:"size"`
	Tif            string `json:"tif"`
	ExpireTime     string `json:"expire_time"`
	StpMode        string `json:"stp_mode"`
	EventType      string `json:"event_type"`
	OrderStatus    string `json:"order_status"`
	LeavesQty      string `json:"leaves_qty"`
	ExecQty        string `json:"exec_qty"`
	AvgPrice       string `json:"avg_price"`
	Message        string `json:"message"`
	Fee            string `json:"fee"`
}

type MarginOverride struct {
	PortfolioId    string `json:"portfolio_id"`
	MarginOverride string `json:"margin_override"`
}

type PaginationParams struct {
	RefDatetime  string `json:"ref_datetime"`
	ResultLimit  int    `json:"result_limit"`
	ResultOffset int    `json:"result_offset"`
}