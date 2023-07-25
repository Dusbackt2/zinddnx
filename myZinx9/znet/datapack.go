package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"myZinx9/utils"
	"myZinx9/ziface"
)

type Datapack struct {
}

func NewDataPack() *Datapack {
	return &Datapack{}
}

func (datapack *Datapack) GetHeadLen() uint32 {
	return 8
}

func (datapack *Datapack) Pack(msg ziface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

func (datapack *Datapack) Unpack(binaryData []byte) (ziface.IMessage, error) {

	dataBuff := bytes.NewReader(binaryData)

	m := &Message{}

	if err := binary.Read(dataBuff, binary.LittleEndian, &m.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuff, binary.LittleEndian, &m.Id); err != nil {
		return nil, err
	}
	if utils.GroubleObject.MaxPackageSize > 0 && m.DataLen > utils.GroubleObject.MaxPackageSize {
		return nil, errors.New("too large")
	}

	return m, nil
}
