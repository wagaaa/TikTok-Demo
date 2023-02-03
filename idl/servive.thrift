include "/idl/model.thrift"
namespace go tiktok.demo

typedef map<string, string> Data

struct Response {
    1:required i32 statusCode; //状态码
    2:required string statusMsg; //状态信息
    3:required Data data;
}

service Greeter {
    Response SayHello(
        1:required User.User user
    )

    Response GetUser(
        1:required i32 uid
    )
}