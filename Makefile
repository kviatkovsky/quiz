
start: #start server
	go run main.go start

compare: #compare your score with other others
	go run main.go compare


submit: # Submit answers for question.
	go run main.go submit 3 2 3

questions: # Get all questions
	go run main.go questions


.PHONY: start compare submit submit
