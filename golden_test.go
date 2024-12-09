package golden

import "testing"

func TestGolden(t *testing.T) {
	Assert(t, []byte("Hello, Golden!\n"))
	t.Run("nested", func(t *testing.T) {
		Assert(t, []byte("Hello again, Golden!\n"))
		t.Run("doubly_nested", func(t *testing.T) {
			Assert(t, []byte("Hello yet again, Golden!\n"))
		})
		t.Run("with_file_ext.json", func(t *testing.T) {
			Assert(t, []byte(`{"msg": "Hello with file ext"}
			`))
		})
	})
}
