// @generated by protobuf-ts 2.9.4
// @generated from protobuf file "user/user.proto" (package "user_service", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { UsersService } from "./user";
import type { UpdateUserLocationResponse } from "./user";
import type { UpdateUserLocationRequest } from "./user";
import type { GetUserResponse } from "./user";
import type { GetUserRequest } from "./user";
import type { UserUpdateResponse } from "./user";
import type { UserUpdateRequest } from "./user";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { RegisterUserResponse } from "./user";
import type { RegisterUserRequest } from "./user";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service user_service.UsersService
 */
export interface IUsersServiceClient {
    /**
     * @generated from protobuf rpc: RegisterUser(user_service.RegisterUserRequest) returns (user_service.RegisterUserResponse);
     */
    registerUser(input: RegisterUserRequest, options?: RpcOptions): UnaryCall<RegisterUserRequest, RegisterUserResponse>;
    /**
     * @generated from protobuf rpc: UpdateUser(user_service.UserUpdateRequest) returns (user_service.UserUpdateResponse);
     */
    updateUser(input: UserUpdateRequest, options?: RpcOptions): UnaryCall<UserUpdateRequest, UserUpdateResponse>;
    /**
     * @generated from protobuf rpc: GetUser(user_service.GetUserRequest) returns (user_service.GetUserResponse);
     */
    getUser(input: GetUserRequest, options?: RpcOptions): UnaryCall<GetUserRequest, GetUserResponse>;
    /**
     * @generated from protobuf rpc: UpdateUserLocation(user_service.UpdateUserLocationRequest) returns (user_service.UpdateUserLocationResponse);
     */
    updateUserLocation(input: UpdateUserLocationRequest, options?: RpcOptions): UnaryCall<UpdateUserLocationRequest, UpdateUserLocationResponse>;
}
/**
 * @generated from protobuf service user_service.UsersService
 */
export class UsersServiceClient implements IUsersServiceClient, ServiceInfo {
    typeName = UsersService.typeName;
    methods = UsersService.methods;
    options = UsersService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: RegisterUser(user_service.RegisterUserRequest) returns (user_service.RegisterUserResponse);
     */
    registerUser(input: RegisterUserRequest, options?: RpcOptions): UnaryCall<RegisterUserRequest, RegisterUserResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<RegisterUserRequest, RegisterUserResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: UpdateUser(user_service.UserUpdateRequest) returns (user_service.UserUpdateResponse);
     */
    updateUser(input: UserUpdateRequest, options?: RpcOptions): UnaryCall<UserUpdateRequest, UserUpdateResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<UserUpdateRequest, UserUpdateResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetUser(user_service.GetUserRequest) returns (user_service.GetUserResponse);
     */
    getUser(input: GetUserRequest, options?: RpcOptions): UnaryCall<GetUserRequest, GetUserResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetUserRequest, GetUserResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: UpdateUserLocation(user_service.UpdateUserLocationRequest) returns (user_service.UpdateUserLocationResponse);
     */
    updateUserLocation(input: UpdateUserLocationRequest, options?: RpcOptions): UnaryCall<UpdateUserLocationRequest, UpdateUserLocationResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateUserLocationRequest, UpdateUserLocationResponse>("unary", this._transport, method, opt, input);
    }
}
