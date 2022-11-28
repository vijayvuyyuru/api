// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppServiceClient interface {
	CreateLocation(ctx context.Context, in *CreateLocationRequest, opts ...grpc.CallOption) (*CreateLocationResponse, error)
	ListOrganizations(ctx context.Context, in *ListOrganizationsRequest, opts ...grpc.CallOption) (*ListOrganizationsResponse, error)
	ListLocations(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error)
	LocationAuth(ctx context.Context, in *LocationAuthRequest, opts ...grpc.CallOption) (*LocationAuthResponse, error)
	// Create a new generated Secret in the Location.
	//  - Succeeds if there are no more than 2 active secrets after creation.
	CreateLocationSecret(ctx context.Context, in *CreateLocationSecretRequest, opts ...grpc.CallOption) (*CreateLocationSecretResponse, error)
	// Delete a Secret from the Location.
	DeleteLocationSecret(ctx context.Context, in *DeleteLocationSecretRequest, opts ...grpc.CallOption) (*DeleteLocationSecretResponse, error)
	// Get a specific robot by ID
	GetRobot(ctx context.Context, in *GetRobotRequest, opts ...grpc.CallOption) (*GetRobotResponse, error)
	GetRobotParts(ctx context.Context, in *GetRobotPartsRequest, opts ...grpc.CallOption) (*GetRobotPartsResponse, error)
	// Get a specific robot part by ID
	GetRobotPart(ctx context.Context, in *GetRobotPartRequest, opts ...grpc.CallOption) (*GetRobotPartResponse, error)
	GetRobotPartLogs(ctx context.Context, in *GetRobotPartLogsRequest, opts ...grpc.CallOption) (*GetRobotPartLogsResponse, error)
	TailRobotPartLogs(ctx context.Context, in *TailRobotPartLogsRequest, opts ...grpc.CallOption) (AppService_TailRobotPartLogsClient, error)
	// Get a specific robot part histy by ID
	GetRobotPartHistory(ctx context.Context, in *GetRobotPartHistoryRequest, opts ...grpc.CallOption) (*GetRobotPartHistoryResponse, error)
	// Update a robot
	UpdateRobotPart(ctx context.Context, in *UpdateRobotPartRequest, opts ...grpc.CallOption) (*UpdateRobotPartResponse, error)
	// Create a new robot part
	NewRobotPart(ctx context.Context, in *NewRobotPartRequest, opts ...grpc.CallOption) (*NewRobotPartResponse, error)
	// Delete a robot part
	DeleteRobotPart(ctx context.Context, in *DeleteRobotPartRequest, opts ...grpc.CallOption) (*DeleteRobotPartResponse, error)
	// Marks the given part as the main part, and all the others as not
	MarkPartAsMain(ctx context.Context, in *MarkPartAsMainRequest, opts ...grpc.CallOption) (*MarkPartAsMainResponse, error)
	// Create a new generated Secret in the Robot Part.
	//  - Succeeds if there are no more than 2 active secrets after creation.
	CreateRobotPartSecret(ctx context.Context, in *CreateRobotPartSecretRequest, opts ...grpc.CallOption) (*CreateRobotPartSecretResponse, error)
	// Delete a Secret from the RobotPart.
	DeleteRobotPartSecret(ctx context.Context, in *DeleteRobotPartSecretRequest, opts ...grpc.CallOption) (*DeleteRobotPartSecretResponse, error)
	// Finds robots given a query
	FindRobots(ctx context.Context, in *FindRobotsRequest, opts ...grpc.CallOption) (*FindRobotsResponse, error)
	// NewRobot creates a new robot
	NewRobot(ctx context.Context, in *NewRobotRequest, opts ...grpc.CallOption) (*NewRobotResponse, error)
	// UpdateRobot updates a robot
	UpdateRobot(ctx context.Context, in *UpdateRobotRequest, opts ...grpc.CallOption) (*UpdateRobotResponse, error)
	// DeleteRobot deletes a robot
	DeleteRobot(ctx context.Context, in *DeleteRobotRequest, opts ...grpc.CallOption) (*DeleteRobotResponse, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) CreateLocation(ctx context.Context, in *CreateLocationRequest, opts ...grpc.CallOption) (*CreateLocationResponse, error) {
	out := new(CreateLocationResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/CreateLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) ListOrganizations(ctx context.Context, in *ListOrganizationsRequest, opts ...grpc.CallOption) (*ListOrganizationsResponse, error) {
	out := new(ListOrganizationsResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/ListOrganizations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) ListLocations(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error) {
	out := new(ListLocationsResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/ListLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) LocationAuth(ctx context.Context, in *LocationAuthRequest, opts ...grpc.CallOption) (*LocationAuthResponse, error) {
	out := new(LocationAuthResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/LocationAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) CreateLocationSecret(ctx context.Context, in *CreateLocationSecretRequest, opts ...grpc.CallOption) (*CreateLocationSecretResponse, error) {
	out := new(CreateLocationSecretResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/CreateLocationSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) DeleteLocationSecret(ctx context.Context, in *DeleteLocationSecretRequest, opts ...grpc.CallOption) (*DeleteLocationSecretResponse, error) {
	out := new(DeleteLocationSecretResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/DeleteLocationSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetRobot(ctx context.Context, in *GetRobotRequest, opts ...grpc.CallOption) (*GetRobotResponse, error) {
	out := new(GetRobotResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/GetRobot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetRobotParts(ctx context.Context, in *GetRobotPartsRequest, opts ...grpc.CallOption) (*GetRobotPartsResponse, error) {
	out := new(GetRobotPartsResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/GetRobotParts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetRobotPart(ctx context.Context, in *GetRobotPartRequest, opts ...grpc.CallOption) (*GetRobotPartResponse, error) {
	out := new(GetRobotPartResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/GetRobotPart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetRobotPartLogs(ctx context.Context, in *GetRobotPartLogsRequest, opts ...grpc.CallOption) (*GetRobotPartLogsResponse, error) {
	out := new(GetRobotPartLogsResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/GetRobotPartLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) TailRobotPartLogs(ctx context.Context, in *TailRobotPartLogsRequest, opts ...grpc.CallOption) (AppService_TailRobotPartLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AppService_ServiceDesc.Streams[0], "/viam.app.v1.AppService/TailRobotPartLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &appServiceTailRobotPartLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AppService_TailRobotPartLogsClient interface {
	Recv() (*TailRobotPartLogsResponse, error)
	grpc.ClientStream
}

type appServiceTailRobotPartLogsClient struct {
	grpc.ClientStream
}

func (x *appServiceTailRobotPartLogsClient) Recv() (*TailRobotPartLogsResponse, error) {
	m := new(TailRobotPartLogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *appServiceClient) GetRobotPartHistory(ctx context.Context, in *GetRobotPartHistoryRequest, opts ...grpc.CallOption) (*GetRobotPartHistoryResponse, error) {
	out := new(GetRobotPartHistoryResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/GetRobotPartHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) UpdateRobotPart(ctx context.Context, in *UpdateRobotPartRequest, opts ...grpc.CallOption) (*UpdateRobotPartResponse, error) {
	out := new(UpdateRobotPartResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/UpdateRobotPart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) NewRobotPart(ctx context.Context, in *NewRobotPartRequest, opts ...grpc.CallOption) (*NewRobotPartResponse, error) {
	out := new(NewRobotPartResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/NewRobotPart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) DeleteRobotPart(ctx context.Context, in *DeleteRobotPartRequest, opts ...grpc.CallOption) (*DeleteRobotPartResponse, error) {
	out := new(DeleteRobotPartResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/DeleteRobotPart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) MarkPartAsMain(ctx context.Context, in *MarkPartAsMainRequest, opts ...grpc.CallOption) (*MarkPartAsMainResponse, error) {
	out := new(MarkPartAsMainResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/MarkPartAsMain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) CreateRobotPartSecret(ctx context.Context, in *CreateRobotPartSecretRequest, opts ...grpc.CallOption) (*CreateRobotPartSecretResponse, error) {
	out := new(CreateRobotPartSecretResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/CreateRobotPartSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) DeleteRobotPartSecret(ctx context.Context, in *DeleteRobotPartSecretRequest, opts ...grpc.CallOption) (*DeleteRobotPartSecretResponse, error) {
	out := new(DeleteRobotPartSecretResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/DeleteRobotPartSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) FindRobots(ctx context.Context, in *FindRobotsRequest, opts ...grpc.CallOption) (*FindRobotsResponse, error) {
	out := new(FindRobotsResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/FindRobots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) NewRobot(ctx context.Context, in *NewRobotRequest, opts ...grpc.CallOption) (*NewRobotResponse, error) {
	out := new(NewRobotResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/NewRobot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) UpdateRobot(ctx context.Context, in *UpdateRobotRequest, opts ...grpc.CallOption) (*UpdateRobotResponse, error) {
	out := new(UpdateRobotResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/UpdateRobot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) DeleteRobot(ctx context.Context, in *DeleteRobotRequest, opts ...grpc.CallOption) (*DeleteRobotResponse, error) {
	out := new(DeleteRobotResponse)
	err := c.cc.Invoke(ctx, "/viam.app.v1.AppService/DeleteRobot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
// All implementations must embed UnimplementedAppServiceServer
// for forward compatibility
type AppServiceServer interface {
	CreateLocation(context.Context, *CreateLocationRequest) (*CreateLocationResponse, error)
	ListOrganizations(context.Context, *ListOrganizationsRequest) (*ListOrganizationsResponse, error)
	ListLocations(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error)
	LocationAuth(context.Context, *LocationAuthRequest) (*LocationAuthResponse, error)
	// Create a new generated Secret in the Location.
	//  - Succeeds if there are no more than 2 active secrets after creation.
	CreateLocationSecret(context.Context, *CreateLocationSecretRequest) (*CreateLocationSecretResponse, error)
	// Delete a Secret from the Location.
	DeleteLocationSecret(context.Context, *DeleteLocationSecretRequest) (*DeleteLocationSecretResponse, error)
	// Get a specific robot by ID
	GetRobot(context.Context, *GetRobotRequest) (*GetRobotResponse, error)
	GetRobotParts(context.Context, *GetRobotPartsRequest) (*GetRobotPartsResponse, error)
	// Get a specific robot part by ID
	GetRobotPart(context.Context, *GetRobotPartRequest) (*GetRobotPartResponse, error)
	GetRobotPartLogs(context.Context, *GetRobotPartLogsRequest) (*GetRobotPartLogsResponse, error)
	TailRobotPartLogs(*TailRobotPartLogsRequest, AppService_TailRobotPartLogsServer) error
	// Get a specific robot part histy by ID
	GetRobotPartHistory(context.Context, *GetRobotPartHistoryRequest) (*GetRobotPartHistoryResponse, error)
	// Update a robot
	UpdateRobotPart(context.Context, *UpdateRobotPartRequest) (*UpdateRobotPartResponse, error)
	// Create a new robot part
	NewRobotPart(context.Context, *NewRobotPartRequest) (*NewRobotPartResponse, error)
	// Delete a robot part
	DeleteRobotPart(context.Context, *DeleteRobotPartRequest) (*DeleteRobotPartResponse, error)
	// Marks the given part as the main part, and all the others as not
	MarkPartAsMain(context.Context, *MarkPartAsMainRequest) (*MarkPartAsMainResponse, error)
	// Create a new generated Secret in the Robot Part.
	//  - Succeeds if there are no more than 2 active secrets after creation.
	CreateRobotPartSecret(context.Context, *CreateRobotPartSecretRequest) (*CreateRobotPartSecretResponse, error)
	// Delete a Secret from the RobotPart.
	DeleteRobotPartSecret(context.Context, *DeleteRobotPartSecretRequest) (*DeleteRobotPartSecretResponse, error)
	// Finds robots given a query
	FindRobots(context.Context, *FindRobotsRequest) (*FindRobotsResponse, error)
	// NewRobot creates a new robot
	NewRobot(context.Context, *NewRobotRequest) (*NewRobotResponse, error)
	// UpdateRobot updates a robot
	UpdateRobot(context.Context, *UpdateRobotRequest) (*UpdateRobotResponse, error)
	// DeleteRobot deletes a robot
	DeleteRobot(context.Context, *DeleteRobotRequest) (*DeleteRobotResponse, error)
	mustEmbedUnimplementedAppServiceServer()
}

// UnimplementedAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppServiceServer struct {
}

func (UnimplementedAppServiceServer) CreateLocation(context.Context, *CreateLocationRequest) (*CreateLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLocation not implemented")
}
func (UnimplementedAppServiceServer) ListOrganizations(context.Context, *ListOrganizationsRequest) (*ListOrganizationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrganizations not implemented")
}
func (UnimplementedAppServiceServer) ListLocations(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLocations not implemented")
}
func (UnimplementedAppServiceServer) LocationAuth(context.Context, *LocationAuthRequest) (*LocationAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocationAuth not implemented")
}
func (UnimplementedAppServiceServer) CreateLocationSecret(context.Context, *CreateLocationSecretRequest) (*CreateLocationSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLocationSecret not implemented")
}
func (UnimplementedAppServiceServer) DeleteLocationSecret(context.Context, *DeleteLocationSecretRequest) (*DeleteLocationSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocationSecret not implemented")
}
func (UnimplementedAppServiceServer) GetRobot(context.Context, *GetRobotRequest) (*GetRobotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRobot not implemented")
}
func (UnimplementedAppServiceServer) GetRobotParts(context.Context, *GetRobotPartsRequest) (*GetRobotPartsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRobotParts not implemented")
}
func (UnimplementedAppServiceServer) GetRobotPart(context.Context, *GetRobotPartRequest) (*GetRobotPartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRobotPart not implemented")
}
func (UnimplementedAppServiceServer) GetRobotPartLogs(context.Context, *GetRobotPartLogsRequest) (*GetRobotPartLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRobotPartLogs not implemented")
}
func (UnimplementedAppServiceServer) TailRobotPartLogs(*TailRobotPartLogsRequest, AppService_TailRobotPartLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method TailRobotPartLogs not implemented")
}
func (UnimplementedAppServiceServer) GetRobotPartHistory(context.Context, *GetRobotPartHistoryRequest) (*GetRobotPartHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRobotPartHistory not implemented")
}
func (UnimplementedAppServiceServer) UpdateRobotPart(context.Context, *UpdateRobotPartRequest) (*UpdateRobotPartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRobotPart not implemented")
}
func (UnimplementedAppServiceServer) NewRobotPart(context.Context, *NewRobotPartRequest) (*NewRobotPartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewRobotPart not implemented")
}
func (UnimplementedAppServiceServer) DeleteRobotPart(context.Context, *DeleteRobotPartRequest) (*DeleteRobotPartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRobotPart not implemented")
}
func (UnimplementedAppServiceServer) MarkPartAsMain(context.Context, *MarkPartAsMainRequest) (*MarkPartAsMainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkPartAsMain not implemented")
}
func (UnimplementedAppServiceServer) CreateRobotPartSecret(context.Context, *CreateRobotPartSecretRequest) (*CreateRobotPartSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRobotPartSecret not implemented")
}
func (UnimplementedAppServiceServer) DeleteRobotPartSecret(context.Context, *DeleteRobotPartSecretRequest) (*DeleteRobotPartSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRobotPartSecret not implemented")
}
func (UnimplementedAppServiceServer) FindRobots(context.Context, *FindRobotsRequest) (*FindRobotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRobots not implemented")
}
func (UnimplementedAppServiceServer) NewRobot(context.Context, *NewRobotRequest) (*NewRobotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewRobot not implemented")
}
func (UnimplementedAppServiceServer) UpdateRobot(context.Context, *UpdateRobotRequest) (*UpdateRobotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRobot not implemented")
}
func (UnimplementedAppServiceServer) DeleteRobot(context.Context, *DeleteRobotRequest) (*DeleteRobotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRobot not implemented")
}
func (UnimplementedAppServiceServer) mustEmbedUnimplementedAppServiceServer() {}

// UnsafeAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServiceServer will
// result in compilation errors.
type UnsafeAppServiceServer interface {
	mustEmbedUnimplementedAppServiceServer()
}

func RegisterAppServiceServer(s grpc.ServiceRegistrar, srv AppServiceServer) {
	s.RegisterService(&AppService_ServiceDesc, srv)
}

func _AppService_CreateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).CreateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/CreateLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).CreateLocation(ctx, req.(*CreateLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_ListOrganizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).ListOrganizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/ListOrganizations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).ListOrganizations(ctx, req.(*ListOrganizationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_ListLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).ListLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/ListLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).ListLocations(ctx, req.(*ListLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_LocationAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).LocationAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/LocationAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).LocationAuth(ctx, req.(*LocationAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_CreateLocationSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLocationSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).CreateLocationSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/CreateLocationSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).CreateLocationSecret(ctx, req.(*CreateLocationSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_DeleteLocationSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLocationSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).DeleteLocationSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/DeleteLocationSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).DeleteLocationSecret(ctx, req.(*DeleteLocationSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GetRobot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRobotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetRobot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/GetRobot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetRobot(ctx, req.(*GetRobotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GetRobotParts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRobotPartsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetRobotParts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/GetRobotParts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetRobotParts(ctx, req.(*GetRobotPartsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GetRobotPart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRobotPartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetRobotPart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/GetRobotPart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetRobotPart(ctx, req.(*GetRobotPartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GetRobotPartLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRobotPartLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetRobotPartLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/GetRobotPartLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetRobotPartLogs(ctx, req.(*GetRobotPartLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_TailRobotPartLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TailRobotPartLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AppServiceServer).TailRobotPartLogs(m, &appServiceTailRobotPartLogsServer{stream})
}

type AppService_TailRobotPartLogsServer interface {
	Send(*TailRobotPartLogsResponse) error
	grpc.ServerStream
}

type appServiceTailRobotPartLogsServer struct {
	grpc.ServerStream
}

func (x *appServiceTailRobotPartLogsServer) Send(m *TailRobotPartLogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AppService_GetRobotPartHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRobotPartHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetRobotPartHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/GetRobotPartHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetRobotPartHistory(ctx, req.(*GetRobotPartHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_UpdateRobotPart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRobotPartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).UpdateRobotPart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/UpdateRobotPart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).UpdateRobotPart(ctx, req.(*UpdateRobotPartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_NewRobotPart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRobotPartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).NewRobotPart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/NewRobotPart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).NewRobotPart(ctx, req.(*NewRobotPartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_DeleteRobotPart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRobotPartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).DeleteRobotPart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/DeleteRobotPart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).DeleteRobotPart(ctx, req.(*DeleteRobotPartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_MarkPartAsMain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkPartAsMainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).MarkPartAsMain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/MarkPartAsMain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).MarkPartAsMain(ctx, req.(*MarkPartAsMainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_CreateRobotPartSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRobotPartSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).CreateRobotPartSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/CreateRobotPartSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).CreateRobotPartSecret(ctx, req.(*CreateRobotPartSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_DeleteRobotPartSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRobotPartSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).DeleteRobotPartSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/DeleteRobotPartSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).DeleteRobotPartSecret(ctx, req.(*DeleteRobotPartSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_FindRobots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRobotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).FindRobots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/FindRobots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).FindRobots(ctx, req.(*FindRobotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_NewRobot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRobotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).NewRobot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/NewRobot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).NewRobot(ctx, req.(*NewRobotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_UpdateRobot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRobotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).UpdateRobot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/UpdateRobot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).UpdateRobot(ctx, req.(*UpdateRobotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_DeleteRobot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRobotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).DeleteRobot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.v1.AppService/DeleteRobot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).DeleteRobot(ctx, req.(*DeleteRobotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppService_ServiceDesc is the grpc.ServiceDesc for AppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "viam.app.v1.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLocation",
			Handler:    _AppService_CreateLocation_Handler,
		},
		{
			MethodName: "ListOrganizations",
			Handler:    _AppService_ListOrganizations_Handler,
		},
		{
			MethodName: "ListLocations",
			Handler:    _AppService_ListLocations_Handler,
		},
		{
			MethodName: "LocationAuth",
			Handler:    _AppService_LocationAuth_Handler,
		},
		{
			MethodName: "CreateLocationSecret",
			Handler:    _AppService_CreateLocationSecret_Handler,
		},
		{
			MethodName: "DeleteLocationSecret",
			Handler:    _AppService_DeleteLocationSecret_Handler,
		},
		{
			MethodName: "GetRobot",
			Handler:    _AppService_GetRobot_Handler,
		},
		{
			MethodName: "GetRobotParts",
			Handler:    _AppService_GetRobotParts_Handler,
		},
		{
			MethodName: "GetRobotPart",
			Handler:    _AppService_GetRobotPart_Handler,
		},
		{
			MethodName: "GetRobotPartLogs",
			Handler:    _AppService_GetRobotPartLogs_Handler,
		},
		{
			MethodName: "GetRobotPartHistory",
			Handler:    _AppService_GetRobotPartHistory_Handler,
		},
		{
			MethodName: "UpdateRobotPart",
			Handler:    _AppService_UpdateRobotPart_Handler,
		},
		{
			MethodName: "NewRobotPart",
			Handler:    _AppService_NewRobotPart_Handler,
		},
		{
			MethodName: "DeleteRobotPart",
			Handler:    _AppService_DeleteRobotPart_Handler,
		},
		{
			MethodName: "MarkPartAsMain",
			Handler:    _AppService_MarkPartAsMain_Handler,
		},
		{
			MethodName: "CreateRobotPartSecret",
			Handler:    _AppService_CreateRobotPartSecret_Handler,
		},
		{
			MethodName: "DeleteRobotPartSecret",
			Handler:    _AppService_DeleteRobotPartSecret_Handler,
		},
		{
			MethodName: "FindRobots",
			Handler:    _AppService_FindRobots_Handler,
		},
		{
			MethodName: "NewRobot",
			Handler:    _AppService_NewRobot_Handler,
		},
		{
			MethodName: "UpdateRobot",
			Handler:    _AppService_UpdateRobot_Handler,
		},
		{
			MethodName: "DeleteRobot",
			Handler:    _AppService_DeleteRobot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TailRobotPartLogs",
			Handler:       _AppService_TailRobotPartLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "app/v1/app.proto",
}