package thread_creator

import (
	"context"
	"errors"
	"testing"
)

// --- Mock implementations ---
type mockImageGenerator struct {
	deployImage string
	deployErr   error
	reviewImage string
	reviewErr   error
}

func (m *mockImageGenerator) GenerateDeployThreadImage(ctx context.Context) (string, error) {
	return m.deployImage, m.deployErr
}
func (m *mockImageGenerator) GenerateReviewThreadImage(ctx context.Context) (string, error) {
	return m.reviewImage, m.reviewErr
}

type mockTextGenerator struct {
	noImageReason   string
	interestingFact string
}

func (m *mockTextGenerator) GenerateNoImageReasonText(ctx context.Context) string {
	return m.noImageReason
}
func (m *mockTextGenerator) GenerateInterestingFactText(ctx context.Context) string {
	return m.interestingFact
}

type mockChannel struct {
	sentWithImage   bool
	sentText        string
	sentTitle       string
	sentDescription string
	imagePath       string
	msgErr          error
}

func (m *mockChannel) SendMessageWithImage(filepath, title, description string, ctx context.Context) error {
	m.sentWithImage = true
	m.imagePath = filepath
	m.sentTitle = title
	m.sentDescription = description
	return m.msgErr
}

func (m *mockChannel) SendMessage(text string, ctx context.Context) error {
	m.sentWithImage = false
	m.sentText = text
	return m.msgErr
}

type mockFileManager struct {
	saveErr error
}

func (m *mockFileManager) SaveBase64File(filepath string, encodedContent string) error {
	return m.saveErr
}
func (m *mockFileManager) DeleteFile(filepath string) error {
	return nil
}

// --- Inject mock file manager ---
func newThreadCreatorWithMocks(img *mockImageGenerator, txt *mockTextGenerator, ch *mockChannel, fm *mockFileManager, withImage bool) *ThreadCreator {
	tc := NewThreadCreator(img, txt, ch, withImage)
	tc.fileManager = fm
	return tc
}

// --- Tests ---
func TestThreadCreator_CreateDeployThread(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		imgGen         *mockImageGenerator
		txtGen         *mockTextGenerator
		channel        *mockChannel
		fileManager    *mockFileManager
		expectImageMsg bool
		expectErr      bool
	}{
		{
			name:           "success with image",
			imgGen:         &mockImageGenerator{deployImage: "imgdata"},
			txtGen:         &mockTextGenerator{interestingFact: "fact"},
			channel:        &mockChannel{},
			fileManager:    &mockFileManager{},
			expectImageMsg: true,
			expectErr:      false,
		},
		{
			name:           "image generation fails",
			imgGen:         &mockImageGenerator{deployErr: errors.New("fail")},
			txtGen:         &mockTextGenerator{interestingFact: "fact", noImageReason: "noimg"},
			channel:        &mockChannel{},
			fileManager:    &mockFileManager{},
			expectImageMsg: false,
			expectErr:      false,
		},
		{
			name:           "file save fails",
			imgGen:         &mockImageGenerator{deployImage: "imgdata"},
			txtGen:         &mockTextGenerator{interestingFact: "fact"},
			channel:        &mockChannel{},
			fileManager:    &mockFileManager{saveErr: errors.New("savefail")},
			expectImageMsg: false,
			expectErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tc := newThreadCreatorWithMocks(tt.imgGen, tt.txtGen, tt.channel, tt.fileManager, true)
			err := tc.CreateDeployThread(ctx)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
			if tt.expectImageMsg && !tt.channel.sentWithImage {
				t.Errorf("expected SendMessageWithImage to be called")
			}
			if !tt.expectImageMsg && tt.channel.sentWithImage {
				t.Errorf("expected SendMessage to be called")
			}
		})
	}
}

func TestThreadCreator_CreateDeployThread_NoImage(t *testing.T) {
	ctx := context.Background()
	imgGen := &mockImageGenerator{}
	textGen := &mockTextGenerator{}
	channel := &mockChannel{sentText: "text"}
	fileManager := &mockFileManager{}

	tc := newThreadCreatorWithMocks(imgGen, textGen, channel, fileManager, false)
	err := tc.CreateDeployThread(ctx)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if channel.sentWithImage {
		t.Errorf("expected SendMessage to be called, but SendMessageWithImage was called")
	}
}

func TestThreadCreator_CreateReviewThread(t *testing.T) {
	ctx := context.Background()

	tc := newThreadCreatorWithMocks(
		&mockImageGenerator{reviewImage: "imgdata"},
		&mockTextGenerator{interestingFact: "fact"},
		&mockChannel{},
		&mockFileManager{},
		true,
	)
	err := tc.CreateReviewThread(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestThreadCreator_CreateReviewThread_NoImage(t *testing.T) {
	ctx := context.Background()
	imgGen := &mockImageGenerator{}
	textGen := &mockTextGenerator{}
	channel := &mockChannel{}
	fileManager := &mockFileManager{}

	tc := newThreadCreatorWithMocks(imgGen, textGen, channel, fileManager, false)
	err := tc.CreateReviewThread(ctx)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if channel.sentWithImage {
		t.Errorf("expected SendMessage to be called, but SendMessageWithImage was called")
	}

}
