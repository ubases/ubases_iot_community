package gorushclient

//import (
//	"context"
//
//	"github.com/appleboy/gorush/rpc/proto"
//	//"github.com/golang/protobuf/proto"
//	"google.golang.org/grpc"
//)
//
//// Client 是gorush客户端的封装结构体
//type Client struct {
//	conn   *grpc.ClientConn
//	client proto.GorushClient
//}
//
//// NewClient 创建一个新的gorush客户端实例
//func NewClient(address string) (*Client, error) {
//	conn, err := grpc.Dial(address, grpc.WithInsecure())
//	if err != nil {
//		return nil, err
//	}
//
//	client := proto.NewGorushClient(conn)
//
//	return &Client{
//		conn:   conn,
//		client: client,
//	}, nil
//}
//
//// PushNotification 发送推送通知
//func (c *Client) PushNotification(tokens []string, platform proto.PlatformType, message string) error {
//	request := &proto.PushNotificationRequest{
//		Tokens:   tokens,
//		Platform: platform,
//		Message:  message,
//	}
//
//	_, err := c.client.PushNotification(context.Background(), request)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// Close 关闭gorush客户端连接
//func (c *Client) Close() {
//	c.conn.Close()
//}
