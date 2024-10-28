package request

// DriverRequest represents the request body for finding nearest driver
type DriverRequest struct {
	Lat      float64 `json:"lat" form:"lat" validate:"required,numeric,min=-90,max=90"`
	Lon      float64 `json:"lon" form:"lon" validate:"required,numeric,min=-180,max=180"`
	Radius   float64 `json:"radius" form:"radius" validate:"required,numeric,min=1"`
	Page     int     `json:"page" form:"page" validate:"omitempty,numeric,min=1"`
	PageSize int     `json:"page_size" form:"page_size" validate:"omitempty,numeric,min=1"`
}
