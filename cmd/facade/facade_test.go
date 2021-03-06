package main

import (
	"testing"

	"github.com/sabrusrin/wildberries_st5/pkg/facade"
	"github.com/sabrusrin/wildberries_st5/pkg/mentor"
	"github.com/sabrusrin/wildberries_st5/pkg/models"
	"github.com/sabrusrin/wildberries_st5/pkg/practice"
	"github.com/sabrusrin/wildberries_st5/pkg/theory"
)

func TestOk(t *testing.T)	{
	th := theory.NewTheory(models.TheoryHeader)
	m := mentor.NewMentor(models.MentorHeader)
	p := practice.NewPractice(models.PracticeHeader)
	wbTrial := facade.NewPlanner(th, m, p)
	// testing for empty string argument, adding a task to the task list and appending task list
	okResult := models.TheoryHeader + models.MentorHeader + models.PracticeHeader
	out := wbTrial.Plan("", "")
	if out != okResult	{
		t.Errorf("Test for empty string argument failed!\nExpected:\n%v\nGot %v", okResult, out)
	}
	okResult = models.TheoryHeader + " Go101\n" + models.MentorHeader + models.PracticeHeader + " Implement facade pattern\n"
	out = wbTrial.Plan("Go101", "Implement facade pattern")
	if out != okResult	{
		t.Errorf("Test for valid argument failed!\nExpected:\n%v\nGot:\n%v", okResult, out)
	}
	okResult = models.TheoryHeader + " Go101, Concurrency in Go\n" + models.MentorHeader + models.PracticeHeader + " Implement facade pattern, Implement builder pattern\n"
	out = wbTrial.Plan("Concurrency in Go", "Implement builder pattern")
	if out != okResult	{
		t.Errorf("Test for correct task list appending failed!\nExpected:\n%v\nGot %v", okResult, out)
	}
}
