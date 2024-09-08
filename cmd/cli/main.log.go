package cli

import (
	"os"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig();
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder;
	encodeConfig.TimeKey = "Time";
	encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder;
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder;
	return zapcore.NewJSONEncoder(encodeConfig);
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile()
}