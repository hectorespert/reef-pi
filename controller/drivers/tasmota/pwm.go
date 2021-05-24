package tasmota

import (
    "errors"
    "fmt"
    "github.com/reef-pi/hal"
    "io/ioutil"
    "log"
    "net/http"
    "sync"
    "time"
)

type pwmDriver struct {
	meta    hal.Metadata
	address string
}

func (m *pwmDriver) Close() error {
	return nil
}

func (m *pwmDriver) Metadata() hal.Metadata {
	return m.meta
}

func (m *pwmDriver) Name() string {
	return "Pwm"
}

func (m *pwmDriver) Number() int {
	return 0
}

func (m *pwmDriver) Pins(capability hal.Capability) ([]hal.Pin, error) {
	switch capability {
	case hal.PWM:
		return []hal.Pin{m}, nil
	default:
		return nil, fmt.Errorf("unsupported capability:%s", capability.String())
	}
}

func (m *pwmDriver) PWMChannels() []hal.PWMChannel {
	return []hal.PWMChannel{m}
}

func (m *pwmDriver) PWMChannel(_ int) (hal.PWMChannel, error) {
	return m, nil
}

func (m *pwmDriver) LastState() bool {
	return false
}

func (m *pwmDriver) Set(value float64) error {
    uri := fmt.Sprintf("http://%s/cm?cmnd=Dimmer%%20%.0f", m.address, value)
    req, err := http.NewRequest("GET", uri, nil)
    if err != nil {
        return err
    }
    c := http.Client{
        Timeout: 5 * time.Second,
    }
    resp, err := c.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    msg, _ := ioutil.ReadAll(resp.Body)
    log.Println("Tasmota: URI: ", req.URL.String(), " Http Code: ", resp.StatusCode, "Channel:", string(msg))
    if resp.StatusCode == 200 {
        return nil
    }
    return fmt.Errorf("HTTP Code:%d. Body:%v", resp.StatusCode, string(msg))
}

func (m *pwmDriver) Write(b bool) error {
	return nil
}

func (m *pwmDriver) DigitalOutputPins() []hal.DigitalOutputPin {
	return []hal.DigitalOutputPin{m}
}

func (m *pwmDriver) DigitalOutputPin(_ int) (hal.DigitalOutputPin, error) {
	return m, nil
}

type factory struct {
	meta       hal.Metadata
	parameters []hal.ConfigParameter
}

var pwmDriverFactory *factory
var once sync.Once

func PwmDriverFactory() hal.DriverFactory {

	once.Do(func() {
		pwmDriverFactory = &factory{
			meta: hal.Metadata{
				Name:         "PWM One Channel",
				Description:  "Tasmota PWM Controller One Channel",
				Capabilities: []hal.Capability{hal.PWM},
			},
			parameters: []hal.ConfigParameter{
				{
					Name:    "Domain or Address",
					Type:    hal.String,
					Order:   0,
					Default: "192.1.168.4",
				},
			},
		}
	})

	return pwmDriverFactory
}

func (f *factory) GetParameters() []hal.ConfigParameter {
	return f.parameters
}

func (f *factory) ValidateParameters(parameters map[string]interface{}) (bool, map[string][]string) {
	var failures = make(map[string][]string)
	return true, failures
}

func (f *factory) Metadata() hal.Metadata {
	return f.meta
}

func (f *factory) NewDriver(parameters map[string]interface{}, hardwareResources interface{}) (hal.Driver, error) {
	if valid, failures := f.ValidateParameters(parameters); !valid {
		return nil, errors.New(hal.ToErrorString(failures))
	}

	driver := &pwmDriver{
		meta: f.meta,
        address: parameters["Domain Or Address"].(string),
	}
	return driver, nil
}

