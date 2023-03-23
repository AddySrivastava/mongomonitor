package types

type TCollStats struct {
	Ns         string
	Size       int32
	Timeseries struct {
		BucketsNs                            string
		BucketCount                          string
		AvgBucketSize                        int32
		NumBucketInserts                     int32
		NumBucketUpdates                     int32
		NumBucketsOpenedDueToMetadata        int32
		NumBucketsClosedDueToCount           int32
		NumBucketsClosedDueToSize            int32
		NumBucketsClosedDueToTimeForward     int32
		NumBucketsClosedDueToTimeBackward    int32
		NumBucketsClosedDueToMemoryThreshold int32
		NumCommits                           int32
		NumWaits                             int32
		NumMeasurementsCommitted             int32
		AvgNumMeasurementsPerCommit          int32
	}
	Count           int32
	AvgObjSize      int32
	NumOrphanDocs   int32 // Available starting in MongoDB 6.0
	StorageSize     int32
	FreeStorageSize int32
	Capped          bool
	Max             int32
	MaxSize         int32
	WiredTiger      TWiredTiger
	Nindexes        int32
	IndexDetails    map[string]TWiredTiger
	TotalIndexSize  int32
	TotalSize       int32 // Available starting in MongoDB 4.4
	IndexSizes      map[string]int32
}
