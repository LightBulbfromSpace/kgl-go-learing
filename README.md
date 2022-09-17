# golang-learning

Exercises made during the course [Learn go with tests](https://quii.gitbook.io/learn-go-with-tests/).  Names of folders match the chapters of course. Inside folders are contained version of code reflecting progress during each lesson.

Упражнения, сделанные во время прохождения [Learn go with tests](https://quii.gitbook.io/learn-go-with-tests/). Названия папок соответствуют разделам курса. Внутри папок содержатся версии кода, отражающие прогресс в течение урока.

# About me

Student of mathematical and information direction. Interested in backend development.

Студент математического и информационного направления. Заинтересован бэкенд-разработке.

# Learned lessons

### arrays_and_slices/sum - 100.0%

<details>
  <summary><code>func Sum(numbers []int) int</code></summary>

    Sum calculates the total from a slice of numbers.
</details>

### arrays_and_slices/sumAll - 100.0%

<details>
  <summary><code>func SumAll(nums ...[]int) (sums []int)</code></summary>

    SumAll calculates the sums in all given slices.
</details>

### arrays_and_slices/sumTails - 100.0%

<details>
  <summary><code>func SumTails(numsToSum ...[]int) []int</code></summary>

    SumTails calculates the sums of all but the first number given a collection
    of slices.
</details>

### concurrency/faster_version_v2 - 100.0%

<details>
  <summary><code>func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool</code></summary>

    TYPES
    type WebsiteChecker func(string) bool
</details>

### concurrency/slow_version_v1 - 100.0%

<details>
  <summary><code>func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool</code></summary>

    TYPES
    type WebsiteChecker func(string) bool
</details>

### context/v1 - 100.0%

<details>
  <summary><code>func Server(store Store) http.HandlerFunc</code></summary>

    TYPES
    type Store interface {
    Fetch() string
    Cancel()
    }
    type StubStore struct {
    // Has unexported fields.
    }
</details>

<details>
  <summary><code>func (s *StubStore) Cancel()</code></summary>

</details>

<details>
  <summary><code>func (s *StubStore) Fetch() string</code></summary>

</details>

### context/v2 - 100.0%

<details>
  <summary><code>func Server(store Store) http.HandlerFunc</code></summary>

    TYPES
    type Store interface {
    Fetch(ctx context.Context) (string, error)
    }
</details>

### dependency_injection/greet - 50.0%

<details>
  <summary><code>func Greet(writer io.Writer, name string)</code></summary>

</details>

### generics - 71.4%

<details>
  <summary><code>func AssertEqual[T comparable](t *testing.T, got, want T)</code></summary>

</details>

<details>
  <summary><code>func AssertFalse(t *testing.T, got bool)</code></summary>

</details>

<details>
  <summary><code>func AssertNotEqual[T comparable](t *testing.T, got, want T)</code></summary>

</details>

<details>
  <summary><code>func AssertTrue(t *testing.T, got bool)</code></summary>

    TYPES
    type Stack[T any] struct {
    // Has unexported fields.
    }
</details>

<details>
  <summary><code>func (s *Stack[T]) IsEmpty() bool</code></summary>

    IsEmpty return true if stack is empty
</details>

<details>
  <summary><code>func (s *Stack[T]) Pop() (T, bool)</code></summary>

    Pop removes the last value in stack, returns as the first parameter value of
    removed item and false as the second parameter if stack is empty
</details>

<details>
  <summary><code>func (s *Stack[T]) Push(value T)</code></summary>

    Push appends value to the end
    type StackOfInts = Stack[int]
    type StackOfStrings = Stack[string]
</details>

### hello - 92.9%

<details>
  <summary><code>func DefaultReceiver(language string) (receiver string)</code></summary>

</details>

<details>
  <summary><code>func GreetingPrefix(language string) (prefix string)</code></summary>

</details>

<details>
  <summary><code>func Hello(name string, language string) string</code></summary>

</details>

### hellogo - 100.0%
Package hellogo contains first steps in language.
<details>
  <summary><code>func Hello() string</code></summary>

    Hello is first function.
</details>

### integers - 100.0%

<details>
  <summary><code>func Add(a, b int) int</code></summary>

</details>

### intro_to_propely_based_tests/v1 - 100.0%

<details>
  <summary><code>func ConvertToRoman(arabic int) string</code></summary>

</details>

### intro_to_propely_based_tests/v2 - 93.3%

<details>
  <summary><code>func ConvertToArabic(roman string) int</code></summary>

</details>

<details>
  <summary><code>func ConvertToRoman(arabic int) string</code></summary>

    TYPES
    type RomanNum struct {
    // Has unexported fields.
    }
</details>

### intro_to_propely_based_tests/v3_to_arabic_new - 96.6%
VARIABLES

var ErrArabicOverflow = errors.New("unacceptable number (more than 3999)")
<details>
  <summary><code>func ConvertToArabic(roman string) (result uint16)</code></summary>

</details>

<details>
  <summary><code>func ConvertToRoman(arabic uint16) (string, error)</code></summary>

    TYPES
    type RomanNumeral struct {
    // Has unexported fields.
    }
    type RomanNumerals []RomanNumeral
</details>

<details>
  <summary><code>func (r RomanNumerals) Exists(symbols ...byte) bool</code></summary>

</details>

<details>
  <summary><code>func (r RomanNumerals) ValueOf(symbols ...byte) uint16</code></summary>

</details>

### iteration - 100.0%

<details>
  <summary><code>func Repeat(str string, num int) string</code></summary>

</details>

### maps - 83.3%
CONSTANTS

const (
	ErrNotFound         = DictionaryErr("can't find the word")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it doesn't exist")
)

TYPES

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) error

func (d Dictionary) Delete(word string)

func (d Dictionary) Search(word string) (string, error)

func (d Dictionary) Update(word, newDefinition string) error

type DictionaryErr string

func (e DictionaryErr) Error() string
### math/v1 - 100.0%
TYPES

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point
### math/v2 - 100.0%

<details>
  <summary><code>func SVGWriter(w io.Writer, t time.Time)</code></summary>

</details>

<details>
  <summary><code>func SecondHand(w io.Writer, t time.Time)</code></summary>

    TYPES
    type Point struct {
    X float64
    Y float64
    }
</details>

### math/v3 - 100.0%

<details>
  <summary><code>func SVGWriter(w io.Writer, t time.Time)</code></summary>

</details>

<details>
  <summary><code>func SecondHand(w io.Writer, t time.Time)</code></summary>

    TYPES
    type Point struct {
    X float64
    Y float64
    }
</details>

### math/v4 - 100.0%

<details>
  <summary><code>func SVGWriter(w io.Writer, t time.Time)</code></summary>

    TYPES
    type Point struct {
    X float64
    Y float64
    }
    A Point represents a two dimensional Cartesian coordinate.
</details>

### mocking/configurable_sleeper - 81.8%

<details>
  <summary><code>func Countdown(out io.Writer, slp Sleeper)</code></summary>

    TYPES
    type ConfigurableSleeper struct {
    // Has unexported fields.
    }
</details>

<details>
  <summary><code>func (c *ConfigurableSleeper) Sleep()</code></summary>

    type Sleeper interface {
    Sleep()
    }
    type SpyCountdownOperations struct {
    Calls []string
    }
</details>

<details>
  <summary><code>func (s *SpyCountdownOperations) Sleep()</code></summary>

</details>

<details>
  <summary><code>func (s *SpyCountdownOperations) Write(p []byte) (n int, err error)</code></summary>

    type SpyTime struct {
    // Has unexported fields.
    }
</details>

<details>
  <summary><code>func (st *SpyTime) Sleep(duration time.Duration)</code></summary>

</details>

### mocking/countdown_v1 - 62.5%

<details>
  <summary><code>func Countdown(out io.Writer, slp Sleeper)</code></summary>

    TYPES
    type RealSleeper struct{}
</details>

<details>
  <summary><code>func (rs *RealSleeper) Sleep()</code></summary>

    type Sleeper interface {
    Sleep()
    }
    type SpySleeper struct {
    Calls int
    }
</details>

<details>
  <summary><code>func (s *SpySleeper) Sleep()</code></summary>

</details>

### pointers_and_errors - 83.3%
VARIABLES

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
var ErrNegativeValue = errors.New("amount of ETH must be positive")

TYPES

type Ethereum float64

func (e Ethereum) String() string

type Stringer interface {
	String() string
}

type Wallet struct {
	// Has unexported fields.
}

func (w *Wallet) Balance() Ethereum

func (w *Wallet) Deposit(sum Ethereum) error

func (w *Wallet) Withdraw(sum Ethereum) error
### reading_files/v1 - 43.8%
TYPES

type Post struct {
	Title string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error)
### reading_files/v2 - 84.2%
TYPES

type Post struct {
	Title string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error)
### reading_files/v3 - 88.9%
TYPES

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPost(postFile io.Reader) (Post, error)
    NewPost parses File contents code

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error)
### reading_files/v4 - 94.6%
VARIABLES

var MetaSeparators = []Separator{
	titleSeparator,
	descriptionSeparator,
	tagsSeparator,
}

TYPES

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPost(postFile io.Reader) (Post, error)
    NewPost parses File contents code

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error)

type Separator string
### reflection/v1 - 100.0%

### reflection/v2 - 100.0%

### reflection/v3 - 100.0%

### revisiting_arrays_and_slices_with_generics - 74.4%

<details>
  <summary><code>func AssertEqual[T comparable](t *testing.T, got, want T)</code></summary>

</details>

<details>
  <summary><code>func AssertFalse(t *testing.T, got bool)</code></summary>

</details>

<details>
  <summary><code>func AssertNotEqual[T comparable](t *testing.T, got, want T)</code></summary>

</details>

<details>
  <summary><code>func AssertTrue(t *testing.T, got bool)</code></summary>

</details>

<details>
  <summary><code>func BalanceFor(transactions []Transaction, name string) float64</code></summary>

</details>

<details>
  <summary><code>func Find[A any](collection []A, parameterFunc func(A) bool) (value A, found bool)</code></summary>

</details>

<details>
  <summary><code>func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B</code></summary>

</details>

<details>
  <summary><code>func Sum(numbers []int) int</code></summary>

    Sum calculates the total from a slice of numbers.
</details>

<details>
  <summary><code>func SumAll(nums ...[]int) []int</code></summary>

    SumAll calculates the sums in all given slices.
</details>

<details>
  <summary><code>func SumAllTails(numsToSum ...[]int) []int</code></summary>

    SumAllTails calculates the sums of all but the first number given a
    collection of slices.
    TYPES
    type Account struct {
    Name    string
    Balance float64
    }
</details>

<details>
  <summary><code>func NewBalanceFor(account Account, transactions []Transaction) Account</code></summary>

    type Color struct {
    R, G, B int64
    A       float32
    }
</details>

<details>
  <summary><code>func ColorMixer(colors ...Color) (Color, error)</code></summary>

</details>

<details>
  <summary><code>func NewColor(RGB string, A float32) (Color, error)</code></summary>

    type Transaction struct {
    From   string
    To     string
    Amount float64
    }
</details>

<details>
  <summary><code>func NewTransaction(from, to Account, amount float64) Transaction</code></summary>

</details>

### select/concurrency_check - 90.0%

<details>
  <summary><code>func ConfigurableWebsiteRacer(a, b string, timeout time.Duration) (winner string, err error)</code></summary>

</details>

<details>
  <summary><code>func WebsiteRacer(a, b string) (winner string, err error)</code></summary>

</details>

### select/server_mocking - 87.5%

<details>
  <summary><code>func WebsiteRacer(a, b string) string</code></summary>

</details>

### select/test_external_sites - 88.9%

<details>
  <summary><code>func WebsiteRacer(a, b string) (winner string)</code></summary>

</details>

### structs_methods_interfaces - 100.0%
TYPES

type Circle struct {
	// Has unexported fields.
}

func (c Circle) Area() float64

func (c Circle) Perimeter() float64

type Rectangle struct {
	// Has unexported fields.
}

func (rect Rectangle) Area() float64

func (rect Rectangle) Perimeter() float64

type RightTriangle struct {
	// Has unexported fields.
}

func (t RightTriangle) Area() float64

func (t RightTriangle) Perimeter() float64

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	// Has unexported fields.
}

func (s Square) Area() float64

func (s Square) Perimeter() float64
### sync - 100.0%
Package sync provides basic synchronization primitives such as mutual
exclusion locks. Other than the Once and WaitGroup types, most are intended
for use by low-level library routines. Higher-level synchronization is
better done via channels and communication.

Values containing the types defined in this package should not be copied.

TYPES

type Cond struct {

	// L is held while observing or changing the condition
	L Locker

	// Has unexported fields.
}
    Cond implements a condition variable, a rendezvous point for goroutines
    waiting for or announcing the occurrence of an event.

    Each Cond has an associated Locker L (often a *Mutex or *RWMutex), which
    must be held when changing the condition and when calling the Wait method.

    A Cond must not be copied after first use.

func NewCond(l Locker) *Cond
    NewCond returns a new Cond with Locker l.

func (c *Cond) Broadcast()
    Broadcast wakes all goroutines waiting on c.

    It is allowed but not required for the caller to hold c.L during the call.

func (c *Cond) Signal()
    Signal wakes one goroutine waiting on c, if there is any.

    It is allowed but not required for the caller to hold c.L during the call.

func (c *Cond) Wait()
    Wait atomically unlocks c.L and suspends execution of the calling goroutine.
    After later resuming execution, Wait locks c.L before returning. Unlike in
    other systems, Wait cannot return unless awoken by Broadcast or Signal.

    Because c.L is not locked when Wait first resumes, the caller typically
    cannot assume that the condition is true when Wait returns. Instead, the
    caller should Wait in a loop:

        c.L.Lock()
        for !condition() {
            c.Wait()
        }
        ... make use of condition ...
        c.L.Unlock()

type Locker interface {
	Lock()
	Unlock()
}
    A Locker represents an object that can be locked and unlocked.

type Map struct {
	// Has unexported fields.
}
    Map is like a Go map[interface{}]interface{} but is safe for concurrent use
    by multiple goroutines without additional locking or coordination. Loads,
    stores, and deletes run in amortized constant time.

    The Map type is specialized. Most code should use a plain Go map instead,
    with separate locking or coordination, for better type safety and to make it
    easier to maintain other invariants along with the map content.

    The Map type is optimized for two common use cases: (1) when the entry for a
    given key is only ever written once but read many times, as in caches that
    only grow, or (2) when multiple goroutines read, write, and overwrite
    entries for disjoint sets of keys. In these two cases, use of a Map may
    significantly reduce lock contention compared to a Go map paired with a
    separate Mutex or RWMutex.

    The zero Map is empty and ready for use. A Map must not be copied after
    first use.

func (m *Map) Delete(key any)
    Delete deletes the value for a key.

func (m *Map) Load(key any) (value any, ok bool)
    Load returns the value stored in the map for a key, or nil if no value is
    present. The ok result indicates whether value was found in the map.

func (m *Map) LoadAndDelete(key any) (value any, loaded bool)
    LoadAndDelete deletes the value for a key, returning the previous value if
    any. The loaded result reports whether the key was present.

func (m *Map) LoadOrStore(key, value any) (actual any, loaded bool)
    LoadOrStore returns the existing value for the key if present. Otherwise, it
    stores and returns the given value. The loaded result is true if the value
    was loaded, false if stored.

func (m *Map) Range(f func(key, value any) bool)
    Range calls f sequentially for each key and value present in the map. If f
    returns false, range stops the iteration.

    Range does not necessarily correspond to any consistent snapshot of the
    Map's contents: no key will be visited more than once, but if the value for
    any key is stored or deleted concurrently (including by f), Range may
    reflect any mapping for that key from any point during the Range call. Range
    does not block other methods on the receiver; even f itself may call any
    method on m.

    Range may be O(N) with the number of elements in the map even if f returns
    false after a constant number of calls.

func (m *Map) Store(key, value any)
    Store sets the value for a key.

type Mutex struct {
	// Has unexported fields.
}
    A Mutex is a mutual exclusion lock. The zero value for a Mutex is an
    unlocked mutex.

    A Mutex must not be copied after first use.

func (m *Mutex) Lock()
    Lock locks m. If the lock is already in use, the calling goroutine blocks
    until the mutex is available.

func (m *Mutex) TryLock() bool
    TryLock tries to lock m and reports whether it succeeded.

    Note that while correct uses of TryLock do exist, they are rare, and use of
    TryLock is often a sign of a deeper problem in a particular use of mutexes.

func (m *Mutex) Unlock()
    Unlock unlocks m. It is a run-time error if m is not locked on entry to
    Unlock.

    A locked Mutex is not associated with a particular goroutine. It is allowed
    for one goroutine to lock a Mutex and then arrange for another goroutine to
    unlock it.

type Once struct {
	// Has unexported fields.
}
    Once is an object that will perform exactly one action.

    A Once must not be copied after first use.

func (o *Once) Do(f func())
    Do calls the function f if and only if Do is being called for the first time
    for this instance of Once. In other words, given

        var once Once

    if once.Do(f) is called multiple times, only the first call will invoke f,
    even if f has a different value in each invocation. A new instance of Once
    is required for each function to execute.

    Do is intended for initialization that must be run exactly once. Since f is
    niladic, it may be necessary to use a function literal to capture the
    arguments to a function to be invoked by Do:

        config.once.Do(func() { config.init(filename) })

    Because no call to Do returns until the one call to f returns, if f causes
    Do to be called, it will deadlock.

    If f panics, Do considers it to have returned; future calls of Do return
    without calling f.

type Pool struct {

	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
	New func() any
	// Has unexported fields.
}
    A Pool is a set of temporary objects that may be individually saved and
    retrieved.

    Any item stored in the Pool may be removed automatically at any time without
    notification. If the Pool holds the only reference when this happens, the
    item might be deallocated.

    A Pool is safe for use by multiple goroutines simultaneously.

    Pool's purpose is to cache allocated but unused items for later reuse,
    relieving pressure on the garbage collector. That is, it makes it easy to
    build efficient, thread-safe free lists. However, it is not suitable for all
    free lists.

    An appropriate use of a Pool is to manage a group of temporary items
    silently shared among and potentially reused by concurrent independent
    clients of a package. Pool provides a way to amortize allocation overhead
    across many clients.

    An example of good use of a Pool is in the fmt package, which maintains a
    dynamically-sized store of temporary output buffers. The store scales under
    load (when many goroutines are actively printing) and shrinks when
    quiescent.

    On the other hand, a free list maintained as part of a short-lived object is
    not a suitable use for a Pool, since the overhead does not amortize well in
    that scenario. It is more efficient to have such objects implement their own
    free list.

    A Pool must not be copied after first use.

func (p *Pool) Get() any
    Get selects an arbitrary item from the Pool, removes it from the Pool, and
    returns it to the caller. Get may choose to ignore the pool and treat it as
    empty. Callers should not assume any relation between values passed to Put
    and the values returned by Get.

    If Get would otherwise return nil and p.New is non-nil, Get returns the
    result of calling p.New.

func (p *Pool) Put(x any)
    Put adds x to the pool.

type RWMutex struct {
	// Has unexported fields.
}
    A RWMutex is a reader/writer mutual exclusion lock. The lock can be held by
    an arbitrary number of readers or a single writer. The zero value for a
    RWMutex is an unlocked mutex.

    A RWMutex must not be copied after first use.

    If a goroutine holds a RWMutex for reading and another goroutine might call
    Lock, no goroutine should expect to be able to acquire a read lock until the
    initial read lock is released. In particular, this prohibits recursive read
    locking. This is to ensure that the lock eventually becomes available; a
    blocked Lock call excludes new readers from acquiring the lock.

func (rw *RWMutex) Lock()
    Lock locks rw for writing. If the lock is already locked for reading or
    writing, Lock blocks until the lock is available.

func (rw *RWMutex) RLock()
    RLock locks rw for reading.

    It should not be used for recursive read locking; a blocked Lock call
    excludes new readers from acquiring the lock. See the documentation on the
    RWMutex type.

func (rw *RWMutex) RLocker() Locker
    RLocker returns a Locker interface that implements the Lock and Unlock
    methods by calling rw.RLock and rw.RUnlock.

func (rw *RWMutex) RUnlock()
    RUnlock undoes a single RLock call; it does not affect other simultaneous
    readers. It is a run-time error if rw is not locked for reading on entry to
    RUnlock.

func (rw *RWMutex) TryLock() bool
    TryLock tries to lock rw for writing and reports whether it succeeded.

    Note that while correct uses of TryLock do exist, they are rare, and use of
    TryLock is often a sign of a deeper problem in a particular use of mutexes.

func (rw *RWMutex) TryRLock() bool
    TryRLock tries to lock rw for reading and reports whether it succeeded.

    Note that while correct uses of TryRLock do exist, they are rare, and use of
    TryRLock is often a sign of a deeper problem in a particular use of mutexes.

func (rw *RWMutex) Unlock()
    Unlock unlocks rw for writing. It is a run-time error if rw is not locked
    for writing on entry to Unlock.

    As with Mutexes, a locked RWMutex is not associated with a particular
    goroutine. One goroutine may RLock (Lock) a RWMutex and then arrange for
    another goroutine to RUnlock (Unlock) it.

type WaitGroup struct {
	// Has unexported fields.
}
    A WaitGroup waits for a collection of goroutines to finish. The main
    goroutine calls Add to set the number of goroutines to wait for. Then each
    of the goroutines runs and calls Done when finished. At the same time, Wait
    can be used to block until all goroutines have finished.

    A WaitGroup must not be copied after first use.

func (wg *WaitGroup) Add(delta int)
    Add adds delta, which may be negative, to the WaitGroup counter. If the
    counter becomes zero, all goroutines blocked on Wait are released. If the
    counter goes negative, Add panics.

    Note that calls with a positive delta that occur when the counter is zero
    must happen before a Wait. Calls with a negative delta, or calls with a
    positive delta that start when the counter is greater than zero, may happen
    at any time. Typically this means the calls to Add should execute before the
    statement creating the goroutine or other event to be waited for. If a
    WaitGroup is reused to wait for several independent sets of events, new Add
    calls must happen after all previous Wait calls have returned. See the
    WaitGroup example.

func (wg *WaitGroup) Done()
    Done decrements the WaitGroup counter by one.

func (wg *WaitGroup) Wait()
    Wait blocks until the WaitGroup counter is zero.