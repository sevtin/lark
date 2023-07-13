package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_msg"
	"unsafe"
)

func Int32ToBytes(n int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

func Int64ToBytes(i int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt32(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func BufferDecode(buf []byte, obj interface{}) (err error) {
	var (
		buffer *bytes.Buffer
		dc     *gob.Decoder
	)
	buffer = bytes.NewBuffer(buf)
	dc = gob.NewDecoder(buffer)
	err = dc.Decode(obj)
	return
}

func ObjEncode(obj interface{}) (buf []byte, err error) {
	var (
		buffer  bytes.Buffer
		encoder *gob.Encoder
	)
	encoder = gob.NewEncoder(&buffer)
	err = encoder.Encode(obj)
	if err != nil {
		return
	}
	buf = buffer.Bytes()
	return
}

// Encode 将消息编码【性能更优】
func Encode(topic int32, subtopic int32, msgType int32, message []byte) (buf []byte, err error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息长度
	err = binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return
	}
	// 写入topic
	err = binary.Write(pkg, binary.LittleEndian, topic)
	if err != nil {
		return
	}
	// 写入subtopic
	err = binary.Write(pkg, binary.LittleEndian, subtopic)
	if err != nil {
		return
	}
	// 写入msgType
	err = binary.Write(pkg, binary.LittleEndian, msgType)
	if err != nil {
		return
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, message)
	if err != nil {
		return
	}
	buf = pkg.Bytes()
	return
}

// Encode 将消息编码【性能更优】
func EncodeCliMessage(topic int32, subtopic int32, msgId int64, message []byte) (buf []byte, err error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息长度
	err = binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return
	}
	// 写入topic
	err = binary.Write(pkg, binary.LittleEndian, topic)
	if err != nil {
		return
	}
	// 写入subtopic
	err = binary.Write(pkg, binary.LittleEndian, subtopic)
	if err != nil {
		return
	}
	// 写入msgId
	err = binary.Write(pkg, binary.LittleEndian, msgId)
	if err != nil {
		return
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, message)
	if err != nil {
		return
	}
	buf = pkg.Bytes()
	return
}

/*
const (

	MessageLength   uint32 = 4
	MessageTopic    uint32 = 8
	MessageSubtopic uint32 = 12
	MessageType     uint32 = 16
	MessageId       uint32 = 24

)

	func Decode(buf []byte) (msg *pb_msg.Packet, endNode uint32) {
		msg = new(pb_msg.Packet)
		var (
			totalLength  uint32
			lengthBuff   []byte
			length       uint32
			topicBuff    []byte
			topic        uint32
			subtopicBuff []byte
			subtopic     uint32
			msgTypeBuff  []byte
			msgType      uint32
			body         []byte
			msgId        uint64
			msgIdBuff    []byte
		)

		totalLength = uint32(len(buf))
		if totalLength < MessageId {
			return
		}
		lengthBuff = buf[:MessageLength]
		length = binary.LittleEndian.Uint32(lengthBuff)
		endNode = MessageId + length
		if totalLength < endNode {
			return
		}
		topicBuff = buf[MessageLength:MessageTopic]
		topic = binary.LittleEndian.Uint32(topicBuff)

		subtopicBuff = buf[MessageTopic:MessageSubtopic]
		subtopic = binary.LittleEndian.Uint32(subtopicBuff)

		msgTypeBuff = buf[MessageSubtopic:MessageType]
		msgType = binary.LittleEndian.Uint32(msgTypeBuff)

		msgIdBuff = buf[MessageType:MessageId]
		msgId = binary.LittleEndian.Uint64(msgIdBuff)

		body = buf[MessageId:endNode]
		msg.Topic = pb_enum.TOPIC(topic)
		msg.SubTopic = pb_enum.SUB_TOPIC(subtopic)
		msg.MsgType = pb_enum.MESSAGE_TYPE(msgType)
		msg.MsgId = int64(msgId)
		msg.Data = body
		return
	}
*/
const (
	CliMessageHeaderLen = 20 // len(4) + topic(4) + subtopic(4) + msgId(8)
	SrvMessageHeaderLen = 16 // len(4) + topic(4) + subtopic(4) + msgType(4)
)

func Decode(pkg *bytes.Buffer, isSrvMsg bool) (msg *pb_msg.Packet) {
	msg = new(pb_msg.Packet)
	var (
		totalLength uint32
		endNode     uint32
		length      uint32
		topic       uint32
		subtopic    uint32
		msgType     uint32
		body        []byte
		msgId       uint64
		err         error
	)
	totalLength = uint32(pkg.Len())
	if isSrvMsg == true {
		if totalLength < SrvMessageHeaderLen {
			return
		}
	} else {
		if totalLength < CliMessageHeaderLen {
			return
		}
	}
	// 读取消息长度
	err = binary.Read(pkg, binary.LittleEndian, &length)
	if err != nil {
		return
	}
	if isSrvMsg == true {
		endNode = SrvMessageHeaderLen + length
	} else {
		endNode = CliMessageHeaderLen + length
	}
	if totalLength < endNode {
		return
	}
	// 读取topic
	err = binary.Read(pkg, binary.LittleEndian, &topic)
	if err != nil {
		return
	}
	// 读取subtopic
	err = binary.Read(pkg, binary.LittleEndian, &subtopic)
	if err != nil {
		return
	}
	if isSrvMsg == true {
		// 读取msgType
		err = binary.Read(pkg, binary.LittleEndian, &msgType)
		if err != nil {
			return
		}
	} else {
		// 读取msgId
		err = binary.Read(pkg, binary.LittleEndian, &msgId)
		if err != nil {
			return
		}
	}
	// 读取body
	body = make([]byte, length)
	err = binary.Read(pkg, binary.LittleEndian, &body)
	if err != nil {
		return
	}
	msg.Topic = pb_enum.TOPIC(topic)
	msg.SubTopic = pb_enum.SUB_TOPIC(subtopic)
	msg.MsgType = pb_enum.MESSAGE_TYPE(msgType)
	msg.MsgId = int64(msgId)
	msg.Data = body
	return
}
