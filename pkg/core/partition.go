package core

type PartitionKey string

type Partition struct {
	CurrentOffset *Offset
}

func NewPartition() *Partition {
	return &Partition{
		CurrentOffset: nil,
	}
}

// NextOffset retrieves the next offset for this partition.
// It assumes that there is only one producer per offset
// but in the future we should probably use a mutex here
// to avoid race conditions as a precaution.
func (partition *Partition) NextOffset() Offset {
	var newOffset Offset
	if partition.CurrentOffset == nil {
		newOffset = InitialOffset()
	} else {
		newOffset = partition.CurrentOffset.Next()
	}

	partition.CurrentOffset = &newOffset

	return newOffset
}
