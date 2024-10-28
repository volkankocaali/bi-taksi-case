package request

// MatchRequest represents the request body for finding nearest driver
type MatchRequest struct {
	Lat    float64 `json:"lat" form:"lat" validate:"required,numeric,min=-90,max=90"`
	Lon    float64 `json:"lon" form:"lon" validate:"required,numeric,min=-180,max=180"`
	Radius int     `json:"radius" form:"radius" validate:"required,numeric,min=1"`
}
