package collector

import (
	"bytes"
	"encoding/binary"
	"syscall"
	"unsafe"
)

const (
	NVM_ADMIN_CMD       = 0xc0484e41
	OPCODE_GET_LOG_PAGE = 0x02
)

type NVMCommand struct {
	DW0     uint32
	DW1     uint32
	DW2     uint32
	DW3     uint32
	DW4     uint32
	DW5     uint32
	DW6     uint32
	DW7     uint32
	DW8     uint32
	DW9     uint32
	DW10    uint32
	DW11    uint32
	DW12    uint32
	DW13    uint32
	DW14    uint32
	DW15    uint32
	Timeout uint32
	Result  uint32
}

func (this *NVMCommand) SetOpcode(opc uint8) {
	this.DW0 = (this.DW0 & 0xFFFFFF00) | (uint32(opc))
}

func (this *NVMCommand) SetNamespaceIdentifier(nsid uint32) {
	this.DW1 = nsid
}

func (this *NVMCommand) SetPRPEntry1(prp1 uint64) {
	this.DW6 = uint32(prp1)
	this.DW7 = uint32(prp1 >> 32)
}
func (this *NVMCommand) SetPRPEntry2(prp2 uint64) {
	this.DW8 = uint32(prp2)
	this.DW9 = uint32(prp2 >> 32)
}

type NVMeAdminCommand struct {
	NVMCommand
}

func (this *NVMeAdminCommand) SetCommandFlags(cflgs uint8) {
	this.DW0 = (this.DW0 & 0xFFFF00FF) | (uint32(cflgs) << 8)
}

func (this *NVMeAdminCommand) SetControllerID(ctlid uint16) {
	this.DW0 = (this.DW0 & 0x0000FFFF) | (uint32(ctlid) << 16)
}

func (this *NVMeAdminCommand) SetDataLength(dlen uint32) {
	this.DW9 = dlen
}

func GetLogPage(fd uintptr, logIdentifier uint8, buffer []byte) error {

	bufPtr := uint64(uintptr(unsafe.Pointer(&buffer[0])))

	cmd := NVMeAdminCommand{}
	cmd.SetOpcode(OPCODE_GET_LOG_PAGE)
	cmd.SetNamespaceIdentifier(0xFFFFFFFF)
	cmd.SetPRPEntry1(bufPtr)
	cmd.SetDataLength(uint32(len(buffer)))
	cmd.DW10 = (uint32((len(buffer)>>2)-1) << 16) | uint32(logIdentifier)

	cmdBuf := new(bytes.Buffer)
	err := binary.Write(cmdBuf, binary.LittleEndian, cmd)
	if err != nil {
		return err
	}

	cmdPtr := uintptr(unsafe.Pointer(&(cmdBuf.Bytes()[0])))
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(NVM_ADMIN_CMD), cmdPtr)
	if e != 0 {
		return e
	} else {
		return nil
	}
}
