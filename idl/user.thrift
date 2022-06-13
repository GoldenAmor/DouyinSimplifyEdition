namespace go user

struct User{
    1:i64 id
    2:string name
    3:string password
}

struct BaseResp {
    1:i64 status_code
    2:string status_message
}

struct ContainsNameReq{
    1:string username
}
struct ContainsNameResp{
    1:BaseResp base_resp
    2:bool contains_name
}

struct CreateUserReq{
    1:string username
    2:string password
}
struct CreateUserResp{
    1:BaseResp base_resp
    2:i64 user_id
}

struct GetUserByNameReq{
    1:string username
}
struct GetUserByNameResp{
    1:BaseResp base_resp
    2:User user
}

struct GetUserByIdReq{
    1:i64 id
}
struct GetUserByIdResp{
    1:BaseResp base_resp
    2:User user
}

service UserService {
    ContainsNameResp ContainsName(1:ContainsNameReq containsNameReq)
    CreateUserResp CreateUser(1:CreateUserReq createUserReq)
    GetUserByNameResp GetUserByName(1:GetUserByNameReq getUserByNameReq)
    GetUserByIdResp GetUserById(1:GetUserByIdReq getUserByIdReq)
}