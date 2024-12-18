# Types in TypeScript PartOne
[Boolean](#boolean) [Number](#number) [String](#string) [Array](#array) [Enum](#enum)

[Type Alias](#type-alias) [Union](#union---)

### 2 Way to defined type

```typescript
// Auto Type Check by TypeScript ( Recommend Way)
let a = "hello";

// defined type byselfs
let b: string = "bye";

let c = [1, 2, 3]; 

let d: number[] = []; // If wants to do like this 

const player = {
    name: "Junbeom"
}
play.name = 1; // Error because name have to be string
```

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


### Enum
```typescript
    enum fruits { apple, banana, orange }
    let fruit: string = fruits[0];
    console.log(fruit) // 'apple'
```

### Optional ( ? )
```typescript
  const player : {
    name: string,
    age?: number // number or undefined
  } = {
      name:"junbeom"
  }
```

### Type Alias
객체나 함수처럼 복잡한 타입을 매번 설정하는 것은 번거롭기 때문에 사용자가 직접 type 키워드를 사용하여 타입 별칭(Type Alias)을 정의할 수 있다.
```typescript
  const player1: {                  type Player = {
    name: string,                       name:string,
    age?: number                        age?:number
    } = {                           }
    name:"Junbeom"
  }                           ->    const player1 : Player = {
  const player2: {                      name:"Junbeom"
    name: string,                   }
    age?: number                    const player2 : Player = {
  } = {                                 name:"Junboom"
      name:"Junboom"                }
  }
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

### Set Return Type in Function
```typescript
    type Player = {
        name: string,
        age?:number
    }
    
    function playerMaker(name:string) : Player {
        return {
            name
        }
    }
    
    // Function Oneline
    const playerMakker = (name:string) : Player => ({name})
```

