This package serves and tracks unique uint16s.
It's intended use is for short-lived incremental id in a small data type.

Usage:
```
id, err := uuint16.Rent()

...

uuint16.Return(id)
```

Errors:
Only one error is available: `uuint16.ErrorNoneAvailable`
it happens when all 65535 values have been rented.

This package uses mutex for thread safety :thumbsup:
