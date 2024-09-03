package genent

import "log/slog"

func LogFn(logger *slog.Logger) func(...any) {
	return func(es ...any) {
		for _, e := range es {
			s, ok := e.(string)
			if !ok {
				continue
			}
			logger.Debug(s)
		}
	}
}