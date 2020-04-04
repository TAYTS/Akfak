package recordpb

import (
	"AKFAK/utils"
	"time"

	"github.com/golang/protobuf/proto"
)

///////////////////////////////////
// 		   Public Methods		 //
///////////////////////////////////

// InitialiseEmptyRecordBatch return a RecordBatch pointer type with default header
func InitialiseEmptyRecordBatch() *RecordBatch {
	// Create empty Recordbatch
	rcdBatch := &RecordBatch{}

	// Set the default field of the RecordBatch
	rcdBatch.writeEmptyHeader()

	return rcdBatch
}

// InitialiseRecordBatchWithData return a RecordBatch pointer type with header set based on the input data
func InitialiseRecordBatchWithData(records []*Record) *RecordBatch {
	rcdBatch := InitialiseEmptyRecordBatch()

	rcdBatch.Records = records
	rcdBatch.updateBatchLength()

	return rcdBatch
}

// GetBatchLengthInt return the BatchLength field in int format
func (rcdBatch *RecordBatch) GetBatchLengthInt() int {
	return utils.BytesToInt(rcdBatch.GetBatchLength())
}

// AppendRecord add one or more new Record to the Recordbatch
func (rcdBatch *RecordBatch) AppendRecord(records ...*Record) {
	for _, record := range records {
		// update record field relative to the batch
		record.TimestampDelta = int32(getCurrentTimeinMs() - rcdBatch.GetFirstTimestamp())
		record.OffsetDelta = int32(len(rcdBatch.GetRecords()))
		// add record
		rcdBatch.Records = append(rcdBatch.GetRecords(), record)
		// update MaxTimestamp
		rcdBatch.MaxTimestamp = getCurrentTimeinMs()
		// update LastOffsetDelta
		rcdBatch.LastOffsetDelta = int32(len(rcdBatch.GetRecords()) - 1)
	}
	rcdBatch.updateBatchLength()
}

///////////////////////////////////
// 		   Private Methods		 //
///////////////////////////////////

// writeEmptyHeader setup the default fields (BatchLength: []byte of size 4 and Magic: 2) in the RecordBatch
func (rcdBatch *RecordBatch) writeEmptyHeader() {
	// Only update the BatchLength and Magic as the rest are using the default value
	rcdBatch.BatchLength = make([]byte, 4)
	rcdBatch.Magic = 2 // Do not change; originally used for handling different Record version
}

// updateBatchLength compute and update the BatchLength
func (rcdBatch *RecordBatch) updateBatchLength() {
	// Calculate new batch length
	length := proto.Size(rcdBatch)

	// Update BatchLength
	rcdBatch.BatchLength = utils.IntToBytes(length)
}

// getCurrentTimeinMs return current unix time in ms
func getCurrentTimeinMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
