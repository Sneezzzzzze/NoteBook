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
```

### User defined Type
객체나 함수처럼 복잡한 타입을 매번 설정하는 것은 번거롭기 때문에 사용자가 직접 type 키워드를 사용하여 타입 별칭(Type Alias)을 정의할 수 있다.
```typescript
    type operation = {
    data: number[],
    output:(num:number)=>number[]
    };
    // 사용자 정의 타입 operation 적용 예시
    let sum:operation = {
        data: [10, 30, 60],
        output(num){
            return this.data.map(n=>n+num);
        }
    };
    let multiply:operation = {
        data: [110, 230, 870, 231],
        output(num){
            return this.data.map(n=>n*num);
        }
    };
```

### Union ( | )
Use when allow multiple type <br>
두 가지 이상의 타입을 모두 허용한는 경우에 사용한다.
```typescript
    function getAge(age: number | string): string {
        if (typeof age === 'number') {
            return age.toString();
        } else {
            return age;
        }
    }
```
