# protocat

* protocat concatenates multiple proto files into 1 file.
  - files must be in a single package.

## Installation

```console
go get github.com/syumai/protocat/cmd/protocat
```

## Usage

```console
# import paths must be specified by `-I` flag.
protocat -I=example,third_party/protobuf/src example-a.proto example-b.proto
```

**example/example-a.proto**

```proto
syntax = "proto3";

package example;

import "example-b.proto";

message A {
    B b = 1;
}
```

**example/example-b.proto**

```proto
syntax = "proto3";

package example;

import "google/protobuf/any.proto";

message B {
    google.protobuf.Any any = 1;
}
```

**output**

```proto
syntax = "proto3";

package example;

import "google/protobuf/any.proto";

message A {
  B b = 1;
}

message B {
  google.protobuf.Any any = 1;
}
```

## License

MIT

## Author

syumai