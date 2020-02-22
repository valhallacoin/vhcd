// Copyright (c) 2020 The Valhalla Coin developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package updater

import (
	"context"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/frankbraun/codechain/secpkg"
	"github.com/frankbraun/codechain/util/homedir"
)

const (
	tickerDuration     = 5 * time.Minute
	checkUpdateTimeout = 5 * time.Second
	maxTimeToUpdate    = 1 * time.Hour // TODO, change:
	// maxTimeToUpdate    = 7 * 24 * time.Hour
)

// UpdateManager provides a concurrency safe update manager.
type UpdateManager struct {
	started                int32
	shutdown               int32
	wg                     sync.WaitGroup
	requestProcessShutdown chan struct{}
	quit                   chan struct{}
	ticker                 *time.Ticker
	packageName            string
	needsUpdate            bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewUpdateManager() (*UpdateManager, error) {
	um := UpdateManager{
		requestProcessShutdown: make(chan struct{}),
		quit:                   make(chan struct{}),
		ticker:                 time.NewTicker(tickerDuration),
	}
	return &um, nil
}

// Start begins the core update handler.
func (u *UpdateManager) Start() {
	// Already started?
	if atomic.AddInt32(&u.started, 1) != 1 {
		return
	}

	log.Trace("Starting update manager")
	u.wg.Add(1)
	go u.updateHandler()
}

// Stop gracefully shuts down the update manager by stopping all asynchronous
// handlers and waiting for them to finish.
func (u *UpdateManager) Stop() error {
	if atomic.AddInt32(&u.shutdown, 1) != 1 {
		log.Warnf("Update manager is already in the process of " +
			"shutting down")
		return nil
	}

	log.Infof("Update manager shutting down")
	close(u.quit)
	u.wg.Wait()
	u.ticker.Stop()
	return nil
}

// taken from github.com/frankbraun/codechain/secpkg/update.go
func upToDateIfInstalled(ctx context.Context, name string) error {
	needsUpdate, err := secpkg.CheckUpdate(ctx, name)
	if err != nil {
		if err == secpkg.ErrNotInstalled {
			log.Warnf("Package '%s' not installed via `secpkg install`", name)
			return nil
		}
		return err
	}
	// determine path of currently running executable
	path, err := os.Executable()
	if err != nil {
		return err
	}
	localDir := filepath.Join(homedir.SecPkg(), "local")
	// issue warning if we don't actually run the installed executable
	if !strings.HasPrefix(path, localDir) {
		log.Warnf("Package '%s' installed via `secpkg install`, but running different executable: %s",
			name, path)
	}
	// now report update needs, if necessary
	if needsUpdate {
		return ErrNeedsUpdate
	}
	return nil
}

func (um *UpdateManager) checkUpdate() error {
	// Determine package name once.
	if um.packageName == "" {
		var err error
		um.packageName, err = getPackageName()
		if err != nil {
			return err
		}
		err = os.RemoveAll("." + um.packageName + "-needs-update")
		if err != nil {
			return err
		}
	}

	// Check for update.
	ctx, cancel := context.WithTimeout(context.Background(), checkUpdateTimeout)
	defer cancel()
	if err := upToDateIfInstalled(ctx, um.packageName); err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			log.Warnf("Update check for package '%s' timed out: %s",
				um.packageName, err)
			return nil
		}
		if err == ErrNeedsUpdate {
			return um.triggerUpdate()
		}
		return err
	}
	log.Info("No update needed")
	return nil
}

func (um *UpdateManager) triggerUpdate() error {
	um.needsUpdate = true
	f, err := os.Create("." + um.packageName + "-needs-update")
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	timeToUpdate := time.Duration(rand.Int63n(int64(maxTimeToUpdate)))
	log.Warnf("Exiting for update at %s",
		time.Now().UTC().Add(timeToUpdate).Format(time.RFC3339))

	go func() {
		time.Sleep(timeToUpdate)
		um.stop(ErrNeedsUpdate)
	}()
	return nil
}

func (um *UpdateManager) updateHandler() {
out:
	for {
		select {
		case <-um.quit:
			break out
		case <-um.ticker.C:
			if !um.needsUpdate {
				if err := um.checkUpdate(); err != nil {
					um.stop(err)
				}
			}
		}
	}

	um.wg.Done()
	log.Trace("Update handler done")
}

// RequestedProcessShutdown returns a channel that is sent to when the updater
// client requests the process to shutdown.  If the request can
// not be read immediately, it is dropped.
func (um *UpdateManager) RequestedProcessShutdown() <-chan struct{} {
	return um.requestProcessShutdown
}

// handleStop implements the stop command.
func (um *UpdateManager) stop(err error) {
	select {
	case um.requestProcessShutdown <- struct{}{}:
		log.Error(err)
	default:
	}
}
