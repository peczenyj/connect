---
title: lfu
type: cache
status: stable
---

<!--
     THIS FILE IS AUTOGENERATED!

     To make changes please edit the corresponding source file under internal/impl/<provider>.
-->

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

Stores key/value pairs in a lfu in-memory cache. This cache is therefore reset every time the service restarts.


<Tabs defaultValue="common" values={[
  { label: 'Common', value: 'common', },
  { label: 'Advanced', value: 'advanced', },
]}>

<TabItem value="common">

```yml
# Common config fields, showing default values
label: ""
lfu:
  size: 1000
  samples: 100000
  init_values: {}
```

</TabItem>
<TabItem value="advanced">

```yml
# All config fields, showing default values
label: ""
lfu:
  size: 1000
  samples: 100000
  init_values: {}
  optimistic: false
```

</TabItem>
</Tabs>

This provides the lfu package which implements a fixed-size thread safe LFU cache.

It uses the package [`lfu`](github.com/vmihailenco/go-tinylfu)

The field init_values can be used to pre-populate the memory cache with any number of key/value pairs:

```yaml
cache_resources:
  - label: foocache
    lfu:
      size: 1024
      init_values:
        foo: bar
```

These values can be overridden during execution, at which point the configured TTL is respected as usual.

## Fields

### `size`

The cache maximum size (number of entries)


Type: `int`  
Default: `1000`  

### `samples`

The cache samples


Type: `int`  
Default: `100000`  

### `init_values`

A table of key/value pairs that should be present in the cache on initialization. This can be used to create static lookup tables.


Type: `object`  
Default: `{}`  

```yml
# Examples

init_values:
  Nickelback: "1995"
  Spice Girls: "1994"
  The Human League: "1977"
```

### `optimistic`

If true, we do not lock on read/write events. The lfu package is thread-safe, however the ADD operation is not atomic.


Type: `bool`  
Default: `false`  

