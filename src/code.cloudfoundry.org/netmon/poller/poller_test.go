package poller_test

import (
	libfakes "code.cloudfoundry.org/lib/fakes"

	"os"
	"time"

	"code.cloudfoundry.org/lager/lagertest"
	"code.cloudfoundry.org/netmon/poller"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Poller Run", func() {
	var (
		iptables *libfakes.IPTablesAdapter
		logger   *lagertest.TestLogger

		metrics      *poller.SystemMetrics
		pollInterval time.Duration
	)

	BeforeEach(func() {
		iptables = &libfakes.IPTablesAdapter{}
		logger = lagertest.NewTestLogger("test")
		pollInterval = 1 * time.Second

		iptables.RuleCountReturnsOnCall(0, 2, nil)
		iptables.RuleCountReturnsOnCall(1, 2, nil)

		metrics = &poller.SystemMetrics{
			Logger:          logger,
			PollInterval:    pollInterval,
			InterfaceName:   "meow",
			IPTablesAdapter: iptables,
		}
	})

	It("should report measurements once within single interval", func() {
		runTest(metrics, pollInterval)
		Expect(logger.LogMessages()).To(Equal([]string{
			"test.measure.measure-start",
			"test.measure.metric-sent",
			"test.measure.metric-sent",
			"test.measure.read-tx-bytes",
			"test.measure.measure-complete",
		}))
	})

	It("should use the iptables adapter when checking the rules", func() {
		runTest(metrics, pollInterval)
		Expect(iptables.RuleCountCallCount()).To(Equal(2))

		table := iptables.RuleCountArgsForCall(0)
		Expect(table).To(Equal("filter"))
		table = iptables.RuleCountArgsForCall(1)
		Expect(table).To(Equal("nat"))

		iptablesLog := logger.Logs()[2]
		Expect(iptablesLog.Data["IPTablesRuleCount"]).To(Equal(float64(4)))
	})

	Context("when a telemetry logger is configured", func() {
		var (
			telemetryLogger *lagertest.TestLogger
		)

		BeforeEach(func() {
			telemetryLogger = lagertest.NewTestLogger("telemetry")
			metrics = &poller.SystemMetrics{
				Logger:          logger,
				TelemetryLogger: telemetryLogger,
				PollInterval:    pollInterval,
				InterfaceName:   "meow",
				IPTablesAdapter: iptables,
			}
		})

		It("logs the number of iptables rules in the telemetry log", func() {
			runTest(metrics, pollInterval)
			Expect(telemetryLogger.LogMessages()).To(Equal([]string{"telemetry.count-iptables-rules"}))
		})
	})
})

func runTest(metrics *poller.SystemMetrics, pollInterval time.Duration) {
	doneCh := make(chan os.Signal)
	readyCh := make(chan struct{})

	go metrics.Run(doneCh, readyCh)

	<-readyCh
	<-time.After(pollInterval + 10*time.Millisecond)
	doneCh <- os.Interrupt
}
