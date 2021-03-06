// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sampling // import "github.com/layasugar/otel-collector/tailsamplingprocessor/internal/sampling"

import (
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
)

type latency struct {
	logger      *zap.Logger
	thresholdMs int64
}

var _ PolicyEvaluator = (*latency)(nil)

// NewLatency creates a policy evaluator sampling traces with a duration higher than a configured threshold
func NewLatency(logger *zap.Logger, thresholdMs int64) PolicyEvaluator {
	return &latency{
		logger:      logger,
		thresholdMs: thresholdMs,
	}
}

// OnLateArrivingSpans notifies the evaluator that the given list of spans arrived
// after the sampling decision was already taken for the trace.
// This gives the evaluator a chance to log any message/metrics and/or update any
// related internal state.
func (l *latency) OnLateArrivingSpans(Decision, []*pdata.Span) error {
	l.logger.Debug("Triggering action for late arriving spans in latency filter")
	return nil
}

// Evaluate looks at the trace data and returns a corresponding SamplingDecision.
func (l *latency) Evaluate(_ pdata.TraceID, traceData *TraceData) (Decision, error) {
	l.logger.Debug("Evaluating spans in latency filter")

	traceData.Lock()
	batches := traceData.ReceivedBatches
	traceData.Unlock()

	var minTime pdata.Timestamp
	var maxTime pdata.Timestamp

	return hasSpanWithCondition(batches, func(span pdata.Span) bool {
		if minTime == 0 || span.StartTimestamp() < minTime {
			minTime = span.StartTimestamp()
		}
		if maxTime == 0 || span.EndTimestamp() > maxTime {
			maxTime = span.EndTimestamp()
		}

		duration := maxTime.AsTime().Sub(minTime.AsTime())
		return duration.Milliseconds() >= l.thresholdMs
	}), nil
}
