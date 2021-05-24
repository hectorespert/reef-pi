package tasmota

import (
    "errors"
    "fmt"
    "github.com/reef-pi/hal"
    "sync"
)


type outputDriver struct {
    meta    hal.Metadata
    address string
}

func (m *outputDriver) Close() error {
    return nil
}

func (m *outputDriver) Metadata() hal.Metadata {
    return m.meta
}

func (m *outputDriver) Pins(capability hal.Capability) ([]hal.Pin, error) {
    switch capability {
    case hal.DigitalOutput:
        return []hal.Pin{m}, nil
    default:
        return nil, fmt.Errorf("unsupported capability:%s", capability.String())
    }
}

func (m *outputDriver) Name() string {
    return "DigitalOutput"
}

func (m *outputDriver) Number() int {
    return 0
}

type outputFactory struct {
    meta       hal.Metadata
    parameters []hal.ConfigParameter
}

var outputDriverFactory *outputFactory
var outputDriverFactoryOnce sync.Once

func OutputDriverFactory() hal.DriverFactory {
    outputDriverFactoryOnce.Do(func() {
        outputDriverFactory = &outputFactory{
            meta: hal.Metadata{
                Name:         "Digital Output Channel",
                Description:  "Tasmota Digital Output Controller One Channel",
                Capabilities: []hal.Capability{hal.DigitalOutput},
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

    return outputDriverFactory
}

func (f *outputFactory) GetParameters() []hal.ConfigParameter {
    return f.parameters
}

func (f *outputFactory) ValidateParameters(parameters map[string]interface{}) (bool, map[string][]string) {
    var failures = make(map[string][]string)
    return true, failures
}

func (f *outputFactory) Metadata() hal.Metadata {
    return f.meta
}

func (f *outputFactory) NewDriver(parameters map[string]interface{}, hardwareResources interface{}) (hal.Driver, error) {
    if valid, failures := f.ValidateParameters(parameters); !valid {
        return nil, errors.New(hal.ToErrorString(failures))
    }

    driver := &outputDriver{
        meta: f.meta,
        address: parameters["Domain Or Address"].(string),
    }
    return driver, nil
}
