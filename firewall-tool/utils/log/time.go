package log

import (
	"time"

	"go.uber.org/zap/zapcore"
)

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	emoji := "\U0001F4AA" + " BharatVigil:"
	enc.AppendString(emoji + " " + t.Format(time.RFC3339) + " ")
}