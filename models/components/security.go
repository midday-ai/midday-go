// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Security struct {
	Token *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=midday_token"`
}

func (o *Security) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}
