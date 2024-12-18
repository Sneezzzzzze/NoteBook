# Types in TypeScript PartThree
[Unknown](#unknown)<br>
[Void](#void)<br>
[Never](#never)

### Unknown
```typescript
    let a:unknown;
    let b = a + 1; // not allow
    if (typeof a === 'number') {
        let b = a + 1
    }
    if (typeof a === 'string') {
        let b = a.toUpperCase();
    }
```

### Void
```typescript
    // Not Return Anythings
    function sayHi() : void {
        console.log('Hi')
    }
    // same
    function sayHi() {  
        console.log('Hi')
    }
    // in variable (변수) can defined only null and undefined
    let a : void = null;
    let b : void = undefined;
```

### Never
`never`는 **"존재하지 않는다"/ "절대 실행될 수 없다"**는 개념을 나타냅니다. 예외를 던지거나 무한 루프에 빠지는 함수 또는 발생할 수 없는 코드를 다룰 때 사용됩니다.

`never`와 `void`의 차이점
- `void`: 값이 반환되지 않을 수는 있지만, 함수가 정상적으로 종료될 경우 사용됩니다.
    - 예: `console.log` 함수의 반환 타입은 `void`입니다.

- `never`: 아예 값을 반환하지 않으며, 함수가 예외를 던지거나 무한 루프에 빠지는 경우 사용됩니다.
```typescript
    // 항상 오류 발생
    function invalid(message:string): never {
        throw new Error(message);
    }
    // 무한 루프
    function infiniteAnimate(): never {
        while ( true ) { infiniteAnimate(); }
    }
    
    function hello(name:string | number) {
        if (typeof name === 'string') {
            name
        } else if (typeof name === 'number') {
            name
        } else {
            name
        }
    }
```