// package: swagchat.protobuf
// file: userMessage.proto

import * as jspb from "google-protobuf";
import * as gogoproto_gogo_pb from "./gogoproto/gogo_pb";
import * as deviceMessage_pb from "./deviceMessage_pb";
import * as roomMessage_pb from "./roomMessage_pb";
import * as commonMessage_pb from "./commonMessage_pb";

export class User extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): number | undefined;
  setId(value: number): void;

  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  hasName(): boolean;
  clearName(): void;
  getName(): string | undefined;
  setName(value: string): void;

  hasPictureUrl(): boolean;
  clearPictureUrl(): void;
  getPictureUrl(): string | undefined;
  setPictureUrl(value: string): void;

  hasInformationUrl(): boolean;
  clearInformationUrl(): void;
  getInformationUrl(): string | undefined;
  setInformationUrl(value: string): void;

  hasUnreadCount(): boolean;
  clearUnreadCount(): void;
  getUnreadCount(): number | undefined;
  setUnreadCount(value: number): void;

  hasMetaData(): boolean;
  clearMetaData(): void;
  getMetaData(): Uint8Array | string;
  getMetaData_asU8(): Uint8Array;
  getMetaData_asB64(): string;
  setMetaData(value: Uint8Array | string): void;

  hasPublic(): boolean;
  clearPublic(): void;
  getPublic(): boolean | undefined;
  setPublic(value: boolean): void;

  hasCanBlock(): boolean;
  clearCanBlock(): void;
  getCanBlock(): boolean | undefined;
  setCanBlock(value: boolean): void;

  hasLang(): boolean;
  clearLang(): void;
  getLang(): string | undefined;
  setLang(value: string): void;

  hasAccessToken(): boolean;
  clearAccessToken(): void;
  getAccessToken(): string | undefined;
  setAccessToken(value: string): void;

  hasLastAccessRoomId(): boolean;
  clearLastAccessRoomId(): void;
  getLastAccessRoomId(): string | undefined;
  setLastAccessRoomId(value: string): void;

  hasLastAccessed(): boolean;
  clearLastAccessed(): void;
  getLastAccessed(): number | undefined;
  setLastAccessed(value: number): void;

  hasCreated(): boolean;
  clearCreated(): void;
  getCreated(): number | undefined;
  setCreated(value: number): void;

  hasModified(): boolean;
  clearModified(): void;
  getModified(): number | undefined;
  setModified(value: number): void;

  hasDeleted(): boolean;
  clearDeleted(): void;
  getDeleted(): number | undefined;
  setDeleted(value: number): void;

  clearBlockUsersList(): void;
  getBlockUsersList(): Array<string>;
  setBlockUsersList(value: Array<string>): void;
  addBlockUsers(value: string, index?: number): string;

  clearDevicesList(): void;
  getDevicesList(): Array<deviceMessage_pb.Device>;
  setDevicesList(value: Array<deviceMessage_pb.Device>): void;
  addDevices(value?: deviceMessage_pb.Device, index?: number): deviceMessage_pb.Device;

  clearRolesList(): void;
  getRolesList(): Array<number>;
  setRolesList(value: Array<number>): void;
  addRoles(value: number, index?: number): number;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id?: number,
    userId?: string,
    name?: string,
    pictureUrl?: string,
    informationUrl?: string,
    unreadCount?: number,
    metaData: Uint8Array | string,
    pb_public?: boolean,
    canBlock?: boolean,
    lang?: string,
    accessToken?: string,
    lastAccessRoomId?: string,
    lastAccessed?: number,
    created?: number,
    modified?: number,
    deleted?: number,
    blockUsersList: Array<string>,
    devicesList: Array<deviceMessage_pb.Device.AsObject>,
    rolesList: Array<number>,
  }
}

export class MiniRoom extends jspb.Message {
  hasRoomId(): boolean;
  clearRoomId(): void;
  getRoomId(): string | undefined;
  setRoomId(value: string): void;

  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  hasName(): boolean;
  clearName(): void;
  getName(): string | undefined;
  setName(value: string): void;

  hasPictureUrl(): boolean;
  clearPictureUrl(): void;
  getPictureUrl(): string | undefined;
  setPictureUrl(value: string): void;

  hasInformationUrl(): boolean;
  clearInformationUrl(): void;
  getInformationUrl(): string | undefined;
  setInformationUrl(value: string): void;

  hasMetaData(): boolean;
  clearMetaData(): void;
  getMetaData(): Uint8Array | string;
  getMetaData_asU8(): Uint8Array;
  getMetaData_asB64(): string;
  setMetaData(value: Uint8Array | string): void;

  hasType(): boolean;
  clearType(): void;
  getType(): roomMessage_pb.RoomType | undefined;
  setType(value: roomMessage_pb.RoomType): void;

  hasLastMessage(): boolean;
  clearLastMessage(): void;
  getLastMessage(): string | undefined;
  setLastMessage(value: string): void;

  hasLastMessageUpdated(): boolean;
  clearLastMessageUpdated(): void;
  getLastMessageUpdated(): number | undefined;
  setLastMessageUpdated(value: number): void;

  hasCanLeft(): boolean;
  clearCanLeft(): void;
  getCanLeft(): boolean | undefined;
  setCanLeft(value: boolean): void;

  hasCreated(): boolean;
  clearCreated(): void;
  getCreated(): number | undefined;
  setCreated(value: number): void;

  hasModified(): boolean;
  clearModified(): void;
  getModified(): number | undefined;
  setModified(value: number): void;

  clearUsersList(): void;
  getUsersList(): Array<roomMessage_pb.MiniUser>;
  setUsersList(value: Array<roomMessage_pb.MiniUser>): void;
  addUsers(value?: roomMessage_pb.MiniUser, index?: number): roomMessage_pb.MiniUser;

  hasRuUnreadCount(): boolean;
  clearRuUnreadCount(): void;
  getRuUnreadCount(): number | undefined;
  setRuUnreadCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MiniRoom.AsObject;
  static toObject(includeInstance: boolean, msg: MiniRoom): MiniRoom.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MiniRoom, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MiniRoom;
  static deserializeBinaryFromReader(message: MiniRoom, reader: jspb.BinaryReader): MiniRoom;
}

export namespace MiniRoom {
  export type AsObject = {
    roomId?: string,
    userId?: string,
    name?: string,
    pictureUrl?: string,
    informationUrl?: string,
    metaData: Uint8Array | string,
    type?: roomMessage_pb.RoomType,
    lastMessage?: string,
    lastMessageUpdated?: number,
    canLeft?: boolean,
    created?: number,
    modified?: number,
    usersList: Array<roomMessage_pb.MiniUser.AsObject>,
    ruUnreadCount?: number,
  }
}

export class CreateUserRequest extends jspb.Message {
  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  hasName(): boolean;
  clearName(): void;
  getName(): string | undefined;
  setName(value: string): void;

  hasPictureUrl(): boolean;
  clearPictureUrl(): void;
  getPictureUrl(): string | undefined;
  setPictureUrl(value: string): void;

  hasInformationUrl(): boolean;
  clearInformationUrl(): void;
  getInformationUrl(): string | undefined;
  setInformationUrl(value: string): void;

  hasMetaData(): boolean;
  clearMetaData(): void;
  getMetaData(): Uint8Array | string;
  getMetaData_asU8(): Uint8Array;
  getMetaData_asB64(): string;
  setMetaData(value: Uint8Array | string): void;

  hasPublic(): boolean;
  clearPublic(): void;
  getPublic(): boolean | undefined;
  setPublic(value: boolean): void;

  hasCanBlock(): boolean;
  clearCanBlock(): void;
  getCanBlock(): boolean | undefined;
  setCanBlock(value: boolean): void;

  hasLang(): boolean;
  clearLang(): void;
  getLang(): string | undefined;
  setLang(value: string): void;

  clearBlockUsersList(): void;
  getBlockUsersList(): Array<string>;
  setBlockUsersList(value: Array<string>): void;
  addBlockUsers(value: string, index?: number): string;

  clearRolesList(): void;
  getRolesList(): Array<number>;
  setRolesList(value: Array<number>): void;
  addRoles(value: number, index?: number): number;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserRequest): CreateUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserRequest;
  static deserializeBinaryFromReader(message: CreateUserRequest, reader: jspb.BinaryReader): CreateUserRequest;
}

export namespace CreateUserRequest {
  export type AsObject = {
    userId?: string,
    name?: string,
    pictureUrl?: string,
    informationUrl?: string,
    metaData: Uint8Array | string,
    pb_public?: boolean,
    canBlock?: boolean,
    lang?: string,
    blockUsersList: Array<string>,
    rolesList: Array<number>,
  }
}

export class GetUsersRequest extends jspb.Message {
  hasLimit(): boolean;
  clearLimit(): void;
  getLimit(): number | undefined;
  setLimit(value: number): void;

  hasOffset(): boolean;
  clearOffset(): void;
  getOffset(): number | undefined;
  setOffset(value: number): void;

  clearOrdersList(): void;
  getOrdersList(): Array<commonMessage_pb.OrderInfo>;
  setOrdersList(value: Array<commonMessage_pb.OrderInfo>): void;
  addOrders(value?: commonMessage_pb.OrderInfo, index?: number): commonMessage_pb.OrderInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUsersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUsersRequest): GetUsersRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUsersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUsersRequest;
  static deserializeBinaryFromReader(message: GetUsersRequest, reader: jspb.BinaryReader): GetUsersRequest;
}

export namespace GetUsersRequest {
  export type AsObject = {
    limit?: number,
    offset?: number,
    ordersList: Array<commonMessage_pb.OrderInfo.AsObject>,
  }
}

export class UsersResponse extends jspb.Message {
  clearUsersList(): void;
  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): void;
  addUsers(value?: User, index?: number): User;

  hasAllcount(): boolean;
  clearAllcount(): void;
  getAllcount(): number | undefined;
  setAllcount(value: number): void;

  hasLimit(): boolean;
  clearLimit(): void;
  getLimit(): number | undefined;
  setLimit(value: number): void;

  hasOffset(): boolean;
  clearOffset(): void;
  getOffset(): number | undefined;
  setOffset(value: number): void;

  clearOrdersList(): void;
  getOrdersList(): Array<commonMessage_pb.OrderInfo>;
  setOrdersList(value: Array<commonMessage_pb.OrderInfo>): void;
  addOrders(value?: commonMessage_pb.OrderInfo, index?: number): commonMessage_pb.OrderInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UsersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UsersResponse): UsersResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UsersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UsersResponse;
  static deserializeBinaryFromReader(message: UsersResponse, reader: jspb.BinaryReader): UsersResponse;
}

export namespace UsersResponse {
  export type AsObject = {
    usersList: Array<User.AsObject>,
    allcount?: number,
    limit?: number,
    offset?: number,
    ordersList: Array<commonMessage_pb.OrderInfo.AsObject>,
  }
}

export class GetUserRequest extends jspb.Message {
  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserRequest): GetUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserRequest;
  static deserializeBinaryFromReader(message: GetUserRequest, reader: jspb.BinaryReader): GetUserRequest;
}

export namespace GetUserRequest {
  export type AsObject = {
    userId?: string,
  }
}

export class UpdateUserRequest extends jspb.Message {
  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  hasName(): boolean;
  clearName(): void;
  getName(): string | undefined;
  setName(value: string): void;

  hasPictureUrl(): boolean;
  clearPictureUrl(): void;
  getPictureUrl(): string | undefined;
  setPictureUrl(value: string): void;

  hasInformationUrl(): boolean;
  clearInformationUrl(): void;
  getInformationUrl(): string | undefined;
  setInformationUrl(value: string): void;

  hasMetaData(): boolean;
  clearMetaData(): void;
  getMetaData(): Uint8Array | string;
  getMetaData_asU8(): Uint8Array;
  getMetaData_asB64(): string;
  setMetaData(value: Uint8Array | string): void;

  hasPublic(): boolean;
  clearPublic(): void;
  getPublic(): boolean | undefined;
  setPublic(value: boolean): void;

  hasCanBlock(): boolean;
  clearCanBlock(): void;
  getCanBlock(): boolean | undefined;
  setCanBlock(value: boolean): void;

  hasLang(): boolean;
  clearLang(): void;
  getLang(): string | undefined;
  setLang(value: string): void;

  clearBlockUsersList(): void;
  getBlockUsersList(): Array<string>;
  setBlockUsersList(value: Array<string>): void;
  addBlockUsers(value: string, index?: number): string;

  clearRolesList(): void;
  getRolesList(): Array<number>;
  setRolesList(value: Array<number>): void;
  addRoles(value: number, index?: number): number;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUserRequest): UpdateUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUserRequest;
  static deserializeBinaryFromReader(message: UpdateUserRequest, reader: jspb.BinaryReader): UpdateUserRequest;
}

export namespace UpdateUserRequest {
  export type AsObject = {
    userId?: string,
    name?: string,
    pictureUrl?: string,
    informationUrl?: string,
    metaData: Uint8Array | string,
    pb_public?: boolean,
    canBlock?: boolean,
    lang?: string,
    blockUsersList: Array<string>,
    rolesList: Array<number>,
  }
}

export class DeleteUserRequest extends jspb.Message {
  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteUserRequest): DeleteUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteUserRequest;
  static deserializeBinaryFromReader(message: DeleteUserRequest, reader: jspb.BinaryReader): DeleteUserRequest;
}

export namespace DeleteUserRequest {
  export type AsObject = {
    userId?: string,
  }
}

export class GetUserRoomsRequest extends jspb.Message {
  hasLimit(): boolean;
  clearLimit(): void;
  getLimit(): number | undefined;
  setLimit(value: number): void;

  hasOffset(): boolean;
  clearOffset(): void;
  getOffset(): number | undefined;
  setOffset(value: number): void;

  clearOrdersList(): void;
  getOrdersList(): Array<commonMessage_pb.OrderInfo>;
  setOrdersList(value: Array<commonMessage_pb.OrderInfo>): void;
  addOrders(value?: commonMessage_pb.OrderInfo, index?: number): commonMessage_pb.OrderInfo;

  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  hasFilter(): boolean;
  clearFilter(): void;
  getFilter(): UserRoomsFilter | undefined;
  setFilter(value: UserRoomsFilter): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserRoomsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserRoomsRequest): GetUserRoomsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUserRoomsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserRoomsRequest;
  static deserializeBinaryFromReader(message: GetUserRoomsRequest, reader: jspb.BinaryReader): GetUserRoomsRequest;
}

export namespace GetUserRoomsRequest {
  export type AsObject = {
    limit?: number,
    offset?: number,
    ordersList: Array<commonMessage_pb.OrderInfo.AsObject>,
    userId?: string,
    filter?: UserRoomsFilter,
  }
}

export class UserRoomsResponse extends jspb.Message {
  clearRoomsList(): void;
  getRoomsList(): Array<MiniRoom>;
  setRoomsList(value: Array<MiniRoom>): void;
  addRooms(value?: MiniRoom, index?: number): MiniRoom;

  hasAllcount(): boolean;
  clearAllcount(): void;
  getAllcount(): number | undefined;
  setAllcount(value: number): void;

  hasLimit(): boolean;
  clearLimit(): void;
  getLimit(): number | undefined;
  setLimit(value: number): void;

  hasOffset(): boolean;
  clearOffset(): void;
  getOffset(): number | undefined;
  setOffset(value: number): void;

  clearOrdersList(): void;
  getOrdersList(): Array<commonMessage_pb.OrderInfo>;
  setOrdersList(value: Array<commonMessage_pb.OrderInfo>): void;
  addOrders(value?: commonMessage_pb.OrderInfo, index?: number): commonMessage_pb.OrderInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserRoomsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UserRoomsResponse): UserRoomsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UserRoomsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserRoomsResponse;
  static deserializeBinaryFromReader(message: UserRoomsResponse, reader: jspb.BinaryReader): UserRoomsResponse;
}

export namespace UserRoomsResponse {
  export type AsObject = {
    roomsList: Array<MiniRoom.AsObject>,
    allcount?: number,
    limit?: number,
    offset?: number,
    ordersList: Array<commonMessage_pb.OrderInfo.AsObject>,
  }
}

export class GetContactsRequest extends jspb.Message {
  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  hasLimit(): boolean;
  clearLimit(): void;
  getLimit(): number | undefined;
  setLimit(value: number): void;

  hasOffset(): boolean;
  clearOffset(): void;
  getOffset(): number | undefined;
  setOffset(value: number): void;

  clearOrdersList(): void;
  getOrdersList(): Array<commonMessage_pb.OrderInfo>;
  setOrdersList(value: Array<commonMessage_pb.OrderInfo>): void;
  addOrders(value?: commonMessage_pb.OrderInfo, index?: number): commonMessage_pb.OrderInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetContactsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetContactsRequest): GetContactsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetContactsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetContactsRequest;
  static deserializeBinaryFromReader(message: GetContactsRequest, reader: jspb.BinaryReader): GetContactsRequest;
}

export namespace GetContactsRequest {
  export type AsObject = {
    userId?: string,
    limit?: number,
    offset?: number,
    ordersList: Array<commonMessage_pb.OrderInfo.AsObject>,
  }
}

export class GetProfileRequest extends jspb.Message {
  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetProfileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetProfileRequest): GetProfileRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetProfileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetProfileRequest;
  static deserializeBinaryFromReader(message: GetProfileRequest, reader: jspb.BinaryReader): GetProfileRequest;
}

export namespace GetProfileRequest {
  export type AsObject = {
    userId?: string,
  }
}

export class GetRoleUsersRequest extends jspb.Message {
  hasRoleId(): boolean;
  clearRoleId(): void;
  getRoleId(): number | undefined;
  setRoleId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRoleUsersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRoleUsersRequest): GetRoleUsersRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetRoleUsersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRoleUsersRequest;
  static deserializeBinaryFromReader(message: GetRoleUsersRequest, reader: jspb.BinaryReader): GetRoleUsersRequest;
}

export namespace GetRoleUsersRequest {
  export type AsObject = {
    roleId?: number,
  }
}

export class RoleUsersResponse extends jspb.Message {
  clearUsersList(): void;
  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): void;
  addUsers(value?: User, index?: number): User;

  clearUserIdsList(): void;
  getUserIdsList(): Array<string>;
  setUserIdsList(value: Array<string>): void;
  addUserIds(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoleUsersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RoleUsersResponse): RoleUsersResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RoleUsersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoleUsersResponse;
  static deserializeBinaryFromReader(message: RoleUsersResponse, reader: jspb.BinaryReader): RoleUsersResponse;
}

export namespace RoleUsersResponse {
  export type AsObject = {
    usersList: Array<User.AsObject>,
    userIdsList: Array<string>,
  }
}

export enum UserRoomsFilter {
  ONLINE = 0,
  UNREAD = 1,
}
