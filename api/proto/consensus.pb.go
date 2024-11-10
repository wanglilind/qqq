// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: api/proto/consensus.proto

package consensus

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId         string         `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`                           // 区块ID
	PreviousBlockId string         `protobuf:"bytes,2,opt,name=previous_block_id,json=previousBlockId,proto3" json:"previous_block_id,omitempty"` // 前一个区块ID
	Timestamp       int64          `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                                     // 时间戳
	ProposerId      string         `protobuf:"bytes,4,opt,name=proposer_id,json=proposerId,proto3" json:"proposer_id,omitempty"`                  // 提议者ID
	Transactions    []*Transaction `protobuf:"bytes,5,rep,name=transactions,proto3" json:"transactions,omitempty"`                                // 交易列表
	MerkleRoot      []byte         `protobuf:"bytes,6,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`                  // Merkle树根
	Signature       []byte         `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`                                      // 区块签名
	Height          int64          `protobuf:"varint,8,opt,name=height,proto3" json:"height,omitempty"`                                           // 区块高度
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_consensus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_consensus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_api_proto_consensus_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *Block) GetPreviousBlockId() string {
	if x != nil {
		return x.PreviousBlockId
	}
	return ""
}

func (x *Block) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Block) GetProposerId() string {
	if x != nil {
		return x.ProposerId
	}
	return ""
}

func (x *Block) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *Block) GetMerkleRoot() []byte {
	if x != nil {
		return x.MerkleRoot
	}
	return nil
}

func (x *Block) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Block) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"` // 交易ID
	SenderId      string `protobuf:"bytes,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`                // 发送者ID
	RecipientId   string `protobuf:"bytes,3,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`       // 接收者ID
	Amount        uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`                                   // 金额
	Timestamp     int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                             // 时间戳
	Signature     []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`                              // 交易签名
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_consensus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_consensus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_api_proto_consensus_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *Transaction) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *Transaction) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

func (x *Transaction) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Transaction) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ProposeBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposerId   string         `protobuf:"bytes,1,opt,name=proposer_id,json=proposerId,proto3" json:"proposer_id,omitempty"` // 提议者ID
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`               // 待打包交易
	Timestamp    int64          `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                    // 提议时间戳
}

func (x *ProposeBlockRequest) Reset() {
	*x = ProposeBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_consensus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposeBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposeBlockRequest) ProtoMessage() {}

func (x *ProposeBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_consensus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposeBlockRequest.ProtoReflect.Descriptor instead.
func (*ProposeBlockRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_consensus_proto_rawDescGZIP(), []int{2}
}

func (x *ProposeBlockRequest) GetProposerId() string {
	if x != nil {
		return x.ProposerId
	}
	return ""
}

func (x *ProposeBlockRequest) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *ProposeBlockRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ProposeBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId     string `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`             // 区块ID
	ConsensusId string `protobuf:"bytes,2,opt,name=consensus_id,json=consensusId,proto3" json:"consensus_id,omitempty"` // 共识ID
	Status      string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`                              // 状态
	Message     string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`                            // 响应信息
}

func (x *ProposeBlockResponse) Reset() {
	*x = ProposeBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_consensus_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposeBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposeBlockResponse) ProtoMessage() {}

func (x *ProposeBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_consensus_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposeBlockResponse.ProtoReflect.Descriptor instead.
func (*ProposeBlockResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_consensus_proto_rawDescGZIP(), []int{3}
}

func (x *ProposeBlockResponse) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *ProposeBlockResponse) GetConsensusId() string {
	if x != nil {
		return x.ConsensusId
	}
	return ""
}

func (x *ProposeBlockResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ProposeBlockResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetConsensusStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"` // 节点ID
}

func (x *GetConsensusStatusRequest) Reset() {
	*x = GetConsensusStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_consensus_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsensusStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsensusStatusRequest) ProtoMessage() {}

func (x *GetConsensusStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_consensus_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsensusStatusRequest.ProtoReflect.Descriptor instead.
func (*GetConsensusStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_consensus_proto_rawDescGZIP(), []int{4}
}

func (x *GetConsensusStatusRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type GetConsensusStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentRound   int64  `protobuf:"varint,1,opt,name=current_round,json=currentRound,proto3" json:"current_round,omitempty"`         // 当前轮次
	ActiveNodes    int32  `protobuf:"varint,2,opt,name=active_nodes,json=activeNodes,proto3" json:"active_nodes,omitempty"`            // 活跃节点数
	ConsensusState string `protobuf:"bytes,3,opt,name=consensus_state,json=consensusState,proto3" json:"consensus_state,omitempty"`    // 共识状态
	LastBlockId    string `protobuf:"bytes,4,opt,name=last_block_id,json=lastBlockId,proto3" json:"last_block_id,omitempty"`           // 最新区块ID
	LastUpdateTime int64  `protobuf:"varint,5,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"` // 最后更新时间
}

func (x *GetConsensusStatusResponse) Reset() {
	*x = GetConsensusStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_consensus_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsensusStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsensusStatusResponse) ProtoMessage() {}

func (x *GetConsensusStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_consensus_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsensusStatusResponse.ProtoReflect.Descriptor instead.
func (*GetConsensusStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_consensus_proto_rawDescGZIP(), []int{5}
}

func (x *GetConsensusStatusResponse) GetCurrentRound() int64 {
	if x != nil {
		return x.CurrentRound
	}
	return 0
}

func (x *GetConsensusStatusResponse) GetActiveNodes() int32 {
	if x != nil {
		return x.ActiveNodes
	}
	return 0
}

func (x *GetConsensusStatusResponse) GetConsensusState() string {
	if x != nil {
		return x.ConsensusState
	}
	return ""
}

func (x *GetConsensusStatusResponse) GetLastBlockId() string {
	if x != nil {
		return x.LastBlockId
	}
	return ""
}

func (x *GetConsensusStatusResponse) GetLastUpdateTime() int64 {
	if x != nil {
		return x.LastUpdateTime
	}
	return 0
}

type RegisterNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId    string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`          // 节点ID
	Endpoint  string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`                    // 节点地址
	PublicKey string `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"` // 公钥
	NodeType  string `protobuf:"bytes,4,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`    // 节点类型
}

func (x *RegisterNodeRequest) Reset() {
	*x = RegisterNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_consensus_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNodeRequest) ProtoMessage() {}

func (x *RegisterNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_consensus_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNodeRequest.ProtoReflect.Descriptor instead.
func (*RegisterNodeRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_consensus_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterNodeRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *RegisterNodeRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *RegisterNodeRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *RegisterNodeRequest) GetNodeType() string {
	if x != nil {
		return x.NodeType
	}
	return ""
}

type RegisterNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success       bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                                 // 注册结果
	Message       string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                                  // 响应信息
	PeerEndpoints []string `protobuf:"bytes,3,rep,name=peer_endpoints,json=peerEndpoints,proto3" json:"peer_endpoints,omitempty"` // 对等节点列表
}

func (x *RegisterNodeResponse) Reset() {
	*x = RegisterNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_consensus_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNodeResponse) ProtoMessage() {}

func (x *RegisterNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_consensus_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNodeResponse.ProtoReflect.Descriptor instead.
func (*RegisterNodeResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_consensus_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterNodeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RegisterNodeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RegisterNodeResponse) GetPeerEndpoints() []string {
	if x != nil {
		return x.PeerEndpoints
	}
	return nil
}

type SyncBlocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId     string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`              // 节点ID
	FromHeight int64  `protobuf:"varint,2,opt,name=from_height,json=fromHeight,proto3" json:"from_height,omitempty"` // 起始高度
	ToHeight   int64  `protobuf:"varint,3,opt,name=to_height,json=toHeight,proto3" json:"to_height,omitempty"`       // 结束高度
}

func (x *SyncBlocksRequest) Reset() {
	*x = SyncBlocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_consensus_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncBlocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncBlocksRequest) ProtoMessage() {}

func (x *SyncBlocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_consensus_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncBlocksRequest.ProtoReflect.Descriptor instead.
func (*SyncBlocksRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_consensus_proto_rawDescGZIP(), []int{8}
}

func (x *SyncBlocksRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *SyncBlocksRequest) GetFromHeight() int64 {
	if x != nil {
		return x.FromHeight
	}
	return 0
}

func (x *SyncBlocksRequest) GetToHeight() int64 {
	if x != nil {
		return x.ToHeight
	}
	return 0
}

var File_api_proto_consensus_proto protoreflect.FileDescriptor

var file_api_proto_consensus_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x22, 0xa0, 0x02, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xc8, 0x01, 0x0a, 0x0b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3a, 0x0a,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x34, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0xdb, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x27, 0x0a,
	0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x71, 0x0a,
	0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x65, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x22, 0x6a, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0xd7, 0x02, 0x0a,
	0x10, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x61, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x30, 0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x6e, 0x67, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x64, 0x2f,
	0x71, 0x71, 0x71, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_consensus_proto_rawDescOnce sync.Once
	file_api_proto_consensus_proto_rawDescData = file_api_proto_consensus_proto_rawDesc
)

func file_api_proto_consensus_proto_rawDescGZIP() []byte {
	file_api_proto_consensus_proto_rawDescOnce.Do(func() {
		file_api_proto_consensus_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_consensus_proto_rawDescData)
	})
	return file_api_proto_consensus_proto_rawDescData
}

var file_api_proto_consensus_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_proto_consensus_proto_goTypes = []interface{}{
	(*Block)(nil),                      // 0: consensus.Block
	(*Transaction)(nil),                // 1: consensus.Transaction
	(*ProposeBlockRequest)(nil),        // 2: consensus.ProposeBlockRequest
	(*ProposeBlockResponse)(nil),       // 3: consensus.ProposeBlockResponse
	(*GetConsensusStatusRequest)(nil),  // 4: consensus.GetConsensusStatusRequest
	(*GetConsensusStatusResponse)(nil), // 5: consensus.GetConsensusStatusResponse
	(*RegisterNodeRequest)(nil),        // 6: consensus.RegisterNodeRequest
	(*RegisterNodeResponse)(nil),       // 7: consensus.RegisterNodeResponse
	(*SyncBlocksRequest)(nil),          // 8: consensus.SyncBlocksRequest
}
var file_api_proto_consensus_proto_depIdxs = []int32{
	1, // 0: consensus.Block.transactions:type_name -> consensus.Transaction
	1, // 1: consensus.ProposeBlockRequest.transactions:type_name -> consensus.Transaction
	2, // 2: consensus.ConsensusService.ProposeBlock:input_type -> consensus.ProposeBlockRequest
	4, // 3: consensus.ConsensusService.GetConsensusStatus:input_type -> consensus.GetConsensusStatusRequest
	6, // 4: consensus.ConsensusService.RegisterNode:input_type -> consensus.RegisterNodeRequest
	8, // 5: consensus.ConsensusService.SyncBlocks:input_type -> consensus.SyncBlocksRequest
	3, // 6: consensus.ConsensusService.ProposeBlock:output_type -> consensus.ProposeBlockResponse
	5, // 7: consensus.ConsensusService.GetConsensusStatus:output_type -> consensus.GetConsensusStatusResponse
	7, // 8: consensus.ConsensusService.RegisterNode:output_type -> consensus.RegisterNodeResponse
	0, // 9: consensus.ConsensusService.SyncBlocks:output_type -> consensus.Block
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_consensus_proto_init() }
func file_api_proto_consensus_proto_init() {
	if File_api_proto_consensus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_consensus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_consensus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_consensus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposeBlockRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_consensus_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposeBlockResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_consensus_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsensusStatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_consensus_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsensusStatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_consensus_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterNodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_consensus_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterNodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_consensus_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncBlocksRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_consensus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_consensus_proto_goTypes,
		DependencyIndexes: file_api_proto_consensus_proto_depIdxs,
		MessageInfos:      file_api_proto_consensus_proto_msgTypes,
	}.Build()
	File_api_proto_consensus_proto = out.File
	file_api_proto_consensus_proto_rawDesc = nil
	file_api_proto_consensus_proto_goTypes = nil
	file_api_proto_consensus_proto_depIdxs = nil
}
