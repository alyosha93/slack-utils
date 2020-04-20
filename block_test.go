package utils

import (
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
	"github.com/slack-go/slack"
)

const (
	mockActionID             = "fAkE123"
	mockValue                = "click_me_btn"
	fakeStyle    slack.Style = "fake"
)

var mockTextObj = &slack.TextBlockObject{
	Type:     slack.PlainTextType,
	Text:     "Do something",
	Emoji:    false,
	Verbatim: false,
}

func TestNewButtonWithStyle(t *testing.T) {
	testCases := []struct {
		description string
		style       slack.Style
		wantButton  *slack.ButtonBlockElement
	}{
		{
			description: "returns default style button",
			style:       slack.StyleDefault,
			wantButton: &slack.ButtonBlockElement{
				Type:     slack.METButton,
				ActionID: mockActionID,
				Text:     mockTextObj,
				Value:    mockValue,
				Style:    slack.StyleDefault,
			},
		},
		{
			description: "returns danger style button",
			style:       slack.StyleDanger,
			wantButton: &slack.ButtonBlockElement{
				Type:     slack.METButton,
				ActionID: mockActionID,
				Text:     mockTextObj,
				Value:    mockValue,
				Style:    slack.StyleDanger,
			},
		},
		{
			description: "returns primary style button",
			style:       slack.StylePrimary,
			wantButton: &slack.ButtonBlockElement{
				Type:     slack.METButton,
				ActionID: mockActionID,
				Text:     mockTextObj,
				Value:    mockValue,
				Style:    slack.StylePrimary,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			btn := NewButtonWithStyle(mockActionID, mockValue, mockTextObj, tc.style)
			if diff := pretty.Compare(btn, tc.wantButton); diff != "" {
				t.Fatalf("expected to receive button: %v, got: %v", tc.wantButton, btn)
			}
		})
	}
}

func TestNewDatePickerAtTime(t *testing.T) {
	now := time.Now()
	expectedStr := now.Format("2006-01-02")

	testCases := []struct {
		description    string
		placeholder    *slack.TextBlockObject
		wantDatePicker *slack.DatePickerBlockElement
	}{
		{
			description: "returns new date picker with placeholder, which trumps initial date",
			placeholder: mockTextObj,
			wantDatePicker: &slack.DatePickerBlockElement{
				Type:        slack.METDatepicker,
				ActionID:    mockActionID,
				Placeholder: mockTextObj,
			},
		},
		{
			description: "returns new date picker with proper initial date",
			wantDatePicker: &slack.DatePickerBlockElement{
				Type:        slack.METDatepicker,
				ActionID:    mockActionID,
				InitialDate: expectedStr,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			picker := NewDatePickerWithOpts(mockActionID, tc.placeholder, now)
			if diff := pretty.Compare(picker, tc.wantDatePicker); diff != "" {
				t.Fatalf("expected to receive datepicker: %v, got: %v", tc.wantDatePicker, picker)
			}
		})
	}

}

func TestNewDateOptToTime(t *testing.T) {
	now := time.Now()
	dateOptStr := now.Format(datePickTimeFmt)
	unrecognizedLayoutOptStr := now.Format(time.ANSIC)
	expectedTime, err := time.Parse(datePickTimeFmt, dateOptStr)
	if err != nil {
		t.Fatalf("received unexpected error: %s", err)
		return
	}

	testCases := []struct {
		description string
		opt         string
		wantErr     bool
	}{
		{
			description: "returns parsed time value with no error",
			opt:         dateOptStr,
		},
		{
			description: "unrecognized layout string causes error",
			opt:         unrecognizedLayoutOptStr,
			wantErr:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			parsedTime, err := DateOptToTime(dateOptStr)
			if err != nil && !tc.wantErr {
				t.Fatalf("received unexpected error: %s", err)
				return
			}
			if diff := pretty.Compare(parsedTime, expectedTime); diff != "" {
				t.Fatalf("expected to receive time: %v, got: %v", expectedTime, parsedTime)
			}
		})
	}
}
