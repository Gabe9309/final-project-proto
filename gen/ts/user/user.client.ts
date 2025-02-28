// @generated by protobuf-ts 2.9.4
// @generated from protobuf file "user/user.proto" (package "user_service", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { UsersService } from "./user";
import type { UpdateUserLocationResponse } from "./user";
import type { UpdateUserLocationRequest } from "./user";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { UserCreateResponse } from "./user";
import type { UserCreateRequest } from "./user";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service user_service.UsersService
 */
export interface IUsersServiceClient {
    /**
     * @generated from protobuf rpc: CreateNewUser(user_service.UserCreateRequest) returns (user_service.UserCreateResponse);
     */
    createNewUser(input: UserCreateRequest, options?: RpcOptions): UnaryCall<UserCreateRequest, UserCreateResponse>;
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
     * @generated from protobuf rpc: CreateNewUser(user_service.UserCreateRequest) returns (user_service.UserCreateResponse);
     */
    createNewUser(input: UserCreateRequest, options?: RpcOptions): UnaryCall<UserCreateRequest, UserCreateResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<UserCreateRequest, UserCreateResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: UpdateUserLocation(user_service.UpdateUserLocationRequest) returns (user_service.UpdateUserLocationResponse);
     */
    updateUserLocation(input: UpdateUserLocationRequest, options?: RpcOptions): UnaryCall<UpdateUserLocationRequest, UpdateUserLocationResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateUserLocationRequest, UpdateUserLocationResponse>("unary", this._transport, method, opt, input);
    }
}
