package bfapi

//Login Status values
const (
	Success    string = "SUCCESS"
	Limited    string = "LIMITED_ACCESS"
	Restricted string = "LOGIN_RESTRICTED"
	Fail       string = "FAIL"
)

//
type CertLoginResult struct {
	SessionToken string `json:"sessionToken"`
	LoginStatus  string `json:"loginStatus"`
}

//
type LoginResult struct {
	SessionToken string `json:"token"`
	AppKey       string `json:"product"`
	Status       string `json:"status"`
	Error        string `json:"error"`
}

//
type LoginError struct {
	Status string
	Err    string
}

//
func (err LoginError) Error() string {
	return err.Status + " " + err.Err
}
