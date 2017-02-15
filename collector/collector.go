package collector

import (
	"github.com/Sirupsen/logrus"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	PLUGIN_VENDOR  = "intel"
	PLUGIN_NAME    = "pciessd"
	PLUGIN_VERSION = 1

	partitionPath = "/proc/partitions"
	devPath       = "/dev/"
)

type PCIeSSD struct {
	logger     *logrus.Logger
	config     *Config
	name2Model map[string]Model
}

func New() plugin.Collector {
	ret := &PCIeSSD{
		logger: logrus.New(),
	}
	return ret
}

func (this *PCIeSSD) init(cfg plugin.Config) error {
	confPath, err := cfg.GetString("conf")
	if err != nil {
		logrus.Fatalf("error %s", err.Error())
		return err
	}
	this.config, err = ReadConfig(confPath)
	if err != nil {
		logrus.Fatalf("read config fail: %s", err.Error())
		return err
	}
	return nil
}

func (this *PCIeSSD) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	err := this.init(cfg)
	if err != nil {
		return nil, err
	}
	ret := []plugin.Metric{}
	metrics := this.CollectAllMetrics()
	for _, m := range metrics {
		ret = append(ret, m)
	}
	return ret, nil
}

func (this *PCIeSSD) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {

	ret := []plugin.Metric{}
	if len(mts) == 0 {
		return ret, nil
	}
	if this.config == nil {
		err := this.init(mts[0].Config)
		if err != nil {
			return nil, err
		}
	}
	metrics := this.CollectAllMetrics()
	for i := 0; i < len(mts); i++ {
		ns := mts[i].Namespace.Strings()
		name := strings.Join(ns, "/")
		m, ok := metrics[name]
		if ok {
			ret = append(ret, m)
		}
	}
	return ret, nil
}

func (this *PCIeSSD) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	return *plugin.NewConfigPolicy(), nil
}

func (this *PCIeSSD) CollectAllMetrics() map[string]plugin.Metric {

	if this.name2Model == nil {
		dev, err := this.GetDevices()
		if err != nil {
			logrus.Errorf("%s", err.Error())
			return map[string]plugin.Metric{}
		}
		this.IdentifyDevice(dev)
	}

	ret := map[string]plugin.Metric{}
	for name, model := range this.name2Model {
		f, err := os.Open(devPath + name)
		if err != nil {
			logrus.Errorf("%s", err.Error())
			continue
		}
		defer f.Close()
		for i := 0; i < len(model.GetLogPage); i++ {
			log := this.config.LogPage[model.GetLogPage[i]]
			buf := make([]byte, int(log.GetBufferLength()))
			err := GetLogPage(f.Fd(), log.Identifier, buf)
			if err != nil {
				logrus.Errorf("fail to get log page %d, err: %s", log.Identifier, err.Error())
				continue
			}

			ns := []string{PLUGIN_VENDOR, PLUGIN_NAME, name}
			m := log.Parse(buf, ns)
			for k, v := range m {
				ret[k] = v
			}
		}
	}
	return ret
}

func (this *PCIeSSD) GetDevices() ([]string, error) {

	b, err := ioutil.ReadFile(partitionPath)
	if err != nil {
		return nil, err
	}

	entrys := bytes.Split(b, []byte{'\n'})

	ret := make([]string, 0)
	for i := 0; i < len(entrys); i++ {
		var major, minor, blocks uint64
		var name string
		s := string(entrys[i])
		n, err := fmt.Sscanf(s, "%d %d %d %s", &major, &minor, &blocks, &name)
		if err != nil || n != 4 {
			continue
		}
		if major == 259 {
			ret = append(ret, name)
		}
	}
	return ret, nil
}

func (this *PCIeSSD) IdentifyDevice(devices []string) {
	this.name2Model = map[string]Model{}
	for i := 0; i < len(devices); i++ {
		dev, err := os.Open(devPath + devices[i])
		if err != nil {
			logrus.Warningf("failed to open device %s: %s", devices[i], err.Error())
			continue
		}
		defer dev.Close()
		for j := 0; j < len(this.config.Models); j++ {
			model := this.config.Models[j]
			log, ok := this.config.LogPage[model.IdentifyLog]
			if !ok {
				logrus.Errorf("failed to find log config %s", model.IdentifyLog)
				continue
			}
			buf := make([]byte, int(log.GetBufferLength()))
			err := GetLogPage(dev.Fd(), log.Identifier, buf)
			if err != nil {
				logrus.Infof("failed to identify device %s as %s: %s", devices[i], model.Name, err.Error())
				continue
			}
			metrics := log.Parse(buf, []string{})
			data, ok := metrics["MarketingName"]
			if !ok || !strings.Contains(data.Data.(string), model.Name) {
				logrus.Infof("failed to identify device %s as %s", devices[i], model.Name)
				continue
			}
			this.name2Model[devices[i]] = model
		}
	}
}
