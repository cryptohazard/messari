package messari

import "time"

// Team informations
type Team struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Description string `json:"description"`
	Github      string `json:"github"`
	Linkedin    string `json:"linkedin"`
	Medium      string `json:"medium"`
	Twitter     string `json:"twitter"`
}

// Profile holds information for a crypto asset.
type Profile struct {
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Tagline           string `json:"tagline"`
	Overview          string `json:"overview"`
	Background        string `json:"background"`
	Technology        string `json:"technology"`
	TokenDistribution struct {
		Description         string `json:"description"`
		SaleStart           string `json:"sale_sart"`
		SaleEnd             string `json:"sale_end"`
		InitialDistribution int    `json:"initial_distribution"`
		CurrentSupply       int    `json:"current_supply"`
		MaxSupply           int    `json:"max_supply"`
	} `json:"token_distribution"`
	Organizations []struct {
		Name                string `json:"name"`
		SoundedDate         string `json:"founded_date"`
		Governance          string `json:"governance"`
		LegalStructure      string `json:"legal_structure"`
		Jurisdiction        string `json:"jurisdiction"`
		OrgCharter          string `json:"org_charter"`
		Descriptions        string `json:"description"`
		PeopleCountEstimate string `json:"people_count_estimate"`
	} `json:"organizations"`
	People struct {
		FoundingTeam []Team `json:"founding_team"`
		Contributors []Team `json:"contributors"`
		Investors    []Team `json:"investors"`
		Advisors     []Team `json:"advisors"`
	} `json:"people"`
	RelevantResources []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"relevant_resources"`
}

// Metrics Get quantitative metrics by asset symbol.
type Metrics struct {
	ID         string `json:"id"`
	Symbol     string `json:"symbol"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	MarketData struct {
		PriceUsd                               float64   `json:"price_usd"`
		PriceBtc                               float64   `json:"price_btc"`
		VolumeLast24Hours                      float64   `json:"volume_last_24_hours"`
		RealVolumeLast24Hours                  float64   `json:"real_volume_last_24_hours"`
		VolumeLast24HoursOverstatementMultiple float64   `json:"volume_last_24_hours_overstatement_multiple"`
		PercentChangeUsdLast24Hours            float64   `json:"percent_change_usd_last_24_hours"`
		PercentChangeBtcLast24Hours            float64   `json:"percent_change_btc_last_24_hours"`
		LastTradeAt                            time.Time `json:"last_trade_at"`
	} `json:"market_data"`
	Marketcap struct {
		CurrentMarketcapUsd              float64 `json:"current_marketcap_usd"`
		Y2050MarketcapUsd                float64 `json:"y_2050_marketcap_usd"`
		YPlus10MarketcapUsd              float64 `json:"y_plus10_marketcap_usd"`
		LiquidMarketcapUsd               float64 `json:"liquid_marketcap_usd"`
		VolumeTurnoverLast24HoursPercent float64 `json:"volume_turnover_last_24_hours_percent"`
	} `json:"marketcap"`
	Supply struct {
		Y2050                  float64 `json:"y_2050"`
		YPlus10                float64 `json:"y_plus10"`
		Liquid                 float64 `json:"liquid"`
		Circulating            float64 `json:"circulating"`
		Y2050IssuedPercent     float64 `json:"y_2050_issued_percent"`
		AnnualInflationPercent float64 `json:"annual_inflation_percent"`
		YPlus10IssuedPercent   float64 `json:"y_plus10_issued_percent"`
	} `json:"supply"`
	BlockchainStats24Hours struct {
		TransactionVolume         float64 `json:"transaction_volume"`
		AdjustedTransactionVolume float64 `json:"adjusted_transaction_volume"`
		Nvt                       float64 `json:"nvt"`
		AdjustedNvt               float64 `json:"adjusted_nvt"`
		SumOfFees                 float64 `json:"sum_of_fees"`
		MedianTxValue             float64 `json:"median_tx_value"`
		MedianTxFee               float64 `json:"median_tx_fee"`
		CountOfActiveAddresses    int     `json:"count_of_active_addresses"`
		CountOfTx                 int     `json:"count_of_tx"`
		CountOfPayments           int     `json:"count_of_payments"`
		NewIssuance               float64 `json:"new_issuance"`
		AverageDifficulty         float64 `json:"average_difficulty"`
		KilobytesAdded            float64 `json:"kilobytes_added"`
		CountOfBlocksAdded        int     `json:"count_of_blocks_added"`
	} `json:"blockchain_stats_24_hours"`
	AllTimeHigh struct {
		Price             float64 `json:"price"`
		At                string  `json:"at"`
		DaysSince         int     `json:"days_since"`
		PercentDown       float64 `json:"percent_down"`
		BreakevenMultiple float64 `json:"breakeven_multiple"`
	} `json:"all_time_high"`
	CycleLow struct {
		Price     float64 `json:"price"`
		At        string  `json:"at"`
		PercentUp float64 `json:"percent_up"`
		DaysSince int     `json:"days_since"`
	} `json:"cycle_low"`
	TokenSaleStats struct {
		SaleProceedsUsd        interface{} `json:"sale_proceeds_usd"`
		SaleStartDate          interface{} `json:"sale_start_date"`
		SaleEndDate            interface{} `json:"sale_end_date"`
		RoiSinceSaleUsdPercent interface{} `json:"roi_since_sale_usd_percent"`
		RoiSinceSaleBtcPercent interface{} `json:"roi_since_sale_btc_percent"`
		RoiSinceSaleEthPercent interface{} `json:"roi_since_sale_eth_percent"`
	} `json:"token_sale_stats"`
	StakingStats struct {
		StakingYieldPercent     interface{} `json:"staking_yield_percent"`
		StakingType             interface{} `json:"staking_type"`
		StakingMinimum          interface{} `json:"staking_minimum"`
		TokensStaked            interface{} `json:"tokens_staked"`
		TokensStakedPercent     float64     `json:"tokens_staked_percent"`
		RealStakingYieldPercent interface{} `json:"real_staking_yield_percent"`
	} `json:"staking_stats"`
	MiningStats struct {
		MiningAlgo                 string  `json:"mining_algo"`
		NetworkHashRate            string  `json:"network_hash_rate"`
		AvailableOnNicehashPercent float64 `json:"available_on_nicehash_percent"`
		OneHourAttackCost          float64 `json:"1_hour_attack_cost"`
		Two4HoursAttackCost        float64 `json:"24_hours_attack_cost"`
		AttackAppeal               float64 `json:"attack_appeal"`
	} `json:"mining_stats"`
	DeveloperActivity struct {
		Stars                   int `json:"stars"`
		Watchers                int `json:"watchers"`
		CommitsLast3Months      int `json:"commits_last_3_months"`
		CommitsLast1Year        int `json:"commits_last_1_year"`
		LinesAddedLast3Months   int `json:"lines_added_last_3_months"`
		LinesAddedLast1Year     int `json:"lines_added_last_1_year"`
		LinesDeletedLast3Months int `json:"lines_deleted_last_3_months"`
		LinesDeletedLast1Year   int `json:"lines_deleted_last_1_year"`
	} `json:"developer_activity"`
	RoiData struct {
		PercentChangeLast1Week      float64 `json:"percent_change_last_1_week"`
		PercentChangeLast1Month     float64 `json:"percent_change_last_1_month"`
		PercentChangeLast3Months    float64 `json:"percent_change_last_3_months"`
		PercentChangeLast1Year      float64 `json:"percent_change_last_1_year"`
		PercentChangeBtcLast1Week   float64 `json:"percent_change_btc_last_1_week"`
		PercentChangeBtcLast1Month  float64 `json:"percent_change_btc_last_1_month"`
		PercentChangeBtcLast3Months float64 `json:"percent_change_btc_last_3_months"`
		PercentChangeBtcLast1Year   float64 `json:"percent_change_btc_last_1_year"`
		PercentChangeMonthToDate    float64 `json:"percent_change_month_to_date"`
		PercentChangeQuarterToDate  float64 `json:"percent_change_quarter_to_date"`
		PercentChangeYearToDate     float64 `json:"percent_change_year_to_date"`
	} `json:"roi_data"`
	RoiByYear struct {
		Two018UsdPercent float64 `json:"2018_usd_percent"`
		Two017UsdPercent float64 `json:"2017_usd_percent"`
		Two016UsdPercent float64 `json:"2016_usd_percent"`
		Two015UsdPercent float64 `json:"2015_usd_percent"`
		Two014UsdPercent float64 `json:"2014_usd_percent"`
		Two013UsdPercent float64 `json:"2013_usd_percent"`
		Two012UsdPercent float64 `json:"2012_usd_percent"`
		Two011UsdPercent float64 `json:"2011_usd_percent"`
	} `json:"roi_by_year"`
	RiskMetrics struct {
		SharpeRatios struct {
			Last30Days float64 `json:"last_30_days"`
			Last90Days float64 `json:"last_90_days"`
			Last1Year  float64 `json:"last_1_year"`
			Last3Years float64 `json:"last_3_years"`
		} `json:"sharpe_ratios"`
		VolatilityStats struct {
			VolatilityLast30Days float64 `json:"volatility_last_30_days"`
			VolatilityLast90Days float64 `json:"volatility_last_90_days"`
			VolatilityLast1Year  float64 `json:"volatility_last_1_year"`
			VolatilityLast3Years float64 `json:"volatility_last_3_years"`
		} `json:"volatility_stats"`
	} `json:"risk_metrics"`
	MiscData struct {
		VladimirClubCost                   float64  `json:"vladimir_club_cost"`
		BtcCurrentNormalizedSupplyPriceUsd float64  `json:"btc_current_normalized_supply_price_usd"`
		BtcY2050NormalizedSupplyPriceUsd   float64  `json:"btc_y2050_normalized_supply_price_usd"`
		AssetCreatedAt                     string   `json:"asset_created_at"`
		AssetAgeDays                       int      `json:"asset_age_days"`
		Categories                         []string `json:"categories"`
		Sectors                            []string `json:"sectors"`
	} `json:"misc_data"`
}
