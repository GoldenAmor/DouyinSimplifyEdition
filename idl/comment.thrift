namespace go comment

struct User{
    1:i64 id
    2:string name
    3:string password
}

struct Comment{
    1:i64 id
    2:User user
    3:string content
    4:string create_date
}

struct BaseResp {
    1:i64 status_code
    2:string status_message
}

struct CreateCommentReq{
    1:i64 user_id
    2:i64 video_id
    3:string content
}
struct CreateCommentResp{
    1:BaseResp base_resp
}

struct DeleteCommentReq{
    1:i64 commentId
    2:i64 video_id
}
struct DeleteCommentResp{
    1:BaseResp base_resp
}

struct GetCommentsByVideoIdReq{
    1:i64 video_id
}
struct GetCommentsByVideoIdResp{
    1:BaseResp base_resp
    2:list<Comment> comments
}

struct CountCommentReq{
    1:i64 videoId
}
struct CountCommentResp{
    1:BaseResp base_resp
    2:i64 comment_count
}

service CommentService {
    CreateCommentResp CreateComment(1:CreateCommentReq createCommentReq)
    DeleteCommentResp DeleteComment(1:DeleteCommentReq deleteCommentReq)
    GetCommentsByVideoIdResp GetCommentsByVideoId(1:GetCommentsByVideoIdReq getCommentsByVideoIdReq)
    CountCommentResp CountComment(1:CountCommentReq countCommentReq)
}