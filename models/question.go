package models

import (
	"database/sql"
	"math/rand"
)

type Question struct {
	QuestionID           int      `json:"questionId"`
	PeopleQuestionID     int      `json:"people_question_id" form:"people_question_id"`
	Question             string   `json:"question"`
	PhotoURL             string   `json:"photo"`
	ShowImageAtBeginning bool     `json:"showImageAtBeginning"`
	Answers              []Answer `json:"answers"`
}

func (q *Question) GetQuestion(db *sql.DB) error {

	err := db.QueryRow("SELECT q.question, i.url, q.showImageAtBeginning FROM Questions q, PeopleQuestions pq, Images i WHERE pq.id = ? AND q.id = pq.id AND pq.photoId = i.id", q.PeopleQuestionID).Scan(&q.Question, &q.PhotoURL, &q.ShowImageAtBeginning)

	if err != nil {
		return err
	}

	answers, err := q.getAnswers(db)

	if err != nil {
		return err
	}

	q.Answers = answers

	return nil

}

func (q *Question) GetRandomQuestion(user int, level int, db *sql.DB) error {

	var questionsIds []int

	results, err := db.Query("SELECT pq.id FROM PeopleQuestions pq, Questions q, QuestionsLevel ql WHERE pq.questionId = q.id AND q.id = ql.questionId AND peopleId != ? AND ql.levelId = ?", user, level)

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		var questionId int

		err = results.Scan(&questionId)

		if err != nil {
			panic(err.Error())
		}

		questionsIds = append(questionsIds, questionId)

	}

	q.PeopleQuestionID = random(questionsIds)

	if err := q.GetQuestion(db); err != nil {
		panic(err.Error())
	}

	return nil

}

func (q *Question) GetUnansweredQuestion(user Legend, relatedUser Legend, db *sql.DB) error {

	var questionsIds []int

	results, err := db.Query("SELECT DISTINCT(q.questionId) FROM PeopleQuestions pq, Questions q WHERE pq.questionId = q.id AND peopleId != ? AND pq.peopleId = ?", user.ID, relatedUser.ID)

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		var questionId int

		err = results.Scan(&questionId)

		if err != nil {
			panic(err.Error())
		}

		questionsIds = append(questionsIds, questionId)

	}

	q.PeopleQuestionID = random(questionsIds)

	if err := q.GetQuestion(db); err != nil {
		panic(err.Error())
	}

	return nil

}

func (q *Question) getAnswers(db *sql.DB) ([]Answer, error) {

	var answers []Answer

	results, err := db.Query("SELECT answer, isCorrect FROM Answers WHERE questionPeopleId = ?", q.PeopleQuestionID)

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		var answer Answer

		err = results.Scan(&answer.Answer, &answer.IsCorrect)

		if err != nil {
			panic(err.Error())
		}

		answers = append(answers, answer)

	}

	return answers, nil

}

func random(ids []int) int {

	totalItems := len(ids)
	rand := rand.Intn(totalItems)

	return ids[rand]

}
