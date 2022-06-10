package proof

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestDivHash(t *testing.T) {

	expected := []string{
		"526f5c1eb6df242199815cc2cbc8fa153e6989b7f81ae3999d22111c85b964b5",
		"8779ac5320f8c8d90847af47d1dcb2366bb042f2ed8771743f985df02a206bd2",
		"bc9eab0fbfe06c67eb593b0b4cd154366cd9939b607a273ed4e8f1ce75f039d0",
		"8b1bcf1515fea74b5d1c4d2011b84d59bce37b8c7069607841a6f0cb550ec55a",
		"02b702f93e07f627f28937a7c599af22e926e2ed5246b58d601356bc66a77d72",
		"af96a09355ac4565b96adfd2d261b785b571edccc5adf1f671cf88ee566a997c",
		"d8c39dceba14049ea8d24625a1acc1bfcb2ab762fbfebef6620f16ee27fe9afe",
		"b85fa8c0be0acb5c2833fb87520229b5ff2e35c976f303930033a2f44a72d373",
		"2b089828dcbd8fb1b3bd039f660ece6a73260bafe4d5f9580f402499f85c6e02",
		"d8a0f7903dafed7c08e17d0d37e7e4b8410404c89ad4c6532a59b57effe42131",
		"c28479415db2770820789a238edcfe1762d654887dbd51a4ee479715646a13c9",
		"b6abab25b9d3371ce285e23221f57b88acb01f8fdb1765e44a3f452a9765ddb8",
		"639d26ccd6524b748ccb828503b1b1d48cb5a3255fefb2bbb34c241b2c328417",
		"5ac2821c4acc949b0b88e5b56e20e97b031d5f6532f08509cbb4cca21cb31f8d",
		"43e8cb865fc99741153e152aeb3a82413b709707af1941049e264f45ed0c96b7",
		"3603d6f8a9f90772c0a84fd592b870567fa3f33f479f9ca57e8a23524565e20e",
		"e195ae36be365f91f2aa9ab34f9c94e2859c84bd02a232ede37ee8e5f6455151",
		"df8905ff06291574fddf4aa1f648c584f554de3c98edaa15d6638e920af8e087",
		"4d9e7bebb44ec86b9ffe9750683c49e8c3505a292dbdae085037a2d595760f48",
		"7bc3fe9715f987272b843242822754bea912d817c528f9cbd5c86e325c59e13d",
		"010b377631fe487dd17a6cdae069e1bb1f10d2d7c110c491ea743b54b911ea2b",
		"c0c0464af7c239e450352de4e377efaed9b86be4a72a9fee30f5658339772db7",
		"b4b20f92993bf543eaa2a030c6bda08e6160c857d767beb4805e04f51f85efe2",
		"a315601df6766abebd802742b65e019b430074d8a03e914ec3c8e277497e54f2",
		"7f68eef58e919f34858d87f6e22fb25a4a9e633a28026fadf99898363909a72c",
		"e59e0c8449d1464c65d7f853d9daf9f7748cf353f6bcd7ccc69fc4128c174e3b",
		"58f50ec5cc167b5ffd20ea2cad3bcf39ef89cf01d8ea3439ee74092561d0eda1",
		"a62eba8572ef950729a5491c8ca3d8b8d19d1dd34296c0710f1aa0cbdb5f7f5f",
		"aba9cd379b0287508c97cf64189bb51ddb8b6341c520045981dd986db3598feb",
		"9c8366d52db5c1a507200fb728611b1ea9429a4ccd5c8078e71a9e24c363123e",
		"c81c43cad276ae2521c822305d3d1e7badf414b2f85bccdb722db0d176837ea6",
		"6e217044df9b43d4028221bd666e569f4b092c51bee603da3e6b09aff160d92f",
	}
	o := "\n\texpected := []string{\n"
	// Standard block size is 138 bytes for a validator block. If another hash is
	// added to the block for proposals or other linked content such as a merkle
	// root for a transaction payload, the result is about 30% more processing time
	// per operation.
	empty := Blake3([]byte{})
	empty = append(empty, empty...)
	empty = append(empty, empty...)
	empty = append(empty, empty[:10]...)
	for i := 0; i < 32; i++ {
		hash := DivHash4(empty[:138])
		empty = hash
		empty = append(empty, empty...)
		empty = append(empty, empty...)
		empty = append(empty, empty[:10]...)
		expect, err := hex.DecodeString(expected[i])
		if err != nil {
			t.Fatalf("error decoding hex string '%s': %v", expected[i], err)
		}
		if string(hash) != string(expect) {
			t.Fatalf("")
		}
		o += fmt.Sprintf("\t\t\"%s\",\n", hex.EncodeToString(hash))
	}
	o += "\t}\n"
	t.Log("generated:\n", o)
}

func benchmarkDivHash(reps int, b *testing.B) {

	// Maximum block length is 138 bytes for not yet implemented proposal blocks
	// this will test the block length defined in pkg/blocks/types.go for a
	// validator token. When proposal blocks are implemented they have a longer time
	// per operation.
	empty := Blake3([]byte{})
	empty = append(empty, empty...)
	empty = append(empty, empty[:32]...)
	empty = append(empty, empty[:10]...)

	for i := 0; i < b.N; i++ {

		hash := DivHash(empty, reps)
		empty = hash
		empty = append(empty, empty...)
		empty = append(empty, empty[:32]...)
		empty = append(empty, empty[:10]...)
	}
}

// Operation time for DivHash is proportional to input size. Thus, for a given
// data size a repetition count is chosen based on the benchmark below:
//
// $ go test -bench=. ./pkg/proof/.
// goos: linux
// goarch: amd64
// pkg: github.com/cybriq/kismet/pkg/proof
// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkDivHash1-16               10000            226545 ns/op
// BenchmarkDivHash2-16                 763           1815126 ns/op
// BenchmarkDivHash3-16                  82          15908234 ns/op
// BenchmarkDivHash4-16                   7         148824596 ns/op
// BenchmarkDivHash5-16                   1        1325919974 ns/op
// PASS
// ok      github.com/cybriq/kismet/pkg/proof      14.759s
//
// In order to render a reasonably integer number per second operations, 4
// repetitions will be selected.

func BenchmarkDivHash1(b *testing.B) { benchmarkDivHash(1, b) }
func BenchmarkDivHash2(b *testing.B) { benchmarkDivHash(2, b) }
func BenchmarkDivHash3(b *testing.B) { benchmarkDivHash(3, b) }
func BenchmarkDivHash4(b *testing.B) { benchmarkDivHash(4, b) }
func BenchmarkDivHash5(b *testing.B) { benchmarkDivHash(5, b) }
