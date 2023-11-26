package yamlreader

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
	"time"
)

func TestLoadReadConf_CorrectYAML(t *testing.T) {
	c := &reader{}

	f, err := os.CreateTemp("", "tempfile.*.yaml")
	assert.NoError(t, err, "cannot create file")
	defer os.Remove(f.Name())

	confSettings := struct {
		Host    string        `yaml:"host"`
		Port    string        `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	}{
		Host:    "localhost",
		Port:    "18080",
		Timeout: time.Second,
	}

	data, err := yaml.Marshal(confSettings)
	assert.NoError(t, err, "cannot marshal struct")
	_, err = f.Write(data)
	assert.NoError(t, err, "cannot write data")
	err = f.Close()
	assert.NoError(t, err, "cannot close temp file")

	err = c.LoadYAML(f.Name())
	assert.NoError(t, err, "cannot load .yaml file")

	res := struct {
		Host    string        `yaml:"host"`
		Port    string        `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	}{}
	err = c.ReadYAML(&res)
	assert.NoError(t, err, "cannot read yaml")

	assert.Equal(t, confSettings.Host, res.Host)
	assert.Equal(t, confSettings.Port, res.Port)
	assert.Equal(t, confSettings.Timeout, res.Timeout)
}

func TestLoadConf_IncorrectFileExtension(t *testing.T) {
	c := &reader{}

	f, err := os.CreateTemp("", "tempfile.*.txt")
	assert.NoError(t, err, "cannot create file")
	defer os.Remove(f.Name())

	assert.NoError(t, err, "cannot marshal struct")
	_, err = f.Write([]byte("incorrect file extension"))
	assert.NoError(t, err, "cannot write data")
	err = f.Close()
	assert.NoError(t, err, "cannot close temp file")

	err = c.LoadYAML(f.Name())
	assert.Error(t, err)
}

func TestReadConf_IncorrectYAML(t *testing.T) {
	c := &reader{}

	f, err := os.CreateTemp("", "tempfile.*.yaml")
	assert.NoError(t, err, "cannot create file")
	defer os.Remove(f.Name())

	assert.NoError(t, err, "cannot marshal struct")
	_, err = f.Write([]byte("incorrect yaml format"))
	assert.NoError(t, err, "cannot write data")
	err = f.Close()
	assert.NoError(t, err, "cannot close temp file")

	err = c.LoadYAML(f.Name())
	assert.NoError(t, err)

	res := struct {
		Host    string        `yaml:"host"`
		Port    string        `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	}{}
	err = c.ReadYAML(&res)
	assert.Error(t, err)
}

func TestNewConf_CorrectYAML(t *testing.T) {
	f, err := os.CreateTemp("", "tempfile.*.yaml")
	assert.NoError(t, err, "cannot create file")
	defer os.Remove(f.Name())

	type Config struct {
		Host    string        `yaml:"host"`
		Port    string        `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	}
	confSettings := struct {
		Conf Config `yaml:"config"`
	}{
		Conf: Config{
			Host:    "localhost",
			Port:    "18080",
			Timeout: time.Second,
		},
	}

	data, err := yaml.Marshal(confSettings)
	assert.NoError(t, err, "cannot marshal struct")
	_, err = f.Write(data)
	assert.NoError(t, err, "cannot write data")
	err = f.Close()
	assert.NoError(t, err, "cannot close temp file")

	res := struct {
		Conf Config `yaml:"config"`
	}{}

	c, err := New(f.Name())
	assert.NoError(t, err)
	err = c.ReadYAML(&res)
	assert.NoError(t, err, "cannot read yaml")

	assert.Equal(t, confSettings.Conf.Host, res.Conf.Host)
	assert.Equal(t, confSettings.Conf.Port, res.Conf.Port)
	assert.Equal(t, confSettings.Conf.Timeout, res.Conf.Timeout)
}
