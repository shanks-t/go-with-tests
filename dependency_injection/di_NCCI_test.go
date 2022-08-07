package main

import "testing"

var MockInterviewer1 = Interviewer{
	name: "Ovidio", interviewQuestions: []Question{
		{
			questionString: "Explain dependency injection",
			difficulty:     9,
		},
		{
			questionString: "Explain polymorphism",
			difficulty:     8,
		},
	},
}

func ConductInterviewTest(t *testing.T) {
	mockInterviewers := []Interview{MockInterviewer1}
	got := ConductInterview(mockInterviewers)
	want := 

}

func BenchmarkConductInterview(b *testing.B) {
	ConductInterview()
}
