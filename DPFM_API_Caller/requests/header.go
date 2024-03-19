package requests

type Header struct {
	InspectionLot		int     `json:"InspectionLot"`
	IsReleased			*bool   `json:"IsReleased"`
}
