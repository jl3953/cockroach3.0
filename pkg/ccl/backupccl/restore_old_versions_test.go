// Copyright 2019 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package backupccl_test

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/base"
	"github.com/cockroachdb/cockroach/pkg/testutils"
	"github.com/cockroachdb/cockroach/pkg/testutils/sqlutils"
	"github.com/cockroachdb/cockroach/pkg/testutils/testcluster"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/stretchr/testify/require"
)

// TestRestoreOldVersions ensures that we can successfully restore tables
// and databases exported by old version.
//
// The files being restored live in testdata and are all made from the same
// input SQL which lives in <testdataBase>/create.sql.
//
// The SSTs were created via the following commands:
//
//  VERSION=...
//  roachprod wipe local
//  roachprod stage local release ${VERSION}
//  roachprod start local
//  # If the version is v1.0.7 then you need to enable enterprise with the
//  # enterprise.enabled cluster setting.
//  roachprod sql local:1 -- -e "$(cat pkg/ccl/backupccl/testdata/restore_old_versions/create.sql)"
//  # Create an S3 bucket to store the backup.
//  roachprod sql local:1 -- -e "BACKUP DATABASE test TO 's3://<bucket-name>/${VERSION}?AWS_ACCESS_KEY_ID=<...>&AWS_SECRET_ACCESS_KEY=<...>'"
//  # Then download the backup from s3 and plop the files into the appropriate
//  # testdata directory.
//
func TestRestoreOldVersions(t *testing.T) {
	defer leaktest.AfterTest(t)()
	const (
		testdataBase = "testdata/restore_old_versions"
		exportDirs   = testdataBase + "/exports"
		clusterDirs  = testdataBase + "/cluster"
	)

	t.Run("table-restore", func(t *testing.T) {
		dirs, err := ioutil.ReadDir(exportDirs)
		require.NoError(t, err)
		for _, dir := range dirs {
			require.True(t, dir.IsDir())
			exportDir, err := filepath.Abs(filepath.Join(exportDirs, dir.Name()))
			require.NoError(t, err)
			t.Run(dir.Name(), restoreOldVersionTest(exportDir))
		}
	})

	t.Run("cluster-restore", func(t *testing.T) {
		dirs, err := ioutil.ReadDir(clusterDirs)
		require.NoError(t, err)
		for _, dir := range dirs {
			require.True(t, dir.IsDir())
			exportDir, err := filepath.Abs(filepath.Join(clusterDirs, dir.Name()))
			require.NoError(t, err)
			t.Run(dir.Name(), restoreOldVersionClusterTest(exportDir))
		}
	})
}

func restoreOldVersionTest(exportDir string) func(t *testing.T) {
	return func(t *testing.T) {
		params := base.TestServerArgs{}
		const numAccounts = 1000
		_, _, sqlDB, dir, cleanup := backupRestoreTestSetupWithParams(t, singleNode, numAccounts,
			initNone, base.TestClusterArgs{ServerArgs: params})
		defer cleanup()
		err := os.Symlink(exportDir, filepath.Join(dir, "foo"))
		require.NoError(t, err)
		sqlDB.Exec(t, `CREATE DATABASE test`)
		var unused string
		var importedRows int
		sqlDB.QueryRow(t, `RESTORE test.* FROM $1`, localFoo).Scan(
			&unused, &unused, &unused, &importedRows, &unused, &unused,
		)
		const totalRows = 12
		if importedRows != totalRows {
			t.Fatalf("expected %d rows, got %d", totalRows, importedRows)
		}
		results := [][]string{
			{"1", "1", "1"},
			{"2", "2", "2"},
			{"3", "3", "3"},
		}
		sqlDB.CheckQueryResults(t, `SELECT * FROM test.t1 ORDER BY k`, results)
		sqlDB.CheckQueryResults(t, `SELECT * FROM test.t2 ORDER BY k`, results)
		sqlDB.CheckQueryResults(t, `SELECT * FROM test.t4 ORDER BY k`, results)
	}
}

func restoreOldVersionClusterTest(exportDir string) func(t *testing.T) {
	return func(t *testing.T) {
		externalDir, dirCleanup := testutils.TempDir(t)
		ctx := context.Background()
		tc := testcluster.StartTestCluster(t, singleNode, base.TestClusterArgs{
			ServerArgs: base.TestServerArgs{
				ExternalIODir: externalDir,
			},
		})
		sqlDB := sqlutils.MakeSQLRunner(tc.Conns[0])
		defer func() {
			tc.Stopper().Stop(ctx)
			dirCleanup()
		}()
		err := os.Symlink(exportDir, filepath.Join(externalDir, "foo"))
		require.NoError(t, err)

		// Ensure that the restore succeeds.
		sqlDB.Exec(t, `RESTORE FROM $1`, localFoo)

		sqlDB.CheckQueryResults(t, "SHOW USERS", [][]string{
			{"admin", "CREATEROLE", "{}"},
			{"craig", "", "{}"},
			{"root", "CREATEROLE", "{admin}"},
		})
		sqlDB.CheckQueryResults(t, "SELECT * FROM system.comments", [][]string{
			{"0", "52", "0", "database comment string"},
			{"1", "53", "0", "table comment string"},
		})
		sqlDB.CheckQueryResults(t, "SELECT * FROM data.bank", [][]string{{"1"}})
	}
}
