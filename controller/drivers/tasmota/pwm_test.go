package tasmota

import (
	"github.com/reef-pi/hal"
	"testing"
)

func TestMagicHome(t *testing.T) {

	params := map[string]interface{}{
		"Ip Address": "",
	}

	f := PwmDriverFactory()
	d, err := f.NewDriver(params, nil)

	if err != nil {
		t.Fatal(err)
	}

	meta := d.Metadata()
	if len(meta.Capabilities) != 1 {
		t.Error("Expected 1 capabilities, found:", len(meta.Capabilities))
	}

	dig, ok := d.(hal.PWMDriver)
	if !ok {
		t.Error("Failed to type driver to PWM driver")
	}

	if len(dig.PWMChannels()) != 1 {
		t.Error("Expected a single pwm channel, found:", len(dig.PWMChannels()))
	}

	pin, err := dig.PWMChannel(0)
	if err != nil {
		t.Error(err)
	}

	if pin.LastState() != false {
		t.Error("Expected initial state to be false")
	}

	if err := pin.Set(50.0); err != nil {
		t.Error(err)
	}

}
