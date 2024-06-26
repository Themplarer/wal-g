package etcd

import (
	"context"
	"os/exec"

	"github.com/wal-g/wal-g/internal"
	"github.com/wal-g/wal-g/pkg/storages/storage"
)

func HandleBackupFetch(ctx context.Context,
	folder storage.Folder,
	targetBackupSelector internal.BackupSelector,
	restoreCmd *exec.Cmd) {
	internal.HandleBackupFetch(folder, targetBackupSelector, internal.GetBackupToCommandFetcher(restoreCmd))
}
