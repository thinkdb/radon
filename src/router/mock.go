/*
 * Radon
 *
 * Copyright 2018 The Radon Authors.
 * Code is licensed under the GPLv3.
 *
 */

package router

import (
	"github.com/thinkdb/radon/src/config"
	"os"

	"github.com/xelabs/go-mysqlstack/xlog"
)

var (
	// MockDefaultConfig config.
	MockDefaultConfig = []*config.PartitionConfig{
		&config.PartitionConfig{
			Table:   "A2",
			Segment: "2-4",
			Backend: "backend2",
		},
		&config.PartitionConfig{
			Table:   "A4",
			Segment: "4-8",
			Backend: "backend4",
		},
	}
)

// MockTableAConfig config.
func MockTableAConfig() *config.TableConfig {
	mock := &config.TableConfig{
		Name:       "A",
		ShardType:  "HASH",
		ShardKey:   "id",
		Partitions: MockDefaultConfig,
	}
	S02 := &config.PartitionConfig{
		Table:   "A0",
		Segment: "0-2",
		Backend: "backend0",
	}
	S81024 := &config.PartitionConfig{
		Table:   "A8",
		Segment: "8-4096",
		Backend: "backend8",
	}

	mock.Partitions = append(mock.Partitions, S02, S81024)
	return mock
}

// MockTableMConfig config.
func MockTableMConfig() *config.TableConfig {
	mock := &config.TableConfig{
		Name:       "A",
		ShardType:  "HASH",
		ShardKey:   "id",
		Partitions: make([]*config.PartitionConfig, 0, 16),
	}

	S032 := &config.PartitionConfig{
		Table:   "A1",
		Segment: "0-32",
		Backend: "backend1",
	}

	S3264 := &config.PartitionConfig{
		Table:   "A2",
		Segment: "32-64",
		Backend: "backend2",
	}

	S6496 := &config.PartitionConfig{
		Table:   "A3",
		Segment: "64-96",
		Backend: "backend3",
	}

	S96256 := &config.PartitionConfig{
		Table:   "A4",
		Segment: "96-256",
		Backend: "backend4",
	}

	S256512 := &config.PartitionConfig{
		Table:   "A5",
		Segment: "256-512",
		Backend: "backend5",
	}

	S5121024 := &config.PartitionConfig{
		Table:   "A6",
		Segment: "512-4096",
		Backend: "backend6",
	}

	mock.Partitions = append(mock.Partitions, S032, S3264, S6496, S96256, S256512, S5121024)
	return mock
}

// MockTableBConfig config.
func MockTableBConfig() *config.TableConfig {
	mock := &config.TableConfig{
		Name:       "B",
		ShardType:  "HASH",
		ShardKey:   "id",
		Partitions: make([]*config.PartitionConfig, 0, 16),
	}
	S0512 := &config.PartitionConfig{
		Table:   "B0",
		Segment: "0-512",
		Backend: "backend0",
	}
	S11024 := &config.PartitionConfig{
		Table:   "B1",
		Segment: "512-4096",
		Backend: "backend512",
	}

	mock.Partitions = append(mock.Partitions, S0512, S11024)
	return mock
}

// MockTableNULLConfig config.
func MockTableNULLConfig() *config.TableConfig {
	mock := &config.TableConfig{
		Name:      "B",
		ShardType: "HASH",
		ShardKey:  "id",
	}
	return mock
}

// MockTableSegmentErr1Config config.
func MockTableSegmentErr1Config() *config.TableConfig {
	mock := &config.TableConfig{
		Name:       "A",
		ShardType:  "HASH",
		ShardKey:   "id",
		Partitions: make([]*config.PartitionConfig, 0, 16),
	}

	S032 := &config.PartitionConfig{
		Table:   "A1",
		Segment: "0",
		Backend: "backend1",
	}
	mock.Partitions = append(mock.Partitions, S032)
	return mock
}

// MockTableSegmentStartErrConfig config.
func MockTableSegmentStartErrConfig() *config.TableConfig {
	mock := &config.TableConfig{
		Name:       "A",
		ShardType:  "HASH",
		ShardKey:   "id",
		Partitions: make([]*config.PartitionConfig, 0, 16),
	}

	S032 := &config.PartitionConfig{
		Table:   "A1",
		Segment: "x-0",
		Backend: "backend1",
	}
	mock.Partitions = append(mock.Partitions, S032)
	return mock
}

// MockTableSegmentEndErrConfig config.
func MockTableSegmentEndErrConfig() *config.TableConfig {
	mock := &config.TableConfig{
		Name:       "A",
		ShardType:  "HASH",
		ShardKey:   "id",
		Partitions: make([]*config.PartitionConfig, 0, 16),
	}

	S032 := &config.PartitionConfig{
		Table:   "A1",
		Segment: "0-x",
		Backend: "backend1",
	}
	mock.Partitions = append(mock.Partitions, S032)
	return mock
}

// MockTable64Config config.
func MockTable64Config() *config.TableConfig {
	mock := &config.TableConfig{
		Name:       "A",
		ShardKey:   "id",
		Partitions: MockDefaultConfig,
	}
	S02 := &config.PartitionConfig{
		Table:   "A1",
		Segment: "0-2",
		Backend: "backend1",
	}
	S864 := &config.PartitionConfig{
		Table:   "A4",
		Segment: "8-64",
		Backend: "backend2",
	}

	mock.Partitions = append(mock.Partitions, S02, S864)
	return mock
}

// MockTableOverlapConfig config.
func MockTableOverlapConfig() *config.TableConfig {
	mock := &config.TableConfig{
		Name:       "A",
		ShardKey:   "id",
		Partitions: MockDefaultConfig,
	}

	S79 := &config.PartitionConfig{
		Table:   "A1",
		Segment: "7-9",
		Backend: "backend1",
	}

	mock.Partitions = append(mock.Partitions, S79)
	return mock
}

// MockTableInvalidConfig config.
func MockTableInvalidConfig() *config.TableConfig {
	mock := &config.TableConfig{
		Name:       "A",
		ShardKey:   "id",
		Partitions: MockDefaultConfig,
	}

	S8X := &config.PartitionConfig{
		Table:   "A1",
		Segment: "8-x",
		Backend: "backend1",
	}

	mock.Partitions = append(mock.Partitions, S8X)
	return mock
}

// MockTableGreaterThanConfig config.
func MockTableGreaterThanConfig() *config.TableConfig {
	mock := &config.TableConfig{
		Name:       "A",
		ShardKey:   "id",
		Partitions: MockDefaultConfig,
	}

	S108 := &config.PartitionConfig{
		Table:   "A1",
		Segment: "10-8",
		Backend: "backend1",
	}

	mock.Partitions = append(mock.Partitions, S108)
	return mock
}

// MockTableE1Config config, unsupport shardtype.
func MockTableE1Config() *config.TableConfig {
	mock := &config.TableConfig{
		Name:       "E1",
		ShardType:  "Range",
		ShardKey:   "id",
		Partitions: MockDefaultConfig,
	}
	S02 := &config.PartitionConfig{
		Table:   "A1",
		Segment: "0-2",
		Backend: "backend1",
	}
	S81024 := &config.PartitionConfig{
		Table:   "A4",
		Segment: "8-4096",
		Backend: "backend2",
	}

	mock.Partitions = append(mock.Partitions, S02, S81024)
	return mock
}

var (
	_mockRouterSchemaDir = "/tmp/router_test"
)

// MockNewRouter mocks router.
func MockNewRouter(log *xlog.Log) (*Router, func()) {
	return NewRouter(log, _mockRouterSchemaDir, config.DefaultRouterConfig()), func() {
		if err := os.RemoveAll(_mockRouterSchemaDir); err != nil {
			panic(err)
		}
	}
}
