// reflections on ideas around di that came up in my interview
package dependencyinjection

import (
	"fmt"
	"time"
)

type Interview interface {
	askQuestion() []string
	getDifficulty() int
	source() string
	getAssessment(answerQuality float32) float32
}

type Question struct {
	questionString string
	difficulty     int
}

type Interviewer struct {
	name               string
	interviewQuestions []Question
}

func (i Interviewer) askQuestion() []string {
	var questions []string
	for _, question := range i.interviewQuestions {
		questions = append(questions, question.questionString)
	}
	return questions
}

func (i Interviewer) source() string {
	return i.name
}

func (i Interviewer) getDifficulty() int {
	return i.interviewQuestions[0].difficulty
}

func (i Interviewer) getAssessment(answerQuality float32) float32 {
	return float32(i.getDifficulty()) * float32(answerQuality)
}

func conductInterview(i []Interview) []float32 {
	const answerQuality float32 = 9.9
	var assessmentOfAnswers []float32
	for _, interview := range i {

		time.Sleep(1 * time.Second)
		fmt.Printf("Here are some questions from %s:\n", interview.source())
		for i, q := range interview.askQuestion() {
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("%d: %s\n", i+1, q)
			assessment := interview.getAssessment(answerQuality)
			assessmentOfAnswers = append(assessmentOfAnswers, assessment)
			time.Sleep(2 * time.Second)
			fmt.Printf("The impressiveness of your answer on a scale of 1 - 100 for question %d from %s: %v\n", i+1, interview.source(), assessment)
		}
	}
	return assessmentOfAnswers
}

func findOutIfYouPassed(assessments []float32) bool {
	var score float32
	var sum float32
	passed := false
	for _, a := range assessments {
		sum += a
		score = sum / float32(len(assessments))
	}
	if score > 75 {
		passed = true
	}
	return passed
}

func main() {
	interviewer1 := Interviewer{
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
	interviewer2 := Interviewer{
		name: "Yohay", interviewQuestions: []Question{
			{
				questionString: "What is the best way to handle errors?",
				difficulty:     9,
			},
			{
				questionString: "Give me a concrete example of polymorphism.",
				difficulty:     9,
			},
		},
	}
	interviewer3 := Interviewer{
		name: "John", interviewQuestions: []Question{
			{
				questionString: "Explain a recent project you've done at your job",
				difficulty:     7,
			},
			{
				questionString: "Explain the cap theorem and the types of tradeoffs of system design, if any, you've encountered at work",
				difficulty:     8,
			},
		},
	}
	interviewers := []Interview{interviewer1, interviewer2, interviewer3}
	totalOfAssessment := conductInterview(interviewers)
	fmt.Println("\n....interviewers are discussing your answers")
	time.Sleep(3 * time.Second)

	if findOutIfYouPassed(totalOfAssessment) {
		fmt.Println("\nYou Passed!")
	} else {
		fmt.Println("\nWe are not moving forward with you in the interview process. Thanks for applying and good luck!")
	}
}
