namespace go video

struct BaseResp {
    1:i64 status_code
    2:string status_message
}

struct User{
    1:i64 id
    2:string name
}

struct Video{
	1:i64 id
   	2:i64 Author_id
	3:User Author
	4:string PlayUrl
	5:string CoverUrl
}

struct CreateVideoReq{
    1:i64 user_id
    2:string play_url
    3:string cover_url
}
struct CreateVideoResp{
    1:BaseResp base_resp
}

struct GetPublishListReq{
    1:i64 user_id
}
struct GetPublishListResp{
    1:BaseResp base_resp
    2:list<Video> videos
}

struct GetVideosReq {
    1:i64 latest_time
}
struct GetVideosResp {
    1:BaseResp base_resp
    2:list<Video> videos
    3:i64 nextTime
}

struct GetVideoByIdReq {
    1:i64 id
}
struct GetVideoByIdResp {
    1:BaseResp base_resp
    2:Video video
}

struct PublishReq{
    1:i64 user_id
    2:string filename
    3:i64 size
    4:list<byte> data
}
struct PublishResp{
    1:BaseResp base_resp
    2:string url
}

service VideoService {
    CreateVideoResp CreateVideo(1:CreateVideoReq createVideoReq)
    GetPublishListResp GetPublishList(1:GetPublishListReq getPublishListReq)
    GetVideosResp GetVideos(1:GetVideosReq getVideosReq)
    GetVideoByIdResp GetVideoById(1:GetVideoByIdReq getVideoByIdReq)
    PublishResp Publish(1:PublishReq publishReq)
}
