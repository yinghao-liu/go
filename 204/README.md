# 204 gin

gin bind

## param oneof

### 方式一

```json
{
  "type": "a",
  "a": {
    "aa": "1",
    "ab": "2"
  }
}
```

或

```json
{
  "type": "b",
  "b": {
    "ba": "1",
    "bb": "2"
  }
}
```

### 方式二

```json
{
  "type": "a",
  "config": {
    "aa": "1",
    "ab": "2"
  }
}
```

或

```json
{
  "type": "b",
  "config": {
    "ba": "1",
    "bb": "2"
  }
}
```

## param anyof

```json
{
  "config": {
    "a": {
      "aa": "1",
      "ab": "2"
    },
    "b": {
      "ba": "1",
      "bb": "2"
    }
  }
}
```
