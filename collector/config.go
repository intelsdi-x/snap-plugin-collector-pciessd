package collector

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Config struct {
	LogPage map[string]LogPage
	Models  []Model
}

func ReadConfig(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	ret := &Config{}
	err = json.Unmarshal(b, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

type Model struct {
	IdentifyLog string
	Name        string
	GetLogPage  []string
}

type LogPage struct {
	Identifier uint8
	Metrics    []Metric
}

func (this LogPage) Parse(data []byte, ns []string) map[string]plugin.Metric {
	now := time.Now()
	ret := make(map[string]plugin.Metric)
	for i := 0; i < len(this.Metrics); i++ {
		m := this.Metrics[i]
		if int(m.Length) > len(data) {
			return ret
		}
		var key []string
		var value interface{}
		rep := this.Metrics[i].Repeat
		if rep == nil {
			key, value = this.Metrics[i].GetKeyValue(data[:m.Length])
			if key != nil {
				namespace := append(ns, key...)
				ret[strings.Join(namespace, "/")] = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace...),
					Version:   PLUGIN_VERSION,
					Data:      value,
					Timestamp: now,
					Unit:      m.Unit,
				}
			}
			data = data[m.Length:]
		} else {
			for j, k := rep.Start, uint32(0); k < rep.Total; j, k = j+rep.Step, k+1 {
				if int(m.Length) > len(data) {
					return ret
				}
				key, value = this.Metrics[i].GetKeyValue(data[:m.Length])
				if key != nil {
					namespace := append(ns, key...)
					namespace = append(namespace, fmt.Sprintf(rep.NameFormat, j))
					ret[strings.Join(namespace, "/")] = plugin.Metric{
						Namespace: plugin.NewNamespace(namespace...),
						Version:   PLUGIN_VERSION,
						Data:      value,
						Timestamp: now,
						Unit:      m.Unit,
					}
				}
				data = data[m.Length:]
			}
		}
	}
	return ret
}

func (this LogPage) GetBufferLength() uint32 {
	var totalLength uint32 = 0
	for i := 0; i < len(this.Metrics); i++ {
		rep := this.Metrics[i].Repeat
		if rep == nil {
			totalLength += uint32(this.Metrics[i].Length)
		} else {
			totalLength += (uint32(this.Metrics[i].Length) * rep.Total)
		}
	}
	return totalLength + 50
}

type Formatter interface {
	Format(data []byte) interface{}
}

type Metric struct {
	Name     string
	Length   uint16
	Unit     string
	Ignore   bool
	*Repeat  `json:"Repeat"`
	*Bit     `json:"Bit"`
	*Uint8   `json:"Uint8"`
	*Uint16  `json:"Uint16"`
	*Uint32  `json:"Uint32"`
	*Uint64  `json:"Uint64"`
	*Uint128 `json:"Uint128"`
	*Int64   `json:"Int64"`
	*String  `json:"String"`
}

func (this Metric) GetKeyValue(data []byte) ([]string, interface{}) {
	var f Formatter
	if this.Ignore {
		return nil, nil
	} else if this.Bit != nil {
		f = this.Bit
	} else if this.Uint8 != nil {
		f = this.Uint8
	} else if this.Uint16 != nil {
		f = this.Uint16
	} else if this.Uint32 != nil {
		f = this.Uint32
	} else if this.Uint64 != nil {
		f = this.Uint64
	} else if this.Uint128 != nil {
		f = this.Uint128
	} else if this.Int64 != nil {
		f = this.Int64
	} else if this.String != nil {
		f = this.String
	} else {
		return nil, nil
	}
	return strings.Split(this.Name, "/"), f.Format(data)
}

type Repeat struct {
	NameFormat string
	Start      uint32
	Step       uint32
	Total      uint32
}

type Bit struct {
	Set []string
}

func (this *Bit) Format(data []byte) interface{} {
	ret := ""
	var i uint
	for i = 0; i < uint(len(this.Set)); i++ {
		if data[i/8]&(1<<(i%8)) != 0 {
			ret += this.Set[i]
		}
	}
	return ret
}

func fixDataLength(data []byte, length uint16) []byte {
	dataLen := uint16(len(data))
	if data == nil {
		return make([]byte, length)
	} else if dataLen < length {
		ret := make([]byte, length)
		copy(ret, data)
		return ret
	} else if dataLen > length {
		return data[:length]
	} else {
		return data
	}
}

type Uint8 struct {
	Scale uint32
}

func (this *Uint8) Format(data []byte) interface{} {
	var tmp uint8

	data = fixDataLength(data, 1)
	binary.Read(bytes.NewReader(data), binary.LittleEndian, &tmp)
	if this.Scale == 0 {
		return fmt.Sprintf("%d", tmp)
	} else {
		return fmt.Sprintf("%d", uint64(tmp)*uint64(this.Scale))
	}
}

type Uint16 struct {
	Scale uint32
}

func (this *Uint16) Format(data []byte) interface{} {
	var tmp uint16

	data = fixDataLength(data, 2)
	binary.Read(bytes.NewReader(data), binary.LittleEndian, &tmp)
	if this.Scale == 0 {
		return fmt.Sprintf("%d", tmp)
	} else {
		return fmt.Sprintf("%d", uint64(tmp)*uint64(this.Scale))
	}
}

type Uint32 struct {
	Scale uint32
}

func (this *Uint32) Format(data []byte) interface{} {
	var tmp uint32

	data = fixDataLength(data, 4)
	binary.Read(bytes.NewReader(data), binary.LittleEndian, &tmp)
	if this.Scale == 0 {
		return fmt.Sprintf("%d", tmp)
	} else {
		return fmt.Sprintf("%d", uint64(tmp)*uint64(this.Scale))
	}
}

type Uint64 struct {
	Scale uint32
}

func (this *Uint64) Format(data []byte) interface{} {
	var tmp uint64

	data = fixDataLength(data, 8)
	binary.Read(bytes.NewReader(data), binary.LittleEndian, &tmp)
	if this.Scale == 0 {
		return fmt.Sprintf("%d", tmp)
	} else {
		return fmt.Sprintf("%.0f", float64(tmp)*float64(this.Scale))
	}
}

type Uint128 struct {
	Scale uint32
}

func (this *Uint128) Format(data []byte) interface{} {
	var tmpLow, tmpHigh uint64

	data = fixDataLength(data, 16)
	binary.Read(bytes.NewReader(data[:8]), binary.LittleEndian, &tmpLow)
	binary.Read(bytes.NewReader(data[8:]), binary.LittleEndian, &tmpHigh)
	value := float64(tmpHigh)*float64(1<<64) + float64(tmpLow)

	if this.Scale == 0 {
		return fmt.Sprintf("%.0f", value)
	} else {
		return fmt.Sprintf("%.0f", value*float64(this.Scale))
	}
}

type Int64 struct {
	Scale uint32
}

func (this *Int64) Format(data []byte) interface{} {
	var tmp int64

	data = fixDataLength(data, 8)
	binary.Read(bytes.NewReader(data), binary.LittleEndian, &tmp)
	if this.Scale == 0 {
		return fmt.Sprintf("%d", tmp)
	} else {
		return fmt.Sprintf("%.0f", float64(tmp)*float64(this.Scale))
	}
}

type String struct {
}

func (this *String) Format(data []byte) interface{} {
	return string(data)
}
