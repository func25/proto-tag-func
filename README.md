# proto-tag-func

Step 1: add inject tags that suit your need by adding a comment line, ex: // #json: "..."

```go
message Message {
  string userId = 1;
  int32 boxId = 2;
  string errorMsg = 12; // #json: "errMsgha" #bson: "errMsgho"
  int32 errorCode = 13;
  int64 timestamp = 14;
}
```

Step 2: generate .go file (use protoc compiler binary file of gg), in my example: ```protoc --go_out=. ./test/taginject/*.proto```

Step 3: run the proto-tag-func binary command, if you follow my demo: ```./proto-tag-func --input="./proto/auditproto/*"```

```go
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	BoxId     int32  `protobuf:"varint,2,opt,name=boxId,proto3" json:"boxId,omitempty"`
	ErrorMsg  string `json:"errMsgha" bson:"errMsgho" protobuf:"bytes,12,opt,name=errorMsg,proto3"` // #json: "errMsgha" #bson: "errMsgho"
	ErrorCode int32  `protobuf:"varint,13,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	Timestamp int64  `protobuf:"varint,14,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}
```
