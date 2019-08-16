package messari

import (
	"testing"
)

//TODO tests more coins since the info can be very different

func TestMarkets(t *testing.T) {
	client := Client{}
	markets, err := client.Markets()
	if err != nil {
		t.Errorf("expected an error")
	}
	if len(markets) == 0 {
		t.Errorf("expected markets data")
	}
}

func TestAssets(t *testing.T) {
	client := Client{}
	assets, err := client.Assets()
	if err != nil {
		t.Errorf("expected an error")
	}
	if len(assets) == 0 {
		t.Errorf("expected assets data")
	}
}

func TestAssetsWithMetrics(t *testing.T) {
	client := Client{}
	assets, err := client.AssetsWithMetrics()
	if err != nil {
		t.Errorf("expected an error")
	}
	if len(assets) == 0 {
		t.Errorf("expected assets data")
	}
}

func TestAssetsWithProfiles(t *testing.T) {
	client := Client{}
	assets, err := client.AssetsWithProfiles()
	if err != nil {
		t.Errorf("expected an error")
	}
	if len(assets) == 0 {
		t.Errorf("expected assets data")
	}
}

func TestAssetsWithMetricsAndProfiles(t *testing.T) {
	client := Client{}
	assets, err := client.AssetsWithMetricsAndProfiles()
	if err != nil {
		t.Errorf("expected an error")
	}
	if len(assets) == 0 {
		t.Errorf("expected assets data")
	}
}

func TestProfileBySymbol(t *testing.T) {
	client := Client{}
	profile, err := client.ProfileBySymbol("btc")
	if err != nil {
		t.Errorf("expected an error\n%v", err)
	}
	if profile.Name != "Bitcoin" || profile.TokenDistribution.MaxSupply != 21000000 {
		t.Errorf("expected profile data")
	}
}

func TestGetMetricsBySymbol(t *testing.T) {
	client := Client{}
	metrics, err := client.GetMetricsBySymbol("eth")
	if err != nil {
		t.Errorf("expected an error\n%v", err)
	}
	if metrics.Name != "Ethereum" || metrics.MiscData.AssetCreatedAt != "2015-07-30" {
		t.Errorf("expected metrics data")
	}
}
