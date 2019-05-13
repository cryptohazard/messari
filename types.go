package messari

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
	People []struct {
		FoundingTeam []Team `json:"founding_team"`
		Contrinutors []Team `json:"contributors"`
		Investors    []Team `json:"investors"`
		Advisors     []Team `json:"advisors"`
	} `json:"people"`
	RelevantResources []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"relevant_resources"`
}
