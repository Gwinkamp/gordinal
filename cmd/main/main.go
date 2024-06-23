package main

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/Gwinkamp/gordinal"
	"github.com/Gwinkamp/gordinal/internal/config"
)

func main() {
	cfg := config.MustReadFromFlag()
	logger := config.MustConfigureLogging(cfg.Logging)

	g := gordinal.New()
	g.SetLoger(logger)
	registerHooks(g, cfg.Hooks)
	g.RegisterDefaultExit()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-stop
		g.Stop()
	}()

	g.Run()
}

func registerHooks(g *gordinal.Gordinal, hooks []config.Hook) {
	for _, hook := range hooks {
		g.Register(
			hook.Name,
			hook.Keys,
			func() error {
				if err := exec.Command(hook.Command, hook.Args...).Run(); err != nil {
					return err
				}
				return nil
			},
		)
	}
}
