# Types in TypeScript

### Boolean
```typescript
    let isLogin : boolean = false;
```

### Number
```typescript
    let num : number = 10;
```

### String
```typescript
    let name : string = "Junbeom";
```

### Array
```typescript
    let arr1 : number[] = [1, 2, 3];
    let arr2 : Array<number> = [10, 20, 30];
```

### String
```typescript
    let name : string = "Junbeom";
```

### Tuple
```typescript
    let arr: [string, number] = ['hi', 10];
```

### Enum
```typescript
    enum fruits { apple, banana, orange }
    let fruit: string = fruits[0];
    console.log(fruit) // 'apple'
```

### Any
```typescript
    let str: any = 'hi';
    let num: any = 10;
    let arr: any = ['a', 2, true];
```

### Void
```typescript
    function sayHi() : void {
    console.log('Hi')
    }
    // in variable (변수) can defind only null and undefined
    let a : void = null;
    let b : void = undefined;
```

### Never
```typescript
    
```