package logstorage

import (
	"reflect"
	"testing"
)

func TestParseStatsCountEmptySuccess(t *testing.T) {
	f := func(pipeStr string) {
		t.Helper()
		expectParseStatsFuncSuccess(t, pipeStr)
	}

	f(`count_empty(*)`)
	f(`count_empty(a)`)
	f(`count_empty(a, b)`)
}

func TestParseStatsCountEmptyFailure(t *testing.T) {
	f := func(pipeStr string) {
		t.Helper()
		expectParseStatsFuncFailure(t, pipeStr)
	}

	f(`count_empty`)
	f(`count_empty(a b)`)
	f(`count_empty(x) y`)
}

func TestStatsCountEmpty(t *testing.T) {
	f := func(pipeStr string, rows, rowsExpected [][]Field) {
		t.Helper()
		expectPipeResults(t, pipeStr, rows, rowsExpected)
	}

	f("stats count_empty(*) as x", [][]Field{
		{
			{"_msg", `abc`},
			{"a", `2`},
			{"b", `3`},
		},
		{
			{"_msg", `def`},
			{"a", `1`},
		},
		{},
		{
			{"a", `3`},
			{"b", `54`},
		},
	}, [][]Field{
		{
			{"x", "1"},
		},
	})

	f("stats count_empty(b) as x", [][]Field{
		{
			{"_msg", `abc`},
			{"a", `2`},
			{"b", `3`},
		},
		{
			{"_msg", `def`},
			{"a", `1`},
		},
		{},
		{
			{"a", `3`},
			{"b", `54`},
		},
	}, [][]Field{
		{
			{"x", "2"},
		},
	})

	f("stats count_empty(a, b) as x", [][]Field{
		{
			{"_msg", `abc`},
			{"a", `2`},
			{"b", `3`},
		},
		{
			{"_msg", `def`},
			{"a", `1`},
		},
		{},
		{
			{"aa", `3`},
			{"bb", `54`},
		},
	}, [][]Field{
		{
			{"x", "2"},
		},
	})

	f("stats count_empty(c) as x", [][]Field{
		{
			{"_msg", `abc`},
			{"a", `2`},
			{"b", `3`},
		},
		{
			{"_msg", `def`},
			{"a", `1`},
		},
		{
			{"a", `3`},
			{"b", `54`},
		},
	}, [][]Field{
		{
			{"x", "3"},
		},
	})

	f("stats count_empty(a) if (b:*) as x", [][]Field{
		{
			{"_msg", `abc`},
			{"a", `2`},
			{"b", `3`},
		},
		{
			{"_msg", `def`},
			{"a", `1`},
		},
		{
			{"b", `54`},
		},
	}, [][]Field{
		{
			{"x", "1"},
		},
	})

	f("stats by (a) count_empty(b) as x", [][]Field{
		{
			{"_msg", `abc`},
			{"a", `1`},
			{"b", `3`},
		},
		{
			{"_msg", `def`},
			{"a", `1`},
		},
		{
			{"a", `3`},
			{"b", `5`},
		},
		{
			{"a", `3`},
			{"b", `7`},
		},
	}, [][]Field{
		{
			{"a", "1"},
			{"x", "1"},
		},
		{
			{"a", "3"},
			{"x", "0"},
		},
	})

	f("stats by (a) count_empty(b) if (!c:foo) as x", [][]Field{
		{
			{"_msg", `abc`},
			{"a", `1`},
			{"b", `3`},
		},
		{
			{"_msg", `def`},
			{"a", `1`},
			{"c", "foo"},
		},
		{
			{"a", `3`},
			{"b", `5`},
			{"c", "bar"},
		},
		{
			{"a", `3`},
		},
	}, [][]Field{
		{
			{"a", "1"},
			{"x", "0"},
		},
		{
			{"a", "3"},
			{"x", "1"},
		},
	})

	f("stats by (a) count_empty(*) as x", [][]Field{
		{
			{"_msg", `abc`},
			{"a", `1`},
			{"b", `3`},
		},
		{
			{"_msg", `def`},
			{"a", `1`},
			{"c", "3"},
		},
		{},
		{
			{"a", `3`},
			{"b", `5`},
		},
		{
			{"a", `3`},
			{"b", `7`},
		},
	}, [][]Field{
		{
			{"a", ""},
			{"x", "1"},
		},
		{
			{"a", "1"},
			{"x", "0"},
		},
		{
			{"a", "3"},
			{"x", "0"},
		},
	})

	f("stats by (a) count_empty(c) as x", [][]Field{
		{
			{"_msg", `abc`},
			{"a", `1`},
			{"b", `3`},
		},
		{
			{"_msg", `def`},
			{"a", `1`},
		},
		{
			{"a", `3`},
			{"c", `5`},
		},
		{
			{"a", `3`},
			{"b", `7`},
		},
	}, [][]Field{
		{
			{"a", "1"},
			{"x", "2"},
		},
		{
			{"a", "3"},
			{"x", "1"},
		},
	})

	f("stats by (a) count_empty(a, b, c) as x", [][]Field{
		{
			{"_msg", `abc`},
			{"a", `1`},
			{"b", `3`},
		},
		{
			{"_msg", `def`},
			{"a", `1`},
			{"c", "3"},
		},
		{
			{"a", `3`},
			{"b", `5`},
		},
		{
			{"a", `3`},
			{"b", `7`},
		},
	}, [][]Field{
		{
			{"a", "1"},
			{"x", "0"},
		},
		{
			{"a", "3"},
			{"x", "0"},
		},
	})

	f("stats by (a, b) count_empty(a) as x", [][]Field{
		{
			{"_msg", `abc`},
			{"a", `1`},
			{"b", `3`},
		},
		{
			{"_msg", `def`},
			{"a", `1`},
			{"c", "3"},
		},
		{
			{"c", `3`},
			{"b", `5`},
		},
	}, [][]Field{
		{
			{"a", "1"},
			{"b", "3"},
			{"x", "0"},
		},
		{
			{"a", "1"},
			{"b", ""},
			{"x", "0"},
		},
		{
			{"a", ""},
			{"b", "5"},
			{"x", "1"},
		},
	})
}

func TestStatsCountEmpty_ExportImportState(t *testing.T) {
	f := func(scp *statsCountEmptyProcessor, dataLenExpected, stateSizeExpected int) {
		t.Helper()

		data := scp.exportState(nil, nil)
		dataLen := len(data)
		if dataLen != dataLenExpected {
			t.Fatalf("unexpected dataLen; got %d; want %d", dataLen, dataLenExpected)
		}

		var scp2 statsCountEmptyProcessor
		stateSize, err := scp2.importState(data, nil)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if stateSize != stateSizeExpected {
			t.Fatalf("unexpected state size; got %d bytes; want %d bytes", stateSize, stateSizeExpected)
		}

		if !reflect.DeepEqual(scp, &scp2) {
			t.Fatalf("unexpected state imported; got %#v; want %#v", &scp2, scp)
		}
	}

	var scp statsCountEmptyProcessor

	f(&scp, 1, 0)

	scp = statsCountEmptyProcessor{
		rowsCount: 234,
	}
	f(&scp, 2, 0)
}
