package common

import "github.com/quartz-technology/agate/indexer/storage/store/dto"

// DataPreprocessorOutput is the default data structure produced as the output of the
// pre-process task.
type DataPreprocessorOutput = struct {
	Output []*PreprocessedRelayData
}

type PreprocessedRelayData struct {
	Bid         *dto.Bid
	Submissions []*dto.Submission
}
