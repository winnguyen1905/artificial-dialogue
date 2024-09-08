package cli

import (
	"os"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core)

	logger.Info("Info Log", zap.Int("line", 1))
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig();
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder;
	encodeConfig.TimeKey = "Time";
	encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder;
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder;
	return zapcore.NewJSONEncoder(encodeConfig);
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm);
	syncFile := zapcore.AddSync(file);
	syncConsole := zapcore.AddSync(os.Stderr);
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile);
}