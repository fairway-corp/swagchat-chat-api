package rest

import (
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/swagchat/chat-api/model"
	"github.com/swagchat/chat-api/service"
	"github.com/swagchat/chat-api/tracer"
	scpb "github.com/swagchat/protobuf/protoc-gen-go"
)

func setBlockUserMux() {
	mux.PostFunc("/users/#userId^[a-z0-9-]$/blockUsers", commonHandler(selfResourceAuthzHandler(postBlockUsers)))
	mux.GetFunc("/users/#userId^[a-z0-9-]$/blockUsers", commonHandler(selfResourceAuthzHandler(getBlockUsers)))
	mux.GetFunc("/users/#userId^[a-z0-9-]$/blockedUsers", commonHandler(selfResourceAuthzHandler(getBlockedUsers)))
	mux.PutFunc("/users/#userId^[a-z0-9-]$/blockUsers", commonHandler(selfResourceAuthzHandler(putBlockUsers)))
	mux.DeleteFunc("/users/#userId^[a-z0-9-]$/blockUsers", commonHandler(selfResourceAuthzHandler(deleteBlockUsers)))
}

func postBlockUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := tracer.Provider(ctx).StartSpan("postBlockUsers", "rest")
	defer tracer.Provider(ctx).Finish(span)

	var req model.CreateBlockUsersRequest
	if err := decodeBody(r, &req); err != nil {
		respondJSONDecodeError(w, r, "")
		return
	}

	req.UserID = bone.GetValue(r, "userId")

	errRes := service.CreateBlockUsers(ctx, &req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusNoContent, "application/json", nil)
}

func getBlockUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := tracer.Provider(ctx).StartSpan("getBlockUsers", "rest")
	defer tracer.Provider(ctx).Finish(span)

	req := &model.GetBlockUsersRequest{}
	req.UserID = bone.GetValue(r, "userId")

	responseType := bone.GetValue(r, "responseType")
	if responseType == "UserIdList" {
		req.ResponseType = scpb.ResponseType_UserIdList
	} else {
		req.ResponseType = scpb.ResponseType_UserList
	}

	blockUsers, errRes := service.GetBlockUsers(ctx, req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusOK, "application/json", blockUsers)
}

func getBlockedUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := tracer.Provider(ctx).StartSpan("getBlockedUsers", "rest")
	defer tracer.Provider(ctx).Finish(span)

	req := &model.GetBlockedUsersRequest{}
	req.UserID = bone.GetValue(r, "userId")

	responseType := bone.GetValue(r, "responseType")
	if responseType == "UserIdList" {
		req.ResponseType = scpb.ResponseType_UserIdList
	} else {
		req.ResponseType = scpb.ResponseType_UserList
	}

	blockedUsers, errRes := service.GetBlockedUsers(ctx, req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusOK, "application/json", blockedUsers)
}

func putBlockUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := tracer.Provider(ctx).StartSpan("putBlockUsers", "rest")
	defer tracer.Provider(ctx).Finish(span)

	var req model.AddBlockUsersRequest
	if err := decodeBody(r, &req); err != nil {
		respondJSONDecodeError(w, r, "")
		return
	}

	req.UserID = bone.GetValue(r, "userId")

	errRes := service.AddBlockUsers(ctx, &req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusNoContent, "application/json", nil)
}

func deleteBlockUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := tracer.Provider(ctx).StartSpan("deleteBlockUsers", "rest")
	defer tracer.Provider(ctx).Finish(span)

	var req model.DeleteBlockUsersRequest
	if err := decodeBody(r, &req); err != nil {
		respondJSONDecodeError(w, r, "")
		return
	}

	req.UserID = bone.GetValue(r, "userId")

	errRes := service.DeleteBlockUsers(ctx, &req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusNoContent, "application/json", nil)
}
