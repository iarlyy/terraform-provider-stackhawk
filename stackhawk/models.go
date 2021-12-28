package stackhawk

type Application struct {
	ApplicationId  string `json:"applicationId,omitempty"`
	Name           string `json:"name"`
	DataType       string `json:"dataType"`
	Env            string `json:"env"`
	EnvId          string `json:"envId,omitempty"`
	RiskLevel      string `json:"riskLevel"`
	OrganizationId string `json:"organizationId,omitempty"`
}
