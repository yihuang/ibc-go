package types

// NewQueryPortResponse creates a new QueryPortResponse instance
func NewQueryPortResponse(portID, app string) *QueryPortResponse {
	return &QueryPortResponse{
		PortId:      portID,
		Application: app,
	}
}

// NewQueryAppVersionResponse creates a new QueryAppVersionResponse instance
func NewQueryAppVersionResponse(portID, version string) *QueryAppVersionResponse {
	return &QueryAppVersionResponse{
		PortId:  portID,
		Version: version,
	}
}
