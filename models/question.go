package models

import (
	"database/sql"
	"math/rand"
)

type Question struct {
	QuestionID           int      `json:"questionId,omitempty"`
	Person               Legend   `json:"legend,omitempty"`
	PeopleQuestionID     int      `json:"people_question_id,omitempty" form:"people_question_id"`
	Question             string   `json:"question"`
	PhotoURL             string   `json:"photo,omitempty"`
	ShowImageAtBeginning bool     `json:"showImageAtBeginning,omitempty"`
	Answers              []Answer `json:"answers,omitempty"`
}

func (q *Question) GetQuestion(db *sql.DB) error {

	err := db.QueryRow("SELECT q.question, i.url, q.showImageAtBeginning FROM Questions q, PeopleQuestions pq, Images i WHERE pq.id = ? AND q.id = pq.questionId AND pq.photoId = i.id", q.PeopleQuestionID).Scan(&q.Question, &q.PhotoURL, &q.ShowImageAtBeginning)

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

func (q *Question) GetUnansweredQuestion(questionFail int, db *sql.DB) error {

	err := q.getPeopleFromQuestion(questionFail, db)

	if err != nil {
		panic(err.Error())
	}

	var questionsIds []int

	results, err := db.Query("SELECT DISTINCT(q.id) FROM Questions q LEFT JOIN PeopleQuestions pq ON pq.questionId = q.id WHERE pq.peopleId != ?", q.Person.ID)

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

	q.QuestionID = random(questionsIds)

	if err := q.GetQuestionDescription(db); err != nil {
		panic(err.Error())
	}

	return nil

}

func (q *Question) getAnswers(db *sql.DB) ([]Answer, error) {

	var answers []Answer

	results, err := db.Query("SELECT id, answer FROM Answers WHERE questionPeopleId = ?", q.PeopleQuestionID)

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		var answer Answer

		err = results.Scan(&answer.ID, &answer.Answer)

		if err != nil {
			panic(err.Error())
		}

		answers = append(answers, answer)

	}

	return answers, nil

}

func (q *Question) getPeopleFromQuestion(question int, db *sql.DB) error {

	err := db.QueryRow("SELECT pq.peopleId, p.name FROM PeopleQuestions pq, People p WHERE pq.peopleId = p.id AND pq.id = ?", question).Scan(&q.Person.ID, &q.Person.Name)

	if err != nil {
		return err
	}

	return nil

}

func (q *Question) GetQuestionDescription(db *sql.DB) error {

	err := db.QueryRow("SELECT q.question FROM Questions q WHERE q.id = ?", q.QuestionID).Scan(&q.Question)

	if err != nil {
		return err
	}

	return nil

}

func random(ids []int) int {

	totalItems := len(ids)
	rand := rand.Intn(totalItems)

	return ids[rand]

}
