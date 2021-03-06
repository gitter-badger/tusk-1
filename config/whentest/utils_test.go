package whentest

import "testing"

func TestTrue(t *testing.T) {
	if err := True.Validate(nil); err != nil {
		t.Errorf(
			"whentest.True did not pass validation. Unexpected err: %s", err,
		)
	}
}

func TestFalse(t *testing.T) {
	if err := False.Validate(nil); err == nil {
		t.Errorf(
			"whentest.False passed validation but should have errored",
		)
	}
}
