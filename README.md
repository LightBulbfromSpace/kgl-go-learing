# golang-learning

Exercises made during the course [Learn go with tests](https://quii.gitbook.io/learn-go-with-tests/).  Names of folders match the chapters of course. Inside folders are contained version of code reflecting progress during each lesson.

Упражнения, сделанные во время прохождения [Learn go with tests](https://quii.gitbook.io/learn-go-with-tests/). Названия папок соответствуют разделам курса. Внутри папок содержатся версии кода, отражающие прогресс в течение урока.

# About me

Student of mathematical and information direction. Interested in backend development.

Студент математического и информационного направления. Заинтересован бэкенд-разработке.

# Learned lessons

### github.com/LightBulbfromSpace/kld-go-learning/arrays_and_slices/sum - 100.0%

<details>
  <summary><code>func Sum(numbers []int) int</code></summary>

    Sum calculates the total from a slice of numbers.
</details>

### github.com/LightBulbfromSpace/kld-go-learning/arrays_and_slices/sumAll - 100.0%

<details>
  <summary><code>func SumAll(nums ...[]int) (sums []int)</code></summary>

    SumAll calculates the sums in all given slices.
</details>

### github.com/LightBulbfromSpace/kld-go-learning/arrays_and_slices/sumTails - 100.0%

<details>
  <summary><code>func SumTails(numsToSum ...[]int) []int</code></summary>

    SumTails calculates the sums of all but the first number given a collection
    of slices.
</details>

### github.com/LightBulbfromSpace/kld-go-learning/concurrency/faster_version_v2 - 100.0%

<details>
  <summary><code>func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool</code></summary>

    TYPES
    type WebsiteChecker func(string) bool
</details>

### github.com/LightBulbfromSpace/kld-go-learning/concurrency/slow_version_v1 - 100.0%

<details>
  <summary><code>func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool</code></summary>

    TYPES
    type WebsiteChecker func(string) bool
</details>

### github.com/LightBulbfromSpace/kld-go-learning/context/v1 - 100.0%

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

### github.com/LightBulbfromSpace/kld-go-learning/context/v2 - 100.0%

<details>
  <summary><code>func Server(store Store) http.HandlerFunc</code></summary>

    TYPES
    type Store interface {
    Fetch(ctx context.Context) (string, error)
    }
</details>

### github.com/LightBulbfromSpace/kld-go-learning/dependency_injection/greet - 50.0%

<details>
  <summary><code>func Greet(writer io.Writer, name string)</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/generics - 71.4%

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

### github.com/LightBulbfromSpace/kld-go-learning/hello - 92.9%

<details>
  <summary><code>func DefaultReceiver(language string) (receiver string)</code></summary>

</details>

<details>
  <summary><code>func GreetingPrefix(language string) (prefix string)</code></summary>

</details>

<details>
  <summary><code>func Hello(name string, language string) string</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/hellogo - 100.0%
Package hellogo contains first steps in language.
<details>
  <summary><code>func Hello() string</code></summary>

    Hello is first function.
</details>

### github.com/LightBulbfromSpace/kld-go-learning/integers - 100.0%

<details>
  <summary><code>func Add(a, b int) int</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/intro_to_propely_based_tests/v1 - 100.0%

<details>
  <summary><code>func ConvertToRoman(arabic int) string</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/intro_to_propely_based_tests/v2 - 93.3%

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

### github.com/LightBulbfromSpace/kld-go-learning/intro_to_propely_based_tests/v3_to_arabic_new - 96.6%
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

### github.com/LightBulbfromSpace/kld-go-learning/iteration - 100.0%

<details>
  <summary><code>func Repeat(str string, num int) string</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/maps - 83.3%
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
### github.com/LightBulbfromSpace/kld-go-learning/math/v1 - 100.0%
TYPES

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point
### github.com/LightBulbfromSpace/kld-go-learning/math/v2 - 100.0%

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

### github.com/LightBulbfromSpace/kld-go-learning/math/v3 - 100.0%

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

### github.com/LightBulbfromSpace/kld-go-learning/math/v4 - 100.0%

<details>
  <summary><code>func SVGWriter(w io.Writer, t time.Time)</code></summary>

    TYPES
    type Point struct {
    X float64
    Y float64
    }
    A Point represents a two dimensional Cartesian coordinate.
</details>

### github.com/LightBulbfromSpace/kld-go-learning/mocking/configurable_sleeper - 81.8%

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

### github.com/LightBulbfromSpace/kld-go-learning/mocking/countdown_v1 - 62.5%

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

### github.com/LightBulbfromSpace/kld-go-learning/pointers_and_errors - 83.3%
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
### github.com/LightBulbfromSpace/kld-go-learning/reading_files/v1 - 43.8%
TYPES

type Post struct {
	Title string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error)
### github.com/LightBulbfromSpace/kld-go-learning/reading_files/v2 - 84.2%
TYPES

type Post struct {
	Title string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error)
### github.com/LightBulbfromSpace/kld-go-learning/reading_files/v3 - 88.9%
TYPES

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPost(postFile io.Reader) (Post, error)
    NewPost parses file contents code

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error)
### github.com/LightBulbfromSpace/kld-go-learning/reading_files/v4 - 94.6%
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
    NewPost parses file contents code

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error)

type Separator string
### github.com/LightBulbfromSpace/kld-go-learning/reflection/v1 - 100.0%

### github.com/LightBulbfromSpace/kld-go-learning/reflection/v2 - 100.0%

### github.com/LightBulbfromSpace/kld-go-learning/reflection/v3 - 100.0%

### github.com/LightBulbfromSpace/kld-go-learning/revisiting_arrays_and_slices_with_generics - 74.4%

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

### github.com/LightBulbfromSpace/kld-go-learning/select/concurrency_check - 90.0%

<details>
  <summary><code>func ConfigurableWebsiteRacer(a, b string, timeout time.Duration) (winner string, err error)</code></summary>

</details>

<details>
  <summary><code>func WebsiteRacer(a, b string) (winner string, err error)</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/select/server_mocking - 87.5%

<details>
  <summary><code>func WebsiteRacer(a, b string) string</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/select/test_external_sites - 88.9%

<details>
  <summary><code>func WebsiteRacer(a, b string) (winner string)</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/structs_methods_interfaces - 100.0%
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
### github.com/LightBulbfromSpace/kld-go-learning/sync - 100.0%
TYPES

type Counter struct {
	// Has unexported fields.
}

func NewCounter() *Counter

func (c *Counter) Inc()

func (c *Counter) Value() int