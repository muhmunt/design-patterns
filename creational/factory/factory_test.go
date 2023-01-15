package factory

import (
	"strings"
	"testing"
)

func TestAssignSMSNotifierFactory(t *testing.T) {
	notification := Notification{"8832313", "Hello factory!"}
	notifier, err := AssignNotifier(SMS)

	if err != nil {
		t.Fatal("Notifier type 'SMS' must be returned")
	}

	resultStr := notifier.SendNotification(notification)
	if !strings.Contains(resultStr, "was sent to number") {
		t.Error("SMS notifier wront message")
	}
}

func TestAssignEmailNotifierFactory(t *testing.T) {
	notification := Notification{"8832313", "Hello factory!"}
	notifier, err := AssignNotifier(Email)

	if err != nil {
		t.Fatal("Notifier type 'SMS' must be returned")
	}

	resultStr := notifier.SendNotification(notification)
	if !strings.Contains(resultStr, "was sent to email") {
		t.Error("SMS notifier wront message")
	}
}

func TestAssignNotExistNotifierFactory(t *testing.T) {
	_, err := AssignNotifier(3)

	if err == nil {
		t.Fatal("Error must be returned cause 3 is unrecognized type")
	}
}
