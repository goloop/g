package g

import (
	"math/rand"
	"runtime"
	"time"
)

// Pair is a generic struct type with two fields: First and Second.
// Used as the result element of the Zip function and can contain
// a pair of values of any T and U types.
type Pair[T, U any] struct {
	First  T
	Second U
}

// MaxRangeSize is the maximum size for range generation
// by the Range function.
const MaxRangeSize = 100_000_000

var (
	// The randomGenerator is a global variable that is used by
	// the Random function to generate random numbers.
	// It is initialized in the init() function.
	randomGenerator *rand.Rand

	// The parallelTasks the number of parallel tasks.
	parallelTasks = 1

	// The maxParallelTasks is the maximum number of parallel tasks.
	maxParallelTasks = runtime.NumCPU() * 3

	// The minLoadPerGoroutine is the minimum slice size for processing
	// in an individual goroutine. Essentially, it delineates the threshold
	// at which it becomes worthwhile to divide the slice processing amongst
	// multiple goroutines. If each goroutine isn't handling a sufficiently
	// large subslice, the overhead of goroutine creation and management
	// may outweigh the benefits of concurrent processing. This variable
	// specifies the minimum number of iterations per goroutine to ensure
	// an efficient division of labor.
	minLoadPerGoroutine = 65536
)

// The init initializes the randomGenerator variable.
func init() {
	parallelTasks = runtime.NumCPU() * 2
	randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// ParallelTasks returns the number of parallel tasks.
//
// If the function is called without parameters, it returns the
// current value of parallelTasks.
//
// A function can receive one or more values for parallelTasks,
// these values are added together to form the final result for
// parallelTasks. If the new value for parallelTasks is less than
// or equal to zero - it will be set to 1, if it is greater than
// maxParallelTasks - it will be set to maxParallelTasks.
func ParallelTasks(v ...int) int {
	if len(v) > 0 {
		n := Sum(v...)
		if n <= 0 {
			parallelTasks = 1
		} else if n > maxParallelTasks {
			parallelTasks = maxParallelTasks
		} else {
			parallelTasks = n
		}
	}

	return parallelTasks
}
