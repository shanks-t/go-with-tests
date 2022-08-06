// reflections on ideas around di that came up in my interview
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Interview interface {
	askQuestion() []string
	getDifficulty(index int) int
	source() string
}

type Question struct {
	questionString string
	difficulty     int
}

type Interviewer struct {
	name               string
	interviewQuestions []Question
}

type Assessment struct {
	sumOfWeights            float32
	sumOfAnswersTimesWeight float32
	sumQuestions            int
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

func (i Interviewer) getDifficulty(index int) int {
	return i.interviewQuestions[index].difficulty
}

func ConductInterview(i []Interview) bool {
	a := Assessment{}
	for _, interview := range i {
		fmt.Printf("Here are some questions from %s:\n", interview.source())

		for i, q := range interview.askQuestion() {
			fmt.Printf("%d: %s\n", i+1, q)
			rand.Seed(int64(time.Now().Nanosecond()))
			answerQuality := rand.Float32() * 10
			a.sumQuestions++
			fmt.Printf("answer difficulty: %v \tanswer quality: %.2f\n", interview.getDifficulty(i), answerQuality)
			questionWeight := float32(interview.getDifficulty(i))
			a.sumOfWeights += questionWeight
			a.sumOfAnswersTimesWeight += (float32(questionWeight) * answerQuality)
		}
	}
	return findOutIfYouPassed(a)
}

func findOutIfYouPassed(a Assessment) bool {
	passed := false
	weightedResult := a.sumOfAnswersTimesWeight / a.sumOfWeights
	fmt.Printf("\nYou answered %v questions, with a weighted average of %.2f\n", a.sumQuestions, weightedResult)
	if weightedResult > 6 {
		passed = true
		fmt.Println("Excellent work, we'd like to invite you back for more hard questions!")
	} else {
		fmt.Println("\nWe are not moving forward with you in the interview process. Thanks for applying and good luck!")
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
	ConductInterview(interviewers)
}
