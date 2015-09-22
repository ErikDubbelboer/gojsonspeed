package openrtb

import (
	"encoding/json"
)

type SeatBid struct {
	Bid   []Bid           `json:"bid"`             // Array of 1+ Bid objects (Section 4.2.3) each related to an impression. Multiple bids can relate to the same impression.
	Seat  json.RawMessage `json:"seat,omitempty"`  // ID of the bidder seat on whose behalf this bid is made.
	Group int             `json:"group,omitempty"` // 0 = impressions can be won individually; 1 = impressions must be won or lost as a group.
	//Ext   Extension `json:"ext,omitempty"`   // Placeholder for exchange-specific extensions to OpenRTB.
}
