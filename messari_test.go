package messari

import (
	"testing"
)

//TODO tests more coins since the info can be very different

func TestMarkets(t *testing.T) {
	client := Client{}
	markets, err := client.Markets()
	if err != nil {
		t.Errorf("Markets: unexpected error")
	}
	if len(markets) == 0 {
		t.Errorf("Markets: expected markets data")
	}
}

func TestAssets(t *testing.T) {
	client := Client{}
	assets, err := client.Assets()
	if err != nil {
		t.Errorf("Assets: unexpected error")
	}
	if len(assets) == 0 {
		t.Errorf("Assets: expected assets data")
	}
}

func TestAssetsWithMetrics(t *testing.T) {
	client := Client{}
	assets, err := client.AssetsWithMetrics()
	if err != nil {
		t.Errorf("AssetsWithMetrics: unexpected error")
	}
	if len(assets) == 0 {
		t.Errorf("AssetsWithMetrics: expected assets data")
	}
}

func TestAssetsWithProfiles(t *testing.T) {
	client := Client{}
	assets, err := client.AssetsWithProfiles()
	if err != nil {
		t.Errorf("AssetsWithProfiles: unexpected error")
	}
	if len(assets) == 0 {
		t.Errorf("AssetsWithProfiles: expected assets data")
	}
}

func TestAssetsWithMetricsAndProfiles(t *testing.T) {
	client := Client{}
	assets, err := client.AssetsWithMetricsAndProfiles()
	if err != nil {
		t.Errorf("AssetsWithMetricsAndProfiles: unexpected error")
	}
	if len(assets) == 0 {
		t.Errorf("AssetsWithMetricsAndProfiles: expected assets data")
	}
}

func TestProfileBySymbol(t *testing.T) {
	client := Client{}
	profile, err := client.ProfileBySymbol("btc")
	if err != nil {
		t.Errorf("ProfileBySymbol: unexpected error\n%v", err)
	}
	if profile.Name != "Bitcoin" || profile.TokenDistribution.MaxSupply != 21000000 {
		t.Errorf("ProfileBySymbol: expected profile data")
	}
}

func TestMetricsBySymbol(t *testing.T) {
	client := Client{}
	metrics, err := client.MetricsBySymbol("eth")
	if err != nil {
		t.Errorf("MetricsBySymbol: unexpected error\n%v", err)
	}
	if metrics.Name != "Ethereum" || metrics.MiscData.AssetCreatedAt != "2015-07-30" {
		t.Errorf("MetricsBySymbol: expected metrics data")
	}
}

func TestNews(t *testing.T) {
	client := Client{}
	news, err := client.News()
	if err != nil {
		t.Errorf("News: unexpected error")
	}
	if len(news) <= 10 {
		t.Errorf("News: expected news data")
	}
}

func TestNewsBySymbol(t *testing.T) {
	client := Client{}
	news, err := client.NewsBySymbol("btc")
	if err != nil {
		t.Errorf("NewsBySymbol: unexpected error")
	}
	if len(news) == 0 {
		t.Errorf("NewsBySymbol: expected news data (or there is no news on that asset)")
	}
}
