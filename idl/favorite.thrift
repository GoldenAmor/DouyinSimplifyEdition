namespace go favorite

struct BaseResp {
    1:i64 status_code
    2:string status_message
}

struct IsFavoriteReq{
    i64 user_id
    i64 video_id
}
struct IsFavoriteResp{
    1:BaseResp base_resp
    2:bool is_favorite
}

struct LikeReq{
    1:i64 user_id
    2:i64 video_id
}
struct LikeResp{
    1:BaseResp base_resp
}

struct UnLikeReq{
    1:i64 user_id
    2:i64 video_id
}
struct UnLikeResp{
    1:BaseResp base_resp
}

struct GetFavoritesByUserIdReq{
    1:i64 user_id
}
struct GetFavoritesByUserIdResp{
    1:BaseResp base_resp
    2:list<i64> favorites
}

struct CountFavoriteReq{
    1:i64 videoId
}
struct CountFavoriteResp{
    1:BaseResp base_resp
    2:i64 favorite_count
}

service FavoriteService {
    IsFavoriteResp IsFavorite(1:IsFavoriteReq isFavoriteReq)
    LikeResp Like(1:LikeReq likeReq)
    UnLikeResp UnLike(1:UnLikeReq unLikeReq)
    GetFavoritesByUserIdResp GetFavoritesByUserId(1:GetFavoritesByUserIdReq getFavoritesByUserIdReq)
    CountFavoriteResp CountFavorite(1:CountFavoriteReq countFavoriteReq)
}
