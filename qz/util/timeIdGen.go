package util

/**
* run in playground: https://go.dev/play/p/vnCZTqOjZp9
*
* Fancy ID generator that creates 20-character string identifiers with the following properties:
*
* 1. They're based on timestamp so that they sort *after* any existing ids.
* 2. They contain 72-bits of random data after the timestamp so that IDs won't collide with other clients' IDs.
* 3. They sort *lexicographically* (so the timestamp is converted to characters that will sort properly).
* 4. They're monotonically increasing. Even if you generate more than one in the same timestamp, the
* latter ones will sort after the former ones. We do this by using the previous random bits
* but "incrementing" them by 1 (only in the case of a timestamp collision).
 */

import (
	crand "crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// Timestamp of last push, used to prevent local collisions if you push twice in one ms.
var lastPushTime int64

// We generate 72-bits of randomness which get turned into 12 characters and appended to the
// timestamp to prevent collisions with other clients. We store the last characters we
// generated because in the event of a collision, we'll use those same characters except
// "incremented" by one.
var lastRandChars []int8

// Modeled after base64 web-safe chars, but ordered by ASCII.
const PUSH_CHARS string = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"

func GeneratePushID() string {
	now := time.Now().UTC().UnixNano() / 1000000
	duplicateTime := now == lastPushTime
	lastPushTime = now

	timeStampChars := make([]string, 8)
	for i := 7; i >= 0; i-- {
		pcIndex := int64(math.Mod(float64(now), 64.0))
		timeStampChars[i] = string(PUSH_CHARS[pcIndex])
		now = int64(math.Floor(float64(now) / 64.0))
	}

	if now != 0 {
		panic("We should have converted the entire timestamp.")
	}

	id := strings.Join(timeStampChars, "")

	if !duplicateTime {
		for i := 0; i < 12; i++ {
			lastRandChars[i] = int8(math.Floor(rand.Float64() * 64.0))
		}
	} else {
		var i int
		for i = 11; i >= 0 && lastRandChars[i] == 63; i-- {
			lastRandChars[i] = 0
		}

		lastRandChars[i]++
	}

	for i := 0; i < 12; i++ {
		id = fmt.Sprintf("%s%s", id, string(PUSH_CHARS[lastRandChars[i]]))
	}

	if len(id) != 20 {
		panic("Length should be 20")
	}

	return id
}

var gb64 = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"

func GenB32(m string) (string, error) {
	encb64 := base64.NewEncoding(gb64)
	dec, err := encb64.DecodeString(m)
	if err != nil {
		return "", err
	}
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(dec), nil
}

const ASIN_CHARS string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
const asinLen = 10

// GenASIN generates a 10 digit Amazon Serial Identifier Number
func GenASIN() string {
	randstr := make([]byte, asinLen)
	b := make([]byte, asinLen)
	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}
	randstr[0] = ASIN_CHARS[b[0]%26] // 1st byte is always an alphabet
	for i := 1; i < asinLen; i++ {
		randstr[i] = ASIN_CHARS[int(b[i])%len(ASIN_CHARS)]
	}
	return string(randstr)
}

/*
 func main() {
	 lastRandChars = make([]int8, 12, 12)

	 //fmt.Println("Fast:")
	 for i := 0; i < 10; i++ {
			 time.Sleep(1 * time.Second / 2)
			 p := generatePushID()
		 fmt.Println(p)
		 b32, err := genB32(p)
		 if err != nil {
		   fmt.Println(err)
		   return
		 }
		 fmt.Println("did:randid:"+b32)
	 }

 }
*/
