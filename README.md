interfaceGUID
=============

interfaceGUID exposes a single function `CalculateGUID` which calculates a
unique deterministic value based on the type provided. It is to be used when
communicating with remote services to quickly assert common knowledge before
starting to communicate, for example via `encoding/gob`.


Example
-------

    type MyService interface {
      ...
    }

    guid := interfaceGUID.CalculateGUID(reflect.TypeOf((*MyService)(nil)).Elem())
    // Then compare the string with the remote service.
