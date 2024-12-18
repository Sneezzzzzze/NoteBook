# Types in TypeScript PartTwo
[ReadOnly](#readonly) <br>
[Tuple](#tuple) <br>
[Any](#any)
### ReadOnly

```typescript
    type Player = {
        readonly name: string, // Can't Do anything
        age?:number
    }
    
    const numbers: readonly number[] = [1,2,3]
```

### Tuple
```typescript
    let arr: [string, number] = ['hi', 10];
```

### Any
```typescript
    let str: any = 'hi';
    let num: any = 10;
    let arr: any = ['a', 2, true];
```