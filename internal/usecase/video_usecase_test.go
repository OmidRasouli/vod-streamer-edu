package usecase

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/OmidRasouli/vod-streamer-edu/internal/domain/model"
	"github.com/OmidRasouli/vod-streamer-edu/internal/entity"
)

// MockStorage implements port.VideoStorage for testing.
type MockStorage struct {
	SaveFunc    func(reader io.Reader, path ...string) (entity.Path, error)
	OpenFunc    func(path ...string) (io.ReadCloser, error)
	GetPathFunc func(path ...string) (entity.Path, error)
}

func (m *MockStorage) Save(reader io.Reader, path ...string) (entity.Path, error) {
	return m.SaveFunc(reader, path...)
}
func (m *MockStorage) Open(path ...string) (io.ReadCloser, error) {
	return m.OpenFunc(path...)
}
func (m *MockStorage) GetPath(path ...string) (entity.Path, error) {
	return m.GetPathFunc(path...)
}

// MockFFmpeg implements port.FFmpeg for testing.
type MockFFmpeg struct {
	GetVideoDetailsFunc func(path entity.Path) (*model.VideoData, error)
	TranscodeFunc       func(input entity.Path, isPortrait bool) error
}

func (m *MockFFmpeg) GetVideoDetails(path entity.Path) (*model.VideoData, error) {
	return m.GetVideoDetailsFunc(path)
}
func (m *MockFFmpeg) Transcode(input entity.Path, isPortrait bool) error {
	return m.TranscodeFunc(input, isPortrait)
}

func TestProcessAndSave_Success(t *testing.T) {
	mockStorage := &MockStorage{
		SaveFunc: func(reader io.Reader, path ...string) (entity.Path, error) {
			return entity.StringPathToPath("/videos/uuid/video.mp4"), nil
		},
	}
	mockFFmpeg := &MockFFmpeg{
		GetVideoDetailsFunc: func(path entity.Path) (*model.VideoData, error) {
			return &model.VideoData{
				Streams: []model.Stream{{CodecType: "video", Width: 1920, Height: 1080}},
			}, nil
		},
		TranscodeFunc: func(input entity.Path, isPortrait bool) error {
			return nil
		},
	}

	uc := NewVideoUsecase(mockStorage, mockFFmpeg)
	err := uc.ProcessAndSave("video.mp4", bytes.NewBufferString("fake video data"))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestProcessAndSave_SaveError(t *testing.T) {
	mockStorage := &MockStorage{
		SaveFunc: func(reader io.Reader, path ...string) (entity.Path, error) {
			return entity.Path{}, errors.New("save failed")
		},
	}
	mockFFmpeg := &MockFFmpeg{}

	uc := NewVideoUsecase(mockStorage, mockFFmpeg)
	err := uc.ProcessAndSave("video.mp4", bytes.NewBufferString("fake video data"))
	if err == nil {
		t.Error("expected error, got nil")
	}
}
