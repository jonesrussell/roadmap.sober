package services

import "fmt"

type Step struct {
	Text       string `json:"text"`
	Content    string `json:"content"`
	PanelTitle string `json:"panel_title"`
}

type StepService struct {
	// Add fields here for database connection, etc.
}

func NewStepService() *StepService {
	// Initialize the service here, e.g. establish a database connection
	return &StepService{}
}

func (s *StepService) GetSteps() []Step {
	// Define the steps in the sobriety journey
	return []Step{
		{Text: "Admit the Problem", Content: "Content for Admit the Problem", PanelTitle: "Admit the Problem"},
		{Text: "Seek Help", Content: "Content for Seek Help", PanelTitle: "Seek Help"},
		{Text: "Detoxification", Content: "Content for Detoxification", PanelTitle: "Detoxification"},
		{Text: "Therapy and Counseling", Content: "Content for Therapy and Counseling", PanelTitle: "Therapy and Counseling"},
		{Text: "Support Groups", Content: "Content for Support Groups", PanelTitle: "Support Groups"},
		{Text: "Healthy Lifestyle", Content: "Content for Healthy Lifestyle", PanelTitle: "Healthy Lifestyle"},
		{Text: "Avoid Triggers", Content: "Content for Avoid Triggers", PanelTitle: "Avoid Triggers"},
		{Text: "Long-Term Maintenance", Content: "Content for Long-Term Maintenance", PanelTitle: "Long-Term Maintenance"},
		{Text: "Forgive Yourself", Content: "Content for Forgive Yourself", PanelTitle: "Forgive Yourself"},
		{Text: "Celebrate Milestones", Content: "Content for Celebrate Milestones", PanelTitle: "Celebrate Milestones"},
	}
}

func (s *StepService) GetStepById(id string) (*Step, error) {
	// Implement your logic here to get a step by its ID.
	// This is just a placeholder implementation.
	for _, step := range s.GetSteps() {
		if step.Text == id {
			return &step, nil
		}
	}
	return nil, fmt.Errorf("step with id %s not found", id)
}
