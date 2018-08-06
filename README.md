# chs

help to replace judging-code in golang

change

```
    if a == 0 {
        a = defaultValue
    }
    if s == "" {
        s = defaultStr
    }
```

to 

```
    a = IntOrDefault(a,defaultValue)
    s = StrOrDefault(s,defaultStr)
```

