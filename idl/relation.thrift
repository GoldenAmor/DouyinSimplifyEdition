namespace go relation

struct BaseResp {
    1:i64 status_code
    2:string status_message
}

struct CreateRelationReq{
    1:i64 user_id
    2:i64 follower_id
}
struct CreateRelationResp{
    1:BaseResp base_resp
}

struct DeleteRelationReq{
    1:i64 user_id
    2:i64 follower_id
}
struct DeleteRelationResp{
    1:BaseResp base_resp
}

struct GetFollowersReq{
    1:i64 user_id
}
struct GetFollowersResp{
    1:BaseResp base_resp
    2:list<i64> followers
}

struct GetFollowsReq{
    1:i64 user_id
}
struct GetFollowsResp{
    1:BaseResp base_resp
    2:list<i64> follows
}

struct CountFollowersReq{
    1:i64 user_id
}
struct CountFollowersResp{
    1:BaseResp base_resp
    2:i64 followers_count
}

struct CountFollowsReq{
    1:i64 user_id
}
struct CountFollowsResp{
    1:BaseResp base_resp
    2:i64 follows_count
}

struct IsFollowReq{
    1:i64 user_id
    2:i64 target_user_id
}
struct IsFollowResp{
    1:BaseResp base_resp
    2:bool is_follow
}

service RelationService {
    CreateRelationResp CreateRelation(1:CreateRelationReq createRelationReq)
    DeleteRelationResp DeleteRelation(1:DeleteRelationReq deleteRelationReq)
    GetFollowersResp GetFollowers(1:GetFollowersReq getFollowersReq)
    GetFollowsResp GetFollows(1:GetFollowsReq getFollowsReq)
    CountFollowersResp CountFollowers(1:CountFollowersReq countFollowersReq)
    CountFollowsResp CountFollows(1:CountFollowsReq countFollowsReq)
    IsFollowResp IsFollow(1:IsFollowReq isFollowReq)
}