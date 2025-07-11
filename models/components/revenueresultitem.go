// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RevenueResultItem struct {
	// Date for the metric (ISO 8601)
	Date       string            `json:"date"`
	Precentage RevenuePercentage `json:"precentage"`
	Current    RevenueCurrent    `json:"current"`
	Previous   RevenuePrevious   `json:"previous"`
}

func (o *RevenueResultItem) GetDate() string {
	if o == nil {
		return ""
	}
	return o.Date
}

func (o *RevenueResultItem) GetPrecentage() RevenuePercentage {
	if o == nil {
		return RevenuePercentage{}
	}
	return o.Precentage
}

func (o *RevenueResultItem) GetCurrent() RevenueCurrent {
	if o == nil {
		return RevenueCurrent{}
	}
	return o.Current
}

func (o *RevenueResultItem) GetPrevious() RevenuePrevious {
	if o == nil {
		return RevenuePrevious{}
	}
	return o.Previous
}
