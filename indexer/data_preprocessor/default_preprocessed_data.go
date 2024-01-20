package data_preprocessor

import "github.com/quartz-technology/agate/indexer/storage_manager/store/dto"

// DefaultDataPreprocessorOutput is the default data structure produced as the output of the
// pre-process task.
type DefaultDataPreprocessorOutput = DataPreprocessorOutput[[]*DefaultPreprocessedRelayData]

type DefaultPreprocessedRelayData struct {
	Bid         *dto.Bid
	Submissions []*dto.Submission
}
