package f3

type Dictionary map[string]interface{}

type ErrorMessage struct {
	Message string `json:"error_message"`
}

type Links struct {
	Self *string `json:"self"`
}

type CreateAccountRequest struct {
	Data *AccountData `json:"data"`
}

type CreateAccountResponse struct {
	Data  *AccountData `json:"data"`
	Links *Links       `json:"links,omitempty"`
}

type FetchAccountResponse struct {
	Data *AccountData `json:"data"`
}

type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string      `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool        `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string       `json:"account_number,omitempty"`
	AlternativeNames        []string     `json:"alternative_names,omitempty"`
	BankID                  string       `json:"bank_id,omitempty"`
	BankIDCode              string       `json:"bank_id_code,omitempty"`
	BaseCurrency            string       `json:"base_currency,omitempty"`
	Bic                     string       `json:"bic,omitempty"`
	Country                 *string      `json:"country,omitempty"`
	Iban                    string       `json:"iban,omitempty"`
	JointAccount            *bool        `json:"joint_account,omitempty"`
	Name                    []string     `json:"name,omitempty"`
	SecondaryIdentification string       `json:"secondary_identification,omitempty"`
	Status                  *string      `json:"status,omitempty"`
	Switched                *bool        `json:"switched,omitempty"`
	ValidationType          *string      `json:"validation_type,omitempty"`
	ReferenceMask           *string      `json:"reference_mask,omitempty"`
	AcceptanceQualifier     *string      `json:"acceptance_qualifier,omitempty"`
	UserDefinedData         []Dictionary `json:"user_defined_data,omitempty"`
}
