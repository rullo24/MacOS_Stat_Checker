# pcstater

A native macOS menu bar performance monitor (CPU, memory, thermals, disk/network I/O).

## Architecture

- **UI**: SwiftUI (`MenuBarExtra`) — status item + popover, no AppKit/WebKit.
- **Logic**: Go — sampling loop, delta calculations, state management.
- **Metrics**: Go via cgo, calling Mach kernel APIs directly (`host_processor_info`, `host_statistics64`, etc.) instead of shelling out to `top`/`ps`.

```
SwiftUI App (Xcode)
      │  C ABI (bridging header)
      ▼
Go package, cgo-exported (//export GetSnapshot, FreeString, ...)
      │  cgo → Mach syscalls
      ▼
host_processor_info, host_statistics64, ...
```

The Go side compiles to a static library with:

```
go build -buildmode=c-archive -o libmetrics.a .
```

This produces `libmetrics.a` and an auto-generated `libmetrics.h`, which get linked into the Xcode project. Swift polls exported Go functions on a `Timer` (Go does not call back into Swift).

## Metrics

| Context | Metrics |
|---|---|
| Menu bar (always visible) | Aggregate CPU %, RAM used |
| Popover | Per-core CPU (P vs E cores), top 5 CPU processes, memory breakdown (wired/active/compressed/free), top 5 RAM processes, thermal state, disk/network throughput |

## Packaging

Bundled as `.app` with `LSUIElement = true` in `Info.plist` (menu bar only, no Dock icon).
