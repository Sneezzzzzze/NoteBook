### Interfaces

```typescript
    type Team = "red" | "blue" | "green"
    type Health = 100 | 200 | 300
    
    // type Player = {
    //     name :string,
    //     team :Team,
    //     health :Health;
    // }
    // Same as above
    //
    interface Player {
        name :string,
        team :Team,
        health :Health;
    }
    const player1: Player = {
        name: "Tom",
        team: "red",
        health: 100
    }
    
    interface User {
        name :string
    }
    
    interface Player extends User {
    }
```

Can Stack Interfaces

```typescript
    interface User {
        name :string
    }
    interface User {
        age :number
    }
    
    const user: User = {
        name: "Tom",
        age: 20
    }
```

```typescript
    abstract class User {
        constructor(
            protected firstName :string,
            protected lastName :string
        )   {}
        abstract sayHi(name :string) :string
        abstract fullName() :string
    }

    class Player extends User {
        fullName() :string {
            return `${this.firstName} ${this.lastName}`
        }
        sayHi(name :string) :string {
            return `Hello ${name}, My name is ${this.fullName()}`
        }
    }
    
    // Change to Interface
    // Difference between Interface and Abstract class
    // Interface is a contract that a class can implement
    // Abstract class is a contract that a class can extend
    
    interface User {
        firstName :string
        lastName :string
        sayHi(name :string) :string
        fullName() :string
    }
    
    interface Human {
        health :number
    }
    
    class Player implements User, Human {
        constructor(
            // Can't use protected, private keyword in Interface
            public firstName :string,
            public lastName :string,
            public health :number
        )   {}
        fullName() :string {
            return `${this.firstName} ${this.lastName}`
        }
        sayHi(name :string) :string {
            return `Hello ${name}, My name is ${this.fullName()}`
        }
    }
```

Example Interfaces x Classes Generics

```typescript
    interface SStorage<T> {
        [key :string] :T;
    }
    class LocalStorage<T> {
        private storage :SStorage<T> = {};
        set(key :string, value :T) {
            this.storage[key] = value
        }
        remove(key :string) {
            delete this.storage[key]
        }
        get(key :string) :T {
            return this.storage[key]
        }
        clear() {
            this.storage = {}
        }
    }
    
    const stringStorage = new LocalStorage<string>();
    stringStorage.get('name');
```

