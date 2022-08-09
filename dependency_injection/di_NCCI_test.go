package main

import "testing"

var (
	mockInterviewers = []Interview{MockInterviewer1, hiringManager1}

	MockInterviewer1 = Interviewer{
		name: "Ovidio", interviewQuestions: []Question{
			{
				questionString: "Explain dependency injection",
				difficulty:     0,
			},
			{
				questionString: "Explain polymorphism",
				difficulty:     0,
			},
		},
	}
	hiringManager1 = HiringManager{
		name: "John", interviewQuestions: []Question{
			{
				questionString: "Explain the cap theorem and the types of tradeoffs of system design, if any, you've encountered at work",
				difficulty:     0,
			},
		},
	}
)

func TestConductInterview(t *testing.T) {

	got := ConductInterview(mockInterviewers)
	want := Assessment{
		sumOfWeights:            0,
		sumOfAnswersTimesWeight: GetAnswerQuality()*float32(MockInterviewer1.getDifficulty(0)) + GetAnswerQuality()*float32(MockInterviewer1.getDifficulty(1)),
		sumQuestions:            3,
	}

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func BenchmarkConductInterview(b *testing.B) {
	ConductInterview(mockInterviewers)
}
