package once_test

import (
	"errors"
	"testing"

	"github.com/hsson/once"
)

func TestOnce(t *testing.T) {
	o := once.Once{}
	count := 0

	for i := 0; i < 100; i++ {
		go o.Do(func() {
			count++
		})
	}

	if count != 1 {
		t.Errorf("unexpected count, got %v want %v", count, 1)
	}
}

func TestOnceError(t *testing.T) {
	o := once.Error{}
	count := 0
	expected := errors.New("some error")

	for i := 0; i < 100; i++ {
		err := o.Do(func() error {
			count++
			return expected
		})
		if err != expected {
			t.Errorf("%d: got unexpected err: %v", i, err)
		}
	}

	if count != 1 {
		t.Errorf("unexpected count, got %v want %v", count, 1)
	}
}

func TestOnceValue(t *testing.T) {
	o := once.Value{}
	count := 0
	expected := "some value"

	for i := 0; i < 100; i++ {
		value := o.Do(func() interface{} {
			count++
			return expected
		})
		if value != expected {
			t.Errorf("%d: got unexpected value: %v", i, value)
		}
	}

	if count != 1 {
		t.Errorf("unexpected count, got %v want %v", count, 1)
	}
}

func TestOnceValueError(t *testing.T) {
	o := once.ValueError{}
	count := 0
	expectedValue := "some value"
	expectedErr := errors.New("some error")

	for i := 0; i < 100; i++ {
		value, err := o.Do(func() (interface{}, error) {
			count++
			return expectedValue, expectedErr
		})
		if value != expectedValue {
			t.Errorf("%d: got unexpected value: %v", i, value)
		}
		if err != expectedErr {
			t.Errorf("%d: got unexpected error: %v", i, err)
		}
	}

	if count != 1 {
		t.Errorf("unexpected count, got %v want %v", count, 1)
	}
}
