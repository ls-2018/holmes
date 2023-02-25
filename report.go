package holmes

import "time"

type ProfileReporter interface {
	Report(pType string, filename string, reason ReasonType, eventID string, sampleTime time.Time, pprofBytes []byte, scene Scene) error
}

// rptEvent stands of the args of report event
type rptEvent struct {
	PType      string
	FileName   string
	Reason     ReasonType
	EventID    string
	SampleTime time.Time
	PprofBytes []byte
	Scene      Scene
}

// Scene 包含profile触发时的场景信息，包括当前值、平均值和配置。
type Scene struct {
	typeOption
	CurVal int // dump时的值
	Avg    int // Avg是过去值的平均值
}

type ReasonType uint8

const (
	ReasonCurlLessMin ReasonType = iota
	ReasonCurlGreaterMin
	ReasonCurGreaterMax
	ReasonCurGreaterAbs
	ReasonDiff
)

func (rt ReasonType) String() string {
	var reason string
	switch rt {
	case ReasonCurlLessMin:
		reason = "curVal < ruleMin"
	case ReasonCurlGreaterMin:
		reason = "curVal >= ruleMin, but don't meet diff trigger condition"
	case ReasonCurGreaterMax:
		reason = "curVal >= ruleMax"
	case ReasonCurGreaterAbs:
		reason = "curVal > ruleAbs"
	case ReasonDiff:
		reason = "curVal >= ruleMin, and meet diff trigger condition"

	}

	return reason
}
