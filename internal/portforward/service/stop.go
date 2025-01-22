package service

import (
	"context"
	"fmt"
	"os"
	"time"
)

func (s *Service) Stop() (err error) {
	s.startStopMutex.Lock()
	defer s.startStopMutex.Unlock()

	s.portMutex.RLock()
	serviceNotRunning := len(s.ports) == 0
	s.portMutex.RUnlock()
	if serviceNotRunning {
		// TODO replace with goservices.ErrAlreadyStopped
		return nil
	}

	s.logger.Info("stopping")

	s.keepPortCancel()
	<-s.keepPortDoneCh

	return s.cleanup()
}

func (s *Service) cleanup() (err error) {
	s.portMutex.Lock()
	defer s.portMutex.Unlock()

	if s.settings.DownCommand != "" {
		const downTimeout = 60 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), downTimeout)
		defer cancel()
		err = runCommand(ctx, s.cmder, s.logger, s.settings.DownCommand, s.ports)
		if err != nil {
			err = fmt.Errorf("running down command: %w", err)
			s.logger.Error(err.Error())
		}
	}

	for _, port := range s.ports {
		err = s.portAllower.RemoveAllowedPort(context.Background(), port)
		if err != nil {
			return fmt.Errorf("blocking previous port in firewall: %w", err)
		}

		if s.settings.ListeningPort != 0 {
			ctx := context.Background()
			const listeningPort = 0 // 0 to clear the redirection
			err = s.portAllower.RedirectPort(ctx, s.settings.Interface, port, listeningPort)
			if err != nil {
				return fmt.Errorf("removing previous port redirection in firewall: %w", err)
			}
		}
	}

	s.ports = nil

	filepath := s.settings.Filepath
	s.logger.Info("removing port file " + filepath)
	err = os.Remove(filepath)
	if err != nil {
		return fmt.Errorf("removing port file: %w", err)
	}

	return nil
}
