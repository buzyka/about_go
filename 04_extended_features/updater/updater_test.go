package updater

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockUpdatingPerson struct {
	mock.Mock
}

func (m *MockUpdatingPerson) GetBirthday() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}

func (m *MockUpdatingPerson) HaveBirthday() {
	m.Called()
}

func TestUpdateBirthdaysForDate(t *testing.T) {
	// Create a mock person with a birthday
	mockPerson1 := &MockUpdatingPerson{}

	mockPerson1.On("GetBirthday").Return(time.Date(2010, 5, 5, 0, 0, 0, 0, time.UTC)).Once()
	mockPerson1.On("HaveBirthday").Return().Once()

	mockPerson2 := &MockUpdatingPerson{}
	mockPerson2.On("GetBirthday").Return(time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC)).Once()
	mockPerson2.AssertNotCalled(t, "HaveBirthday")

	// Create an updater and add the mock person
	updater := &Updater{
		Persons: []UpdatingPerson{
			mockPerson1,
			mockPerson2,
		},
	}
	updater.UpdateBirthdaysForDate(time.Date(2010, 5, 5, 0, 0, 0, 0, time.UTC))

	// Assert that HaveBirthday was called for the first person
	mockPerson1.AssertExpectations(t)
	mockPerson2.AssertExpectations(t)
}
