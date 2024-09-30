package writer

import (
	"context"
	"alexdenkk/applications/internal/entity"
	"github.com/go-telegram/bot"
)

func (w *Writer) SendMessage(
	ctx context.Context, application *entity.Application,
) {
	w.Bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: w.ChatID,
		Text: application.ToText(),
	})
}
