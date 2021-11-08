// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
)

func (s AnswerShippingQueryPostReqApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.ShippingOptions {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "shipping_options",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CallbackQuery) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Message // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "message",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s Chat) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Type // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.PinnedMessage // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "pinned_message",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s EncryptedPassportElement) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Type // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s Game) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Photo == nil {
			return errors.New("required, can't be nil")
		}
		_ = s.Photo // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "photo",
			Error: err,
		})
	}
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.TextEntities {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "text_entities",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s GetGameHighScoresPostOK) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Result == nil {
			return errors.New("required, can't be nil")
		}
		_ = s.Result // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "result",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s GetMyCommandsPostOK) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Result == nil {
			return errors.New("required, can't be nil")
		}
		_ = s.Result // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "result",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s GetStickerSetPostOK) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Result.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "result",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s GetUpdatesPostOK) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Result == nil {
			return errors.New("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.Result {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "result",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s GetUpdatesPostReqApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Limit // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "limit",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s GetUserProfilePhotosPostOK) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Result.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "result",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s GetUserProfilePhotosPostReqApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Limit // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "limit",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s InlineKeyboardMarkup) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.InlineKeyboard == nil {
			return errors.New("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.InlineKeyboard {
			if err := func() error {
				if elem == nil {
					return errors.New("required, can't be nil")
				}
				_ = elem // validation expected, but not supported
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "inline_keyboard",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s MaskPosition) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Point // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "point",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s Message) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.SenderChat // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sender_chat",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Chat.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "chat",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.ForwardFromChat // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "forward_from_chat",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.ReplyToMessage // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "reply_to_message",
			Error: err,
		})
	}
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Entities {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "entities",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Sticker // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sticker",
			Error: err,
		})
	}
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.CaptionEntities {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "caption_entities",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Game // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "game",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Poll // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "poll",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.PinnedMessage // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "pinned_message",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.PassportData // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "passport_data",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.ReplyMarkup // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "reply_markup",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s MessageEntity) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Type // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s PassportData) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Data == nil {
			return errors.New("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.Data {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "data",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s Poll) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Options == nil {
			return errors.New("required, can't be nil")
		}
		_ = s.Options // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "options",
			Error: err,
		})
	}
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.ExplanationEntities {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "explanation_entities",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s PollAnswer) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.OptionIds == nil {
			return errors.New("required, can't be nil")
		}
		_ = s.OptionIds // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "option_ids",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s SendGamePostOK) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Result.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "result",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s SendGamePostReqApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.ReplyMarkup // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "reply_markup",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s SendInvoicePostOK) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Result.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "result",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s SendInvoicePostReqApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Prices == nil {
			return errors.New("required, can't be nil")
		}
		_ = s.Prices // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "prices",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.ReplyMarkup // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "reply_markup",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s SetMyCommandsPostReqApplicationJSON) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Commands == nil {
			return errors.New("required, can't be nil")
		}
		_ = s.Commands // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "commands",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ShippingOption) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Prices == nil {
			return errors.New("required, can't be nil")
		}
		_ = s.Prices // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "prices",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s Sticker) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.MaskPosition // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "mask_position",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s StickerSet) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Stickers == nil {
			return errors.New("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.Stickers {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "stickers",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s Update) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Message // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "message",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.EditedMessage // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "edited_message",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.ChannelPost // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "channel_post",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.EditedChannelPost // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "edited_channel_post",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.CallbackQuery // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "callback_query",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Poll // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "poll",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.PollAnswer // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "poll_answer",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s UserProfilePhotos) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Photos == nil {
			return errors.New("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.Photos {
			if err := func() error {
				if elem == nil {
					return errors.New("required, can't be nil")
				}
				_ = elem // validation expected, but not supported
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "photos",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
