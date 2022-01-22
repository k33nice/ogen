// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
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
	_ = bits.LeadingZeros64
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

func encodeAddStickerToSetRequestJSON(req AddStickerToSet, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeAnswerCallbackQueryRequestJSON(req AnswerCallbackQuery, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeAnswerInlineQueryRequestJSON(req AnswerInlineQuery, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeAnswerPreCheckoutQueryRequestJSON(req AnswerPreCheckoutQuery, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeAnswerShippingQueryRequestJSON(req AnswerShippingQuery, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeApproveChatJoinRequestRequestJSON(req ApproveChatJoinRequest, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeBanChatMemberRequestJSON(req BanChatMember, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeCopyMessageRequestJSON(req CopyMessage, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeCreateChatInviteLinkRequestJSON(req CreateChatInviteLink, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeCreateNewStickerSetRequestJSON(req CreateNewStickerSet, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeDeclineChatJoinRequestRequestJSON(req DeclineChatJoinRequest, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeDeleteChatPhotoRequestJSON(req DeleteChatPhoto, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeDeleteChatStickerSetRequestJSON(req DeleteChatStickerSet, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeDeleteMessageRequestJSON(req DeleteMessage, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeDeleteMyCommandsRequestJSON(req DeleteMyCommands, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeDeleteStickerFromSetRequestJSON(req DeleteStickerFromSet, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeDeleteWebhookRequestJSON(req DeleteWebhook, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeEditChatInviteLinkRequestJSON(req EditChatInviteLink, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeEditMessageCaptionRequestJSON(req EditMessageCaption, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeEditMessageLiveLocationRequestJSON(req EditMessageLiveLocation, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeEditMessageMediaRequestJSON(req EditMessageMedia, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeEditMessageReplyMarkupRequestJSON(req EditMessageReplyMarkup, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeEditMessageTextRequestJSON(req EditMessageText, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeExportChatInviteLinkRequestJSON(req ExportChatInviteLink, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeForwardMessageRequestJSON(req ForwardMessage, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeGetChatRequestJSON(req GetChat, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeGetChatAdministratorsRequestJSON(req GetChatAdministrators, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeGetChatMemberRequestJSON(req GetChatMember, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeGetChatMemberCountRequestJSON(req GetChatMemberCount, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeGetFileRequestJSON(req GetFile, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeGetGameHighScoresRequestJSON(req GetGameHighScores, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeGetMyCommandsRequestJSON(req GetMyCommands, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeGetStickerSetRequestJSON(req GetStickerSet, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeGetUpdatesRequestJSON(req GetUpdates, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeGetUserProfilePhotosRequestJSON(req GetUserProfilePhotos, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeLeaveChatRequestJSON(req LeaveChat, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodePinChatMessageRequestJSON(req PinChatMessage, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodePromoteChatMemberRequestJSON(req PromoteChatMember, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeRestrictChatMemberRequestJSON(req RestrictChatMember, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeRevokeChatInviteLinkRequestJSON(req RevokeChatInviteLink, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendAnimationRequestJSON(req SendAnimation, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendAudioRequestJSON(req SendAudio, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendChatActionRequestJSON(req SendChatAction, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendContactRequestJSON(req SendContact, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendDiceRequestJSON(req SendDice, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendDocumentRequestJSON(req SendDocument, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendGameRequestJSON(req SendGame, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendInvoiceRequestJSON(req SendInvoice, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendLocationRequestJSON(req SendLocation, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendMediaGroupRequestJSON(req SendMediaGroup, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendMessageRequestJSON(req SendMessage, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendPhotoRequestJSON(req SendPhoto, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendPollRequestJSON(req SendPoll, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendStickerRequestJSON(req SendSticker, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendVenueRequestJSON(req SendVenue, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendVideoRequestJSON(req SendVideo, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendVideoNoteRequestJSON(req SendVideoNote, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSendVoiceRequestJSON(req SendVoice, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSetChatAdministratorCustomTitleRequestJSON(req SetChatAdministratorCustomTitle, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSetChatDescriptionRequestJSON(req SetChatDescription, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSetChatPermissionsRequestJSON(req SetChatPermissions, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSetChatPhotoRequestJSON(req SetChatPhoto, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSetChatStickerSetRequestJSON(req SetChatStickerSet, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSetChatTitleRequestJSON(req SetChatTitle, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSetGameScoreRequestJSON(req SetGameScore, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSetMyCommandsRequestJSON(req SetMyCommands, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSetPassportDataErrorsRequestJSON(req SetPassportDataErrors, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSetStickerPositionInSetRequestJSON(req SetStickerPositionInSet, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSetStickerSetThumbRequestJSON(req SetStickerSetThumb, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeSetWebhookRequestJSON(req SetWebhook, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeStopMessageLiveLocationRequestJSON(req StopMessageLiveLocation, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeStopPollRequestJSON(req StopPoll, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeUnbanChatMemberRequestJSON(req UnbanChatMember, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeUnpinAllChatMessagesRequestJSON(req UnpinAllChatMessages, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeUnpinChatMessageRequestJSON(req UnpinChatMessage, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}

func encodeUploadStickerFileRequestJSON(req UploadStickerFile, span trace.Span) (data *bytes.Buffer, err error) {
	buf := getBuf()
	e := jx.GetWriter()
	defer jx.PutWriter(e)

	req.Encode(e)
	if _, err := e.WriteTo(buf); err != nil {
		putBuf(buf)
		return nil, err
	}

	return buf, nil
}
