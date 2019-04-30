package setup

import (
	"fmt"
	"io"
)

type progressTracker struct {
}

func (p *progressTracker) TrackProgress(src string, currentSize int64, totalSize int64, stream io.ReadCloser) (body io.ReadCloser) {
	fmt.Sprintf("%.2f", currentSize/totalSize)
	return stream
}
