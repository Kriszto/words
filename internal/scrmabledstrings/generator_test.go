package scrmabledstrings

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator_scanTestCase(t *testing.T) {
	r := strings.NewReader(`12
zf fvtfzfndlxlfr vnbrfr vvddzzhr dptdnjnjbx vznthzpfjdt xtlnlvdjfjdj rjbndrd xdjz pnjbhfrlh dnh hvtnb
l o 1000 975016853 972311008 925928190 999999236`)
	scanner := bufio.NewScanner(r)
	g := NewGenerator()
	expectedWords := []string{"zf", "fvtfzfndlxlfr", "vnbrfr", "vvddzzhr", "dptdnjnjbx", "vznthzpfjdt", "xtlnlvdjfjdj", "rjbndrd", "xdjz", "pnjbhfrlh", "dnh", "hvtnb"} //nolint:lll
	expectedSeeds := []string{"l", "o", "1000", "975016853", "972311008", "925928190", "999999236"}
	words, seeds := g.scanTestCase(scanner)
	assert.Equal(t, expectedWords, words)
	assert.Equal(t, expectedSeeds, seeds)
}

func TestGenerator_convertParams(t *testing.T) {
	g := NewGenerator()
	params := []string{"l", "o", "1000", "975016853", "972311008", "925928190", "999999236"}
	gi := g.convertParams(params)
	assert.Equal(t, 'l', gi.S1)
	assert.Equal(t, 'o', gi.S2)
	assert.EqualValues(t, 1000, gi.N)
	assert.EqualValues(t, 975016853, gi.A)
	assert.EqualValues(t, 972311008, gi.B)
	assert.EqualValues(t, 925928190, gi.C)
	assert.EqualValues(t, 999999236, gi.D)
}

func TestGenerator_GenerateInput(t *testing.T) {
	g := NewGenerator()
	s := g.GenerateInput(GeneratorInput{
		S1: 'l',
		S2: 'o',
		N:  1000,
		A:  975016853,
		B:  972311008,
		C:  925928190,
		D:  999999236,
	})
	expected := "lonzrbfrrbpjrtlzfvjzdflzhjthnjdhzbntbjxfbfbpzvvlfxbvxxjvnnphzvdhjxrdnlxdjrbxpxxvzpvljbtdtffnzbnxfbxbtpzpjjrzhhrlbjbnjztnfpxrbtfxhvdnxznzjjfrjxphdvjrdxdnhjhrvnxbtzbpzlvfvvlhjnxbzhpfrjrdrzlvtnrtbrlzvlbznzdvnprpfjxpztbjxxrxtrjrxdvprnxzfnzttxtnrjhvzlhvfjvpdrrrfxvxfftjtrzvtzftxfznxdtjhlbthxzlpftfrvdrlvfvfhbzttdxxdpvtfttxnbdtrlxfdtnjnpbjdxrphhtxdnxdjzddfhhfzzxrdxvjfdnjdlltjvttzvbrptpvxlnjzvzdfdzhjhpdldzjdbdjxltrjtdxvdnhjftpzztxpnrbjfhlhxvhbdztldvztxlljnbvrdxxvpjrtfrvxdjfbtrxflbztxjlrvhxtzzbfxdfptvflzxxnppfvxpbltvjlddjtrjflnrblvdnttbjvzbzbxdjlrpvhdbnzjhdtjplpfzrvttdrtrtrxzppltfnbdpbtxjvtrjdjnprtvzvzxbtzbvztxjbdhjhxndrzplrnnjdnpptvrdpnhbxjxztpxbnldbjxhnnpjtvpfrhznpffphrhjzdfnffllfnlzffxdtrndlhjrnflzldhdftlvjxfbllfffztdzlnfzrvthzhjvvhdzdzrhrzbhvtnbdxhpnrzjnrzzxpnzprfbpnjxbzvdbrfrnhhdzfbphfljjvlfjjllflfftfrzdrhbpbrfzdbrzhxhxxxdpbjvhpxrvjtfjlnjrtbtxxhrdjrbjxdfrjrzfftdplzlptzrltzpddffnvlxzbvfbxprlfbpbdbdbfxdnttbtdpzfthnhpddlrvrrjnrxnxrjbndrdhzbfrdnbbzzdhzxdjzjzlpldlnfvdpptptxrrtrxblnjfrjzfbnhrlvjl" //nolint:lll
	assert.Equal(t, expected, s)
}
