package gordinal

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

	gohook "github.com/robotn/gohook"
)

type Gordinal struct {
	log *slog.Logger
	ch  chan gohook.Event
}

// New creates new Gordinal instance
func New(log *slog.Logger) *Gordinal {
	return &Gordinal{
		log: log,
		ch:  nil,
	}
}

// Register registers function hot key hook
func (g *Gordinal) Register(name string, keys []string, hook func() error) {
	log := g.log.With(
		slog.String("operation", fmt.Sprintf("hotkey.%s", name)),
		slog.String("hotkey", strings.Join(keys, "+")),
	)

	gohook.Register(gohook.KeyDown, keys, func(e gohook.Event) {
		startTime := time.Now()
		log.Debug(fmt.Sprintf("start command %s", name))

		if err := hook(); err != nil {
			log.Error("failed command %s: %w", name, err)
		} else {
			log.
				With(slog.Duration("duration", time.Since(startTime))).
				Debug(fmt.Sprintf("complete command %s", name))
		}
	})
}

// RegisterExit registers hot key hook for exit from program
func (g *Gordinal) RegisterExit(keys []string) {
	const operation = "hotkey.exit"

	log := g.log.With(
		slog.String("operation", operation),
		slog.String("hotkey", strings.Join(keys, "+")),
	)

	gohook.Register(gohook.KeyDown, keys, func(e gohook.Event) {
		log.Info("service has completed work by exit hot key")
		gohook.End()
	})
}

// RegisterDefaultExit registers default hot key hook for exit from program
func (g *Gordinal) RegisterDefaultExit() {
	g.RegisterExit([]string{"ctrl", "shift", "q"})
}

// Start starts gordinal service
func (g *Gordinal) Start() {
	const operation = "hotkey.start"

	g.log.
		With(slog.String("operation", operation)).
		Info("gordinal service has started")

	g.ch = gohook.Start()
}

// Stop stops gordinal service
func (g *Gordinal) Stop() {
	const operation = "hotkey.stop"

	g.log.
		With(slog.String("operation", operation)).
		Info("gordinal service has stopped")

	gohook.End()
}

// Wait waits for gordinal service to stop
func (g *Gordinal) Wait() {
	<-gohook.Process(g.ch)
}

// Run starts gordinal service and waits service to stop
func (g *Gordinal) Run() {
	g.Start()
	g.Wait()
}
