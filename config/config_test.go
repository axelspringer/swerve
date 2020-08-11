package config_test

import (
	"os"
	"testing"

	"github.com/axelspringer/swerve/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Configuration Suite")
}

var _ = Describe("Config", func() {
	var conf *config.Configuration

	BeforeEach(func() {
		conf = config.NewConfiguration()
	})

	Describe("Pulling configuration from the environment", func() {
		Context("with no env vars defined", func() {
			It("should not panic and use default values", func() {
				err := conf.FromEnv()
				Expect(*conf).Should(Equal(*config.NewConfiguration()))
				Expect(err).Should(BeNil())
			})
		})
		Context("with env vars of wrong type", func() {
			It("should throw error", func() {
				os.Setenv("EVADE_API_LISTENER", ":foo")
				err := conf.FromEnv()
				Expect(err.Error()).Should(Equal(config.ErrAPIPortInvalid))
			})
			defer os.Clearenv()
		})
	})
})
