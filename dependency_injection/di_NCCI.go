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
	//getAssessment(questionIndex int, answerQuality float32) float32
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

// implement weighted average for scores
// func (i Interviewer) getAssessment(questionIndex int, answerQuality float32) map[string]float32 {
// 	//get sum of all weights
// 	var sumOfWeights float32
// 	var answerSumTimesWeight float32
// 	sumOfWeights += float32(i.getDifficulty(questionIndex))
// 	answerSumTimesWeight += float32(i.getDifficulty(questionIndex)) * answerQuality
// 	answerMap := map[string]float32{
// 		"sumOfWeights":         sumOfWeights,
// 		"answerSumTimesWeight": answerSumTimesWeight,
// 	}
// 	return answerMap
// }

func conductInterview(i []Interview) bool {
	a := Assessment{}
	for _, interview := range i {
		//time.Sleep(1 * time.Second)
		fmt.Printf("Here are some questions from %s:\n", interview.source())

		for i, q := range interview.askQuestion() {
			rand.Seed(int64(time.Now().Nanosecond()))
			answerQuality := rand.Float32() * 10
			a.sumQuestions++
			fmt.Printf("%d: %s\n", i+1, q)
			fmt.Printf("answer quality: %v\n", answerQuality)
			weightSum := float32(interview.getDifficulty(i))
			a.sumOfWeights += weightSum
			a.sumOfAnswersTimesWeight += (float32(weightSum) * answerQuality)
		}
	}
	return findOutIfYouPassed(a)
}

func findOutIfYouPassed(a Assessment) bool {
	passed := false
	fmt.Printf("sumquestions: %v\n", a.sumQuestions)
	weightedResult := a.sumOfAnswersTimesWeight / a.sumOfWeights
	fmt.Printf("weighted result: %v\n", weightedResult)
	fmt.Printf("sumOfAnswerTimesWeight: %v\n", a.sumOfAnswersTimesWeight)

	if weightedResult > 5 {
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
	didPass := conductInterview(interviewers)
	fmt.Println("\n....interviewers are discussing your answers")
	//time.Sleep(1 * time.Second)

	if didPass {
		fmt.Println("\nYou Passed!")
	} else {
		fmt.Println("\nWe are not moving forward with you in the interview process. Thanks for applying and good luck!")
	}
}
