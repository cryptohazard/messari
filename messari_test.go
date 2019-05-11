package messari

import (
	"testing"
)

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
